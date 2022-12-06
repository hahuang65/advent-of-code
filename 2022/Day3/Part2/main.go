package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		counter = 0
		sum     = 0
		group   []string
	)

	input, err := os.Open("inputs/2022-3.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewScanner(input)
	for buf.Scan() {
		counter += 1
		group = append(group, buf.Text())

		if counter == 3 {
			intersection := intersection([]rune(group[0]), []rune(group[1]), []rune(group[2]))
			sum += score(intersection)
			counter = 0
			group = []string{}
		}
	}

	fmt.Println(sum)
}

func intersection(first []rune, second []rune, third []rune) rune {
	var res rune

out:
	for _, i := range first {
		for _, j := range second {
			if i == j {
				for _, k := range third {
					if i == k {
						res = i
						break out
					}
				}
			}
		}
	}

	return res
}

func score(r rune) int {
	items := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var val int

	for i, rr := range items {
		if r == rr {
			val = i
		}
	}

	return val + 1
}
