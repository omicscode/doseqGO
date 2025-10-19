package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FastaSequence struct {
	Header   string
	Sequence string
}

func readtranscriptome(result chan []FastaSequence) {

	filename := "./serverfiles/transcriptome.fasta"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	header := []string{}
	sequence := []string{}

	newseq := []FastaSequence{}

	for scanner.Scan() {
		lineread := scanner.Text()
		if strings.HasPrefix(lineread, ">") {
			lineinsert := lineread
			header = append(header, lineinsert)
		}
		if !strings.HasPrefix(lineread, ">") {
			linesequence := lineread
			sequence = append(sequence, linesequence)
		}
	}

	for i := 0; i <= len(header); i++ {
		newseq = append(newseq, FastaSequence{
			Header:   header[i],
			Sequence: sequence[i],
		})
	}
	result <- newseq
}
