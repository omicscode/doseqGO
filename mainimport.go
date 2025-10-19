package main


import ("transcriptome"
	"cds"
	"protein"
	"diamonannotate"
	"diamondresult"
)


type last struct {
	sequence string
	id string
	nucleotide string
	cds string
	protein string
	GeneAnnotation string
	AnnotationType string
	IPR            []string
	annotation string
	annotate []string
}

func outputmain(

transcriptome := make(chan []FastaSequence, 1)
	go readtranscriptome(result)

cds := make(chan []FastaSequence, 1)
	go readcds(result)

protein := make(chan []FastaSequence, 1)
	go readprotein(result)

annotate := make(chan []POAnnotate, 1)
		go readannotate(result)

diamonresultan := make(chan []DiamondResult, 1)
				go diamonsresult(result)

diamondannotation := make(chan []DiamondAnnoate, 1)
 go diamondResult(result)
)
