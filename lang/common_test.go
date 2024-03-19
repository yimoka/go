package lang

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/yimoka/go/config"
)

func TestGetCommonBundle(t *testing.T) {
	bundle := GetCommonBundle(nil)
	// Add your assertions here
	assert.NotNil(t, bundle)
	// Add more assertions here if needed

	l := NewCommonLang(map[string]*config.Lang{
		"zh": {Messages: []*config.LangMessage{
			// 新增的
			{Id: "add", Other: "新增的"},
			// 替换原来的
			{Id: dataNotFoundKey, Other: "新的 找不到数据"},
		}},
	})
	localizer := i18n.NewLocalizer(l.Bundle, "zh")
	value, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "add"})
	assert.Equal(t, "新增的", value)
	lang := "zh"
	expected := "新的 找不到数据"
	actual := l.GetNotDataFoundMsg(lang)
	assert.Equal(t, expected, actual)
	actual = l.GetNotDataDuplicateMsg("zh")
	assert.Equal(t, "该数据已存在,请勿重复添加", actual)
}

func TestGetNotDataFoundMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data not found"
	actual := l.GetNotDataFoundMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "找不到数据"
	actual = l.GetNotDataFoundMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataDuplicateMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data duplicate"
	actual := l.GetNotDataDuplicateMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "该数据已存在,请勿重复添加"
	actual = l.GetNotDataDuplicateMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataConstraintMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data constraint check failed"
	actual := l.GetNotDataConstraintMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据约束检查失败，请检查您的参数"
	actual = l.GetNotDataConstraintMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataNotLoadedMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Database not loaded"
	actual := l.GetNotDataNotLoadedMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据库未加载，请联系管理员"
	actual = l.GetNotDataNotLoadedMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataNotSingularMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data error Not Singular"
	actual := l.GetNotDataNotSingularMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据出错了 Not Singular,请联系管理员"
	actual = l.GetNotDataNotSingularMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataValidationErrorMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data validation failed"
	actual := l.GetNotDataValidationErrorMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据校验失败，请检查您的参数"
	actual = l.GetNotDataValidationErrorMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetNotDataErrorMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data layer error"
	actual := l.GetNotDataErrorMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据层出错了,请联系管理员"
	actual = l.GetNotDataErrorMsg(lang)
	assert.Equal(t, expected, actual)
}
