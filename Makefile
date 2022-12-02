include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

YEAR ?= $(shell date +"%Y")
DAY ?= $(shell date +"%-d")
PART ?= 1

## run: runs the problem solver for the given year, date and part. Defaults to part 1 of the current year and day.
.PHONY: run
run: download_input
	go run ${YEAR}/Day${DAY}/Part${PART}/main.go

## new: creates a folder for the given year and day, date and part. Defaults to the current year and day.
.PHONY: new
new:
	mkdir -p ${YEAR}/Day${DAY}/Part1
	mkdir -p ${YEAR}/Day${DAY}/Part2
	touch ${YEAR}/Day${DAY}/Part1/main.go
	touch ${YEAR}/Day${DAY}/Part2/main.go

## download_input: downloads the input file for the given year, date and part. Defaults the current year and day.
.PHONY: download_input
download_input:
	@mkdir -p inputs/
	@AOC_SESSION=${AOC_SESSION} go run util/input_downloader.go --year ${YEAR} --day ${DAY}
