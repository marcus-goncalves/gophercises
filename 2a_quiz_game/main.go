package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "static/problems.csv", "a csv file that contains 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open csv file: %s \n", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Sprintf("Failed to parse file.")
		os.Exit(1)
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		var answer string

		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		fmt.Scanf("%s \n", &answer)

		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.. \n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}
