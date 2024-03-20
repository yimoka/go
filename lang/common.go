// Package lang bundle common
package lang

import (
	"bytes"
	"html/template"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yimoka/api/fault"
	"github.com/yimoka/go/config"
	"golang.org/x/text/language"
)

// CommonLang 公共语言包
type CommonLang struct {
	Bundle *i18n.Bundle
}

// GetCommonBundle 获取公共语言包
func GetCommonBundle(langs map[string]*config.Lang) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	BundleMessage(bundle, commonLangs)
	BundleMessage(bundle, langs)
	return bundle
}

// NewCommonLang 创建公共语言包
func NewCommonLang(langs map[string]*config.Lang) *CommonLang {
	return &CommonLang{
		Bundle: GetCommonBundle(langs),
	}
}

// HandleMetadataError 处理 metadata 错误
func (c *CommonLang) HandleMetadataError(err *errors.Error, langs ...string) error {
	if err == nil {
		return nil
	}
	if err.Metadata == nil {
		return fault.ErrorBadRequest(c.GetMetadataFailMsg(langs...))
	}
	source, sOk := err.Metadata["source"]
	target, tOk := err.Metadata["target"]

	if !sOk || !tOk {
		return fault.ErrorBadRequest(c.GetMetadataFailMsg(langs...))
	}
	return fault.ErrorBadRequest(c.GetMetadataConversionFailMsg(source, target, langs...))
}

const (
	// 公共
	parameterErrorKey            = "parameter_error"
	parameterErrorMsg            = "Parameter error, please check your parameters"
	getMetadataFailKey           = "get_metadata_fail"
	getMetadataFailMsg           = "Get metadata failed"
	getMetadataConversionFailKey = "get_metadata_int_fail"
	getMetadataConversionFailMsg = "metadata {{.Source}} conversion to {{.Target}} failed"
	missingMetadataKey           = "missing_metadata"
	missingMetadataMsg           = "Missing metadata {{.Name}}, please check whether the transmission link is enabled for metadata transmission and pass the value"
	encryptFailKey               = "encrypt_fail"
	encryptFailMsg               = "Encryption failed"
	paramCanNotEmptyKey          = "param_can_not_empty"
	paramCanNotEmptyMsg          = "Parameter {{.Name}} cannot be empty"
	notEditableKey               = "not_editable"
	notEditableMsg               = "The data is not editable"

	// 数据库
	dataNotFoundKey        = "data_not_found"
	dataNotFoundMsg        = "Data not found"
	dataDuplicateKey       = "data_duplicate_key"
	dataDuplicateMsg       = "The data already exists, please do not add it repeatedly"
	dataConstraintKey      = "data_constraint_error"
	dataConstraintMsg      = "Data constraint check failed, please check your parameters"
	dataNotLoadedKey       = "data_not_loaded"
	dataNotLoadedMsg       = "Database not loaded, please contact the administrator"
	dataNotSingularKey     = "data_not_singular"
	dataNotSingularMsg     = "Data error Not Singular, please contact the administrator"
	dataValidationErrorKey = "data_validation_error"
	dataValidationErrorMsg = "Data validation failed, please check your parameters"
	dataErrorKey           = "data_error"
	dataErrorMsg           = "Data layer error, please contact the administrator"
	// 缓存
	cacheNotFoundKey        = "cache_not_found"
	cacheNotFoundMsg        = "Cache not found"
	cachePreMatchGetFailKey = "cache_pre_match_get_fail"
	cachePreMatchGetFailMsg = "Pre-match cache get failed"
	cacheSetFailKey         = "cache_set_fail"
	cacheSetFailMsg         = "Set cache failed"
	cacheMSetFailKey        = "cache_m_set_fail"
	cacheMSetFailMsg        = "Batch setting cache failed"
	cacheDelFailKey         = "cache_del_fail"
	cacheDelFailMsg         = "Delete cache failed"
	cachePreMatchDelFailKey = "cache_pre_match_del_fail"
	cachePreMatchDelFailMsg = "Pre-match delete cache failed"
	cacheFlushFailKey       = "cache_flush_fail"
	cacheFlushFailMsg       = "Flush cache failed"
	cacheMGetFailKey        = "cache_m_get_fail"
	cacheMGetFailMsg        = "Batch get cache failed"
	cacheMDelFailKey        = "cache_m_del_fail"
	cacheMDelFailMsg        = "Batch delete cache failed"
)

