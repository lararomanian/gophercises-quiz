package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type questions struct {
	question []SingleQuestion
}

type SingleQuestion struct {
	question string
	answer   string
}

type guesses struct {
	correct   int
	incorrect int
	skipped   int
}

var TotalQuestions questions
var TotalGuesses guesses
var input_time int

func importCSV(filename string) {

	if filename == "" {
		filename = "./problems.csv"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	createQuiz(lines)
}

func createQuiz(questions [][]string) {

	var quiz []SingleQuestion

	for _, value := range questions {
		if len(value) == 2 {
			quiz = append(quiz, SingleQuestion{
				question: value[0],
				answer:   value[1],
			})
		}
	}

	TotalQuestions.question = quiz

}

func checkResult(c chan int) {

	line := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the answers Based on the questions")
	fmt.Println("Press enter to submit answer")
	fmt.Println("---------------------------------------------------------------------------------")

	fmt.Println("Enter READY to start the quiz")
	fmt.Println("---------------------------------------------------------------------------------")

	for _, item := range TotalQuestions.question {

		select {
		case <-c:
			printFinalResult()
			fmt.Println("DC : Timer has Expired")
			os.Exit(0)

		default:

			fmt.Print("->")
			fmt.Printf("%s ", item.question)

			answer, _, err := line.ReadLine()

			if err != nil {
				log.Fatal(err)
			}

			// fmt.Printf("The answer is %s \n", strings.ToLower(strings.TrimSpace(item.answer)))
			// fmt.Printf("Your answer was %s \n", strings.ToLower(strings.TrimSpace(string(answer))))

			matchQuestionToAnswer(strings.ToLower(strings.TrimSpace(item.answer)), strings.ToLower(strings.TrimSpace(string(answer))))
			fmt.Println()
		}

	}

	c <- 1

}

func matchQuestionToAnswer(question string, answer string) {

	fmt.Println(question)
	fmt.Println(answer)

	if question == answer {
		TotalGuesses.correct = TotalGuesses.correct + 1
	} else {
		TotalGuesses.incorrect = TotalGuesses.incorrect + 1
	}

	fmt.Println()
	fmt.Printf("%#v", TotalGuesses)
}

func printFinalResult() {
	total_questions := len(TotalQuestions.question)
	total_correct_answers := TotalGuesses.correct

	fmt.Printf("\n You've scored %d out of %d\n", total_correct_answers, total_questions)
}

func shuffleTotalQuestions() {

	n := len(TotalQuestions.question)

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		TotalQuestions.question[i], TotalQuestions.question[j] = TotalQuestions.question[j], TotalQuestions.question[i]
	}
}

func main() {
	fmt.Println("Reading CSV Data from ./problems.csv")
	TotalGuesses = guesses{correct: 0, incorrect: 0, skipped: 0}
	filename := ""

	if len(os.Args) >= 2  && os.Args[1] == "--file" {
		filename = os.Args[2]
		importCSV("./" + filename + ".csv")
	} else {
        importCSV("")
    }

	if len(os.Args) > 3 && os.Args[3] == "--time" {
		input_time, _ = strconv.Atoi(os.Args[4])
	} else {
		// fmt.Print("setting def inputtime")
		input_time = 30
	}

	if len(os.Args) == 6 && os.Args[5] == "--shuffle" {
		shuffleTotalQuestions()
	}

	timerChannel := make(chan int)
	go func() {
		timer := time.NewTimer(time.Duration(input_time) * time.Second)
		<-timer.C
		timerChannel <- 1
	}()

	go checkResult(timerChannel)

	<-timerChannel
	printFinalResult()
}
