// Package meta key
package meta

const (
	globalPrefix = "x-md-global-"
	localPrefix  = "x-md-local-"
)

const staffIDKey = globalPrefix + "staff-id"
const userIDKey = globalPrefix + "user-id"

// 客户端 ID
const clientIDKey = globalPrefix + "client-id"

// 签名
const signKey = globalPrefix + "sign"

// 认证信息
const authKey = globalPrefix + "auth"

// 平台
const platformKey = globalPrefix + "platform"

// 渠道
const channelKey = globalPrefix + "channel"

// 用户类型
const userTypeKey = globalPrefix + "user-type"

// 业务
const businessKey = globalPrefix + "business"

// 微信小程序 appid
const wechatMiniAppIDKey = globalPrefix + "wechat-mini-appid"

// 微信小程序 openid
const wechatMiniOpenIDKey = globalPrefix + "wechat-mini-openid"

// 微信公众号 appid
const wechatOfficialAccountAppIDKey = globalPrefix + "wechat-offiaccount-appid"

// 微信公众号 openid
const wechatOfficialAccountOpenIDKey = globalPrefix + "wechat-offiaccount-openid"

// 微信开放平台 appid
const wechatPlatformAppIDKey = globalPrefix + "wechat-platform-appid"

// 微信开放平台 openid
const wechatPlatformOpenIDKey = globalPrefix + "wechat-platform-openid"

// 企业微信 appid
const wechatWorkAppIDKey = globalPrefix + "wechat-work-appid"

// 企业微信 openid
const wechatWorkOpenIDKey = globalPrefix + "wechat-work-openid"

// 支付宝小程序 appid
const alipayMiniAppIDKey = globalPrefix + "alipay-mini-appid"

// 支付宝小程序 openid
const alipayMiniOpenIDKey = globalPrefix + "alipay-mini-openid"
