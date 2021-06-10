package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
	"os"
)

type PackConfig struct {
	Base    string
	Path    string
	Objects []interface{}
}

func packObjectMdo(path string, object interface{}) {
	xmlFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("Ошибка создания файла %s", err.Error())
		return
	}
	xmlFile.WriteString(xml.Header)
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "  ")
	err = encoder.Encode(&object)
	if err != nil {
		log.Fatalf("Ошибка сохранения объекта в файл", object)
		return
	}
	xmlFile.WriteString("\n")
}
