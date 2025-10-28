package main

/*
Gaurav Sablok
codeprog@icloud.com
*/

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func variantbrowse() []VCF {

	vcfvector := []VCF{}

	pathfile := "./testfiles/sample.vcf"

	file, err := os.Open(pathfile)
	if err != nil {
		log.Fatal("file not found")
	}
	fileread := bufio.NewScanner(file)
	if err != nil {
		log.Fatal("cant scan the file")
	}
	for fileread.Scan() {
		line := fileread.Text()
		if strings.HasPrefix(line, "#") {
			continue
		} else if !strings.HasPrefix(line, "#") {
			linestring := strings.Split(line, "\t")
			posindex, err := strconv.Atoi(linestring[1])
			qualityvalue, err := strconv.ParseFloat(linestring[6], 64)
			if err != nil {
				log.Fatal("position not found")
			}
			vcfvector = append(vcfvector, VCF{
				Chrom:       linestring[0],
				Pos:         posindex,
				Id:          linestring[3],
				Ref:         linestring[4],
				Alt:         linestring[5],
				Quality:     qualityvalue,
				Filter:      linestring[7],
				Information: strings.Join(strings.Split(linestring[8], ";"), "-"),
			})
		}
	}
	return vcfvector
}
