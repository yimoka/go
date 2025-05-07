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
		"color", "opacity",
		"font-family", "font-size", "font-weight", "font-style",
		"text-decoration", "text-align", "text-shadow", "text-indent", "text-transform", "text-overflow", "text-wrap", "text-justify", "text-line-clamp",
		"background-color", "background-image", "background-size", "background-position", "background-repeat", "background-attachment", "background-blend-mode",
		"border", "border-radius", "border-color", "border-width", "border-style", "border-top", "border-top-color", "border-top-width", "border-top-style",
		"min-width", "max-width", "min-height", "max-height", "width", "height",
		"margin", "margin-top", "margin-right", "margin-bottom", "margin-left",
		"padding", "padding-top", "padding-right", "padding-bottom", "padding-left",
		"box-shadow", "box-sizing", "box-align", "box-direction", "box-flex", "box-flex-group", "box-lines", "box-ordinal-group", "box-orient", "box-pack", "box-reflect",
		"line-height", "letter-spacing",
		"list-style", "list-style-type", "list-style-position", "list-style-image",
		"outline", "outline-color", "outline-width", "outline-style", "outline-offset",
		"vertical-align", "white-space", "word-break", "word-spacing", "word-wrap",
		"z-index",
	).Globally()

	return p.Sanitize(s)
}
