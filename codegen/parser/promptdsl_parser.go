// Code generated from ./codegen/grammar/P/PromptDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"'goimport'", "'runexe'", "'it'", "'it_ctx'", "'loopRange'", "", "",
		"", "", "'{'", "", "'('", "')'", "':'", "'='", "','", "'.'", "'=='",
		"'!='", "'@'", "'md'", "'json'", "'['", "']'", "'++'", "'--'", "'-'",
		"'*'", "'/'", "'%'", "'_'", "'+='", "'-='", "'*='", "'/='", "'%='",
		"':='", "'<'", "'<='", "'>'", "'>='", "", "", "", "", "'|'", "';'",
		"'+'",
	}
	staticData.SymbolicNames = []string{
		"", "STRING_TYPE", "FLOAT_TYPE", "INT_TYPE", "PROMPT", "PARAMS", "SYSTEM",
		"USER", "NOTE", "INPUT", "OUTPUT", "FORMAT", "TYPE", "STRUCT", "SCHEMA",
		"PARSE", "JSONFIX", "MARKDOWN", "IF", "ELSE", "OUTPUTSPEC", "FOR", "RANGE",
		"SWITCH", "DEFAULT", "CASE", "GOIMPORT", "RUNEXE", "IT", "IT_CTX", "LOOPRANGE",
		"BEFORE", "FIX", "AFTER", "ARRAY_OUTPUTSPEC", "LBRACE", "RBRACE", "LPAREN",
		"RPAREN", "COLON", "EQUAL", "COMMA", "DOT", "EQEQ", "NOTEQ", "AT", "MD",
		"JSON", "LBRACK", "RBRACK", "INCREMENT", "DECREMENT", "MINUS", "STAR",
		"SLASH", "MOD", "UNDERSCORE", "PLUSEQ", "MINUSEQ", "MULTEQ", "DIVEQ",
		"MODEQ", "DECL_ASSIGN", "LT", "LTE", "GT", "GTE", "ID", "STRING", "NUMBER",
		"BOOL", "PIPE", "SEMI", "PLUS", "DASH_STRING", "WS", "LINE_COMMENT",
		"BLOCK_COMMENT", "CODE_STRING_SINGLE", "CODE_STRING", "CODE_TEXT",
	}
	staticData.RuleNames = []string{
		"promptFile", "runexeDef", "promptDef", "promptBlock", "goimportSection",
		"goimportEntry", "inputSection", "iterater", "loopRangeSection", "outputSection",
		"outputStruct", "outputMarkdown", "beforeSection", "systemSection",
		"sysContent", "userSection", "userContent", "moduleDef", "moduleContent",
		"thencontent", "elsecontent", "forcontent", "ifStatement", "condition",
		"forStatement", "assignExpr", "updateExpr", "switchStatement", "switchCase",
		"switchDefault", "typeCase", "typeDefault", "typeName", "dslCallExpr",
		"expr", "fieldDef", "textLine", "paramPath", "structDef", "annotation",
		"annotationArgs", "annotationValue", "arrayLiteral", "defaultAnnotation",
		"fixSection", "afterSection", "codeBlockContent", "type", "value", "formatType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 620, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12, 0,
		105, 9, 0, 1, 0, 4, 0, 108, 8, 0, 11, 0, 12, 0, 109, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 122, 8, 2, 11, 2, 12, 2,
		123, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 139, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 144, 8, 4, 10, 4,
		12, 4, 147, 9, 4, 1, 4, 1, 4, 1, 5, 3, 5, 152, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 4, 6, 159, 8, 6, 11, 6, 12, 6, 160, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 167, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 180, 8, 8, 1, 8, 1, 8, 1, 9, 5, 9, 185, 8, 9, 10, 9,
		12, 9, 188, 9, 9, 1, 9, 1, 9, 1, 9, 3, 9, 193, 8, 9, 1, 10, 1, 10, 4, 10,
		197, 8, 10, 11, 10, 12, 10, 198, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 4, 13, 213, 8, 13, 11, 13,
		12, 13, 214, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 221, 8, 13, 11, 13, 12,
		13, 222, 1, 13, 1, 13, 3, 13, 227, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 237, 8, 14, 1, 15, 1, 15, 1, 15, 4, 15,
		242, 8, 15, 11, 15, 12, 15, 243, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 250,
		8, 15, 11, 15, 12, 15, 251, 1, 15, 1, 15, 3, 15, 256, 8, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 266, 8, 16, 1, 17,
		1, 17, 1, 17, 5, 17, 271, 8, 17, 10, 17, 12, 17, 274, 9, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 286,
		8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 5, 22, 300, 8, 22, 10, 22, 12, 22, 303, 9, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 5, 22, 309, 8, 22, 10, 22, 12, 22, 312, 9, 22, 1,
		22, 3, 22, 315, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 322, 8,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 332,
		8, 24, 10, 24, 12, 24, 335, 9, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 348, 8, 24, 10, 24, 12, 24,
		351, 9, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 362, 8, 24, 10, 24, 12, 24, 365, 9, 24, 1, 24, 1, 24, 3, 24,
		369, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 388,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 394, 8, 27, 10, 27, 12, 27, 397,
		9, 27, 1, 27, 3, 27, 400, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 5, 28, 408, 8, 28, 10, 28, 12, 28, 411, 9, 28, 1, 29, 1, 29, 1, 29,
		5, 29, 416, 8, 29, 10, 29, 12, 29, 419, 9, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 5, 30, 425, 8, 30, 10, 30, 12, 30, 428, 9, 30, 1, 31, 1, 31, 1, 31,
		5, 31, 433, 8, 31, 10, 31, 12, 31, 436, 9, 31, 1, 32, 3, 32, 439, 8, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 448, 8, 33, 10,
		33, 12, 33, 451, 9, 33, 3, 33, 453, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 466, 8, 34, 1,
		34, 1, 34, 1, 34, 5, 34, 471, 8, 34, 10, 34, 12, 34, 474, 9, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 481, 8, 35, 1, 35, 5, 35, 484, 8, 35,
		10, 35, 12, 35, 487, 9, 35, 1, 35, 3, 35, 490, 8, 35, 1, 36, 1, 36, 1,
		36, 1, 36, 3, 36, 496, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 5, 37, 505, 8, 37, 10, 37, 12, 37, 508, 9, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 517, 8, 37, 10, 37, 12, 37, 520,
		9, 37, 3, 37, 522, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 4, 38, 528, 8, 38,
		11, 38, 12, 38, 529, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 538,
		8, 39, 1, 39, 3, 39, 541, 8, 39, 1, 40, 1, 40, 1, 40, 5, 40, 546, 8, 40,
		10, 40, 12, 40, 549, 9, 40, 1, 41, 1, 41, 3, 41, 553, 8, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 5, 42, 559, 8, 42, 10, 42, 12, 42, 562, 9, 42, 3, 42,
		564, 8, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 572, 8, 43,
		1, 43, 3, 43, 575, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 592,
		8, 46, 10, 46, 12, 46, 595, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 601,
		8, 47, 10, 47, 12, 47, 604, 9, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 3, 47, 614, 8, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49,
		0, 1, 68, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 6, 2,
		0, 43, 44, 62, 66, 2, 0, 56, 56, 67, 67, 2, 0, 52, 55, 73, 73, 4, 0, 9,
		10, 28, 28, 31, 31, 33, 33, 2, 0, 68, 70, 74, 74, 1, 0, 46, 47, 678, 0,
		103, 1, 0, 0, 0, 2, 113, 1, 0, 0, 0, 4, 117, 1, 0, 0, 0, 6, 138, 1, 0,
		0, 0, 8, 140, 1, 0, 0, 0, 10, 151, 1, 0, 0, 0, 12, 166, 1, 0, 0, 0, 14,
		168, 1, 0, 0, 0, 16, 173, 1, 0, 0, 0, 18, 186, 1, 0, 0, 0, 20, 194, 1,
		0, 0, 0, 22, 202, 1, 0, 0, 0, 24, 205, 1, 0, 0, 0, 26, 226, 1, 0, 0, 0,
		28, 236, 1, 0, 0, 0, 30, 255, 1, 0, 0, 0, 32, 265, 1, 0, 0, 0, 34, 267,
		1, 0, 0, 0, 36, 285, 1, 0, 0, 0, 38, 287, 1, 0, 0, 0, 40, 289, 1, 0, 0,
		0, 42, 291, 1, 0, 0, 0, 44, 293, 1, 0, 0, 0, 46, 321, 1, 0, 0, 0, 48, 368,
		1, 0, 0, 0, 50, 370, 1, 0, 0, 0, 52, 374, 1, 0, 0, 0, 54, 389, 1, 0, 0,
		0, 56, 403, 1, 0, 0, 0, 58, 412, 1, 0, 0, 0, 60, 420, 1, 0, 0, 0, 62, 429,
		1, 0, 0, 0, 64, 438, 1, 0, 0, 0, 66, 442, 1, 0, 0, 0, 68, 465, 1, 0, 0,
		0, 70, 475, 1, 0, 0, 0, 72, 495, 1, 0, 0, 0, 74, 521, 1, 0, 0, 0, 76, 523,
		1, 0, 0, 0, 78, 533, 1, 0, 0, 0, 80, 542, 1, 0, 0, 0, 82, 552, 1, 0, 0,
		0, 84, 554, 1, 0, 0, 0, 86, 567, 1, 0, 0, 0, 88, 576, 1, 0, 0, 0, 90, 580,
		1, 0, 0, 0, 92, 593, 1, 0, 0, 0, 94, 613, 1, 0, 0, 0, 96, 615, 1, 0, 0,
		0, 98, 617, 1, 0, 0, 0, 100, 102, 3, 2, 1, 0, 101, 100, 1, 0, 0, 0, 102,
		105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 107,
		1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 108, 3, 4, 2, 0, 107, 106, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 112, 5, 0, 0, 1, 112, 1, 1, 0, 0, 0, 113, 114,
		5, 27, 0, 0, 114, 115, 5, 39, 0, 0, 115, 116, 5, 68, 0, 0, 116, 3, 1, 0,
		0, 0, 117, 118, 5, 4, 0, 0, 118, 119, 5, 67, 0, 0, 119, 121, 5, 35, 0,
		0, 120, 122, 3, 6, 3, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126,
		5, 36, 0, 0, 126, 5, 1, 0, 0, 0, 127, 139, 3, 12, 6, 0, 128, 139, 3, 14,
		7, 0, 129, 139, 3, 16, 8, 0, 130, 139, 3, 18, 9, 0, 131, 139, 3, 24, 12,
		0, 132, 139, 3, 26, 13, 0, 133, 139, 3, 30, 15, 0, 134, 139, 3, 90, 45,
		0, 135, 139, 3, 88, 44, 0, 136, 139, 3, 34, 17, 0, 137, 139, 3, 8, 4, 0,
		138, 127, 1, 0, 0, 0, 138, 128, 1, 0, 0, 0, 138, 129, 1, 0, 0, 0, 138,
		130, 1, 0, 0, 0, 138, 131, 1, 0, 0, 0, 138, 132, 1, 0, 0, 0, 138, 133,
		1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 135, 1, 0, 0, 0, 138, 136, 1, 0,
		0, 0, 138, 137, 1, 0, 0, 0, 139, 7, 1, 0, 0, 0, 140, 141, 5, 26, 0, 0,
		141, 145, 5, 35, 0, 0, 142, 144, 3, 10, 5, 0, 143, 142, 1, 0, 0, 0, 144,
		147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148,
		1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 36, 0, 0, 149, 9, 1, 0,
		0, 0, 150, 152, 5, 67, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 154, 5, 68, 0, 0, 154, 11, 1, 0, 0, 0, 155,
		156, 5, 9, 0, 0, 156, 158, 5, 35, 0, 0, 157, 159, 3, 70, 35, 0, 158, 157,
		1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 36, 0, 0, 163, 167, 1, 0, 0, 0,
		164, 165, 5, 9, 0, 0, 165, 167, 3, 70, 35, 0, 166, 155, 1, 0, 0, 0, 166,
		164, 1, 0, 0, 0, 167, 13, 1, 0, 0, 0, 168, 169, 5, 29, 0, 0, 169, 170,
		5, 37, 0, 0, 170, 171, 5, 68, 0, 0, 171, 172, 5, 38, 0, 0, 172, 15, 1,
		0, 0, 0, 173, 174, 5, 30, 0, 0, 174, 179, 5, 48, 0, 0, 175, 176, 5, 69,
		0, 0, 176, 177, 5, 41, 0, 0, 177, 180, 5, 69, 0, 0, 178, 180, 5, 69, 0,
		0, 179, 175, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 5, 49, 0, 0, 182, 17, 1, 0, 0, 0, 183, 185, 3, 86, 43, 0, 184, 183,
		1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0,
		0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 192, 5, 10, 0, 0,
		190, 193, 3, 20, 10, 0, 191, 193, 3, 22, 11, 0, 192, 190, 1, 0, 0, 0, 192,
		191, 1, 0, 0, 0, 193, 19, 1, 0, 0, 0, 194, 196, 5, 35, 0, 0, 195, 197,
		3, 70, 35, 0, 196, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1,
		0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 5, 36, 0,
		0, 201, 21, 1, 0, 0, 0, 202, 203, 5, 39, 0, 0, 203, 204, 5, 17, 0, 0, 204,
		23, 1, 0, 0, 0, 205, 206, 5, 31, 0, 0, 206, 207, 3, 92, 46, 0, 207, 208,
		5, 36, 0, 0, 208, 25, 1, 0, 0, 0, 209, 210, 5, 6, 0, 0, 210, 212, 5, 35,
		0, 0, 211, 213, 5, 67, 0, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0,
		214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216,
		227, 5, 36, 0, 0, 217, 218, 5, 6, 0, 0, 218, 220, 5, 35, 0, 0, 219, 221,
		3, 28, 14, 0, 220, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 220, 1,
		0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 5, 36, 0,
		0, 225, 227, 1, 0, 0, 0, 226, 209, 1, 0, 0, 0, 226, 217, 1, 0, 0, 0, 227,
		27, 1, 0, 0, 0, 228, 237, 3, 44, 22, 0, 229, 237, 3, 74, 37, 0, 230, 237,
		3, 48, 24, 0, 231, 237, 3, 54, 27, 0, 232, 237, 5, 34, 0, 0, 233, 237,
		5, 20, 0, 0, 234, 237, 3, 68, 34, 0, 235, 237, 3, 72, 36, 0, 236, 228,
		1, 0, 0, 0, 236, 229, 1, 0, 0, 0, 236, 230, 1, 0, 0, 0, 236, 231, 1, 0,
		0, 0, 236, 232, 1, 0, 0, 0, 236, 233, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0,
		236, 235, 1, 0, 0, 0, 237, 29, 1, 0, 0, 0, 238, 239, 5, 7, 0, 0, 239, 241,
		5, 35, 0, 0, 240, 242, 5, 67, 0, 0, 241, 240, 1, 0, 0, 0, 242, 243, 1,
		0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0, 0,
		0, 245, 256, 5, 36, 0, 0, 246, 247, 5, 7, 0, 0, 247, 249, 5, 35, 0, 0,
		248, 250, 3, 32, 16, 0, 249, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251,
		249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254,
		5, 36, 0, 0, 254, 256, 1, 0, 0, 0, 255, 238, 1, 0, 0, 0, 255, 246, 1, 0,
		0, 0, 256, 31, 1, 0, 0, 0, 257, 266, 3, 44, 22, 0, 258, 266, 3, 74, 37,
		0, 259, 266, 3, 48, 24, 0, 260, 266, 3, 54, 27, 0, 261, 266, 5, 34, 0,
		0, 262, 266, 5, 20, 0, 0, 263, 266, 3, 68, 34, 0, 264, 266, 3, 72, 36,
		0, 265, 257, 1, 0, 0, 0, 265, 258, 1, 0, 0, 0, 265, 259, 1, 0, 0, 0, 265,
		260, 1, 0, 0, 0, 265, 261, 1, 0, 0, 0, 265, 262, 1, 0, 0, 0, 265, 263,
		1, 0, 0, 0, 265, 264, 1, 0, 0, 0, 266, 33, 1, 0, 0, 0, 267, 268, 5, 67,
		0, 0, 268, 272, 5, 35, 0, 0, 269, 271, 3, 36, 18, 0, 270, 269, 1, 0, 0,
		0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273,
		275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 276, 5, 36, 0, 0, 276, 35,
		1, 0, 0, 0, 277, 286, 3, 44, 22, 0, 278, 286, 3, 74, 37, 0, 279, 286, 3,
		48, 24, 0, 280, 286, 3, 54, 27, 0, 281, 286, 5, 34, 0, 0, 282, 286, 5,
		20, 0, 0, 283, 286, 3, 68, 34, 0, 284, 286, 3, 72, 36, 0, 285, 277, 1,
		0, 0, 0, 285, 278, 1, 0, 0, 0, 285, 279, 1, 0, 0, 0, 285, 280, 1, 0, 0,
		0, 285, 281, 1, 0, 0, 0, 285, 282, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285,
		284, 1, 0, 0, 0, 286, 37, 1, 0, 0, 0, 287, 288, 3, 32, 16, 0, 288, 39,
		1, 0, 0, 0, 289, 290, 3, 32, 16, 0, 290, 41, 1, 0, 0, 0, 291, 292, 3, 32,
		16, 0, 292, 43, 1, 0, 0, 0, 293, 294, 5, 18, 0, 0, 294, 295, 5, 37, 0,
		0, 295, 296, 3, 46, 23, 0, 296, 297, 5, 38, 0, 0, 297, 301, 5, 35, 0, 0,
		298, 300, 3, 38, 19, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301,
		299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301,
		1, 0, 0, 0, 304, 314, 5, 36, 0, 0, 305, 306, 5, 19, 0, 0, 306, 310, 5,
		35, 0, 0, 307, 309, 3, 40, 20, 0, 308, 307, 1, 0, 0, 0, 309, 312, 1, 0,
		0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313, 1, 0, 0, 0,
		312, 310, 1, 0, 0, 0, 313, 315, 5, 36, 0, 0, 314, 305, 1, 0, 0, 0, 314,
		315, 1, 0, 0, 0, 315, 45, 1, 0, 0, 0, 316, 317, 3, 68, 34, 0, 317, 318,
		7, 0, 0, 0, 318, 319, 3, 68, 34, 0, 319, 322, 1, 0, 0, 0, 320, 322, 3,
		68, 34, 0, 321, 316, 1, 0, 0, 0, 321, 320, 1, 0, 0, 0, 322, 47, 1, 0, 0,
		0, 323, 324, 5, 21, 0, 0, 324, 325, 3, 50, 25, 0, 325, 326, 5, 72, 0, 0,
		326, 327, 3, 46, 23, 0, 327, 328, 5, 72, 0, 0, 328, 329, 3, 52, 26, 0,
		329, 333, 5, 35, 0, 0, 330, 332, 3, 42, 21, 0, 331, 330, 1, 0, 0, 0, 332,
		335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 336,
		1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337, 5, 36, 0, 0, 337, 369, 1, 0,
		0, 0, 338, 339, 5, 21, 0, 0, 339, 340, 7, 1, 0, 0, 340, 341, 5, 41, 0,
		0, 341, 342, 5, 67, 0, 0, 342, 343, 5, 62, 0, 0, 343, 344, 5, 22, 0, 0,
		344, 345, 3, 68, 34, 0, 345, 349, 5, 35, 0, 0, 346, 348, 3, 42, 21, 0,
		347, 346, 1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349,
		350, 1, 0, 0, 0, 350, 352, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 352, 353,
		5, 36, 0, 0, 353, 369, 1, 0, 0, 0, 354, 355, 5, 21, 0, 0, 355, 356, 5,
		67, 0, 0, 356, 357, 5, 62, 0, 0, 357, 358, 5, 22, 0, 0, 358, 359, 3, 68,
		34, 0, 359, 363, 5, 35, 0, 0, 360, 362, 3, 42, 21, 0, 361, 360, 1, 0, 0,
		0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364,
		366, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 367, 5, 36, 0, 0, 367, 369,
		1, 0, 0, 0, 368, 323, 1, 0, 0, 0, 368, 338, 1, 0, 0, 0, 368, 354, 1, 0,
		0, 0, 369, 49, 1, 0, 0, 0, 370, 371, 3, 74, 37, 0, 371, 372, 5, 40, 0,
		0, 372, 373, 3, 68, 34, 0, 373, 51, 1, 0, 0, 0, 374, 387, 3, 74, 37, 0,
		375, 388, 5, 50, 0, 0, 376, 388, 5, 51, 0, 0, 377, 378, 5, 57, 0, 0, 378,
		388, 3, 68, 34, 0, 379, 380, 5, 58, 0, 0, 380, 388, 3, 68, 34, 0, 381,
		382, 5, 59, 0, 0, 382, 388, 3, 68, 34, 0, 383, 384, 5, 60, 0, 0, 384, 388,
		3, 68, 34, 0, 385, 386, 5, 61, 0, 0, 386, 388, 3, 68, 34, 0, 387, 375,
		1, 0, 0, 0, 387, 376, 1, 0, 0, 0, 387, 377, 1, 0, 0, 0, 387, 379, 1, 0,
		0, 0, 387, 381, 1, 0, 0, 0, 387, 383, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0,
		388, 53, 1, 0, 0, 0, 389, 390, 5, 23, 0, 0, 390, 391, 3, 46, 23, 0, 391,
		395, 5, 35, 0, 0, 392, 394, 3, 56, 28, 0, 393, 392, 1, 0, 0, 0, 394, 397,
		1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 399, 1, 0,
		0, 0, 397, 395, 1, 0, 0, 0, 398, 400, 3, 58, 29, 0, 399, 398, 1, 0, 0,
		0, 399, 400, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 5, 36, 0, 0, 402,
		55, 1, 0, 0, 0, 403, 404, 5, 25, 0, 0, 404, 405, 3, 46, 23, 0, 405, 409,
		5, 39, 0, 0, 406, 408, 3, 32, 16, 0, 407, 406, 1, 0, 0, 0, 408, 411, 1,
		0, 0, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 57, 1, 0, 0,
		0, 411, 409, 1, 0, 0, 0, 412, 413, 5, 24, 0, 0, 413, 417, 5, 39, 0, 0,
		414, 416, 3, 32, 16, 0, 415, 414, 1, 0, 0, 0, 416, 419, 1, 0, 0, 0, 417,
		415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 59, 1, 0, 0, 0, 419, 417, 1,
		0, 0, 0, 420, 421, 5, 25, 0, 0, 421, 422, 3, 64, 32, 0, 422, 426, 5, 39,
		0, 0, 423, 425, 3, 32, 16, 0, 424, 423, 1, 0, 0, 0, 425, 428, 1, 0, 0,
		0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 61, 1, 0, 0, 0, 428,
		426, 1, 0, 0, 0, 429, 430, 5, 24, 0, 0, 430, 434, 5, 39, 0, 0, 431, 433,
		3, 32, 16, 0, 432, 431, 1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434, 432, 1,
		0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 63, 1, 0, 0, 0, 436, 434, 1, 0, 0,
		0, 437, 439, 5, 53, 0, 0, 438, 437, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439,
		440, 1, 0, 0, 0, 440, 441, 5, 67, 0, 0, 441, 65, 1, 0, 0, 0, 442, 443,
		3, 74, 37, 0, 443, 452, 5, 37, 0, 0, 444, 449, 3, 68, 34, 0, 445, 446,
		5, 41, 0, 0, 446, 448, 3, 68, 34, 0, 447, 445, 1, 0, 0, 0, 448, 451, 1,
		0, 0, 0, 449, 447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 453, 1, 0, 0,
		0, 451, 449, 1, 0, 0, 0, 452, 444, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453,
		454, 1, 0, 0, 0, 454, 455, 5, 38, 0, 0, 455, 67, 1, 0, 0, 0, 456, 457,
		6, 34, -1, 0, 457, 466, 3, 74, 37, 0, 458, 466, 5, 68, 0, 0, 459, 466,
		5, 69, 0, 0, 460, 466, 5, 70, 0, 0, 461, 462, 5, 37, 0, 0, 462, 463, 3,
		68, 34, 0, 463, 464, 5, 38, 0, 0, 464, 466, 1, 0, 0, 0, 465, 456, 1, 0,
		0, 0, 465, 458, 1, 0, 0, 0, 465, 459, 1, 0, 0, 0, 465, 460, 1, 0, 0, 0,
		465, 461, 1, 0, 0, 0, 466, 472, 1, 0, 0, 0, 467, 468, 10, 6, 0, 0, 468,
		469, 7, 2, 0, 0, 469, 471, 3, 68, 34, 7, 470, 467, 1, 0, 0, 0, 471, 474,
		1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 69, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 475, 476, 5, 67, 0, 0, 476, 477, 5, 39, 0,
		0, 477, 480, 3, 94, 47, 0, 478, 479, 5, 40, 0, 0, 479, 481, 3, 96, 48,
		0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 485, 1, 0, 0, 0, 482,
		484, 3, 78, 39, 0, 483, 482, 1, 0, 0, 0, 484, 487, 1, 0, 0, 0, 485, 483,
		1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 489, 1, 0, 0, 0, 487, 485, 1, 0,
		0, 0, 488, 490, 5, 72, 0, 0, 489, 488, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0,
		490, 71, 1, 0, 0, 0, 491, 496, 5, 74, 0, 0, 492, 496, 5, 68, 0, 0, 493,
		496, 5, 76, 0, 0, 494, 496, 3, 74, 37, 0, 495, 491, 1, 0, 0, 0, 495, 492,
		1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 494, 1, 0, 0, 0, 496, 73, 1, 0,
		0, 0, 497, 506, 5, 67, 0, 0, 498, 499, 5, 42, 0, 0, 499, 505, 5, 67, 0,
		0, 500, 501, 5, 48, 0, 0, 501, 502, 3, 68, 34, 0, 502, 503, 5, 49, 0, 0,
		503, 505, 1, 0, 0, 0, 504, 498, 1, 0, 0, 0, 504, 500, 1, 0, 0, 0, 505,
		508, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 522,
		1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 509, 518, 7, 3, 0, 0, 510, 511, 5, 42,
		0, 0, 511, 517, 5, 67, 0, 0, 512, 513, 5, 48, 0, 0, 513, 514, 3, 68, 34,
		0, 514, 515, 5, 49, 0, 0, 515, 517, 1, 0, 0, 0, 516, 510, 1, 0, 0, 0, 516,
		512, 1, 0, 0, 0, 517, 520, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 518, 519,
		1, 0, 0, 0, 519, 522, 1, 0, 0, 0, 520, 518, 1, 0, 0, 0, 521, 497, 1, 0,
		0, 0, 521, 509, 1, 0, 0, 0, 522, 75, 1, 0, 0, 0, 523, 524, 5, 67, 0, 0,
		524, 525, 5, 13, 0, 0, 525, 527, 5, 35, 0, 0, 526, 528, 3, 70, 35, 0, 527,
		526, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 527, 1, 0, 0, 0, 529, 530,
		1, 0, 0, 0, 530, 531, 1, 0, 0, 0, 531, 532, 5, 36, 0, 0, 532, 77, 1, 0,
		0, 0, 533, 534, 5, 45, 0, 0, 534, 540, 5, 67, 0, 0, 535, 537, 5, 37, 0,
		0, 536, 538, 3, 80, 40, 0, 537, 536, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0,
		538, 539, 1, 0, 0, 0, 539, 541, 5, 38, 0, 0, 540, 535, 1, 0, 0, 0, 540,
		541, 1, 0, 0, 0, 541, 79, 1, 0, 0, 0, 542, 547, 3, 82, 41, 0, 543, 544,
		5, 41, 0, 0, 544, 546, 3, 82, 41, 0, 545, 543, 1, 0, 0, 0, 546, 549, 1,
		0, 0, 0, 547, 545, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 81, 1, 0, 0,
		0, 549, 547, 1, 0, 0, 0, 550, 553, 5, 68, 0, 0, 551, 553, 3, 84, 42, 0,
		552, 550, 1, 0, 0, 0, 552, 551, 1, 0, 0, 0, 553, 83, 1, 0, 0, 0, 554, 563,
		5, 35, 0, 0, 555, 560, 5, 68, 0, 0, 556, 557, 5, 41, 0, 0, 557, 559, 5,
		68, 0, 0, 558, 556, 1, 0, 0, 0, 559, 562, 1, 0, 0, 0, 560, 558, 1, 0, 0,
		0, 560, 561, 1, 0, 0, 0, 561, 564, 1, 0, 0, 0, 562, 560, 1, 0, 0, 0, 563,
		555, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 565, 1, 0, 0, 0, 565, 566,
		5, 36, 0, 0, 566, 85, 1, 0, 0, 0, 567, 568, 5, 45, 0, 0, 568, 574, 5, 67,
		0, 0, 569, 571, 5, 37, 0, 0, 570, 572, 3, 80, 40, 0, 571, 570, 1, 0, 0,
		0, 571, 572, 1, 0, 0, 0, 572, 573, 1, 0, 0, 0, 573, 575, 5, 38, 0, 0, 574,
		569, 1, 0, 0, 0, 574, 575, 1, 0, 0, 0, 575, 87, 1, 0, 0, 0, 576, 577, 5,
		32, 0, 0, 577, 578, 3, 92, 46, 0, 578, 579, 5, 36, 0, 0, 579, 89, 1, 0,
		0, 0, 580, 581, 5, 33, 0, 0, 581, 582, 3, 92, 46, 0, 582, 583, 5, 36, 0,
		0, 583, 91, 1, 0, 0, 0, 584, 592, 5, 80, 0, 0, 585, 592, 5, 78, 0, 0, 586,
		592, 5, 79, 0, 0, 587, 588, 5, 35, 0, 0, 588, 589, 3, 92, 46, 0, 589, 590,
		5, 36, 0, 0, 590, 592, 1, 0, 0, 0, 591, 584, 1, 0, 0, 0, 591, 585, 1, 0,
		0, 0, 591, 586, 1, 0, 0, 0, 591, 587, 1, 0, 0, 0, 592, 595, 1, 0, 0, 0,
		593, 591, 1, 0, 0, 0, 593, 594, 1, 0, 0, 0, 594, 93, 1, 0, 0, 0, 595, 593,
		1, 0, 0, 0, 596, 614, 1, 0, 0, 0, 597, 598, 5, 13, 0, 0, 598, 602, 5, 35,
		0, 0, 599, 601, 3, 70, 35, 0, 600, 599, 1, 0, 0, 0, 601, 604, 1, 0, 0,
		0, 602, 600, 1, 0, 0, 0, 602, 603, 1, 0, 0, 0, 603, 605, 1, 0, 0, 0, 604,
		602, 1, 0, 0, 0, 605, 614, 5, 36, 0, 0, 606, 614, 5, 2, 0, 0, 607, 614,
		5, 3, 0, 0, 608, 609, 5, 48, 0, 0, 609, 610, 5, 49, 0, 0, 610, 614, 3,
		94, 47, 0, 611, 614, 5, 1, 0, 0, 612, 614, 5, 67, 0, 0, 613, 596, 1, 0,
		0, 0, 613, 597, 1, 0, 0, 0, 613, 606, 1, 0, 0, 0, 613, 607, 1, 0, 0, 0,
		613, 608, 1, 0, 0, 0, 613, 611, 1, 0, 0, 0, 613, 612, 1, 0, 0, 0, 614,
		95, 1, 0, 0, 0, 615, 616, 7, 4, 0, 0, 616, 97, 1, 0, 0, 0, 617, 618, 7,
		5, 0, 0, 618, 99, 1, 0, 0, 0, 64, 103, 109, 123, 138, 145, 151, 160, 166,
		179, 186, 192, 198, 214, 222, 226, 236, 243, 251, 255, 265, 272, 285, 301,
		310, 314, 321, 333, 349, 363, 368, 387, 395, 399, 409, 417, 426, 434, 438,
		449, 452, 465, 472, 480, 485, 489, 495, 504, 506, 516, 518, 521, 529, 537,
		540, 547, 552, 560, 563, 571, 574, 591, 593, 602, 613,
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
	PromptDSLParserRUNEXE             = 27
	PromptDSLParserIT                 = 28
	PromptDSLParserIT_CTX             = 29
	PromptDSLParserLOOPRANGE          = 30
	PromptDSLParserBEFORE             = 31
	PromptDSLParserFIX                = 32
	PromptDSLParserAFTER              = 33
	PromptDSLParserARRAY_OUTPUTSPEC   = 34
	PromptDSLParserLBRACE             = 35
	PromptDSLParserRBRACE             = 36
	PromptDSLParserLPAREN             = 37
	PromptDSLParserRPAREN             = 38
	PromptDSLParserCOLON              = 39
	PromptDSLParserEQUAL              = 40
	PromptDSLParserCOMMA              = 41
	PromptDSLParserDOT                = 42
	PromptDSLParserEQEQ               = 43
	PromptDSLParserNOTEQ              = 44
	PromptDSLParserAT                 = 45
	PromptDSLParserMD                 = 46
	PromptDSLParserJSON               = 47
	PromptDSLParserLBRACK             = 48
	PromptDSLParserRBRACK             = 49
	PromptDSLParserINCREMENT          = 50
	PromptDSLParserDECREMENT          = 51
	PromptDSLParserMINUS              = 52
	PromptDSLParserSTAR               = 53
	PromptDSLParserSLASH              = 54
	PromptDSLParserMOD                = 55
	PromptDSLParserUNDERSCORE         = 56
	PromptDSLParserPLUSEQ             = 57
	PromptDSLParserMINUSEQ            = 58
	PromptDSLParserMULTEQ             = 59
	PromptDSLParserDIVEQ              = 60
	PromptDSLParserMODEQ              = 61
	PromptDSLParserDECL_ASSIGN        = 62
	PromptDSLParserLT                 = 63
	PromptDSLParserLTE                = 64
	PromptDSLParserGT                 = 65
	PromptDSLParserGTE                = 66
	PromptDSLParserID                 = 67
	PromptDSLParserSTRING             = 68
	PromptDSLParserNUMBER             = 69
	PromptDSLParserBOOL               = 70
	PromptDSLParserPIPE               = 71
	PromptDSLParserSEMI               = 72
	PromptDSLParserPLUS               = 73
	PromptDSLParserDASH_STRING        = 74
	PromptDSLParserWS                 = 75
	PromptDSLParserLINE_COMMENT       = 76
	PromptDSLParserBLOCK_COMMENT      = 77
	PromptDSLParserCODE_STRING_SINGLE = 78
	PromptDSLParserCODE_STRING        = 79
	PromptDSLParserCODE_TEXT          = 80
)

