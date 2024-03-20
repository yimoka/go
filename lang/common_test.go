package lang

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/yimoka/api/fault"
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
	actual := l.GetDataFoundMsg(lang)
	assert.Equal(t, expected, actual)
	actual = l.GetDataDuplicateMsg("zh")
	assert.Equal(t, "该数据已存在,请勿重复添加", actual)
}

func TestHandleMetadataError(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	err := fault.ErrorBadRequest("Get metadata failed")
	expected := "error: code = 400 reason = BAD_REQUEST message = Get metadata failed metadata = map[] cause = <nil>"
	actual := l.HandleMetadataError(err, lang)
	assert.Equal(t, expected, actual.Error())

	err.Metadata = map[string]string{"source": "string", "target": "int"}
	expected = "error: code = 400 reason = BAD_REQUEST message = metadata string conversion to int failed metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(err, lang)
	assert.Equal(t, expected, actual.Error())

	lang = "zh-CN"
	expected = "error: code = 400 reason = BAD_REQUEST message = 元数据 string 转换为 int 失败 metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(err, lang)
	assert.Equal(t, expected, actual.Error())

	lang = "xx"
	expected = "error: code = 400 reason = BAD_REQUEST message = metadata string conversion to int failed metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(err, lang)
	assert.Equal(t, expected, actual.Error())
}

func TestGetParameterErrorMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Parameter error, please check your parameters"
	actual := l.GetParameterErrorMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "参数错误,请检查您的参数"
	actual = l.GetParameterErrorMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetMetadataFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Get metadata failed"
	actual := l.GetMetadataFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "获取元数据失败"
	actual = l.GetMetadataFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetMetadataConversionFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "metadata string conversion to int failed"
	actual := l.GetMetadataConversionFailMsg("string", "int", lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "元数据 string 转换为 int 失败"
	actual = l.GetMetadataConversionFailMsg("string", "int", lang)
	assert.Equal(t, expected, actual)

	lang = "xx"
	expected = "metadata string conversion to int failed"
	actual = l.GetMetadataConversionFailMsg("string", "int", lang)
	assert.Equal(t, expected, actual)
}

func TestGetMissingMetadataMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Missing metadata {.Name}}, please check whether the transmission link is enabled for metadata transmission and pass the value"
	actual := l.GetMissingMetadataMsg("name", lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "缺少元数据 xxxx,请检查传输链路是否启用元数据传递,并传值。"
	actual = l.GetMissingMetadataMsg("xxxx", lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataFoundMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data not found"
	actual := l.GetDataFoundMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "找不到数据"
	actual = l.GetDataFoundMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataDuplicateMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "The data already exists, please do not add it repeatedly"
	actual := l.GetDataDuplicateMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "该数据已存在,请勿重复添加"
	actual = l.GetDataDuplicateMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataConstraintMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data constraint check failed, please check your parameters"
	actual := l.GetDataConstraintMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据约束检查失败，请检查您的参数"
	actual = l.GetDataConstraintMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataNotLoadedMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Database not loaded, please contact the administrator"
	actual := l.GetDataNotLoadedMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据库未加载，请联系管理员"
	actual = l.GetDataNotLoadedMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataNotSingularMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data error Not Singular, please contact the administrator"
	actual := l.GetDataNotSingularMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据出错了 Not Singular,请联系管理员"
	actual = l.GetDataNotSingularMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataValidationErrorMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data validation failed, please check your parameters"
	actual := l.GetDataValidationErrorMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据校验失败，请检查您的参数"
	actual = l.GetDataValidationErrorMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataErrorMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Data layer error, please contact the administrator"
	actual := l.GetDataErrorMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据层出错了,请联系管理员"
	actual = l.GetDataErrorMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheNotFoundMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Cache not found"
	actual := l.GetCacheNotFoundMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "缓存不存在"
	actual = l.GetCacheNotFoundMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCachePreMatchGetFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Pre-match cache get failed"
	actual := l.GetCachePreMatchGetFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "前置匹配获取缓存失败"
	actual = l.GetCachePreMatchGetFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheSetFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Set cache failed"
	actual := l.GetCacheSetFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "设置缓存失败"
	actual = l.GetCacheSetFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMSetFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Batch setting cache failed"
	actual := l.GetCacheMSetFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量设置缓存失败"
	actual = l.GetCacheMSetFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Delete cache failed"
	actual := l.GetCacheDelFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "删除缓存失败"
	actual = l.GetCacheDelFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCachePreMatchDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Pre-match delete cache failed"
	actual := l.GetCachePreMatchDelFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "前置匹配删除缓存失败"
	actual = l.GetCachePreMatchDelFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheFlushFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Flush cache failed"
	actual := l.GetCacheFlushFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "清空缓存失败"
	actual = l.GetCacheFlushFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMGetFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Batch get cache failed"
	actual := l.GetCacheMGetFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量获取缓存失败"
	actual = l.GetCacheMGetFailMsg(lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil)
	lang := "en-US"
	expected := "Batch delete cache failed"
	actual := l.GetCacheMDelFailMsg(lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量删除缓存失败"
	actual = l.GetCacheMDelFailMsg(lang)
	assert.Equal(t, expected, actual)
}
