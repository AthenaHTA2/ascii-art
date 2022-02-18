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
	var word []string

	for i := 1; i < len(os.Args); i++ {
		inputToSlice = append(inputToSlice, userInput[i]) // put user input into a string slice
	}
	fmt.Println(inputToSlice)

	var bannerLines []string // a slice of strings containing the 95 images from standard.txt file

	banner, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer banner.Close() // the file descriptor is closed at the end of the main function

	scanner := bufio.NewScanner(banner) // make new Scanner

	for scanner.Scan() { // scan the 'standard.txt' file
		line := scanner.Text()                       // store 'standard.txt' into the variable 'line'
		bannerLines = append(bannerLines, line+"\n") // separate 'standard.txt' after each instance of a new line and populate bannerSlices string slice
	}
	var bannerImages []string

	// fmt.Println(bannerLines[8:16])

	for i := 0; i < len(bannerLines); i += 9 {
		bannerImages = append(bannerImages, strings.Join(bannerLines[i:i+9], ""))
	}

	// bannerImagesString := strings.Join(bannerImagesSlice, "")

	// bannerImages = strings.SplitAfter(bannerImagesString, "\n")  // I am trying to drop the 'new line' from line 34 but is not working :-)

	// fmt.Println(bannerImages[1:95])
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	imagesMap := make(map[rune]string) // matching 95 runes in the ascii table - from dec 32 to dec 126 - to 95 images from 'standard.txt'

	// for i := 32; i <= 126; i++ {
	//	for j := 1; j <= 95; j++ {
	imagesMap[32] = bannerImages[0]
	imagesMap[33] = bannerImages[1]
	imagesMap[34] = bannerImages[2]
	imagesMap[35] = bannerImages[3]
	imagesMap[36] = bannerImages[4]
	imagesMap[37] = bannerImages[5]
	imagesMap[38] = bannerImages[6]
	imagesMap[39] = bannerImages[7]
	imagesMap[40] = bannerImages[8]
	imagesMap[41] = bannerImages[9]
	imagesMap[42] = bannerImages[10]
	imagesMap[43] = bannerImages[11]
	imagesMap[44] = bannerImages[12]
	imagesMap[45] = bannerImages[13]
	imagesMap[46] = bannerImages[14]
	imagesMap[47] = bannerImages[15]
	imagesMap[48] = bannerImages[16]
	imagesMap[49] = bannerImages[17]
	imagesMap[50] = bannerImages[18]
	imagesMap[51] = bannerImages[19]
	imagesMap[52] = bannerImages[20]
	imagesMap[53] = bannerImages[21]
	imagesMap[54] = bannerImages[22]
	imagesMap[55] = bannerImages[23]
	imagesMap[56] = bannerImages[24]
	imagesMap[57] = bannerImages[25]
	imagesMap[58] = bannerImages[26]
	imagesMap[59] = bannerImages[27]
	imagesMap[60] = bannerImages[28]
	imagesMap[61] = bannerImages[29]
	imagesMap[62] = bannerImages[30]
	imagesMap[63] = bannerImages[31]
	imagesMap[64] = bannerImages[32]
	imagesMap[65] = bannerImages[33]
	imagesMap[66] = bannerImages[34]
	imagesMap[67] = bannerImages[35]
	imagesMap[68] = bannerImages[36]
	imagesMap[69] = bannerImages[37]
	imagesMap[70] = bannerImages[38]
	imagesMap[71] = bannerImages[39]
	imagesMap[72] = bannerImages[40]
	imagesMap[73] = bannerImages[41]
	imagesMap[74] = bannerImages[42]
	imagesMap[75] = bannerImages[43]
	imagesMap[76] = bannerImages[44]
	imagesMap[77] = bannerImages[45]
	imagesMap[78] = bannerImages[46]
	imagesMap[79] = bannerImages[47]
	imagesMap[80] = bannerImages[48]
	imagesMap[81] = bannerImages[49]
	imagesMap[82] = bannerImages[50]
	imagesMap[83] = bannerImages[51]
	imagesMap[84] = bannerImages[52]
	imagesMap[85] = bannerImages[53]
	imagesMap[86] = bannerImages[54]
	imagesMap[87] = bannerImages[55]
	imagesMap[88] = bannerImages[56]
	imagesMap[89] = bannerImages[57]
	imagesMap[90] = bannerImages[58]
	imagesMap[91] = bannerImages[59]
	imagesMap[92] = bannerImages[60]
	imagesMap[93] = bannerImages[61]
	imagesMap[94] = bannerImages[62]
	imagesMap[95] = bannerImages[63]
	imagesMap[96] = bannerImages[64]
	imagesMap[97] = bannerImages[65]
	imagesMap[98] = bannerImages[66]
	imagesMap[99] = bannerImages[67]
	imagesMap[100] = bannerImages[68]
	imagesMap[101] = bannerImages[69]
	imagesMap[102] = bannerImages[70]
	imagesMap[103] = bannerImages[71]
	imagesMap[104] = bannerImages[72]
	imagesMap[105] = bannerImages[73]
	imagesMap[106] = bannerImages[74]
	imagesMap[107] = bannerImages[75]
	imagesMap[108] = bannerImages[76]
	imagesMap[109] = bannerImages[77]
	imagesMap[110] = bannerImages[78]
	imagesMap[111] = bannerImages[79]
	imagesMap[112] = bannerImages[80]
	imagesMap[113] = bannerImages[81]
	imagesMap[114] = bannerImages[82]
	imagesMap[115] = bannerImages[83]
	imagesMap[116] = bannerImages[84]
	imagesMap[117] = bannerImages[85]
	imagesMap[118] = bannerImages[86]
	imagesMap[119] = bannerImages[87]
	imagesMap[120] = bannerImages[88]
	imagesMap[121] = bannerImages[89]
	imagesMap[122] = bannerImages[90]
	imagesMap[123] = bannerImages[91]
	imagesMap[124] = bannerImages[92]
	imagesMap[125] = bannerImages[93]
	imagesMap[126] = bannerImages[94]

	// fmt.Println("imagesMap: \n", imagesMap)

	for i := 0; i < len(inputToSlice); i++ { // cycle through each string slice from terminal input
		runeInput := []rune(inputToSlice[i]) // cast each input slice to a rune slice

		for j := 0; j < len(runeInput); j++ { // cycle through each letter in the rune slice
			if runeInput[j] == 10 {
				fmt.Printf(word[j])
				fmt.Println()
			}
			if len(runeInput) != 0 && inputToSlice[i] != "\\n" {
				if runeInput[j] >= ' ' && runeInput[j] <= '~' { // if the letter matches any of the ASCII characters in the 32 to 126 range
					word = append(word, imagesMap[runeInput[j]]) // print the corresponding image in imagesMap
				}
			}
		}

	}
	result := strings.Join(word, " ")
	fmt.Printf("%s", result)
}
