package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "year, defaults to the current year")
	day := flag.Int("day", time.Now().Day(), "day, defaults to the current day")
	flag.Parse()

	filename := fmt.Sprintf("inputs/%d-%d.txt", *year, *day)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		path := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *year, *day)
		request, err := http.NewRequest(http.MethodGet, path, nil)
		if err != nil {
			log.Fatal(err)
		}

		session := &http.Cookie{
			Name:     "session",
			Value:    os.Getenv("AOC_SESSION"),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		}
		request.AddCookie(session)

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
