package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var userInput string = os.Args[1] // user input from the VSC terminal
	fmt.Println("you typed: ", userInput)
	inputToSlice := strings.Fields(userInput) // separates the user input around each instance of a white space and populates a slice of strings
	fmt.Println("first slice: ", inputToSlice[0])
	var bannerSlice []string

	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close() // the file descriptor is closed at the end of the main function

	scanner := bufio.NewScanner(banner) // scan the 'standard.txt' file
	line := scanner.Split(bufio.ScanWords) // split the content of 'standard.txt'file by line (= image)
	for scanner.Scan() {
		bannerSlice = append(bannerSlice, line) // populate the bannerSlice
	}
	if err := scanner.Err(); err != nil { 
        fmt.Println(err)
    }
	  
	theMap := make(map[string]rune) // matching the strings from 'standard.txt' to the runes in each userInput sub-slice
		
	m[bannerSlice[0]] = rune(31)


	}
}
