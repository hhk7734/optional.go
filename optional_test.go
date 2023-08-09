package optional_test

import (
	"testing"

	"github.com/hhk7734/optional.go"
	"github.com/stretchr/testify/assert"
)

func stringPtr(v string) *string {
	return &v
}

func TestNil(t *testing.T) {
	var sPtr optional.Optional[*string]
	assert.Nil(t, sPtr)
	sPtr = optional.New[*string](nil)
	assert.NotNil(t, sPtr)
}
