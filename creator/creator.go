package creator

import "prac04/document"

type DocumentCreator interface {
	CreateDocument() document.Document
}
