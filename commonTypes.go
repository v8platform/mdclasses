package mdclasses

//AttributeTypedValue - значение реквизитов
type AttributeTypedValue struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value"`
}

//StandardAttribute - стандартные реквизиты
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

//AttributeType - тип реквизита
type AttributeType struct {
	Types            string `xml:"types"`
	NumberQualifiers struct {
		Precision string `xml:"precision"`
	} `xml:"numberQualifiers"`
	StringQualifiers struct {
		Length string `xml:"length"`
	} `xml:"stringQualifiers"`

	DateQualifiers struct {
		DateFractions string `xml:"dateFractions"`
	} `xml:"dateQualifiers"`
}

//Attribute - Реквизиты
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

//TabularSectionProducedTypes - описание типов табличной части
type TabularSectionProducedTypes struct {
	ObjectType ObjectTypeRef `xml:"objectType"`
	RowType    ObjectTypeRef `xml:"rowType"`
}

//TabularSection - Табличные части
type TabularSection struct {
	MDOBaseType
	ProducedTypes      TabularSectionProducedTypes `xml:"producedTypes"`
	ToolTip            ObjectKeyValueType          `xml:"toolTip"`
	StandardAttributes []StandardAttribute         `xml:"standardAttributes"`
	Attributes         []Attribute                 `xml:"attributes"`
}

//MDOProducedTypes - описание типов объекта метаданных
type MDOProducedTypes struct {
	ObjectType    ObjectTypeRef `xml:"objectType"`
	RefType       ObjectTypeRef `xml:"refType"`
	SelectionType ObjectTypeRef `xml:"selectionType"`
	ListType      ObjectTypeRef `xml:"listType"`
	ManagerType   ObjectTypeRef `xml:"managerType"`
}