// PromptDSLParser rules.
const (
	PromptDSLParserRULE_promptFile        = 0
	PromptDSLParserRULE_runexeDef         = 1
	PromptDSLParserRULE_promptDef         = 2
	PromptDSLParserRULE_promptBlock       = 3
	PromptDSLParserRULE_goimportSection   = 4
	PromptDSLParserRULE_goimportEntry     = 5
	PromptDSLParserRULE_inputSection      = 6
	PromptDSLParserRULE_iterater          = 7
	PromptDSLParserRULE_loopRangeSection  = 8
	PromptDSLParserRULE_outputSection     = 9
	PromptDSLParserRULE_outputStruct      = 10
	PromptDSLParserRULE_outputMarkdown    = 11
	PromptDSLParserRULE_beforeSection     = 12
	PromptDSLParserRULE_systemSection     = 13
	PromptDSLParserRULE_sysContent        = 14
	PromptDSLParserRULE_userSection       = 15
	PromptDSLParserRULE_userContent       = 16
	PromptDSLParserRULE_moduleDef         = 17
	PromptDSLParserRULE_moduleContent     = 18
	PromptDSLParserRULE_thencontent       = 19
	PromptDSLParserRULE_elsecontent       = 20
	PromptDSLParserRULE_forcontent        = 21
	PromptDSLParserRULE_ifStatement       = 22
	PromptDSLParserRULE_condition         = 23
	PromptDSLParserRULE_forStatement      = 24
	PromptDSLParserRULE_assignExpr        = 25
	PromptDSLParserRULE_updateExpr        = 26
	PromptDSLParserRULE_switchStatement   = 27
	PromptDSLParserRULE_switchCase        = 28
	PromptDSLParserRULE_switchDefault     = 29
	PromptDSLParserRULE_typeCase          = 30
	PromptDSLParserRULE_typeDefault       = 31
	PromptDSLParserRULE_typeName          = 32
	PromptDSLParserRULE_dslCallExpr       = 33
	PromptDSLParserRULE_expr              = 34
	PromptDSLParserRULE_fieldDef          = 35
	PromptDSLParserRULE_textLine          = 36
	PromptDSLParserRULE_paramPath         = 37
	PromptDSLParserRULE_structDef         = 38
	PromptDSLParserRULE_annotation        = 39
	PromptDSLParserRULE_annotationArgs    = 40
	PromptDSLParserRULE_annotationValue   = 41
	PromptDSLParserRULE_arrayLiteral      = 42
	PromptDSLParserRULE_defaultAnnotation = 43
	PromptDSLParserRULE_fixSection        = 44
	PromptDSLParserRULE_afterSection      = 45
	PromptDSLParserRULE_codeBlockContent  = 46
	PromptDSLParserRULE_type              = 47
	PromptDSLParserRULE_value             = 48
	PromptDSLParserRULE_formatType        = 49
)

