// Code generated from ./codegen/grammar/W/WorkflowDSLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package Wparser // WorkflowDSLParser
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

type WorkflowDSLParser struct {
	*antlr.BaseParser
}

var WorkflowDSLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workflowdslparserParserInit() {
	staticData := &WorkflowDSLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'workflow'", "'out'", "'models'", "'if'", "'else'", "'string'",
		"'int'", "'float'", "'auto'", "'struct'", "'{'", "'}'", "'['", "']'",
		"'('", "')'", "':'", "'<-'", "','", "'.'", "'/'", "'+'", "'-'", "'*'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'$'", "'&'", "'|'", "'%'",
		"'@'", "'#'", "'^'", "'~'", "'&&'", "'||'", "'!'", "'sum'",
	}
	staticData.SymbolicNames = []string{
		"", "WORKFLOW", "OUT", "MODELS", "IF", "ELSE", "STRING_TYPE", "INT_TYPE",
		"FLOAT_TYPE", "AUTO", "STRUCT", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"LPAREN", "RPAREN", "COLON", "ARROW", "COMMA", "DOT", "SLASH", "PLUS",
		"MINUS", "MULT", "EQ", "NEQ", "LT", "LTE", "GT", "GTE", "DOLLAR", "AMP",
		"PIPE", "PERCENT", "AT", "HASH", "CARET", "TILDE", "AND", "OR", "NOT",
		"SUM", "IN_ID", "IN_ID_ARRAY", "ID", "INT", "FLOAT", "STRING", "WS",
		"COMMENT",
	}
	staticData.RuleNames = []string{
		"workflow", "workflowBody", "inputs", "inputBlock", "fieldDef", "type",
		"outtype", "out", "expr", "atom", "functionCall", "conditionalExpr",
		"condition", "comparator", "arithmeticExpr", "parenExpr", "arithExprContent",
		"models", "modelEntry", "outField",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 232, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 51, 8, 2, 11, 2, 12,
		2, 52, 1, 3, 1, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 76, 8, 5, 10, 5, 12, 5, 79, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		85, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 93, 8, 6, 10, 6, 12,
		6, 96, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 103, 8, 6, 1, 7, 1, 7,
		1, 7, 5, 7, 108, 8, 7, 10, 7, 12, 7, 111, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 121, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 126, 8,
		9, 10, 9, 12, 9, 129, 9, 9, 1, 9, 1, 9, 1, 9, 5, 9, 134, 8, 9, 10, 9, 12,
		9, 137, 9, 9, 3, 9, 139, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		146, 8, 10, 10, 10, 12, 10, 149, 9, 10, 3, 10, 151, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 174,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 183, 8,
		12, 1, 12, 1, 12, 3, 12, 187, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 209, 8, 16, 1, 17, 1, 17, 1, 17,
		5, 17, 214, 8, 17, 10, 17, 12, 17, 217, 9, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 230, 8, 19,
		1, 19, 0, 0, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 0, 4, 1, 0, 43, 44, 2, 0, 43, 43, 45, 45, 2, 0, 25,
		30, 39, 40, 2, 0, 9, 9, 46, 46, 242, 0, 40, 1, 0, 0, 0, 2, 45, 1, 0, 0,
		0, 4, 50, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 84, 1,
		0, 0, 0, 12, 102, 1, 0, 0, 0, 14, 104, 1, 0, 0, 0, 16, 120, 1, 0, 0, 0,
		18, 138, 1, 0, 0, 0, 20, 140, 1, 0, 0, 0, 22, 154, 1, 0, 0, 0, 24, 186,
		1, 0, 0, 0, 26, 188, 1, 0, 0, 0, 28, 190, 1, 0, 0, 0, 30, 196, 1, 0, 0,
		0, 32, 208, 1, 0, 0, 0, 34, 210, 1, 0, 0, 0, 36, 220, 1, 0, 0, 0, 38, 224,
		1, 0, 0, 0, 40, 41, 5, 1, 0, 0, 41, 42, 5, 11, 0, 0, 42, 43, 3, 2, 1, 0,
		43, 44, 5, 12, 0, 0, 44, 1, 1, 0, 0, 0, 45, 46, 3, 4, 2, 0, 46, 47, 3,
		14, 7, 0, 47, 48, 3, 34, 17, 0, 48, 3, 1, 0, 0, 0, 49, 51, 3, 6, 3, 0,
		50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1,
		0, 0, 0, 53, 5, 1, 0, 0, 0, 54, 55, 7, 0, 0, 0, 55, 59, 5, 11, 0, 0, 56,
		58, 3, 8, 4, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0,
		0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63,
		5, 12, 0, 0, 63, 7, 1, 0, 0, 0, 64, 65, 5, 45, 0, 0, 65, 67, 5, 17, 0,
		0, 66, 68, 3, 10, 5, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 9,
		1, 0, 0, 0, 69, 85, 5, 7, 0, 0, 70, 85, 5, 8, 0, 0, 71, 85, 5, 6, 0, 0,
		72, 73, 5, 10, 0, 0, 73, 77, 5, 11, 0, 0, 74, 76, 3, 8, 4, 0, 75, 74, 1,
		0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 85, 5, 12, 0, 0, 81, 82, 5, 13,
		0, 0, 82, 83, 5, 14, 0, 0, 83, 85, 3, 10, 5, 0, 84, 69, 1, 0, 0, 0, 84,
		70, 1, 0, 0, 0, 84, 71, 1, 0, 0, 0, 84, 72, 1, 0, 0, 0, 84, 81, 1, 0, 0,
		0, 85, 11, 1, 0, 0, 0, 86, 103, 5, 7, 0, 0, 87, 103, 5, 8, 0, 0, 88, 103,
		5, 6, 0, 0, 89, 90, 5, 10, 0, 0, 90, 94, 5, 11, 0, 0, 91, 93, 3, 38, 19,
		0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 103, 5, 12, 0,
		0, 98, 99, 5, 13, 0, 0, 99, 100, 5, 14, 0, 0, 100, 103, 3, 12, 6, 0, 101,
		103, 1, 0, 0, 0, 102, 86, 1, 0, 0, 0, 102, 87, 1, 0, 0, 0, 102, 88, 1,
		0, 0, 0, 102, 89, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0,
		103, 13, 1, 0, 0, 0, 104, 105, 5, 2, 0, 0, 105, 109, 5, 11, 0, 0, 106,
		108, 3, 38, 19, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0,
		0, 0, 112, 113, 5, 12, 0, 0, 113, 15, 1, 0, 0, 0, 114, 121, 3, 20, 10,
		0, 115, 121, 3, 22, 11, 0, 116, 121, 3, 28, 14, 0, 117, 121, 3, 18, 9,
		0, 118, 121, 5, 46, 0, 0, 119, 121, 3, 30, 15, 0, 120, 114, 1, 0, 0, 0,
		120, 115, 1, 0, 0, 0, 120, 116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120,
		118, 1, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 17, 1, 0, 0, 0, 122, 127, 3,
		20, 10, 0, 123, 124, 5, 20, 0, 0, 124, 126, 5, 45, 0, 0, 125, 123, 1, 0,
		0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0,
		128, 139, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 135, 7, 1, 0, 0, 131,
		132, 5, 20, 0, 0, 132, 134, 5, 45, 0, 0, 133, 131, 1, 0, 0, 0, 134, 137,
		1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 139, 1, 0,
		0, 0, 137, 135, 1, 0, 0, 0, 138, 122, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0,
		139, 19, 1, 0, 0, 0, 140, 141, 5, 45, 0, 0, 141, 150, 5, 15, 0, 0, 142,
		147, 3, 16, 8, 0, 143, 144, 5, 19, 0, 0, 144, 146, 3, 16, 8, 0, 145, 143,
		1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150, 142, 1, 0, 0, 0,
		150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 5, 16, 0, 0, 153,
		21, 1, 0, 0, 0, 154, 155, 5, 13, 0, 0, 155, 156, 5, 4, 0, 0, 156, 157,
		3, 24, 12, 0, 157, 158, 5, 14, 0, 0, 158, 159, 5, 15, 0, 0, 159, 160, 3,
		16, 8, 0, 160, 161, 5, 16, 0, 0, 161, 162, 5, 13, 0, 0, 162, 163, 5, 5,
		0, 0, 163, 164, 5, 14, 0, 0, 164, 165, 5, 15, 0, 0, 165, 166, 3, 16, 8,
		0, 166, 167, 5, 16, 0, 0, 167, 23, 1, 0, 0, 0, 168, 169, 5, 15, 0, 0, 169,
		173, 3, 16, 8, 0, 170, 171, 3, 26, 13, 0, 171, 172, 3, 16, 8, 0, 172, 174,
		1, 0, 0, 0, 173, 170, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 175, 176, 5, 16, 0, 0, 176, 187, 1, 0, 0, 0, 177, 178, 5, 15, 0,
		0, 178, 182, 3, 24, 12, 0, 179, 180, 3, 26, 13, 0, 180, 181, 3, 24, 12,
		0, 181, 183, 1, 0, 0, 0, 182, 179, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		184, 1, 0, 0, 0, 184, 185, 5, 16, 0, 0, 185, 187, 1, 0, 0, 0, 186, 168,
		1, 0, 0, 0, 186, 177, 1, 0, 0, 0, 187, 25, 1, 0, 0, 0, 188, 189, 7, 2,
		0, 0, 189, 27, 1, 0, 0, 0, 190, 191, 5, 13, 0, 0, 191, 192, 5, 31, 0, 0,
		192, 193, 5, 42, 0, 0, 193, 194, 3, 32, 16, 0, 194, 195, 5, 14, 0, 0, 195,
		29, 1, 0, 0, 0, 196, 197, 5, 15, 0, 0, 197, 198, 3, 16, 8, 0, 198, 199,
		5, 16, 0, 0, 199, 31, 1, 0, 0, 0, 200, 201, 3, 30, 15, 0, 201, 202, 5,
		24, 0, 0, 202, 203, 3, 30, 15, 0, 203, 209, 1, 0, 0, 0, 204, 205, 3, 30,
		15, 0, 205, 206, 5, 22, 0, 0, 206, 207, 3, 30, 15, 0, 207, 209, 1, 0, 0,
		0, 208, 200, 1, 0, 0, 0, 208, 204, 1, 0, 0, 0, 209, 33, 1, 0, 0, 0, 210,
		211, 5, 3, 0, 0, 211, 215, 5, 11, 0, 0, 212, 214, 3, 36, 18, 0, 213, 212,
		1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 218, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 219, 5, 12, 0, 0,
		219, 35, 1, 0, 0, 0, 220, 221, 5, 45, 0, 0, 221, 222, 5, 17, 0, 0, 222,
		223, 7, 3, 0, 0, 223, 37, 1, 0, 0, 0, 224, 225, 5, 45, 0, 0, 225, 226,
		5, 17, 0, 0, 226, 229, 3, 12, 6, 0, 227, 228, 5, 18, 0, 0, 228, 230, 3,
		16, 8, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 39, 1, 0, 0,
		0, 20, 52, 59, 67, 77, 84, 94, 102, 109, 120, 127, 135, 138, 147, 150,
		173, 182, 186, 208, 215, 229,
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

// WorkflowDSLParserInit initializes any static state used to implement WorkflowDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWorkflowDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkflowDSLParserInit() {
	staticData := &WorkflowDSLParserParserStaticData
	staticData.once.Do(workflowdslparserParserInit)
}

// NewWorkflowDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWorkflowDSLParser(input antlr.TokenStream) *WorkflowDSLParser {
	WorkflowDSLParserInit()
	this := new(WorkflowDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &WorkflowDSLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "WorkflowDSLParser.g4"

	return this
}

// WorkflowDSLParser tokens.
const (
	WorkflowDSLParserEOF         = antlr.TokenEOF
	WorkflowDSLParserWORKFLOW    = 1
	WorkflowDSLParserOUT         = 2
	WorkflowDSLParserMODELS      = 3
	WorkflowDSLParserIF          = 4
	WorkflowDSLParserELSE        = 5
	WorkflowDSLParserSTRING_TYPE = 6
	WorkflowDSLParserINT_TYPE    = 7
	WorkflowDSLParserFLOAT_TYPE  = 8
	WorkflowDSLParserAUTO        = 9
	WorkflowDSLParserSTRUCT      = 10
	WorkflowDSLParserLBRACE      = 11
	WorkflowDSLParserRBRACE      = 12
	WorkflowDSLParserLBRACK      = 13
	WorkflowDSLParserRBRACK      = 14
	WorkflowDSLParserLPAREN      = 15
	WorkflowDSLParserRPAREN      = 16
	WorkflowDSLParserCOLON       = 17
	WorkflowDSLParserARROW       = 18
	WorkflowDSLParserCOMMA       = 19
	WorkflowDSLParserDOT         = 20
	WorkflowDSLParserSLASH       = 21
	WorkflowDSLParserPLUS        = 22
	WorkflowDSLParserMINUS       = 23
	WorkflowDSLParserMULT        = 24
	WorkflowDSLParserEQ          = 25
	WorkflowDSLParserNEQ         = 26
	WorkflowDSLParserLT          = 27
	WorkflowDSLParserLTE         = 28
	WorkflowDSLParserGT          = 29
	WorkflowDSLParserGTE         = 30
	WorkflowDSLParserDOLLAR      = 31
	WorkflowDSLParserAMP         = 32
	WorkflowDSLParserPIPE        = 33
	WorkflowDSLParserPERCENT     = 34
	WorkflowDSLParserAT          = 35
	WorkflowDSLParserHASH        = 36
	WorkflowDSLParserCARET       = 37
	WorkflowDSLParserTILDE       = 38
	WorkflowDSLParserAND         = 39
	WorkflowDSLParserOR          = 40
	WorkflowDSLParserNOT         = 41
	WorkflowDSLParserSUM         = 42
	WorkflowDSLParserIN_ID       = 43
	WorkflowDSLParserIN_ID_ARRAY = 44
	WorkflowDSLParserID          = 45
	WorkflowDSLParserINT         = 46
	WorkflowDSLParserFLOAT       = 47
	WorkflowDSLParserSTRING      = 48
	WorkflowDSLParserWS          = 49
	WorkflowDSLParserCOMMENT     = 50
)

