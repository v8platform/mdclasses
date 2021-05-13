package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Report struct {
	XMLName                   xml.Name           `xml:"Report"`
	Mdclass                   string             `xml:"mdclass,attr"`
	Uuid                      string             `xml:"uuid,attr"`
	ProducedTypes             MDOProducedTypes   `xml:"producedTypes"`
	Name                      string             `xml:"name"`
	Synonym                   ObjectKeyValueType `xml:"synonym"`
	DefaultForm               string             `xml:"defaultForm"`
	MainDataCompositionSchema string             `xml:"mainDataCompositionSchema"`
	DefaultSettingsForm       string             `xml:"defaultSettingsForm"`
	DefaultVariantForm        string             `xml:"defaultVariantForm"`
	Help                      Help               `xml:"help"`
	Templates                 struct {
		Uuid         string             `xml:"uuid,attr"`
		Name         string             `xml:"name"`
		Synonym      ObjectKeyValueType `xml:"synonym"`
		TemplateType string             `xml:"templateType"`
	} `xml:"templates"`
}
