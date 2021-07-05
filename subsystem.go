package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
	"strings"
)

type Subsystem struct {
	MDOBaseType
	SubsystemChildSubsystems
	XMLName                   xml.Name       `xml:"Subsystem"`
	IncludeHelpInContents     bool           `xml:"includeHelpInContents"`
	IncludeInCommandInterface bool           `xml:"includeInCommandInterface"`
	Subsystems                []*Subsystem   `xml:"-"`
	Content                   MDOTypeRefList `xml:"content"`
	ParentSubsystem           MDOTypeRef     `xml:"parentSubsystem"`
}

type SubsystemChildSubsystems struct {
	Subsystems []string `xml:"subsystems"`
}

func (conf *Subsystem) Unpack(cfg UnpackConfig) error {

	parentMDO := NewMDOTypeRef(SUBSYSTEM, conf.Name, conf.ParentSubsystem)

	var subsystems []MDOTypeRef

	for _, name := range conf.SubsystemChildSubsystems.Subsystems {

		subsystems = append(subsystems, NewMDOTypeRef(
			SUBSYSTEM,
			name,
			parentMDO,
		))

	}

	err := MDOTypeRefList(subsystems).Unpack(cfg, &conf.Subsystems)
	if err != nil {
		return err
	}

	//
	// for _, name := range conf.ContentNames {
	//
	// 	names := strings.Split(name, ".")
	//
	// 	if len(names) != 2 {
	// 		log.Warnf("Error parce subsystems: %s", name)
	// 		continue
	// 	}
	//
	// 	contentType := names[0]
	//
	// 	value := objectByType(contentType)
	//
	// 	// TODO Временно убрать, после реализации
	// 	if value == nil {
	// 		continue
	// 	}
	//
	// 	err := Unpack(cfg.WithName(name, contentType), value)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	conf.Content = append(conf.Content, value)
	// }

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
		log.Warnf("Unknown content type: %s", contentType)
	}

	return nil
}

func (c Subsystem) GetContents() MDOTypeRefList {
	list := MDOTypeRefList{}
	addRecurse(&c, &list)
	return list
}

func addRecurse(c *Subsystem, list *MDOTypeRefList) {
	addContents(c, list)
	for _, childSubsystem := range c.Subsystems {
		addRecurse(childSubsystem, list)
	}
}

func addContents(c *Subsystem, list *MDOTypeRefList) {
	for _, content := range c.Content {
		for _, v := range *list {
			if v.raw == content.raw {
				return
			}
		}
		*list = append(*list, content)
	}
}