// IPromptFileContext is an interface to support dynamic dispatch.
type IPromptFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllRunexeDef() []IRunexeDefContext
	RunexeDef(i int) IRunexeDefContext
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

func (s *PromptFileContext) AllRunexeDef() []IRunexeDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunexeDefContext); ok {
			len++
		}
	}

	tst := make([]IRunexeDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunexeDefContext); ok {
			tst[i] = t.(IRunexeDefContext)
			i++
		}
	}

	return tst
}

func (s *PromptFileContext) RunexeDef(i int) IRunexeDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunexeDefContext); ok {
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

	return t.(IRunexeDefContext)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserRUNEXE {
		{
			p.SetState(100)
			p.RunexeDef()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserPROMPT {
		{
			p.SetState(106)
			p.PromptDef()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(111)
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

// IRunexeDefContext is an interface to support dynamic dispatch.
type IRunexeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RUNEXE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsRunexeDefContext differentiates from other interfaces.
	IsRunexeDefContext()
}

type RunexeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunexeDefContext() *RunexeDefContext {
	var p = new(RunexeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_runexeDef
	return p
}

func InitEmptyRunexeDefContext(p *RunexeDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_runexeDef
}

func (*RunexeDefContext) IsRunexeDefContext() {}

func NewRunexeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunexeDefContext {
	var p = new(RunexeDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_runexeDef

	return p
}

func (s *RunexeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *RunexeDefContext) RUNEXE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRUNEXE, 0)
}

func (s *RunexeDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOLON, 0)
}

func (s *RunexeDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *RunexeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunexeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunexeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterRunexeDef(s)
	}
}

func (s *RunexeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitRunexeDef(s)
	}
}

func (p *PromptDSLParser) RunexeDef() (localctx IRunexeDefContext) {
	localctx = NewRunexeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromptDSLParserRULE_runexeDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(PromptDSLParserRUNEXE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(PromptDSLParserSTRING)
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
	p.EnterRule(localctx, 4, PromptDSLParserRULE_promptDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(PromptDSLParserPROMPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-6)) & ^0x3f) == 0 && ((int64(1)<<(_la-6))&2305843559230603291) != 0) {
		{
			p.SetState(120)
			p.PromptBlock()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
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
	Iterater() IIteraterContext
	LoopRangeSection() ILoopRangeSectionContext
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

func (s *PromptBlockContext) Iterater() IIteraterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteraterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteraterContext)
}

func (s *PromptBlockContext) LoopRangeSection() ILoopRangeSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopRangeSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopRangeSectionContext)
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
	p.EnterRule(localctx, 6, PromptDSLParserRULE_promptBlock)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.InputSection()
		}

	case PromptDSLParserIT_CTX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Iterater()
		}

	case PromptDSLParserLOOPRANGE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.LoopRangeSection()
		}

	case PromptDSLParserOUTPUT, PromptDSLParserAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.OutputSection()
		}

	case PromptDSLParserBEFORE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.BeforeSection()
		}

	case PromptDSLParserSYSTEM:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.SystemSection()
		}

	case PromptDSLParserUSER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.UserSection()
		}

	case PromptDSLParserAFTER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			p.AfterSection()
		}

	case PromptDSLParserFIX:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(135)
			p.FixSection()
		}

	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(136)
			p.ModuleDef()
		}

	case PromptDSLParserGOIMPORT:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(137)
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
	p.EnterRule(localctx, 8, PromptDSLParserRULE_goimportSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(PromptDSLParserGOIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserID || _la == PromptDSLParserSTRING {
		{
			p.SetState(142)
			p.GoimportEntry()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(148)
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
	p.EnterRule(localctx, 10, PromptDSLParserRULE_goimportEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserID {
		{
			p.SetState(150)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*GoimportEntryContext).alias = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(153)

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
	p.EnterRule(localctx, 12, PromptDSLParserRULE_inputSection)
	var _la int

	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(PromptDSLParserINPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(157)
				p.FieldDef()
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(162)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(PromptDSLParserINPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
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

// IIteraterContext is an interface to support dynamic dispatch.
type IIteraterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IT_CTX() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsIteraterContext differentiates from other interfaces.
	IsIteraterContext()
}

type IteraterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteraterContext() *IteraterContext {
	var p = new(IteraterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_iterater
	return p
}

func InitEmptyIteraterContext(p *IteraterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_iterater
}

func (*IteraterContext) IsIteraterContext() {}

func NewIteraterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteraterContext {
	var p = new(IteraterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_iterater

	return p
}

func (s *IteraterContext) GetParser() antlr.Parser { return s.parser }

func (s *IteraterContext) IT_CTX() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserIT_CTX, 0)
}

func (s *IteraterContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLPAREN, 0)
}

func (s *IteraterContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserSTRING, 0)
}

func (s *IteraterContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRPAREN, 0)
}

func (s *IteraterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteraterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IteraterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterIterater(s)
	}
}

func (s *IteraterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitIterater(s)
	}
}

func (p *PromptDSLParser) Iterater() (localctx IIteraterContext) {
	localctx = NewIteraterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromptDSLParserRULE_iterater)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(PromptDSLParserIT_CTX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(PromptDSLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
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

// ILoopRangeSectionContext is an interface to support dynamic dispatch.
type ILoopRangeSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOPRANGE() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsLoopRangeSectionContext differentiates from other interfaces.
	IsLoopRangeSectionContext()
}

type LoopRangeSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopRangeSectionContext() *LoopRangeSectionContext {
	var p = new(LoopRangeSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_loopRangeSection
	return p
}

func InitEmptyLoopRangeSectionContext(p *LoopRangeSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromptDSLParserRULE_loopRangeSection
}

func (*LoopRangeSectionContext) IsLoopRangeSectionContext() {}

func NewLoopRangeSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopRangeSectionContext {
	var p = new(LoopRangeSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromptDSLParserRULE_loopRangeSection

	return p
}

func (s *LoopRangeSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopRangeSectionContext) LOOPRANGE() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLOOPRANGE, 0)
}

func (s *LoopRangeSectionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserLBRACK, 0)
}

func (s *LoopRangeSectionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserRBRACK, 0)
}

func (s *LoopRangeSectionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserNUMBER)
}

