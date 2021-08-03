package main

import (
	"github.com/uudashr/structfield"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(structfield.Analyzer)
}
