// Package meta common
package meta

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
)

// GetClientID 获取客户端 id
func GetClientID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, clientIDKey)
}

// GetSign 获取签名
func GetSign(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, signKey)
}

// GetAuth 获取认证信息
func GetAuth(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, authKey)
}

// SetAuth 设置认证信息
func SetAuth(ctx context.Context, auth string) context.Context {
	return SetValue(ctx, authKey, auth)
}

// GetPlatform 获取平台
func GetPlatform(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, platformKey)
}

// GetChannel 获取渠道
func GetChannel(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, channelKey)
}

// GetUserType 获取用户类型
func GetUserType(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, userTypeKey)
}

// GetBusiness 获取业务
func GetBusiness(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, businessKey)
}

// GetWechatMiniAppID 获取微信小程序 appid
func GetWechatMiniAppID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatMiniAppIDKey)
}

// GetWechatMiniOpenID 获取微信小程序 openid
func GetWechatMiniOpenID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatMiniOpenIDKey)
}

// GetWechatOfficialAccountAppID 获取微信公众号 appid
func GetWechatOfficialAccountAppID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatOfficialAccountAppIDKey)
}

// GetWechatOfficialAccountOpenID 获取微信公众号 openid
func GetWechatOfficialAccountOpenID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatOfficialAccountOpenIDKey)
}

// GetWechatPlatformAppID 获取微信开放平台 appid
func GetWechatPlatformAppID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatPlatformAppIDKey)
}

// GetWechatPlatformOpenID 获取微信开放平台 openid
func GetWechatPlatformOpenID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatPlatformOpenIDKey)
}

// GetWechatWorkAppID 获取企业微信 appid
func GetWechatWorkAppID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatWorkAppIDKey)
}

// GetWechatWorkOpenID 获取企业微信 openid
func GetWechatWorkOpenID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, wechatWorkOpenIDKey)
}

// GetAlipayMiniAppID 获取支付宝小程序 appid
func GetAlipayMiniAppID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, alipayMiniAppIDKey)
}

// GetAlipayMiniOpenID 获取支付宝小程序 openid
func GetAlipayMiniOpenID(ctx context.Context) (string, *errors.Error) {
	return GetValue(ctx, alipayMiniOpenIDKey)
}
