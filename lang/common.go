// Package lang bundle common
package lang

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yimoka/api/fault"
	"github.com/yimoka/go/config"
	"golang.org/x/text/language"
)

// CommonLang 公共语言包
type CommonLang struct {
	Bundle *i18n.Bundle
	log    *log.Helper
}

// GetCommonBundle 获取公共语言包
func GetCommonBundle(langMapConfig map[string]*config.Lang) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	LoadMessage(bundle, langMap)
	LoadMessageForConfig(bundle, langMapConfig)
	return bundle
}

// NewCommonLang 创建公共语言包
func NewCommonLang(langMapConfig map[string]*config.Lang, logger log.Logger) *CommonLang {
	return &CommonLang{
		Bundle: GetCommonBundle(langMapConfig),
		log:    log.NewHelper(log.With(logger, "layer", "commonLang")),
	}
}

// HandleMetadataError 处理 metadata 错误
func (c *CommonLang) HandleMetadataError(ctx context.Context, err *errors.Error, langs ...string) error {
	if err == nil {
		return nil
	}
	if err.Metadata == nil {
		return fault.ErrorBadRequest(c.GetMetadataFailMsg(ctx, langs...))
	}
	source, sOk := err.Metadata["source"]
	target, tOk := err.Metadata["target"]

	if !sOk || !tOk {
		return fault.ErrorBadRequest(c.GetMetadataFailMsg(ctx, langs...))
	}
	return fault.ErrorBadRequest(c.GetMetadataConversionFailMsg(ctx, source, target, langs...))
}

// GetLocalizer _
func (c *CommonLang) getLocalizer(ctx context.Context, langs ...string) *i18n.Localizer {
	if len(langs) == 0 {
		return i18n.NewLocalizer(c.Bundle, GetAcceptArr(ctx)...)
	}
	return i18n.NewLocalizer(c.Bundle, langs...)
}

// getMsg _
func (c *CommonLang) getMsg(ctx context.Context, key MsgKey, templateData interface{}, langs ...string) string {
	localizer := c.getLocalizer(ctx, langs...)
	v, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:      key.String(),
		TemplateData:   templateData,
		DefaultMessage: dfMsgMap[key],
	})
	if err != nil {
		c.log.Error("GetMsg", err)
		return HandleError(key, dfMsgMap[key], templateData)
	}
	return v
}

// GetParameterErrorMsg 获取参数错误消息
func (c *CommonLang) GetParameterErrorMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, parameterErrorKey, nil, langs...)
}

// GetMetadataFailMsg 获取获取元数据失败消息
func (c *CommonLang) GetMetadataFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, getMetadataFailKey, nil, langs...)
}

// GetMetadataConversionFailMsg 获取获取元数据转换失败消息
func (c *CommonLang) GetMetadataConversionFailMsg(ctx context.Context, source string, target string, langs ...string) string {
	return c.getMsg(ctx, getMetadataConversionFailKey, map[string]string{"Source": source, "Target": target}, langs...)
}

// GetMissingMetadataMsg 获取缺少元数据消息
func (c *CommonLang) GetMissingMetadataMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, missingMetadataKey, map[string]string{"Name": name}, langs...)
}

// GetEncryptFailMsg 获取加密失败消息
func (c *CommonLang) GetEncryptFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, encryptFailKey, nil, langs...)
}

// GetDecryptFailMsg 获取解密失败消息
func (c *CommonLang) GetDecryptFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, decryptFailKey, nil, langs...)
}

// GetParamCanNotEmptyMsg 获取参数不能为空消息
func (c *CommonLang) GetParamCanNotEmptyMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, paramCanNotEmptyKey, map[string]string{"Name": name}, langs...)
}

// GetNotEditableMsg 获取数据不可编辑消息
func (c *CommonLang) GetNotEditableMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, notEditableKey, nil, langs...)
}

// GetRequestErrorMsg 获取请求错误消息
func (c *CommonLang) GetRequestErrorMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, requestErrorKey, nil, langs...)
}

// GetPleaseLoginMsg 获取请登录消息
func (c *CommonLang) GetPleaseLoginMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, pleaseLoginKey, nil, langs...)
}

// GetNeedReLoginMsg 获取需要重新登录消息
func (c *CommonLang) GetNeedReLoginMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, needReLoginKey, nil, langs...)
}

