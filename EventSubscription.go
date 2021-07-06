package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Source struct {
	Types []string `xml:"types"`
}

type EventSubscription struct {
	MDOBaseType
	XMLName xml.Name           `xml:"EventSubscription"`
	Mdclass string             `xml:"mdclass,attr"`
	Uuid    string             `xml:"uuid,attr"`
	Name    string             `xml:"name"`
	Synonym ObjectKeyValueType `xml:"synonym"`
	Source  Source             `xml:"source"`
	Event   string             `xml:"event"`
	Handler string             `xml:"handler"`
}