var commonLangs = map[string]*config.Lang{
	"en": {
		Messages: []*config.LangMessage{
			{Id: parameterErrorKey, Other: parameterErrorMsg},
			{Id: getMetadataFailKey, Other: getMetadataFailMsg},
			{Id: getMetadataConversionFailKey, Other: getMetadataConversionFailMsg},
			{Id: missingMetadataKey, Other: missingMetadataMsg},
			{Id: encryptFailKey, Other: encryptFailMsg},
			{Id: paramCanNotEmptyKey, Other: paramCanNotEmptyMsg},
			{Id: notEditableKey, Other: notEditableMsg},

			{Id: dataNotFoundKey, Other: dataNotFoundMsg},
			{Id: dataDuplicateKey, Other: dataDuplicateMsg},
			{Id: dataConstraintKey, Other: dataConstraintMsg},
			{Id: dataNotLoadedKey, Other: dataNotLoadedMsg},
			{Id: dataNotSingularKey, Other: dataNotSingularMsg},
			{Id: dataValidationErrorKey, Other: dataValidationErrorMsg},
			{Id: dataErrorKey, Other: dataErrorMsg},

			{Id: cacheNotFoundKey, Other: cacheNotFoundMsg},
			{Id: cachePreMatchGetFailKey, Other: cachePreMatchGetFailMsg},
			{Id: cacheSetFailKey, Other: cacheSetFailMsg},
			{Id: cacheMSetFailKey, Other: cacheMSetFailMsg},
			{Id: cacheDelFailKey, Other: cacheDelFailMsg},
			{Id: cachePreMatchDelFailKey, Other: cachePreMatchDelFailMsg},
			{Id: cacheFlushFailKey, Other: cacheFlushFailMsg},
			{Id: cacheMGetFailKey, Other: cacheMGetFailMsg},
			{Id: cacheMDelFailKey, Other: cacheMDelFailMsg},
		}},
	"zh": {
		Messages: []*config.LangMessage{
			{Id: parameterErrorKey, Other: "参数错误,请检查您的参数"},
			{Id: getMetadataFailKey, Other: "获取元数据失败"},
			{Id: getMetadataConversionFailKey, Other: "元数据 {{.Source}} 转换为 {{.Target}} 失败"},
			{Id: missingMetadataKey, Other: "缺少元数据 {{.Name}},请检查传输链路是否启用元数据传递,并传值。"},
			{Id: encryptFailKey, Other: "加密失败"},
			{Id: paramCanNotEmptyKey, Other: "参数 {{.Name}} 不能为空"},
			{Id: notEditableKey, Other: "数据不可编辑"},

			{Id: dataNotFoundKey, Other: "找不到数据"},
			{Id: dataDuplicateKey, Other: "该数据已存在,请勿重复添加"},
			{Id: dataConstraintKey, Other: "数据约束检查失败，请检查您的参数"},
			{Id: dataNotLoadedKey, Other: "数据库未加载，请联系管理员"},
			{Id: dataNotSingularKey, Other: "数据出错了 Not Singular,请联系管理员"},
			{Id: dataValidationErrorKey, Other: "数据校验失败，请检查您的参数"},
			{Id: dataErrorKey, Other: "数据层出错了,请联系管理员"},

			{Id: cacheNotFoundKey, Other: "缓存不存在"},
			{Id: cachePreMatchGetFailKey, Other: "前置匹配获取缓存失败"},
			{Id: cacheSetFailKey, Other: "设置缓存失败"},
			{Id: cacheMSetFailKey, Other: "批量设置缓存失败"},
			{Id: cacheDelFailKey, Other: "删除缓存失败"},
			{Id: cachePreMatchDelFailKey, Other: "前置匹配删除缓存失败"},
			{Id: cacheFlushFailKey, Other: "清空缓存失败"},
			{Id: cacheMGetFailKey, Other: "批量获取缓存失败"},
			{Id: cacheMDelFailKey, Other: "批量删除缓存失败"},
		}},
	"ru": {
		Messages: []*config.LangMessage{
			{Id: parameterErrorKey, Other: "Ошибка параметра, пожалуйста, проверьте ваши параметры"},
			{Id: getMetadataFailKey, Other: "Ошибка получения метаданных"},
			{Id: getMetadataConversionFailKey, Other: "Ошибка преобразования метаданных {{.Source}} в {{.Target}}"},
			{Id: missingMetadataKey, Other: "Отсутствует метаданные {{.Name}}, пожалуйста, проверьте, включена ли передача метаданных в цепи передачи и передайте значение"},
			{Id: encryptFailKey, Other: "Ошибка шифрования"},
			{Id: paramCanNotEmptyKey, Other: "Параметр {{.Name}} не может быть пустым"},
			{Id: notEditableKey, Other: "Данные нельзя редактировать"},

			{Id: dataNotFoundKey, Other: "Данные не найдены"},
			{Id: dataDuplicateKey, Other: "Эти данные уже существуют, пожалуйста, не добавляйте их повторно"},
			{Id: dataConstraintKey, Other: "Ошибка проверки ограничений данных, пожалуйста, проверьте ваши параметры"},
			{Id: dataNotLoadedKey, Other: "База данных не загружена, пожалуйста, свяжитесь с администратором"},
			{Id: dataNotSingularKey, Other: "Ошибка данных Not Singular, пожалуйста, свяжитесь с администратором"},
			{Id: dataValidationErrorKey, Other: "Ошибка проверки данных, пожалуйста, проверьте ваши параметры"},
			{Id: dataErrorKey, Other: "Ошибка слоя данных, пожалуйста, свяжитесь с администратором"},

			{Id: cacheNotFoundKey, Other: "Кэш не найден"},
			{Id: cachePreMatchGetFailKey, Other: "Предварительное сопоставление получения кэша не удалось"},
			{Id: cacheSetFailKey, Other: "Ошибка установки кэша"},
			{Id: cacheMSetFailKey, Other: "Ошибка установки кэша"},
			{Id: cacheDelFailKey, Other: "Ошибка удаления кэша"},
			{Id: cachePreMatchDelFailKey, Other: "Предварительное сопоставление удаления кэша не удалось"},
			{Id: cacheFlushFailKey, Other: "Ошибка очистки кэша"},
			{Id: cacheMGetFailKey, Other: "Ошибка получения кэша"},
			{Id: cacheMDelFailKey, Other: "Ошибка удаления кэша"},
		}},
}

