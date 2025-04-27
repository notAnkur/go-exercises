package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"strings"
	"io"
	"bufio"
)

type Quiz struct {
	Question string
	Answer string
}

func main() {
	var helpFlag = flag.Bool("h", false, "help")
	var questionCSVPathFlag = flag.String("csv", "questions.csv", "Specify the question csv filename.")
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

		score := 0

		fmt.Println(len(quizArr))
		for _, qz := range quizArr {
			// print question
			fmt.Printf("%s\n", qz.Question)

			// wait for answer
			answer, _ := stdReader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if strings.ToLower(answer) == strings.ToLower(qz.Answer) {
				fmt.Println("Correct answer")
				score ++
			} else {
				fmt.Printf("Incorrect answer, input-%s :: correct answer-%s\n", answer, qz.Answer)
			}

		}

		if score <=0 {
			fmt.Printf("Your score is %d\n", 0)
		} else {
			fmt.Printf("Your score is %d\n", score)
		}
	}

}
