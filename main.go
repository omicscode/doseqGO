package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

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

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	finalseq := []lastseq{}

	app.Get("/submit", func(c *fiber.Ctx) error {
		f := new(PageData)
		for _, v := range readseqannotate() {
			if f.Message == v.id {
				finalseq = append(finalseq, lastseq{
					id:             v.id,
					transcriptseq:  v.transcriptseq,
					cdsseq:         v.cdsseq,
					proteinseq:     v.proteinseq,
					idtype:         v.idtype,
					annotation:     v.annotation,
					annotate:       v.annotate,
					GeneAnnotation: v.GeneAnnotation,
					AnnotationType: v.AnnotationType,
					IPR:            v.IPR,
				})
			}
		}
		return c.Render("finalseq", fiber.Map{
			"Finalseq": finalseq,
		})
	})
	log.Fatal(app.Listen(":3000"))
}

func readseqannotate() []lastseq {
	cds := readcds()
	transcript := readtranscriptome()
	diamondresult := diamondResult()
	diamondannotate := diamondAnnotate()
	protein := readprotein()
	poannotate := readannotate()

	returnvecannotate := []lastseq{}

	for _, v := range cds {
		for _, id := range transcript {
			for _, diaresult := range diamondresult {
				for _, hint := range poannotate {
					for _, prot := range protein {
						for _, dia := range diamondannotate {
							if v.header == id.header || prot.header == diaresult.id || dia.id == hint.UnigeneID {
								returnvecannotate = append(returnvecannotate, lastseq{
									id:             v.header,
									transcriptseq:  id.sequence,
									cdsseq:         v.sequence,
									proteinseq:     prot.sequence,
									idtype:         dia.idtype,
									annotation:     dia.annotation,
									annotate:       diaresult.annotate,
									GeneAnnotation: hint.GeneAnnotation,
									AnnotationType: hint.AnnotationType,
									IPR:            hint.IPR,
								})
							}
						}
					}
				}
			}
		}
	}
	return returnvecannotate
}
