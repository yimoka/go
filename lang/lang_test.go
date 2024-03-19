package lang

import (
	"testing"

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
