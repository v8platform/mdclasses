package mdclasses

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Unpacker interface {
	Unpack(cfg UnpackConfig) error
}

type UnpackConfig struct {
	Tag        string
	Name       string // Subsystems
	Base       string
	Path       string
	IdxObjects *sync.Map
}

func NewUnpackConfig(name string, base string) UnpackConfig {
	return UnpackConfig{
		"",
		name,
		base,
		base,
		&sync.Map{},
	}
}

const ExtMdo = ".mdo"

var dirMap = map[string]string{
	"Configuration": "Configuration", // Для конифгурации
	"Subsystem":     "Subsystems",
	"Catalog":       "Catalogs",
	"Document":      "Documents",
}

func (cfg UnpackConfig) WithName(name string, tag string) UnpackConfig {

	return UnpackConfig{
		tag,
		name,
		cfg.Base,
		cfg.Path,
		cfg.IdxObjects,
	}
}

func (cfg UnpackConfig) WithPath(path string) UnpackConfig {

	return UnpackConfig{
		cfg.Tag,
		cfg.Name,
		cfg.Base,
		path,
		cfg.IdxObjects,
	}
}

func (cfg UnpackConfig) HasUnpacked() (interface{}, bool) {

	return cfg.IdxObjects.Load(cfg.Name)

}

func (cfg UnpackConfig) StoreUnpacked(value interface{}) {

	// TODO Надо добавить тег для в имя ключа, возможно пересечение имен
	cfg.IdxObjects.Store(cfg.Name, value)

}

var NoNameFile = errors.New("no name unpack config")

func getFilename(cfg UnpackConfig) (string, error) {

	names := strings.Split(cfg.Name, ".")

	switch len(names) {
	case 0:
		return "", NoNameFile
	case 1:
		name := names[0]
		dir := dirMap[cfg.Tag]

		filename := filepath.Join(cfg.Path, dir, name, name+ExtMdo)
		_, err := os.Stat(filename)
		if err != nil {
			return "", err
		}

		return filename, nil
	case 2:

		dir := dirMap[names[0]]
		name := names[1]

		filename := filepath.Join(cfg.Base, dir, name, name+ExtMdo)
		_, err := os.Stat(filename)
		if err != nil {
			return "", err
		}
		return filename, nil
	}

	return "", NoNameFile

}

func Unpack(cfg UnpackConfig, value interface{}) error {
	// TODO Сделать проверку что value это поинтер

	if val, ok := cfg.HasUnpacked(); ok {
		value = val
		return nil
	}

	filename, err := getFilename(cfg)
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
