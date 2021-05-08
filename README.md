# go-cron-parser

A CLI application for parsing cron expressions. More details about cron expressions can be found in the man page: https://man7.org/linux/man-pages/man5/crontab.5.html

The application will parse a given cron expression string consisting of 5 time parts plus a command then pretty print to screen.

Input:
  ```bash
  go-cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
  ```

Output:
  ```bash
    minute        0 15 30 45
    hour          0
    day of month  1 15
    month         1 2 3 4 5 6 7 8 9 10 11 12
    day of week   1 2 3 4 5
    command       /usr/bin/find
   ```

## Makefile

A [Makefile](./Makefile) is included:

 - `make build` - build application locally to `/build`
 - `make clean` - Go clean, format and vet
 - `make test` - run all tests

## Usage

Building application:
 - build: `make build`
 - run: `./build/go-cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"`

You can also run without building:
 - via Go: `go run .  "*/15 0 1,15 * 1-5 /usr/bin/find"`

## Application Design

 - `./cmd` - commands for CLI (note: Cobra is used as a CLI framework: https://github.com/spf13/cobra)
 - `./parser` - expression parser code. Made up of following:
    - a lexer to tokenise the input - currently just split by space
    - a parser to process the tokens

## Task only docs

 - [Task](./TASK.md) some notes around the task - planning/thoughts/discussion points