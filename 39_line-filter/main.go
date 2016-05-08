package main

/*
line filter: common type of program that reads input on stdin, processes it, and
then prints some derived result to stdout like grep or sed
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// os.Stdin is unbuffered so wrapping it into
	// a buffered scanner advances it to the next line
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		// Text is the current token (line)
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// end of file is expected and not reported as an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
