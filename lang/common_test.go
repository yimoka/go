// nolint:goconst
package lang

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
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
			{Id: dataNotFoundKey.String(), Other: "新的 找不到数据"},
		}},
	}, log.DefaultLogger)
	localizer := i18n.NewLocalizer(l.Bundle, "zh")
	value, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: "add"})
	assert.Equal(t, "新增的", value)
	lang := "zh"
	expected := "新的 找不到数据"
	actual := l.GetDataNotFoundMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
	actual = l.GetDataDuplicateKeyMsg(context.Background(), "zh")
	assert.Equal(t, "该数据已存在,请勿重复添加", actual)
}

func TestHandleMetadataError(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	err := fault.ErrorBadRequest("Get metadata failed")
	expected := "error: code = 400 reason = BAD_REQUEST message = Get metadata failed metadata = map[] cause = <nil>"
	actual := l.HandleMetadataError(context.Background(), err, lang)
	assert.Equal(t, expected, actual.Error())

	err.Metadata = map[string]string{"source": "string", "target": "int"}
	expected = "error: code = 400 reason = BAD_REQUEST message = metadata string conversion to int failed metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(context.Background(), err, lang)
	assert.Equal(t, expected, actual.Error())

	lang = "zh-CN"
	expected = "error: code = 400 reason = BAD_REQUEST message = 元数据 string 转换为 int 失败 metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(context.Background(), err, lang)
	assert.Equal(t, expected, actual.Error())

	lang = "xx"
	expected = "error: code = 400 reason = BAD_REQUEST message = metadata string conversion to int failed metadata = map[] cause = <nil>"
	actual = l.HandleMetadataError(context.Background(), err, lang)
	assert.Equal(t, expected, actual.Error())
}

func TestGetParameterErrorMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Parameter error, please check your parameters"
	actual := l.GetParameterErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "参数错误,请检查您的参数"
	actual = l.GetParameterErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetMetadataFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Get metadata failed"
	actual := l.GetMetadataFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "获取元数据失败"
	actual = l.GetMetadataFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetMetadataConversionFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "metadata string conversion to int failed"
	actual := l.GetMetadataConversionFailMsg(context.Background(), "string", "int", lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "元数据 string 转换为 int 失败"
	actual = l.GetMetadataConversionFailMsg(context.Background(), "string", "int", lang)
	assert.Equal(t, expected, actual)

	lang = "xx"
	expected = "metadata string conversion to int failed"
	actual = l.GetMetadataConversionFailMsg(context.Background(), "string", "int", lang)
	assert.Equal(t, expected, actual)
}

func TestGetMissingMetadataMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Missing metadata name, please check whether the transmission link is enabled for metadata transmission and pass the value"
	actual := l.GetMissingMetadataMsg(context.Background(), "name", lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "缺少元数据 xxxx,请检查传输链路是否启用元数据传递,并传值。"
	actual = l.GetMissingMetadataMsg(context.Background(), "xxxx", lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataNotFoundMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Data not found"
	actual := l.GetDataNotFoundMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "找不到数据"
	actual = l.GetDataNotFoundMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataDuplicateKeyMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "The data already exists, please do not add it repeatedly"
	actual := l.GetDataDuplicateKeyMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "该数据已存在,请勿重复添加"
	actual = l.GetDataDuplicateKeyMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataConstraintMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Data constraint check failed, please check your parameters"
	actual := l.GetDataConstraintMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据约束检查失败，请检查您的参数"
	actual = l.GetDataConstraintMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataNotLoadedMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Database not loaded, please contact the administrator"
	actual := l.GetDataNotLoadedMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据库未加载，请联系管理员"
	actual = l.GetDataNotLoadedMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataNotSingularMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Data error Not Singular, please contact the administrator"
	actual := l.GetDataNotSingularMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据出错了 Not Singular,请联系管理员"
	actual = l.GetDataNotSingularMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataValidationErrorMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Data validation failed, please check your parameters"
	actual := l.GetDataValidationErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据校验失败，请检查您的参数"
	actual = l.GetDataValidationErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetDataErrorMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Data layer error, please contact the administrator"
	actual := l.GetDataErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "数据层出错了,请联系管理员"
	actual = l.GetDataErrorMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheNotFoundMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Cache not found"
	actual := l.GetCacheNotFoundMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "缓存不存在"
	actual = l.GetCacheNotFoundMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCachePreMatchGetFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Pre-match cache get failed"
	actual := l.GetCachePreMatchGetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "前置匹配获取缓存失败"
	actual = l.GetCachePreMatchGetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheSetFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Set cache failed"
	actual := l.GetCacheSetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "设置缓存失败"
	actual = l.GetCacheSetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMSetFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Batch setting cache failed"
	actual := l.GetCacheMSetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量设置缓存失败"
	actual = l.GetCacheMSetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Delete cache failed"
	actual := l.GetCacheDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "删除缓存失败"
	actual = l.GetCacheDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCachePreMatchDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Pre-match delete cache failed"
	actual := l.GetCachePreMatchDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "前置匹配删除缓存失败"
	actual = l.GetCachePreMatchDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheFlushFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Flush cache failed"
	actual := l.GetCacheFlushFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "清空缓存失败"
	actual = l.GetCacheFlushFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMGetFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Batch get cache failed"
	actual := l.GetCacheMGetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量获取缓存失败"
	actual = l.GetCacheMGetFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}

func TestGetCacheMDelFailMsg(t *testing.T) {
	l := NewCommonLang(nil, log.DefaultLogger)
	lang := "en-US"
	expected := "Batch delete cache failed"
	actual := l.GetCacheMDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)

	lang = "zh-CN"
	expected = "批量删除缓存失败"
	actual = l.GetCacheMDelFailMsg(context.Background(), lang)
	assert.Equal(t, expected, actual)
}
