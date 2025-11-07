package main

import (
	"fmt"
	"net/url"
	"strings"
)

func getDomainUsingParser(raw string) string {
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "http://" + raw // чтобы parser не ругался
	}

	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}

	host := strings.TrimPrefix(u.Host, "www.")
	return host
}

func main() {
	tests := []string{
		"https://www.google.com/maps?q=test",
		"http://example.org/about",
		"https://my-site.net",
		"www.github.com/user/repo",
	}

	for _, t := range tests {
		fmt.Printf("🔗 %s → 🌐 %s\n", t, getDomainUsingParser(t))
	}
}
