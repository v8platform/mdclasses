package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Content []struct {
	MdObject string `xml:"mdObject"`
}

type ExchangePlan struct {
	MDOBaseType
	XMLName                       xml.Name            `xml:"ExchangePlan"`
	Xsi                           string              `xml:"xsi,attr"`
	Core                          string              `xml:"core,attr"`
	Mdclass                       string              `xml:"mdclass,attr"`
	Uuid                          string              `xml:"uuid,attr"`
	ThisNode                      string              `xml:"thisNode,attr"`
	ProducedTypes                 MDOProducedTypes    `xml:"producedTypes"`
	Name                          string              `xml:"name"`
	Synonym                       ObjectKeyValueType  `xml:"synonym"`
	UseStandardCommands           string              `xml:"useStandardCommands"`
	InputByString                 []string            `xml:"inputByString"`
	FullTextSearchOnInputByString string              `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []StandardAttribute `xml:"standardAttributes"`
	CreateOnInput                 string              `xml:"createOnInput"`
	DataLockFields                string              `xml:"dataLockFields"`
	DataLockControlMode           string              `xml:"dataLockControlMode"`
	FullTextSearch                string              `xml:"fullTextSearch"`
	ObjectPresentation            ObjectKeyValueType  `xml:"objectPresentation"`
	ListPresentation              ObjectKeyValueType  `xml:"listPresentation"`
	CodeLength                    string              `xml:"codeLength"`
	CodeAllowedLength             string              `xml:"codeAllowedLength"`
	DescriptionLength             string              `xml:"descriptionLength"`
	Content                       Content             `xml:"content"`
	DefaultPresentation           string              `xml:"defaultPresentation"`
	EditType                      string              `xml:"editType"`
	ChoiceMode                    string              `xml:"choiceMode"`
	Attributes                    []Attribute         `xml:"attributes"`
}
