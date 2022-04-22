package main

import (
	"errors"
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "buildssactrlflow",
	Doc:  `Experimental`,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	ssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		return nil, errors.New("buildssa: buildssa does not provide data")
	}
	for _, f := range ssa.SrcFuncs {
		for _, block := range f.Blocks {
			fmt.Println(block.Index, " -> ", block.Succs)
			for _, instr := range block.Instrs {
				fmt.Println(instr.String(), pass.Fset.Position(instr.Pos()))
			}
			fmt.Println("")
		}
	}
	return nil, nil
}
