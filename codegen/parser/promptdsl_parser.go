// Code generated from ./codegen/grammar/PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromptDSLParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PromptDSLParser struct {
	*antlr.BaseParser
}

var PromptDSLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promptdslparserParserInit() {
	staticData := &PromptDSLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'string'", "'float'", "'int'", "'prompt'", "'params'", "'sys'",
		"'user'", "'note'", "'in'", "'output'", "'format'", "'type'", "'struct'",
		"'schema'", "'parse'", "'jsonfix'", "'markdown'", "'if'", "'else'",
		"'outputspec'", "'for'", "'range'", "'switch'", "'default'", "'case'",
		"'goimport'", "", "", "", "", "'{'", "", "'('", "')'", "':'", "'='",
		"','", "'.'", "'=='", "'!='", "'@'", "'md'", "'json'", "'['", "']'",
		"'++'", "'--'", "'-'", "'*'", "'/'", "'%'", "'_'", "'+='", "'-='", "'*='",
		"'/='", "'%='", "':='", "'<'", "'<='", "'>'", "'>='", "", "", "", "",
		"'|'", "';'", "'+'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING_TYPE", "FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM",
		"USER", "NOTE", "INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "SCHEMA",
		"PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC", "FOR", "RANGE",
		"SWITCH", "DEFAULT", "CASE", "GOIMPORT", "BEFORE", "FIX", "AFTER", "ARRAY_OUTPUTSPEC",
		"LBRACE", "RBRACE", "LPAREN", "RPAREN", "COLON", "EQUAL", "COMMA", "DOT",
		"EQEQ", "NOTEQ", "AT", "MD", "JSON", "LBRACK", "RBRACK", "INCREMENT",
		"DECREMENT", "MINUS", "STAR", "SLASH", "MOD", "UNDERSCORE", "PLUSEQ",
		"MINUSEQ", "MULTEQ", "DIVEQ", "MODEQ", "DECL_ASSIGN", "LT", "LTE", "GT",
		"GTE", "ID", "STRING", "NUMBER", "BOOL", "PIPE", "SEMI", "PLUS", "DASH_STRING",
		"WS", "LINE_COMMENT", "BLOCK_COMMENT", "CODE_STRING_SINGLE", "CODE_STRING",
		"CODE_TEXT",
	}
	staticData.RuleNames = []string{
		"promptFile", "promptDef", "promptBlock", "goimportSection", "goimportEntry",
		"inputSection", "outputSection", "outputStruct", "outputMarkdown", "beforeSection",
		"systemSection", "sysContent", "userSection", "userContent", "moduleDef",
		"moduleContent", "thencontent", "elsecontent", "forcontent", "ifStatement",
		"condition", "forStatement", "assignExpr", "updateExpr", "switchStatement",
		"switchCase", "switchDefault", "typeCase", "typeDefault", "typeName",
		"dslCallExpr", "expr", "fieldDef", "textLine", "paramPath", "structDef",
		"annotation", "annotationArgs", "annotationValue", "arrayLiteral", "defaultAnnotation",
		"fixSection", "afterSection", "codeBlockContent", "type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 76, 571, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		4, 0, 96, 8, 0, 11, 0, 12, 0, 97, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4,
		1, 106, 8, 1, 11, 1, 12, 1, 107, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 121, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 126,
		8, 3, 10, 3, 12, 3, 129, 9, 3, 1, 3, 1, 3, 1, 4, 3, 4, 134, 8, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 4, 5, 141, 8, 5, 11, 5, 12, 5, 142, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 149, 8, 5, 1, 6, 5, 6, 152, 8, 6, 10, 6, 12, 6, 155,
		9, 6, 1, 6, 1, 6, 1, 6, 3, 6, 160, 8, 6, 1, 7, 1, 7, 4, 7, 164, 8, 7, 11,
		7, 12, 7, 165, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 4, 10, 180, 8, 10, 11, 10, 12, 10, 181, 1, 10, 1, 10,
		1, 10, 1, 10, 4, 10, 188, 8, 10, 11, 10, 12, 10, 189, 1, 10, 1, 10, 3,
		10, 194, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		203, 8, 11, 1, 12, 1, 12, 1, 12, 4, 12, 208, 8, 12, 11, 12, 12, 12, 209,
		1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 216, 8, 12, 11, 12, 12, 12, 217, 1,
		12, 1, 12, 3, 12, 222, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 231, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 236, 8, 14, 10, 14,
		12, 14, 239, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 3, 15, 251, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 265, 8, 19, 10,
		19, 12, 19, 268, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 274, 8, 19,
		10, 19, 12, 19, 277, 9, 19, 1, 19, 3, 19, 280, 8, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 287, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 5, 21, 297, 8, 21, 10, 21, 12, 21, 300, 9, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		5, 21, 313, 8, 21, 10, 21, 12, 21, 316, 9, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 327, 8, 21, 10, 21, 12, 21,
		330, 9, 21, 1, 21, 1, 21, 3, 21, 334, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 353, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24,
		359, 8, 24, 10, 24, 12, 24, 362, 9, 24, 1, 24, 3, 24, 365, 8, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 373, 8, 25, 10, 25, 12, 25, 376,
		9, 25, 1, 26, 1, 26, 1, 26, 5, 26, 381, 8, 26, 10, 26, 12, 26, 384, 9,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 390, 8, 27, 10, 27, 12, 27, 393,
		9, 27, 1, 28, 1, 28, 1, 28, 5, 28, 398, 8, 28, 10, 28, 12, 28, 401, 9,
		28, 1, 29, 3, 29, 404, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 5, 30, 413, 8, 30, 10, 30, 12, 30, 416, 9, 30, 3, 30, 418, 8, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 3, 31, 432, 8, 31, 1, 31, 1, 31, 1, 31, 5, 31, 437, 8, 31, 10,
		31, 12, 31, 440, 9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 447,
		8, 32, 1, 32, 5, 32, 450, 8, 32, 10, 32, 12, 32, 453, 9, 32, 1, 32, 3,
		32, 456, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 462, 8, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 471, 8, 34, 10, 34, 12, 34,
		474, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 4, 35, 480, 8, 35, 11, 35, 12,
		35, 481, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 490, 8, 36, 1,
		36, 3, 36, 493, 8, 36, 1, 37, 1, 37, 1, 37, 5, 37, 498, 8, 37, 10, 37,
		12, 37, 501, 9, 37, 1, 38, 1, 38, 3, 38, 505, 8, 38, 1, 39, 1, 39, 1, 39,
		1, 39, 5, 39, 511, 8, 39, 10, 39, 12, 39, 514, 9, 39, 3, 39, 516, 8, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 524, 8, 40, 1, 40, 3,
		40, 527, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 543, 8, 43, 10, 43, 12,
		43, 546, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 5, 44, 552, 8, 44, 10, 44,
		12, 44, 555, 9, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 3, 44, 565, 8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 0, 1, 62, 47,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
		74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 6, 2, 0, 39, 40, 58, 62, 2,
		0, 52, 52, 63, 63, 2, 0, 48, 51, 69, 69, 4, 0, 9, 10, 27, 27, 29, 29, 63,
		63, 2, 0, 64, 66, 70, 70, 1, 0, 42, 43, 623, 0, 95, 1, 0, 0, 0, 2, 101,
		1, 0, 0, 0, 4, 120, 1, 0, 0, 0, 6, 122, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0,
		10, 148, 1, 0, 0, 0, 12, 153, 1, 0, 0, 0, 14, 161, 1, 0, 0, 0, 16, 169,
		1, 0, 0, 0, 18, 172, 1, 0, 0, 0, 20, 193, 1, 0, 0, 0, 22, 202, 1, 0, 0,
		0, 24, 221, 1, 0, 0, 0, 26, 230, 1, 0, 0, 0, 28, 232, 1, 0, 0, 0, 30, 250,
		1, 0, 0, 0, 32, 252, 1, 0, 0, 0, 34, 254, 1, 0, 0, 0, 36, 256, 1, 0, 0,
		0, 38, 258, 1, 0, 0, 0, 40, 286, 1, 0, 0, 0, 42, 333, 1, 0, 0, 0, 44, 335,
		1, 0, 0, 0, 46, 339, 1, 0, 0, 0, 48, 354, 1, 0, 0, 0, 50, 368, 1, 0, 0,
		0, 52, 377, 1, 0, 0, 0, 54, 385, 1, 0, 0, 0, 56, 394, 1, 0, 0, 0, 58, 403,
		1, 0, 0, 0, 60, 407, 1, 0, 0, 0, 62, 431, 1, 0, 0, 0, 64, 441, 1, 0, 0,
		0, 66, 461, 1, 0, 0, 0, 68, 463, 1, 0, 0, 0, 70, 475, 1, 0, 0, 0, 72, 485,
		1, 0, 0, 0, 74, 494, 1, 0, 0, 0, 76, 504, 1, 0, 0, 0, 78, 506, 1, 0, 0,
		0, 80, 519, 1, 0, 0, 0, 82, 528, 1, 0, 0, 0, 84, 532, 1, 0, 0, 0, 86, 544,
		1, 0, 0, 0, 88, 564, 1, 0, 0, 0, 90, 566, 1, 0, 0, 0, 92, 568, 1, 0, 0,
		0, 94, 96, 3, 2, 1, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95,
		1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 5, 0, 0, 1,
		100, 1, 1, 0, 0, 0, 101, 102, 5, 4, 0, 0, 102, 103, 5, 63, 0, 0, 103, 105,
		5, 31, 0, 0, 104, 106, 3, 4, 2, 0, 105, 104, 1, 0, 0, 0, 106, 107, 1, 0,
		0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 110, 5, 32, 0, 0, 110, 3, 1, 0, 0, 0, 111, 121, 3, 10, 5, 0, 112,
		121, 3, 12, 6, 0, 113, 121, 3, 18, 9, 0, 114, 121, 3, 20, 10, 0, 115, 121,
		3, 24, 12, 0, 116, 121, 3, 84, 42, 0, 117, 121, 3, 82, 41, 0, 118, 121,
		3, 28, 14, 0, 119, 121, 3, 6, 3, 0, 120, 111, 1, 0, 0, 0, 120, 112, 1,
		0, 0, 0, 120, 113, 1, 0, 0, 0, 120, 114, 1, 0, 0, 0, 120, 115, 1, 0, 0,
		0, 120, 116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		119, 1, 0, 0, 0, 121, 5, 1, 0, 0, 0, 122, 123, 5, 26, 0, 0, 123, 127, 5,
		31, 0, 0, 124, 126, 3, 8, 4, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0,
		0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 130, 1, 0, 0, 0, 129,
		127, 1, 0, 0, 0, 130, 131, 5, 32, 0, 0, 131, 7, 1, 0, 0, 0, 132, 134, 5,
		63, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0,
		0, 135, 136, 5, 64, 0, 0, 136, 9, 1, 0, 0, 0, 137, 138, 5, 9, 0, 0, 138,
		140, 5, 31, 0, 0, 139, 141, 3, 64, 32, 0, 140, 139, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 144, 145, 5, 32, 0, 0, 145, 149, 1, 0, 0, 0, 146, 147, 5, 9, 0, 0,
		147, 149, 3, 64, 32, 0, 148, 137, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		11, 1, 0, 0, 0, 150, 152, 3, 80, 40, 0, 151, 150, 1, 0, 0, 0, 152, 155,
		1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 1, 0,
		0, 0, 155, 153, 1, 0, 0, 0, 156, 159, 5, 10, 0, 0, 157, 160, 3, 14, 7,
		0, 158, 160, 3, 16, 8, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160,
		13, 1, 0, 0, 0, 161, 163, 5, 31, 0, 0, 162, 164, 3, 64, 32, 0, 163, 162,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 32, 0, 0, 168, 15, 1, 0, 0, 0,
		169, 170, 5, 35, 0, 0, 170, 171, 5, 17, 0, 0, 171, 17, 1, 0, 0, 0, 172,
		173, 5, 27, 0, 0, 173, 174, 3, 86, 43, 0, 174, 175, 5, 32, 0, 0, 175, 19,
		1, 0, 0, 0, 176, 177, 5, 6, 0, 0, 177, 179, 5, 31, 0, 0, 178, 180, 5, 63,
		0, 0, 179, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 194, 5, 32, 0, 0, 184,
		185, 5, 6, 0, 0, 185, 187, 5, 31, 0, 0, 186, 188, 3, 22, 11, 0, 187, 186,
		1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 5, 32, 0, 0, 192, 194, 1, 0, 0, 0,
		193, 176, 1, 0, 0, 0, 193, 184, 1, 0, 0, 0, 194, 21, 1, 0, 0, 0, 195, 203,
		3, 38, 19, 0, 196, 203, 3, 68, 34, 0, 197, 203, 3, 42, 21, 0, 198, 203,
		5, 30, 0, 0, 199, 203, 5, 20, 0, 0, 200, 203, 3, 62, 31, 0, 201, 203, 3,
		66, 33, 0, 202, 195, 1, 0, 0, 0, 202, 196, 1, 0, 0, 0, 202, 197, 1, 0,
		0, 0, 202, 198, 1, 0, 0, 0, 202, 199, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0,
		202, 201, 1, 0, 0, 0, 203, 23, 1, 0, 0, 0, 204, 205, 5, 7, 0, 0, 205, 207,
		5, 31, 0, 0, 206, 208, 5, 63, 0, 0, 207, 206, 1, 0, 0, 0, 208, 209, 1,
		0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0,
		0, 211, 222, 5, 32, 0, 0, 212, 213, 5, 7, 0, 0, 213, 215, 5, 31, 0, 0,
		214, 216, 3, 26, 13, 0, 215, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217,
		215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220,
		5, 32, 0, 0, 220, 222, 1, 0, 0, 0, 221, 204, 1, 0, 0, 0, 221, 212, 1, 0,
		0, 0, 222, 25, 1, 0, 0, 0, 223, 231, 3, 38, 19, 0, 224, 231, 3, 68, 34,
		0, 225, 231, 3, 42, 21, 0, 226, 231, 5, 30, 0, 0, 227, 231, 5, 20, 0, 0,
		228, 231, 3, 62, 31, 0, 229, 231, 3, 66, 33, 0, 230, 223, 1, 0, 0, 0, 230,
		224, 1, 0, 0, 0, 230, 225, 1, 0, 0, 0, 230, 226, 1, 0, 0, 0, 230, 227,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231, 27, 1, 0,
		0, 0, 232, 233, 5, 63, 0, 0, 233, 237, 5, 31, 0, 0, 234, 236, 3, 30, 15,
		0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 241,
		5, 32, 0, 0, 241, 29, 1, 0, 0, 0, 242, 251, 3, 38, 19, 0, 243, 251, 3,
		68, 34, 0, 244, 251, 3, 42, 21, 0, 245, 251, 3, 48, 24, 0, 246, 251, 5,
		30, 0, 0, 247, 251, 5, 20, 0, 0, 248, 251, 3, 62, 31, 0, 249, 251, 3, 66,
		33, 0, 250, 242, 1, 0, 0, 0, 250, 243, 1, 0, 0, 0, 250, 244, 1, 0, 0, 0,
		250, 245, 1, 0, 0, 0, 250, 246, 1, 0, 0, 0, 250, 247, 1, 0, 0, 0, 250,
		248, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 31, 1, 0, 0, 0, 252, 253, 3,
		26, 13, 0, 253, 33, 1, 0, 0, 0, 254, 255, 3, 26, 13, 0, 255, 35, 1, 0,
		0, 0, 256, 257, 3, 26, 13, 0, 257, 37, 1, 0, 0, 0, 258, 259, 5, 18, 0,
		0, 259, 260, 5, 33, 0, 0, 260, 261, 3, 40, 20, 0, 261, 262, 5, 34, 0, 0,
		262, 266, 5, 31, 0, 0, 263, 265, 3, 32, 16, 0, 264, 263, 1, 0, 0, 0, 265,
		268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269,
		1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 279, 5, 32, 0, 0, 270, 271, 5, 19,
		0, 0, 271, 275, 5, 31, 0, 0, 272, 274, 3, 34, 17, 0, 273, 272, 1, 0, 0,
		0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276,
		278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 280, 5, 32, 0, 0, 279, 270,
		1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 39, 1, 0, 0, 0, 281, 282, 3, 62,
		31, 0, 282, 283, 7, 0, 0, 0, 283, 284, 3, 62, 31, 0, 284, 287, 1, 0, 0,
		0, 285, 287, 3, 62, 31, 0, 286, 281, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0,
		287, 41, 1, 0, 0, 0, 288, 289, 5, 21, 0, 0, 289, 290, 3, 44, 22, 0, 290,
		291, 5, 68, 0, 0, 291, 292, 3, 40, 20, 0, 292, 293, 5, 68, 0, 0, 293, 294,
		3, 46, 23, 0, 294, 298, 5, 31, 0, 0, 295, 297, 3, 36, 18, 0, 296, 295,
		1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0,
		0, 0, 299, 301, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 302, 5, 32, 0, 0,
		302, 334, 1, 0, 0, 0, 303, 304, 5, 21, 0, 0, 304, 305, 7, 1, 0, 0, 305,
		306, 5, 37, 0, 0, 306, 307, 5, 63, 0, 0, 307, 308, 5, 58, 0, 0, 308, 309,
		5, 22, 0, 0, 309, 310, 3, 62, 31, 0, 310, 314, 5, 31, 0, 0, 311, 313, 3,
		36, 18, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0,
		0, 0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0,
		317, 318, 5, 32, 0, 0, 318, 334, 1, 0, 0, 0, 319, 320, 5, 21, 0, 0, 320,
		321, 5, 63, 0, 0, 321, 322, 5, 58, 0, 0, 322, 323, 5, 22, 0, 0, 323, 324,
		3, 62, 31, 0, 324, 328, 5, 31, 0, 0, 325, 327, 3, 36, 18, 0, 326, 325,
		1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0,
		0, 0, 329, 331, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 332, 5, 32, 0, 0,
		332, 334, 1, 0, 0, 0, 333, 288, 1, 0, 0, 0, 333, 303, 1, 0, 0, 0, 333,
		319, 1, 0, 0, 0, 334, 43, 1, 0, 0, 0, 335, 336, 3, 68, 34, 0, 336, 337,
		5, 36, 0, 0, 337, 338, 3, 62, 31, 0, 338, 45, 1, 0, 0, 0, 339, 352, 3,
		68, 34, 0, 340, 353, 5, 46, 0, 0, 341, 353, 5, 47, 0, 0, 342, 343, 5, 53,
		0, 0, 343, 353, 3, 62, 31, 0, 344, 345, 5, 54, 0, 0, 345, 353, 3, 62, 31,
		0, 346, 347, 5, 55, 0, 0, 347, 353, 3, 62, 31, 0, 348, 349, 5, 56, 0, 0,
		349, 353, 3, 62, 31, 0, 350, 351, 5, 57, 0, 0, 351, 353, 3, 62, 31, 0,
		352, 340, 1, 0, 0, 0, 352, 341, 1, 0, 0, 0, 352, 342, 1, 0, 0, 0, 352,
		344, 1, 0, 0, 0, 352, 346, 1, 0, 0, 0, 352, 348, 1, 0, 0, 0, 352, 350,
		1, 0, 0, 0, 353, 47, 1, 0, 0, 0, 354, 355, 5, 23, 0, 0, 355, 356, 3, 40,
		20, 0, 356, 360, 5, 31, 0, 0, 357, 359, 3, 50, 25, 0, 358, 357, 1, 0, 0,
		0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361,
		364, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 365, 3, 52, 26, 0, 364, 363,
		1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 32,
		0, 0, 367, 49, 1, 0, 0, 0, 368, 369, 5, 25, 0, 0, 369, 370, 3, 40, 20,
		0, 370, 374, 5, 35, 0, 0, 371, 373, 3, 26, 13, 0, 372, 371, 1, 0, 0, 0,
		373, 376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375,
		51, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 378, 5, 24, 0, 0, 378, 382,
		5, 35, 0, 0, 379, 381, 3, 26, 13, 0, 380, 379, 1, 0, 0, 0, 381, 384, 1,
		0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 53, 1, 0, 0,
		0, 384, 382, 1, 0, 0, 0, 385, 386, 5, 25, 0, 0, 386, 387, 3, 58, 29, 0,
		387, 391, 5, 35, 0, 0, 388, 390, 3, 26, 13, 0, 389, 388, 1, 0, 0, 0, 390,
		393, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 55, 1,
		0, 0, 0, 393, 391, 1, 0, 0, 0, 394, 395, 5, 24, 0, 0, 395, 399, 5, 35,
		0, 0, 396, 398, 3, 26, 13, 0, 397, 396, 1, 0, 0, 0, 398, 401, 1, 0, 0,
		0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 57, 1, 0, 0, 0, 401,
		399, 1, 0, 0, 0, 402, 404, 5, 49, 0, 0, 403, 402, 1, 0, 0, 0, 403, 404,
		1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 406, 5, 63, 0, 0, 406, 59, 1, 0,
		0, 0, 407, 408, 3, 68, 34, 0, 408, 417, 5, 33, 0, 0, 409, 414, 3, 62, 31,
		0, 410, 411, 5, 37, 0, 0, 411, 413, 3, 62, 31, 0, 412, 410, 1, 0, 0, 0,
		413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415,
		418, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 409, 1, 0, 0, 0, 417, 418,
		1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 5, 34, 0, 0, 420, 61, 1, 0,
		0, 0, 421, 422, 6, 31, -1, 0, 422, 432, 3, 68, 34, 0, 423, 432, 5, 70,
		0, 0, 424, 432, 5, 64, 0, 0, 425, 432, 5, 65, 0, 0, 426, 432, 5, 66, 0,
		0, 427, 428, 5, 33, 0, 0, 428, 429, 3, 62, 31, 0, 429, 430, 5, 34, 0, 0,
		430, 432, 1, 0, 0, 0, 431, 421, 1, 0, 0, 0, 431, 423, 1, 0, 0, 0, 431,
		424, 1, 0, 0, 0, 431, 425, 1, 0, 0, 0, 431, 426, 1, 0, 0, 0, 431, 427,
		1, 0, 0, 0, 432, 438, 1, 0, 0, 0, 433, 434, 10, 7, 0, 0, 434, 435, 7, 2,
		0, 0, 435, 437, 3, 62, 31, 8, 436, 433, 1, 0, 0, 0, 437, 440, 1, 0, 0,
		0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 63, 1, 0, 0, 0, 440,
		438, 1, 0, 0, 0, 441, 442, 5, 63, 0, 0, 442, 443, 5, 35, 0, 0, 443, 446,
		3, 88, 44, 0, 444, 445, 5, 36, 0, 0, 445, 447, 3, 90, 45, 0, 446, 444,
		1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 451, 1, 0, 0, 0, 448, 450, 3, 72,
		36, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0,
		451, 452, 1, 0, 0, 0, 452, 455, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454,
		456, 5, 68, 0, 0, 455, 454, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 65,
		1, 0, 0, 0, 457, 462, 5, 70, 0, 0, 458, 462, 5, 64, 0, 0, 459, 462, 5,
		72, 0, 0, 460, 462, 3, 68, 34, 0, 461, 457, 1, 0, 0, 0, 461, 458, 1, 0,
		0, 0, 461, 459, 1, 0, 0, 0, 461, 460, 1, 0, 0, 0, 462, 67, 1, 0, 0, 0,
		463, 472, 7, 3, 0, 0, 464, 465, 5, 38, 0, 0, 465, 471, 5, 63, 0, 0, 466,
		467, 5, 44, 0, 0, 467, 468, 3, 62, 31, 0, 468, 469, 5, 45, 0, 0, 469, 471,
		1, 0, 0, 0, 470, 464, 1, 0, 0, 0, 470, 466, 1, 0, 0, 0, 471, 474, 1, 0,
		0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 69, 1, 0, 0, 0,
		474, 472, 1, 0, 0, 0, 475, 476, 5, 63, 0, 0, 476, 477, 5, 13, 0, 0, 477,
		479, 5, 31, 0, 0, 478, 480, 3, 64, 32, 0, 479, 478, 1, 0, 0, 0, 480, 481,
		1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 483, 1, 0,
		0, 0, 483, 484, 5, 32, 0, 0, 484, 71, 1, 0, 0, 0, 485, 486, 5, 41, 0, 0,
		486, 492, 5, 63, 0, 0, 487, 489, 5, 33, 0, 0, 488, 490, 3, 74, 37, 0, 489,
		488, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 493,
		5, 34, 0, 0, 492, 487, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 73, 1, 0,
		0, 0, 494, 499, 3, 76, 38, 0, 495, 496, 5, 37, 0, 0, 496, 498, 3, 76, 38,
		0, 497, 495, 1, 0, 0, 0, 498, 501, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 499,
		500, 1, 0, 0, 0, 500, 75, 1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 502, 505, 5,
		64, 0, 0, 503, 505, 3, 78, 39, 0, 504, 502, 1, 0, 0, 0, 504, 503, 1, 0,
		0, 0, 505, 77, 1, 0, 0, 0, 506, 515, 5, 31, 0, 0, 507, 512, 5, 64, 0, 0,
		508, 509, 5, 37, 0, 0, 509, 511, 5, 64, 0, 0, 510, 508, 1, 0, 0, 0, 511,
		514, 1, 0, 0, 0, 512, 510, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 516,
		1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 515, 507, 1, 0, 0, 0, 515, 516, 1, 0,
		0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 5, 32, 0, 0, 518, 79, 1, 0, 0, 0,
		519, 520, 5, 41, 0, 0, 520, 526, 5, 63, 0, 0, 521, 523, 5, 33, 0, 0, 522,
		524, 3, 74, 37, 0, 523, 522, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524, 525,
		1, 0, 0, 0, 525, 527, 5, 34, 0, 0, 526, 521, 1, 0, 0, 0, 526, 527, 1, 0,
		0, 0, 527, 81, 1, 0, 0, 0, 528, 529, 5, 28, 0, 0, 529, 530, 3, 86, 43,
		0, 530, 531, 5, 32, 0, 0, 531, 83, 1, 0, 0, 0, 532, 533, 5, 29, 0, 0, 533,
		534, 3, 86, 43, 0, 534, 535, 5, 32, 0, 0, 535, 85, 1, 0, 0, 0, 536, 543,
		5, 76, 0, 0, 537, 543, 5, 75, 0, 0, 538, 539, 5, 31, 0, 0, 539, 540, 3,
		86, 43, 0, 540, 541, 5, 32, 0, 0, 541, 543, 1, 0, 0, 0, 542, 536, 1, 0,
		0, 0, 542, 537, 1, 0, 0, 0, 542, 538, 1, 0, 0, 0, 543, 546, 1, 0, 0, 0,
		544, 542, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 87, 1, 0, 0, 0, 546, 544,
		1, 0, 0, 0, 547, 565, 1, 0, 0, 0, 548, 549, 5, 13, 0, 0, 549, 553, 5, 31,
		0, 0, 550, 552, 3, 64, 32, 0, 551, 550, 1, 0, 0, 0, 552, 555, 1, 0, 0,
		0, 553, 551, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 556, 1, 0, 0, 0, 555,
		553, 1, 0, 0, 0, 556, 565, 5, 32, 0, 0, 557, 565, 5, 2, 0, 0, 558, 565,
		5, 3, 0, 0, 559, 560, 5, 44, 0, 0, 560, 561, 5, 45, 0, 0, 561, 565, 3,
		88, 44, 0, 562, 565, 5, 1, 0, 0, 563, 565, 5, 63, 0, 0, 564, 547, 1, 0,
		0, 0, 564, 548, 1, 0, 0, 0, 564, 557, 1, 0, 0, 0, 564, 558, 1, 0, 0, 0,
		564, 559, 1, 0, 0, 0, 564, 562, 1, 0, 0, 0, 564, 563, 1, 0, 0, 0, 565,
		89, 1, 0, 0, 0, 566, 567, 7, 4, 0, 0, 567, 91, 1, 0, 0, 0, 568, 569, 7,
		5, 0, 0, 569, 93, 1, 0, 0, 0, 59, 97, 107, 120, 127, 133, 142, 148, 153,
		159, 165, 181, 189, 193, 202, 209, 217, 221, 230, 237, 250, 266, 275, 279,
		286, 298, 314, 328, 333, 352, 360, 364, 374, 382, 391, 399, 403, 414, 417,
		431, 438, 446, 451, 455, 461, 470, 472, 481, 489, 492, 499, 504, 512, 515,
		523, 526, 542, 544, 553, 564,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PromptDSLParserInit initializes any static state used to implement PromptDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromptDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromptDSLParserInit() {
	staticData := &PromptDSLParserParserStaticData
	staticData.once.Do(promptdslparserParserInit)
}

// NewPromptDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromptDSLParser(input antlr.TokenStream) *PromptDSLParser {
	PromptDSLParserInit()
	this := new(PromptDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromptDSLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromptDSLParser.g4"

	return this
}

// PromptDSLParser tokens.
const (
	PromptDSLParserEOF                = antlr.TokenEOF
	PromptDSLParserSTRING_TYPE        = 1
	PromptDSLParserFLOAT_TYPE         = 2
	PromptDSLParserINT_TYPE           = 3
	PromptDSLParserPROMPT             = 4
	PromptDSLParserPARAMS             = 5
	PromptDSLParserSYSTEM             = 6
	PromptDSLParserUSER               = 7
	PromptDSLParserNOTE               = 8
	PromptDSLParserINPUT              = 9
	PromptDSLParserOUTPUT             = 10
	PromptDSLParserFORMAT             = 11
	PromptDSLParserTYPE               = 12
	PromptDSLParserSTRUCT             = 13
	PromptDSLParserSCHEMA             = 14
	PromptDSLParserPARSE              = 15
	PromptDSLParserJSONFIX            = 16
	PromptDSLParserMARKDOWN           = 17
	PromptDSLParserIF                 = 18
	PromptDSLParserELSE               = 19
	PromptDSLParserOUTPUTSPEC         = 20
	PromptDSLParserFOR                = 21
	PromptDSLParserRANGE              = 22
	PromptDSLParserSWITCH             = 23
	PromptDSLParserDEFAULT            = 24
	PromptDSLParserCASE               = 25
	PromptDSLParserGOIMPORT           = 26
	PromptDSLParserBEFORE             = 27
	PromptDSLParserFIX                = 28
	PromptDSLParserAFTER              = 29
	PromptDSLParserARRAY_OUTPUTSPEC   = 30
	PromptDSLParserLBRACE             = 31
	PromptDSLParserRBRACE             = 32
	PromptDSLParserLPAREN             = 33
	PromptDSLParserRPAREN             = 34
	PromptDSLParserCOLON              = 35
	PromptDSLParserEQUAL              = 36
	PromptDSLParserCOMMA              = 37
	PromptDSLParserDOT                = 38
	PromptDSLParserEQEQ               = 39
	PromptDSLParserNOTEQ              = 40
	PromptDSLParserAT                 = 41
	PromptDSLParserMD                 = 42
	PromptDSLParserJSON               = 43
	PromptDSLParserLBRACK             = 44
	PromptDSLParserRBRACK             = 45
	PromptDSLParserINCREMENT          = 46
	PromptDSLParserDECREMENT          = 47
	PromptDSLParserMINUS              = 48
	PromptDSLParserSTAR               = 49
	PromptDSLParserSLASH              = 50
	PromptDSLParserMOD                = 51
	PromptDSLParserUNDERSCORE         = 52
	PromptDSLParserPLUSEQ             = 53
	PromptDSLParserMINUSEQ            = 54
	PromptDSLParserMULTEQ             = 55
	PromptDSLParserDIVEQ              = 56
	PromptDSLParserMODEQ              = 57
	PromptDSLParserDECL_ASSIGN        = 58
	PromptDSLParserLT                 = 59
	PromptDSLParserLTE                = 60
	PromptDSLParserGT                 = 61
	PromptDSLParserGTE                = 62
	PromptDSLParserID                 = 63
	PromptDSLParserSTRING             = 64
	PromptDSLParserNUMBER             = 65
	PromptDSLParserBOOL               = 66
	PromptDSLParserPIPE               = 67
	PromptDSLParserSEMI               = 68
	PromptDSLParserPLUS               = 69
	PromptDSLParserDASH_STRING        = 70
	PromptDSLParserWS                 = 71
	PromptDSLParserLINE_COMMENT       = 72
	PromptDSLParserBLOCK_COMMENT      = 73
	PromptDSLParserCODE_STRING_SINGLE = 74
	PromptDSLParserCODE_STRING        = 75
	PromptDSLParserCODE_TEXT          = 76
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile        = 0
	PromptDSLParserRULE_promptDef         = 1
	PromptDSLParserRULE_promptBlock       = 2
	PromptDSLParserRULE_goimportSection   = 3
	PromptDSLParserRULE_goimportEntry     = 4
	PromptDSLParserRULE_inputSection      = 5
	PromptDSLParserRULE_outputSection     = 6
	PromptDSLParserRULE_outputStruct      = 7
	PromptDSLParserRULE_outputMarkdown    = 8
	PromptDSLParserRULE_beforeSection     = 9
	PromptDSLParserRULE_systemSection     = 10
	PromptDSLParserRULE_sysContent        = 11
	PromptDSLParserRULE_userSection       = 12
	PromptDSLParserRULE_userContent       = 13
	PromptDSLParserRULE_moduleDef         = 14
	PromptDSLParserRULE_moduleContent     = 15
	PromptDSLParserRULE_thencontent       = 16
	PromptDSLParserRULE_elsecontent       = 17
	PromptDSLParserRULE_forcontent        = 18
	PromptDSLParserRULE_ifStatement       = 19
	PromptDSLParserRULE_condition         = 20
	PromptDSLParserRULE_forStatement      = 21
	PromptDSLParserRULE_assignExpr        = 22
	PromptDSLParserRULE_updateExpr        = 23
	PromptDSLParserRULE_switchStatement   = 24
	PromptDSLParserRULE_switchCase        = 25
	PromptDSLParserRULE_switchDefault     = 26
	PromptDSLParserRULE_typeCase          = 27
	PromptDSLParserRULE_typeDefault       = 28
	PromptDSLParserRULE_typeName          = 29
	PromptDSLParserRULE_dslCallExpr       = 30
	PromptDSLParserRULE_expr              = 31
	PromptDSLParserRULE_fieldDef          = 32
	PromptDSLParserRULE_textLine          = 33
	PromptDSLParserRULE_paramPath         = 34
	PromptDSLParserRULE_structDef         = 35
	PromptDSLParserRULE_annotation        = 36
	PromptDSLParserRULE_annotationArgs    = 37
	PromptDSLParserRULE_annotationValue   = 38
	PromptDSLParserRULE_arrayLiteral      = 39
	PromptDSLParserRULE_defaultAnnotation = 40
	PromptDSLParserRULE_fixSection        = 41
	PromptDSLParserRULE_afterSection      = 42
	PromptDSLParserRULE_codeBlockContent  = 43
	PromptDSLParserRULE_type              = 44
	PromptDSLParserRULE_value             = 45
	PromptDSLParserRULE_formatType        = 46
)

// IPromptFileContext is an interface to support dynamic dispatch.
type IPromptFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllPromptDef() []IPromptDefContext
	PromptDef(i int) IPromptDefContext

	// IsPromptFileContext differentiates from other interfaces.
	IsPromptFileContext()
}

type PromptFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptFileContext() *PromptFileContext {
	var p = new(PromptFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptFile
	return p
}

func InitEmptyPromptFileContext(p *PromptFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptFile
}

func (*PromptFileContext) IsPromptFileContext() {}

func NewPromptFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptFileContext {
	var p = new(PromptFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptFile

	return p
}

func (s *PromptFileContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEOF, 0)
}

func (s *PromptFileContext) AllPromptDef() []IPromptDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPromptDefContext); ok {
			len++
		}
	}

	tst := make([]IPromptDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPromptDefContext); ok {
			tst[i] = t.(IPromptDefContext)
			i++
		}
	}

	return tst
}

func (s *PromptFileContext) PromptDef(i int) IPromptDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromptDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromptDefContext)
}

func (s *PromptFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptFile(s)
	}
}

func (s *PromptFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptFile(s)
	}
}

func (p *PromptDSLParser) PromptFile() (localctx IPromptFileContext) {
	localctx = NewPromptFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromptDSLParserRULE_promptFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(94)
			p.PromptDef()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(PromptDSLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPromptDefContext is an interface to support dynamic dispatch.
type IPromptDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROMPT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllPromptBlock() []IPromptBlockContext
	PromptBlock(i int) IPromptBlockContext

	// IsPromptDefContext differentiates from other interfaces.
	IsPromptDefContext()
}

type PromptDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptDefContext() *PromptDefContext {
	var p = new(PromptDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptDef
	return p
}

func InitEmptyPromptDefContext(p *PromptDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptDef
}

func (*PromptDefContext) IsPromptDefContext() {}

func NewPromptDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptDefContext {
	var p = new(PromptDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptDef

	return p
}

func (s *PromptDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptDefContext) PROMPT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPROMPT, 0)
}

func (s *PromptDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *PromptDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *PromptDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *PromptDefContext) AllPromptBlock() []IPromptBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPromptBlockContext); ok {
			len++
		}
	}

	tst := make([]IPromptBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPromptBlockContext); ok {
			tst[i] = t.(IPromptBlockContext)
			i++
		}
	}

	return tst
}

