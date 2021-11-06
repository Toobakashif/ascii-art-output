package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Flag()
	Banner()
	data, _ := os.ReadFile(os.Args[2] + ".txt")

	words := strings.Split(os.Args[1], "\\n")
	lines := strings.Split(string(data), "\n")
	vastus := ""
	for index := range words {
		if words[index] == "" { // break if there is no word
			break
		}

		letters := strings.Split(words[index], "") //slice the words into symbols
		var ascii []int
		for i := range letters {
			l := int([]rune(letters[i])[0]) //store the rune value of the letter in l
			ascii = append(ascii, l)        //append l value to ascii untill all letters are appended
		}
		//vastus = ""
		for j := 1; j < 9; j++ { //loop 8 times because a row is 8 lines long
			s := ""
			for k := range ascii { //loop through all runes in ascii
				line := (ascii[k] - 32) * 9 //this finds the row before the ascii symbol in the txt file
				s += lines[line+j]        //gives the string all the lines into the string
			}
			vastus += s + "\n"
		}

	}
	f, _ := os.Create(strings.ReplaceAll(os.Args[3], "--output=", ""))

	defer f.Close()
	_, err2 := f.WriteString(vastus)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Done! Check the", (strings.ReplaceAll(os.Args[3], "--output=", "")), "for answer")
}

func Flag() {
	switch {
	case len(os.Args) < 4, len(os.Args) > 5, strings.HasPrefix(os.Args[3], "--output=") == false:
	
		cRt := "\033[0m"
		cRd := "\033[31m"
		fmt.Println(string(cRd), "Usage: go run . [STRING] [OPTION]\n\nEx: go run . something standard --output=<fileName.txt>")
		fmt.Print(string(cRt))
		os.Exit(0)
	}

	switch {
	case len(os.Args) < 4, len(os.Args) > 5, strings.HasSuffix(os.Args[3], ".txt") == false:
		cRt := "\033[0m"
		cRd := "\033[31m"
		fmt.Println(string(cRd), "Usage: go run . [STRING] [OPTION]\n\nEx: go run . something standard --output=<fileName.txt>")
		fmt.Print(string(cRt))
		os.Exit(0)
	}
}

func Banner() {
	if strings.Contains(os.Args[2], "standard") == true {
		return
	} else if strings.Contains(os.Args[2], "shadow") == true {
		return
	} else if strings.Contains(os.Args[2], "thinkertoy") == true {
		return
	} else {
		cRd := "\033[31m"
		fmt.Println(string(cRd), "Error! Use correct banner name!")
		os.Exit(0)
	}
}
