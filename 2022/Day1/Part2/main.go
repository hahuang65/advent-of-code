package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var (
		list []int
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
			list = append(list, sum)
			sum = 0
		} else {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}

			sum += num
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	res := 0
	for _, val := range list[:3] {
		res += val
	}

	fmt.Println(res)
}