// GetParameterErrorMsg 获取参数错误消息
func (c *CommonLang) GetParameterErrorMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      parameterErrorKey,
		DefaultMessage: &i18n.Message{ID: parameterErrorKey, Other: parameterErrorMsg},
	})
	if err != nil {
		return parameterErrorMsg
	}
	return value
}

// GetMetadataFailMsg 获取获取元数据失败消息
func (c *CommonLang) GetMetadataFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      getMetadataFailKey,
		DefaultMessage: &i18n.Message{ID: getMetadataFailKey, Other: getMetadataFailMsg},
	})
	if err != nil {
		return getMetadataFailMsg
	}
	return value
}

// GetMetadataConversionFailMsg 获取获取元数据转换失败消息
func (c *CommonLang) GetMetadataConversionFailMsg(source string, target string, langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      getMetadataConversionFailKey,
		DefaultMessage: &i18n.Message{ID: getMetadataConversionFailKey, Other: getMetadataConversionFailMsg},
		TemplateData:   map[string]string{"Source": source, "Target": target},
	})
	if err != nil {
		t := template.New("getMetadataConversionFailMsg")
		t, pErr := t.Parse(getMetadataConversionFailMsg)
		if pErr != nil {
			return getMetadataConversionFailMsg
		}
		p := map[string]string{"Source": source, "Target": target}
		var buf bytes.Buffer
		eErr := t.Execute(&buf, p)
		if eErr != nil {
			return getMetadataConversionFailMsg
		}
		return buf.String()
	}
	return value
}

