#  GOPHERCISES-QUIZ


Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```csv
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.
At the end of the quiz the program should output the total number of questions correct and how many questions there were in total.
Questions given invalid answers are considered incorrect.

# FEATURES:
CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV.


Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.
- [x] Create a program that will read in a quiz provided via a CSV file
- [x] Output the total number of questions correct and how many questions there were
- [x] Keep track of how many questions are right and how many are incorrect
- [x] Users should be asked to press enter (or some other key) before the timer starts
- [x] Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint: Check out the strings package.
- [x] Add an option (a new flag) to shuffle the quiz order each time it is run.

# RUNNING:

use go run main.go --file problems --time 2 --shuffle in the order to have the selected flags as i couldn't make them dynamic (yet)

you can use run main.go , it runs the default problems.csv at a timer for 30 secs using non shuffled index


to read custom files

```bash
go run main.go --file "filename (without extension)" to read custom files
```

to have a timer of said time

```bash
go run main.go --file -filename --time "time in secs" to have a timer of said time
```


to shuffle the questions

```bash
go run main.go --shuffle to shuffle the questions
```
##
