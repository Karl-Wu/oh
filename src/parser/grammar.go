//line grammar.y:16
package parser

import __yyfmt__ "fmt"

//line grammar.y:16
import (
	. "github.com/michaelmacinnis/oh/src/cell"
	"github.com/michaelmacinnis/oh/src/task"
)

type yySymType struct {
	yys int
	c   Cell
	s   string
}

const CTRLC = 57346
const DOUBLE_QUOTED = 57347
const SINGLE_QUOTED = 57348
const SYMBOL = 57349
const BACKGROUND = 57350
const ORF = 57351
const ANDF = 57352
const PIPE = 57353
const REDIRECT = 57354
const SUBSTITUTE = 57355
const CONS = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CTRLC",
	"DOUBLE_QUOTED",
	"SINGLE_QUOTED",
	"SYMBOL",
	"BACKGROUND",
	"ORF",
	"ANDF",
	"PIPE",
	"REDIRECT",
	"SUBSTITUTE",
	"^",
	"@",
	"`",
	"CONS",
	"\n",
	")",
	";",
	":",
	"{",
	"}",
	"%",
	"(",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line grammar.y:197

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	5, 16,
	6, 16,
	7, 16,
	15, 16,
	16, 16,
	18, 5,
	21, 16,
	22, 16,
	24, 16,
	25, 16,
	-2, 0,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	8, 14,
	9, 14,
	10, 14,
	11, 14,
	12, 14,
	13, 14,
	18, 14,
	19, 14,
	-2, 17,
	-1, 10,
	1, 1,
	5, 16,
	6, 16,
	7, 16,
	15, 16,
	16, 16,
	18, 5,
	21, 16,
	22, 16,
	24, 16,
	25, 16,
	-2, 0,
	-1, 47,
	18, 34,
	-2, 16,
	-1, 67,
	18, 34,
	-2, 16,
}

const yyNprod = 49
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 115

var yyAct = [...]int{

	5, 8, 18, 61, 7, 68, 21, 11, 12, 13,
	14, 15, 16, 33, 34, 35, 9, 37, 64, 70,
	39, 9, 36, 40, 59, 45, 42, 9, 51, 52,
	9, 48, 49, 17, 29, 30, 31, 4, 20, 67,
	47, 10, 44, 55, 24, 25, 15, 16, 62, 58,
	56, 57, 63, 26, 27, 9, 65, 43, 41, 66,
	44, 28, 29, 30, 31, 14, 15, 16, 62, 69,
	50, 71, 24, 25, 29, 30, 31, 17, 22, 23,
	60, 26, 27, 46, 24, 25, 13, 14, 15, 16,
	22, 23, 3, 26, 27, 11, 12, 13, 14, 15,
	16, 19, 54, 32, 38, 6, 53, 11, 12, 13,
	14, 15, 16, 2, 1,
}
var yyPact = [...]int{

	35, -1000, 23, -1000, -1000, 99, -1000, 13, 69, -1000,
	35, -1000, 7, 7, 7, 29, 7, -1000, 7, 69,
	-1000, 43, 69, 22, 29, 29, 63, 10, -1000, -1000,
	-1000, -1000, -1000, 76, 54, 34, 43, 87, -1000, -1000,
	57, -1000, 43, 29, 29, -1000, 69, 1, 25, 25,
	45, -1, -1000, -1000, 7, -1000, 25, -1000, -1000, -1000,
	21, -1000, 99, -19, -1000, -1000, 57, -4, -1000, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 114, 113, 92, 0, 6, 105, 4, 1, 2,
	104, 102, 101, 38, 83, 80, 3, 61,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 6, 6, 8, 8, 7, 7,
	10, 10, 11, 11, 9, 9, 9, 13, 13, 13,
	14, 14, 15, 15, 16, 16, 12, 12, 5, 5,
	5, 5, 5, 5, 5, 5, 17, 17, 17,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 0, 1, 2, 3, 3,
	3, 3, 4, 1, 1, 3, 0, 1, 1, 2,
	1, 3, 1, 3, 1, 2, 1, 2, 3, 2,
	2, 4, 1, 3, 0, 1, 1, 2, 3, 2,
	2, 3, 4, 3, 2, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 2, -4, -6, -7, -8, 20,
	18, 8, 9, 10, 11, 12, 13, 20, -9, -12,
	-13, -5, 21, 22, 15, 16, 24, 25, -17, 5,
	6, 7, -3, -4, -4, -4, -5, -4, -10, -8,
	-7, -13, -5, 14, 17, -9, -14, 18, -5, -5,
	7, -4, 19, 19, -11, -9, -5, -5, -9, 23,
	-15, -16, -4, 7, 19, -8, -7, 18, 24, -9,
	23, -16,
}
var yyDef = [...]int{

	-2, -2, 0, 2, 4, 6, 13, -2, 0, 18,
	-2, 7, 16, 16, 16, 0, 16, 19, 16, 24,
	26, 36, 0, 0, 0, 0, 0, 16, 45, 46,
	47, 48, 3, 8, 9, 10, 11, 0, 15, 20,
	17, 25, 37, 0, 0, 27, 29, -2, 39, 40,
	0, 0, 44, 12, 16, 22, 38, 41, 28, 30,
	0, 32, 35, 0, 43, 21, 17, -2, 42, 23,
	31, 33,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	18, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 24, 3, 3,
	25, 19, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 21, 20,
	3, 3, 3, 3, 15, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 14, 3, 16, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 22, 3, 23,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 17,
}
var yyTok3 = [...]int{
	0,
}

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
	state     func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
		state:     func() int { return -1 },
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
	yyS := make([]yySymType, yyMaxDepth)

	startyyVAL := yyVAL
