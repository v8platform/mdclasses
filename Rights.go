package mdclasses

type Right struct {
	Name                   string `xml:"name"`
	Value                  string `xml:"value"`
	RestrictionByCondition struct {
		Condition string `xml:"condition"`
	} `xml:"restrictionByCondition"`
}

type Object struct {
	Name  string  `xml:"name"`
	Right []Right `xml:"right"`
}

type Rights struct {
	XMLName                         xml.Name `xml:"Rights"`
	Xsi                             string   `xml:"xsi,attr"`
	Xmlns                           string   `xml:"xmlns,attr"`
	Type                            string   `xml:"type,attr"`
	SetForNewObjects                string   `xml:"setForNewObjects"`
	SetForAttributesByDefault       string   `xml:"setForAttributesByDefault"`
	IndependentRightsOfChildObjects string   `xml:"independentRightsOfChildObjects"`
	Object                          []Object `xml:"object"`
}

func (conf *Rights) Unpack(cfg UnpackConfig) error {
	return nil
}