func (s *PromptDefContext) PromptBlock(i int) IPromptBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromptBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromptBlockContext)
}

func (s *PromptDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptDef(s)
	}
}

func (s *PromptDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptDef(s)
	}
}

func (p *PromptDSLParser) PromptDef() (localctx IPromptDefContext) {
	localctx = NewPromptDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromptDSLParserRULE_promptDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9223369836824885568) != 0) {
		{
			p.SetState(104)
			p.PromptBlock()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPromptBlockContext is an interface to support dynamic dispatch.
type IPromptBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InputSection() IInputSectionContext
	OutputSection() IOutputSectionContext
	BeforeSection() IBeforeSectionContext
	SystemSection() ISystemSectionContext
	UserSection() IUserSectionContext
	AfterSection() IAfterSectionContext
	FixSection() IFixSectionContext
	ModuleDef() IModuleDefContext
	GoimportSection() IGoimportSectionContext

	// IsPromptBlockContext differentiates from other interfaces.
	IsPromptBlockContext()
}

type PromptBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromptBlockContext() *PromptBlockContext {
	var p = new(PromptBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptBlock
	return p
}

func InitEmptyPromptBlockContext(p *PromptBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_promptBlock
}

func (*PromptBlockContext) IsPromptBlockContext() {}

func NewPromptBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromptBlockContext {
	var p = new(PromptBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_promptBlock

	return p
}

func (s *PromptBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PromptBlockContext) InputSection() IInputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputSectionContext)
}

func (s *PromptBlockContext) OutputSection() IOutputSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputSectionContext)
}

