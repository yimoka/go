// Package utils xss 过滤
package utils

import (
	"github.com/microcosm-cc/bluemonday"
)

// XSS 过滤
func XSS(s string) string {
	p := bluemonday.NewPolicy()
	// 允许常见的标签
	p.AllowElements(
		"p", "a", "span", "div", "h1", "h2", "h3", "h4", "h5", "h6",
		"article", "aside", "audio", "b", "br", "button", "code",
		"del", "em", "embed", "header", "hr", "i", "label", "link",
		"main", "nav", "pre", "section", "small", "strong", "pre",
		"source", "sub", "sup", "u", "video",
	)
	p.AllowAttrs("href").OnElements("a")
	// 媒体标签的常见属性
	p.AllowAttrs(
		"src", "controls", "autoplay", "loop", "muted", "poster",
		"preload", "width", "height", "controlsList", "playsinline",
		"webkit-playsinline", "crossorigin", "allowfullscreen",
		"allowpaymentrequest", "disablepictureinpicture",
		"disableRemotePlayback",
	).OnElements("img", "video", "audio", "source")
	p.AllowStandardAttributes()
	p.AllowStandardURLs()
	p.AllowImages()
	p.AllowDataURIImages()
	p.AllowLists()
	p.AllowTables()
	p.AllowRelativeURLs(true)

	// 富文本的常见样式属性
	p.AllowStyles(
		"color", "font-family", "font-size", "font-weight", "font-style",
		"text-decoration", "text-align", "background-color", "width",
		"height", "margin", "padding", "border", "border-radius",
		"min-width", "max-width", "min-height", "max-height",
		"box-shadow", "text-shadow", "line-height",
	).Globally()

	return p.Sanitize(s)
}
