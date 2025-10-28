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

	app.Get("/variant", func(c *fiber.Ctx) error {
		return c.Render("form", nil)
	})

	app.Post("/variant", func(c *fiber.Ctx) error {
		fvalue := new(FormChrom)
		if err := c.BodyParser(fvalue); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("id not found")
		}
		idreturn := []VCF{}
		valuereturn := variantbrowse()
		for _, key := range valuereturn {
			if fvalue.Value == key.Chrom {
				idreturn = append(idreturn, VCF{
					Chrom:       key.Chrom,
					Pos:         key.Pos,
					Id:          key.Id,
					Ref:         key.Ref,
					Alt:         key.Alt,
					Quality:     key.Quality,
					Filter:      key.Filter,
					Information: key.Information,
				})
			}
		}
		return c.Render("formid", fiber.Map{
			"idreturnvalue": idreturn,
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
