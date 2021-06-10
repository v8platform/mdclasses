package mdclasses

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/v8platform/mdclasses/encoding/xml"
	"io"
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

func TestMarshalConfigurationMDO(t *testing.T) {
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
			require.True(t, deepCompare(t, filepath.Join(tt.dir, "Configuration", "Configuration.mdo"), xmlFile.Name()))
		})

	}
}

func TestMarshalCatalog(t *testing.T){
	ws, _ := os.Getwd()
	srcDir := filepath.Join(ws, "tests", "metadata", "edt", "src")
	got, err := UnpackConfiguration(srcDir)
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}

	ObjectName := "ПрофилиГруппДоступа"

	TempDir, err := ioutil.TempDir(os.TempDir(), "prefix")
	if err != nil {
		t.Fatalf("Error create temp dir %s", TempDir)
	}

	os.MkdirAll(filepath.Join(TempDir, "Catalogs", ObjectName), os.ModePerm)
	srcTestFileName := filepath.Join(srcDir, "Catalogs", ObjectName, ObjectName +".mdo")
	DestTestFileName := filepath.Join(TempDir, "Catalogs", ObjectName, ObjectName +".mdo")

	xmlFile, err := os.Create(DestTestFileName)
	if err != nil {
		t.Fatalf("Ошибка тестирования %v", got)
		return
	}

	for _, catalog := range got.Catalogs {
		if catalog.MDOBaseType.Name == ObjectName {
			catalog.Pack(xmlFile.Name())
		}
	}
	require.True(t, deepCompare(t, srcTestFileName, xmlFile.Name()))
}

const chunkSize = 64000

func deepCompare(t *testing.T, file1, file2 string) bool {
	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				t.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func TestCheckPackCatalog(t *testing.T) {

}
