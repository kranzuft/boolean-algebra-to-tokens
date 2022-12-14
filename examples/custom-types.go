package main

import (
	"github.com/kranzuft/boolean-algebra-to-tokens/cmd/com/nodlim/batt/types"
)

func main() {
	def := types.TokensDefinition{}
	def.DefineTokenInfo(types.AND, "and", "and").
		DefineTokenInfo(types.OR, "or", "or").
		DefineTokenInfo(types.NOT, "not", "not").
		DefineTokenInfo(types.ANDNOT, "and not", "and not").
		DefineTokenInfo(types.ORNOT, "or not", "or not").
		DefineTokenInfo(types.TRUE, "True", "true").
		DefineTokenInfo(types.LBR, "{", "left bracket").
		DefineTokenInfo(types.RBR, "}", "right bracket").
		DefineTokenInfo(types.EOL, "\n", "end of line").
		DefineTokenInfo(types.EXP, "", "expression").
		DefineTokenInfo(types.DQUOTE, "\"", "double inverted comma").
		DefineTokenInfo(types.SQUOTE, "'", "single inverted comma").
		Finalise()
}
