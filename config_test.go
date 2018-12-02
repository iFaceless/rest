package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOptions(t *testing.T) {
	SetOptions(DefaultLimit(123),
		DefaultOffset(456))
	assert.Equal(t, 123, defaultLimit)
	assert.Equal(t, 456, defaultOffset)
}