// GetMissingMetadataMsg 获取缺少元数据消息
func (c *CommonLang) GetMissingMetadataMsg(name string, langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      missingMetadataKey,
		DefaultMessage: &i18n.Message{ID: missingMetadataKey, Other: missingMetadataMsg},
		TemplateData:   map[string]string{"Name": name},
	})
	if err != nil {
		t := template.New("missingMetadataMsg")
		t, pErr := t.Parse(missingMetadataMsg)
		if pErr != nil {
			return missingMetadataMsg
		}
		p := map[string]string{"Name": name}
		var buf bytes.Buffer
		eErr := t.Execute(&buf, p)
		if eErr != nil {
			return missingMetadataMsg
		}
		return buf.String()
	}
	return value
}

// GetEncryptFailMsg 获取加密失败消息
func (c *CommonLang) GetEncryptFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      encryptFailKey,
		DefaultMessage: &i18n.Message{ID: encryptFailKey, Other: encryptFailMsg},
	})
	if err != nil {
		return encryptFailMsg
	}
	return value
}

// GetParamCanNotEmptyMsg 获取参数不能为空消息
func (c *CommonLang) GetParamCanNotEmptyMsg(name string, langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      paramCanNotEmptyKey,
		DefaultMessage: &i18n.Message{ID: paramCanNotEmptyKey, Other: paramCanNotEmptyMsg},
		TemplateData:   map[string]string{"Name": name},
	})
	if err != nil {
		t := template.New("paramCanNotEmptyMsg")
		t, pErr := t.Parse(paramCanNotEmptyMsg)
		if pErr != nil {
			return paramCanNotEmptyMsg
		}
		p := map[string]string{"Name": name}
		var buf bytes.Buffer
		eErr := t.Execute(&buf, p)
		if eErr != nil {
			return paramCanNotEmptyMsg
		}
		return buf.String()
	}
	return value
}

// GetNotEditableMsg 获取数据不可编辑消息
func (c *CommonLang) GetNotEditableMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      notEditableKey,
		DefaultMessage: &i18n.Message{ID: notEditableKey, Other: notEditableMsg},
	})
	if err != nil {
		return notEditableMsg
	}
	return value
}

// GetDataFoundMsg 获取数据未找到消息
func (c *CommonLang) GetDataFoundMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotFoundKey,
		DefaultMessage: &i18n.Message{ID: dataNotFoundKey, Other: dataNotFoundMsg},
	})
	if err != nil {
		return dataNotFoundMsg
	}
	return value
}

// GetDataDuplicateMsg 获取数据重复消息
func (c *CommonLang) GetDataDuplicateMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataDuplicateKey,
		DefaultMessage: &i18n.Message{ID: dataDuplicateKey, Other: dataDuplicateMsg},
	})
	if err != nil {
		return dataDuplicateMsg
	}
	return value
}

// GetDataConstraintMsg 获取数据约束消息
func (c *CommonLang) GetDataConstraintMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataConstraintKey,
		DefaultMessage: &i18n.Message{ID: dataConstraintKey, Other: dataConstraintMsg},
	})
	if err != nil {
		return dataConstraintMsg
	}
	return value
}

// GetDataNotLoadedMsg 获取数据未加载消息
func (c *CommonLang) GetDataNotLoadedMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotLoadedKey,
		DefaultMessage: &i18n.Message{ID: dataNotLoadedKey, Other: dataNotLoadedMsg},
	})
	if err != nil {
		return dataNotLoadedMsg
	}
	return value
}

// GetDataNotSingularMsg 获取数据不是单数消息
func (c *CommonLang) GetDataNotSingularMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotSingularKey,
		DefaultMessage: &i18n.Message{ID: dataNotSingularKey, Other: dataNotSingularMsg},
	})
	if err != nil {
		return dataNotSingularMsg
	}
	return value
}

