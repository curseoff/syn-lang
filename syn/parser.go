//line syn/parser.go.yy:3
package syn

import __yyfmt__ "fmt"

//line syn/parser.go.yy:3
import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

//line syn/parser.go.yy:17
type yySymType struct {
	yys        int
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

const IDENT = 57346
const NUMBER = 57347
const VAR = 57348
const ECHO = 57349
const CLASS = 57350
const FUNCTION = 57351
const UNARY = 57352

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"VAR",
	"ECHO",
	"CLASS",
	"FUNCTION",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"';'",
	"'='",
	"'{'",
	"'}'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line syn/parser.go.yy:109

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

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 17
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 74

var yyAct = [...]int{

	1, 36, 34, 10, 12, 13, 14, 15, 16, 41,
	3, 37, 20, 38, 31, 32, 18, 17, 42, 21,
	22, 39, 2, 23, 24, 25, 26, 27, 28, 0,
	4, 7, 35, 5, 6, 19, 7, 8, 0, 40,
	29, 0, 8, 19, 7, 0, 9, 14, 15, 16,
	8, 9, 12, 13, 14, 15, 16, 0, 33, 9,
	12, 13, 14, 15, 16, 0, 30, 12, 13, 14,
	15, 16, 0, 11,
}
var yyPact = [...]int{

	26, -1000, 26, 57, 0, 39, 8, -1000, 39, 39,
	-1000, -1000, 39, 39, 39, 39, 39, 31, 50, -1000,
	-4, -1000, -6, 35, 35, -1000, -1000, -1000, 42, -18,
	-1000, 26, -1000, -1000, -20, -8, -5, 5, 26, -1000,
	-10, 2, -1000,
}
var yyPgo = [...]int{

	0, 0, 22, 10,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3,
}
var yyR2 = [...]int{

	0, 0, 2, 2, 4, 3, 6, 9, 1, 1,
	2, 3, 3, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 4, 7, 8, 5, 11, 20,
	-1, 16, 10, 11, 12, 13, 14, 17, -3, 4,
	4, -3, -3, -3, -3, -3, -3, -3, -3, 9,
	16, 18, 21, 16, 20, -1, 21, 19, 18, 16,
	-1, 19, 16,
}
var yyDef = [...]int{

	1, -2, 1, 0, 9, 0, 0, 8, 0, 0,
	2, 3, 0, 0, 0, 0, 0, 0, 0, 9,
	0, 10, 0, 12, 13, 14, 15, 16, 0, 0,
	5, 1, 11, 4, 0, 0, 0, 0, 1, 6,
	0, 0, 7,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 14, 3, 3,
	20, 21, 12, 10, 3, 11, 3, 13, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 16,
	3, 17, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 3, 19,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 15,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line syn/parser.go.yy:44
		{
			yyVAL.statements = nil
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line syn/parser.go.yy:51
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line syn/parser.go.yy:60
		{
			yyVAL.statement = &ExpressionStatement{Expr: yyDollar[1].expr}
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line syn/parser.go.yy:64
		{
			yyVAL.statement = &VarDefStatement{VarName: yyDollar[1].tok.lit, Expr: yyDollar[3].expr}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:68
		{
			yyVAL.statement = &EchoStatement{Expr: yyDollar[2].expr}
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line syn/parser.go.yy:72
		{
			yyVAL.statement = &ClassStatement{VarName: yyDollar[2].tok.lit, Statements: yyDollar[4].statements}
		}
	case 7:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line syn/parser.go.yy:76
		{
			yyVAL.statement = &FunctionStatement{VarName: yyDollar[1].tok.lit, Statements: yyDollar[7].statements}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line syn/parser.go.yy:83
		{
			yyVAL.expr = &NumberExpression{Lit: yyDollar[1].tok.lit}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line syn/parser.go.yy:87
		{
			yyVAL.expr = &IdentifierExpression{Lit: yyDollar[1].tok.lit}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line syn/parser.go.yy:91
		{
			yyVAL.expr = &UnaryMinusExpression{SubExpr: yyDollar[2].expr}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:95
		{
			yyVAL.expr = &ParenExpression{SubExpr: yyDollar[2].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:99
		{
			yyVAL.expr = &BinOpExpression{LHS: yyDollar[1].expr, Operator: int('+'), RHS: yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:101
		{
			yyVAL.expr = &BinOpExpression{LHS: yyDollar[1].expr, Operator: int('-'), RHS: yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:103
		{
			yyVAL.expr = &BinOpExpression{LHS: yyDollar[1].expr, Operator: int('*'), RHS: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:105
		{
			yyVAL.expr = &BinOpExpression{LHS: yyDollar[1].expr, Operator: int('/'), RHS: yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line syn/parser.go.yy:107
		{
			yyVAL.expr = &BinOpExpression{LHS: yyDollar[1].expr, Operator: int('%'), RHS: yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
