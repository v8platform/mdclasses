package mdclasses

import (
	"encoding/xml"
	"log"
	"strings"
)

type Subsystem struct {
	XMLName                   xml.Name    `xml:"Subsystem"`
	UUID                      string      `xml:"uuid,attr"`
	Name                      string      `xml:"name"`
	IncludeHelpInContents     bool        `xml:"includeHelpInContents"`
	IncludeInCommandInterface bool        `xml:"includeInCommandInterface"`
	SubsystemsNames           []string    `xml:"subsystems"`
	Subsystems                []Subsystem `xml:"-"`

	СontentNames []string      `xml:"content"`
	Content      []interface{} `xml:"-"`
	// Добавить другие обхекты
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

	for _, name := range conf.СontentNames {

		names := strings.Split(name, ".")

		if len(names) != 2 {
			log.Panic("Error")
		}

		contentType := names[0]

		value := objectByType(contentType)

		// TODO Временно убрать, после реализации
		if value == nil {
			continue
		}

		err := Unpack(cfg.WithName(name, contentType), value)
		if err != nil {
			return err
		}

		conf.Content = append(conf.Content, value)
	}

	return nil
}

func objectByType(contentType string) interface{} {

	switch strings.ToLower(contentType) {
	case "businessprocess":
		return nil
	case "catalog":
		return &Catalog{}
	case "document":
		return &Document{}
	default:
		log.Print("Unknown content type")
	}

	return nil
}