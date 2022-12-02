package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		high = 0
		sum  = 0
	)

	input, err := os.Open("inputs/2022-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	buf := bufio.NewScanner(input)
	for buf.Scan() {
		text := buf.Text()
		if text == "" {
			if sum > high {
				high = sum
			}
			sum = 0
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}

			sum += num
		}
	}

	fmt.Println(high)
}
