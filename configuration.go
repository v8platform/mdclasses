package mdclasses

import (
	"github.com/khorevaa/logos"
	"github.com/v8platform/mdclasses/encoding/xml"
)

var log = logos.New("github.com/v8platform/mdclasses").Sugar()

type MDOBaseType struct {
	Mdclass xml.Attr           `xml:"mdclass,attr"`
	Uuid    string             `xml:"uuid,attr"`
	Name    string             `xml:"name"`
	Synonym ObjectKeyValueType `xml:"synonym"`
}
type ContainedObject struct {
	ClassId  string `xml:"classId,attr,allowempty"`
	ObjectId string `xml:"objectId,attr,allowempty"`
}

type Language struct {
	Uuid         string             `xml:"uuid,attr"`
	Name         string             `xml:"name"`
	Synonym      ObjectKeyValueType `xml:"synonym"`
	LanguageCode string             `xml:"languageCode"`
}

type Configuration struct {
	MDOBaseType
	XMLName                                 xml.Name          `xml:"mdclass Configuration"`
	ContainedObjects                        []ContainedObject `xml:"containedObjects,allowempty"`
	ConfigurationExtensionCompatibilityMode string            `xml:"configurationExtensionCompatibilityMode"`
	DefaultRunMode                          string            `xml:"defaultRunMode"`
	UsePurposes                             string            `xml:"usePurposes"`
	ScriptVariant                           string            `xml:"scriptVariant"`
	DefaultRoles                            []MDOTypeRef      `xml:"defaultRoles"`
	Vendor                                  string            `xml:"vendor"`
	Version                                 string            `xml:"version"`
	UpdateCatalogAddress                    string            `xml:"updateCatalogAddress,omitempty"`
	UseManagedFormInOrdinaryApplication     string            `xml:"useManagedFormInOrdinaryApplication"`
	UseOrdinaryFormInManagedApplication     string            `xml:"useOrdinaryFormInManagedApplication"`
	ReportsVariantsStorage                  string            `xml:"reportsVariantsStorage"`
	DefaultReportForm                       MDOTypeRef        `xml:"defaultReportForm"`
	DefaultReportVariantForm                MDOTypeRef        `xml:"defaultReportVariantForm"`
	DefaultReportSettingsForm               MDOTypeRef        `xml:"defaultReportSettingsForm"`
	DefaultSearchForm                       MDOTypeRef        `xml:"defaultSearchForm"`
	UsedMobileApplicationFunctionalities    struct {
		Functionality []struct {
			Functionality string `xml:"functionality,omitempty"`
			Use           string `xml:"use"`
		} `xml:"functionality"`
	} `xml:"usedMobileApplicationFunctionalities"`
	MainSectionPicture              string             `xml:"mainSectionPicture,allowempty"`
	DefaultLanguage                 MDOTypeRef         `xml:"defaultLanguage"`
	BriefInformation                ObjectKeyValueType `xml:"briefInformation"`
	DetailedInformation             ObjectKeyValueType `xml:"detailedInformation"`
	Splash                          string             `xml:"splash,omitempty"`
	Copyright                       ObjectKeyValueType `xml:"copyright"`
	VendorInformationAddress        ObjectKeyValueType `xml:"vendorInformationAddress"`
	ConfigurationInformationAddress ObjectKeyValueType `xml:"configurationInformationAddress,omitempty"`
	DataLockControlMode             string             `xml:"dataLockControlMode"`
	ObjectAutonumerationMode        string             `xml:"objectAutonumerationMode"`
	ModalityUseMode                 string             `xml:"modalityUseMode"`
	InterfaceCompatibilityMode      string             `xml:"interfaceCompatibilityMode"`
	CompatibilityMode               string             `xml:"compatibilityMode"`
	Languages                       []Language         `xml:"languages"`

	Subsystems []Subsystem `xml:"-"`
	// SubsystemsNames []MDOTypeRef `xml:"subsystems"`

	// StyleItems                  []MDOTypeRef `xml:"styleItems"`
	// CommonPictures              []MDOTypeRef `xml:"commonPictures"`
	// SessionParameters           []MDOTypeRef `xml:"sessionParameters"`
	// Roles                       []MDOTypeRef `xml:"roles"`
	// CommonTemplates             []MDOTypeRef `xml:"commonTemplates"`
	// FilterCriteria              []MDOTypeRef `xml:"filterCriteria"`
	// CommonModules               []MDOTypeRef `xml:"commonModules"`
	// CommonAttributes            []MDOTypeRef `xml:"commonAttributes"`
	// ExchangePlans               []MDOTypeRef `xml:"exchangePlans"`
	// XDTOPackages                []MDOTypeRef `xml:"xDTOPackages"`
	// WebServices                 []MDOTypeRef `xml:"webServices"`
	// HttpServices                []MDOTypeRef `xml:"httpServices"`
	// WsReferences                []MDOTypeRef `xml:"wsReferences"`
	// EventSubscriptions          []MDOTypeRef `xml:"eventSubscriptions"`
	// ScheduledJobs               []MDOTypeRef `xml:"scheduledJobs"`
	// SettingsStorages            []MDOTypeRef `xml:"settingsStorages"`
	// FunctionalOptions           []MDOTypeRef `xml:"functionalOptions"`
	// FunctionalOptionsParameters []MDOTypeRef `xml:"functionalOptionsParameters"`
	// DefinedTypes                []MDOTypeRef `xml:"definedTypes"`
	// CommonCommands              []MDOTypeRef `xml:"commonCommands"`
	// CommandGroups               []MDOTypeRef `xml:"commandGroups"`
	// Constants                   []MDOTypeRef `xml:"constants"`
	// CommonForms                 []MDOTypeRef `xml:"commonForms"`
	// CatalogsNames               []MDOTypeRef `xml:"catalogs"`
	Catalogs []Catalog `xml:"-"`
	// Documents                   []MDOTypeRef `xml:"documents"`
	// DocumentNumerators          []MDOTypeRef `xml:"documentNumerators"`
	// DocumentJournals            []MDOTypeRef `xml:"documentJournals"`
	// Enums                       []MDOTypeRef `xml:"enums"`
	// Reports                     []MDOTypeRef `xml:"reports"`
	// DataProcessors              []MDOTypeRef `xml:"dataProcessors"`
	// InformationRegisters        []MDOTypeRef `xml:"informationRegisters"`
	// AccumulationRegisters       []MDOTypeRef `xml:"accumulationRegisters"`
	// ChartsOfCharacteristicTypes []MDOTypeRef `xml:"chartsOfCharacteristicTypes"`
	// BusinessProcesses           []MDOTypeRef `xml:"businessProcesses"`
	// Tasks                       []MDOTypeRef `xml:"tasks"`

	ConfigurationChildObjects
}

