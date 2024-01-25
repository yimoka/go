package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpType(t *testing.T) {
	assert.Equal(t, OpType("").String(), "")
	assert.Equal(t, OpType("").Label(), "")
	assert.Equal(t, OpType("add").Label(), "添加")
	assert.Equal(t, OpAdd.Label(), "添加")
}
