package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := flag.String("file", "problems.csv", "csv file in the format of question and answer")
	flag.Parse()

	file, err := os.Open(*fileName)
	check(err)

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	check(err)

	problems := evaluateLines(lines)

	correctCounter := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("correct!")
			correctCounter++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctCounter, len(problems))
}

func evaluateLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}
