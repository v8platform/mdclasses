package mdclasses

import (
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
