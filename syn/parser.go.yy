// vim: noet sw=8 sts=8
%{
package syn

import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

%}

%union{
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

%type<statements> statements
%type<statement> statement
%type<expr> expr

%token<tok> IDENT
%token<tok> NUMBER
%token<tok> VAR
%token<tok> ECHO
%token<tok> CLASS
%token<tok> FUNCTION

%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%


statements
	:
	{
		$$ = nil
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}
	| statement statements
	{
		$$ = append([]Statement{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}

statement
	: expr ';'
	{
		$$ = &ExpressionStatement{Expr: $1}
	}
	| IDENT '=' expr ';'
	{
		$$ = &VarDefStatement{VarName: $1.lit, Expr: $3}
	}
	| ECHO expr ';'
	{
		$$ = &EchoStatement{Expr: $2}
	}
	| CLASS IDENT '{'  statements '}' ';'
	{
		$$ = &ClassStatement{VarName: $2.lit, Statements: $4}
	}
	| IDENT '=' FUNCTION '(' ')' '{' statements '}' ';'
	{
		$$ = &FunctionStatement{VarName: $1.lit, Statements: $7}
	}



expr	: NUMBER
	{
		$$ = &NumberExpression{Lit: $1.lit}
	}
	| IDENT
	{
		$$ = &IdentifierExpression{Lit: $1.lit}
	}
	| '-' expr      %prec UNARY
	{
		$$ = &UnaryMinusExpression{SubExpr: $2}
	}
	| '(' expr ')'
	{
		$$ = &ParenExpression{SubExpr: $2}
	}
	| expr '+' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('+'), RHS: $3} }
	| expr '-' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('-'), RHS: $3} }
	| expr '*' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('*'), RHS: $3} }
	| expr '/' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('/'), RHS: $3} }
	| expr '%' expr
	{ $$ = &BinOpExpression{LHS: $1, Operator: int('%'), RHS: $3} }

%%

type LexerWrapper struct {
	s          *Scanner
	recentLit  string
	recentPos  Position
	statements []Statement
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []Statement {
	l := LexerWrapper{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.statements
}
