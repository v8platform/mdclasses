package mdclasses

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/v8platform/mdclasses/encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
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

func TestCompareFile(t *testing.T) {
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

			TempDir, err := ioutil.TempDir(os.TempDir(), "prefix")
			if err != nil {
				t.Fatalf("Error create temp dir %s", TempDir)
			}
			os.MkdirAll(filepath.Join(TempDir, "Configuration"), os.ModePerm)
			xmlFile, err := os.Create(filepath.Join(TempDir, "Configuration", "Configuration.mdo"))
			if err != nil {
				t.Fatalf("Ошибка тестирования %v", got)
				return
			}
			xmlFile.WriteString(xml.Header)
			encoder := xml.NewEncoder(xmlFile)
			encoder.Indent("", "  ")
			err = encoder.Encode(&got)
			xmlFile.WriteString("\n")

			if err != nil {
				t.Fatalf("Ошибка тестирования %v", got)
				return
			}
			require.True(t, fileCompare(t, filepath.Join(tt.dir, "Configuration", "Configuration.mdo"), xmlFile.Name()))
		})

	}
}

func fileCompare(t *testing.T, file1, file2 string) bool {
	// per comment, better to not read an entire file into memory
	// this is simply a trivial example.
	f1, err1 := ioutil.ReadFile(file1)

	if err1 != nil {
		t.Fatalf("Ошибка тестирования %v", err1)
	}

	f2, err2 := ioutil.ReadFile(file2)

	if err2 != nil {
		t.Fatalf("Ошибка тестирования %v", err1)
	}

	return bytes.Equal(f1, f2) // Per comment, this is significantly more performant.
}
