package mdclasses

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"reflect"
	"strings"
)

type MDOType string

func (m MDOType) String() string {
	return string(m)
}

func (m MDOType) IsNull() bool {
	return len(m) == 0
}

func (m MDOType) MarshalText() ([]byte, error) {

	return []byte(m.String()), nil

}

func (m MDOType) isUnknown() bool {

	switch m {
	case ACCOUNTING_REGISTER, ACCUMULATION_REGISTER, BUSINESS_PROCESS, CALCULATION_REGISTER, CATALOG,
		CHART_OF_ACCOUNTS, CHART_OF_CALCULATION_TYPES, CHART_OF_CHARACTERISTIC_TYPES,
		COMMAND_GROUP, COMMON_ATTRIBUTE, COMMON_COMMAND, COMMON_FORM,
		COMMON_MODULE, COMMON_PICTURE, COMMON_TEMPLATE, CONFIGURATION, CONSTANT,
		DATA_PROCESSOR, DEFINED_TYPE, DOCUMENT_JOURNAL, DOCUMENT_NUMERATOR, DOCUMENT,
		ENUM, EVENT_SUBSCRIPTION, EXCHANGE_PLAN, FILTER_CRITERION,
		FUNCTIONAL_OPTION, FUNCTIONAL_OPTIONS_PARAMETER,
		HTTP_SERVICE, INFORMATION_REGISTER, INTERFACE, LANGUAGE,
		REPORT, ROLE, SCHEDULED_JOB, SEQUENCE, SESSION_PARAMETER,
		SETTINGS_STORAGE, STYLE_ITEM, STYLE, SUBSYSTEM, TASK,
		WEB_SERVICE, WS_REFERENCE, XDTO_PACKAGE,
		// MDO Для вложенных объектов
		FORM, COMMAND, TEMPLATE, ATTRIBUTE,
		RECALCULATION, WS_OPERATION, HTTP_SERVICE_URL_TEMPLATE, HTTP_SERVICE_METHOD,
		REF:

		return false

	default:
		return true
	}
}

func (m *MDOType) UnmarshalText(text []byte) error {

	value := MDOType(text)

	if value.isUnknown() {
		return errors.New("unknown <MDOType> " + string(text))
	}

	*m = value

	return nil
}

func (m MDOType) Group() string {

	switch m {
	case ACCOUNTING_REGISTER:
		return "AccountingRegisters"
	case CATALOG:
		return "Catalogs"
	case DOCUMENT:
		return "Documents"
	case SUBSYSTEM:
		return "Subsystems"
	case COMMON_TEMPLATE:
		return "CommonTemplates"
	case COMMON_MODULE:
		return "CommonModules"
	default:
		return ""
	}
}

type MDORef interface {
	Type() MDOType
	Ref() string
	Parent() MDORef
	String() string
	Dir() string
	Empty() bool
}

type MDOTypeRef struct {
	mdoType MDOType
	ref     string
	parent  MDORef
	raw     string
}

func (m MDOTypeRef) Type() MDOType {
	return m.mdoType
}

func (m MDOTypeRef) Ref() string {
	return m.ref
}

func (m MDOTypeRef) Parent() MDORef {
	return m.parent
}
func (m MDOTypeRef) Empty() bool {
	return m.mdoType.IsNull() && len(m.ref) == 0
}

func (m MDOTypeRef) String() string {

	if m.parent == nil || m.parent.Empty() {
		return fmt.Sprintf("%s.%s", m.mdoType, m.ref)
	}
	return fmt.Sprintf("%s.%s.%s", m.parent, m.mdoType, m.ref)

}

func (m MDOTypeRef) Unpack(cfg UnpackConfig, value interface{}) error {

	if val, ok := cfg.HasUnpacked(m); ok {
		value = val
		return nil
	}

	if m.mdoType == REF {
		return nil
	}

	filename := m.Filename()
	err := Unpack(cfg, filename, value)

	cfg.StoreUnpacked(m, value)

	return err
}

type EachMDOTypeRef []MDOTypeRef

func (e EachMDOTypeRef) Unpack(cfg UnpackConfig, value interface{}) error {

	var v reflect.Value

	v = reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("ErrInvalidPointer")
	}
	v = v.Elem()

	isSlice := v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8
	if !isSlice {
		return errors.New("NotSlice")
	}

	for _, item := range e {
		var ielem interface{}

		elem := reflectAlloc(v.Type().Elem())

		if v.Type().Elem().Kind() == reflect.Ptr {
			ielem = elem.Interface()
		} else {
			ielem = elem.Elem().Interface()
		}

		err := item.Unpack(cfg, ielem)
		if err != nil {
			return err
		}

		v.Set(reflect.Append(v, elem))
	}

	return nil
}

func reflectAlloc(typ reflect.Type) reflect.Value {
	if typ.Kind() == reflect.Ptr {
		return reflect.New(typ.Elem())
	}
	return reflect.New(typ).Elem()
}

