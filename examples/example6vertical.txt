package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputToSlice := os.Args[1:] // user input from the VSC terminal
	inputSlices := strings.Split(inputToSlice[0], "\\n")

	var bannerLines []string // a slice of strings containing the 95 images from standard.txt file
	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close()                // the file descriptor is closed at the end of the main function
	scanner := bufio.NewScanner(banner) // make new Scanner

	for scanner.Scan() { // scan the 'standard.txt' file

		line := scanner.Text() // store a line from 'standard.txt' into the variable 'line'

		bannerLines = append(bannerLines, line+"\n") // separate 'standard.txt' after each instance of a new line and populate bannerSlices string slice
	}

	var bannerImages []string

	for a := 0; a <= 95; a++ {
		image := ""
		var step int
		for i := 1; i < 9; i++ {
			if a == 0 {
				step = a + 1
				image += bannerLines[step+i]
			} else if a <= 94 {
				step = a * 9
				image += bannerLines[step+i] // bannerImages = append(bannerImages, strings.Join(bannerLines[i:i+9], ""))
			} else {
				step = a * 8
				image += bannerLines[step+i]
			}
		}

		bannerImages = append(bannerImages, image)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	imagesMap := make(map[rune]string) // matching 95 runes in the ascii table - from dec 32 to dec 126 - to 95 images from 'standard.txt'

	for i := 32; i < 127; i++ {
		imagesMap[rune(i)] = bannerImages[i-32]
	}

	var words []string
	for _, k := range inputSlices {

		var word string
		var chunk string
		for i := 0; i < len(k); i++ { // cycle through each string slice from inputSlices
			runeInput := []rune(string(k[i]))     // cast each input slice to a rune slice
			for j := 0; j < len(runeInput); j++ { // cycle through each letter in the rune slice

				if len(runeInput) != 0 && runeInput[j] != '\n' {
					if runeInput[j] >= ' ' && runeInput[j] <= '~' { // if the letter matches any of the ASCII characters in the 32 to 126 range
						chunk = imagesMap[runeInput[j]]
					}
				}
			}
			word += chunk // print the corresponding image in imagesMap
			chunk = ""
		}
		words = append(words, word)

	}

	for i, word := range words {
		if len(words) > 1 && i != len(words) {

			fmt.Print(word)
			fmt.Println()
		} else {
			fmt.Print(word)
		}
	}
}
