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
		me := split[1]

		total += score(opponent, me)
	}

	fmt.Println(total)
}

func score(opponent string, me string) int {
	var (
		shapeScore   int
		outcomeScore int
	)

	switch me {
	case "X":
		shapeScore = 1
		switch opponent {
		case "A":
			outcomeScore = 3
		case "B":
			outcomeScore = 0
		case "C":
			outcomeScore = 6
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	case "Y":
		shapeScore = 2
		switch opponent {
		case "A":
			outcomeScore = 6
		case "B":
			outcomeScore = 3
		case "C":
			outcomeScore = 0
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	case "Z":
		shapeScore = 3
		switch opponent {
		case "A":
			outcomeScore = 0
		case "B":
			outcomeScore = 6
		case "C":
			outcomeScore = 3
		default:
			log.Fatalf("Unsupported shape: %s", opponent)
		}
	default:
		log.Fatalf("Unsupported shape: %s", me)
	}

	return shapeScore + outcomeScore
}
