package gocovistanbul

// https://github.com/istanbuljs/istanbuljs/blob/main/docs/raw-output.md

type IstanbulCoverage map[string]IstanbulFile

func NewIstanbulCoverage() IstanbulCoverage {
	return make(IstanbulCoverage)
}

type IstanbulFile struct {
	Path string `json:"path"`

	StatementMap      map[string]IstanbulSpan `json:"statementMap"`
	StatementCounters map[string]int          `json:"s"`

	FunctionMap      map[string]IstanbulFunction `json:"fnMap"`
	FunctionCounters map[string]int              `json:"f"`

	BranchMap      map[string]any `json:"branchMap"`
	BranchCounters map[string]int `json:"b"`
}

func NewIstanbulFile(path string) IstanbulFile {
	return IstanbulFile{
		Path:              path,
		StatementMap:      make(map[string]IstanbulSpan),
		StatementCounters: make(map[string]int),

		FunctionMap:      make(map[string]IstanbulFunction),
		FunctionCounters: make(map[string]int),

		BranchMap:      make(map[string]any),
		BranchCounters: make(map[string]int),
	}
}

type IstanbulSpan struct {
	Start IstanbulLocation `json:"start"`
	End   IstanbulLocation `json:"end"`
}

type IstanbulLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type IstanbulFunction struct {
	Name string       `json:"name"`
	Line int          `json:"line"`
	Decl IstanbulSpan `json:"decl"`
	Body IstanbulSpan `json:"body"`
}
