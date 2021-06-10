package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
)

//AttributeTypedValue - значение реквизитов
type AttributeTypedValue struct {
	Type  xml.Attr `xml:"type,attr"`
	Value string   `xml:"value"`
}

type TypeValue struct {
	Type xml.Attr `xml:"attr"`
}

//StandardAttribute - стандартные реквизиты
type StandardAttribute struct {
	DataHistory          string               `xml:"dataHistory"`
	Name                 string               `xml:"name"`
	Synonym              []ObjectKeyValueType `xml:"synonym,omitempty"`
	ToolTip              ObjectKeyValueType   `xml:"toolTip,omitempty"`
	FillFromFillingValue string               `xml:"fillFromFillingValue,omitempty"`
	FillValue            AttributeTypedValue  `xml:"fillValue,allowempty"`
	FillChecking         string               `xml:"fillChecking,omitempty"`
	FullTextSearch       string               `xml:"fullTextSearch"`
	MinValue             AttributeTypedValue  `xml:"minValue,allowempty"`
	MaxValue             AttributeTypedValue  `xml:"maxValue,allowempty"`
}

//AttributeType - тип реквизита
type NumberQualifiers struct {
	Precision string `xml:"precision,omitempty"`
}

type StringQualifiers struct {
	Length string `xml:"length,omitempty"`
}

type DateQualifiers struct {
	DateFractions string `xml:"dateFractions,omitempty"`
}

type AttributeType struct {
	Types            []string         `xml:"types"`
	NumberQualifiers NumberQualifiers `xml:"numberQualifiers,omitempty"`
	StringQualifiers StringQualifiers `xml:"stringQualifiers,allowempty"`
	DateQualifiers   DateQualifiers   `xml:"dateQualifiers,omitempty"`
}

//Attribute - Реквизиты
type Attribute struct {
	Uuid           string              `xml:"uuid,attr"`
	Name           string              `xml:"name"`
	Synonym        ObjectKeyValueType  `xml:"synonym"`
	Comment        string              `xml:"comment,omitempty"`
	Type           []AttributeType     `xml:"type"`
	ToolTip        ObjectKeyValueType  `xml:"toolTip"`
	MultiLine      string              `xml:"multiLine,omitempty"`
	MinValue       AttributeTypedValue `xml:"minValue"`
	MaxValue       AttributeTypedValue `xml:"maxValue"`
	FillChecking   string              `xml:"fillChecking,omitempty"`
	QuickChoice    string              `xml:"quickChoice,omitempty"`
	CreateOnInput  string              `xml:"createOnInput,omitempty"`
	FillValue      AttributeTypedValue `xml:"fillValue,omitempty"`
	DataHistory    string              `xml:"dataHistory"`
	Indexing       string              `xml:"indexing,omitempty"`
	FullTextSearch string              `xml:"fullTextSearch,omitempty"`
	Use            string              `xml:"use,omitempty"`
	Format         ObjectKeyValueType  `xml:"format,omitempty"`
	EditFormat     ObjectKeyValueType  `xml:"editFormat,omitempty"`
	Mask           string              `xml:"mask,omitempty"`
}

//TabularSectionProducedTypes - описание типов табличной части
type TabularSectionProducedTypes struct {
	ObjectType ObjectTypeRef `xml:"objectType"`
	RowType    ObjectTypeRef `xml:"rowType"`
}

//TabularSection - Табличные части
type TabularSection struct {
	MDOBaseType
	ProducedTypes      TabularSectionProducedTypes `xml:"producedTypes"`
	Name               string                      `xml:"name"`
	Synonym            []ObjectKeyValueType        `xml:"synonym,omitempty"`
	FillChecking       string                      `xml:"fillChecking,omitempty"`
	ToolTip            ObjectKeyValueType          `xml:"toolTip,omitempty"`
	StandardAttributes []StandardAttribute         `xml:"standardAttributes"`
	Attributes         []Attribute                 `xml:"attributes"`
}

//MDOProducedTypes - описание типов объекта метаданных
type MDOProducedTypes struct {
	ObjectType    ObjectTypeRef `xml:"objectType,allowempty"`
	RefType       ObjectTypeRef `xml:"refType,allowempty"`
	SelectionType ObjectTypeRef `xml:"selectionType,allowempty"`
	ListType      ObjectTypeRef `xml:"listType,allowempty"`
	ManagerType   ObjectTypeRef `xml:"managerType,allowempty"`
}

type Help struct {
	Pages struct {
		Lang string `xml:"lang"`
	} `xml:"pages"`
}
type Content struct {
	ID          string              `xml:"id,attr"`
	Name        string              `xml:"name"`
	Description string              `xml:"description"`
	Code        AttributeTypedValue `xml:"code"`
}

type Item struct {
	ID          string              `xml:"id,attr"`
	Name        string              `xml:"name"`
	Description string              `xml:"description"`
	Code        AttributeTypedValue `xml:"code"`
	IsFolder    string              `xml:"isFolder"`
	Content     []Content           `xml:"content"`
}

type Predefined struct {
	Items []Item `xml:"items"`
}


type Form struct {
	Uuid        string               `xml:"uuid,attr"`
	Name        string               `xml:"name"`
	Synonym     []ObjectKeyValueType `xml:"synonym"`
	Help        Help                 `xml:"help"`
	UsePurposes []string             `xml:"usePurposes"`
}

type Command struct {
	Uuid                 string               `xml:"uuid,attr"`
	Name                 string               `xml:"name"`
	Synonym              []ObjectKeyValueType `xml:"synonym"`
	Group                string               `xml:"group"`
	CommandParameterType Type                 `xml:"commandParameterType"`
	Representation       string               `xml:"representation"`
	Shortcut             string               `xml:"shortcut"`
}

type Type struct {
	Types []string `xml:"types"`
}