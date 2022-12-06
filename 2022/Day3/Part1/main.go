package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sum := 0

	input, err := os.Open("inputs/2022-3.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewScanner(input)
	for buf.Scan() {
		text := buf.Text()
		first, second := halveString(text)
		intersect := intersection(first, second)
		sum += score(intersect)
	}

	fmt.Println(sum)
}

func halveString(str string) ([]rune, []rune) {
	var (
		first  []rune
		second []rune
	)

	strLen := len([]rune(str))
	halfSize := strLen / 2

	for i, r := range str {
		if i < halfSize {
			first = append(first, rune(r))
		} else {
			second = append(second, rune(r))
		}
	}
	return first, second
}

func intersection(first []rune, second []rune) rune {
	var res rune

out:
	for _, i := range first {
		for _, j := range second {
			if i == j {
				res = i
				break out
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
