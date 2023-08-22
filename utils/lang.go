package utils

import (
	"os"
	"strings"
)

// handle languages based on https://github.com/tldr-pages/tldr/blob/main/CLIENT-SPECIFICATION.md#language
func GetLanguages() []string {
	lang, exists := os.LookupEnv("LANG")
	if !exists {
		return []string{"en"}
	}
	languages := os.Getenv("LANGUAGE")
	langs := []string{lang}
	for _, locale := range strings.Split(languages, ":") {
		if len(locale) >= 5 && strings.Contains(locale, "_") {
			langs = append(langs, string([]rune(locale)[:5]))
		}
		if len(locale) >= 2 && locale != "POSIX" {
			langs = append(langs, string([]rune(locale)[:2]))
		}
	}
	langs = append(langs, "en")
	return langs
}
