package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func printReader(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		for _, r := range line {
			z01.PrintRune(r)
		}
		if err != nil {
			break
		}
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printReader(bufio.NewReader(os.Stdin))
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		printReader(bufio.NewReader(file))
		file.Close()
	}
}
