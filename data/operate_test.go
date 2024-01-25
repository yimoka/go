package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpType(t *testing.T) {
	assert.Equal(t, OpType("").String(), "")
}
