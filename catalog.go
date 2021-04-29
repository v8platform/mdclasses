package mdclasses

import "encoding/xml"

type Catalog struct {
	MDObjectBase
	XMLName xml.Name `xml:"Catalog"`

	Xsi  string `xml:"xsi,attr"`
	Core string `xml:"core,attr"`

	ProducedTypes struct {
		ObjectType    ObjectTypeRef `xml:"objectType"`
		RefType       ObjectTypeRef `xml:"refType"`
		SelectionType ObjectTypeRef `xml:"selectionType"`
		ListType      ObjectTypeRef `xml:"listType"`
		ManagerType   ObjectTypeRef `xml:"managerType"`
	} `xml:"producedTypes"`
	Name                          string             `xml:"name"`
	Synonym                       ObjectKeyValueType `xml:"synonym"`
	InputByString                 string             `xml:"inputByString"`
	FullTextSearchOnInputByString string             `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []struct {
		DataHistory string `xml:"dataHistory"`
		Name        string `xml:"name"`
		FillValue   struct {
			Type  string `xml:"type,attr"`
			Value string `xml:"value"`
		} `xml:"fillValue"`
		FullTextSearch string `xml:"fullTextSearch"`
		MinValue       struct {
			Type string `xml:"type,attr"`
		} `xml:"minValue"`
		MaxValue struct {
			Type string `xml:"type,attr"`
		} `xml:"maxValue"`
		FillFromFillingValue string             `xml:"fillFromFillingValue"`
		FillChecking         string             `xml:"fillChecking"`
		ToolTip              ObjectKeyValueType `xml:"toolTip"`
	} `xml:"standardAttributes"`
	CreateOnInput       string             `xml:"createOnInput"`
	DataLockControlMode string             `xml:"dataLockControlMode"`
	FullTextSearch      string             `xml:"fullTextSearch"`
	ObjectPresentation  ObjectKeyValueType `xml:"objectPresentation"`
	Hierarchical        string             `xml:"hierarchical"`
	LevelCount          string             `xml:"levelCount"`
	FoldersOnTop        string             `xml:"foldersOnTop"`
	DescriptionLength   string             `xml:"descriptionLength"`
	CodeType            string             `xml:"codeType"`
	CodeAllowedLength   string             `xml:"codeAllowedLength"`
	DefaultPresentation string             `xml:"defaultPresentation"`
	EditType            string             `xml:"editType"`
	ChoiceMode          string             `xml:"choiceMode"`
	Attributes          []struct {
		Uuid    string             `xml:"uuid,attr"`
		Name    string             `xml:"name"`
		Synonym ObjectKeyValueType `xml:"synonym"`
		Comment string             `xml:"comment"`
		Type    struct {
			Types            string `xml:"types"`
			DateQualifiers   string `xml:"dateQualifiers"`
			NumberQualifiers struct {
				Precision string `xml:"precision"`
			} `xml:"numberQualifiers"`
			StringQualifiers struct {
				Length string `xml:"length"`
			} `xml:"stringQualifiers"`
		} `xml:"type"`
		ToolTip  ObjectKeyValueType `xml:"toolTip"`
		MinValue struct {
			Type string `xml:"type,attr"`
		} `xml:"minValue"`
		MaxValue struct {
			Type string `xml:"type,attr"`
		} `xml:"maxValue"`
		FillChecking  string `xml:"fillChecking"`
		QuickChoice   string `xml:"quickChoice"`
		CreateOnInput string `xml:"createOnInput"`
		FillValue     struct {
			Type  string `xml:"type,attr"`
			Value string `xml:"value"`
		} `xml:"fillValue"`
		Indexing       string             `xml:"indexing"`
		FullTextSearch string             `xml:"fullTextSearch"`
		Use            string             `xml:"use"`
		DataHistory    string             `xml:"dataHistory"`
		Format         ObjectKeyValueType `xml:"format"`
		EditFormat     ObjectKeyValueType `xml:"editFormat"`
		Mask           string             `xml:"mask"`
		MultiLine      string             `xml:"multiLine"`
	} `xml:"attributes"`
	TabularSections []struct {
		Uuid          string `xml:"uuid,attr"`
		ProducedTypes struct {
			ObjectType ObjectTypeRef `xml:"objectType"`
			RowType    ObjectTypeRef `xml:"rowType"`
		} `xml:"producedTypes"`
		Name               string             `xml:"name"`
		Synonym            ObjectKeyValueType `xml:"synonym"`
		ToolTip            ObjectKeyValueType `xml:"toolTip"`
		StandardAttributes struct {
			DataHistory string `xml:"dataHistory"`
			Name        string `xml:"name"`
			FillValue   struct {
				Type string `xml:"type,attr"`
			} `xml:"fillValue"`
			FullTextSearch string `xml:"fullTextSearch"`
			MinValue       struct {
				Type string `xml:"type,attr"`
			} `xml:"minValue"`
			MaxValue struct {
				Type string `xml:"type,attr"`
			} `xml:"maxValue"`
		} `xml:"standardAttributes"`
		Attributes []struct {
			Uuid    string             `xml:"uuid,attr"`
			Name    string             `xml:"name"`
			Synonym ObjectKeyValueType `xml:"synonym"`
			Type    struct {
				Types            string `xml:"types"`
				DateQualifiers   string `xml:"dateQualifiers"`
				StringQualifiers struct {
					Length string `xml:"length"`
				} `xml:"stringQualifiers"`
			} `xml:"type"`
			ToolTip  ObjectKeyValueType `xml:"toolTip"`
			MinValue struct {
				Type string `xml:"type,attr"`
			} `xml:"minValue"`
			MaxValue struct {
				Type string `xml:"type,attr"`
			} `xml:"maxValue"`
			DataHistory    string `xml:"dataHistory"`
			FullTextSearch string `xml:"fullTextSearch"`
			Comment        string `xml:"comment"`
			Indexing       string `xml:"indexing"`
		} `xml:"attributes"`
	} `xml:"tabularSections"`
}
