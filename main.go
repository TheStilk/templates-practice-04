package main

import (
	"fmt"

	"prac04/creator"
	"prac04/factory"
)

func main() {
	fmt.Printf("Document System with Factory Method:\n")

	creators := []creator.DocumentCreator{
		&factory.ReportCreator{},
		&factory.ResumeCreator{},
		&factory.LetterCreator{},
		&factory.InvoiceCreator{},
	}

	for _, c := range creators {
		doc := c.CreateDocument()
		doc.Open()
	}

	fmt.Printf("\nDynamic Document Creation\n")

	docTypes := []string{"report", "resume", "letter", "invoice", "unknown"}

	for _, docType := range docTypes {
		creator := factory.CreateCreatorByType(docType)
		if creator == nil {
			fmt.Printf("Unknown document type: %s\n", docType)
			continue
		}
		doc := creator.CreateDocument()
		doc.Open()
	}
}
