package mdclasses

import "github.com/v8platform/mdclasses/encoding/xml"

type ProducedTypes struct {
	SelectionType     ObjectTypeRef `xml:"selectionType"`
	ListType          ObjectTypeRef `xml:"listType"`
	ManagerType       ObjectTypeRef `xml:"managerType"`
	RecordSetType     ObjectTypeRef `xml:"recordSetType"`
	RecordKeyType     ObjectTypeRef `xml:"recordKeyType"`
	RecordType        ObjectTypeRef `xml:"recordType"`
	RecordManagerType ObjectTypeRef `xml:"recordManagerType"`
}

type Resource struct {
	Uuid    string              `xml:"uuid,attr"`
	Name    string              `xml:"name"`
	Synonym AttributeTypedValue `xml:"synonym"`
	Type    struct {
		Types []string `xml:"types"`
	} `xml:"type"`
	ToolTip  AttributeTypedValue `xml:"toolTip"`
	MinValue struct {
		Type string `xml:"type,attr"`
	} `xml:"minValue"`
	MaxValue struct {
		Type string `xml:"type,attr"`
	} `xml:"maxValue"`
	FillChecking     string `xml:"fillChecking"`
	ChoiceParameters struct {
		Name  string              `xml:"name"`
		Value AttributeTypedValue `xml:"value"`
	} `xml:"choiceParameters"`
	FullTextSearch       string              `xml:"fullTextSearch"`
	DataHistory          string              `xml:"dataHistory"`
	FillFromFillingValue string              `xml:"fillFromFillingValue"`
	FillValue            AttributeTypedValue `xml:"fillValue"`
}

type Dimension struct {
	Uuid    string             `xml:"uuid,attr"`
	Name    string             `xml:"name"`
	Synonym ObjectKeyValueType `xml:"synonym"`
	Type    struct {
		Types string `xml:"types"`
	} `xml:"type"`
	ToolTip  ObjectKeyValueType `xml:"toolTip"`
	MinValue struct {
		Type string `xml:"type,attr"`
	} `xml:"minValue"`
	MaxValue struct {
		Type string `xml:"type,attr"`
	} `xml:"maxValue"`
	FullTextSearch       string              `xml:"fullTextSearch"`
	DataHistory          string              `xml:"dataHistory"`
	FillFromFillingValue string              `xml:"fillFromFillingValue"`
	FillValue            AttributeTypedValue `xml:"fillValue"`
	Master               string              `xml:"master"`
	MainFilter           string              `xml:"mainFilter"`
}

type InformationRegister struct {
	XMLName                        xml.Name            `xml:"InformationRegister"`
	Xsi                            string              `xml:"xsi,attr"`
	Core                           string              `xml:"core,attr"`
	Mdclass                        string              `xml:"mdclass,attr"`
	Uuid                           string              `xml:"uuid,attr"`
	ProducedTypes                  ProducedTypes       `xml:"producedTypes"`
	Name                           string              `xml:"name"`
	Synonym                        ObjectKeyValueType  `xml:"synonym"`
	UseStandardCommands            string              `xml:"useStandardCommands"`
	DefaultRecordForm              string              `xml:"defaultRecordForm"`
	DefaultListForm                string              `xml:"defaultListForm"`
	StandardAttributes             []StandardAttribute `xml:"standardAttributes"`
	InformationRegisterPeriodicity string              `xml:"informationRegisterPeriodicity"`
	MainFilterOnPeriod             string              `xml:"mainFilterOnPeriod"`
	Help                           Help                `xml:"help"`
	RecordPresentation             ObjectKeyValueType  `xml:"recordPresentation"`
	ListPresentation               ObjectKeyValueType  `xml:"listPresentation"`
	Explanation                    ObjectKeyValueType  `xml:"explanation"`
	Resources                      []Resource          `xml:"resources"`
	Dimensions                     []Dimension         `xml:"dimensions"`
	Forms                          []Form              `xml:"forms"`
}
