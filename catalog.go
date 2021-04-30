package mdclasses

import "encoding/xml"

type AttributeTypedValue struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value"`
}

type StandardAttribute struct {
	Name                 string              `xml:"name"`
	DataHistory          string              `xml:"dataHistory"`
	FillValue            AttributeTypedValue `xml:"fillValue"`
	FullTextSearch       string              `xml:"fullTextSearch"`
	MinValue             AttributeTypedValue `xml:"minValue"`
	MaxValue             AttributeTypedValue `xml:"maxValue"`
	FillFromFillingValue string              `xml:"fillFromFillingValue"`
	FillChecking         string              `xml:"fillChecking"`
	ToolTip              ObjectKeyValueType  `xml:"toolTip"`
}

type AttributeType struct {
	Types            string `xml:"types"`
	DateQualifiers   string `xml:"dateQualifiers"`
	NumberQualifiers struct {
		Precision string `xml:"precision"`
	} `xml:"numberQualifiers"`
	StringQualifiers struct {
		Length string `xml:"length"`
	} `xml:"stringQualifiers"`
}

type Attribute struct {
	Uuid           string              `xml:"uuid,attr"`
	Name           string              `xml:"name"`
	Synonym        ObjectKeyValueType  `xml:"synonym"`
	Comment        string              `xml:"comment"`
	Type           AttributeType       `xml:"type"`
	ToolTip        ObjectKeyValueType  `xml:"toolTip"`
	MinValue       AttributeTypedValue `xml:"minValue"`
	MaxValue       AttributeTypedValue `xml:"maxValue"`
	FillChecking   string              `xml:"fillChecking"`
	QuickChoice    string              `xml:"quickChoice"`
	CreateOnInput  string              `xml:"createOnInput"`
	FillValue      AttributeTypedValue `xml:"fillValue"`
	Indexing       string              `xml:"indexing"`
	FullTextSearch string              `xml:"fullTextSearch"`
	Use            string              `xml:"use"`
	DataHistory    string              `xml:"dataHistory"`
	Format         ObjectKeyValueType  `xml:"format"`
	EditFormat     ObjectKeyValueType  `xml:"editFormat"`
	Mask           string              `xml:"mask"`
	MultiLine      string              `xml:"multiLine"`
}

type Catalog struct {
	MDOBaseType
	XMLName xml.Name `xml:"Catalog"`

	ProducedTypes struct {
		ObjectType    ObjectTypeRef `xml:"objectType"`
		RefType       ObjectTypeRef `xml:"refType"`
		SelectionType ObjectTypeRef `xml:"selectionType"`
		ListType      ObjectTypeRef `xml:"listType"`
		ManagerType   ObjectTypeRef `xml:"managerType"`
	} `xml:"producedTypes"`
	Name                          string              `xml:"name"`
	Synonym                       ObjectKeyValueType  `xml:"synonym"`
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
	TabularSections               []struct {
		MDOBaseType
		ProducedTypes struct {
			ObjectType ObjectTypeRef `xml:"objectType"`
			RowType    ObjectTypeRef `xml:"rowType"`
		} `xml:"producedTypes"`
		ToolTip            ObjectKeyValueType  `xml:"toolTip"`
		StandardAttributes []StandardAttribute `xml:"standardAttributes"`
		Attributes         []Attribute         `xml:"attributes"`
	} `xml:"tabularSections"`
}
