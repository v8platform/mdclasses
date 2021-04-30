package mdclasses

// CommonModule ...
type CommonModule struct {
	MDOBaseType
	Comment                   string `xml:"comment"`
	Global                    bool   `xml:"global"`
	ClientManagedApplication  bool   `xml:"clientManagedApplication"`
	Server                    bool   `xml:"server"`
	ExternalConnection        bool   `xml:"externalConnection"`
	ClientOrdinaryApplication bool   `xml:"clientOrdinaryApplication"`
	Client                    bool   `xml:"client"`
	ServerCall                bool   `xml:"server_call"`
	Privileged                bool   `xml:"privileged"`
	ReturnValuesReuse         string `xml:"returnValuesReuse"`
}
