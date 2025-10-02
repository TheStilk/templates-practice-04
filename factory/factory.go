package factory

import (
	"prac04/creator"
	"prac04/document"
	"prac04/invoice"
	"prac04/letter"
	"prac04/report"
	"prac04/resume"
)

type ReportCreator struct{}

func (rc *ReportCreator) CreateDocument() document.Document {
	return &report.Report{}
}

type ResumeCreator struct{}

func (rc *ResumeCreator) CreateDocument() document.Document {
	return &resume.Resume{}
}

type LetterCreator struct{}

func (lc *LetterCreator) CreateDocument() document.Document {
	return &letter.Letter{}
}

type InvoiceCreator struct{}

func (ic *InvoiceCreator) CreateDocument() document.Document {
	return &invoice.Invoice{}
}

func CreateCreatorByType(docType string) creator.DocumentCreator {
	switch docType {
	case "report":
		return &ReportCreator{}
	case "resume":
		return &ResumeCreator{}
	case "letter":
		return &LetterCreator{}
	case "invoice":
		return &InvoiceCreator{}
	default:
		return nil
	}
}
