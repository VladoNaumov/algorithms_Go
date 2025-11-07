package main

import (
	"errors"
	"fmt"
	"strings"
)

// ParseCSVLine парсит одну строку CSV в срез полей.
func ParseCSVLine(line string) ([]string, error) {

	if strings.HasSuffix(line, "\r") {
		line = strings.TrimSuffix(line, "\r")
	}

	var (
		fields   []string
		sb       strings.Builder
		inQuotes bool
		i        int
		n        = len(line)
	)

	for i < n {
		ch := line[i]

		if inQuotes {
			if ch == '"' {

				if i+1 < n && line[i+1] == '"' {
					sb.WriteByte('"')
					i += 2
					continue
				}
				inQuotes = false
				i++
				for i < n && (line[i] == ' ' || line[i] == '\t') {
					i++
				}
				continue
			} else {
				sb.WriteByte(ch)
				i++
				continue
			}
		} else {
			if ch == ',' {
				fields = append(fields, sb.String())
				sb.Reset()
				i++
				continue
			}
			if ch == '"' {

				inQuotes = true
				i++
				continue
			}

			sb.WriteByte(ch)
			i++
			continue
		}
	}

	if inQuotes {
		return nil, errors.New("unterminated quoted field")
	}

	fields = append(fields, sb.String())
	return fields, nil
}

func main() {
	tests := []string{
		`a,b,c`,
		`"a,b",c,,"d""e"`,
		`"He said ""Hi""",hello`,
		`one,"two, with comma",three`,
		`,,`,
		`"unterminated`,
		` " spaced " , plain `,
		`"x" , "y"`,
		`"","",`,
		`"a""b","c""","d"`,
		`"a,b", "c, d" ,e`,
	}

	for _, t := range tests {
		fmt.Printf("Input: %q\n", t)
		fields, err := ParseCSVLine(t)
		if err != nil {
			fmt.Printf("  Error: %v\n\n", err)
			continue
		}
		fmt.Printf("  Fields (%d):\n", len(fields))
		for i, f := range fields {
			fmt.Printf("    [%d] %q\n", i, f)
		}
		fmt.Println()
	}
}
