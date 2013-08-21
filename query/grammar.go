//line grammar.y:2
package query

import __yyfmt__ "fmt"

//line grammar.y:3
import (
	"github.com/skydb/sky/core"
)

//line grammar.y:11
type yySymType struct {
	yys              int
	token            int
	integer          int
	str              string
	strs             []string
	query            *Query
	variable         *Variable
	variables        []*Variable
	statement        Statement
	statements       Statements
	selection        *Selection
	selection_field  *SelectionField
	selection_fields []*SelectionField
	condition        *Condition
	condition_within *within
	expr             Expression
	var_ref          *VarRef
	integer_literal  *IntegerLiteral
	boolean_literal  *BooleanLiteral
	string_literal   *StringLiteral
}

const TSTARTQUERY = 57346
const TSTARTSTATEMENT = 57347
const TSTARTSTATEMENTS = 57348
const TSTARTEXPRESSION = 57349
const TFACTOR = 57350
const TSTRING = 57351
const TINTEGER = 57352
const TFLOAT = 57353
const TBOOLEAN = 57354
const TDECLARE = 57355
const TAS = 57356
const TSELECT = 57357
const TGROUP = 57358
const TBY = 57359
const TINTO = 57360
const TWHEN = 57361
const TWITHIN = 57362
const TTHEN = 57363
const TEND = 57364
const TSEMICOLON = 57365
const TCOMMA = 57366
const TLPAREN = 57367
const TRPAREN = 57368
const TRANGE = 57369
const TEQUALS = 57370
const TNOTEQUALS = 57371
const TLT = 57372
const TLTE = 57373
const TGT = 57374
const TGTE = 57375
const TAND = 57376
const TOR = 57377
const TPLUS = 57378
const TMINUS = 57379
const TMUL = 57380
const TDIV = 57381
const TTRUE = 57382
const TFALSE = 57383
const TIDENT = 57384
const TQUOTEDSTRING = 57385
const TWITHINUNITS = 57386
const TINT = 57387

