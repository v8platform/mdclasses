package mdclasses

import (
	"github.com/khorevaa/logos"
	"github.com/v8platform/mdclasses/encoding/xml"
)

var log = logos.New("github.com/v8platform/mdclasses").Sugar()

type MDOBaseType struct {
	Name    string             `xml:"name"`
	Synonym ObjectKeyValueType `xml:"synonym"`
	Uuid    string             `xml:"uuid,attr"`
	Mdclass xml.Attr           `xml:"mdclass,attr"`
	Xsi     xml.Attr           `xml:"xsi,attr"`
	Core    xml.Attr           `xml:"core,attr"`
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
	ConfigurationProperties
	ConfigurationChildObjects

	XMLName          xml.Name          `xml:"mdclass Configuration"`
	ContainedObjects []ContainedObject `xml:"containedObjects,allowempty"`

	Languages  []*Language  `xml:"languages"`
	Subsystems []*Subsystem `xml:"-"`

	// StyleItems                  []MDOTypeRef `xml:"styleItems"`
	// CommonPictures              []MDOTypeRef `xml:"commonPictures"`
	// SessionParameters           []MDOTypeRef `xml:"sessionParameters"`
	// Roles                       []MDOTypeRef `xml:"roles"`
	CommonTemplates []*CommonTemplates
	// FilterCriteria              []MDOTypeRef `xml:"filterCriteria"`
	CommonModules []*CommonModule
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
	Catalogs  []*Catalog
	Documents []*Document
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

	idxMDOType idxMDOTypeRef
}

type idxMDOTypeRef map[MDOTypeRef]interface{}

type MobileApplicationFunctionality struct {
	Functionality string `xml:"functionality,omitempty"`
	Use           string `xml:"use"`
}

type MobileApplicationFunctionalities struct {
	Functionality []MobileApplicationFunctionality `xml:"functionality,omitempty"`
}

type ConfigurationProperties struct {
	ConfigurationExtensionCompatibilityMode string                           `xml:"configurationExtensionCompatibilityMode"`
	DefaultRunMode                          string                           `xml:"defaultRunMode"`
	UsePurposes                             string                           `xml:"usePurposes"`
	ScriptVariant                           string                           `xml:"scriptVariant"`
	DefaultRoles                            MDOTypeRefList                   `xml:"defaultRoles"`
	Vendor                                  string                           `xml:"vendor"`
	Version                                 string                           `xml:"version"`
	UpdateCatalogAddress                    string                           `xml:"updateCatalogAddress,omitempty"`
	UseManagedFormInOrdinaryApplication     string                           `xml:"useManagedFormInOrdinaryApplication"`
	UseOrdinaryFormInManagedApplication     string                           `xml:"useOrdinaryFormInManagedApplication"`
	ReportsVariantsStorage                  string                           `xml:"reportsVariantsStorage"`
	DefaultReportForm                       MDOTypeRef                       `xml:"defaultReportForm"`
	DefaultReportVariantForm                MDOTypeRef                       `xml:"defaultReportVariantForm"`
	DefaultReportSettingsForm               MDOTypeRef                       `xml:"defaultReportSettingsForm"`
	DefaultSearchForm                       MDOTypeRef                       `xml:"defaultSearchForm"`
	UsedMobileApplicationFunctionalities    MobileApplicationFunctionalities `xml:"usedMobileApplicationFunctionalities,omitempty"`
	RequiredMobileApplicationPermissions    []string                         `xml:"requiredMobileApplicationPermissions"`
	MainSectionPicture                      string                           `xml:"mainSectionPicture,allowempty"`
	DefaultLanguage                         MDOTypeRef                       `xml:"defaultLanguage"`
	BriefInformation                        ObjectKeyValueType               `xml:"briefInformation"`
	DetailedInformation                     ObjectKeyValueType               `xml:"detailedInformation"`
	Splash                                  string                           `xml:"splash,omitempty"`
	Copyright                               ObjectKeyValueType               `xml:"copyright"`
	VendorInformationAddress                ObjectKeyValueType               `xml:"vendorInformationAddress"`
	ConfigurationInformationAddress         ObjectKeyValueType               `xml:"configurationInformationAddress,omitempty"`
	DataLockControlMode                     string                           `xml:"dataLockControlMode"`
	ObjectAutonumerationMode                string                           `xml:"objectAutonumerationMode"`
	ModalityUseMode                         string                           `xml:"modalityUseMode"`
	InterfaceCompatibilityMode              string                           `xml:"interfaceCompatibilityMode"`
	CompatibilityMode                       string                           `xml:"compatibilityMode"`
}