func (s *LoopRangeSectionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserNUMBER, i)
}

func (s *LoopRangeSectionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCOMMA, 0)
}

func (s *LoopRangeSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopRangeSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopRangeSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.EnterLoopRangeSection(s)
	}
}

func (s *LoopRangeSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromptDSLParserListener); ok {
		listenerT.ExitLoopRangeSection(s)
	}
}

func (p *PromptDSLParser) LoopRangeSection() (localctx ILoopRangeSectionContext) {
	localctx = NewLoopRangeSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromptDSLParserRULE_loopRangeSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(PromptDSLParserLOOPRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(PromptDSLParserLBRACK)
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

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(175)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(178)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(181)
		p.Match(PromptDSLParserRBRACK)
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
	p.EnterRule(localctx, 18, PromptDSLParserRULE_outputSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserAT {
		{
			p.SetState(183)
			p.DefaultAnnotation()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Match(PromptDSLParserOUTPUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserLBRACE:
		{
			p.SetState(190)
			p.OutputStruct()
		}

	case PromptDSLParserCOLON:
		{
			p.SetState(191)
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
	p.EnterRule(localctx, 20, PromptDSLParserRULE_outputStruct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(195)
			p.FieldDef()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(200)
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
	p.EnterRule(localctx, 22, PromptDSLParserRULE_outputMarkdown)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
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
	p.EnterRule(localctx, 24, PromptDSLParserRULE_beforeSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(PromptDSLParserBEFORE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.CodeBlockContent()
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 26, PromptDSLParserRULE_systemSection)
	var _la int

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(211)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(216)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(PromptDSLParserSYSTEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(219)
				p.SysContent()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(224)
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
	SwitchStatement() ISwitchStatementContext
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

func (s *SysContentContext) SwitchStatement() ISwitchStatementContext {
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
	p.EnterRule(localctx, 28, PromptDSLParserRULE_sysContent)
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.SwitchStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.expr(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(235)
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
	p.EnterRule(localctx, 30, PromptDSLParserRULE_userSection)
	var _la int

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PromptDSLParserID {
			{
				p.SetState(240)
				p.Match(PromptDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(245)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Match(PromptDSLParserUSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(248)
				p.UserContent()
			}

			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(253)
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
	SwitchStatement() ISwitchStatementContext
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

func (s *UserContentContext) SwitchStatement() ISwitchStatementContext {
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
	p.EnterRule(localctx, 32, PromptDSLParserRULE_userContent)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(259)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(260)
			p.SwitchStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(261)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(262)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(263)
			p.expr(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(264)
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
	p.EnterRule(localctx, 34, PromptDSLParserRULE_moduleDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(269)
			p.ModuleContent()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(275)
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
	p.EnterRule(localctx, 36, PromptDSLParserRULE_moduleContent)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.IfStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.ParamPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.ForStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.SwitchStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(281)
			p.Match(PromptDSLParserARRAY_OUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(282)
			p.Match(PromptDSLParserOUTPUTSPEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(283)
			p.expr(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(284)
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
	p.EnterRule(localctx, 38, PromptDSLParserRULE_thencontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 40, PromptDSLParserRULE_elsecontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
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
	p.EnterRule(localctx, 42, PromptDSLParserRULE_forcontent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
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
	p.EnterRule(localctx, 44, PromptDSLParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(PromptDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Condition()
	}
	{
		p.SetState(296)
		p.Match(PromptDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(298)
			p.Thencontent()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(304)
		p.Match(PromptDSLParserRBRACE)
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

	if _la == PromptDSLParserELSE {
		{
			p.SetState(305)
			p.Match(PromptDSLParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(307)
				p.Elsecontent()
			}

			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(313)
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
	p.EnterRule(localctx, 46, PromptDSLParserRULE_condition)
	var _la int

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)

			var _x = p.expr(0)

			localctx.(*ConditionContext).lhs = _x
		}
		{
			p.SetState(317)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConditionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-43)) & ^0x3f) == 0 && ((int64(1)<<(_la-43))&16252931) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConditionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(318)

			var _x = p.expr(0)

			localctx.(*ConditionContext).rhs = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)

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
	p.EnterRule(localctx, 48, PromptDSLParserRULE_forStatement)
	var _la int

	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForTraditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)

			var _x = p.AssignExpr()

			localctx.(*ForTraditionalContext).init = _x
		}
		{
			p.SetState(325)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Condition()
		}
		{
			p.SetState(327)
			p.Match(PromptDSLParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)

			var _x = p.UpdateExpr()

			localctx.(*ForTraditionalContext).update = _x
		}
		{
			p.SetState(329)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(330)
				p.Forcontent()
			}

			p.SetState(335)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(336)
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
			p.SetState(338)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)

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
			p.SetState(340)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*ForRangeWithIndexContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(PromptDSLParserDECL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(PromptDSLParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)

			var _x = p.expr(0)

			localctx.(*ForRangeWithIndexContext).iterable = _x
		}
		{
			p.SetState(345)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(346)
				p.Forcontent()
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(352)
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
			p.SetState(354)
			p.Match(PromptDSLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _m = p.Match(PromptDSLParserID)

			localctx.(*ForRangeNoIndexContext).val = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(PromptDSLParserDECL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(PromptDSLParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)

			var _x = p.expr(0)

			localctx.(*ForRangeNoIndexContext).iterable = _x
		}
		{
			p.SetState(359)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
			{
				p.SetState(360)
				p.Forcontent()
			}

			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(366)
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
	p.EnterRule(localctx, 50, PromptDSLParserRULE_assignExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)

		var _x = p.ParamPath()

		localctx.(*AssignExprContext).lhs = _x
	}
	{
		p.SetState(371)
		p.Match(PromptDSLParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)

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
	p.EnterRule(localctx, 52, PromptDSLParserRULE_updateExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.ParamPath()
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINCREMENT:
		{
			p.SetState(375)
			p.Match(PromptDSLParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserDECREMENT:
		{
			p.SetState(376)
			p.Match(PromptDSLParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserPLUSEQ:
		{
			p.SetState(377)
			p.Match(PromptDSLParserPLUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.expr(0)
		}

	case PromptDSLParserMINUSEQ:
		{
			p.SetState(379)
			p.Match(PromptDSLParserMINUSEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.expr(0)
		}

	case PromptDSLParserMULTEQ:
		{
			p.SetState(381)
			p.Match(PromptDSLParserMULTEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.expr(0)
		}

	case PromptDSLParserDIVEQ:
		{
			p.SetState(383)
			p.Match(PromptDSLParserDIVEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.expr(0)
		}

	case PromptDSLParserMODEQ:
		{
			p.SetState(385)
			p.Match(PromptDSLParserMODEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
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
	p.EnterRule(localctx, 54, PromptDSLParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(PromptDSLParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Condition()
	}
	{
		p.SetState(391)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserCASE {
		{
			p.SetState(392)
			p.SwitchCase()
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserDEFAULT {
		{
			p.SetState(398)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(401)
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
	p.EnterRule(localctx, 56, PromptDSLParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(PromptDSLParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Condition()
	}
	{
		p.SetState(405)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(406)
			p.UserContent()
		}

		p.SetState(411)
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
	p.EnterRule(localctx, 58, PromptDSLParserRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(PromptDSLParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(PromptDSLParserCOLON)
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

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(414)
			p.UserContent()
		}

		p.SetState(419)
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
	p.EnterRule(localctx, 60, PromptDSLParserRULE_typeCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Match(PromptDSLParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.TypeName()
	}
	{
		p.SetState(422)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(423)
			p.UserContent()
		}

		p.SetState(428)
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
	p.EnterRule(localctx, 62, PromptDSLParserRULE_typeDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Match(PromptDSLParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(430)
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&165636474368) != 0) || ((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&655) != 0) {
		{
			p.SetState(431)
			p.UserContent()
		}

		p.SetState(436)
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
	p.EnterRule(localctx, 64, PromptDSLParserRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTAR {
		{
			p.SetState(437)
			p.Match(PromptDSLParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(440)
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
	p.EnterRule(localctx, 66, PromptDSLParserRULE_dslCallExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.ParamPath()
	}
	{
		p.SetState(443)
		p.Match(PromptDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&4323455642565607427) != 0 {
		{
			p.SetState(444)
			p.expr(0)
		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(445)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(446)
				p.expr(0)
			}

			p.SetState(451)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(454)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, PromptDSLParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserIT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		{
			p.SetState(457)
			p.ParamPath()
		}

	case PromptDSLParserSTRING:
		{
			p.SetState(458)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserNUMBER:
		{
			p.SetState(459)
			p.Match(PromptDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserBOOL:
		{
			p.SetState(460)
			p.Match(PromptDSLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLPAREN:
		{
			p.SetState(461)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.expr(0)
		}
		{
			p.SetState(463)
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
	p.SetState(472)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PromptDSLParserRULE_expr)
			p.SetState(467)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				goto errorExit
			}
			{
				p.SetState(468)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&2097167) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(469)
				p.expr(7)
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, PromptDSLParserRULE_fieldDef)
	var _la int

	var _alt int

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
		p.Match(PromptDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
		p.Type_()
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserEQUAL {
		{
			p.SetState(478)
			p.Match(PromptDSLParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.Value()
		}

	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(482)
				p.Annotation()
			}

		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSEMI {
		{
			p.SetState(488)
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
	p.EnterRule(localctx, 72, PromptDSLParserRULE_textLine)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserDASH_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Match(PromptDSLParserDASH_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(493)
			p.Match(PromptDSLParserLINE_COMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserIT, PromptDSLParserBEFORE, PromptDSLParserAFTER, PromptDSLParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(494)
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
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	IT() antlr.TerminalNode
	INPUT() antlr.TerminalNode
	OUTPUT() antlr.TerminalNode
	AFTER() antlr.TerminalNode
	BEFORE() antlr.TerminalNode

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

func (s *ParamPathContext) IT() antlr.TerminalNode {
	return s.GetToken(PromptDSLParserIT, 0)
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
	p.EnterRule(localctx, 74, PromptDSLParserRULE_paramPath)
	var _la int

	var _alt int

	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)
			p.Match(PromptDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(504)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case PromptDSLParserDOT:
					{
						p.SetState(498)
						p.Match(PromptDSLParserDOT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(499)
						p.Match(PromptDSLParserID)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case PromptDSLParserLBRACK:
					{
						p.SetState(500)
						p.Match(PromptDSLParserLBRACK)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(501)
						p.expr(0)
					}
					{
						p.SetState(502)
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
			p.SetState(508)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case PromptDSLParserINPUT, PromptDSLParserOUTPUT, PromptDSLParserIT, PromptDSLParserBEFORE, PromptDSLParserAFTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(509)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&11005855232) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(516)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case PromptDSLParserDOT:
					{
						p.SetState(510)
						p.Match(PromptDSLParserDOT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(511)
						p.Match(PromptDSLParserID)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case PromptDSLParserLBRACK:
					{
						p.SetState(512)
						p.Match(PromptDSLParserLBRACK)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(513)
						p.expr(0)
					}
					{
						p.SetState(514)
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
			p.SetState(520)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
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
	p.EnterRule(localctx, 76, PromptDSLParserRULE_structDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(524)
		p.Match(PromptDSLParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromptDSLParserID {
		{
			p.SetState(526)
			p.FieldDef()
		}

		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 78, PromptDSLParserRULE_annotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(534)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(535)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(536)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(539)
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
	p.EnterRule(localctx, 80, PromptDSLParserRULE_annotationArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		p.AnnotationValue()
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromptDSLParserCOMMA {
		{
			p.SetState(543)
			p.Match(PromptDSLParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.AnnotationValue()
		}

		p.SetState(549)
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
	p.EnterRule(localctx, 82, PromptDSLParserRULE_annotationValue)
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromptDSLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromptDSLParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(551)
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
	p.EnterRule(localctx, 84, PromptDSLParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(554)
		p.Match(PromptDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserSTRING {
		{
			p.SetState(555)
			p.Match(PromptDSLParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserCOMMA {
			{
				p.SetState(556)
				p.Match(PromptDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(557)
				p.Match(PromptDSLParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(562)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(565)
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
	p.EnterRule(localctx, 86, PromptDSLParserRULE_defaultAnnotation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		p.Match(PromptDSLParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(568)
		p.Match(PromptDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromptDSLParserLPAREN {
		{
			p.SetState(569)
			p.Match(PromptDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromptDSLParserLBRACE || _la == PromptDSLParserSTRING {
			{
				p.SetState(570)
				p.AnnotationArgs()
			}

		}
		{
			p.SetState(573)
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
	p.EnterRule(localctx, 88, PromptDSLParserRULE_fixSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.Match(PromptDSLParserFIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(577)
		p.CodeBlockContent()
	}
	{
		p.SetState(578)
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
	p.EnterRule(localctx, 90, PromptDSLParserRULE_afterSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Match(PromptDSLParserAFTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(581)
		p.CodeBlockContent()
	}
	{
		p.SetState(582)
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
	AllCODE_STRING_SINGLE() []antlr.TerminalNode
	CODE_STRING_SINGLE(i int) antlr.TerminalNode
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

func (s *CodeBlockContentContext) AllCODE_STRING_SINGLE() []antlr.TerminalNode {
	return s.GetTokens(PromptDSLParserCODE_STRING_SINGLE)
}

func (s *CodeBlockContentContext) CODE_STRING_SINGLE(i int) antlr.TerminalNode {
	return s.GetToken(PromptDSLParserCODE_STRING_SINGLE, i)
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
	p.EnterRule(localctx, 92, PromptDSLParserRULE_codeBlockContent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&61572651155457) != 0 {
		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromptDSLParserCODE_TEXT:
			{
				p.SetState(584)
				p.Match(PromptDSLParserCODE_TEXT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserCODE_STRING_SINGLE:
			{
				p.SetState(585)
				p.Match(PromptDSLParserCODE_STRING_SINGLE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserCODE_STRING:
			{
				p.SetState(586)
				p.Match(PromptDSLParserCODE_STRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case PromptDSLParserLBRACE:
			{
				p.SetState(587)
				p.Match(PromptDSLParserLBRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(588)
				p.CodeBlockContent()
			}
			{
				p.SetState(589)
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

		p.SetState(595)
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
	p.EnterRule(localctx, 94, PromptDSLParserRULE_type)
	var _la int

	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(597)
			p.Match(PromptDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)
			p.Match(PromptDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromptDSLParserID {
			{
				p.SetState(599)
				p.FieldDef()
			}

			p.SetState(604)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(605)
			p.Match(PromptDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(606)
			p.Match(PromptDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(607)
			p.Match(PromptDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(608)
			p.Match(PromptDSLParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Match(PromptDSLParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.Type_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(611)
			p.Match(PromptDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(612)
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
	p.EnterRule(localctx, 96, PromptDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(615)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&71) != 0) {
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
	p.EnterRule(localctx, 98, PromptDSLParserRULE_formatType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
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
	case 34:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