// WorkflowDSLParser rules.
const (
	WorkflowDSLParserRULE_workflow         = 0
	WorkflowDSLParserRULE_workflowBody     = 1
	WorkflowDSLParserRULE_inputs           = 2
	WorkflowDSLParserRULE_inputBlock       = 3
	WorkflowDSLParserRULE_fieldDef         = 4
	WorkflowDSLParserRULE_type             = 5
	WorkflowDSLParserRULE_outtype          = 6
	WorkflowDSLParserRULE_out              = 7
	WorkflowDSLParserRULE_expr             = 8
	WorkflowDSLParserRULE_atom             = 9
	WorkflowDSLParserRULE_functionCall     = 10
	WorkflowDSLParserRULE_conditionalExpr  = 11
	WorkflowDSLParserRULE_condition        = 12
	WorkflowDSLParserRULE_comparator       = 13
	WorkflowDSLParserRULE_arithmeticExpr   = 14
	WorkflowDSLParserRULE_parenExpr        = 15
	WorkflowDSLParserRULE_arithExprContent = 16
	WorkflowDSLParserRULE_models           = 17
	WorkflowDSLParserRULE_modelEntry       = 18
	WorkflowDSLParserRULE_outField         = 19
)

// IWorkflowContext is an interface to support dynamic dispatch.
type IWorkflowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORKFLOW() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	WorkflowBody() IWorkflowBodyContext
	RBRACE() antlr.TerminalNode

	// IsWorkflowContext differentiates from other interfaces.
	IsWorkflowContext()
}

type WorkflowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorkflowContext() *WorkflowContext {
	var p = new(WorkflowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_workflow
	return p
}

func InitEmptyWorkflowContext(p *WorkflowContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_workflow
}

func (*WorkflowContext) IsWorkflowContext() {}

func NewWorkflowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorkflowContext {
	var p = new(WorkflowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_workflow

	return p
}

func (s *WorkflowContext) GetParser() antlr.Parser { return s.parser }

func (s *WorkflowContext) WORKFLOW() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserWORKFLOW, 0)
}

func (s *WorkflowContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *WorkflowContext) WorkflowBody() IWorkflowBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWorkflowBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWorkflowBodyContext)
}

func (s *WorkflowContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
}

func (s *WorkflowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkflowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorkflowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterWorkflow(s)
	}
}

func (s *WorkflowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitWorkflow(s)
	}
}

func (p *WorkflowDSLParser) Workflow() (localctx IWorkflowContext) {
	localctx = NewWorkflowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WorkflowDSLParserRULE_workflow)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(WorkflowDSLParserWORKFLOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(WorkflowDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.WorkflowBody()
	}
	{
		p.SetState(43)
		p.Match(WorkflowDSLParserRBRACE)
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

// IWorkflowBodyContext is an interface to support dynamic dispatch.
type IWorkflowBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Inputs() IInputsContext
	Out() IOutContext
	Models() IModelsContext

	// IsWorkflowBodyContext differentiates from other interfaces.
	IsWorkflowBodyContext()
}

type WorkflowBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorkflowBodyContext() *WorkflowBodyContext {
	var p = new(WorkflowBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_workflowBody
	return p
}

func InitEmptyWorkflowBodyContext(p *WorkflowBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_workflowBody
}

func (*WorkflowBodyContext) IsWorkflowBodyContext() {}

func NewWorkflowBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorkflowBodyContext {
	var p = new(WorkflowBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_workflowBody

	return p
}

func (s *WorkflowBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *WorkflowBodyContext) Inputs() IInputsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputsContext)
}

func (s *WorkflowBodyContext) Out() IOutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutContext)
}

func (s *WorkflowBodyContext) Models() IModelsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelsContext)
}

func (s *WorkflowBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkflowBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorkflowBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterWorkflowBody(s)
	}
}

func (s *WorkflowBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitWorkflowBody(s)
	}
}

func (p *WorkflowDSLParser) WorkflowBody() (localctx IWorkflowBodyContext) {
	localctx = NewWorkflowBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WorkflowDSLParserRULE_workflowBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Inputs()
	}
	{
		p.SetState(46)
		p.Out()
	}
	{
		p.SetState(47)
		p.Models()
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

// IInputsContext is an interface to support dynamic dispatch.
type IInputsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInputBlock() []IInputBlockContext
	InputBlock(i int) IInputBlockContext

	// IsInputsContext differentiates from other interfaces.
	IsInputsContext()
}

type InputsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputsContext() *InputsContext {
	var p = new(InputsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_inputs
	return p
}

func InitEmptyInputsContext(p *InputsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_inputs
}

func (*InputsContext) IsInputsContext() {}

func NewInputsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputsContext {
	var p = new(InputsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_inputs

	return p
}

func (s *InputsContext) GetParser() antlr.Parser { return s.parser }

func (s *InputsContext) AllInputBlock() []IInputBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInputBlockContext); ok {
			len++
		}
	}

	tst := make([]IInputBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInputBlockContext); ok {
			tst[i] = t.(IInputBlockContext)
			i++
		}
	}

	return tst
}

func (s *InputsContext) InputBlock(i int) IInputBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputBlockContext); ok {
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

	return t.(IInputBlockContext)
}

func (s *InputsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterInputs(s)
	}
}

func (s *InputsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitInputs(s)
	}
}

func (p *WorkflowDSLParser) Inputs() (localctx IInputsContext) {
	localctx = NewInputsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WorkflowDSLParserRULE_inputs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WorkflowDSLParserIN_ID || _la == WorkflowDSLParserIN_ID_ARRAY {
		{
			p.SetState(49)
			p.InputBlock()
		}

		p.SetState(52)
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

// IInputBlockContext is an interface to support dynamic dispatch.
type IInputBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IN_ID() antlr.TerminalNode
	IN_ID_ARRAY() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext

	// IsInputBlockContext differentiates from other interfaces.
	IsInputBlockContext()
}

type InputBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputBlockContext() *InputBlockContext {
	var p = new(InputBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_inputBlock
	return p
}

func InitEmptyInputBlockContext(p *InputBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_inputBlock
}

func (*InputBlockContext) IsInputBlockContext() {}

func NewInputBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputBlockContext {
	var p = new(InputBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_inputBlock

	return p
}

func (s *InputBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InputBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *InputBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
}

func (s *InputBlockContext) IN_ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserIN_ID, 0)
}

func (s *InputBlockContext) IN_ID_ARRAY() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserIN_ID_ARRAY, 0)
}

func (s *InputBlockContext) AllFieldDef() []IFieldDefContext {
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

func (s *InputBlockContext) FieldDef(i int) IFieldDefContext {
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

func (s *InputBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterInputBlock(s)
	}
}

func (s *InputBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitInputBlock(s)
	}
}

