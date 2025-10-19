package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

/*
 *
 * list := set.From[string](items).Slice()
 */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DiamondResult struct {
	id         string
	idtype     string
	annotation string
}

func diamondResult(result chan []DiamondResult) {

	filename := "./serverfiles/diamonsresult.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	finalannotatewrite := []DiamondResult{}
	for scanner.Scan() {
		line := scanner.Text()
		linevec := strings.Split(line, "\t")
		linevecID := string(linevec[0])
		linevecannotation := string(linevec[1])
		restline := strings.Replace(strings.Replace(strings.Join(linevec[2:], "-"), "[", "", -1), "]", "", -1)
		finalannotatewrite = append(finalannotatewrite, DiamondResult{
			id:         linevecID,
			idtype:     linevecannotation,
			annotation: restline,
		})
	}

}
