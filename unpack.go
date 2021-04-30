package mdclasses

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/v8platform/mdclasses/encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Unpacker interface {
	Unpack(cfg UnpackConfig) error
}

type UnpackConfig struct {
	Base       string
	Path       string
	IdxObjects idxMDOTypeRef
}

func NewUnpackConfig(base string) UnpackConfig {
	return UnpackConfig{
		base,
		base,
		make(idxMDOTypeRef),
	}
}

const ExtMdo = ".mdo"

var dirMap = map[string]string{
	"Configuration": "Configuration", // Для конифгурации
	"Subsystem":     "Subsystems",
	"Catalog":       "Catalogs",
	"Document":      "Documents",
}

func (cfg UnpackConfig) WithPath(path string) UnpackConfig {

	return UnpackConfig{
		cfg.Base,
		path,
		cfg.IdxObjects,
	}
}

func (cfg UnpackConfig) HasUnpacked(ref MDOTypeRef) (interface{}, bool) {
	val, ok := cfg.IdxObjects[ref]
	return val, ok

}

func (cfg UnpackConfig) StoreUnpacked(ref MDOTypeRef, value interface{}) {

	// TODO Надо добавить тег для в имя ключа, возможно пересечение имен
	cfg.IdxObjects[ref] = value

}

var NoNameFile = errors.New("no name unpack config")

func (cfg UnpackConfig) getFilename(name string) (string, error) {

	var filename string

	if strings.HasPrefix(name, ".") {
		filename = filepath.Join(cfg.Path, name)
	}

	filename = filepath.Join(cfg.Base, name)

	_, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	return filename, nil

}

func Unpack(cfg UnpackConfig, filename string, value interface{}) error {
	// TODO Сделать проверку что value это поинтер

	filename, err := cfg.getFilename(filename)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Ошибка чтения файла %s", err.Error())
	}

	err = xml.Unmarshal(data, value)

	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	if un, ok := value.(Unpacker); ok {
		err := un.Unpack(cfg.WithPath(filepath.Dir(filename)))
		if err != nil {
			return err
		}
	}

	return nil
}
