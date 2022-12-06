package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("inputs/2022-2.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	buf := bufio.NewScanner(input)
	for buf.Scan() {
		text := buf.Text()
		split := strings.Split(text, " ")
		opponent := split[0]
		outcoem := split[1]

		total += score(opponent, outcoem)
	}

	fmt.Println(total)
}

func score(opponent string, outcome string) int {
	var (
		shapeScore   int
		outcomeScore int
	)

	switch outcome {
	case "X":
		outcomeScore = 0
		switch opponent {
		case "A":
			shapeScore = 3
		case "B":
			shapeScore = 1
		case "C":
			shapeScore = 2
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	case "Y":
		outcomeScore = 3
		switch opponent {
		case "A":
			shapeScore = 1
		case "B":
			shapeScore = 2
		case "C":
			shapeScore = 3
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	case "Z":
		outcomeScore = 6
		switch opponent {
		case "A":
			shapeScore = 2
		case "B":
			shapeScore = 3
		case "C":
			shapeScore = 1
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	default:
		log.Fatalf("Unsupported outcome: %s", outcome)
	}

	return shapeScore + outcomeScore
}
