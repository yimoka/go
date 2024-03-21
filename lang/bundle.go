// Package lang bundle
package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/lo"
	"github.com/yimoka/go/config"
	"golang.org/x/text/language"
)

// GetBundle 获取语言包
func GetBundle(langs map[string]*config.Lang) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	LoadMessageForConfig(bundle, langs)
	return bundle
}

// LoadMessageForConfig 从配置加载语言包
func LoadMessageForConfig(bundle *i18n.Bundle, langMap map[string]*config.Lang) {
	if langMap == nil {
		return
	}
	for key, l := range langMap {
		tag, err := language.Parse(key)
		if err == nil {
			msgs := lo.Map(l.Messages, func(item *config.LangMessage, index int) *i18n.Message { return MessageToI18n(item) })
			_ = bundle.AddMessages(tag, msgs...)
		}
	}
}

// LoadMessage 从加载语言包
func LoadMessage(bundle *i18n.Bundle, langMap map[language.Tag]map[MsgKey]*i18n.Message) {
	if langMap == nil {
		return
	}
	for tag, l := range langMap {
		_ = bundle.AddMessages(tag, lo.Values(l)...)
	}
}

// MessageToI18n LangMessage 转为i18n.Message
func MessageToI18n(configMes *config.LangMessage) *i18n.Message {
	hash := configMes.Hash
	if hash == "" {
		hash = configMes.Id
	}
	return &i18n.Message{
		ID:          configMes.Id,
		Hash:        hash,
		Description: configMes.Description,
		LeftDelim:   configMes.LeftDelim,
		RightDelim:  configMes.RightDelim,
		Zero:        configMes.Zero,
		One:         configMes.One,
		Two:         configMes.Two,
		Few:         configMes.Few,
		Many:        configMes.Many,
		Other:       configMes.Other,
	}
}
