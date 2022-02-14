package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var inputToSlice []string
	userInput := os.Args // user input from the VSC terminal

	for i := 1; i < len(os.Args); i++ {
		inputToSlice = append(inputToSlice, userInput[i]) // put user input into a string slice
	}

	var bannerSlices []string // a slice of strings containing the 95 images from standard.txt file

	// var strLine []string

	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close() // the file descriptor is closed at the end of the main function

	scanner := bufio.NewScanner(banner) // make new Scanner
	for scanner.Scan() {                // scan the 'standard.txt' file
		line := scanner.Text() // store 'standard.txt' into the variable 'line'
		// fmt.Println("Line: ", line)
		bannerSlices = append(bannerSlices, line+"\n")
		// separate 'standard.txt' after each instance of a new line and populate a slice of strings
		// fmt.Println("length of bannerSlices: ", counter, bannerSlices)

	}
	var bannerArray []string
	fmt.Println(bannerSlices[8:16])

	for i := 0; i < len(bannerSlices); i += 9 {
		bannerArray = append(bannerArray, strings.Join(bannerSlices[i:i+9], ""))
	}
	fmt.Println(bannerArray[3])
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// theMap := make(map[string]rune) // matching 95 images from 'standard.txt' to 95 runes in the ascii table, from dec 32 to dec 126.

	// m[bannerSlice[0]] = rune(31)
}
