package mdclasses

import "encoding/xml"

type Catalog struct {
	MDOBaseType
	XMLName                       xml.Name            `xml:"Catalog"`
	ProducedTypes                 MDOProducedTypes    `xml:"producedTypes"`
	InputByString                 string              `xml:"inputByString"`
	FullTextSearchOnInputByString string              `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []StandardAttribute `xml:"standardAttributes"`
	CreateOnInput                 string              `xml:"createOnInput"`
	DataLockControlMode           string              `xml:"dataLockControlMode"`
	FullTextSearch                string              `xml:"fullTextSearch"`
	ObjectPresentation            ObjectKeyValueType  `xml:"objectPresentation"`
	Hierarchical                  string              `xml:"hierarchical"`
	LevelCount                    string              `xml:"levelCount"`
	FoldersOnTop                  string              `xml:"foldersOnTop"`
	DescriptionLength             string              `xml:"descriptionLength"`
	CodeType                      string              `xml:"codeType"`
	CodeAllowedLength             string              `xml:"codeAllowedLength"`
	DefaultPresentation           string              `xml:"defaultPresentation"`
	EditType                      string              `xml:"editType"`
	ChoiceMode                    string              `xml:"choiceMode"`
	Attributes                    []Attribute         `xml:"attributes"`
	TabularSections               []TabularSection    `xml:"tabularSections"`
}
