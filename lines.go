package gocovistanbul

import (
	"sort"
)

// LineMapping provides an optimized way of going from byte indices into line/column information.
type LineMapping struct {
	lineStarters []int
}

// BuildLineMapping parses the byte slice for newlines and builds a [LineMapping].
func BuildLineMapping(data []byte) LineMapping {
	lineStarters := make([]int, 0, 128)
	for idx, ch := range data {
		if ch == '\n' {
			lineStarters = append(lineStarters, idx)
		}
	}
	return LineMapping{lineStarters: lineStarters}
}

// Resolve takes a byte position into a file and returns its line/column.
func (l *LineMapping) Resolve(pos int) IstanbulLocation {
	line := sort.Search(len(l.lineStarters), func(i int) bool {
		return pos <= l.lineStarters[i]
	})

	if line == 0 {
		return IstanbulLocation{
			Line:   1,
			Column: pos,
		}
	}

	return IstanbulLocation{
		Line:   line + 1,
		Column: pos - l.lineStarters[line-1] - 1,
	}
}
