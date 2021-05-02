package mdclasses

import (
	"github.com/khorevaa/logos"
	"github.com/v8platform/mdclasses/encoding/xml"
)

var log = logos.New("github.com/v8platform/mdclasses").Sugar()

type MDOBaseType struct {
	Name    string   `xml:"name"`
	Mdclass xml.Attr `xml:"mdclass,attr"`
	Xsi     xml.Attr `xml:"xsi,attr"`
	Core    xml.Attr `xml:"core,attr"`
	Uuid    string   `xml:"uuid,attr"`
}
type ContainedObject struct {
	ClassId  string `xml:"classId,attr,allowempty"`
	ObjectId string `xml:"objectId,attr,allowempty"`
}

type Language struct {
	Uuid         string               `xml:"uuid,attr"`
	Name         string               `xml:"name"`
	Synonym      []ObjectKeyValueType `xml:"synonym"`
	LanguageCode string               `xml:"languageCode"`
}

type Configuration struct {
	MDOBaseType

	XMLName          xml.Name          `xml:"mdclass Configuration"`
	ContainedObjects []ContainedObject `xml:"containedObjects,allowempty"`
	ConfigurationProperties

	Languages  []*Language  `xml:"languages"`
	Subsystems []*Subsystem `xml:"-"`

	// StyleItems                  []MDOTypeRef `xml:"styleItems"`
	// CommonPictures              []MDOTypeRef `xml:"commonPictures"`
	// SessionParameters           []MDOTypeRef `xml:"sessionParameters"`
	// Roles                       []MDOTypeRef `xml:"roles"`
	CommonTemplates []*CommonTemplates `xml:"-"`
	// FilterCriteria              []MDOTypeRef `xml:"filterCriteria"`
	CommonModules []*CommonModule `xml:"-"`
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
	Catalogs  []*Catalog  `xml:"-"`
	Documents []*Document `xml:"-"`
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

	idxMDOType idxMDOTypeRef `xml:"-"`
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
	ConfigurationExtensionCompatibilityMode         string                            `xml:"configurationExtensionCompatibilityMode,omitempty"`
	DefaultRunMode                                  string                            `xml:"defaultRunMode,omitempty"`
	UsePurposes                                     []string                          `xml:"usePurposes,omitempty"`
	ScriptVariant                                   string                            `xml:"scriptVariant,omitempty"`
	Vendor                                          string                            `xml:"vendor,omitempty"`
	Version                                         string                            `xml:"version,omitempty"`
	UpdateCatalogAddress                            string                            `xml:"updateCatalogAddress,omitempty"`
	UseManagedFormInOrdinaryApplication             string                            `xml:"useManagedFormInOrdinaryApplication,omitempty"`
	UseOrdinaryFormInManagedApplication             string                            `xml:"useOrdinaryFormInManagedApplication,omitempty"`
	DefaultRoles                                    MDOTypeRefList                    `xml:"defaultRoles,omitempty"`
	ReportsVariantsStorage                          string                            `xml:"reportsVariantsStorage,omitempty"`
	DefaultReportForm                               MDOTypeRef                        `xml:"defaultReportForm,omitempty"`
	DefaultReportVariantForm                        MDOTypeRef                        `xml:"defaultReportVariantForm,omitempty"`
	DefaultReportSettingsForm                       MDOTypeRef                        `xml:"defaultReportSettingsForm,omitempty"`
	DefaultSearchForm                               MDOTypeRef                        `xml:"defaultSearchForm,omitempty"`
	UsedMobileApplicationFunctionalities            *MobileApplicationFunctionalities `xml:"usedMobileApplicationFunctionalities,omitempty"`
	RequiredMobileApplicationPermissions            []string                          `xml:"requiredMobileApplicationPermissions,omitempty"`
	MainSectionPicture                              string                            `xml:"mainSectionPicture,omitempty"`
	DefaultLanguage                                 MDOTypeRef                        `xml:"defaultLanguage,omitempty"`
	BriefInformation                                ObjectKeyValueType                `xml:"briefInformation,omitempty"`
	DetailedInformation                             ObjectKeyValueType                `xml:"detailedInformation,omitempty"`
	Splash                                          string                            `xml:"splash,omitempty"`
	Copyright                                       ObjectKeyValueType                `xml:"copyright,omitempty"`
	VendorInformationAddress                        ObjectKeyValueType                `xml:"vendorInformationAddress,omitempty"`
	ConfigurationInformationAddress                 ObjectKeyValueType                `xml:"configurationInformationAddress,omitempty"`
	DataLockControlMode                             string                            `xml:"dataLockControlMode,omitempty"`
	ObjectAutonumerationMode                        string                            `xml:"objectAutonumerationMode,omitempty"`
	SynchronousExtensionAndAddInCallUseMode         string                            `xml:"synchronousExtensionAndAddInCallUseMode,omitempty"`
	SynchronousPlatformExtensionAndAddInCallUseMode string                            `xml:"synchronousPlatformExtensionAndAddInCallUseMode,omitempty"`
	ModalityUseMode                                 string                            `xml:"modalityUseMode,omitempty"`
	InterfaceCompatibilityMode                      string                            `xml:"interfaceCompatibilityMode,omitempty"`
	CompatibilityMode                               string                            `xml:"compatibilityMode,omitempty"`
}

type ConfigurationChildObjects struct {
	Subsystems                  MDOTypeRefList `xml:"subsystems"`
	StyleItems                  MDOTypeRefList `xml:"styleItems"`
	Styles                      MDOTypeRefList `xml:"styles"`
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
	Sequences                   MDOTypeRefList `xml:"sequences"`
	DocumentJournals            MDOTypeRefList `xml:"documentJournals"`
	Enums                       MDOTypeRefList `xml:"enums"`
	Reports                     MDOTypeRefList `xml:"reports"`
	DataProcessors              MDOTypeRefList `xml:"dataProcessors"`
	InformationRegisters        MDOTypeRefList `xml:"informationRegisters"`
	AccumulationRegisters       MDOTypeRefList `xml:"accumulationRegisters"`
	ChartsOfCharacteristicTypes MDOTypeRefList `xml:"chartsOfCharacteristicTypes"`
	ChartsOfAccounts            MDOTypeRefList `xml:"chartsOfAccounts"`
	AccountingRegisters         MDOTypeRefList `xml:"accountingRegisters"`
	ChartsOfCalculationTypes    MDOTypeRefList `xml:"chartsOfCalculationTypes"`
	CalculationRegisters        MDOTypeRefList `xml:"calculationRegisters"`
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
