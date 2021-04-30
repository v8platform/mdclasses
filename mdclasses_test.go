package mdclasses

import (
	"github.com/v8platform/mdclasses/encoding/xml"
	"os"
	"reflect"
	"testing"
)

func TestUnpackConfiguration(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		dir     string
		want    Configuration
		wantErr bool
	}{
		{
			"simple",
			"./tests/metadata/edt/src",
			Configuration{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackConfiguration(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackConfiguration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name    string
		dir     string
		want    Configuration
		wantErr bool
	}{
		{
			"simple",
			"./tests/metadata/edt/src",
			Configuration{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackConfiguration(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			xmlFile, err := os.Create("./tests/metadata/edt/src/Configuration/Configuration.mdo")
			if err != nil {
				t.Fatalf("Ошибка тестирования %v", got)
				return
			}
			xmlFile.WriteString(xml.Header)
			encoder := xml.NewEncoder(xmlFile)
			encoder.Indent("", "  ")
			err = encoder.Encode(&got)
			if err != nil {
				t.Fatalf("Ошибка тестирования %v", got)
				return
			}
		})
	}
}
