package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type readStore struct {
	idreadref    string
	idreadsample string
	tag          string
	seq          string
}

func sam(result chan []readStore) {

	openFile, err := os.Open("readFile")
	if err != nil {
		log.Fatal(err)
	}

	readBuffer := bufio.NewScanner(openFile)
	for readBuffer.Scan() {
		line := readBuffer.Text()
		if strings.HasPrefix(line, "@") {
			continue
		} else {
			data := []readStore{}
			data = append(data, readStore{
				idreadref:    strings.Split(line, "\t")[0],
				idreadsample: strings.Split(line, "\t")[2],
				tag:          strings.Split(line, "\t")[5],
				seq:          strings.Split(line, "\t")[9],
			})
		}
	}
}
