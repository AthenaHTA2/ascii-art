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

	var bannerSlices []string //a slice of strings containing the 95 images from standard.txt file
	counter := 0
	//var strLine []string

	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close() // the file descriptor is closed at the end of the main function

	scanner := bufio.NewScanner(banner) // make new Scanner
	for scanner.Scan() {                //scan the 'standard.txt' file
		line := scanner.Text()                   // store 'standard.txt' into the variable 'line'
		bannerSlices = strings.Split(line, "\n") //separate 'standard.txt' after each instance of a new line and populate a slice of strings
		fmt.Println("length of bannerSlices: ", counter, bannerSlices)
		counter = counter + 1
		//fmt.Println(bannerSlices)
		//for j := 10; j <= 35; j++ {
		//	fmt.Println("printing line by line", bannerSlices[j])
		//}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	//theMap := make(map[string]rune) // matching 95 images from 'standard.txt' to 95 runes in the ascii table, from dec 32 to dec 126.

	//m[bannerSlice[0]] = rune(31)

}
