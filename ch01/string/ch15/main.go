package main

import (
	"fmt"
	"strings"
	"unicode"
)

// EncodeSpacesSimple заменяет каждый ASCII-пробел ' ' на "%20".
func EncodeSpacesSimple(s string) string {

	if !strings.ContainsRune(s, ' ') {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if r == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// EncodeSpacesWhitespace заменяет все Unicode-пробельные символы
// (unicode.IsSpace) на "%20". Полезно если строки содержат \t, \n, NBSP и т.д.
func EncodeSpacesWhitespace(s string) string {

	hasSpace := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			hasSpace = true
			break
		}
	}
	if !hasSpace {
		return s
	}

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if unicode.IsSpace(r) {
			b.WriteString("%20")
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func main() {
	tests := []string{
		"Hello World",
		"a b  c",
		"no_spaces_here",
		"line 1\nline 2",
		"tab\tseparated",
		"NBSP\u00A0here",
		"%20 already encoded",
	}

	fmt.Println("=== EncodeSpacesSimple (только ' ') ===")
	for _, t := range tests {
		fmt.Printf("%q -> %q\n", t, EncodeSpacesSimple(t))
	}

	fmt.Println("\n=== EncodeSpacesWhitespace (все unicode-пробелы) ===")
	for _, t := range tests {
		fmt.Printf("%q -> %q\n", t, EncodeSpacesWhitespace(t))
	}

	fmt.Println("\nПримечание: для полного URL-encoding (не только пробелы) используйте net/url:")
	fmt.Println(`  import "net/url"  -> url.QueryEscape(s)  или  url.PathEscape(s)`)
}
