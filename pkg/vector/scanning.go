package vector

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func scanVectors(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	// Skip commas.
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if r != ',' {
			break
		}
	}
	// We're not at a comma so we need to see where the next comma is.
	// Find next comma.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[start:i], nil
		}
	}
	// If we are at EOF then return what we have.
	if atEOF && len(data) > start {
		return len(data), bytes.TrimRightFunc(data[start:], isLineEnding), nil
	}
	// We need more data.
	return start, nil, nil
}

func isLineEnding(r rune) bool {
	lineEnding := &unicode.RangeTable{
		R16: []unicode.Range16{
			{Lo: 0x000A, Hi: 0x000D, Stride: 1},
			{Lo: 0x0085, Hi: 0x0085, Stride: 1},
			{Lo: 0x2028, Hi: 0x2029, Stride: 1},
		},
	}
	return unicode.Is(lineEnding, r)
}