func (p *WorkflowDSLParser) InputBlock() (localctx IInputBlockContext) {
	localctx = NewInputBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WorkflowDSLParserRULE_inputBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WorkflowDSLParserIN_ID || _la == WorkflowDSLParserIN_ID_ARRAY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(55)
		p.Match(WorkflowDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkflowDSLParserID {
		{
			p.SetState(56)
			p.FieldDef()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(WorkflowDSLParserRBRACE)
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

// IFieldDefContext is an interface to support dynamic dispatch.
type IFieldDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext

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
	p.RuleIndex = WorkflowDSLParserRULE_fieldDef
	return p
}

func InitEmptyFieldDefContext(p *FieldDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_fieldDef
}

func (*FieldDefContext) IsFieldDefContext() {}

func NewFieldDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefContext {
	var p = new(FieldDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_fieldDef

	return p
}

func (s *FieldDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserID, 0)
}

func (s *FieldDefContext) COLON() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserCOLON, 0)
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

func (s *FieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterFieldDef(s)
	}
}

func (s *FieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitFieldDef(s)
	}
}

func (p *WorkflowDSLParser) FieldDef() (localctx IFieldDefContext) {
	localctx = NewFieldDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WorkflowDSLParserRULE_fieldDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(WorkflowDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(WorkflowDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9664) != 0 {
		{
			p.SetState(66)
			p.Type_()
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFieldDef() []IFieldDefContext
	FieldDef(i int) IFieldDefContext
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Type_() ITypeContext

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
	p.RuleIndex = WorkflowDSLParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserINT_TYPE, 0)
}

func (s *TypeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserFLOAT_TYPE, 0)
}

func (s *TypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserSTRING_TYPE, 0)
}

func (s *TypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserSTRUCT, 0)
}

func (s *TypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *TypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
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

func (s *TypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACK, 0)
}

func (s *TypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACK, 0)
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

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *WorkflowDSLParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WorkflowDSLParserRULE_type)
	var _la int

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkflowDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(WorkflowDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(WorkflowDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(WorkflowDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserSTRUCT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.Match(WorkflowDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(WorkflowDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == WorkflowDSLParserID {
			{
				p.SetState(74)
				p.FieldDef()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)
			p.Match(WorkflowDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserLBRACK:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(81)
			p.Match(WorkflowDSLParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(WorkflowDSLParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Type_()
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

// IOuttypeContext is an interface to support dynamic dispatch.
type IOuttypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllOutField() []IOutFieldContext
	OutField(i int) IOutFieldContext
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Outtype() IOuttypeContext

	// IsOuttypeContext differentiates from other interfaces.
	IsOuttypeContext()
}

type OuttypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOuttypeContext() *OuttypeContext {
	var p = new(OuttypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_outtype
	return p
}

func InitEmptyOuttypeContext(p *OuttypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_outtype
}

func (*OuttypeContext) IsOuttypeContext() {}

func NewOuttypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OuttypeContext {
	var p = new(OuttypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_outtype

	return p
}

func (s *OuttypeContext) GetParser() antlr.Parser { return s.parser }

func (s *OuttypeContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserINT_TYPE, 0)
}

func (s *OuttypeContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserFLOAT_TYPE, 0)
}

func (s *OuttypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserSTRING_TYPE, 0)
}

func (s *OuttypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserSTRUCT, 0)
}

func (s *OuttypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *OuttypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
}

func (s *OuttypeContext) AllOutField() []IOutFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutFieldContext); ok {
			len++
		}
	}

	tst := make([]IOutFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutFieldContext); ok {
			tst[i] = t.(IOutFieldContext)
			i++
		}
	}

	return tst
}

func (s *OuttypeContext) OutField(i int) IOutFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutFieldContext); ok {
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

	return t.(IOutFieldContext)
}

func (s *OuttypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACK, 0)
}

func (s *OuttypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACK, 0)
}

func (s *OuttypeContext) Outtype() IOuttypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOuttypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOuttypeContext)
}

func (s *OuttypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OuttypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OuttypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterOuttype(s)
	}
}

func (s *OuttypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitOuttype(s)
	}
}

