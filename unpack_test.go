package mdclasses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Тест, возвращающий путь к объекту метаданных относитьельно каталога исходников
func TestGetPathToObject(t *testing.T) {
	got := unpackTestConf(t)

	mdoType := got.ConfigurationChildObjects.Roles
	role, err := mdoType.GetByName("АдминистраторСистемы")
	if err != nil {

	}
	path := role.Filename()
	require.True(t, path == "Roles\\АдминистраторСистемы\\АдминистраторСистемы.mdo")
}
