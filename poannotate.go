package main

/*
 Gaurav Sablok
 codeprog@icloud.com
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type POAnnotate struct {
	GeneAnnotation string
	AnnotationType string
	UnigeneID      string
	IPR            []string
}

func readannotate() []POAnnotate {

	filename := "./serverfiles/poannotate.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	poannotate := []POAnnotate{}

	for scanner.Scan() {
		line := scanner.Text()
		linevector := strings.Split(line, "\t")
		IPRvector := strings.Split(linevector[4], ",")
		poannotate = append(poannotate, POAnnotate{
			GeneAnnotation: linevector[0],
			AnnotationType: linevector[1],
			UnigeneID:      linevector[2],
			IPR:            IPRvector,
		})
	}
	return poannotate
}