type ConfigurationChildObjects struct {
	Subsystems                  MDOTypeRefList `xml:"subsystems"`
	StyleItems                  MDOTypeRefList `xml:"styleItems"`
	CommonPictures              MDOTypeRefList `xml:"commonPictures"`
	SessionParameters           MDOTypeRefList `xml:"sessionParameters"`
	Roles                       MDOTypeRefList `xml:"roles"`
	CommonTemplates             MDOTypeRefList `xml:"commonTemplates"`
	FilterCriteria              MDOTypeRefList `xml:"filterCriteria"`
	CommonModules               MDOTypeRefList `xml:"commonModules"`
	CommonAttributes            MDOTypeRefList `xml:"commonAttributes"`
	ExchangePlans               MDOTypeRefList `xml:"exchangePlans"`
	XDTOPackages                MDOTypeRefList `xml:"xDTOPackages"`
	WebServices                 MDOTypeRefList `xml:"webServices"`
	HttpServices                MDOTypeRefList `xml:"httpServices"`
	WsReferences                MDOTypeRefList `xml:"wsReferences"`
	EventSubscriptions          MDOTypeRefList `xml:"eventSubscriptions"`
	ScheduledJobs               MDOTypeRefList `xml:"scheduledJobs"`
	SettingsStorages            MDOTypeRefList `xml:"settingsStorages"`
	FunctionalOptions           MDOTypeRefList `xml:"functionalOptions"`
	FunctionalOptionsParameters MDOTypeRefList `xml:"functionalOptionsParameters"`
	DefinedTypes                MDOTypeRefList `xml:"definedTypes"`
	CommonCommands              MDOTypeRefList `xml:"commonCommands"`
	CommandGroups               MDOTypeRefList `xml:"commandGroups"`
	Constants                   MDOTypeRefList `xml:"constants"`
	CommonForms                 MDOTypeRefList `xml:"commonForms"`
	Catalogs                    MDOTypeRefList `xml:"catalogs"`
	Documents                   MDOTypeRefList `xml:"documents"`
	DocumentNumerators          MDOTypeRefList `xml:"documentNumerators"`
	DocumentJournals            MDOTypeRefList `xml:"documentJournals"`
	Enums                       MDOTypeRefList `xml:"enums"`
	Reports                     MDOTypeRefList `xml:"reports"`
	DataProcessors              MDOTypeRefList `xml:"dataProcessors"`
	InformationRegisters        MDOTypeRefList `xml:"informationRegisters"`
	AccumulationRegisters       MDOTypeRefList `xml:"accumulationRegisters"`
	ChartsOfCharacteristicTypes MDOTypeRefList `xml:"chartsOfCharacteristicTypes"`
	BusinessProcesses           MDOTypeRefList `xml:"businessProcesses"`
	Tasks                       MDOTypeRefList `xml:"tasks"`
	WebService                  MDOTypeRefList `xml:"serviceService"`
	WSReference                 MDOTypeRefList `xml:"wsReferenceReference"`
	XDTOPackage                 MDOTypeRefList `xml:"xdtoPackage"`
}

const ConfigurationFile = "Configuration.mdo"

func (conf *Configuration) Unpack(cfg UnpackConfig) error {

	conf.idxMDOType = make(idxMDOTypeRef)
	cfg.IdxObjects = conf.idxMDOType

	err := conf.ConfigurationChildObjects.Subsystems.Unpack(cfg, &conf.Subsystems)
	if err != nil {
		return err
	}

	err = conf.ConfigurationChildObjects.Catalogs.Unpack(cfg, &conf.Catalogs)
	if err != nil {
		return err
	}

	err = conf.ConfigurationChildObjects.Documents.Unpack(cfg, &conf.Documents)
	if err != nil {
		return err
	}

	err = conf.ConfigurationChildObjects.CommonTemplates.Unpack(cfg, &conf.CommonTemplates)
	if err != nil {
		return err
	}

	err = conf.ConfigurationChildObjects.CommonModules.Unpack(cfg, &conf.CommonModules)
	if err != nil {
		return err
	}
	// for _, mdoTypeRef := range conf.ConfigurationChildObjects.Subsystems {
	//
	// 	subsystem := Subsystem{}
	// 	err := mdoTypeRef.Unpack(cfg, &subsystem)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	conf.Subsystems = append(conf.Subsystems, subsystem)
	//
	// 	conf.idxMDOType[mdoTypeRef] = subsystem
	//
	// }

	// for _, name := range conf.ConfigurationChildObjects.Catalogs {
	//
	// 	value := Catalog{}
	// 	err := Unpack(cfg.WithName(name.String(), "Catalog"), &value)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	conf.Catalogs = append(conf.Catalogs, value)
	// 	conf.idxMDOType[name] = value
	//
	// }

	return nil
}
