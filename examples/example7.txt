package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// var inputToSlice []string
	inputToSlice := os.Args[1:] // user input from the VSC terminal
	inputSlices := strings.Split(inputToSlice[0], "\\n")
	// fmt.Print("the input slices --> ", inputSlices, "\n")
	// fmt.Print("length of input slices --> ", len(inputSlices), "\n")
	//	var result string
	// fmt.Print(len(inputToSlice), "\n")
	// for i := 0; i < len(os.Args)-1; i++ {
	// 	inputToSlice = append(inputToSlice, userInput[i]) // put user input into a string slice
	// }
	// fmt.Println(inputToSlice)

	var bannerLines []string // a slice of strings containing the 95 images from standard.txt file

	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close() // the file descriptor is closed at the end of the main function

	scanner := bufio.NewScanner(banner) // make new Scanner
	z := 0
	for scanner.Scan() { // scan the 'standard.txt' file
		z++
		line := scanner.Text() // store a line from 'standard.txt' into the variable 'line'
		for len(line) < 16 {   // making each line 11 spaces long
			line = line + " "
		}
		// fmt.Print("line = ", line, "\n")
		bannerLines = append(bannerLines, line+"\n") // separate 'standard.txt' after each instance of a new line and populate bannerSlices string slice

	}
	// bannerLines = bannerLines[1:]
	/*	fmt.Print("banner line -->  ", bannerLines[9], "\n")
		fmt.Print("banner line -->  ", bannerLines[10], "\n")
		fmt.Print("banner line -->  ", bannerLines[11], "\n")*/
	var bannerImages []string
	/*	fmt.Print(z)
		fmt.Println()
		fmt.Print(len(bannerLines), "\n")*/

	// fmt.Println(bannerLines[8:16])

	for i := 0; i < len(bannerLines); i += 9 {
		bannerImages = append(bannerImages, strings.Join(bannerLines[i:i+9], ""))
	}
	// fmt.Print("banner line length-->  ", len(bannerImages[94]), " ", bannerImages[94], "\n")
	// bannerImagesString := strings.Join(bannerImagesSlice, "")

	// bannerImages = strings.SplitAfter(bannerImagesString, "\n")  // I am trying to drop the 'new line' from line 34 but is not working :-)

	// fmt.Println(bannerImages[1:95])
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	imagesMap := make(map[rune]string) // matching 95 runes in the ascii table - from dec 32 to dec 126 - to 95 images from 'standard.txt'

	// for i := 32; i <= 126; i++ {
	//	for j := 1; j <= 95; j++ {
	for i := 32; i < 127; i++ {
		imagesMap[rune(i)] = bannerImages[i-32]
	}

	//for a := 32; a < 127; a++ {
	//fmt.Println("image length: \n", len(imagesMap[rune(a)]), imagesMap[rune(a)])
	//}

	// fmt.Println("W length: \n", len(imagesMap[87]))
	var words [][]string
	for _, k := range inputSlices {
		// fmt.Print("k is equal to: ", k, "\n")
		var word []string
		for i := 0; i < len(k); i++ { // cycle through each string slice from terminal input
			runeInput := []rune(string(k[i])) // cast each input slice to a rune slice

			for j := 0; j < len(runeInput); j++ { // cycle through each letter in the rune slice
				/*	if runeInput[j] == '\n' {
					fmt.Printf(word[j])
					fmt.Println()
				}*/
				if len(runeInput) != 0 && runeInput[j] != '\n' {
					if runeInput[j] >= ' ' && runeInput[j] <= '~' { // if the letter matches any of the ASCII characters in the 32 to 126 range
						word = append(word, imagesMap[runeInput[j]]) // print the corresponding image in imagesMap
						// fmt.Print("length of word  -->", len(word), "\n")
						// fmt.Print("word  -->", word, "\n")
					}
				}
			}
		}
		words = append(words, word)
		//	word = []string{}
	}
	//fmt.Print("Words printout --> ", words, "\n")
	/*for i := 0; i < len(inputSlices[0]); i++ { // cycle through each string slice from terminal input
		runeInput := []rune(string(inputSlices[0][i])) // cast each input slice to a rune slice

		for j := 0; j < len(runeInput); j++ { // cycle through each letter in the rune slice
			if runeInput[j] == '\n' {
				fmt.Printf(word[j])
				fmt.Println()
			}
			if len(runeInput) != 0 && inputSlices[0][i] != '\n' {
				if runeInput[j] >= ' ' && runeInput[j] <= '~' { // if the letter matches any of the ASCII characters in the 32 to 126 range
					word = append(word, imagesMap[runeInput[j]]) // print the corresponding image in imagesMap
					fmt.Print("first input word  -->", len(word[j]), "\n")
				}
			}
		}

	}*/
	// fmt.Print("first input word  -->", len(word[3]), "\n")
	for _, word := range words {
		for k := 0; k < 137; k += 17 { // cycle through every space in each image, 97 in total
			for _, i := range word { // take each single word from user's input
				// fmt.Print("a simbol within word", i)
				for j, _ := range i { // look at each string in word [i]'s slice
					// if i[j] == '\n' { // if it is a new line
					// 	fmt.Print(" ") // print a space instead
					// 	break          // exit
					// }
					if i[j+k] != '\n' {
						fmt.Print(string(i[j+k]))
						// fmt.Print("j = ", j+k, "\n")
					} else {
						fmt.Print(" ") // print a space instead
						break          // exit
					}
				}
			} // for loop of a line of the whole input from the map
			fmt.Println()
		}
		fmt.Println()
	}

	//fmt.Print(word)
	// result := strings.Join(word, " ")
	//for _, i := range word {
	//for _, j := range i {
	//if j == '\n' {
	//	fmt.Print(" ")
	//	break
	//}
	// fmt.Print(string(j))
	//}
	//}
}
