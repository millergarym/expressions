package expressions

import (
	"github.com/millergarym/expressions/parser"
	"github.com/wxio/antlr4-go"
)

type ExpressionsVisitor struct {
	*antlr.BaseParseTreeVisitor
	Expressions map[string]parser.IExprContext
}

var _ parser.CodelineContextVisitor = &ExpressionsVisitor{}

func NewExpressionsVisitor() *ExpressionsVisitor {
	visitor := new(ExpressionsVisitor)
	visitor.Expressions = make(map[string]parser.IExprContext)
	return visitor
}

func (ev *ExpressionsVisitor) VisitCodeline(ctx parser.ICodelineContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	//func (ev *ExpressionsVisitor) VisitCodeline(ctx *parser.CodelineContext) interface{} {
	ev.Expressions[ctx.(*parser.CodelineContext).GetLabel().GetText()] = ctx.(*parser.CodelineContext).GetCode()
	return nil
}

func GetExpressions(filename string) map[string]parser.IExprContext {
	input := antlr.NewFileStream(filename)
	lexer := parser.NewExpressionsLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionsParser(stream)
	tree := p.Start()

	tv := NewExpressionsVisitor()
	tree.Visit(tv)

	return tv.Expressions
}
