package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer string
}

func main() {
	var helpFlag = flag.Bool("h", false, "help")
	var questionCSVPathFlag = flag.String("csv", "questions.csv", "Specify the question csv filename.")
	var timerFlag = flag.Int("timer", 30, "Quiz timer")
	flag.Parse();

	var quizArr []Quiz

	fmt.Printf("%v\n", *helpFlag)

	if *helpFlag {
		// display the help text
		flag.Usage();
	} else {

		// try open the file
		file, err := os.Open(*questionCSVPathFlag);
		if err != nil {
			log.Fatal(err);
		}

		defer file.Close();

		csvReader := csv.NewReader(file)
		// For custom delimiter. comma is default
		csvReader.Comma = ','

		for {
			line, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Printf("%s - %s\n", line[0], line[1])
			tempQuiz := Quiz {Question: strings.TrimSpace(line[0]), Answer: strings.TrimSpace(line[1])}
			quizArr = append(quizArr, tempQuiz)

		}

		fmt.Println("Starting Quiz")

		// new stdin reader
		stdReader := bufio.NewReader(os.Stdin)
		
		fmt.Println("Press <Enter> to start")
		_, _ = stdReader.ReadString('\n')

		scoreChan := make(chan int)
		timerChan := make(chan bool)

		go startQuiz(quizArr, stdReader, scoreChan)
		go runTimer(*timerFlag, timerChan)

		select {
		case <- timerChan:
			fmt.Println("Quiz time is over!")
		case score := <- scoreChan:
			fmt.Printf("You scored %d", score)
		}
	}
	
}

func runTimer(timeInSec int, timerChan chan<- bool) {
	time.Sleep(time.Duration(timeInSec) * time.Second)
	timerChan <- true
}

func startQuiz(quizArr []Quiz, stdReader *bufio.Reader, scoreChan chan<- int) {
	
	score := 0
	for _, qz := range quizArr {
		fmt.Printf("%s\n", qz.Question)
	
		answer, _ := stdReader.ReadString('\n')
		answer = strings.TrimSpace(answer)
	
		if strings.ToLower(answer) == strings.ToLower(qz.Answer) {
			fmt.Println("Correct answer")
			score ++
		} else {
			fmt.Printf("Incorrect answer, input-%s :: correct answer-%s\n", answer, qz.Answer)
		}
	
	}
	
	scoreChan <- score
}
