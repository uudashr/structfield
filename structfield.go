package structfield

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `Find struct literals using non-labeled fields.

The structfield analysis reports the usage of struct literal using non-labeled fields more than defined limit (default: 2).
The non-labeled field above the limit considered has higher cognitive load (harder to understand and remember).
`

var Analyzer = &analysis.Analyzer{
	Name:     "structfield",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var (
	limit int // -limit flag
)

func init() {
	Analyzer.Flags.IntVar(&limit, "limit", 2, "Limit of non-labeled fields allowed on struct literal")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		lit := n.(*ast.CompositeLit)
		ok, count := CountNonLabeledFields(lit)
		if !ok {
			return
		}

		if count > limit {
			pass.Reportf(lit.Pos(), "Found %d non-labeled fields on struct literal (> %d)", count, limit)
		}
	})

	return nil, nil
}

// CountNonLabeledFields count usage of non-labeled fields in struct literal,
// returned as n.
// If the lit is not struct literal, then it will return ok as false.
func CountNonLabeledFields(lit *ast.CompositeLit) (ok bool, n int) {
	ident, ok := lit.Type.(*ast.Ident)
	if !ok {
		return false, 0
	}

	if ident.Obj.Kind != ast.Typ {
		return false, 0
	}

	var noLabelCount int
	for _, e := range lit.Elts {
		if _, ok := e.(*ast.KeyValueExpr); !ok {
			noLabelCount++
		}
	}

	return true, noLabelCount
}
