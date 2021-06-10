package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
)
type ObjectTypeRef struct {
	TypeId      xml.Attr `xml:"typeId,attr"`
	ValueTypeId xml.Attr `xml:"valueTypeId,attr"`
}

type ObjectKeyValueType struct {
	Key   string `xml:"key,omitempty"`
	Value string `xml:"value,omitempty"`
}
