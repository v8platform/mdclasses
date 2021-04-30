package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type CommonTemplates struct {
	MDOBaseType
	XMLName      xml.Name `xml:"CommonTemplate"`
	TemplateType string   `xml:"templateType"`
	Comment      string   `xml:"comment"`
}
