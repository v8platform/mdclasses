package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
)

type Catalog struct {
	MDOBaseType
	XMLName                       xml.Name            `xml:"xsi Catalog"`
	ProducedTypes                 MDOProducedTypes    `xml:"producedTypes"`
	Synonym                       ObjectKeyValueType  `xml:"synonym"`
	UseStandardCommands           string              `xml:"useStandardCommands"`
	InputByString                 string              `xml:"inputByString"`
	FullTextSearchOnInputByString string              `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []StandardAttribute `xml:"standardAttributes"`
	CreateOnInput                 string              `xml:"createOnInput"`
	IncludeHelpInContents         string              `xml:"includeHelpInContents"`
	Help                          Help                `xml:"help"`
	DataLockControlMode           string              `xml:"dataLockControlMode"`
	FullTextSearch                string              `xml:"fullTextSearch"`
	ObjectPresentation            ObjectKeyValueType  `xml:"objectPresentation"`
	Hierarchical                  string              `xml:"hierarchical"`
	LevelCount                    string              `xml:"levelCount"`
	FoldersOnTop                  string              `xml:"foldersOnTop"`
	DescriptionLength             string              `xml:"descriptionLength"`
	CodeType                      string              `xml:"codeType"`
	CodeAllowedLength             string              `xml:"codeAllowedLength,omitempty"`
	DefaultPresentation           string              `xml:"defaultPresentation"`
	Predefined                    Predefined          `xml:"predefined"`
	EditType                      string              `xml:"editType"`
	ChoiceMode                    string              `xml:"choiceMode"`
	DefaultObjectForm             string              `xml:"defaultObjectForm"`
	DefaultFolderForm             string              `xml:"defaultFolderForm"`
	DefaultListForm               string              `xml:"defaultListForm"`
	DefaultChoiceForm             string              `xml:"defaultChoiceForm"`
	DefaultFolderChoiceForm       string              `xml:"defaultFolderChoiceForm"`
	Attributes                    []Attribute         `xml:"attributes"`
	TabularSections               []TabularSection    `xml:"tabularSections"`
	Forms                         []Form              `xml:"forms"`
	Commands                      []Command           `xml:"commands"`
}

func (c *Catalog) Pack(path string) {
	packObjectMdo(path, c)
}
