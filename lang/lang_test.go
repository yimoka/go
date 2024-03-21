package lang

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
)

func TestMatchContent(t *testing.T) {
	langMap := map[string]string{
		"en":      "English",
		"zh":      "Chinese",
		"zh-CN":   "Simplified Chinese",
		"ja":      "Japanese",
		"ko":      "Korean",
		"default": "Default Language",
	}

	// Test case 1: Matching language found
	result, found := MatchContent(langMap, []string{"zh-CN", "en"})
	assert.True(t, found)
	assert.Equal(t, "Simplified Chinese", result)

	result, found = MatchContent(langMap, []string{"zh", "en"})
	assert.True(t, found)
	assert.Equal(t, "Chinese", result)

	// Test case 2: Matching language not found, fallback to default
	result, found = MatchContent(langMap, []string{"fr", "ja"})
	assert.True(t, found)
	assert.Equal(t, "Japanese", result)

	// Test case 3: No matching language found, fallback to default
	result, found = MatchContent(langMap, []string{"fr", "es"})
	assert.False(t, found)
	assert.Equal(t, "Default Language", result)
}

func TestHandleError(t *testing.T) {
	// Test case 1: msg is nil, should return key string
	result := HandleError(MsgKey("test"), nil, nil)
	assert.Equal(t, "test", result)

	// Test case 2: templateData is nil, should return msg.Other
	msg := &i18n.Message{
		Other: "other message",
	}
	result = HandleError(MsgKey("test"), msg, nil)
	assert.Equal(t, "other message", result)

	// Test case 3: templateData is not nil, should return rendered template
	msg = &i18n.Message{
		Other: "Hello, {{.Name}}!",
	}
	templateData := struct {
		Name string
	}{
		Name: "John",
	}
	result = HandleError(MsgKey("test"), msg, templateData)
	assert.Equal(t, "Hello, John!", result)

	// Test case 4: template parsing error, should return key string
	msg = &i18n.Message{
		Other: "Hello, {{.Name}!",
	}
	result = HandleError(MsgKey("test"), msg, templateData)
	assert.Equal(t, "test", result)

	// Test case 5: template execution error, should return key string
	msg = &i18n.Message{
		Other: "Hello, {{.NonExistent}}!",
	}
	result = HandleError(MsgKey("test"), msg, templateData)
	assert.Equal(t, "test", result)
}
