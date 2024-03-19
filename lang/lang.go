// Package lang fn
package lang

import (
	"context"
	"sort"
	"strconv"
	"strings"

	"github.com/yimoka/go/middleware/meta"
)

// GetStr 获取语言字符串
func GetStr(ctx context.Context) string {
	lang, err := meta.GetValue(ctx, "lang")
	if err == nil && lang != "" {
		return lang
	}
	lang = meta.GetAcceptLanguage(ctx)
	if lang != "" {
		meta.SetValue(ctx, "lang", lang)
		return lang
	}
	return ""
}

// GetArr 获取语言数组
func GetArr(ctx context.Context) []string {
	str := GetStr(ctx)
	if str == "" {
		return []string{}
	}
	// 转为数组
	lang := strings.Split(str, ",")
	// 根据权重排序
	sort.Slice(lang, func(i, j int) bool {
		iArr := strings.Split(lang[i], ";")
		iQ := 1.0
		if len(iArr) > 1 {
			iQArr := strings.Split(iArr[1], "=")
			if len(iQArr) > 1 {
				iQ, _ = strconv.ParseFloat(iQArr[1], 64)
			}
		}
		jArr := strings.Split(lang[j], ";")
		jQ := 1.0
		if len(jArr) > 1 {
			jQArr := strings.Split(jArr[1], "=")
			if len(jQArr) > 1 {
				jQ, _ = strconv.ParseFloat(jQArr[1], 64)
			}
		}
		return iQ > jQ
	})
	// 去掉权重
	for i, l := range lang {
		lang[i] = strings.Split(l, ";")[0]
	}
	return lang
}

// MatchContent 匹配语言内容 泛型  输入 map[string]T  输出 T
func MatchContent[T any](langMap map[string]T, lang []string) (T, bool) {
	// 备用值 当语言-地区匹配不到 但语言代码匹配到时使用
	spareValue := langMap["default"]
	// 是否匹配到备用
	spare := false
	for _, item := range lang {
		if v, ok := langMap[item]; ok {
			return v, true
		}
		// 折分 语言-地区
		l := ""
		// 匹配分割线 - _
		if strings.Contains(item, "-") {
			arr := strings.Split(item, "-")
			l = arr[0]
		} else if strings.Contains(item, "_") {
			arr := strings.Split(item, "_")
			l = arr[0]
		} else {
			continue
		}
		if v, ok := langMap[l]; ok {
			spareValue = v
			spare = true
		}
	}
	return spareValue, spare
}