type ConfigurationChildObjects struct {
	Subsystems                  []MDOTypeRef `xml:"subsystems"`
	StyleItems                  []MDOTypeRef `xml:"styleItems"`
	CommonPictures              []MDOTypeRef `xml:"commonPictures"`
	SessionParameters           []MDOTypeRef `xml:"sessionParameters"`
	Roles                       []MDOTypeRef `xml:"roles"`
	CommonTemplates             []MDOTypeRef `xml:"commonTemplates"`
	FilterCriteria              []MDOTypeRef `xml:"filterCriteria"`
	CommonModules               []MDOTypeRef `xml:"commonModules"`
	CommonAttributes            []MDOTypeRef `xml:"commonAttributes"`
	ExchangePlans               []MDOTypeRef `xml:"exchangePlans"`
	XDTOPackages                []MDOTypeRef `xml:"xDTOPackages"`
	WebServices                 []MDOTypeRef `xml:"webServices"`
	HttpServices                []MDOTypeRef `xml:"httpServices"`
	WsReferences                []MDOTypeRef `xml:"wsReferences"`
	EventSubscriptions          []MDOTypeRef `xml:"eventSubscriptions"`
	ScheduledJobs               []MDOTypeRef `xml:"scheduledJobs"`
	SettingsStorages            []MDOTypeRef `xml:"settingsStorages"`
	FunctionalOptions           []MDOTypeRef `xml:"functionalOptions"`
	FunctionalOptionsParameters []MDOTypeRef `xml:"functionalOptionsParameters"`
	DefinedTypes                []MDOTypeRef `xml:"definedTypes"`
	CommonCommands              []MDOTypeRef `xml:"commonCommands"`
	CommandGroups               []MDOTypeRef `xml:"commandGroups"`
	Constants                   []MDOTypeRef `xml:"constants"`
	CommonForms                 []MDOTypeRef `xml:"commonForms"`
	Catalogs                    []MDOTypeRef `xml:"catalogs"`
	Documents                   []MDOTypeRef `xml:"documents"`
	DocumentNumerators          []MDOTypeRef `xml:"documentNumerators"`
	DocumentJournals            []MDOTypeRef `xml:"documentJournals"`
	Enums                       []MDOTypeRef `xml:"enums"`
	Reports                     []MDOTypeRef `xml:"reports"`
	DataProcessors              []MDOTypeRef `xml:"dataProcessors"`
	InformationRegisters        []MDOTypeRef `xml:"informationRegisters"`
	AccumulationRegisters       []MDOTypeRef `xml:"accumulationRegisters"`
	ChartsOfCharacteristicTypes []MDOTypeRef `xml:"chartsOfCharacteristicTypes"`
	BusinessProcesses           []MDOTypeRef `xml:"businessProcesses"`
	Tasks                       []MDOTypeRef `xml:"tasks"`
	WebService                  []MDOTypeRef `xml:"serviceService"`
	WSReference                 []MDOTypeRef `xml:"wsReferenceReference"`
	XDTOPackage                 []MDOTypeRef `xml:"xdtoPackage"`
}

const ConfigurationFile = "Configuration.mdo"

func (conf *Configuration) Unpack(cfg UnpackConfig) error {

	for _, name := range conf.ConfigurationChildObjects.Subsystems {

		subsystem := Subsystem{}
		err := Unpack(cfg.WithName(name.String(), "Subsystem"), &subsystem)
		if err != nil {
			return err
		}

		conf.Subsystems = append(conf.Subsystems, subsystem)
	}

	for _, name := range conf.ConfigurationChildObjects.Catalogs {

		value := Catalog{}
		err := Unpack(cfg.WithName(name.String(), "Catalog"), &value)
		if err != nil {
			return err
		}

		conf.Catalogs = append(conf.Catalogs, value)
	}

	return nil
}
