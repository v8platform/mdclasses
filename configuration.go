package mdclasses

import (
	"github.com/khorevaa/logos"
	"github.com/v8platform/mdclasses/encoding/xml"
)

var log = logos.New("github.com/v8platform/mdclasses").Sugar()

type Mdclass struct {
	xml.Name
	value string
}

type Name struct {
	xml.Name
	value string
}

type Configuration struct {
	XMLName xml.Name `xml:"mdclass Configuration"`
	Mdclass xml.Attr `xml:"mdclass,attr"`
	Uuid    string   `xml:"uuid,attr"`
	Name    string   `xml:"name"`
	Synonym ObjectKeyValueType   `xml:"synonym"`
	ContainedObjects []struct {
		ClassId  string `xml:"classId,attr,allowempty"`
		ObjectId string `xml:"objectId,attr,allowempty"`
	} `xml:"containedObjects,allowempty"`
	ConfigurationExtensionCompatibilityMode string   `xml:"configurationExtensionCompatibilityMode"`
	DefaultRunMode                          string   `xml:"defaultRunMode"`
	UsePurposes                             string   `xml:"usePurposes"`
	ScriptVariant                           string   `xml:"scriptVariant"`
	DefaultRoles                            []string `xml:"defaultRoles"`
	Vendor                                  string   `xml:"vendor"`
	Version                                 string   `xml:"version"`
	UpdateCatalogAddress                    string   `xml:"updateCatalogAddress,omitempty"`
	UseManagedFormInOrdinaryApplication     string   `xml:"useManagedFormInOrdinaryApplication"`
	UseOrdinaryFormInManagedApplication     string   `xml:"useOrdinaryFormInManagedApplication"`
	ReportsVariantsStorage                  string   `xml:"reportsVariantsStorage"`
	DefaultReportForm                       string   `xml:"defaultReportForm"`
	DefaultReportVariantForm                string   `xml:"defaultReportVariantForm"`
	DefaultReportSettingsForm               string   `xml:"defaultReportSettingsForm"`
	DefaultSearchForm                       string   `xml:"defaultSearchForm"`
	UsedMobileApplicationFunctionalities    struct {
		Functionality []struct {
			Functionality string `xml:"functionality,omitempty"`
			Use           string `xml:"use"`
		} `xml:"functionality"`
	} `xml:"usedMobileApplicationFunctionalities"`
	MainSectionPicture 				string 				`xml:"mainSectionPicture,allowempty"`
	DefaultLanguage    				string 				`xml:"defaultLanguage"`
	BriefInformation   				ObjectKeyValueType `xml:"briefInformation"`
	DetailedInformation 			ObjectKeyValueType `xml:"detailedInformation"`
	Splash    						string 				`xml:"splash,omitempty"`
	Copyright 						ObjectKeyValueType `xml:"copyright"`
	VendorInformationAddress 		ObjectKeyValueType `xml:"vendorInformationAddress"`
	ConfigurationInformationAddress ObjectKeyValueType `xml:"configurationInformationAddress,omitempty"`
	DataLockControlMode             string             `xml:"dataLockControlMode"`
	ObjectAutonumerationMode        string             `xml:"objectAutonumerationMode"`
	ModalityUseMode                 string             `xml:"modalityUseMode"`
	InterfaceCompatibilityMode      string             `xml:"interfaceCompatibilityMode"`
	CompatibilityMode               string             `xml:"compatibilityMode"`
	Languages                       struct {
		Uuid         string             `xml:"uuid,attr"`
		Name         string             `xml:"name"`
		Synonym      ObjectKeyValueType `xml:"synonym"`
		LanguageCode string             `xml:"languageCode"`
	} `xml:"languages"`

	Subsystems      []Subsystem `xml:"-"`
	SubsystemsNames []string    `xml:"subsystems"`

	StyleItems                  []string  `xml:"styleItems"`
	CommonPictures              []string  `xml:"commonPictures"`
	SessionParameters           []string  `xml:"sessionParameters"`
	Roles                       []string  `xml:"roles"`
	CommonTemplates             []string  `xml:"commonTemplates"`
	FilterCriteria              []string  `xml:"filterCriteria"`
	CommonModules               []string  `xml:"commonModules"`
	CommonAttributes            []string  `xml:"commonAttributes"`
	ExchangePlans               []string  `xml:"exchangePlans"`
	XDTOPackages                []string  `xml:"xDTOPackages"`
	WebServices                 []string  `xml:"webServices"`
	HttpServices                []string  `xml:"httpServices"`
	WsReferences                []string  `xml:"wsReferences"`
	EventSubscriptions          []string  `xml:"eventSubscriptions"`
	ScheduledJobs               []string  `xml:"scheduledJobs"`
	SettingsStorages            []string  `xml:"settingsStorages"`
	FunctionalOptions           []string  `xml:"functionalOptions"`
	FunctionalOptionsParameters []string  `xml:"functionalOptionsParameters"`
	DefinedTypes                []string  `xml:"definedTypes"`
	CommonCommands              []string  `xml:"commonCommands"`
	CommandGroups               []string  `xml:"commandGroups"`
	Constants                   []string  `xml:"constants"`
	CommonForms                 []string  `xml:"commonForms"`
	CatalogsNames               []string  `xml:"catalogs"`
	Catalogs                    []Catalog `xml:"-"`
	Documents                   []string  `xml:"documents"`
	DocumentNumerators          []string  `xml:"documentNumerators"`
	DocumentJournals            []string  `xml:"documentJournals"`
	Enums                       []string  `xml:"enums"`
	Reports                     []string  `xml:"reports"`
	DataProcessors              []string  `xml:"dataProcessors"`
	InformationRegisters        []string  `xml:"informationRegisters"`
	AccumulationRegisters       []string  `xml:"accumulationRegisters"`
	ChartsOfCharacteristicTypes []string  `xml:"chartsOfCharacteristicTypes"`
	BusinessProcesses           []string  `xml:"businessProcesses"`
	Tasks                       []string  `xml:"tasks"`
}

const ConfigurationFile = "Configuration.mdo"

func (conf *Name) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	panic("implement me")
}

func (conf *Mdclass) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	//pp.Println(name)
	return xml.Attr{
		conf.Name,
		conf.value,
	}, nil
}

func (conf *Mdclass) UnmarshalXMLAttr(attr xml.Attr) error {

	conf.Name = attr.Name
	conf.value = attr.Value

	return nil
}

func (conf *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	//pp.Println(name)
	return xml.Attr{
		conf.Name,
		conf.value,
	}, nil
}

func (conf *Name) UnmarshalXMLAttr(attr xml.Attr) error {

	conf.Name = attr.Name
	conf.value = attr.Value

	return nil
}

func (conf *Configuration) Unpack(cfg UnpackConfig) error {

	for _, name := range conf.SubsystemsNames {

		subsystem := Subsystem{}
		err := Unpack(cfg.WithName(name, "Subsystem"), &subsystem)
		if err != nil {
			return err
		}

		conf.Subsystems = append(conf.Subsystems, subsystem)
	}

	for _, name := range conf.CatalogsNames {

		value := Catalog{}
		err := Unpack(cfg.WithName(name, "Catalog"), &value)
		if err != nil {
			return err
		}

		conf.Catalogs = append(conf.Catalogs, value)
	}

	return nil
}
