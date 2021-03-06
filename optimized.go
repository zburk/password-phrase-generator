package main

import (
	"bufio"
	"bytes"
	crand "crypto/rand"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func optimized() []string {
	fileLineCount, err := os.Open("words")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLineCount.Close()
	totalLineCount, _ := lineCounter(fileLineCount)

	fileLineCount.Seek(0, 0)
	reader := bufio.NewReader(fileLineCount)

	var wordLookup []string
	var wordsInPhrase []string

	for len(wordsInPhrase) < 3 {
		// Generate a random index
		randomIndex, err := crand.Int(crand.Reader, big.NewInt(int64(totalLineCount)))
		if err != nil {
			log.Println(err)
		}

		randomIndexInt, err := strconv.Atoi(randomIndex.String())
		if err != nil {
			log.Println(err)
		}

		var attemptedWord string
		if randomIndexInt < len(wordLookup) {
			attemptedWord = wordLookup[randomIndexInt]
		} else {
			for randomIndexInt > len(wordLookup) {
				line, err := reader.ReadString('\n')
				line = strings.TrimSuffix(line, "\n")

				if err != nil {
					break
				}

				wordLookup = append(wordLookup, line)
				if randomIndexInt == len(wordLookup) {
					attemptedWord = line
				}
			}
		}

		if allowWord(attemptedWord) {
			wordsInPhrase = append(wordsInPhrase, attemptedWord)
		}
	}

	return wordsInPhrase
}

// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