func (s *PromptBlockContext) BeforeSection() IBeforeSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBeforeSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBeforeSectionContext)
}

func (s *PromptBlockContext) SystemSection() ISystemSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISystemSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISystemSectionContext)
}

func (s *PromptBlockContext) UserSection() IUserSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserSectionContext)
}

func (s *PromptBlockContext) AfterSection() IAfterSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAfterSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAfterSectionContext)
}

func (s *PromptBlockContext) FixSection() IFixSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFixSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFixSectionContext)
}

func (s *PromptBlockContext) ModuleDef() IModuleDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleDefContext)
}

func (s *PromptBlockContext) GoimportSection() IGoimportSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoimportSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoimportSectionContext)
}

func (s *PromptBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromptBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromptBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterPromptBlock(s)
	}
}

func (s *PromptBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitPromptBlock(s)
	}
}

func (p *PromptDSLParser) PromptBlock() (localctx IPromptBlockContext) {
	localctx = NewPromptBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromptDSLParserRULE_promptBlock)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.InputSection()
		}

	case PromptDSLParserOUTPUT, PromptDSLParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.OutputSection()
		}

	case PromptDSLParserBEFORE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.BeforeSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.UserSection()
		}

	case PromptDSLParserAFTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.AfterSection()
		}

	case PromptDSLParserFIX:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.FixSection()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(118)
			p.ModuleDef()
		}

	case PromptDSLParserGOIMPORT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(119)
			p.GoimportSection()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoimportSectionContext is an interface to support dynamic dispatch.
type IGoimportSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GOIMPORT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllGoimportEntry() []IGoimportEntryContext
	GoimportEntry(i int) IGoimportEntryContext

	// IsGoimportSectionContext differentiates from other interfaces.
	IsGoimportSectionContext()
}

type GoimportSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoimportSectionContext() *GoimportSectionContext {
	var p = new(GoimportSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_goimportSection
	return p
}

func InitEmptyGoimportSectionContext(p *GoimportSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_goimportSection
}

func (*GoimportSectionContext) IsGoimportSectionContext() {}

func NewGoimportSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoimportSectionContext {
	var p = new(GoimportSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_goimportSection

	return p
}

func (s *GoimportSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoimportSectionContext) GOIMPORT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserGOIMPORT, 0)
}

func (s *GoimportSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *GoimportSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *GoimportSectionContext) AllGoimportEntry() []IGoimportEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGoimportEntryContext); ok {
			len++
		}
	}

	tst := make([]IGoimportEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGoimportEntryContext); ok {
			tst[i] = t.(IGoimportEntryContext)
			i++
		}
	}

	return tst
}

func (s *GoimportSectionContext) GoimportEntry(i int) IGoimportEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoimportEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoimportEntryContext)
}

func (s *GoimportSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoimportSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoimportSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterGoimportSection(s)
	}
}

func (s *GoimportSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitGoimportSection(s)
	}
}

func (p *PromptDSLParser) GoimportSection() (localctx IGoimportSectionContext) {
	localctx = NewGoimportSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromptDSLParserRULE_goimportSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(PromptDSLParserGOIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserID || _la == PromptDSLParserSTRING {
		{
			p.SetState(124)
			p.GoimportEntry()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(130)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoimportEntryContext is an interface to support dynamic dispatch.
type IGoimportEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetPath returns the path token.
	GetPath() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetPath sets the path token.
	SetPath(antlr.Token)

	// Getter signatures
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsGoimportEntryContext differentiates from other interfaces.
	IsGoimportEntryContext()
}

type GoimportEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	path   antlr.Token
}

func NewEmptyGoimportEntryContext() *GoimportEntryContext {
	var p = new(GoimportEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_goimportEntry
	return p
}

func InitEmptyGoimportEntryContext(p *GoimportEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_goimportEntry
}

func (*GoimportEntryContext) IsGoimportEntryContext() {}

func NewGoimportEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoimportEntryContext {
	var p = new(GoimportEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_goimportEntry

	return p
}

func (s *GoimportEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *GoimportEntryContext) GetAlias() antlr.Token { return s.alias }

func (s *GoimportEntryContext) GetPath() antlr.Token { return s.path }

func (s *GoimportEntryContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *GoimportEntryContext) SetPath(v antlr.Token) { s.path = v }

func (s *GoimportEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *GoimportEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *GoimportEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoimportEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoimportEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterGoimportEntry(s)
	}
}

func (s *GoimportEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitGoimportEntry(s)
	}
}

func (p *PromptDSLParser) GoimportEntry() (localctx IGoimportEntryContext) {
	localctx = NewGoimportEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromptDSLParserRULE_goimportEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserID {
		{
			p.SetState(132)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*GoimportEntryContext).alias = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(135)

		var _m = p.Match(PromptDSLParserSTRING)

		localctx.(*GoimportEntryContext).path = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInputSectionContext is an interface to support dynamic dispatch.
type IInputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INPUT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsInputSectionContext differentiates from other interfaces.
	IsInputSectionContext()
}

type InputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputSectionContext() *InputSectionContext {
	var p = new(InputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_inputSection
	return p
}

func InitEmptyInputSectionContext(p *InputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_inputSection
}

func (*InputSectionContext) IsInputSectionContext() {}

func NewInputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputSectionContext {
	var p = new(InputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_inputSection

	return p
}

func (s *InputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputSectionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINPUT, 0)
}

func (s *InputSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *InputSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *InputSectionContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *InputSectionContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *InputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterInputSection(s)
	}
}

func (s *InputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitInputSection(s)
	}
}

func (p *PromptDSLParser) InputSection() (localctx IInputSectionContext) {
	localctx = NewInputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromptDSLParserRULE_inputSection)
	var _la int

	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(PromptDSLParserINPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(139)
				p.FieldDef()
			}

			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(PromptDSLParserINPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.FieldDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutputSectionContext is an interface to support dynamic dispatch.
type IOutputSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTPUT() antlr.TerminalNode
	OutputStruct() IOutputStructContext
	OutputMarkdown() IOutputMarkdownContext
	AllDefaultAnnotation() []IDefaultAnnotationContext
	DefaultAnnotation(i int) IDefaultAnnotationContext

	// IsOutputSectionContext differentiates from other interfaces.
	IsOutputSectionContext()
}

type OutputSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputSectionContext() *OutputSectionContext {
	var p = new(OutputSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputSection
	return p
}

func InitEmptyOutputSectionContext(p *OutputSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputSection
}

func (*OutputSectionContext) IsOutputSectionContext() {}

func NewOutputSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputSectionContext {
	var p = new(OutputSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputSection

	return p
}

func (s *OutputSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputSectionContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUT, 0)
}

func (s *OutputSectionContext) OutputStruct() IOutputStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputStructContext)
}

func (s *OutputSectionContext) OutputMarkdown() IOutputMarkdownContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutputMarkdownContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutputMarkdownContext)
}

func (s *OutputSectionContext) AllDefaultAnnotation() []IDefaultAnnotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefaultAnnotationContext); ok {
			len++
		}
	}

	tst := make([]IDefaultAnnotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefaultAnnotationContext); ok {
			tst[i] = t.(IDefaultAnnotationContext)
			i++
		}
	}

	return tst
}

func (s *OutputSectionContext) DefaultAnnotation(i int) IDefaultAnnotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultAnnotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultAnnotationContext)
}

func (s *OutputSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputSection(s)
	}
}

func (s *OutputSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputSection(s)
	}
}

func (p *PromptDSLParser) OutputSection() (localctx IOutputSectionContext) {
	localctx = NewOutputSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromptDSLParserRULE_outputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserAT {
		{
			p.SetState(150)
			p.DefaultAnnotation()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserLBRACE:
		{
			p.SetState(157)
			p.OutputStruct()
		}

	case PromptDSLParserCOLON:
		{
			p.SetState(158)
			p.OutputMarkdown()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutputStructContext is an interface to support dynamic dispatch.
type IOutputStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsOutputStructContext differentiates from other interfaces.
	IsOutputStructContext()
}

type OutputStructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputStructContext() *OutputStructContext {
	var p = new(OutputStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputStruct
	return p
}

func InitEmptyOutputStructContext(p *OutputStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputStruct
}

func (*OutputStructContext) IsOutputStructContext() {}

func NewOutputStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputStructContext {
	var p = new(OutputStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputStruct

	return p
}

func (s *OutputStructContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputStructContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *OutputStructContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *OutputStructContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *OutputStructContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *OutputStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputStruct(s)
	}
}

func (s *OutputStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputStruct(s)
	}
}

func (p *PromptDSLParser) OutputStruct() (localctx IOutputStructContext) {
	localctx = NewOutputStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromptDSLParserRULE_outputStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(162)
			p.FieldDef()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutputMarkdownContext is an interface to support dynamic dispatch.
type IOutputMarkdownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	MARKDOWN() antlr.TerminalNode

	// IsOutputMarkdownContext differentiates from other interfaces.
	IsOutputMarkdownContext()
}

type OutputMarkdownContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutputMarkdownContext() *OutputMarkdownContext {
	var p = new(OutputMarkdownContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown
	return p
}

func InitEmptyOutputMarkdownContext(p *OutputMarkdownContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown
}

func (*OutputMarkdownContext) IsOutputMarkdownContext() {}

func NewOutputMarkdownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutputMarkdownContext {
	var p = new(OutputMarkdownContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_outputMarkdown

	return p
}

func (s *OutputMarkdownContext) GetParser() antlr.Parser { return s.parser }

func (s *OutputMarkdownContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *OutputMarkdownContext) MARKDOWN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMARKDOWN, 0)
}

func (s *OutputMarkdownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutputMarkdownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutputMarkdownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterOutputMarkdown(s)
	}
}

func (s *OutputMarkdownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitOutputMarkdown(s)
	}
}

func (p *PromptDSLParser) OutputMarkdown() (localctx IOutputMarkdownContext) {
	localctx = NewOutputMarkdownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromptDSLParserRULE_outputMarkdown)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(PromptDSLParserMARKDOWN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBeforeSectionContext is an interface to support dynamic dispatch.
type IBeforeSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEFORE() antlr.TerminalNode
	CodeBlockContent() ICodeBlockContentContext
	RBRACE() antlr.TerminalNode

	// IsBeforeSectionContext differentiates from other interfaces.
	IsBeforeSectionContext()
}

type BeforeSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeforeSectionContext() *BeforeSectionContext {
	var p = new(BeforeSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeSection
	return p
}

func InitEmptyBeforeSectionContext(p *BeforeSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_beforeSection
}

func (*BeforeSectionContext) IsBeforeSectionContext() {}

func NewBeforeSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeforeSectionContext {
	var p = new(BeforeSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_beforeSection

	return p
}

func (s *BeforeSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *BeforeSectionContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBEFORE, 0)
}

func (s *BeforeSectionContext) CodeBlockContent() ICodeBlockContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *BeforeSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *BeforeSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeforeSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeforeSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterBeforeSection(s)
	}
}