func (m MDOTypeRef) MarshalText() ([]byte, error) {

	return []byte(m.String()), nil

}

func (m *MDOTypeRef) UnmarshalText(text []byte) error {

	*m = NewMDOTypeRefFromString(string(text))

	return nil

}

func newMDOTypeRef(modType MDOType, ref string, parent MDOTypeRef) MDOTypeRef {

	raw := fmt.Sprintf("%s.%s", modType, ref)
	if modType == REF {
		raw = ref
	}

	return MDOTypeRef{
		mdoType: modType,
		ref:     ref,
		parent:  parent,
		raw:     raw,
	}
}

func NewMDOTypeRefFromString(raw string) MDOTypeRef {

	values := strings.Split(raw, ".")
	mdoType := REF
	ref := raw

	parent := MDOTypeRef{}

	if len(values) > 2 {
		//add := len(values)%2
		parent = getParentMDOTypeRef(raw, parent)
	}

	if len(values)%2 == 0 {
		mdoType = MDOType(values[len(values)-2])
		ref = values[len(values)-1]
	}

	return newMDOTypeRef(mdoType, ref, parent)
}

func getParentMDOTypeRef(name string, parent MDOTypeRef) MDOTypeRef {

	val := parent

	n := 0

	for i, c := range name {

		if c == '.' {
			n++
		}

		if n == 2 {
			n = 0
			parentName := name[0:i]
			if parentName != "" {
				values := strings.Split(parentName, ".")
				val = newMDOTypeRef(MDOType(values[0]), values[1], val)
			}
		}

	}
	return val

}

func (m MDOTypeRef) IsNull() bool {
	return m.mdoType.isUnknown() && len(m.ref) == 0
}

func (m MDOTypeRef) Filename() string {

	return filepath.Join(m.Dir(), m.ref+ExtMdo)

}

func (m MDOTypeRef) Dir() string {

	if m.parent == nil || m.parent.Empty() {
		return filepath.Join(m.mdoType.Group(), m.ref)
	}

	return filepath.Join(m.parent.Dir(), m.mdoType.Group(), m.ref)

}

