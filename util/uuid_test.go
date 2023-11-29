package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShortIDString(t *testing.T) {
	expId := NewShortIDString("exp")
	assert.NotEqual(t, expId, "")
}