func (s *BeforeSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitBeforeSection(s)
	}
}

func (p *PromptDSLParser) BeforeSection() (localctx IBeforeSectionContext) {
	localctx = NewBeforeSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromptDSLParserRULE_beforeSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(PromptDSLParserBEFORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.CodeBlockContent()
	}
	{
		p.SetState(174)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISystemSectionContext is an interface to support dynamic dispatch.
type ISystemSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYSTEM() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllSysContent() []ISysContentContext
	SysContent(i int) ISysContentContext

	// IsSystemSectionContext differentiates from other interfaces.
	IsSystemSectionContext()
}

type SystemSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemSectionContext() *SystemSectionContext {
	var p = new(SystemSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_systemSection
	return p
}

func InitEmptySystemSectionContext(p *SystemSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_systemSection
}

func (*SystemSectionContext) IsSystemSectionContext() {}

func NewSystemSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemSectionContext {
	var p = new(SystemSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_systemSection

	return p
}

func (s *SystemSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemSectionContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSYSTEM, 0)
}

func (s *SystemSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *SystemSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *SystemSectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *SystemSectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *SystemSectionContext) AllSysContent() []ISysContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISysContentContext); ok {
			len++
		}
	}

	tst := make([]ISysContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISysContentContext); ok {
			tst[i] = t.(ISysContentContext)
			i++
		}
	}

	return tst
}

func (s *SystemSectionContext) SysContent(i int) ISysContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISysContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISysContentContext)
}

func (s *SystemSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSystemSection(s)
	}
}

func (s *SystemSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSystemSection(s)
	}
}

func (p *PromptDSLParser) SystemSection() (localctx ISystemSectionContext) {
	localctx = NewSystemSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromptDSLParserRULE_systemSection)
	var _la int

	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(178)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(183)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0) {
			{
				p.SetState(186)
				p.SysContent()
			}

			p.SetState(189)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(191)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISysContentContext is an interface to support dynamic dispatch.
type ISysContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsSysContentContext differentiates from other interfaces.
	IsSysContentContext()
}

type SysContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySysContentContext() *SysContentContext {
	var p = new(SysContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysContent
	return p
}

func InitEmptySysContentContext(p *SysContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_sysContent
}

func (*SysContentContext) IsSysContentContext() {}

func NewSysContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SysContentContext {
	var p = new(SysContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_sysContent

	return p
}

func (s *SysContentContext) GetParser() antlr.Parser { return s.parser }

func (s *SysContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *SysContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *SysContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *SysContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *SysContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *SysContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SysContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *SysContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SysContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SysContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSysContent(s)
	}
}

func (s *SysContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSysContent(s)
	}
}

func (p *PromptDSLParser) SysContent() (localctx ISysContentContext) {
	localctx = NewSysContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromptDSLParserRULE_sysContent)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(199)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(200)
			p.expr(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(201)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUserSectionContext is an interface to support dynamic dispatch.
type IUserSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsUserSectionContext differentiates from other interfaces.
	IsUserSectionContext()
}

type UserSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserSectionContext() *UserSectionContext {
	var p = new(UserSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userSection
	return p
}

func InitEmptyUserSectionContext(p *UserSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userSection
}

func (*UserSectionContext) IsUserSectionContext() {}

func NewUserSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserSectionContext {
	var p = new(UserSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_userSection

	return p
}

func (s *UserSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *UserSectionContext) USER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserUSER, 0)
}

func (s *UserSectionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *UserSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *UserSectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *UserSectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *UserSectionContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *UserSectionContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *UserSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUserSection(s)
	}
}

func (s *UserSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUserSection(s)
	}
}

func (p *PromptDSLParser) UserSection() (localctx IUserSectionContext) {
	localctx = NewUserSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromptDSLParserRULE_userSection)
	var _la int

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(206)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(211)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0) {
			{
				p.SetState(214)
				p.UserContent()
			}

			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUserContentContext is an interface to support dynamic dispatch.
type IUserContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsUserContentContext differentiates from other interfaces.
	IsUserContentContext()
}

type UserContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContentContext() *UserContentContext {
	var p = new(UserContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userContent
	return p
}

func InitEmptyUserContentContext(p *UserContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_userContent
}

func (*UserContentContext) IsUserContentContext() {}

func NewUserContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContentContext {
	var p = new(UserContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_userContent

	return p
}

func (s *UserContentContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *UserContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *UserContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *UserContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *UserContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *UserContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UserContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *UserContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUserContent(s)
	}
}

func (s *UserContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUserContent(s)
	}
}

func (p *PromptDSLParser) UserContent() (localctx IUserContentContext) {
	localctx = NewUserContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromptDSLParserRULE_userContent)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(227)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(228)
			p.expr(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(229)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModuleDefContext is an interface to support dynamic dispatch.
type IModuleDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllModuleContent() []IModuleContentContext
	ModuleContent(i int) IModuleContentContext

	// IsModuleDefContext differentiates from other interfaces.
	IsModuleDefContext()
}

type ModuleDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDefContext() *ModuleDefContext {
	var p = new(ModuleDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleDef
	return p
}

func InitEmptyModuleDefContext(p *ModuleDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleDef
}

func (*ModuleDefContext) IsModuleDefContext() {}

func NewModuleDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDefContext {
	var p = new(ModuleDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_moduleDef

	return p
}

func (s *ModuleDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *ModuleDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ModuleDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ModuleDefContext) AllModuleContent() []IModuleContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModuleContentContext); ok {
			len++
		}
	}

	tst := make([]IModuleContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModuleContentContext); ok {
			tst[i] = t.(IModuleContentContext)
			i++
		}
	}

	return tst
}

func (s *ModuleDefContext) ModuleContent(i int) IModuleContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleContentContext)
}

func (s *ModuleDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterModuleDef(s)
	}
}

func (s *ModuleDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitModuleDef(s)
	}
}

func (p *PromptDSLParser) ModuleDef() (localctx IModuleDefContext) {
	localctx = NewModuleDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromptDSLParserRULE_moduleDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978643965) != 0 {
		{
			p.SetState(234)
			p.ModuleContent()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(240)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModuleContentContext is an interface to support dynamic dispatch.
type IModuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfStatement() IIfStatementContext
	ParamPath() IParamPathContext
	ForStatement() IForStatementContext
	SwitchStatement() ISwitchStatementContext
	ARRAY_OUTPUTSPEC() antlr.TerminalNode
	OUTPUTSPEC() antlr.TerminalNode
	Expr() IExprContext
	TextLine() ITextLineContext

	// IsModuleContentContext differentiates from other interfaces.
	IsModuleContentContext()
}

type ModuleContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContentContext() *ModuleContentContext {
	var p = new(ModuleContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleContent
	return p
}

func InitEmptyModuleContentContext(p *ModuleContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_moduleContent
}

func (*ModuleContentContext) IsModuleContentContext() {}

func NewModuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContentContext {
	var p = new(ModuleContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_moduleContent

	return p
}

func (s *ModuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContentContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *ModuleContentContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *ModuleContentContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *ModuleContentContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *ModuleContentContext) ARRAY_OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserARRAY_OUTPUTSPEC, 0)
}

func (s *ModuleContentContext) OUTPUTSPEC() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUTSPEC, 0)
}

func (s *ModuleContentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModuleContentContext) TextLine() ITextLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextLineContext)
}

func (s *ModuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterModuleContent(s)
	}
}

func (s *ModuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitModuleContent(s)
	}
}

func (p *PromptDSLParser) ModuleContent() (localctx IModuleContentContext) {
	localctx = NewModuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromptDSLParserRULE_moduleContent)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(245)
			p.SwitchStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(246)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(247)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(248)
			p.expr(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(249)
			p.TextLine()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IThencontentContext is an interface to support dynamic dispatch.
type IThencontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsThencontentContext differentiates from other interfaces.
	IsThencontentContext()
}

type ThencontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThencontentContext() *ThencontentContext {
	var p = new(ThencontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_thencontent
	return p
}

func InitEmptyThencontentContext(p *ThencontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_thencontent
}

func (*ThencontentContext) IsThencontentContext() {}

func NewThencontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThencontentContext {
	var p = new(ThencontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_thencontent

	return p
}

func (s *ThencontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ThencontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ThencontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThencontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThencontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterThencontent(s)
	}
}

func (s *ThencontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitThencontent(s)
	}
}

func (p *PromptDSLParser) Thencontent() (localctx IThencontentContext) {
	localctx = NewThencontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromptDSLParserRULE_thencontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.UserContent()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElsecontentContext is an interface to support dynamic dispatch.
type IElsecontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsElsecontentContext differentiates from other interfaces.
	IsElsecontentContext()
}

type ElsecontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsecontentContext() *ElsecontentContext {
	var p = new(ElsecontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_elsecontent
	return p
}

func InitEmptyElsecontentContext(p *ElsecontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_elsecontent
}

func (*ElsecontentContext) IsElsecontentContext() {}

func NewElsecontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsecontentContext {
	var p = new(ElsecontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_elsecontent

	return p
}

func (s *ElsecontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsecontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ElsecontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsecontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsecontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterElsecontent(s)
	}
}

func (s *ElsecontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitElsecontent(s)
	}
}

func (p *PromptDSLParser) Elsecontent() (localctx IElsecontentContext) {
	localctx = NewElsecontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromptDSLParserRULE_elsecontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.UserContent()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForcontentContext is an interface to support dynamic dispatch.
type IForcontentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UserContent() IUserContentContext

	// IsForcontentContext differentiates from other interfaces.
	IsForcontentContext()
}

type ForcontentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForcontentContext() *ForcontentContext {
	var p = new(ForcontentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forcontent
	return p
}

func InitEmptyForcontentContext(p *ForcontentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forcontent
}

func (*ForcontentContext) IsForcontentContext() {}

func NewForcontentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForcontentContext {
	var p = new(ForcontentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_forcontent

	return p
}

func (s *ForcontentContext) GetParser() antlr.Parser { return s.parser }

func (s *ForcontentContext) UserContent() IUserContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *ForcontentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForcontentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForcontentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForcontent(s)
	}
}

func (s *ForcontentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForcontent(s)
	}
}

func (p *PromptDSLParser) Forcontent() (localctx IForcontentContext) {
	localctx = NewForcontentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromptDSLParserRULE_forcontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.UserContent()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Condition() IConditionContext
	RPAREN() antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode
	AllThencontent() []IThencontentContext
	Thencontent(i int) IThencontentContext
	ELSE() antlr.TerminalNode
	AllElsecontent() []IElsecontentContext
	Elsecontent(i int) IElsecontentContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserIF, 0)
}

func (s *IfStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *IfStatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *IfStatementContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserLBRACE)
}

func (s *IfStatementContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, i)
}

func (s *IfStatementContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserRBRACE)
}

func (s *IfStatementContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, i)
}

func (s *IfStatementContext) AllThencontent() []IThencontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThencontentContext); ok {
			len++
		}
	}

	tst := make([]IThencontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThencontentContext); ok {
			tst[i] = t.(IThencontentContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Thencontent(i int) IThencontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThencontentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThencontentContext)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserELSE, 0)
}