func (p *WorkflowDSLParser) Outtype() (localctx IOuttypeContext) {
	localctx = NewOuttypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WorkflowDSLParserRULE_outtype)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case WorkflowDSLParserINT_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(WorkflowDSLParserINT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserFLOAT_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(WorkflowDSLParserFLOAT_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(WorkflowDSLParserSTRING_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserSTRUCT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(WorkflowDSLParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(WorkflowDSLParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == WorkflowDSLParserID {
			{
				p.SetState(91)
				p.OutField()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(WorkflowDSLParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case WorkflowDSLParserLBRACK:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(WorkflowDSLParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(WorkflowDSLParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Outtype()
		}

	case WorkflowDSLParserRBRACE, WorkflowDSLParserARROW, WorkflowDSLParserID:
		p.EnterOuterAlt(localctx, 6)

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

// IOutContext is an interface to support dynamic dispatch.
type IOutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllOutField() []IOutFieldContext
	OutField(i int) IOutFieldContext

	// IsOutContext differentiates from other interfaces.
	IsOutContext()
}

type OutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutContext() *OutContext {
	var p = new(OutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_out
	return p
}

func InitEmptyOutContext(p *OutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_out
}

func (*OutContext) IsOutContext() {}

func NewOutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutContext {
	var p = new(OutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_out

	return p
}

func (s *OutContext) GetParser() antlr.Parser { return s.parser }

func (s *OutContext) OUT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserOUT, 0)
}

func (s *OutContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *OutContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
}

func (s *OutContext) AllOutField() []IOutFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutFieldContext); ok {
			len++
		}
	}

	tst := make([]IOutFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutFieldContext); ok {
			tst[i] = t.(IOutFieldContext)
			i++
		}
	}

	return tst
}

func (s *OutContext) OutField(i int) IOutFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutFieldContext); ok {
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

	return t.(IOutFieldContext)
}

func (s *OutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterOut(s)
	}
}

func (s *OutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitOut(s)
	}
}

func (p *WorkflowDSLParser) Out() (localctx IOutContext) {
	localctx = NewOutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WorkflowDSLParserRULE_out)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(WorkflowDSLParserOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(WorkflowDSLParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WorkflowDSLParserID {
		{
			p.SetState(106)
			p.OutField()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Match(WorkflowDSLParserRBRACE)
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

	// Getter signatures
	FunctionCall() IFunctionCallContext
	ConditionalExpr() IConditionalExprContext
	ArithmeticExpr() IArithmeticExprContext
	Atom() IAtomContext
	INT() antlr.TerminalNode
	ParenExpr() IParenExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExprContext) ConditionalExpr() IConditionalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalExprContext)
}

func (s *ExprContext) ArithmeticExpr() IArithmeticExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticExprContext)
}

func (s *ExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserINT, 0)
}

func (s *ExprContext) ParenExpr() IParenExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *WorkflowDSLParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WorkflowDSLParserRULE_expr)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.ConditionalExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.ArithmeticExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.Match(WorkflowDSLParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.ParenExpr()
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

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	IN_ID() antlr.TerminalNode

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AtomContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserDOT)
}

func (s *AtomContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserDOT, i)
}

func (s *AtomContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserID)
}

func (s *AtomContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserID, i)
}

func (s *AtomContext) IN_ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserIN_ID, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *WorkflowDSLParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WorkflowDSLParserRULE_atom)
	var _la int

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.FunctionCall()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == WorkflowDSLParserDOT {
			{
				p.SetState(123)
				p.Match(WorkflowDSLParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.Match(WorkflowDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			_la = p.GetTokenStream().LA(1)

			if !(_la == WorkflowDSLParserIN_ID || _la == WorkflowDSLParserID) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == WorkflowDSLParserDOT {
			{
				p.SetState(131)
				p.Match(WorkflowDSLParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(132)
				p.Match(WorkflowDSLParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserID, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRPAREN, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
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

func (s *FunctionCallContext) Expr(i int) IExprContext {
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

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserCOMMA, i)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *WorkflowDSLParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WorkflowDSLParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(WorkflowDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(WorkflowDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114349209329664) != 0 {
		{
			p.SetState(142)
			p.Expr()
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == WorkflowDSLParserCOMMA {
			{
				p.SetState(143)
				p.Match(WorkflowDSLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.Expr()
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(152)
		p.Match(WorkflowDSLParserRPAREN)
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

// IConditionalExprContext is an interface to support dynamic dispatch.
type IConditionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	Condition() IConditionContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsConditionalExprContext differentiates from other interfaces.
	IsConditionalExprContext()
}

type ConditionalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExprContext() *ConditionalExprContext {
	var p = new(ConditionalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_conditionalExpr
	return p
}

func InitEmptyConditionalExprContext(p *ConditionalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_conditionalExpr
}

func (*ConditionalExprContext) IsConditionalExprContext() {}

func NewConditionalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_conditionalExpr

	return p
}

func (s *ConditionalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExprContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserLBRACK)
}

func (s *ConditionalExprContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACK, i)
}

func (s *ConditionalExprContext) IF() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserIF, 0)
}

func (s *ConditionalExprContext) Condition() IConditionContext {
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

func (s *ConditionalExprContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserRBRACK)
}

func (s *ConditionalExprContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACK, i)
}

func (s *ConditionalExprContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserLPAREN)
}

func (s *ConditionalExprContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLPAREN, i)
}

func (s *ConditionalExprContext) AllExpr() []IExprContext {
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

func (s *ConditionalExprContext) Expr(i int) IExprContext {
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

func (s *ConditionalExprContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(WorkflowDSLParserRPAREN)
}

func (s *ConditionalExprContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRPAREN, i)
}

func (s *ConditionalExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserELSE, 0)
}

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterConditionalExpr(s)
	}
}

func (s *ConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitConditionalExpr(s)
	}
}

func (p *WorkflowDSLParser) ConditionalExpr() (localctx IConditionalExprContext) {
	localctx = NewConditionalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WorkflowDSLParserRULE_conditionalExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(WorkflowDSLParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(WorkflowDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Condition()
	}
	{
		p.SetState(157)
		p.Match(WorkflowDSLParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(WorkflowDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Expr()
	}
	{
		p.SetState(160)
		p.Match(WorkflowDSLParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(WorkflowDSLParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(WorkflowDSLParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(WorkflowDSLParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(WorkflowDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Expr()
	}
	{
		p.SetState(166)
		p.Match(WorkflowDSLParserRPAREN)
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	Comparator() IComparatorContext
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLPAREN, 0)
}

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

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRPAREN, 0)
}

func (s *ConditionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ConditionContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *WorkflowDSLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WorkflowDSLParserRULE_condition)
	var _la int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(WorkflowDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Expr()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1651381370880) != 0 {
			{
				p.SetState(170)
				p.Comparator()
			}
			{
				p.SetState(171)
				p.Expr()
			}

		}
		{
			p.SetState(175)
			p.Match(WorkflowDSLParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(WorkflowDSLParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Condition()
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1651381370880) != 0 {
			{
				p.SetState(179)
				p.Comparator()
			}
			{
				p.SetState(180)
				p.Condition()
			}

		}
		{
			p.SetState(184)
			p.Match(WorkflowDSLParserRPAREN)
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserEQ, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserNEQ, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLT, 0)
}

func (s *ComparatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLTE, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserGT, 0)
}

func (s *ComparatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserGTE, 0)
}

func (s *ComparatorContext) AND() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserAND, 0)
}

func (s *ComparatorContext) OR() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserOR, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *WorkflowDSLParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, WorkflowDSLParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1651381370880) != 0) {
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

// IArithmeticExprContext is an interface to support dynamic dispatch.
type IArithmeticExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArithmeticExprContext differentiates from other interfaces.
	IsArithmeticExprContext()
}

type ArithmeticExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticExprContext() *ArithmeticExprContext {
	var p = new(ArithmeticExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_arithmeticExpr
	return p
}

func InitEmptyArithmeticExprContext(p *ArithmeticExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_arithmeticExpr
}

func (*ArithmeticExprContext) IsArithmeticExprContext() {}

func NewArithmeticExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticExprContext {
	var p = new(ArithmeticExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_arithmeticExpr

	return p
}

func (s *ArithmeticExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticExprContext) CopyAll(ctx *ArithmeticExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArithmeticExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DollarSumExprContext struct {
	ArithmeticExprContext
}

func NewDollarSumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DollarSumExprContext {
	var p = new(DollarSumExprContext)

	InitEmptyArithmeticExprContext(&p.ArithmeticExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArithmeticExprContext))

	return p
}

func (s *DollarSumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DollarSumExprContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACK, 0)
}

func (s *DollarSumExprContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserDOLLAR, 0)
}

func (s *DollarSumExprContext) SUM() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserSUM, 0)
}

func (s *DollarSumExprContext) ArithExprContent() IArithExprContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithExprContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithExprContentContext)
}