const (
	UNKNOWN                       MDOType = ""
	REF                           MDOType = "_RefType_"
	ACCOUNTING_REGISTER           MDOType = "AccountingRegister"
	ACCUMULATION_REGISTER         MDOType = "AccumulationRegister"       // "AccumulationRegisters", "РегистрНакопления", "РегистрыНакопления"),
	BUSINESS_PROCESS              MDOType = "BusinessProcess"            // "BusinessProcesses", "БизнесПроцесс", "БизнесПроцессы"),
	CALCULATION_REGISTER          MDOType = "CalculationRegister"        //, "CalculationRegisters", "РегистрРасчета", "РегистрыРасчета"),
	CATALOG                       MDOType = "Catalog"                    //, "Catalogs", "Справочник", "Справочники"),
	CHART_OF_ACCOUNTS             MDOType = "ChartOfAccounts"            //, "ChartsOfAccounts", "ПланСчетов", "ПланыСчетов"),
	CHART_OF_CALCULATION_TYPES    MDOType = "ChartOfCalculationTypes"    //, "ChartsOfCalculationTypes", "ПланВидовРасчета", "ПланыВидовРасчета"),
	CHART_OF_CHARACTERISTIC_TYPES MDOType = "ChartOfCharacteristicTypes" //, "ChartsOfCharacteristicTypes",  "ПланВидовХарактеристик", "ПланыВидовХарактеристик"),
	COMMAND_GROUP                 MDOType = "CommandGroup"               //, "CommandGroups", "ГруппаКоманд", "ГруппыКоманд"),
	COMMON_ATTRIBUTE              MDOType = "CommonAttribute"            //, "CommonAttributes", "ОбщийРеквизит", "ОбщиеРеквизиты"),
	COMMON_COMMAND                MDOType = "CommonCommand"              //, "CommonCommands", "ОбщаяКоманда", "ОбщиеКоманды"),
	COMMON_FORM                   MDOType = "CommonForm"                 //, "CommonForms", "ОбщаяФорма", "ОбщиеФормы"),
	COMMON_MODULE                 MDOType = "CommonModule"               //, "CommonModules", "ОбщийМодуль", "ОбщиеМодули"),
	COMMON_PICTURE                MDOType = "CommonPicture"              //, "CommonPictures", "ОбщаяКартинка", "ОбщиеКартинки"),
	COMMON_TEMPLATE               MDOType = "CommonTemplate"             //, "CommonTemplates", "ОбщийМакет", "ОбщиеМакеты"),
	CONFIGURATION                 MDOType = "Configuration"              //, "", "Конфигурация", ""),
	CONSTANT                      MDOType = "Constant"                   //, "Constants", "Константа", "Константы"),
	DATA_PROCESSOR                MDOType = "DataProcessor"              //, "DataProcessors", "Обработка", "Обработки"),
	DEFINED_TYPE                  MDOType = "DefinedType"                //, "DefinedTypes", "ОпределяемыйТип", "ОпределяемыеТипы"),
	DOCUMENT_JOURNAL              MDOType = "DocumentJournal"            //, "DocumentJournals", "ЖурналДокументов", "ЖурналыДокументов"),
	DOCUMENT_NUMERATOR            MDOType = "DocumentNumerator"          //, "DocumentNumerators", "Нумератор", "Нумераторы"),
	DOCUMENT                      MDOType = "Document"                   //, "Documents", "Документ", "Документы"),
	ENUM                          MDOType = "Enum"                       //, "Enums", "Перечисление", "Перечисления"),
	EVENT_SUBSCRIPTION            MDOType = "EventSubscription"          //, "EventSubscriptions", "ПодпискаНаСобытие", "ПодпискиНаСобытия"),
	EXCHANGE_PLAN                 MDOType = "ExchangePlan"               //, "ExchangePlans", "ПланОбмена", "ПланыОбмена"),
	FILTER_CRITERION              MDOType = "FilterCriterion"            //, "FilterCriteria", "КритерийОтбора", "КритерииОтбора"),
	FUNCTIONAL_OPTION             MDOType = "FunctionalOption"           //, "FunctionalOptions", "ФункциональнаяОпция", "ФункциональныеОпции"),
	FUNCTIONAL_OPTIONS_PARAMETER  MDOType = "FunctionalOptionsParameter" //, "FunctionalOptionsParameters", "ПараметрФункциональныхОпций", "ПараметрыФункциональныхОпций"),
	HTTP_SERVICE                  MDOType = "HTTPService"                //, "HTTPServices", "HTTPСервис", "HTTPСервисы"),
	INFORMATION_REGISTER          MDOType = "InformationRegister"        //, "InformationRegisters", "РегистрСведений", "РегистрыСведений"),
	INTERFACE                     MDOType = "Interface"                  //, "Interfaces", "Интерфейс", "Интерфейсы"),
	LANGUAGE                      MDOType = "Language"                   //, "Languages", "Язык", "Языки"),
	REPORT                        MDOType = "Report"                     //, "Reports", "Отчет", "Отчеты"),
	ROLE                          MDOType = "Role"                       //, "Roles", "Роль", "Роли"),
	SCHEDULED_JOB                 MDOType = "ScheduledJob"               //, "ScheduledJobs", "РегламентноеЗадание", "РегламентныеЗадания"),
	SEQUENCE                      MDOType = "Sequence"                   //, "Sequences", "Последовательность", "Последовательности"),
	SESSION_PARAMETER             MDOType = "SessionParameter"           //, "SessionParameters", "ПараметрСеанса", "ПараметрыСеанса"),
	SETTINGS_STORAGE              MDOType = "SettingsStorage"            //, "SettingsStorages", "ХранилищеНастроек", "ХранилищаНастроек"),
	STYLE_ITEM                    MDOType = "StyleItem"                  //, "StyleItems", "ЭлементСтиля", "ЭлементыСтиля"),
	STYLE                         MDOType = "Style"                      //, "Styles", "Стиль", "Стили"),
	SUBSYSTEM                     MDOType = "Subsystem"                  //, "Subsystems", "Подсистема", "Подсистемы"),
	TASK                          MDOType = "Task"                       //, "Tasks", "Задача", "Задачи"),
	WEB_SERVICE                   MDOType = "WebService"                 //, "WebServices", "WebСервис", "WebСервисы"),
	WS_REFERENCE                  MDOType = "WSReference"                //, "WSReferences", "WSСсылка", "WSСсылки"),
	XDTO_PACKAGE                  MDOType = "XDTOPackage"                //, "XDTOPackages", "ПакетXDTO", "ПакетыXDTO"),

	FORM                      MDOType = "Form"          //, "Forms", "Форма", "Формы"),
	COMMAND                   MDOType = "Command"       //, "Commands", "Команда", "Команды"),
	TEMPLATE                  MDOType = "Template"      //, "Templates", "Макет", "Макеты"),
	ATTRIBUTE                 MDOType = "Attribute"     //, "Attributes", "Реквизит", "Реквизиты"),
	RECALCULATION             MDOType = "Recalculation" //, "Recalculations", "Перерасчет", "Перерасчеты"),
	WS_OPERATION              MDOType = "Operation"     //, "Operations", "Операция", "Операции"),
	HTTP_SERVICE_URL_TEMPLATE MDOType = "URLTemplate"   //, "URLTemplates", "ШаблонURL", "ШаблоныURL"),
	HTTP_SERVICE_METHOD       MDOType = "Method"        //, "Methods", "Метод", "Методы"),

)
