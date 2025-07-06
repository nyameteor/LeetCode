package main

import (
	"fmt"
	"strings"
)

func entityParser(text string) string {
	entityMap := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}

	var res, buffer strings.Builder

	const (
		Text = iota
		Entity
	)
	state := Text

	for i := 0; i < len(text); i++ {
		ch := text[i]

		switch state {
		case Text:
			if ch == '&' {
				buffer.WriteByte(ch)
				state = Entity
			} else {
				res.WriteByte(ch)
			}
		case Entity:
			switch ch {
			case '&':
				res.WriteString(buffer.String())
				buffer.Reset()
				buffer.WriteByte(ch)
			case ';':
				buffer.WriteByte(ch)
				entity := buffer.String()
				if val, ok := entityMap[entity]; ok {
					res.WriteString(val)
				} else {
					res.WriteString(entity)
				}
				buffer.Reset()
				state = Text
			default:
				buffer.WriteByte(ch)
			}
		}
	}

	// Flush unprocessed buffer
	if state == Entity {
		entity := buffer.String()
		if val, ok := entityMap[entity]; ok {
			res.WriteString(val)
		} else {
			res.WriteString(entity)
		}
		buffer.Reset()
	}

	return res.String()
}

func main() {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "&amp; is an HTML entity but &ambassador; is not.",
			want:  "& is an HTML entity but &ambassador; is not.",
		},
		{
			input: "and I quote: &quot;...&quot;",
			want:  "and I quote: \"...\"",
		},
		{
			input: "&&gt;",
			want:  "&>",
		},
		{
			input: "&amp;gt;",
			want:  "&gt;",
		},
		{
			input: "&a&gt;",
			want:  "&a>",
		},
	}

	for _, tc := range testCases {
		got := entityParser(tc.input)
		if got != tc.want {
			fmt.Printf("entityParser(%q) = %q, want = %q\n", tc.input, got, tc.want)
		}
	}
}
