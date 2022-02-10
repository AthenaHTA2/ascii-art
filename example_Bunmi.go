package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	word_Arg := os.Args

	result := strings.Split(word_Arg[1], "\\n") // break down the input string into a number of substrings using "\\n" (= newline?) as separator

	if word_Arg[1] == "\\n" {
		result = append(result[:1])
	}

	f, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for a := 0; a < len(result); a++ { // result is a string slice
		wordRune := []rune(result[a])
		if len(wordRune) != 0 && result[a] != "\\n" {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(wordRune); j++ {
					if lines[int(wordRune[j])*9-287+i] == "        " {
						fmt.Printf("        ")
					} else {
						fmt.Printf(lines[int(wordRune[j])*9-287+i])
					}
				}
				fmt.Print("\n")
			}
		} else {
			fmt.Print("\n")
		}
	}
}