func (s *DollarSumExprContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACK, 0)
}

func (s *DollarSumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterDollarSumExpr(s)
	}
}

func (s *DollarSumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitDollarSumExpr(s)
	}
}

func (p *WorkflowDSLParser) ArithmeticExpr() (localctx IArithmeticExprContext) {
	localctx = NewArithmeticExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, WorkflowDSLParserRULE_arithmeticExpr)
	localctx = NewDollarSumExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(WorkflowDSLParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(WorkflowDSLParserDOLLAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Match(WorkflowDSLParserSUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.ArithExprContent()
	}
	{
		p.SetState(194)
		p.Match(WorkflowDSLParserRBRACK)
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

// IParenExprContext is an interface to support dynamic dispatch.
type IParenExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsParenExprContext differentiates from other interfaces.
	IsParenExprContext()
}

type ParenExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenExprContext() *ParenExprContext {
	var p = new(ParenExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_parenExpr
	return p
}

func InitEmptyParenExprContext(p *ParenExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_parenExpr
}

func (*ParenExprContext) IsParenExprContext() {}

func NewParenExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenExprContext {
	var p = new(ParenExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_parenExpr

	return p
}

func (s *ParenExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
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

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRPAREN, 0)
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (p *WorkflowDSLParser) ParenExpr() (localctx IParenExprContext) {
	localctx = NewParenExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, WorkflowDSLParserRULE_parenExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(WorkflowDSLParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Expr()
	}
	{
		p.SetState(198)
		p.Match(WorkflowDSLParserRPAREN)
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

// IArithExprContentContext is an interface to support dynamic dispatch.
type IArithExprContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParenExpr() []IParenExprContext
	ParenExpr(i int) IParenExprContext
	MULT() antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsArithExprContentContext differentiates from other interfaces.
	IsArithExprContentContext()
}

type ArithExprContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithExprContentContext() *ArithExprContentContext {
	var p = new(ArithExprContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_arithExprContent
	return p
}

func InitEmptyArithExprContentContext(p *ArithExprContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_arithExprContent
}

func (*ArithExprContentContext) IsArithExprContentContext() {}

func NewArithExprContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithExprContentContext {
	var p = new(ArithExprContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_arithExprContent

	return p
}

func (s *ArithExprContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithExprContentContext) AllParenExpr() []IParenExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParenExprContext); ok {
			len++
		}
	}

	tst := make([]IParenExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParenExprContext); ok {
			tst[i] = t.(IParenExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithExprContentContext) ParenExpr(i int) IParenExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenExprContext); ok {
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

	return t.(IParenExprContext)
}

func (s *ArithExprContentContext) MULT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserMULT, 0)
}

func (s *ArithExprContentContext) PLUS() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserPLUS, 0)
}

func (s *ArithExprContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithExprContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithExprContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterArithExprContent(s)
	}
}

