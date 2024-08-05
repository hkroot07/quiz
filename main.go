package main

import (
	"fmt"
	"os"
	"quiz/questions"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't load questions: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(questions)
}
