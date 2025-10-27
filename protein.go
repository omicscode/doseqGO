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

func readprotein() []FastaSequence {

	filename := "./serverfiles/protein.fasta"
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

	newseqprotein := []FastaSequence{}

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
		newseqprotein = append(newseqprotein, FastaSequence{
			header:   header[i],
			sequence: sequence[i],
		})
	}
	return newseqprotein
}
