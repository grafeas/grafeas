// Code generated from FilterExpressionLexer.g4 by ANTLR 4.11.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FilterExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var filterexpressionlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func filterexpressionlexerLexerInit() {
	staticData := &filterexpressionlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'.'", "':'", "'OR'", "'AND'", "'NOT'", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "','", "'<'", "'<='", "'>'", "'>='", "'!='", "'='", "'!'",
		"'-'", "'+'", "", "", "", "", "", "", "'\\'",
	}
	staticData.symbolicNames = []string{
		"", "DOT", "HAS", "OR", "AND", "NOT", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "COMMA", "LESS_THAN", "LESS_EQUALS",
		"GREATER_THAN", "GREATER_EQUALS", "NOT_EQUALS", "EQUALS", "EXCLAIM",
		"MINUS", "PLUS", "STRING", "WS", "DIGIT", "HEX_DIGIT", "EXPONENT", "TEXT",
		"BACKSLASH",
	}
	staticData.ruleNames = []string{
		"DOT", "HAS", "OR", "AND", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"LBRACKET", "RBRACKET", "COMMA", "LESS_THAN", "LESS_EQUALS", "GREATER_THAN",
		"GREATER_EQUALS", "NOT_EQUALS", "EQUALS", "EXCLAIM", "MINUS", "PLUS",
		"STRING", "WS", "DIGIT", "HEX_DIGIT", "EXPONENT", "TEXT", "BACKSLASH",
		"Character", "TextEsc", "UnicodeEsc", "OctalEsc", "HexEsc", "Digit",
		"Exponent", "HexDigit", "OctalDigit", "StartChar", "MidChar", "EscapedChar",
		"Whitespace", "CharactersFromU00A1",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 244, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 138, 8, 21,
		10, 21, 12, 21, 141, 9, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 4, 24, 153, 8, 24, 11, 24, 12, 24, 154, 1, 25,
		1, 25, 1, 26, 1, 26, 3, 26, 161, 8, 26, 1, 26, 1, 26, 5, 26, 165, 8, 26,
		10, 26, 12, 26, 168, 9, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 3, 28, 177, 8, 28, 1, 28, 3, 28, 180, 8, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 3, 29, 186, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 3, 31, 197, 8, 31, 1, 31, 3, 31, 200, 8, 31, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 3, 34, 215, 8, 34, 1, 34, 4, 34, 218, 8, 34, 11, 34, 12, 34, 219, 1,
		35, 1, 35, 3, 35, 224, 8, 35, 1, 36, 1, 36, 1, 37, 1, 37, 3, 37, 230, 8,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 236, 8, 38, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 41, 1, 41, 0, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 0, 59, 0, 61, 0, 63, 0, 65, 0, 67,
		0, 69, 0, 71, 0, 73, 0, 75, 0, 77, 0, 79, 0, 81, 0, 83, 0, 1, 0, 10, 3,
		0, 32, 33, 35, 91, 93, 126, 6, 0, 97, 98, 102, 102, 110, 110, 114, 114,
		116, 116, 118, 118, 1, 0, 48, 51, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101,
		2, 0, 65, 70, 97, 102, 1, 0, 48, 55, 7, 0, 35, 39, 42, 42, 47, 47, 59,
		59, 63, 90, 94, 122, 124, 124, 7, 0, 34, 34, 42, 43, 46, 46, 58, 58, 60,
		62, 92, 92, 126, 126, 3, 0, 9, 10, 12, 13, 32, 32, 252, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 1, 85, 1, 0, 0, 0, 3, 87, 1, 0, 0, 0, 5, 89, 1, 0, 0, 0, 7, 92, 1, 0,
		0, 0, 9, 96, 1, 0, 0, 0, 11, 100, 1, 0, 0, 0, 13, 102, 1, 0, 0, 0, 15,
		104, 1, 0, 0, 0, 17, 106, 1, 0, 0, 0, 19, 108, 1, 0, 0, 0, 21, 110, 1,
		0, 0, 0, 23, 112, 1, 0, 0, 0, 25, 114, 1, 0, 0, 0, 27, 116, 1, 0, 0, 0,
		29, 119, 1, 0, 0, 0, 31, 121, 1, 0, 0, 0, 33, 124, 1, 0, 0, 0, 35, 127,
		1, 0, 0, 0, 37, 129, 1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 133, 1, 0, 0,
		0, 43, 135, 1, 0, 0, 0, 45, 144, 1, 0, 0, 0, 47, 146, 1, 0, 0, 0, 49, 148,
		1, 0, 0, 0, 51, 156, 1, 0, 0, 0, 53, 160, 1, 0, 0, 0, 55, 169, 1, 0, 0,
		0, 57, 179, 1, 0, 0, 0, 59, 185, 1, 0, 0, 0, 61, 187, 1, 0, 0, 0, 63, 194,
		1, 0, 0, 0, 65, 203, 1, 0, 0, 0, 67, 209, 1, 0, 0, 0, 69, 211, 1, 0, 0,
		0, 71, 223, 1, 0, 0, 0, 73, 225, 1, 0, 0, 0, 75, 229, 1, 0, 0, 0, 77, 235,
		1, 0, 0, 0, 79, 237, 1, 0, 0, 0, 81, 240, 1, 0, 0, 0, 83, 242, 1, 0, 0,
		0, 85, 86, 5, 46, 0, 0, 86, 2, 1, 0, 0, 0, 87, 88, 5, 58, 0, 0, 88, 4,
		1, 0, 0, 0, 89, 90, 5, 79, 0, 0, 90, 91, 5, 82, 0, 0, 91, 6, 1, 0, 0, 0,
		92, 93, 5, 65, 0, 0, 93, 94, 5, 78, 0, 0, 94, 95, 5, 68, 0, 0, 95, 8, 1,
		0, 0, 0, 96, 97, 5, 78, 0, 0, 97, 98, 5, 79, 0, 0, 98, 99, 5, 84, 0, 0,
		99, 10, 1, 0, 0, 0, 100, 101, 5, 40, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103,
		5, 41, 0, 0, 103, 14, 1, 0, 0, 0, 104, 105, 5, 91, 0, 0, 105, 16, 1, 0,
		0, 0, 106, 107, 5, 93, 0, 0, 107, 18, 1, 0, 0, 0, 108, 109, 5, 123, 0,
		0, 109, 20, 1, 0, 0, 0, 110, 111, 5, 125, 0, 0, 111, 22, 1, 0, 0, 0, 112,
		113, 5, 44, 0, 0, 113, 24, 1, 0, 0, 0, 114, 115, 5, 60, 0, 0, 115, 26,
		1, 0, 0, 0, 116, 117, 5, 60, 0, 0, 117, 118, 5, 61, 0, 0, 118, 28, 1, 0,
		0, 0, 119, 120, 5, 62, 0, 0, 120, 30, 1, 0, 0, 0, 121, 122, 5, 62, 0, 0,
		122, 123, 5, 61, 0, 0, 123, 32, 1, 0, 0, 0, 124, 125, 5, 33, 0, 0, 125,
		126, 5, 61, 0, 0, 126, 34, 1, 0, 0, 0, 127, 128, 5, 61, 0, 0, 128, 36,
		1, 0, 0, 0, 129, 130, 5, 33, 0, 0, 130, 38, 1, 0, 0, 0, 131, 132, 5, 45,
		0, 0, 132, 40, 1, 0, 0, 0, 133, 134, 5, 43, 0, 0, 134, 42, 1, 0, 0, 0,
		135, 139, 5, 34, 0, 0, 136, 138, 3, 57, 28, 0, 137, 136, 1, 0, 0, 0, 138,
		141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 142,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 5, 34, 0, 0, 143, 44, 1, 0,
		0, 0, 144, 145, 3, 81, 40, 0, 145, 46, 1, 0, 0, 0, 146, 147, 3, 67, 33,
		0, 147, 48, 1, 0, 0, 0, 148, 149, 5, 48, 0, 0, 149, 150, 5, 120, 0, 0,
		150, 152, 1, 0, 0, 0, 151, 153, 3, 71, 35, 0, 152, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 50, 1,
		0, 0, 0, 156, 157, 3, 69, 34, 0, 157, 52, 1, 0, 0, 0, 158, 161, 3, 75,
		37, 0, 159, 161, 3, 59, 29, 0, 160, 158, 1, 0, 0, 0, 160, 159, 1, 0, 0,
		0, 161, 166, 1, 0, 0, 0, 162, 165, 3, 77, 38, 0, 163, 165, 3, 59, 29, 0,
		164, 162, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166,
		164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 54, 1, 0, 0, 0, 168, 166, 1,
		0, 0, 0, 169, 170, 5, 92, 0, 0, 170, 56, 1, 0, 0, 0, 171, 180, 7, 0, 0,
		0, 172, 180, 3, 83, 41, 0, 173, 180, 3, 59, 29, 0, 174, 176, 5, 92, 0,
		0, 175, 177, 7, 1, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177,
		180, 1, 0, 0, 0, 178, 180, 3, 81, 40, 0, 179, 171, 1, 0, 0, 0, 179, 172,
		1, 0, 0, 0, 179, 173, 1, 0, 0, 0, 179, 174, 1, 0, 0, 0, 179, 178, 1, 0,
		0, 0, 180, 58, 1, 0, 0, 0, 181, 186, 3, 79, 39, 0, 182, 186, 3, 61, 30,
		0, 183, 186, 3, 63, 31, 0, 184, 186, 3, 65, 32, 0, 185, 181, 1, 0, 0, 0,
		185, 182, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186,
		60, 1, 0, 0, 0, 187, 188, 5, 92, 0, 0, 188, 189, 5, 117, 0, 0, 189, 190,
		3, 71, 35, 0, 190, 191, 3, 71, 35, 0, 191, 192, 3, 71, 35, 0, 192, 193,
		3, 71, 35, 0, 193, 62, 1, 0, 0, 0, 194, 196, 5, 92, 0, 0, 195, 197, 7,
		2, 0, 0, 196, 195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0,
		0, 198, 200, 3, 73, 36, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 201, 1, 0, 0, 0, 201, 202, 3, 73, 36, 0, 202, 64, 1, 0, 0, 0, 203,
		204, 5, 92, 0, 0, 204, 205, 5, 120, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207,
		3, 71, 35, 0, 207, 208, 3, 71, 35, 0, 208, 66, 1, 0, 0, 0, 209, 210, 7,
		3, 0, 0, 210, 68, 1, 0, 0, 0, 211, 214, 7, 4, 0, 0, 212, 215, 3, 41, 20,
		0, 213, 215, 3, 39, 19, 0, 214, 212, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0,
		214, 215, 1, 0, 0, 0, 215, 217, 1, 0, 0, 0, 216, 218, 3, 67, 33, 0, 217,
		216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220,
		1, 0, 0, 0, 220, 70, 1, 0, 0, 0, 221, 224, 3, 67, 33, 0, 222, 224, 7, 5,
		0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0, 224, 72, 1, 0, 0, 0,
		225, 226, 7, 6, 0, 0, 226, 74, 1, 0, 0, 0, 227, 230, 7, 7, 0, 0, 228, 230,
		3, 83, 41, 0, 229, 227, 1, 0, 0, 0, 229, 228, 1, 0, 0, 0, 230, 76, 1, 0,
		0, 0, 231, 236, 3, 75, 37, 0, 232, 236, 3, 67, 33, 0, 233, 236, 3, 41,
		20, 0, 234, 236, 3, 39, 19, 0, 235, 231, 1, 0, 0, 0, 235, 232, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 78, 1, 0, 0, 0, 237,
		238, 5, 92, 0, 0, 238, 239, 7, 8, 0, 0, 239, 80, 1, 0, 0, 0, 240, 241,
		7, 9, 0, 0, 241, 82, 1, 0, 0, 0, 242, 243, 2, 161, 65534, 0, 243, 84, 1,
		0, 0, 0, 16, 0, 139, 154, 160, 164, 166, 176, 179, 185, 196, 199, 214,
		219, 223, 229, 235, 0,
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

// FilterExpressionLexerInit initializes any static state used to implement FilterExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFilterExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterExpressionLexerInit() {
	staticData := &filterexpressionlexerLexerStaticData
	staticData.once.Do(filterexpressionlexerLexerInit)
}

// NewFilterExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFilterExpressionLexer(input antlr.CharStream) *FilterExpressionLexer {
	FilterExpressionLexerInit()
	l := new(FilterExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &filterexpressionlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "FilterExpressionLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FilterExpressionLexer tokens.
const (
	FilterExpressionLexerDOT            = 1
	FilterExpressionLexerHAS            = 2
	FilterExpressionLexerOR             = 3
	FilterExpressionLexerAND            = 4
	FilterExpressionLexerNOT            = 5
	FilterExpressionLexerLPAREN         = 6
	FilterExpressionLexerRPAREN         = 7
	FilterExpressionLexerLBRACE         = 8
	FilterExpressionLexerRBRACE         = 9
	FilterExpressionLexerLBRACKET       = 10
	FilterExpressionLexerRBRACKET       = 11
	FilterExpressionLexerCOMMA          = 12
	FilterExpressionLexerLESS_THAN      = 13
	FilterExpressionLexerLESS_EQUALS    = 14
	FilterExpressionLexerGREATER_THAN   = 15
	FilterExpressionLexerGREATER_EQUALS = 16
	FilterExpressionLexerNOT_EQUALS     = 17
	FilterExpressionLexerEQUALS         = 18
	FilterExpressionLexerEXCLAIM        = 19
	FilterExpressionLexerMINUS          = 20
	FilterExpressionLexerPLUS           = 21
	FilterExpressionLexerSTRING         = 22
	FilterExpressionLexerWS             = 23
	FilterExpressionLexerDIGIT          = 24
	FilterExpressionLexerHEX_DIGIT      = 25
	FilterExpressionLexerEXPONENT       = 26
	FilterExpressionLexerTEXT           = 27
	FilterExpressionLexerBACKSLASH      = 28
)
