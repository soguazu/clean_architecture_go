package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	carDetailsService = NewCarDetailsService()
)

func TestGetData(t *testing.T) {
	carDetails := carDetailsService.GetDetails()

	assert.NotNil(t, carDetails)
}