func (s *IfStatementContext) AllElsecontent() []IElsecontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElsecontentContext); ok {
			len++
		}
	}

	tst := make([]IElsecontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElsecontentContext); ok {
			tst[i] = t.(IElsecontentContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Elsecontent(i int) IElsecontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsecontentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsecontentContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *PromptDSLParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromptDSLParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(PromptDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Condition()
	}
	{
		p.SetState(261)
		p.Match(PromptDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
		{
			p.SetState(263)
			p.Thencontent()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserELSE {
		{
			p.SetState(270)
			p.Match(PromptDSLParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
			{
				p.SetState(272)
				p.Elsecontent()
			}

			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(278)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLhs returns the lhs rule contexts.
	GetLhs() IExprContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// GetSingle returns the single rule contexts.
	GetSingle() IExprContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IExprContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// SetSingle sets the single rule contexts.
	SetSingle(IExprContext)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	EQEQ() antlr.TerminalNode
	NOTEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	DECL_ASSIGN() antlr.TerminalNode

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IExprContext
	op     antlr.Token
	rhs    IExprContext
	single IExprContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetOp() antlr.Token { return s.op }

func (s *ConditionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ConditionContext) GetLhs() IExprContext { return s.lhs }

func (s *ConditionContext) GetRhs() IExprContext { return s.rhs }

func (s *ConditionContext) GetSingle() IExprContext { return s.single }

func (s *ConditionContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ConditionContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ConditionContext) SetSingle(v IExprContext) { s.single = v }

func (s *ConditionContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQEQ, 0)
}

func (s *ConditionContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNOTEQ, 0)
}

func (s *ConditionContext) LT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLT, 0)
}

func (s *ConditionContext) LTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLTE, 0)
}

func (s *ConditionContext) GT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserGT, 0)
}

func (s *ConditionContext) GTE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserGTE, 0)
}

func (s *ConditionContext) DECL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDECL_ASSIGN, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *PromptDSLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromptDSLParserRULE_condition)
	var _la int

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)

			var _x = p.expr(0)

			localctx.(*ConditionContext).lhs = _x
		}
		{
			p.SetState(282)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConditionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8935143309970505728) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConditionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(283)

			var _x = p.expr(0)

			localctx.(*ConditionContext).rhs = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)

			var _x = p.expr(0)

			localctx.(*ConditionContext).single = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) CopyAll(ctx *ForStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForTraditionalContext struct {
	ForStatementContext
	init   IAssignExprContext
	update IUpdateExprContext
}

func NewForTraditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForTraditionalContext {
	var p = new(ForTraditionalContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForTraditionalContext) GetInit() IAssignExprContext { return s.init }

func (s *ForTraditionalContext) GetUpdate() IUpdateExprContext { return s.update }

func (s *ForTraditionalContext) SetInit(v IAssignExprContext) { s.init = v }

func (s *ForTraditionalContext) SetUpdate(v IUpdateExprContext) { s.update = v }

func (s *ForTraditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForTraditionalContext) FOR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFOR, 0)
}

func (s *ForTraditionalContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSEMI)
}

func (s *ForTraditionalContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSEMI, i)
}

func (s *ForTraditionalContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ForTraditionalContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ForTraditionalContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ForTraditionalContext) AssignExpr() IAssignExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignExprContext)
}

func (s *ForTraditionalContext) UpdateExpr() IUpdateExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateExprContext)
}

func (s *ForTraditionalContext) AllForcontent() []IForcontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForcontentContext); ok {
			len++
		}
	}

	tst := make([]IForcontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForcontentContext); ok {
			tst[i] = t.(IForcontentContext)
			i++
		}
	}

	return tst
}

func (s *ForTraditionalContext) Forcontent(i int) IForcontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForcontentContext)
}

func (s *ForTraditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForTraditional(s)
	}
}

func (s *ForTraditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForTraditional(s)
	}
}

type ForRangeWithIndexContext struct {
	ForStatementContext
	key      antlr.Token
	val      antlr.Token
	iterable IExprContext
}

func NewForRangeWithIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeWithIndexContext {
	var p = new(ForRangeWithIndexContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForRangeWithIndexContext) GetKey() antlr.Token { return s.key }

func (s *ForRangeWithIndexContext) GetVal() antlr.Token { return s.val }

func (s *ForRangeWithIndexContext) SetKey(v antlr.Token) { s.key = v }

func (s *ForRangeWithIndexContext) SetVal(v antlr.Token) { s.val = v }

func (s *ForRangeWithIndexContext) GetIterable() IExprContext { return s.iterable }

func (s *ForRangeWithIndexContext) SetIterable(v IExprContext) { s.iterable = v }

func (s *ForRangeWithIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeWithIndexContext) FOR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFOR, 0)
}

func (s *ForRangeWithIndexContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, 0)
}

func (s *ForRangeWithIndexContext) DECL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDECL_ASSIGN, 0)
}

func (s *ForRangeWithIndexContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRANGE, 0)
}

func (s *ForRangeWithIndexContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ForRangeWithIndexContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ForRangeWithIndexContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *ForRangeWithIndexContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *ForRangeWithIndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeWithIndexContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserUNDERSCORE, 0)
}

func (s *ForRangeWithIndexContext) AllForcontent() []IForcontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForcontentContext); ok {
			len++
		}
	}

	tst := make([]IForcontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForcontentContext); ok {
			tst[i] = t.(IForcontentContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeWithIndexContext) Forcontent(i int) IForcontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForcontentContext)
}

func (s *ForRangeWithIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForRangeWithIndex(s)
	}
}

func (s *ForRangeWithIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForRangeWithIndex(s)
	}
}

type ForRangeNoIndexContext struct {
	ForStatementContext
	val      antlr.Token
	iterable IExprContext
}

func NewForRangeNoIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeNoIndexContext {
	var p = new(ForRangeNoIndexContext)

	InitEmptyForStatementContext(&p.ForStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForStatementContext))

	return p
}

func (s *ForRangeNoIndexContext) GetVal() antlr.Token { return s.val }

func (s *ForRangeNoIndexContext) SetVal(v antlr.Token) { s.val = v }

func (s *ForRangeNoIndexContext) GetIterable() IExprContext { return s.iterable }

func (s *ForRangeNoIndexContext) SetIterable(v IExprContext) { s.iterable = v }

func (s *ForRangeNoIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeNoIndexContext) FOR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFOR, 0)
}

func (s *ForRangeNoIndexContext) DECL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDECL_ASSIGN, 0)
}

func (s *ForRangeNoIndexContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRANGE, 0)
}

func (s *ForRangeNoIndexContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ForRangeNoIndexContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ForRangeNoIndexContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *ForRangeNoIndexContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeNoIndexContext) AllForcontent() []IForcontentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForcontentContext); ok {
			len++
		}
	}

	tst := make([]IForcontentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForcontentContext); ok {
			tst[i] = t.(IForcontentContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeNoIndexContext) Forcontent(i int) IForcontentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForcontentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForcontentContext)
}

func (s *ForRangeNoIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterForRangeNoIndex(s)
	}
}

func (s *ForRangeNoIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitForRangeNoIndex(s)
	}
}

func (p *PromptDSLParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PromptDSLParserRULE_forStatement)
	var _la int

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForTraditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _x = p.AssignExpr()

			localctx.(*ForTraditionalContext).init = _x
		}
		{
			p.SetState(290)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Condition()
		}
		{
			p.SetState(292)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _x = p.UpdateExpr()

			localctx.(*ForTraditionalContext).update = _x
		}
		{
			p.SetState(294)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
			{
				p.SetState(295)
				p.Forcontent()
			}

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(301)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForRangeWithIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ForRangeWithIndexContext).key = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PromptDSLParserUNDERSCORE || _la == PromptDSLParserID) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ForRangeWithIndexContext).key = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(305)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*ForRangeWithIndexContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(PromptDSLParserDECL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(PromptDSLParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _x = p.expr(0)

			localctx.(*ForRangeWithIndexContext).iterable = _x
		}
		{
			p.SetState(310)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
			{
				p.SetState(311)
				p.Forcontent()
			}

			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(317)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewForRangeNoIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*ForRangeNoIndexContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(PromptDSLParserDECL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(PromptDSLParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)

			var _x = p.expr(0)

			localctx.(*ForRangeNoIndexContext).iterable = _x
		}
		{
			p.SetState(324)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
			{
				p.SetState(325)
				p.Forcontent()
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(331)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignExprContext is an interface to support dynamic dispatch.
type IAssignExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLhs returns the lhs rule contexts.
	GetLhs() IParamPathContext

	// GetRhs returns the rhs rule contexts.
	GetRhs() IExprContext

	// SetLhs sets the lhs rule contexts.
	SetLhs(IParamPathContext)

	// SetRhs sets the rhs rule contexts.
	SetRhs(IExprContext)

	// Getter signatures
	EQUAL() antlr.TerminalNode
	ParamPath() IParamPathContext
	Expr() IExprContext

	// IsAssignExprContext differentiates from other interfaces.
	IsAssignExprContext()
}

type AssignExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lhs    IParamPathContext
	rhs    IExprContext
}

func NewEmptyAssignExprContext() *AssignExprContext {
	var p = new(AssignExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_assignExpr
	return p
}

func InitEmptyAssignExprContext(p *AssignExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_assignExpr
}

func (*AssignExprContext) IsAssignExprContext() {}

func NewAssignExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExprContext {
	var p = new(AssignExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_assignExpr

	return p
}

func (s *AssignExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExprContext) GetLhs() IParamPathContext { return s.lhs }

func (s *AssignExprContext) GetRhs() IExprContext { return s.rhs }

func (s *AssignExprContext) SetLhs(v IParamPathContext) { s.lhs = v }

func (s *AssignExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *AssignExprContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQUAL, 0)
}

func (s *AssignExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *AssignExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (p *PromptDSLParser) AssignExpr() (localctx IAssignExprContext) {
	localctx = NewAssignExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PromptDSLParserRULE_assignExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)

		var _x = p.ParamPath()

		localctx.(*AssignExprContext).lhs = _x
	}
	{
		p.SetState(336)
		p.Match(PromptDSLParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)

		var _x = p.expr(0)

		localctx.(*AssignExprContext).rhs = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdateExprContext is an interface to support dynamic dispatch.
type IUpdateExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamPath() IParamPathContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	PLUSEQ() antlr.TerminalNode
	Expr() IExprContext
	MINUSEQ() antlr.TerminalNode
	MULTEQ() antlr.TerminalNode
	DIVEQ() antlr.TerminalNode
	MODEQ() antlr.TerminalNode

	// IsUpdateExprContext differentiates from other interfaces.
	IsUpdateExprContext()
}

type UpdateExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateExprContext() *UpdateExprContext {
	var p = new(UpdateExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_updateExpr
	return p
}

func InitEmptyUpdateExprContext(p *UpdateExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_updateExpr
}

func (*UpdateExprContext) IsUpdateExprContext() {}

func NewUpdateExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateExprContext {
	var p = new(UpdateExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_updateExpr

	return p
}

func (s *UpdateExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *UpdateExprContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINCREMENT, 0)
}

func (s *UpdateExprContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDECREMENT, 0)
}

func (s *UpdateExprContext) PLUSEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUSEQ, 0)
}

func (s *UpdateExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UpdateExprContext) MINUSEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMINUSEQ, 0)
}

func (s *UpdateExprContext) MULTEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMULTEQ, 0)
}

func (s *UpdateExprContext) DIVEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDIVEQ, 0)
}

func (s *UpdateExprContext) MODEQ() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMODEQ, 0)
}

func (s *UpdateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterUpdateExpr(s)
	}
}

func (s *UpdateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitUpdateExpr(s)
	}
}

