package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXSS(t *testing.T) {
	assert.Equal(t, XSS("<script>alert('XSS')</script>"), "")
	assert.Equal(t, XSS("<p>Hello, World!</p>"), "<p>Hello, World!</p>")
	assert.Equal(t, XSS("<span>Hello, World!</span>"), "<span>Hello, World!</span>")
	// video
	assert.Equal(t, XSS("<video src='video.mp4' controls></video>"), `<video src="video.mp4" controls=""></video>`)
	assert.Equal(t, XSS("<a href='https://example.com'>Link</a>"), `<a href="https://example.com" rel="nofollow">Link</a>`)
	assert.Equal(t, XSS("<img src='image.jpg' alt='Image'>"), "<img src=\"image.jpg\" alt=\"Image\">")
	assert.Equal(t, XSS("<ul><li>Item 1</li><li>Item 2</li></ul>"), "<ul><li>Item 1</li><li>Item 2</li></ul>")
	assert.Equal(t, XSS("<table><tr><td>Cell 1</td><td>Cell 2</td></tr></table>"), "<table><tr><td>Cell 1</td><td>Cell 2</td></tr></table>")
	// 样式
	assert.Equal(t, XSS("<p style='color: red'>Hello, World!</p>"), "<p style=\"color: red\">Hello, World!</p>")
}
