package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func countWords(file *os.File) int {
	words := 0
	inWord := false

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if unicode.IsSpace(char) {
			if inWord {
				words++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		words++
	}

	return words
}

func countLines(file *os.File) int {
	lines := 0

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if char == '\n' {
			lines++
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	return lines
}

func countCharacters(file *os.File) int {
	characters := 0

	reader := bufio.NewReader(file)
	for {
		_, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		characters++
	}

	return characters
}

func countBytes(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return fileInfo.Size()
}

func main() {

	lastArg := len(os.Args) - 1
	fileName := string(os.Args[lastArg])

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if len(os.Args) == 1 {
		fmt.Println("Usage: wc [-lwc] [file]")
	} else if len(os.Args) == 2 {
		words := countWords(file)
		file.Seek(0, 0)
		lines := countLines(file)
		bytes := countBytes(file)

		fmt.Printf("%d %d %d\n", lines, words, bytes)
	} else {
		output := ""
		for _, arr := range os.Args[1:lastArg] {
			if arr == "-l" {
				lines := countLines(file)
				output += fmt.Sprintf("%d ", lines)
			}
			if arr == "-w" {
				file.Seek(0, 0)
				words := countWords(file)
				output += fmt.Sprintf("%d ", words)
			}
			if arr == "-c" {
				file.Seek(0, 0)
				bytes := countBytes(file)
				output += fmt.Sprintf("%d ", bytes)
			}
			if arr == "-m" {
				file.Seek(0, 0)
				characters := countCharacters(file)
				output += fmt.Sprintf("%d ", characters)
			}
		}
		fmt.Println(output)
	}

}
