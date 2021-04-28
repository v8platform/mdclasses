package mdclasses

import "encoding/xml"

type Catalog struct {
	XMLName       xml.Name `xml:"Catalog"`
	Text          string   `xml:",chardata"`
	Xsi           string   `xml:"xsi,attr"`
	Core          string   `xml:"core,attr"`
	Mdclass       string   `xml:"mdclass,attr"`
	Uuid          string   `xml:"uuid,attr"`
	ProducedTypes struct {
		Text       string `xml:",chardata"`
		ObjectType struct {
			Text        string `xml:",chardata"`
			TypeId      string `xml:"typeId,attr"`
			ValueTypeId string `xml:"valueTypeId,attr"`
		} `xml:"objectType"`
		RefType struct {
			Text        string `xml:",chardata"`
			TypeId      string `xml:"typeId,attr"`
			ValueTypeId string `xml:"valueTypeId,attr"`
		} `xml:"refType"`
		SelectionType struct {
			Text        string `xml:",chardata"`
			TypeId      string `xml:"typeId,attr"`
			ValueTypeId string `xml:"valueTypeId,attr"`
		} `xml:"selectionType"`
		ListType struct {
			Text        string `xml:",chardata"`
			TypeId      string `xml:"typeId,attr"`
			ValueTypeId string `xml:"valueTypeId,attr"`
		} `xml:"listType"`
		ManagerType struct {
			Text        string `xml:",chardata"`
			TypeId      string `xml:"typeId,attr"`
			ValueTypeId string `xml:"valueTypeId,attr"`
		} `xml:"managerType"`
	} `xml:"producedTypes"`
	Name    string `xml:"name"`
	Synonym struct {
		Text  string `xml:",chardata"`
		Key   string `xml:"key"`
		Value string `xml:"value"`
	} `xml:"synonym"`
	InputByString                 string `xml:"inputByString"`
	FullTextSearchOnInputByString string `xml:"fullTextSearchOnInputByString"`
	StandardAttributes            []struct {
		Text        string `xml:",chardata"`
		DataHistory string `xml:"dataHistory"`
		Name        string `xml:"name"`
		FillValue   struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Value string `xml:"value"`
		} `xml:"fillValue"`
		FullTextSearch string `xml:"fullTextSearch"`
		MinValue       struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"minValue"`
		MaxValue struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"maxValue"`
		FillFromFillingValue string `xml:"fillFromFillingValue"`
		FillChecking         string `xml:"fillChecking"`
		ToolTip              struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"toolTip"`
	} `xml:"standardAttributes"`
	CreateOnInput       string `xml:"createOnInput"`
	DataLockControlMode string `xml:"dataLockControlMode"`
	FullTextSearch      string `xml:"fullTextSearch"`
	ObjectPresentation  struct {
		Text  string `xml:",chardata"`
		Key   string `xml:"key"`
		Value string `xml:"value"`
	} `xml:"objectPresentation"`
	Hierarchical        string `xml:"hierarchical"`
	LevelCount          string `xml:"levelCount"`
	FoldersOnTop        string `xml:"foldersOnTop"`
	DescriptionLength   string `xml:"descriptionLength"`
	CodeType            string `xml:"codeType"`
	CodeAllowedLength   string `xml:"codeAllowedLength"`
	DefaultPresentation string `xml:"defaultPresentation"`
	EditType            string `xml:"editType"`
	ChoiceMode          string `xml:"choiceMode"`
	Attributes          []struct {
		Text    string `xml:",chardata"`
		Uuid    string `xml:"uuid,attr"`
		Name    string `xml:"name"`
		Synonym struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"synonym"`
		Comment string `xml:"comment"`
		Type    struct {
			Text             string `xml:",chardata"`
			Types            string `xml:"types"`
			DateQualifiers   string `xml:"dateQualifiers"`
			NumberQualifiers struct {
				Text      string `xml:",chardata"`
				Precision string `xml:"precision"`
			} `xml:"numberQualifiers"`
			StringQualifiers struct {
				Text   string `xml:",chardata"`
				Length string `xml:"length"`
			} `xml:"stringQualifiers"`
		} `xml:"type"`
		ToolTip struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"toolTip"`
		MinValue struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"minValue"`
		MaxValue struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"maxValue"`
		FillChecking  string `xml:"fillChecking"`
		QuickChoice   string `xml:"quickChoice"`
		CreateOnInput string `xml:"createOnInput"`
		FillValue     struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Value string `xml:"value"`
		} `xml:"fillValue"`
		Indexing       string `xml:"indexing"`
		FullTextSearch string `xml:"fullTextSearch"`
		Use            string `xml:"use"`
		DataHistory    string `xml:"dataHistory"`
		Format         struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"format"`
		EditFormat struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"editFormat"`
		Mask      string `xml:"mask"`
		MultiLine string `xml:"multiLine"`
	} `xml:"attributes"`
	TabularSections []struct {
		Text          string `xml:",chardata"`
		Uuid          string `xml:"uuid,attr"`
		ProducedTypes struct {
			Text       string `xml:",chardata"`
			ObjectType struct {
				Text        string `xml:",chardata"`
				TypeId      string `xml:"typeId,attr"`
				ValueTypeId string `xml:"valueTypeId,attr"`
			} `xml:"objectType"`
			RowType struct {
				Text        string `xml:",chardata"`
				TypeId      string `xml:"typeId,attr"`
				ValueTypeId string `xml:"valueTypeId,attr"`
			} `xml:"rowType"`
		} `xml:"producedTypes"`
		Name    string `xml:"name"`
		Synonym struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"synonym"`
		ToolTip struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"toolTip"`
		StandardAttributes struct {
			Text        string `xml:",chardata"`
			DataHistory string `xml:"dataHistory"`
			Name        string `xml:"name"`
			FillValue   struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"fillValue"`
			FullTextSearch string `xml:"fullTextSearch"`
			MinValue       struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"minValue"`
			MaxValue struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"maxValue"`
		} `xml:"standardAttributes"`
		Attributes []struct {
			Text    string `xml:",chardata"`
			Uuid    string `xml:"uuid,attr"`
			Name    string `xml:"name"`
			Synonym struct {
				Text  string `xml:",chardata"`
				Key   string `xml:"key"`
				Value string `xml:"value"`
			} `xml:"synonym"`
			Type struct {
				Text             string `xml:",chardata"`
				Types            string `xml:"types"`
				DateQualifiers   string `xml:"dateQualifiers"`
				StringQualifiers struct {
					Text   string `xml:",chardata"`
					Length string `xml:"length"`
				} `xml:"stringQualifiers"`
			} `xml:"type"`
			ToolTip struct {
				Text  string `xml:",chardata"`
				Key   string `xml:"key"`
				Value string `xml:"value"`
			} `xml:"toolTip"`
			MinValue struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"minValue"`
			MaxValue struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"maxValue"`
			DataHistory    string `xml:"dataHistory"`
			FullTextSearch string `xml:"fullTextSearch"`
			Comment        string `xml:"comment"`
			Indexing       string `xml:"indexing"`
		} `xml:"attributes"`
	} `xml:"tabularSections"`
}
