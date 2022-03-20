package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main(){
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("failed to open CSV file: %s\n", *csvFile))
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the file.")
	}
	problems := parse(line)

	var score int
	for i, p := range problems { 
		fmt.Printf("\nCurrent Score: %d\n", score)
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a { 
			fmt.Println("Correct!")
			score += 1
		}else{
			fmt.Println("incorrect :(")
		}
	fmt.Printf("Score Overall: %d/%d", score, len(problems))
	}

	

}
func parse(lines [][]string) []problem {
	retrn := make([]problem, len(lines))
	for i, line := range lines {
		retrn[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return retrn
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}