start:
	yyVAL = startyyVAL

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.state = func() int { return yystate }
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
		if yychar == CTRLC {
			goto start
		}

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
			if yychar == CTRLC {
				goto start
			}

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

	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:40
		{
			yyVAL.c = Null
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:42
		{
			yyVAL.c = yyDollar[1].c
			if yyDollar[1].c != Null {
				yylex.(*scanner).process(yyDollar[1].c)
			}
			goto start
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:50
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[1].c)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:54
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[1].c, yyDollar[3].c)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:58
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[1].c, yyDollar[3].c)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:62
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[1].c, yyDollar[3].c)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:66
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[3].c, yyDollar[1].c)
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:70
		{
			yyVAL.c = List(NewSymbol(yyDollar[2].s), yyDollar[3].c, yyDollar[1].c)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:74
		{
			yyVAL.c = yyDollar[1].c
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:76
		{
			yyVAL.c = Null
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:78
		{
			if yyDollar[3].c == Null {
				yyVAL.c = yyDollar[2].c
			} else {
				yyVAL.c = Cons(NewSymbol("block"), Cons(yyDollar[2].c, yyDollar[3].c))
			}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:94
		{
			yyVAL.c = Null
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:96
		{
			yyVAL.c = yyDollar[2].c
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:98
		{
			yyVAL.c = Cons(yyDollar[1].c, Null)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:100
		{
			yyVAL.c = AppendTo(yyDollar[1].c, yyDollar[3].c)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:102
		{
			yyVAL.c = yyDollar[1].c
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:104
		{
			yyVAL.c = JoinTo(yyDollar[1].c, yyDollar[2].c)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:108
		{
			yyVAL.c = yyDollar[1].c
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:110
		{
			yyVAL.c = Cons(yyDollar[2].c, Null)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:112
		{
			if yyDollar[2].c == Null {
				yyVAL.c = yyDollar[3].c
			} else {
				yyVAL.c = JoinTo(yyDollar[2].c, yyDollar[3].c)
			}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:120
		{
			yyVAL.c = yyDollar[2].c
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:124
		{
			yyVAL.c = Null
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:126
		{
			yyVAL.c = yyDollar[2].c
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:128
		{
			if yyDollar[1].c == Null {
				yyVAL.c = yyDollar[1].c
			} else {
				yyVAL.c = Cons(yyDollar[1].c, Null)
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:136
		{
			if yyDollar[1].c == Null {
				if yyDollar[3].c == Null {
					yyVAL.c = yyDollar[3].c
				} else {
					yyVAL.c = Cons(yyDollar[3].c, Null)
				}
			} else {
				if yyDollar[3].c == Null {
					yyVAL.c = yyDollar[1].c
				} else {
					yyVAL.c = AppendTo(yyDollar[1].c, yyDollar[3].c)
				}
			}
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line grammar.y:152
		{
			yyVAL.c = Null
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:154
		{
			yyVAL.c = yyDollar[1].c
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:156
		{
			yyVAL.c = Cons(yyDollar[1].c, Null)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:158
		{
			yyVAL.c = AppendTo(yyDollar[1].c, yyDollar[2].c)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:160
		{
			s := Cons(task.NewString(yylex.(*scanner).task, ""), NewSymbol("join"))
			yyVAL.c = List(s, yyDollar[1].c, yyDollar[3].c)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:165
		{
			yyVAL.c = List(NewSymbol("splice"), yyDollar[2].c)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:169
		{
			yyVAL.c = List(NewSymbol("backtick"), yyDollar[2].c)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:173
		{
			yyVAL.c = Cons(yyDollar[1].c, yyDollar[3].c)
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line grammar.y:177
		{
			yyVAL.c = yylex.(*scanner).deref(yyDollar[2].s, yyDollar[3].s)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line grammar.y:181
		{
			yyVAL = yyDollar[2]
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line grammar.y:183
		{
			yyVAL.c = Null
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:185
		{
			yyVAL = yyDollar[1]
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:187
		{
			yyVAL.c = task.NewString(yylex.(*scanner).task, yyDollar[1].s[1:len(yyDollar[1].s)-1])
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:191
		{
			yyVAL.c = task.NewRawString(yylex.(*scanner).task, yyDollar[1].s[1:len(yyDollar[1].s)-1])
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line grammar.y:195
		{
			yyVAL.c = NewSymbol(yyDollar[1].s)
		}
	}
	goto yystack /* stack new state and value */
}