// GetDataValidationErrorMsg 获取数据验证错误消息
func (c *CommonLang) GetDataValidationErrorMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataValidationErrorKey,
		DefaultMessage: &i18n.Message{ID: dataValidationErrorKey, Other: dataValidationErrorMsg},
	})
	if err != nil {
		return dataValidationErrorMsg
	}
	return value
}

// GetDataErrorMsg 获取数据错误消息
func (c *CommonLang) GetDataErrorMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataErrorKey,
		DefaultMessage: &i18n.Message{ID: dataErrorKey, Other: dataErrorMsg},
	})
	if err != nil {
		return dataErrorMsg
	}
	return value
}

// GetCacheNotFoundMsg 获取缓存未找到消息
func (c *CommonLang) GetCacheNotFoundMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheNotFoundKey,
		DefaultMessage: &i18n.Message{ID: cacheNotFoundKey, Other: cacheNotFoundMsg},
	})
	if err != nil {
		return cacheNotFoundMsg
	}
	return value
}

// GetCachePreMatchGetFailMsg 获取缓存前置匹配获取失败消息
func (c *CommonLang) GetCachePreMatchGetFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cachePreMatchGetFailKey,
		DefaultMessage: &i18n.Message{ID: cachePreMatchGetFailKey, Other: cachePreMatchGetFailMsg},
	})
	if err != nil {
		return cachePreMatchGetFailMsg
	}
	return value
}

// GetCacheSetFailMsg 获取缓存设置失败消息
func (c *CommonLang) GetCacheSetFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheSetFailKey,
		DefaultMessage: &i18n.Message{ID: cacheSetFailKey, Other: cacheSetFailMsg},
	})
	if err != nil {
		return cacheSetFailMsg
	}
	return value
}

// GetCacheMSetFailMsg 获取缓存批量设置失败消息
func (c *CommonLang) GetCacheMSetFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheMSetFailKey,
		DefaultMessage: &i18n.Message{ID: cacheMSetFailKey, Other: cacheMSetFailMsg},
	})
	if err != nil {
		return cacheMSetFailMsg
	}
	return value
}

// GetCacheDelFailMsg 获取缓存删除失败消息
func (c *CommonLang) GetCacheDelFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheDelFailKey,
		DefaultMessage: &i18n.Message{ID: cacheDelFailKey, Other: cacheDelFailMsg},
	})
	if err != nil {
		return cacheDelFailMsg
	}
	return value
}

// GetCachePreMatchDelFailMsg 获取缓存前置匹配删除失败消息
func (c *CommonLang) GetCachePreMatchDelFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cachePreMatchDelFailKey,
		DefaultMessage: &i18n.Message{ID: cachePreMatchDelFailKey, Other: cachePreMatchDelFailMsg},
	})
	if err != nil {
		return cachePreMatchDelFailMsg
	}
	return value
}

// GetCacheFlushFailMsg 获取缓存清空失败消息
func (c *CommonLang) GetCacheFlushFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheFlushFailKey,
		DefaultMessage: &i18n.Message{ID: cacheFlushFailKey, Other: cacheFlushFailMsg},
	})
	if err != nil {
		return cacheFlushFailMsg
	}
	return value
}

// GetCacheMGetFailMsg 获取缓存批量获取失败消息
func (c *CommonLang) GetCacheMGetFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheMGetFailKey,
		DefaultMessage: &i18n.Message{ID: cacheMGetFailKey, Other: cacheMGetFailMsg},
	})
	if err != nil {
		return cacheMGetFailMsg
	}
	return value
}

// GetCacheMDelFailMsg 获取缓存批量删除失败消息
func (c *CommonLang) GetCacheMDelFailMsg(langs ...string) string {
	localizer := i18n.NewLocalizer(c.Bundle, langs...)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      cacheMDelFailKey,
		DefaultMessage: &i18n.Message{ID: cacheMDelFailKey, Other: cacheMDelFailMsg},
	})
	if err != nil {
		return cacheMDelFailMsg
	}
	return value
}
