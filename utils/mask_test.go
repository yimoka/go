package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMaskPhone 测试手机号加星
func TestMaskPhone(t *testing.T) {
	assert.Equal(t, MaskPhone(""), "")
	assert.Equal(t, MaskPhone("1"), "*")
	assert.Equal(t, MaskPhone("12"), "**")
	assert.Equal(t, MaskPhone("123"), "***")
	assert.Equal(t, MaskPhone("1234"), "****")
	assert.Equal(t, MaskPhone("12345"), "****5")
	assert.Equal(t, MaskPhone("123456"), "1****6")
	assert.Equal(t, MaskPhone("1234567"), "1****67")
	assert.Equal(t, MaskPhone("12345678"), "12****78")
	assert.Equal(t, MaskPhone("123456789"), "12****789")
	assert.Equal(t, MaskPhone("1234567890"), "123****890")
	assert.Equal(t, MaskPhone("12345678901"), "123****8901")
	assert.Equal(t, MaskPhone("123456789012"), "123*****9012")
	assert.Equal(t, MaskPhone("1234567890123"), "123******0123")
}

func TestMaskMail(t *testing.T) {
	assert.Equal(t, MaskMail(""), "***@***")
	assert.Equal(t, MaskMail("i"), "i***@***")
	assert.Equal(t, MaskMail("ixxx@qq.com"), "i***@***.com")
	assert.Equal(t, MaskMail("ixxx@163.qq.com"), "i***@***.com")
	assert.Equal(t, MaskMail("ixxx@163163"), "i***@***")
}
