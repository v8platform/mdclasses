package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Document struct {
	MDOBaseType
	XMLName xml.Name `xml:"Document"`
}