// GetAccountDisabledMsg 获取账号已禁用消息
func (c *CommonLang) GetAccountDisabledMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, accountDisabledKey, nil, langs...)
}

// GetPleaseChangePasswordMsg 获取请修改密码消息
func (c *CommonLang) GetPleaseChangePasswordMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, pleaseChangePasswordKey, nil, langs...)
}

// GetPasswordErrorMsg 获取密码错误消息
func (c *CommonLang) GetPasswordErrorMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, passwordErrorKey, nil, langs...)
}

// GetNoPermissionMsg 获取没有权限消息
func (c *CommonLang) GetNoPermissionMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, noPermissionKey, nil, langs...)
}

// GetNotConfiguredMsg 获取未配置消息
func (c *CommonLang) GetNotConfiguredMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, notConfiguredKey, map[string]string{"Name": name}, langs...)
}

// GetCanNotEmptyMsg 获取不能为空消息
func (c *CommonLang) GetCanNotEmptyMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, canNotEmptyKey, map[string]string{"Name": name}, langs...)
}

// GetExpiredMsg 获取已过期消息
func (c *CommonLang) GetExpiredMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, expiredKey, map[string]string{"Name": name}, langs...)
}

// GetDataAbnormalMsg 获取数据异常消息
func (c *CommonLang) GetDataAbnormalMsg(ctx context.Context, name string, langs ...string) string {
	return c.getMsg(ctx, dataAbnormalKey, map[string]string{"Name": name}, langs...)
}

// GetDataFoundMsg 获取数据未找到消息
func (c *CommonLang) GetDataNotFoundMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataNotFoundKey, nil, langs...)
}

// GetDataDuplicateMsg 获取数据重复消息
func (c *CommonLang) GetDataDuplicateKeyMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataDuplicateKey, nil, langs...)
}

// GetDataConstraintMsg 获取数据约束消息
func (c *CommonLang) GetDataConstraintMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataConstraintKey, nil, langs...)
}

// GetDataNotLoadedMsg 获取数据未加载消息
func (c *CommonLang) GetDataNotLoadedMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataNotLoadedKey, nil, langs...)
}

// GetDataNotSingularMsg 获取数据不是单数消息
func (c *CommonLang) GetDataNotSingularMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataNotSingularKey, nil, langs...)
}

// GetDataValidationErrorMsg 获取数据验证错误消息
func (c *CommonLang) GetDataValidationErrorMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataValidationErrorKey, nil, langs...)
}

// GetDataErrorMsg 获取数据错误消息
func (c *CommonLang) GetDataErrorMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, dataErrorKey, nil, langs...)
}

// GetCacheNotFoundMsg 获取缓存未找到消息
func (c *CommonLang) GetCacheNotFoundMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheNotFoundKey, nil, langs...)
}

// GetCachePreMatchGetFailMsg 获取缓存前置匹配获取失败消息
func (c *CommonLang) GetCachePreMatchGetFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cachePreMatchGetFailKey, nil, langs...)
}

// GetCacheSetFailMsg 获取缓存设置失败消息
func (c *CommonLang) GetCacheSetFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheSetFailKey, nil, langs...)
}

// GetCacheMSetFailMsg 获取缓存批量设置失败消息
func (c *CommonLang) GetCacheMSetFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheMSetFailKey, nil, langs...)
}

// GetCacheDelFailMsg 获取缓存删除失败消息
func (c *CommonLang) GetCacheDelFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheDelFailKey, nil, langs...)
}

// GetCachePreMatchDelFailMsg 获取缓存前置匹配删除失败消息
func (c *CommonLang) GetCachePreMatchDelFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cachePreMatchDelFailKey, nil, langs...)
}

// GetCacheFlushFailMsg 获取缓存清空失败消息
func (c *CommonLang) GetCacheFlushFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheFlushFailKey, nil, langs...)
}

// GetCacheMGetFailMsg 获取缓存批量获取失败消息
func (c *CommonLang) GetCacheMGetFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheMGetFailKey, nil, langs...)
}

// GetCacheMDelFailMsg 获取缓存批量删除失败消息
func (c *CommonLang) GetCacheMDelFailMsg(ctx context.Context, langs ...string) string {
	return c.getMsg(ctx, cacheMDelFailKey, nil, langs...)
}
