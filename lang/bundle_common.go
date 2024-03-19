// Package lang bundle common
package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yimoka/go/config"
	"golang.org/x/text/language"
)

const (
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
)

var commonLangs = map[string]*config.Lang{
	"en": {
		Messages: []*config.LangMessage{
			{Id: dataNotFoundKey, Other: dataNotFoundMsg},
			{Id: dataDuplicateKey, Other: dataDuplicateMsg},
			{Id: dataConstraintKey, Other: dataConstraintMsg},
			{Id: dataNotLoadedKey, Other: dataNotLoadedMsg},
			{Id: dataNotSingularKey, Other: dataNotSingularMsg},
			{Id: dataValidationErrorKey, Other: dataValidationErrorMsg},
			{Id: dataErrorKey, Other: dataErrorMsg},
		}},
	"zh": {
		Messages: []*config.LangMessage{
			{Id: dataNotFoundKey, Other: "找不到数据"},
			{Id: dataDuplicateKey, Other: "该数据已存在,请勿重复添加"},
			{Id: dataConstraintKey, Other: "数据约束检查失败，请检查您的参数"},
			{Id: dataNotLoadedKey, Other: "数据库未加载，请联系管理员"},
			{Id: dataNotSingularKey, Other: "数据出错了 Not Singular,请联系管理员"},
			{Id: dataValidationErrorKey, Other: "数据校验失败，请检查您的参数"},
			{Id: dataErrorKey, Other: "数据层出错了,请联系管理员"},
		}},
	"ru": {
		Messages: []*config.LangMessage{
			{Id: dataNotFoundKey, Other: "Данные не найдены"},
			{Id: dataDuplicateKey, Other: "Эти данные уже существуют, пожалуйста, не добавляйте их повторно"},
			{Id: dataConstraintKey, Other: "Ошибка проверки ограничений данных, пожалуйста, проверьте ваши параметры"},
			{Id: dataNotLoadedKey, Other: "База данных не загружена, пожалуйста, свяжитесь с администратором"},
			{Id: dataNotSingularKey, Other: "Ошибка данных Not Singular, пожалуйста, свяжитесь с администратором"},
			{Id: dataValidationErrorKey, Other: "Ошибка проверки данных, пожалуйста, проверьте ваши параметры"},
			{Id: dataErrorKey, Other: "Ошибка слоя данных, пожалуйста, свяжитесь с администратором"},
		}},
}

// GetCommonBundle 获取公共语言包
func GetCommonBundle(langs map[string]*config.Lang) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	BundleMessage(bundle, commonLangs)
	BundleMessage(bundle, langs)
	return bundle
}

// GetNotDataFoundMsg 获取数据未找到消息
func GetNotDataFoundMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotFoundKey,
		DefaultMessage: &i18n.Message{ID: dataNotFoundKey, Other: dataNotFoundMsg},
	})
	if err != nil {
		return dataNotFoundMsg
	}
	return value
}

// GetNotDataDuplicateMsg 获取数据重复消息
func GetNotDataDuplicateMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataDuplicateKey,
		DefaultMessage: &i18n.Message{ID: dataDuplicateKey, Other: dataDuplicateMsg},
	})
	if err != nil {
		return dataDuplicateMsg
	}
	return value
}

// GetNotDataConstraintMsg 获取数据约束消息
func GetNotDataConstraintMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataConstraintKey,
		DefaultMessage: &i18n.Message{ID: dataConstraintKey, Other: dataConstraintMsg},
	})
	if err != nil {
		return dataConstraintMsg
	}
	return value
}

// GetNotDataNotLoadedMsg 获取数据未加载消息
func GetNotDataNotLoadedMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotLoadedKey,
		DefaultMessage: &i18n.Message{ID: dataNotLoadedKey, Other: dataNotLoadedMsg},
	})
	if err != nil {
		return dataNotLoadedMsg
	}
	return value
}

// GetNotDataNotSingularMsg 获取数据不是单数消息
func GetNotDataNotSingularMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataNotSingularKey,
		DefaultMessage: &i18n.Message{ID: dataNotSingularKey, Other: dataNotSingularMsg},
	})
	if err != nil {
		return dataNotSingularMsg
	}
	return value
}

// GetNotDataValidationErrorMsg 获取数据验证错误消息
func GetNotDataValidationErrorMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataValidationErrorKey,
		DefaultMessage: &i18n.Message{ID: dataValidationErrorKey, Other: dataValidationErrorMsg},
	})
	if err != nil {
		return dataValidationErrorMsg
	}
	return value
}

// GetNotDataErrorMsg 获取数据错误消息
func GetNotDataErrorMsg(bundle *i18n.Bundle, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	value, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      dataErrorKey,
		DefaultMessage: &i18n.Message{ID: dataErrorKey, Other: dataErrorMsg},
	})
	if err != nil {
		return dataErrorMsg
	}
	return value
}
