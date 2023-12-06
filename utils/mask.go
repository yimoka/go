// Package utils mask.go
package utils

import (
	"math"
	"strings"
)

// MaskType 敏感数据类型
type MaskType string

const (
	// MaskTypePhone 用于手机号
	MaskTypePhone MaskType = "phone"
	// MaskTypeEmail 用于邮箱账号
	MaskTypeEmail MaskType = "email"
	// MaskTypeIDCard 用于身份证号
	MaskTypeIDCard MaskType = "idCard"
	// MaskTypeBankCard 用于银行卡号
	MaskTypeBankCard MaskType = "bankCard"
	// MaskTypeOther 用于其他敏感数据
	MaskTypeOther MaskType = "other"
)

// Mask 敏感数据加星
func Mask(str string, maskType MaskType) string {
	switch maskType {
	case MaskTypePhone:
		return MaskPhone(str)
	case MaskTypeEmail:
		return MaskMail(str)
	case MaskTypeIDCard:
		return MaskIDCard(str)
	case MaskTypeBankCard:
		return MaskBankCard(str)
	default:
		return MaskOther(str)
	}
}

// MaskPhone 手机号加星 手机长度不确定，中间至少4位替换成 *  前面展示 1-3位 后1-4位  例如  138****1234
func MaskPhone(phone string) string {
	strLen := len(phone)
	if strLen <= 4 {
		return strings.Repeat("*", strLen)
	}
	mastStr := "****"
	if strLen <= 11 {
		showLen := strLen - 4
		// 前后一半后优先 向上取整
		endLen := int(math.Ceil(float64(showLen) / 2))
		startLen := showLen - endLen
		if startLen < 1 {
			return mastStr + phone[strLen-endLen:]
		}
		return phone[:startLen] + mastStr + phone[strLen-endLen:]
	}
	return phone[:3] + strings.Repeat("*", strLen-7) + phone[strLen-4:]
}

// MaskMail 邮箱加星
func MaskMail(mail string) string {
	arr := strings.Split(mail, "@")
	str := ""
	if len(arr[0]) > 0 {
		str += arr[0][:1]
	}
	str += "***@***"
	if len(arr) > 1 {
		sl := strings.Split(arr[1], ".")
		if len(sl) > 1 {
			str += "." + sl[len(sl)-1]
		}
	}
	return str
}

// TODO 待测定与完善

// MaskIDCard 身份证号加星
func MaskIDCard(idCard string) string {
	strLen := len(idCard)
	if strLen <= 8 {
		return strings.Repeat("*", strLen)
	}
	return idCard[:4] + strings.Repeat("*", strLen-8) + idCard[strLen-4:]
}

// TODO 待测定与完善

// MaskBankCard 银行卡号加星
func MaskBankCard(bankCard string) string {
	strLen := len(bankCard)
	if strLen <= 8 {
		return strings.Repeat("*", strLen)
	}
	return bankCard[:4] + strings.Repeat("*", strLen-8) + bankCard[strLen-4:]
}

// TODO 待测定与完善

// MaskOther 其他敏感数据加星
func MaskOther(str string) string {
	strLen := len(str)
	if strLen <= 4 {
		return strings.Repeat("*", strLen)
	}
	return str[:2] + strings.Repeat("*", strLen-4) + str[strLen-2:]
}