var yyToknames = []string{
	"TSTARTQUERY",
	"TSTARTSTATEMENT",
	"TSTARTSTATEMENTS",
	"TSTARTEXPRESSION",
	"TFACTOR",
	"TSTRING",
	"TINTEGER",
	"TFLOAT",
	"TBOOLEAN",
	"TDECLARE",
	"TAS",
	"TSELECT",
	"TGROUP",
	"TBY",
	"TINTO",
	"TWHEN",
	"TWITHIN",
	"TTHEN",
	"TEND",
	"TSEMICOLON",
	"TCOMMA",
	"TLPAREN",
	"TRPAREN",
	"TRANGE",
	"TEQUALS",
	"TNOTEQUALS",
	"TLT",
	"TLTE",
	"TGT",
	"TGTE",
	"TAND",
	"TOR",
	"TPLUS",
	"TMINUS",
	"TMUL",
	"TDIV",
	"TTRUE",
	"TFALSE",
	"TIDENT",
	"TQUOTEDSTRING",
	"TWITHINUNITS",
	"TINT",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line grammar.y:312
type within struct {
	start int
	end   int
	units string
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 55
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 163

var yyAct = []int{

	8, 14, 93, 74, 30, 96, 82, 95, 25, 94,
	35, 36, 37, 38, 90, 32, 41, 42, 43, 44,
	65, 45, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 19, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 52, 84, 22,
	23, 20, 24, 69, 21, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 41, 42, 43,
	44, 31, 46, 88, 87, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 33, 34, 35,
	36, 37, 38, 39, 86, 41, 42, 43, 44, 33,
	34, 35, 36, 37, 38, 89, 71, 41, 42, 43,
	44, 37, 38, 43, 44, 41, 42, 43, 44, 50,
	49, 81, 72, 12, 73, 12, 70, 13, 48, 13,
	92, 68, 76, 77, 78, 79, 80, 91, 1, 85,
	66, 27, 2, 4, 3, 5, 28, 15, 18, 17,
	16, 9, 51, 11, 75, 67, 83, 47, 29, 10,
	26, 7, 6,
}
var yyPact = []int{

	138, -1000, -1000, -1000, 110, 9, -1000, 128, 110, -1000,
	-1000, -1000, 29, 9, 47, -1000, -1000, -1000, -1000, 9,
	-1000, -1000, -1000, -1000, -1000, 110, -1000, 30, -1000, 104,
	-1000, 94, 27, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, -6, 126, 113, 29, 109,
	80, 103, -42, -20, -20, 79, 79, 31, 31, 71,
	59, 75, 75, -1000, -1000, -1000, 124, 98, -37, -1000,
	6, 125, 68, -1000, 46, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 81, -1000, -28, 123, 108, -43, -33,
	-1000, -35, -1000, -39, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 162, 161, 160, 159, 146, 0, 4, 158, 157,
	156, 155, 154, 153, 152, 1, 150, 149, 148, 147,
	138,
}
var yyR1 = []int{

	0, 20, 20, 20, 20, 1, 6, 6, 5, 5,
	2, 2, 3, 12, 12, 12, 12, 12, 4, 8,
	8, 8, 7, 7, 9, 9, 10, 10, 11, 11,
	13, 14, 14, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	19, 16, 17, 17, 18,
}
var yyR2 = []int{

	0, 2, 2, 2, 2, 2, 0, 2, 1, 1,
	0, 2, 4, 1, 1, 1, 1, 1, 5, 0,
	1, 3, 5, 6, 0, 3, 1, 3, 0, 2,
	6, 0, 5, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -20, 4, 6, 5, 7, -1, -2, -6, -5,
	-4, -13, 15, 19, -15, -19, -16, -17, -18, 25,
	42, 45, 40, 41, 43, -6, -3, 13, -5, -8,
	-7, 42, -15, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, -15, 42, -9, 24, 16,
	25, -14, 20, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, 26, 14, -11, 18, -7,
	17, 26, 42, 21, 45, -12, 8, 9, 10, 11,
	12, 23, 43, -10, 42, 14, 26, -6, 27, 24,
	42, 14, 22, 45, 42, 42, 44,
}
var yyDef = []int{

	0, -2, 10, 6, 0, 0, 1, 6, 2, 3,
	8, 9, 19, 0, 4, 45, 46, 47, 48, 0,
	50, 51, 52, 53, 54, 5, 11, 0, 7, 24,
	20, 0, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 28, 0, 0,
	0, 0, 0, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 49, 0, 0, 0, 21,
	0, 0, 0, 6, 0, 12, 13, 14, 15, 16,
	17, 18, 29, 25, 26, 0, 0, 0, 0, 0,
	22, 0, 30, 0, 27, 23, 32,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
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
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line grammar.y:80
		{
			l := yylex.(*yylexer)
			l.query = yyS[yypt-0].query
		}
	case 2:
		//line grammar.y:85
		{
			l := yylex.(*yylexer)
			l.statements = yyS[yypt-0].statements
		}
	case 3:
		//line grammar.y:90
		{
			l := yylex.(*yylexer)
			l.statement = yyS[yypt-0].statement
		}
	case 4:
		//line grammar.y:95
		{
			l := yylex.(*yylexer)
			l.expression = yyS[yypt-0].expr
		}
	case 5:
		//line grammar.y:103
		{
			yyVAL.query = &Query{}
			yyVAL.query.SetVariables(yyS[yypt-1].variables)
			yyVAL.query.SetStatements(yyS[yypt-0].statements)
		}
	case 6:
		//line grammar.y:112
		{
			yyVAL.statements = make(Statements, 0)
		}
	case 7:
		//line grammar.y:116
		{
			yyVAL.statements = append(yyS[yypt-1].statements, yyS[yypt-0].statement)
		}
	case 8:
		//line grammar.y:123
		{
			yyVAL.statement = Statement(yyS[yypt-0].selection)
		}
	case 9:
		//line grammar.y:127
		{
			yyVAL.statement = Statement(yyS[yypt-0].condition)
		}
	case 10:
		//line grammar.y:134
		{
			yyVAL.variables = make([]*Variable, 0)
		}
	case 11:
		//line grammar.y:138
		{
			yyVAL.variables = append(yyS[yypt-1].variables, yyS[yypt-0].variable)
		}
	case 12:
		//line grammar.y:145
		{
			yyVAL.variable = NewVariable(yyS[yypt-2].str, yyS[yypt-0].str)
		}
	case 13:
		//line grammar.y:151
		{
			yyVAL.str = core.FactorDataType
		}
	case 14:
		//line grammar.y:152
		{
			yyVAL.str = core.StringDataType
		}
	case 15:
		//line grammar.y:153
		{
			yyVAL.str = core.IntegerDataType
		}
	case 16:
		//line grammar.y:154
		{
			yyVAL.str = core.FloatDataType
		}
	case 17:
		//line grammar.y:155
		{
			yyVAL.str = core.BooleanDataType
		}
	case 18:
		//line grammar.y:160
		{
			yyVAL.selection = NewSelection()
			yyVAL.selection.SetFields(yyS[yypt-3].selection_fields)
			yyVAL.selection.Dimensions = yyS[yypt-2].strs
			yyVAL.selection.Name = yyS[yypt-1].str
		}
	case 19:
		//line grammar.y:170
		{
			yyVAL.selection_fields = make([]*SelectionField, 0)
		}
	case 20:
		//line grammar.y:174
		{
			yyVAL.selection_fields = make([]*SelectionField, 0)
			yyVAL.selection_fields = append(yyVAL.selection_fields, yyS[yypt-0].selection_field)
		}
	case 21:
		//line grammar.y:179
		{
			yyVAL.selection_fields = append(yyS[yypt-2].selection_fields, yyS[yypt-0].selection_field)
		}
	case 22:
		//line grammar.y:186
		{
			yyVAL.selection_field = NewSelectionField(yyS[yypt-0].str, yyS[yypt-4].str+"()")
		}
	case 23:
		//line grammar.y:190
		{
			yyVAL.selection_field = NewSelectionField(yyS[yypt-0].str, yyS[yypt-5].str+"("+yyS[yypt-3].str+")")
		}
	case 24:
		//line grammar.y:197
		{
			yyVAL.strs = make([]string, 0)
		}
	case 25:
		//line grammar.y:201
		{
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 26:
		//line grammar.y:208
		{
			yyVAL.strs = make([]string, 0)
			yyVAL.strs = append(yyVAL.strs, yyS[yypt-0].str)
		}
	case 27:
		//line grammar.y:213
		{
			yyVAL.strs = append(yyS[yypt-2].strs, yyS[yypt-0].str)
		}
	case 28:
		//line grammar.y:220
		{
			yyVAL.str = ""
		}
	case 29:
		//line grammar.y:224
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 30:
		//line grammar.y:231
		{
			yyVAL.condition = NewCondition()
			yyVAL.condition.SetExpression(yyS[yypt-4].expr)
			yyVAL.condition.WithinRangeStart = yyS[yypt-3].condition_within.start
			yyVAL.condition.WithinRangeEnd = yyS[yypt-3].condition_within.end
			yyVAL.condition.WithinUnits = yyS[yypt-3].condition_within.units
			yyVAL.condition.SetStatements(yyS[yypt-1].statements)
		}
	case 31:
		//line grammar.y:243
		{
			yyVAL.condition_within = &within{start: 0, end: 0, units: UnitSteps}
		}
	case 32:
		//line grammar.y:247
		{
			yyVAL.condition_within = &within{start: yyS[yypt-3].integer, end: yyS[yypt-1].integer}
			switch yyS[yypt-0].str {
			case "STEPS":
				yyVAL.condition_within.units = UnitSteps
			case "SESSIONS":
				yyVAL.condition_within.units = UnitSessions
			case "SECONDS":
				yyVAL.condition_within.units = UnitSeconds
			}
		}
	case 33:
		//line grammar.y:261
		{
			yyVAL.expr = NewBinaryExpression(OpEquals, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 34:
		//line grammar.y:262
		{
			yyVAL.expr = NewBinaryExpression(OpNotEquals, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 35:
		//line grammar.y:263
		{
			yyVAL.expr = NewBinaryExpression(OpLessThan, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 36:
		//line grammar.y:264
		{
			yyVAL.expr = NewBinaryExpression(OpLessThanOrEqualTo, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 37:
		//line grammar.y:265
		{
			yyVAL.expr = NewBinaryExpression(OpGreaterThan, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 38:
		//line grammar.y:266
		{
			yyVAL.expr = NewBinaryExpression(OpGreaterThanOrEqualTo, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 39:
		//line grammar.y:267
		{
			yyVAL.expr = NewBinaryExpression(OpAnd, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 40:
		//line grammar.y:268
		{
			yyVAL.expr = NewBinaryExpression(OpOr, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 41:
		//line grammar.y:269
		{
			yyVAL.expr = NewBinaryExpression(OpPlus, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 42:
		//line grammar.y:270
		{
			yyVAL.expr = NewBinaryExpression(OpMinus, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 43:
		//line grammar.y:271
		{
			yyVAL.expr = NewBinaryExpression(OpMultiply, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 44:
		//line grammar.y:272
		{
			yyVAL.expr = NewBinaryExpression(OpDivide, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 45:
		//line grammar.y:273
		{
			yyVAL.expr = Expression(yyS[yypt-0].var_ref)
		}
	case 46:
		//line grammar.y:274
		{
			yyVAL.expr = Expression(yyS[yypt-0].integer_literal)
		}
	case 47:
		//line grammar.y:275
		{
			yyVAL.expr = Expression(yyS[yypt-0].boolean_literal)
		}
	case 48:
		//line grammar.y:276
		{
			yyVAL.expr = Expression(yyS[yypt-0].string_literal)
		}
	case 49:
		//line grammar.y:277
		{
			yyVAL.expr = yyS[yypt-1].expr
		}
	case 50:
		//line grammar.y:282
		{
			yyVAL.var_ref = &VarRef{value: yyS[yypt-0].str}
		}
	case 51:
		//line grammar.y:289
		{
			yyVAL.integer_literal = &IntegerLiteral{value: yyS[yypt-0].integer}
		}
	case 52:
		//line grammar.y:296
		{
			yyVAL.boolean_literal = &BooleanLiteral{value: true}
		}
	case 53:
		//line grammar.y:300
		{
			yyVAL.boolean_literal = &BooleanLiteral{value: false}
		}
	case 54:
		//line grammar.y:307
		{
			yyVAL.string_literal = &StringLiteral{value: yyS[yypt-0].str}
		}
	}
	goto yystack /* stack new state and value */
}
