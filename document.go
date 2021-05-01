package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Form struct {
	MDOBaseType
	UsePurposes []string `xml:"usePurposes"`
}

type Command struct {
	MDOBaseType
	Group          string `xml:"group"`
	Representation string `xml:"representation"`
}

type Document struct {
	MDOBaseType
	XMLName       xml.Name         `xml:"Document"`
	ProducedTypes MDOProducedTypes `xml:"producedTypes"`

	UseStandardCommands           bool                `xml:"useStandardCommands"`
	InputByString                 MDOTypeRef          `xml:"inputByString"`
	FullTextSearchOnInputByString string              `xml:"fullTextSearchOnInputByString"`
	CreateOnInput                 string              `xml:"createOnInput"`
	DataLockControlMode           string              `xml:"dataLockControlMode"`
	FullTextSearch                string              `xml:"fullTextSearch"`
	NumberType                    string              `xml:"numberType"`
	NumberLength                  string              `xml:"numberLength"`
	NumberAllowedLength           string              `xml:"numberAllowedLength"`
	CheckUnique                   bool                `xml:"checkUnique"`
	Autonumbering                 bool                `xml:"autonumbering"`
	DefaultObjectForm             MDOTypeRef          `xml:"defaultObjectForm"`
	DefaultListForm               MDOTypeRef          `xml:"defaultListForm"`
	DefaultChoiceForm             MDOTypeRef          `xml:"defaultChoiceForm"`
	RegisterRecords               MDOTypeRefList      `xml:"registerRecords"`
	PostInPrivilegedMode          bool                `xml:"postInPrivilegedMode"`
	UnpostInPrivilegedMode        bool                `xml:"unpostInPrivilegedMode"`
	StandardAttributes            []StandardAttribute `xml:"standardAttributes"`
	Attributes                    []Attribute         `xml:"attributes"`
	Forms                         []Form              `xml:"forms"`
	TabularSections               []TabularSection    `xml:"tabularSections"`
	Commands                      []Command           `xml:"commands"`
}
