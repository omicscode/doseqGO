package main

/*
 Gaurav Sablok
 codeprog@icloud.com
*/

type FormChrom struct {
	Value string
}

type FormPos struct {
	Value string
}

type FormID struct {
	Value string
}

type FormRef struct {
	value string
}

type FormAlt struct {
	Value string
}

type FormQuality struct {
	Value string
}

type FormFilter struct {
	Value string
}

type FormInfo struct {
	Value string
}

type VCF struct {
	Chrom       string
	Pos         int
	Id          string
	Ref         string
	Alt         string
	Quality     float64
	Filter      string
	Info        string
	Information string
}

type PageData struct {
	Message string
}

type lastseq struct {
	id             string
	transcriptseq  string
	cdsseq         string
	proteinseq     string
	idtype         string
	annotation     string
	annotate       []string
	GeneAnnotation string
	AnnotationType string
	UnigeneID      string
	IPR            []string
}
