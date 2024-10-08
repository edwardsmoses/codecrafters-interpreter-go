package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	rawFileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	for _, token := range rawFileContents {
		if rune(token) == LEFT_PAREN {
			fmt.Println("LEFT_PAREN", string(token), "null")
		} else if rune(token) == RIGHT_PAREN {
			fmt.Println("RIGHT_PAREN", string(token), "null")
		} else if rune(token) == LEFT_BRACE {
			fmt.Println("LEFT_BRACE", string(token), "null")
		} else if rune(token) == RIGHT_BRACE {
			fmt.Println("RIGHT_BRACE", string(token), "null")
		}
	}

	fmt.Println("EOF  null")
}
