package gocovistanbul

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLineMapping(t *testing.T) {
	m := BuildLineMapping([]byte("abc\ndef\nghi"))

	expectedLocations := []IstanbulLocation{
		{Line: 1, Column: 0},
		{Line: 1, Column: 1},
		{Line: 1, Column: 2},
		{Line: 1, Column: 3},
		{Line: 2, Column: 0},
		{Line: 2, Column: 1},
		{Line: 2, Column: 2},
		{Line: 2, Column: 3},
		{Line: 3, Column: 0},
		{Line: 3, Column: 1},
		{Line: 3, Column: 2},
		// Out of bounds:
		{Line: 3, Column: 3},
		{Line: 3, Column: 4},
	}

	for idx, expectedLoc := range expectedLocations {
		actualLoc := m.Resolve(idx)
		require.Equal(t, expectedLoc, actualLoc)
	}
}
