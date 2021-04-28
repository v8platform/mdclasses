package mdclasses

import "encoding/xml"

type Subsystem struct {
	XMLName                   xml.Name    `xml:"Subsystem"`
	UUID                      string      `xml:"uuid,attr"`
	Name                      string      `xml:"name"`
	IncludeHelpInContents     bool        `xml:"includeHelpInContents"`
	IncludeInCommandInterface bool        `xml:"includeInCommandInterface"`
	SubsystemsNames           []string    `xml:"subsystems"`
	Subsystems                []Subsystem `xml:"-"`
}

func (conf *Subsystem) Unpack(cfg UnpackConfig) error {

	for _, name := range conf.SubsystemsNames {

		subsystem := Subsystem{}
		err := Unpack(cfg.WithName(name, "Subsystem"), &subsystem)
		if err != nil {
			return err
		}

		conf.Subsystems = append(conf.Subsystems, subsystem)
	}

	return nil
}
