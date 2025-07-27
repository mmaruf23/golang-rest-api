package test

import (
	"testing"

	"github.com/mmaruf23/golang-rest-api/internal/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InisializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InisializeService(true)

	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}
