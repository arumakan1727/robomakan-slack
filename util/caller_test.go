package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCallerInfo(t *testing.T) {
	caller, err := GetCallerInfo()
	require.NoError(t, err)
	assert.Equal(t, "testing.tRunner", caller.FnName)
	assert.Equal(t, "testing.go", caller.File)
	assert.NotZero(t, caller.Line)
}

func TestGetCallerInfoStr(t *testing.T) {
	s := GetCallerInfoStr()
	assert.Regexp(t, `testing.go:\d+:testing.tRunner`, s)
}
