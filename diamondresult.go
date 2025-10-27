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

type DiamondResult struct {
	id       string
	annotate []string
}

func diamondResult() []DiamondResult {

	filename := "./serverfiles/diamondresult.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	linemape := make(map[string]struct{})
	linemaplist := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.Split(line, "\t")
		vectorinsert := text[0]
		linemape[vectorinsert] = struct{}{}
	}

	for item := range linemape {
		linemaplist = append(linemaplist, item)
	}

	finalvec := []DiamondResult{}
	for scanner.Scan() {
		for i, _ := range linemaplist {
			linevec := scanner.Text()
			linevecfirst := strings.Split(linevec, "\t")
			if linevecfirst[0] == linemaplist[i] {
				valuestore := []string{}
				valuestore = append(valuestore, linemaplist[i])
				finalvec = append(finalvec, DiamondResult{
					id:       linemaplist[i],
					annotate: valuestore,
				})
			}
		}
	}
	return finalvec
}
