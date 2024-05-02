package gocovistanbul

import (
	"fmt"
	"os"

	"github.com/axw/gocov"
)

// Converter converts from gocov's JSON format into Istanbul's JSON format.
type Converter struct {
	lineMappers map[string]LineMapping
	output      IstanbulCoverage
}

// NewConverter initializes a new converter.
func NewConverter() *Converter {
	return &Converter{
		output:      NewIstanbulCoverage(),
		lineMappers: make(map[string]LineMapping),
	}
}

// Finalize returns the final Istanbul coverage (which can be serialized).
func (c *Converter) Finalize() IstanbulCoverage {
	return c.output
}

func (c *Converter) getLineMapper(path string) (LineMapping, error) {
	m, ok := c.lineMappers[path]
	if !ok {
		data, err := os.ReadFile(path)
		if err != nil {
			return LineMapping{}, err
		}
		m = BuildLineMapping(data)
		c.lineMappers[path] = m
	}
	return m, nil
}

// AddPackage adds a gocov Package to the coverage.
func (c *Converter) AddPackage(pkg *gocov.Package) error {
	for _, fn := range pkg.Functions {
		lines, err := c.getLineMapper(fn.File)
		if err != nil {
			return err
		}

		file, ok := c.output[fn.File]
		if !ok {
			file = NewIstanbulFile(fn.File)
		}

		fnHit := 0

		for _, stmt := range fn.Statements {
			stmtID := fmt.Sprintf("%d", len(file.StatementMap))
			if stmt.Reached > 0 {
				fnHit = 1
			}
			file.StatementCounters[stmtID] = int(stmt.Reached)
			file.StatementMap[stmtID] = IstanbulSpan{
				Start: lines.Resolve(stmt.Start),
				End:   lines.Resolve(stmt.End),
			}
		}

		fnID := fmt.Sprintf("%d", len(file.FunctionMap))
		file.FunctionCounters[fnID] = fnHit

		body := IstanbulSpan{
			Start: lines.Resolve(fn.Start),
			End:   lines.Resolve(fn.End),
		}

		file.FunctionMap[fnID] = IstanbulFunction{
			Name: fn.Name,
			Line: body.Start.Line,
			Decl: body,
			Body: body,
		}
		c.output[fn.File] = file
	}

	return nil
}
