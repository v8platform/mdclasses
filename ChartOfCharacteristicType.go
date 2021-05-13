package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type Predefined struct {
	Items []struct {
		ID          string `xml:"id,attr"`
		Name        string `xml:"name"`
		Description string `xml:"description"`
		Code        string `xml:"code"`
		Type        struct {
			Types []string `xml:"types"`
		} `xml:"type"`
	} `xml:"items"`
}

type ChartOfCharacteristicType struct {
	XMLName                       xml.Name            `xml:"ChartOfCharacteristicTypes"`
	Xsi                           string              `xml:"xsi,attr"`
	Core                          string              `xml:"core,attr"`
	Mdclass                       string              `xml:"mdclass,attr"`
	Uuid                          string              `xml:"uuid,attr"`
	ProducedTypes                 MDOProducedTypes    `xml:"producedTypes"`
	Name                          string              `xml:"name"`
	Synonym                       ObjectKeyValueType  `xml:"synonym"`
	UseStandardCommands           string              `xml:"useStandardCommands"`
	InputByString                 []string            `xml:"inputByString"`
	FullTextSearchOnInputByString string              `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []StandardAttribute `xml:"standardAttributes"`
	Characteristics               []struct {
		CharacteristicTypes  string              `xml:"characteristicTypes"`
		KeyField             string              `xml:"keyField"`
		TypesFilterField     string              `xml:"typesFilterField"`
		TypesFilterValue     AttributeTypedValue `xml:"typesFilterValue"`
		CharacteristicValues string              `xml:"characteristicValues"`
		ObjectField          string              `xml:"objectField"`
		TypeField            string              `xml:"typeField"`
		ValueField           string              `xml:"valueField"`
	} `xml:"characteristics"`
	CreateOnInput  string `xml:"createOnInput"`
	Help           Help   `xml:"help"`
	FullTextSearch string `xml:"fullTextSearch"`
	Type           struct {
		Types []string `xml:"types"`
	} `xml:"type"`
	Hierarchical         string         `xml:"hierarchical"`
	FoldersOnTop         string         `xml:"foldersOnTop"`
	CodeLength           string         `xml:"codeLength"`
	DescriptionLength    string         `xml:"descriptionLength"`
	CheckUnique          string         `xml:"checkUnique"`
	Autonumbering        string         `xml:"autonumbering"`
	DefaultPresentation  string         `xml:"defaultPresentation"`
	Predefined           Predefined     `xml:"predefined"`
	EditType             string         `xml:"editType"`
	ChoiceMode           string         `xml:"choiceMode"`
	ChoiceHistoryOnInput string         `xml:"choiceHistoryOnInput"`
	DefaultObjectForm    string         `xml:"defaultObjectForm"`
	DefaultFolderForm    string         `xml:"defaultFolderForm"`
	DefaultListForm      string         `xml:"defaultListForm"`
	DefaultChoiceForm    string         `xml:"defaultChoiceForm"`
	Attributes           []Attribute    `xml:"attributes"`
	TabularSections      TabularSection `xml:"tabularSections"`
	Forms                []Form         `xml:"forms"`
}