func (p *PromptDSLParser) UpdateExpr() (localctx IUpdateExprContext) {
	localctx = NewUpdateExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PromptDSLParserRULE_updateExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.ParamPath()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINCREMENT:
		{
			p.SetState(340)
			p.Match(PromptDSLParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserDECREMENT:
		{
			p.SetState(341)
			p.Match(PromptDSLParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserPLUSEQ:
		{
			p.SetState(342)
			p.Match(PromptDSLParserPLUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.expr(0)
		}

	case PromptDSLParserMINUSEQ:
		{
			p.SetState(344)
			p.Match(PromptDSLParserMINUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.expr(0)
		}

	case PromptDSLParserMULTEQ:
		{
			p.SetState(346)
			p.Match(PromptDSLParserMULTEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.expr(0)
		}

	case PromptDSLParserDIVEQ:
		{
			p.SetState(348)
			p.Match(PromptDSLParserDIVEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.expr(0)
		}

	case PromptDSLParserMODEQ:
		{
			p.SetState(350)
			p.Match(PromptDSLParserMODEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.expr(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Condition() IConditionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	SwitchDefault() ISwitchDefaultContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSWITCH, 0)
}

func (s *SwitchStatementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *SwitchStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *SwitchStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) SwitchDefault() ISwitchDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDefaultContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *PromptDSLParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PromptDSLParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(PromptDSLParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.Condition()
	}
	{
		p.SetState(356)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserCASE {
		{
			p.SetState(357)
			p.SwitchCase()
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserDEFAULT {
		{
			p.SetState(363)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(366)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Condition() IConditionContext
	COLON() antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCASE, 0)
}

func (s *SwitchCaseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *SwitchCaseContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (p *PromptDSLParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PromptDSLParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(PromptDSLParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Condition()
	}
	{
		p.SetState(370)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
		{
			p.SetState(371)
			p.UserContent()
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchDefaultContext is an interface to support dynamic dispatch.
type ISwitchDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsSwitchDefaultContext differentiates from other interfaces.
	IsSwitchDefaultContext()
}

type SwitchDefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchDefaultContext() *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchDefault
	return p
}

func InitEmptySwitchDefaultContext(p *SwitchDefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_switchDefault
}

func (*SwitchDefaultContext) IsSwitchDefaultContext() {}

func NewSwitchDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_switchDefault

	return p
}

func (s *SwitchDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDEFAULT, 0)
}

func (s *SwitchDefaultContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *SwitchDefaultContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *SwitchDefaultContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *SwitchDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterSwitchDefault(s)
	}
}

func (s *SwitchDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitSwitchDefault(s)
	}
}

func (p *PromptDSLParser) SwitchDefault() (localctx ISwitchDefaultContext) {
	localctx = NewSwitchDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromptDSLParserRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(PromptDSLParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
		{
			p.SetState(379)
			p.UserContent()
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeCaseContext is an interface to support dynamic dispatch.
type ITypeCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	TypeName() ITypeNameContext
	COLON() antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsTypeCaseContext differentiates from other interfaces.
	IsTypeCaseContext()
}

type TypeCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeCaseContext() *TypeCaseContext {
	var p = new(TypeCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeCase
	return p
}

func InitEmptyTypeCaseContext(p *TypeCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeCase
}

func (*TypeCaseContext) IsTypeCaseContext() {}

func NewTypeCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeCaseContext {
	var p = new(TypeCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_typeCase

	return p
}

func (s *TypeCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCASE, 0)
}

func (s *TypeCaseContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *TypeCaseContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *TypeCaseContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *TypeCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterTypeCase(s)
	}
}

func (s *TypeCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitTypeCase(s)
	}
}

func (p *PromptDSLParser) TypeCase() (localctx ITypeCaseContext) {
	localctx = NewTypeCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromptDSLParserRULE_typeCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(PromptDSLParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.TypeName()
	}
	{
		p.SetState(387)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
		{
			p.SetState(388)
			p.UserContent()
		}

		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDefaultContext is an interface to support dynamic dispatch.
type ITypeDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AllUserContent() []IUserContentContext
	UserContent(i int) IUserContentContext

	// IsTypeDefaultContext differentiates from other interfaces.
	IsTypeDefaultContext()
}

type TypeDefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefaultContext() *TypeDefaultContext {
	var p = new(TypeDefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeDefault
	return p
}

func InitEmptyTypeDefaultContext(p *TypeDefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeDefault
}

func (*TypeDefaultContext) IsTypeDefaultContext() {}

func NewTypeDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefaultContext {
	var p = new(TypeDefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_typeDefault

	return p
}

func (s *TypeDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDEFAULT, 0)
}

func (s *TypeDefaultContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *TypeDefaultContext) AllUserContent() []IUserContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUserContentContext); ok {
			len++
		}
	}

	tst := make([]IUserContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUserContentContext); ok {
			tst[i] = t.(IUserContentContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefaultContext) UserContent(i int) IUserContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContentContext)
}

func (s *TypeDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterTypeDefault(s)
	}
}

func (s *TypeDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitTypeDefault(s)
	}
}

func (p *PromptDSLParser) TypeDefault() (localctx ITypeDefaultContext) {
	localctx = NewTypeDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromptDSLParserRULE_typeDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(PromptDSLParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&-6647313049978660349) != 0 {
		{
			p.SetState(396)
			p.UserContent()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STAR() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *TypeNameContext) STAR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTAR, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *PromptDSLParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PromptDSLParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTAR {
		{
			p.SetState(402)
			p.Match(PromptDSLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(405)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDslCallExprContext is an interface to support dynamic dispatch.
type IDslCallExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParamPath() IParamPathContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDslCallExprContext differentiates from other interfaces.
	IsDslCallExprContext()
}

type DslCallExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDslCallExprContext() *DslCallExprContext {
	var p = new(DslCallExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr
	return p
}

func InitEmptyDslCallExprContext(p *DslCallExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr
}

func (*DslCallExprContext) IsDslCallExprContext() {}

func NewDslCallExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DslCallExprContext {
	var p = new(DslCallExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_dslCallExpr

	return p
}

func (s *DslCallExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DslCallExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *DslCallExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *DslCallExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *DslCallExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DslCallExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DslCallExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *DslCallExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *DslCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DslCallExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DslCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterDslCallExpr(s)
	}
}

func (s *DslCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitDslCallExpr(s)
	}
}

func (p *PromptDSLParser) DslCallExpr() (localctx IDslCallExprContext) {
	localctx = NewDslCallExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.ParamPath()
	}
	{
		p.SetState(408)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&2576058986874011651) != 0 {
		{
			p.SetState(409)
			p.expr(0)
		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(410)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(411)
				p.expr(0)
			}

			p.SetState(416)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(419)
		p.Match(PromptDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ParamPath() IParamPathContext
	DASH_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *ExprContext) DASH_STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDASH_STRING, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, 0)
}

func (s *ExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBOOL, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMINUS, 0)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTAR, 0)
}

func (s *ExprContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSLASH, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMOD, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *PromptDSLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PromptDSLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, PromptDSLParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		{
			p.SetState(422)
			p.ParamPath()
		}

	case PromptDSLParserDASH_STRING:
		{
			p.SetState(423)
			p.Match(PromptDSLParserDASH_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserSTRING:
		{
			p.SetState(424)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		{
			p.SetState(425)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		{
			p.SetState(426)
			p.Match(PromptDSLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLPAREN:
		{
			p.SetState(427)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.expr(0)
		}
		{
			p.SetState(429)
			p.Match(PromptDSLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PromptDSLParserRULE_expr)
			p.SetState(433)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				goto errorExit
			}
			{
				p.SetState(434)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-48)) & ^0x3f) == 0 && ((int64(1)<<(_la-48))&2097167) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(435)
				p.expr(8)
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldDefContext is an interface to support dynamic dispatch.
type IFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EQUAL() antlr.TerminalNode
	Value() IValueContext
	AllAnnotation() []IAnnotationContext
	Annotation(i int) IAnnotationContext
	SEMI() antlr.TerminalNode

	// IsFieldDefContext differentiates from other interfaces.
	IsFieldDefContext()
}

type FieldDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefContext() *FieldDefContext {
	var p = new(FieldDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fieldDef
	return p
}

func InitEmptyFieldDefContext(p *FieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fieldDef
}

func (*FieldDefContext) IsFieldDefContext() {}

func NewFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefContext {
	var p = new(FieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_fieldDef

	return p
}

func (s *FieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *FieldDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *FieldDefContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserEQUAL, 0)
}

func (s *FieldDefContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FieldDefContext) AllAnnotation() []IAnnotationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationContext); ok {
			tst[i] = t.(IAnnotationContext)
			i++
		}
	}

	return tst
}

func (s *FieldDefContext) Annotation(i int) IAnnotationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FieldDefContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSEMI, 0)
}

func (s *FieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFieldDef(s)
	}
}

func (s *FieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFieldDef(s)
	}
}

func (p *PromptDSLParser) FieldDef() (localctx IFieldDefContext) {
	localctx = NewFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromptDSLParserRULE_fieldDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
		p.Type_()
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserEQUAL {
		{
			p.SetState(444)
			p.Match(PromptDSLParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Value()
		}

	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(448)
				p.Annotation()
			}

		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(454)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextLineContext is an interface to support dynamic dispatch.
type ITextLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DASH_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LINE_COMMENT() antlr.TerminalNode
	ParamPath() IParamPathContext

	// IsTextLineContext differentiates from other interfaces.
	IsTextLineContext()
}

type TextLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextLineContext() *TextLineContext {
	var p = new(TextLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textLine
	return p
}

func InitEmptyTextLineContext(p *TextLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_textLine
}

func (*TextLineContext) IsTextLineContext() {}

func NewTextLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextLineContext {
	var p = new(TextLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_textLine

	return p
}

func (s *TextLineContext) GetParser() antlr.Parser { return s.parser }

func (s *TextLineContext) DASH_STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDASH_STRING, 0)
}

func (s *TextLineContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *TextLineContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLINE_COMMENT, 0)
}

func (s *TextLineContext) ParamPath() IParamPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamPathContext)
}

func (s *TextLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterTextLine(s)
	}
}

func (s *TextLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitTextLine(s)
	}
}

func (p *PromptDSLParser) TextLine() (localctx ITextLineContext) {
	localctx = NewTextLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromptDSLParserRULE_textLine)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserDASH_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(457)
			p.Match(PromptDSLParserDASH_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(458)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(459)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(460)
			p.ParamPath()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamPathContext is an interface to support dynamic dispatch.
type IParamPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	INPUT() antlr.TerminalNode
	OUTPUT() antlr.TerminalNode
	AFTER() antlr.TerminalNode
	BEFORE() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode

	// IsParamPathContext differentiates from other interfaces.
	IsParamPathContext()
}

type ParamPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamPathContext() *ParamPathContext {
	var p = new(ParamPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramPath
	return p
}

func InitEmptyParamPathContext(p *ParamPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_paramPath
}

func (*ParamPathContext) IsParamPathContext() {}

func NewParamPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamPathContext {
	var p = new(ParamPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_paramPath

	return p
}

func (s *ParamPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamPathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserID)
}

func (s *ParamPathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, i)
}

func (s *ParamPathContext) INPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINPUT, 0)
}

func (s *ParamPathContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserOUTPUT, 0)
}

func (s *ParamPathContext) AFTER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAFTER, 0)
}

func (s *ParamPathContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBEFORE, 0)
}

func (s *ParamPathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserDOT)
}

func (s *ParamPathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDOT, i)
}

func (s *ParamPathContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserLBRACK)
}

func (s *ParamPathContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACK, i)
}

func (s *ParamPathContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ParamPathContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamPathContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserRBRACK)
}

func (s *ParamPathContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACK, i)
}

func (s *ParamPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterParamPath(s)
	}
}

func (s *ParamPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitParamPath(s)
	}
}

