package mdclasses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetChildByTypeName(t *testing.T) {
	got := unpackTestConf(t)
	children := got.ConfigurationChildObjects
	exchangePlan := children.GetChildByType(EVENT_SUBSCRIPTION)
	require.True(t, exchangePlan != nil)
}
