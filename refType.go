package mdclasses

type ObjectTypeRef struct {
	TypeId      string `xml:"typeId,attr"`
	ValueTypeId string `xml:"valueTypeId,attr"`
}

type ValueTypeRef struct {
	Key   string `xml:"key,,omitempty"`
	Value string `xml:"value,omitempty"`
}
