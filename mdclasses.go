package mdclasses

import (
	"fmt"
	"github.com/v8platform/mdclasses/encoding/xml"
	"path/filepath"
)

func UnpackConfiguration(dir string) (Configuration, error) {

	cfg := NewUnpackConfig(dir)
	conf := Configuration{}

	err := Unpack(cfg, filepath.Join("Configuration", "Configuration.mdo"), &conf)
	if err != nil {
		return Configuration{}, err
	}

	bytes, err := xml.Marshal(conf)
	if err != nil {
		return Configuration{}, err
	}
	str := string(bytes)
	fmt.Println(str)

	return conf, nil

}
