package mdclasses

import (
	"encoding/xml"
	"fmt"
)

type Subsystems []Subsystem

type Configuration struct {
	XMLName    xml.Name   `xml:"Configuration"`
	UUID       string     `xml:"uuid,attr"`
	Subsystems Subsystems `xml:"subsystems"`
}
type Subsystem struct {
	XMLName                   xml.Name   `xml:"Subsystem"`
	UUID                      string     `xml:"uuid,attr"`
	Name                      string     `xml:"name"`
	IncludeHelpInContents     bool       `xml:"includeHelpInContents"`
	IncludeInCommandInterface bool       `xml:"includeInCommandInterface"`
	Subsystems                Subsystems `xml:"subsystems"`
}

func (s *Subsystems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var filepath string
	if err := d.DecodeElement(&filepath, &start); err != nil {
		return err
	}

	fmt.Printf("Subsystem: %q\n", filepath)

	// читаем данные из файла
	data := `
<?xml version="1.0" encoding="UTF-8"?>
<mdclass:Subsystem xmlns:mdclass="http://g5.1c.ru/v8/dt/metadata/mdclass" uuid="3d00f7d6-e3b0-49cf-8093-e2e4f6ea2293">
  <name>ПерваяПодсистема</name>
  <synonym>
    <key>ru</key>
    <value>Первая подсистема</value>
  </synonym>
  <includeHelpInContents>true</includeHelpInContents>
  <includeInCommandInterface>true</includeInCommandInterface>

</mdclass:Subsystem>`

	subsystem := Subsystem{}
	err := xml.Unmarshal([]byte(data), &subsystem)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	// Собственно тут делаем распаковку подсистемы
	*s = append(*s, subsystem)

	return nil
}

func Parse(file string) error {

	cfg := &Configuration{}
	data := `
	<?xml version="1.0" encoding="UTF-8"?>
	<mdclass:Configuration xmlns:mdclass="http://g5.1c.ru/v8/dt/metadata/mdclass" uuid="46c7c1d0-b04d-4295-9b04-ae3207c18d29">
  		 <containedObjects classId="9cd510cd-abfc-11d4-9434-004095e12fc7" objectId="6b9cfff9-b36e-4dc0-947d-9ae7d8eafbff"/>
    <containedObjects classId="9fcd25a0-4822-11d4-9414-008048da11f9" objectId="7cf1638d-e924-4cda-9c28-e32ac23e9bd7"/>
    <containedObjects classId="e3687481-0a87-462c-a166-9f34594f9bba" objectId="3982f5e9-e852-4ca4-bac4-0dcb8fc6119a"/>
    <containedObjects classId="9de14907-ec23-4a07-96f0-85521cb6b53b" objectId="fe20ab46-0f06-4ed5-90a6-9190787ac88a"/>
    <containedObjects classId="51f2d5d8-ea4d-4064-8892-82951750031e" objectId="25eee44c-dea9-490b-a98a-7bdb81f3cdc3"/>
    <containedObjects classId="e68182ea-4237-4383-967f-90c1e3370bc7" objectId="acb89856-4d76-4913-a9d7-e352e9d1873e"/>
    <configurationExtensionCompatibilityMode>8.3.10</configurationExtensionCompatibilityMode>
    <defaultRunMode>ManagedApplication</defaultRunMode>
    <usePurposes>PersonalComputer</usePurposes>
    <usePurposes>MobileDevice</usePurposes>
    <scriptVariant>Russian</scriptVariant>
    <useManagedFormInOrdinaryApplication>true</useManagedFormInOrdinaryApplication>
    <useOrdinaryFormInManagedApplication>true</useOrdinaryFormInManagedApplication>
    <defaultRoles>Role.Роль1</defaultRoles>
    
		<subsystems>Subsystem.ПерваяПодсистема</subsystems>
    	<subsystems>Subsystem.ВтораяПодсистема</subsystems>
	</mdclass:Configuration>
	`
	err := xml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	fmt.Printf("XMLName: %#v\n", cfg.XMLName)
	fmt.Printf("UUID: %#v\n", cfg.UUID)
	fmt.Printf("Subsystems: %v\n", cfg.Subsystems)
	return nil
}