func (s *ArithExprContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitArithExprContent(s)
	}
}

func (p *WorkflowDSLParser) ArithExprContent() (localctx IArithExprContentContext) {
	localctx = NewArithExprContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, WorkflowDSLParserRULE_arithExprContent)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.ParenExpr()
		}
		{
			p.SetState(201)
			p.Match(WorkflowDSLParserMULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.ParenExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.ParenExpr()
		}
		{
			p.SetState(205)
			p.Match(WorkflowDSLParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.ParenExpr()
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

// IModelsContext is an interface to support dynamic dispatch.
type IModelsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODELS() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllModelEntry() []IModelEntryContext
	ModelEntry(i int) IModelEntryContext

	// IsModelsContext differentiates from other interfaces.
	IsModelsContext()
}

type ModelsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelsContext() *ModelsContext {
	var p = new(ModelsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_models
	return p
}

func InitEmptyModelsContext(p *ModelsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_models
}

func (*ModelsContext) IsModelsContext() {}

func NewModelsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelsContext {
	var p = new(ModelsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_models

	return p
}

func (s *ModelsContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelsContext) MODELS() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserMODELS, 0)
}

func (s *ModelsContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserLBRACE, 0)
}

func (s *ModelsContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserRBRACE, 0)
}

func (s *ModelsContext) AllModelEntry() []IModelEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModelEntryContext); ok {
			len++
		}
	}

	tst := make([]IModelEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModelEntryContext); ok {
			tst[i] = t.(IModelEntryContext)
			i++
		}
	}

	return tst
}

func (s *ModelsContext) ModelEntry(i int) IModelEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelEntryContext); ok {
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

	return t.(IModelEntryContext)
}

func (s *ModelsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterModels(s)
	}
}

func (s *ModelsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitModels(s)
	}
}

func (p *WorkflowDSLParser) Models() (localctx IModelsContext) {
	localctx = NewModelsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, WorkflowDSLParserRULE_models)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(WorkflowDSLParserMODELS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(WorkflowDSLParserLBRACE)
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

	for _la == WorkflowDSLParserID {
		{
			p.SetState(212)
			p.ModelEntry()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(218)
		p.Match(WorkflowDSLParserRBRACE)
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

// IModelEntryContext is an interface to support dynamic dispatch.
type IModelEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	INT() antlr.TerminalNode
	AUTO() antlr.TerminalNode

	// IsModelEntryContext differentiates from other interfaces.
	IsModelEntryContext()
}

type ModelEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelEntryContext() *ModelEntryContext {
	var p = new(ModelEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_modelEntry
	return p
}

func InitEmptyModelEntryContext(p *ModelEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_modelEntry
}

func (*ModelEntryContext) IsModelEntryContext() {}

func NewModelEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelEntryContext {
	var p = new(ModelEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_modelEntry

	return p
}

func (s *ModelEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserID, 0)
}

func (s *ModelEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserCOLON, 0)
}

func (s *ModelEntryContext) INT() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserINT, 0)
}

func (s *ModelEntryContext) AUTO() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserAUTO, 0)
}

func (s *ModelEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterModelEntry(s)
	}
}

func (s *ModelEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitModelEntry(s)
	}
}

func (p *WorkflowDSLParser) ModelEntry() (localctx IModelEntryContext) {
	localctx = NewModelEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, WorkflowDSLParserRULE_modelEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(WorkflowDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(WorkflowDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WorkflowDSLParserAUTO || _la == WorkflowDSLParserINT) {
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

// IOutFieldContext is an interface to support dynamic dispatch.
type IOutFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Outtype() IOuttypeContext
	ARROW() antlr.TerminalNode
	Expr() IExprContext

	// IsOutFieldContext differentiates from other interfaces.
	IsOutFieldContext()
}

type OutFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutFieldContext() *OutFieldContext {
	var p = new(OutFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_outField
	return p
}

func InitEmptyOutFieldContext(p *OutFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WorkflowDSLParserRULE_outField
}

func (*OutFieldContext) IsOutFieldContext() {}

func NewOutFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutFieldContext {
	var p = new(OutFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkflowDSLParserRULE_outField

	return p
}

func (s *OutFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OutFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserID, 0)
}

func (s *OutFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserCOLON, 0)
}

func (s *OutFieldContext) Outtype() IOuttypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOuttypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOuttypeContext)
}

func (s *OutFieldContext) ARROW() antlr.TerminalNode {
	return s.GetToken(WorkflowDSLParserARROW, 0)
}

func (s *OutFieldContext) Expr() IExprContext {
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

func (s *OutFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.EnterOutField(s)
	}
}

func (s *OutFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkflowDSLParserListener); ok {
		listenerT.ExitOutField(s)
	}
}

func (p *WorkflowDSLParser) OutField() (localctx IOutFieldContext) {
	localctx = NewOutFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, WorkflowDSLParserRULE_outField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(WorkflowDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(WorkflowDSLParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Outtype()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WorkflowDSLParserARROW {
		{
			p.SetState(227)
			p.Match(WorkflowDSLParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Expr()
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
