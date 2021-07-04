package mdclasses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetContents(t *testing.T) {
	got, err := UnpackConfiguration("tests/metadata/ssl/src")
	if err != nil {
		t.Errorf("UnpackConfiguration() error = %v", err)
		return
	}
	var list MDOTypeRefList
	for _, subsystem := range got.Subsystems {
		if subsystem.MDOBaseType.Name == "СтандартныеПодсистемы" {
			list = subsystem.GetContents()
		}
	}
	log.Infof("Количество элементов: %d", len(list))
	require.True(t, len(list) == 2054)
}
