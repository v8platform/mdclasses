package mdclasses

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewMDOTypeRefFromString(t *testing.T) {

	tests := []struct {
		name string
		raw  string
		want MDOTypeRef
	}{
		{
			"simple",
			"AccountingRegister.РегистрНакопления1",
			MDOTypeRef{
				ACCOUNTING_REGISTER,
				"РегистрНакопления1",
				MDOTypeRef{},
				"AccountingRegister.РегистрНакопления1",
			},
		},
		{
			"ref",
			"9cd510cd-abfc-11d4-9434-004095e12fc7",
			MDOTypeRef{
				REF,
				"9cd510cd-abfc-11d4-9434-004095e12fc7",
				MDOTypeRef{},
				"9cd510cd-abfc-11d4-9434-004095e12fc7",
			},
		},
		{
			"with parent",
			"Catalog.Справочник1.Form.ФормаЭлемента",
			MDOTypeRef{
				FORM,
				"ФормаЭлемента",
				MDOTypeRef{
					mdoType: CATALOG,
					ref:     "Справочник1",
				},

				"Catalog.Справочник1.Form.ФормаЭлемента",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMDOTypeRefFromString(tt.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMDOTypeRefFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMDOTypeRefExist(t *testing.T) {
	got, err := UnpackConfiguration("tests/metadata/edt/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}
	subsystem := MDOTypeRef{
		mdoType: "Subsystem",
		ref:     "ПерваяПодсистема",
		parent:  nil,
		raw:     "Subsystem.ПерваяПодсистема",
	}
	require.True(t, got.ConfigurationChildObjects.Subsystems.Exist(subsystem))
}

func TestMDOTypeRefGetIndex(t *testing.T) {
	got, err := UnpackConfiguration("tests/metadata/edt/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}
	mdo := MDOTypeRef{
		mdoType: "Subsystem",
		ref:     "ВтораяПодсистема",
		parent:  nil,
		raw:     "",
	}
	index := got.ConfigurationChildObjects.Subsystems.GetIndex(mdo)
	if index < 0 {
		t.Errorf("Ошибка поиска дочернего объекта = %v", err)
	}
	require.True(t, index == 1)
}

func TestMDOTypeRefDelete(t *testing.T) {
	got, err := UnpackConfiguration("tests/metadata/edt/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}
	ref := got.ConfigurationChildObjects.Subsystems
	newChild, err := ref.Delete(1)
	if err != nil {
		t.Errorf("Ошибка поиска дочернего объекта = %v", err)
	}
	got.ConfigurationChildObjects.Subsystems = newChild
	require.True(t, len(got.ConfigurationChildObjects.Subsystems) == 1)
}

func TestMDOTypeRefCreate(t *testing.T) {
	got, err := UnpackConfiguration("tests/metadata/edt/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}

	MDOTypeRef := NewMDOTypeRefFromString("Subsystem.NewSubsystem")
	got.ConfigurationChildObjects.Subsystems = append(got.ConfigurationChildObjects.Subsystems, MDOTypeRef)
	require.True(t, true)
}
