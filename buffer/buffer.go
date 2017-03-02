// Package buffer provides three simple things for a buffer. (1) an array of strings which make up the text, (2) a
// cursor position, and (3) a highlight.
package buffer

// Position is a position in the buffer where the cursor currently lives.
type Position struct {
	LineNo int
	CharNo int
}

// Range is a start and end in the buffer using two positions as values. It could be used for selections, highlights,
// or anything else. The Start and End positions are generally classed as inclusive (Note: this may change if it turns
// out to be better to include Start but exclude End).
type Range struct {
	Start Position
	End   Position
}

// Buffer defines the buffer we actually work on when editing. All methods on Buffer are essentially known as
// 'built-in' commands within the editor. All other commands that make up the editors functions end up calling these
// primitives.
type Buffer struct {
	Lines      []string
	Selection  Range
	Highlights []Range
}

func NewBuffer() *Buffer {
	b := Buffer{}
	b.Lines = make([]string, 0)
	b.Highlights = make([]Range, 0)
	return &b
}
