package main

import (
	"fmt"
	"regexp"
	"runtime"
)

var (
	fnRegex = regexp.MustCompile(`^.*\.(.*)$`)
	debug   = false
)

func enter() string {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "<UNK>"
	}

	funcName := fnRegex.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	fmt.Printf("Entering: %v\n", funcName)
	if debug {
		fmt.Println("File:", file)
		fmt.Println("Line:", line)
	}

	return funcName
}

func exit(fn string) {
	fmt.Printf("Exiting: %v\n", fn)
}

func foo() {
	defer exit(enter())
	fmt.Println("Do something...")
}

func main() {
	defer exit(enter())
	foo()
}
