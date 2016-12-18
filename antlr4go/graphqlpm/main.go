package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

type myListener struct {
	*BaseGraphQLPMListener
}

func newMyListener() *myListener {
	return new(myListener)
}

func (this *myListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := NewGraphQLPMLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewGraphQLPMParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	antlr.ParseTreeWalkerDefault.Walk(newMyListener(), tree)
}
