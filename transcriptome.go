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
	header   string
	sequence string
}

func readtranscriptome() []FastaSequence {

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

	newseqtranscript := []FastaSequence{}

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
		newseqtranscript = append(newseqtranscript, FastaSequence{
			header:   header[i],
			sequence: sequence[i],
		})
	}
	return newseqtranscript
}
