package mdclasses

type ObjectTypeRef struct {
	TypeId      string `xml:"typeId,attr"`
	ValueTypeId string `xml:"valueTypeId,attr"`
}

type ObjectKeyValueType struct {
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}

func (v ObjectKeyValueType) IsNull() bool {
	return len(v.Key) == 0 && len(v.Value) == 0
}
