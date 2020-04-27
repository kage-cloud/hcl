package fuzzexpr

import (
	"github.com/kage-cloud/hcl/v2"
	"github.com/kage-cloud/hcl/v2/hclsyntax"
)

func Fuzz(data []byte) int {
	_, diags := hclsyntax.ParseExpression(data, "<fuzz-expr>", hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		return 0
	}

	return 1
}