func (p *PromptDSLParser) ParamPath() (localctx IParamPathContext) {
	localctx = NewParamPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PromptDSLParserRULE_paramPath)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-9223372036183685632) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(470)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case PromptDSLParserDOT:
				{
					p.SetState(464)
					p.Match(PromptDSLParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(465)
					p.Match(PromptDSLParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case PromptDSLParserLBRACK:
				{
					p.SetState(466)
					p.Match(PromptDSLParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(467)
					p.expr(0)
				}
				{
					p.SetState(468)
					p.Match(PromptDSLParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDefContext is an interface to support dynamic dispatch.
type IStructDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsStructDefContext differentiates from other interfaces.
	IsStructDefContext()
}

type StructDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefContext() *StructDefContext {
	var p = new(StructDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_structDef
	return p
}

func InitEmptyStructDefContext(p *StructDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_structDef
}

func (*StructDefContext) IsStructDefContext() {}

func NewStructDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefContext {
	var p = new(StructDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_structDef

	return p
}

func (s *StructDefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *StructDefContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
}

func (s *StructDefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *StructDefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *StructDefContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *StructDefContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *StructDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterStructDef(s)
	}
}

func (s *StructDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitStructDef(s)
	}
}

func (p *PromptDSLParser) StructDef() (localctx IStructDefContext) {
	localctx = NewStructDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(478)
			p.FieldDef()
		}

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(483)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AnnotationArgs() IAnnotationArgsContext

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotation
	return p
}

func InitEmptyAnnotationContext(p *AnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotation
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) AT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAT, 0)
}

func (s *AnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *AnnotationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *AnnotationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *AnnotationContext) AnnotationArgs() IAnnotationArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationArgsContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *PromptDSLParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(486)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(487)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(488)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(491)
			p.Match(PromptDSLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationArgsContext is an interface to support dynamic dispatch.
type IAnnotationArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnnotationValue() []IAnnotationValueContext
	AnnotationValue(i int) IAnnotationValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAnnotationArgsContext differentiates from other interfaces.
	IsAnnotationArgsContext()
}

type AnnotationArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationArgsContext() *AnnotationArgsContext {
	var p = new(AnnotationArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationArgs
	return p
}

func InitEmptyAnnotationArgsContext(p *AnnotationArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationArgs
}

func (*AnnotationArgsContext) IsAnnotationArgsContext() {}

func NewAnnotationArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationArgsContext {
	var p = new(AnnotationArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotationArgs

	return p
}

func (s *AnnotationArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationArgsContext) AllAnnotationValue() []IAnnotationValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotationValueContext); ok {
			len++
		}
	}

	tst := make([]IAnnotationValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotationValueContext); ok {
			tst[i] = t.(IAnnotationValueContext)
			i++
		}
	}

	return tst
}

func (s *AnnotationArgsContext) AnnotationValue(i int) IAnnotationValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationValueContext)
}

func (s *AnnotationArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *AnnotationArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *AnnotationArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotationArgs(s)
	}
}

func (s *AnnotationArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotationArgs(s)
	}
}

func (p *PromptDSLParser) AnnotationArgs() (localctx IAnnotationArgsContext) {
	localctx = NewAnnotationArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.AnnotationValue()
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserCOMMA {
		{
			p.SetState(495)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)
			p.AnnotationValue()
		}

		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnnotationValueContext is an interface to support dynamic dispatch.
type IAnnotationValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsAnnotationValueContext differentiates from other interfaces.
	IsAnnotationValueContext()
}

type AnnotationValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationValueContext() *AnnotationValueContext {
	var p = new(AnnotationValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationValue
	return p
}

func InitEmptyAnnotationValueContext(p *AnnotationValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_annotationValue
}

func (*AnnotationValueContext) IsAnnotationValueContext() {}

func NewAnnotationValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationValueContext {
	var p = new(AnnotationValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_annotationValue

	return p
}

func (s *AnnotationValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *AnnotationValueContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *AnnotationValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAnnotationValue(s)
	}
}

func (s *AnnotationValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAnnotationValue(s)
	}
}

func (p *PromptDSLParser) AnnotationValue() (localctx IAnnotationValueContext) {
	localctx = NewAnnotationValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PromptDSLParserRULE_annotationValue)
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(502)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)
			p.ArrayLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *ArrayLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *ArrayLiteralContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserSTRING)
}

func (s *ArrayLiteralContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, i)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *PromptDSLParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(507)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(508)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(509)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(514)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(517)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultAnnotationContext is an interface to support dynamic dispatch.
type IDefaultAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AnnotationArgs() IAnnotationArgsContext

	// IsDefaultAnnotationContext differentiates from other interfaces.
	IsDefaultAnnotationContext()
}

type DefaultAnnotationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultAnnotationContext() *DefaultAnnotationContext {
	var p = new(DefaultAnnotationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation
	return p
}

func InitEmptyDefaultAnnotationContext(p *DefaultAnnotationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation
}

func (*DefaultAnnotationContext) IsDefaultAnnotationContext() {}

func NewDefaultAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultAnnotationContext {
	var p = new(DefaultAnnotationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_defaultAnnotation

	return p
}

func (s *DefaultAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultAnnotationContext) AT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAT, 0)
}

func (s *DefaultAnnotationContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *DefaultAnnotationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *DefaultAnnotationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *DefaultAnnotationContext) AnnotationArgs() IAnnotationArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotationArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotationArgsContext)
}

func (s *DefaultAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterDefaultAnnotation(s)
	}
}

func (s *DefaultAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitDefaultAnnotation(s)
	}
}

func (p *PromptDSLParser) DefaultAnnotation() (localctx IDefaultAnnotationContext) {
	localctx = NewDefaultAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PromptDSLParserRULE_defaultAnnotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(521)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(522)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(525)
			p.Match(PromptDSLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFixSectionContext is an interface to support dynamic dispatch.
type IFixSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIX() antlr.TerminalNode
	CodeBlockContent() ICodeBlockContentContext
	RBRACE() antlr.TerminalNode

	// IsFixSectionContext differentiates from other interfaces.
	IsFixSectionContext()
}

type FixSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixSectionContext() *FixSectionContext {
	var p = new(FixSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fixSection
	return p
}

func InitEmptyFixSectionContext(p *FixSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_fixSection
}

func (*FixSectionContext) IsFixSectionContext() {}

func NewFixSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FixSectionContext {
	var p = new(FixSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_fixSection

	return p
}

func (s *FixSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *FixSectionContext) FIX() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFIX, 0)
}

func (s *FixSectionContext) CodeBlockContent() ICodeBlockContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *FixSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *FixSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FixSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FixSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFixSection(s)
	}
}

func (s *FixSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFixSection(s)
	}
}

func (p *PromptDSLParser) FixSection() (localctx IFixSectionContext) {
	localctx = NewFixSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PromptDSLParserRULE_fixSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(PromptDSLParserFIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
		p.CodeBlockContent()
	}
	{
		p.SetState(530)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAfterSectionContext is an interface to support dynamic dispatch.
type IAfterSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFTER() antlr.TerminalNode
	CodeBlockContent() ICodeBlockContentContext
	RBRACE() antlr.TerminalNode

	// IsAfterSectionContext differentiates from other interfaces.
	IsAfterSectionContext()
}

type AfterSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAfterSectionContext() *AfterSectionContext {
	var p = new(AfterSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterSection
	return p
}

func InitEmptyAfterSectionContext(p *AfterSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_afterSection
}

func (*AfterSectionContext) IsAfterSectionContext() {}

func NewAfterSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AfterSectionContext {
	var p = new(AfterSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_afterSection

	return p
}

func (s *AfterSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *AfterSectionContext) AFTER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserAFTER, 0)
}

func (s *AfterSectionContext) CodeBlockContent() ICodeBlockContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *AfterSectionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *AfterSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AfterSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AfterSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterAfterSection(s)
	}
}

func (s *AfterSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitAfterSection(s)
	}
}

func (p *PromptDSLParser) AfterSection() (localctx IAfterSectionContext) {
	localctx = NewAfterSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PromptDSLParserRULE_afterSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(PromptDSLParserAFTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(533)
		p.CodeBlockContent()
	}
	{
		p.SetState(534)
		p.Match(PromptDSLParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeBlockContentContext is an interface to support dynamic dispatch.
type ICodeBlockContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCODE_TEXT() []antlr.TerminalNode
	CODE_TEXT(i int) antlr.TerminalNode
	AllCODE_STRING() []antlr.TerminalNode
	CODE_STRING(i int) antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	AllCodeBlockContent() []ICodeBlockContentContext
	CodeBlockContent(i int) ICodeBlockContentContext
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode

	// IsCodeBlockContentContext differentiates from other interfaces.
	IsCodeBlockContentContext()
}

type CodeBlockContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContentContext() *CodeBlockContentContext {
	var p = new(CodeBlockContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent
	return p
}

func InitEmptyCodeBlockContentContext(p *CodeBlockContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent
}

func (*CodeBlockContentContext) IsCodeBlockContentContext() {}

func NewCodeBlockContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContentContext {
	var p = new(CodeBlockContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_codeBlockContent

	return p
}

func (s *CodeBlockContentContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContentContext) AllCODE_TEXT() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCODE_TEXT)
}

func (s *CodeBlockContentContext) CODE_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCODE_TEXT, i)
}

func (s *CodeBlockContentContext) AllCODE_STRING() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCODE_STRING)
}

func (s *CodeBlockContentContext) CODE_STRING(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCODE_STRING, i)
}

func (s *CodeBlockContentContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserLBRACE)
}

func (s *CodeBlockContentContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, i)
}

func (s *CodeBlockContentContext) AllCodeBlockContent() []ICodeBlockContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			len++
		}
	}

	tst := make([]ICodeBlockContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICodeBlockContentContext); ok {
			tst[i] = t.(ICodeBlockContentContext)
			i++
		}
	}

	return tst
}

func (s *CodeBlockContentContext) CodeBlockContent(i int) ICodeBlockContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContentContext)
}

func (s *CodeBlockContentContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserRBRACE)
}

func (s *CodeBlockContentContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, i)
}

func (s *CodeBlockContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterCodeBlockContent(s)
	}
}

func (s *CodeBlockContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitCodeBlockContent(s)
	}
}

func (p *PromptDSLParser) CodeBlockContent() (localctx ICodeBlockContentContext) {
	localctx = NewCodeBlockContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PromptDSLParserRULE_codeBlockContent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-31)) & ^0x3f) == 0 && ((int64(1)<<(_la-31))&52776558133249) != 0 {
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromptDSLParserCODE_TEXT:
			{
				p.SetState(536)
				p.Match(PromptDSLParserCODE_TEXT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserCODE_STRING:
			{
				p.SetState(537)
				p.Match(PromptDSLParserCODE_STRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserLBRACE:
			{
				p.SetState(538)
				p.Match(PromptDSLParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(539)
				p.CodeBlockContent()
			}
			{
				p.SetState(540)
				p.Match(PromptDSLParserRBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
	FLOAT_TYPE() antlr.TerminalNode
	INT_TYPE() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() ITypeContext
	STRING_TYPE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRUCT, 0)
}

func (s *TypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACE, 0)
}

func (s *TypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACE, 0)
}

func (s *TypeContext) AllFieldDef() []IFieldDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefContext); ok {
			tst[i] = t.(IFieldDefContext)
			i++
		}
	}

	return tst
}

func (s *TypeContext) FieldDef(i int) IFieldDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefContext)
}

func (s *TypeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserFLOAT_TYPE, 0)
}

func (s *TypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserINT_TYPE, 0)
}

func (s *TypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACK, 0)
}

func (s *TypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACK, 0)
}

func (s *TypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING_TYPE, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *PromptDSLParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PromptDSLParserRULE_type)
	var _la int

	p.SetState(564)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserID {
			{
				p.SetState(550)
				p.FieldDef()
			}

			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(556)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(557)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(558)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(559)
			p.Match(PromptDSLParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
			p.Match(PromptDSLParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
			p.Type_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(562)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(563)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DASH_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) DASH_STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserDASH_STRING, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, 0)
}

func (s *ValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserBOOL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *PromptDSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(566)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&71) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormatTypeContext is an interface to support dynamic dispatch.
type IFormatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MD() antlr.TerminalNode
	JSON() antlr.TerminalNode

	// IsFormatTypeContext differentiates from other interfaces.
	IsFormatTypeContext()
}

type FormatTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatTypeContext() *FormatTypeContext {
	var p = new(FormatTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_formatType
	return p
}

func InitEmptyFormatTypeContext(p *FormatTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_formatType
}

func (*FormatTypeContext) IsFormatTypeContext() {}

func NewFormatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatTypeContext {
	var p = new(FormatTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_formatType

	return p
}

func (s *FormatTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatTypeContext) MD() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserMD, 0)
}

func (s *FormatTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserJSON, 0)
}

func (s *FormatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterFormatType(s)
	}
}

func (s *FormatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitFormatType(s)
	}
}

func (p *PromptDSLParser) FormatType() (localctx IFormatTypeContext) {
	localctx = NewFormatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(568)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromptDSLParserMD || _la == PromptDSLParserJSON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PromptDSLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 31:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromptDSLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
