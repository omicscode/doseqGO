package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

import (
	"os"
)

func main() {

	argread := os.Args
	argument := argread[1]
	if len(argument) == 0 {
		panic("argument cant be empty")
	} else {

	}
	result := make(chan []FastaSequence, 1)
	go readFasta(argument, result)

}
