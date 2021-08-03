package structfield_test

import (
	"testing"

	"github.com/uudashr/structfield"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	structfield.Analyzer.Flags.Set("limit", "0")
	analysistest.Run(t, testdata, structfield.Analyzer, "a")
}

func TestAnalyzerLimit2(t *testing.T) {
	testdata := analysistest.TestData()
	structfield.Analyzer.Flags.Set("limit", "2")
	analysistest.Run(t, testdata, structfield.Analyzer, "b")
}

func TestAnalyzerDefaultLimit(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, structfield.Analyzer, "b")
}
