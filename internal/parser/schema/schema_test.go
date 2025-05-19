package schema

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/obj-p/cross/internal/parser/schema/antlr"
)

func TestParseSchema(t *testing.T) {
	input := antlr.NewInputStream(
		`
		data Person(
			name String
			age Int
		)
		`,
	)
	lexer := gen.NewSchemaLexer(input)
	fmt.Println(lexer.GetAllTokens())
}
