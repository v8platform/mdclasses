package mdclasses

func UnpackConfiguration(dir string) (Configuration, error) {

	cfg := NewUnpackConfig("Configuration", dir)
	conf := Configuration{}

	err := Unpack(cfg, &conf)
	if err != nil {
		return Configuration{}, err
	}

	return conf, nil

}
