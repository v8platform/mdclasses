package mdclasses

import "testing"

func unpackTestConf(t *testing.T) Configuration {
	got, err := UnpackConfiguration("tests/metadata/ssl/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return Configuration{}
	}
	return got
}
