package mdclasses

import (
	"path/filepath"
)

func UnpackConfiguration(dir string) (Configuration, error) {

	cfg := NewUnpackConfig(dir)
	conf := Configuration{}

	err := Unpack(cfg, filepath.Join("Configuration", "Configuration.mdo"), &conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil

}
