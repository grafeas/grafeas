// Code generated from FilterExpression.g4 by ANTLR 4.11.1. DO NOT EDIT.

package gen // FilterExpression
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FilterExpression struct {
	*antlr.BaseParser
}

var filterexpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func filterexpressionParserInit() {
	staticData := &filterexpressionParserStaticData
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
		"filter", "expression", "sequence", "factor", "term", "restriction",
		"comparable", "comparator", "value", "primary", "argList", "composite",
		"text", "field", "number", "intVal", "floatVal", "notOp", "andOp", "orOp",
		"sep", "keyword",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 301, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 1, 0, 3, 0, 46, 8, 0, 1, 0, 5, 0, 49, 8, 0, 10, 0, 12, 0, 52,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 60, 8, 1, 10, 1, 12, 1,
		63, 9, 1, 1, 2, 1, 2, 4, 2, 67, 8, 2, 11, 2, 12, 2, 68, 1, 2, 5, 2, 72,
		8, 2, 10, 2, 12, 2, 75, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 81, 8, 3, 10,
		3, 12, 3, 84, 9, 3, 1, 4, 3, 4, 87, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5,
		93, 8, 5, 10, 5, 12, 5, 96, 9, 5, 1, 5, 1, 5, 5, 5, 100, 8, 5, 10, 5, 12,
		5, 103, 9, 5, 1, 5, 1, 5, 3, 5, 107, 8, 5, 1, 6, 1, 6, 3, 6, 111, 8, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 123,
		8, 8, 1, 8, 3, 8, 126, 8, 8, 1, 8, 1, 8, 1, 8, 5, 8, 131, 8, 8, 10, 8,
		12, 8, 134, 9, 8, 1, 8, 1, 8, 5, 8, 138, 8, 8, 10, 8, 12, 8, 141, 9, 8,
		1, 8, 1, 8, 5, 8, 145, 8, 8, 10, 8, 12, 8, 148, 9, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 3, 9, 154, 8, 9, 1, 9, 3, 9, 157, 8, 9, 1, 9, 3, 9, 160, 8, 9, 1,
		10, 5, 10, 163, 8, 10, 10, 10, 12, 10, 166, 9, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 172, 8, 10, 10, 10, 12, 10, 175, 9, 10, 1, 10, 5, 10, 178,
		8, 10, 10, 10, 12, 10, 181, 9, 10, 1, 11, 1, 11, 5, 11, 185, 8, 11, 10,
		11, 12, 11, 188, 9, 11, 1, 11, 1, 11, 5, 11, 192, 8, 11, 10, 11, 12, 11,
		195, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 201, 8, 12, 10, 12, 12,
		12, 204, 9, 12, 1, 13, 1, 13, 1, 13, 3, 13, 209, 8, 13, 1, 14, 1, 14, 3,
		14, 213, 8, 14, 1, 15, 3, 15, 216, 8, 15, 1, 15, 4, 15, 219, 8, 15, 11,
		15, 12, 15, 220, 1, 15, 3, 15, 224, 8, 15, 1, 15, 3, 15, 227, 8, 15, 1,
		16, 3, 16, 230, 8, 16, 1, 16, 4, 16, 233, 8, 16, 11, 16, 12, 16, 234, 1,
		16, 1, 16, 5, 16, 239, 8, 16, 10, 16, 12, 16, 242, 9, 16, 1, 16, 1, 16,
		4, 16, 246, 8, 16, 11, 16, 12, 16, 247, 3, 16, 250, 8, 16, 1, 16, 3, 16,
		253, 8, 16, 1, 17, 1, 17, 1, 17, 4, 17, 258, 8, 17, 11, 17, 12, 17, 259,
		3, 17, 262, 8, 17, 1, 18, 4, 18, 265, 8, 18, 11, 18, 12, 18, 266, 1, 18,
		1, 18, 4, 18, 271, 8, 18, 11, 18, 12, 18, 272, 1, 19, 4, 19, 276, 8, 19,
		11, 19, 12, 19, 277, 1, 19, 1, 19, 4, 19, 282, 8, 19, 11, 19, 12, 19, 283,
		1, 20, 5, 20, 287, 8, 20, 10, 20, 12, 20, 290, 9, 20, 1, 20, 1, 20, 5,
		20, 294, 8, 20, 10, 20, 12, 20, 297, 9, 20, 1, 21, 1, 21, 1, 21, 0, 1,
		16, 22, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 0, 4, 2, 0, 2, 2, 13, 18, 3, 0, 19, 19, 24, 24, 27,
		27, 3, 0, 19, 20, 24, 24, 27, 27, 1, 0, 3, 5, 326, 0, 45, 1, 0, 0, 0, 2,
		55, 1, 0, 0, 0, 4, 64, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8, 86, 1, 0, 0, 0,
		10, 90, 1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 112, 1, 0, 0, 0, 16, 114,
		1, 0, 0, 0, 18, 159, 1, 0, 0, 0, 20, 164, 1, 0, 0, 0, 22, 182, 1, 0, 0,
		0, 24, 198, 1, 0, 0, 0, 26, 208, 1, 0, 0, 0, 28, 212, 1, 0, 0, 0, 30, 226,
		1, 0, 0, 0, 32, 229, 1, 0, 0, 0, 34, 261, 1, 0, 0, 0, 36, 264, 1, 0, 0,
		0, 38, 275, 1, 0, 0, 0, 40, 288, 1, 0, 0, 0, 42, 298, 1, 0, 0, 0, 44, 46,
		3, 2, 1, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 50, 1, 0, 0, 0,
		47, 49, 5, 23, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53,
		54, 5, 0, 0, 1, 54, 1, 1, 0, 0, 0, 55, 61, 3, 4, 2, 0, 56, 57, 3, 36, 18,
		0, 57, 58, 3, 4, 2, 0, 58, 60, 1, 0, 0, 0, 59, 56, 1, 0, 0, 0, 60, 63,
		1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 3, 1, 0, 0, 0,
		63, 61, 1, 0, 0, 0, 64, 73, 3, 6, 3, 0, 65, 67, 5, 23, 0, 0, 66, 65, 1,
		0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69,
		70, 1, 0, 0, 0, 70, 72, 3, 6, 3, 0, 71, 66, 1, 0, 0, 0, 72, 75, 1, 0, 0,
		0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0, 75, 73, 1,
		0, 0, 0, 76, 82, 3, 8, 4, 0, 77, 78, 3, 38, 19, 0, 78, 79, 3, 8, 4, 0,
		79, 81, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 7, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85,
		87, 3, 34, 17, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 89, 3, 10, 5, 0, 89, 9, 1, 0, 0, 0, 90, 106, 3, 12, 6, 0, 91,
		93, 5, 23, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0,
		0, 0, 94, 95, 1, 0, 0, 0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 101,
		3, 14, 7, 0, 98, 100, 5, 23, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0,
		0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0,
		103, 101, 1, 0, 0, 0, 104, 105, 3, 12, 6, 0, 105, 107, 1, 0, 0, 0, 106,
		94, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 11, 1, 0, 0, 0, 108, 111, 3,
		28, 14, 0, 109, 111, 3, 16, 8, 0, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0,
		0, 0, 111, 13, 1, 0, 0, 0, 112, 113, 7, 0, 0, 0, 113, 15, 1, 0, 0, 0, 114,
		115, 6, 8, -1, 0, 115, 116, 3, 18, 9, 0, 116, 146, 1, 0, 0, 0, 117, 118,
		10, 2, 0, 0, 118, 119, 5, 1, 0, 0, 119, 125, 3, 26, 13, 0, 120, 122, 5,
		6, 0, 0, 121, 123, 3, 20, 10, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 5, 7, 0, 0, 125, 120, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 145, 1, 0, 0, 0, 127, 128, 10, 1, 0, 0, 128,
		132, 5, 8, 0, 0, 129, 131, 5, 23, 0, 0, 130, 129, 1, 0, 0, 0, 131, 134,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 135, 139, 3, 12, 6, 0, 136, 138, 5, 23, 0,
		0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139,
		140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143,
		5, 9, 0, 0, 143, 145, 1, 0, 0, 0, 144, 117, 1, 0, 0, 0, 144, 127, 1, 0,
		0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0,
		147, 17, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 160, 3, 22, 11, 0, 150,
		156, 3, 24, 12, 0, 151, 153, 5, 6, 0, 0, 152, 154, 3, 20, 10, 0, 153, 152,
		1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 5, 7,
		0, 0, 156, 151, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0,
		158, 160, 5, 22, 0, 0, 159, 149, 1, 0, 0, 0, 159, 150, 1, 0, 0, 0, 159,
		158, 1, 0, 0, 0, 160, 19, 1, 0, 0, 0, 161, 163, 5, 23, 0, 0, 162, 161,
		1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 167, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 173, 3, 12, 6, 0,
		168, 169, 3, 40, 20, 0, 169, 170, 3, 12, 6, 0, 170, 172, 1, 0, 0, 0, 171,
		168, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 179, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 178, 5, 23,
		0, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0,
		179, 180, 1, 0, 0, 0, 180, 21, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 186,
		5, 6, 0, 0, 183, 185, 5, 23, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0,
		0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 186, 1, 0, 0, 0, 189, 193, 3, 2, 1, 0, 190, 192, 5, 23, 0, 0, 191,
		190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194,
		1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 5, 7,
		0, 0, 197, 23, 1, 0, 0, 0, 198, 202, 7, 1, 0, 0, 199, 201, 7, 2, 0, 0,
		200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202,
		203, 1, 0, 0, 0, 203, 25, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 209, 3,
		24, 12, 0, 206, 209, 5, 22, 0, 0, 207, 209, 3, 42, 21, 0, 208, 205, 1,
		0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 207, 1, 0, 0, 0, 209, 27, 1, 0, 0,
		0, 210, 213, 3, 32, 16, 0, 211, 213, 3, 30, 15, 0, 212, 210, 1, 0, 0, 0,
		212, 211, 1, 0, 0, 0, 213, 29, 1, 0, 0, 0, 214, 216, 5, 20, 0, 0, 215,
		214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 1, 0, 0, 0, 217, 219,
		5, 24, 0, 0, 218, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 221, 1, 0, 0, 0, 221, 227, 1, 0, 0, 0, 222, 224, 5, 20, 0, 0,
		223, 222, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225,
		227, 5, 25, 0, 0, 226, 215, 1, 0, 0, 0, 226, 223, 1, 0, 0, 0, 227, 31,
		1, 0, 0, 0, 228, 230, 5, 20, 0, 0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0,
		0, 0, 230, 249, 1, 0, 0, 0, 231, 233, 5, 24, 0, 0, 232, 231, 1, 0, 0, 0,
		233, 234, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		236, 1, 0, 0, 0, 236, 240, 5, 1, 0, 0, 237, 239, 5, 24, 0, 0, 238, 237,
		1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0,
		0, 0, 241, 250, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 245, 5, 1, 0, 0,
		244, 246, 5, 24, 0, 0, 245, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 232,
		1, 0, 0, 0, 249, 243, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 253, 5, 26,
		0, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 33, 1, 0, 0, 0,
		254, 262, 5, 20, 0, 0, 255, 257, 5, 5, 0, 0, 256, 258, 5, 23, 0, 0, 257,
		256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260,
		1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 254, 1, 0, 0, 0, 261, 255, 1, 0,
		0, 0, 262, 35, 1, 0, 0, 0, 263, 265, 5, 23, 0, 0, 264, 263, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		268, 1, 0, 0, 0, 268, 270, 5, 4, 0, 0, 269, 271, 5, 23, 0, 0, 270, 269,
		1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0,
		0, 0, 273, 37, 1, 0, 0, 0, 274, 276, 5, 23, 0, 0, 275, 274, 1, 0, 0, 0,
		276, 277, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278,
		279, 1, 0, 0, 0, 279, 281, 5, 3, 0, 0, 280, 282, 5, 23, 0, 0, 281, 280,
		1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 39, 1, 0, 0, 0, 285, 287, 5, 23, 0, 0, 286, 285, 1, 0, 0, 0,
		287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		291, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 295, 5, 12, 0, 0, 292, 294,
		5, 23, 0, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0,
		0, 0, 295, 296, 1, 0, 0, 0, 296, 41, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0,
		298, 299, 7, 3, 0, 0, 299, 43, 1, 0, 0, 0, 46, 45, 50, 61, 68, 73, 82,
		86, 94, 101, 106, 110, 122, 125, 132, 139, 144, 146, 153, 156, 159, 164,
		173, 179, 186, 193, 202, 208, 212, 215, 220, 223, 226, 229, 234, 240, 247,
		249, 252, 259, 261, 266, 272, 277, 283, 288, 295,
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

// FilterExpressionInit initializes any static state used to implement FilterExpression. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFilterExpression(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterExpressionInit() {
	staticData := &filterexpressionParserStaticData
	staticData.once.Do(filterexpressionParserInit)
}

// NewFilterExpression produces a new parser instance for the optional input antlr.TokenStream.
func NewFilterExpression(input antlr.TokenStream) *FilterExpression {
	FilterExpressionInit()
	this := new(FilterExpression)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &filterexpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FilterExpression.g4"

	return this
}

// FilterExpression tokens.
const (
	FilterExpressionEOF            = antlr.TokenEOF
	FilterExpressionDOT            = 1
	FilterExpressionHAS            = 2
	FilterExpressionOR             = 3
	FilterExpressionAND            = 4
	FilterExpressionNOT            = 5
	FilterExpressionLPAREN         = 6
	FilterExpressionRPAREN         = 7
	FilterExpressionLBRACE         = 8
	FilterExpressionRBRACE         = 9
	FilterExpressionLBRACKET       = 10
	FilterExpressionRBRACKET       = 11
	FilterExpressionCOMMA          = 12
	FilterExpressionLESS_THAN      = 13
	FilterExpressionLESS_EQUALS    = 14
	FilterExpressionGREATER_THAN   = 15
	FilterExpressionGREATER_EQUALS = 16
	FilterExpressionNOT_EQUALS     = 17
	FilterExpressionEQUALS         = 18
	FilterExpressionEXCLAIM        = 19
	FilterExpressionMINUS          = 20
	FilterExpressionPLUS           = 21
	FilterExpressionSTRING         = 22
	FilterExpressionWS             = 23
	FilterExpressionDIGIT          = 24
	FilterExpressionHEX_DIGIT      = 25
	FilterExpressionEXPONENT       = 26
	FilterExpressionTEXT           = 27
	FilterExpressionBACKSLASH      = 28
)

// FilterExpression rules.
const (
	FilterExpressionRULE_filter      = 0
	FilterExpressionRULE_expression  = 1
	FilterExpressionRULE_sequence    = 2
	FilterExpressionRULE_factor      = 3
	FilterExpressionRULE_term        = 4
	FilterExpressionRULE_restriction = 5
	FilterExpressionRULE_comparable  = 6
	FilterExpressionRULE_comparator  = 7
	FilterExpressionRULE_value       = 8
	FilterExpressionRULE_primary     = 9
	FilterExpressionRULE_argList     = 10
	FilterExpressionRULE_composite   = 11
	FilterExpressionRULE_text        = 12
	FilterExpressionRULE_field       = 13
	FilterExpressionRULE_number      = 14
	FilterExpressionRULE_intVal      = 15
	FilterExpressionRULE_floatVal    = 16
	FilterExpressionRULE_notOp       = 17
	FilterExpressionRULE_andOp       = 18
	FilterExpressionRULE_orOp        = 19
	FilterExpressionRULE_sep         = 20
	FilterExpressionRULE_keyword     = 21
)

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) EOF() antlr.TerminalNode {
	return s.GetToken(FilterExpressionEOF, 0)
}

func (s *FilterContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *FilterContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitFilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Filter() (localctx IFilterContext) {
	this := p
	_ = this

	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterExpressionRULE_filter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&190316642) != 0 {
		{
			p.SetState(44)
			p.Expression()
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(47)
			p.Match(FilterExpressionWS)
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(FilterExpressionEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() ISequenceContext

	// Get_andOp returns the _andOp rule contexts.
	Get_andOp() IAndOpContext

	// Get_sequence returns the _sequence rule contexts.
	Get_sequence() ISequenceContext

	// SetExpr sets the expr rule contexts.
	SetExpr(ISequenceContext)

	// Set_andOp sets the _andOp rule contexts.
	Set_andOp(IAndOpContext)

	// Set_sequence sets the _sequence rule contexts.
	Set_sequence(ISequenceContext)

	// GetOp returns the op rule context list.
	GetOp() []IAndOpContext

	// GetRest returns the rest rule context list.
	GetRest() []ISequenceContext

	// SetOp sets the op rule context list.
	SetOp([]IAndOpContext)

	// SetRest sets the rest rule context list.
	SetRest([]ISequenceContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	expr      ISequenceContext
	_andOp    IAndOpContext
	op        []IAndOpContext
	_sequence ISequenceContext
	rest      []ISequenceContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetExpr() ISequenceContext { return s.expr }

func (s *ExpressionContext) Get_andOp() IAndOpContext { return s._andOp }

func (s *ExpressionContext) Get_sequence() ISequenceContext { return s._sequence }

func (s *ExpressionContext) SetExpr(v ISequenceContext) { s.expr = v }

func (s *ExpressionContext) Set_andOp(v IAndOpContext) { s._andOp = v }

func (s *ExpressionContext) Set_sequence(v ISequenceContext) { s._sequence = v }

func (s *ExpressionContext) GetOp() []IAndOpContext { return s.op }

func (s *ExpressionContext) GetRest() []ISequenceContext { return s.rest }

func (s *ExpressionContext) SetOp(v []IAndOpContext) { s.op = v }

func (s *ExpressionContext) SetRest(v []ISequenceContext) { s.rest = v }

func (s *ExpressionContext) AllSequence() []ISequenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISequenceContext); ok {
			len++
		}
	}

	tst := make([]ISequenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISequenceContext); ok {
			tst[i] = t.(ISequenceContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Sequence(i int) ISequenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequenceContext); ok {
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

	return t.(ISequenceContext)
}

func (s *ExpressionContext) AllAndOp() []IAndOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndOpContext); ok {
			len++
		}
	}

	tst := make([]IAndOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndOpContext); ok {
			tst[i] = t.(IAndOpContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) AndOp(i int) IAndOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndOpContext); ok {
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

	return t.(IAndOpContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FilterExpressionRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)

		var _x = p.Sequence()

		localctx.(*ExpressionContext).expr = _x
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(56)

				var _x = p.AndOp()

				localctx.(*ExpressionContext)._andOp = _x
			}
			localctx.(*ExpressionContext).op = append(localctx.(*ExpressionContext).op, localctx.(*ExpressionContext)._andOp)
			{
				p.SetState(57)

				var _x = p.Sequence()

				localctx.(*ExpressionContext)._sequence = _x
			}
			localctx.(*ExpressionContext).rest = append(localctx.(*ExpressionContext).rest, localctx.(*ExpressionContext)._sequence)

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IFactorContext

	// Get_factor returns the _factor rule contexts.
	Get_factor() IFactorContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IFactorContext)

	// Set_factor sets the _factor rule contexts.
	Set_factor(IFactorContext)

	// GetRest returns the rest rule context list.
	GetRest() []IFactorContext

	// SetRest sets the rest rule context list.
	SetRest([]IFactorContext)

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	expr    IFactorContext
	_factor IFactorContext
	rest    []IFactorContext
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) GetExpr() IFactorContext { return s.expr }

func (s *SequenceContext) Get_factor() IFactorContext { return s._factor }

func (s *SequenceContext) SetExpr(v IFactorContext) { s.expr = v }

func (s *SequenceContext) Set_factor(v IFactorContext) { s._factor = v }

func (s *SequenceContext) GetRest() []IFactorContext { return s.rest }

func (s *SequenceContext) SetRest(v []IFactorContext) { s.rest = v }

func (s *SequenceContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *SequenceContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
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

	return t.(IFactorContext)
}

func (s *SequenceContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *SequenceContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Sequence() (localctx ISequenceContext) {
	this := p
	_ = this

	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterExpressionRULE_sequence)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)

		var _x = p.Factor()

		localctx.(*SequenceContext).expr = _x
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == FilterExpressionWS {
				{
					p.SetState(65)
					p.Match(FilterExpressionWS)
				}

				p.SetState(68)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(70)

				var _x = p.Factor()

				localctx.(*SequenceContext)._factor = _x
			}
			localctx.(*SequenceContext).rest = append(localctx.(*SequenceContext).rest, localctx.(*SequenceContext)._factor)

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() ITermContext

	// Get_orOp returns the _orOp rule contexts.
	Get_orOp() IOrOpContext

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// SetExpr sets the expr rule contexts.
	SetExpr(ITermContext)

	// Set_orOp sets the _orOp rule contexts.
	Set_orOp(IOrOpContext)

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetOp returns the op rule context list.
	GetOp() []IOrOpContext

	// GetRest returns the rest rule context list.
	GetRest() []ITermContext

	// SetOp sets the op rule context list.
	SetOp([]IOrOpContext)

	// SetRest sets the rest rule context list.
	SetRest([]ITermContext)

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   ITermContext
	_orOp  IOrOpContext
	op     []IOrOpContext
	_term  ITermContext
	rest   []ITermContext
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) GetExpr() ITermContext { return s.expr }

func (s *FactorContext) Get_orOp() IOrOpContext { return s._orOp }

func (s *FactorContext) Get_term() ITermContext { return s._term }

func (s *FactorContext) SetExpr(v ITermContext) { s.expr = v }

func (s *FactorContext) Set_orOp(v IOrOpContext) { s._orOp = v }

func (s *FactorContext) Set_term(v ITermContext) { s._term = v }

func (s *FactorContext) GetOp() []IOrOpContext { return s.op }

func (s *FactorContext) GetRest() []ITermContext { return s.rest }

func (s *FactorContext) SetOp(v []IOrOpContext) { s.op = v }

func (s *FactorContext) SetRest(v []ITermContext) { s.rest = v }

func (s *FactorContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *FactorContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
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

	return t.(ITermContext)
}

func (s *FactorContext) AllOrOp() []IOrOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrOpContext); ok {
			len++
		}
	}

	tst := make([]IOrOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrOpContext); ok {
			tst[i] = t.(IOrOpContext)
			i++
		}
	}

	return tst
}

func (s *FactorContext) OrOp(i int) IOrOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrOpContext); ok {
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

	return t.(IOrOpContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterExpressionRULE_factor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)

		var _x = p.Term()

		localctx.(*FactorContext).expr = _x
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(77)

				var _x = p.OrOp()

				localctx.(*FactorContext)._orOp = _x
			}
			localctx.(*FactorContext).op = append(localctx.(*FactorContext).op, localctx.(*FactorContext)._orOp)
			{
				p.SetState(78)

				var _x = p.Term()

				localctx.(*FactorContext)._term = _x
			}
			localctx.(*FactorContext).rest = append(localctx.(*FactorContext).rest, localctx.(*FactorContext)._term)

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() INotOpContext

	// GetExpr returns the expr rule contexts.
	GetExpr() IRestrictionContext

	// SetOp sets the op rule contexts.
	SetOp(INotOpContext)

	// SetExpr sets the expr rule contexts.
	SetExpr(IRestrictionContext)

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     INotOpContext
	expr   IRestrictionContext
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetOp() INotOpContext { return s.op }

func (s *TermContext) GetExpr() IRestrictionContext { return s.expr }

func (s *TermContext) SetOp(v INotOpContext) { s.op = v }

func (s *TermContext) SetExpr(v IRestrictionContext) { s.expr = v }

func (s *TermContext) Restriction() IRestrictionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRestrictionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRestrictionContext)
}

func (s *TermContext) NotOp() INotOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotOpContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterExpressionRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(85)

			var _x = p.NotOp()

			localctx.(*TermContext).op = _x
		}

	}
	{
		p.SetState(88)

		var _x = p.Restriction()

		localctx.(*TermContext).expr = _x
	}

	return localctx
}

// IRestrictionContext is an interface to support dynamic dispatch.
type IRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IComparableContext

	// GetOp returns the op rule contexts.
	GetOp() IComparatorContext

	// GetRest returns the rest rule contexts.
	GetRest() IComparableContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IComparableContext)

	// SetOp sets the op rule contexts.
	SetOp(IComparatorContext)

	// SetRest sets the rest rule contexts.
	SetRest(IComparableContext)

	// IsRestrictionContext differentiates from other interfaces.
	IsRestrictionContext()
}

type RestrictionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IComparableContext
	op     IComparatorContext
	rest   IComparableContext
}

func NewEmptyRestrictionContext() *RestrictionContext {
	var p = new(RestrictionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_restriction
	return p
}

func (*RestrictionContext) IsRestrictionContext() {}

func NewRestrictionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestrictionContext {
	var p = new(RestrictionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_restriction

	return p
}

func (s *RestrictionContext) GetParser() antlr.Parser { return s.parser }

func (s *RestrictionContext) GetExpr() IComparableContext { return s.expr }

func (s *RestrictionContext) GetOp() IComparatorContext { return s.op }

func (s *RestrictionContext) GetRest() IComparableContext { return s.rest }

func (s *RestrictionContext) SetExpr(v IComparableContext) { s.expr = v }

func (s *RestrictionContext) SetOp(v IComparatorContext) { s.op = v }

func (s *RestrictionContext) SetRest(v IComparableContext) { s.rest = v }

func (s *RestrictionContext) AllComparable() []IComparableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparableContext); ok {
			len++
		}
	}

	tst := make([]IComparableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparableContext); ok {
			tst[i] = t.(IComparableContext)
			i++
		}
	}

	return tst
}

func (s *RestrictionContext) Comparable(i int) IComparableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparableContext); ok {
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

	return t.(IComparableContext)
}

func (s *RestrictionContext) Comparator() IComparatorContext {
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

func (s *RestrictionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *RestrictionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *RestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestrictionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestrictionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitRestriction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Restriction() (localctx IRestrictionContext) {
	this := p
	_ = this

	localctx = NewRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterExpressionRULE_restriction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)

		var _x = p.Comparable()

		localctx.(*RestrictionContext).expr = _x
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FilterExpressionWS {
			{
				p.SetState(91)
				p.Match(FilterExpressionWS)
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)

			var _x = p.Comparator()

			localctx.(*RestrictionContext).op = _x
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FilterExpressionWS {
			{
				p.SetState(98)
				p.Match(FilterExpressionWS)
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(104)

			var _x = p.Comparable()

			localctx.(*RestrictionContext).rest = _x
		}

	}

	return localctx
}

// IComparableContext is an interface to support dynamic dispatch.
type IComparableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparableContext differentiates from other interfaces.
	IsComparableContext()
}

type ComparableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparableContext() *ComparableContext {
	var p = new(ComparableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_comparable
	return p
}

func (*ComparableContext) IsComparableContext() {}

func NewComparableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparableContext {
	var p = new(ComparableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_comparable

	return p
}

func (s *ComparableContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparableContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ComparableContext) Value() IValueContext {
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

func (s *ComparableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitComparable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Comparable() (localctx IComparableContext) {
	this := p
	_ = this

	localctx = NewComparableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FilterExpressionRULE_comparable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.value(0)
		}

	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) LESS_EQUALS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLESS_EQUALS, 0)
}

func (s *ComparatorContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLESS_THAN, 0)
}

func (s *ComparatorContext) GREATER_EQUALS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionGREATER_EQUALS, 0)
}

func (s *ComparatorContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionGREATER_THAN, 0)
}

func (s *ComparatorContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionNOT_EQUALS, 0)
}

func (s *ComparatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionEQUALS, 0)
}

func (s *ComparatorContext) HAS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionHAS, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Comparator() (localctx IComparatorContext) {
	this := p
	_ = this

	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FilterExpressionRULE_comparator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516100) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectOrCallContext struct {
	*ValueContext
	op   antlr.Token
	open antlr.Token
}

func NewSelectOrCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectOrCallContext {
	var p = new(SelectOrCallContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *SelectOrCallContext) GetOp() antlr.Token { return s.op }

func (s *SelectOrCallContext) GetOpen() antlr.Token { return s.open }

func (s *SelectOrCallContext) SetOp(v antlr.Token) { s.op = v }

func (s *SelectOrCallContext) SetOpen(v antlr.Token) { s.open = v }

func (s *SelectOrCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectOrCallContext) Value() IValueContext {
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

func (s *SelectOrCallContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectOrCallContext) DOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionDOT, 0)
}

func (s *SelectOrCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionRPAREN, 0)
}

func (s *SelectOrCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLPAREN, 0)
}

func (s *SelectOrCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *SelectOrCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitSelectOrCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicIndexContext struct {
	*ValueContext
	op    antlr.Token
	index IComparableContext
}

func NewDynamicIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicIndexContext {
	var p = new(DynamicIndexContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *DynamicIndexContext) GetOp() antlr.Token { return s.op }

func (s *DynamicIndexContext) SetOp(v antlr.Token) { s.op = v }

func (s *DynamicIndexContext) GetIndex() IComparableContext { return s.index }

func (s *DynamicIndexContext) SetIndex(v IComparableContext) { s.index = v }

func (s *DynamicIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicIndexContext) Value() IValueContext {
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

func (s *DynamicIndexContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionRBRACE, 0)
}

func (s *DynamicIndexContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLBRACE, 0)
}

func (s *DynamicIndexContext) Comparable() IComparableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparableContext)
}

func (s *DynamicIndexContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *DynamicIndexContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *DynamicIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitDynamicIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	*ValueContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Value() (localctx IValueContext) {
	return p.value(0)
}

func (p *FilterExpression) value(_p int) (localctx IValueContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, FilterExpressionRULE_value, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPrimaryExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(115)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectOrCallContext(p, NewValueContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionRULE_value)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(118)

					var _m = p.Match(FilterExpressionDOT)

					localctx.(*SelectOrCallContext).op = _m
				}
				{
					p.SetState(119)
					p.Field()
				}
				p.SetState(125)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(120)

						var _m = p.Match(FilterExpressionLPAREN)

						localctx.(*SelectOrCallContext).open = _m
					}
					p.SetState(122)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&198705218) != 0 {
						{
							p.SetState(121)
							p.ArgList()
						}

					}
					{
						p.SetState(124)
						p.Match(FilterExpressionRPAREN)
					}

				}

			case 2:
				localctx = NewDynamicIndexContext(p, NewValueContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FilterExpressionRULE_value)
				p.SetState(127)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(128)

					var _m = p.Match(FilterExpressionLBRACE)

					localctx.(*DynamicIndexContext).op = _m
				}
				p.SetState(132)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == FilterExpressionWS {
					{
						p.SetState(129)
						p.Match(FilterExpressionWS)
					}

					p.SetState(134)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(135)

					var _x = p.Comparable()

					localctx.(*DynamicIndexContext).index = _x
				}
				p.SetState(139)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == FilterExpressionWS {
					{
						p.SetState(136)
						p.Match(FilterExpressionWS)
					}

					p.SetState(141)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(142)
					p.Match(FilterExpressionRBRACE)
				}

			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyFrom(ctx *PrimaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringValContext struct {
	*PrimaryContext
	quotedText antlr.Token
}

func NewStringValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValContext {
	var p = new(StringValContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *StringValContext) GetQuotedText() antlr.Token { return s.quotedText }

func (s *StringValContext) SetQuotedText(v antlr.Token) { s.quotedText = v }

func (s *StringValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSTRING, 0)
}

func (s *StringValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitStringVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExprContext struct {
	*PrimaryContext
}

func NewNestedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExprContext {
	var p = new(NestedExprContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *NestedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExprContext) Composite() ICompositeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompositeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompositeContext)
}

func (s *NestedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitNestedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentOrGlobalCallContext struct {
	*PrimaryContext
	id   ITextContext
	open antlr.Token
}

func NewIdentOrGlobalCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentOrGlobalCallContext {
	var p = new(IdentOrGlobalCallContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *IdentOrGlobalCallContext) GetOpen() antlr.Token { return s.open }

func (s *IdentOrGlobalCallContext) SetOpen(v antlr.Token) { s.open = v }

func (s *IdentOrGlobalCallContext) GetId() ITextContext { return s.id }

func (s *IdentOrGlobalCallContext) SetId(v ITextContext) { s.id = v }

func (s *IdentOrGlobalCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentOrGlobalCallContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *IdentOrGlobalCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionRPAREN, 0)
}

func (s *IdentOrGlobalCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLPAREN, 0)
}

func (s *IdentOrGlobalCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *IdentOrGlobalCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitIdentOrGlobalCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FilterExpressionRULE_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionLPAREN:
		localctx = NewNestedExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Composite()
		}

	case FilterExpressionEXCLAIM, FilterExpressionDIGIT, FilterExpressionTEXT:
		localctx = NewIdentOrGlobalCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)

			var _x = p.Text()

			localctx.(*IdentOrGlobalCallContext).id = _x
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(151)

				var _m = p.Match(FilterExpressionLPAREN)

				localctx.(*IdentOrGlobalCallContext).open = _m
			}
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&198705218) != 0 {
				{
					p.SetState(152)
					p.ArgList()
				}

			}
			{
				p.SetState(155)
				p.Match(FilterExpressionRPAREN)
			}

		}

	case FilterExpressionSTRING:
		localctx = NewStringValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)

			var _m = p.Match(FilterExpressionSTRING)

			localctx.(*StringValContext).quotedText = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_comparable returns the _comparable rule contexts.
	Get_comparable() IComparableContext

	// Set_comparable sets the _comparable rule contexts.
	Set_comparable(IComparableContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IComparableContext

	// SetArgs sets the args rule context list.
	SetArgs([]IComparableContext)

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_comparable IComparableContext
	args        []IComparableContext
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) Get_comparable() IComparableContext { return s._comparable }

func (s *ArgListContext) Set_comparable(v IComparableContext) { s._comparable = v }

func (s *ArgListContext) GetArgs() []IComparableContext { return s.args }

func (s *ArgListContext) SetArgs(v []IComparableContext) { s.args = v }

func (s *ArgListContext) AllComparable() []IComparableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparableContext); ok {
			len++
		}
	}

	tst := make([]IComparableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparableContext); ok {
			tst[i] = t.(IComparableContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Comparable(i int) IComparableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparableContext); ok {
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

	return t.(IComparableContext)
}

func (s *ArgListContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *ArgListContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *ArgListContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
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

	return t.(ISepContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) ArgList() (localctx IArgListContext) {
	this := p
	_ = this

	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FilterExpressionRULE_argList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(161)
			p.Match(FilterExpressionWS)
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)

		var _x = p.Comparable()

		localctx.(*ArgListContext)._comparable = _x
	}
	localctx.(*ArgListContext).args = append(localctx.(*ArgListContext).args, localctx.(*ArgListContext)._comparable)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(168)
				p.Sep()
			}
			{
				p.SetState(169)

				var _x = p.Comparable()

				localctx.(*ArgListContext)._comparable = _x
			}
			localctx.(*ArgListContext).args = append(localctx.(*ArgListContext).args, localctx.(*ArgListContext)._comparable)

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(176)
			p.Match(FilterExpressionWS)
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICompositeContext is an interface to support dynamic dispatch.
type ICompositeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositeContext differentiates from other interfaces.
	IsCompositeContext()
}

type CompositeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositeContext() *CompositeContext {
	var p = new(CompositeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_composite
	return p
}

func (*CompositeContext) IsCompositeContext() {}

func NewCompositeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositeContext {
	var p = new(CompositeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_composite

	return p
}

func (s *CompositeContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionLPAREN, 0)
}

func (s *CompositeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompositeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterExpressionRPAREN, 0)
}

func (s *CompositeContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *CompositeContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *CompositeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitComposite(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Composite() (localctx ICompositeContext) {
	this := p
	_ = this

	localctx = NewCompositeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FilterExpressionRULE_composite)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(FilterExpressionLPAREN)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(183)
			p.Match(FilterExpressionWS)
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Expression()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(190)
			p.Match(FilterExpressionWS)
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(FilterExpressionRPAREN)
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionTEXT)
}

func (s *TextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionTEXT, i)
}

func (s *TextContext) AllEXCLAIM() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionEXCLAIM)
}

func (s *TextContext) EXCLAIM(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionEXCLAIM, i)
}

func (s *TextContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionDIGIT)
}

func (s *TextContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionDIGIT, i)
}

func (s *TextContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionMINUS)
}

func (s *TextContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionMINUS, i)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Text() (localctx ITextContext) {
	this := p
	_ = this

	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FilterExpressionRULE_text)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&151519232) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(199)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&152567808) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetQuotedText returns the quotedText token.
	GetQuotedText() antlr.Token

	// SetQuotedText sets the quotedText token.
	SetQuotedText(antlr.Token)

	// GetId returns the id rule contexts.
	GetId() ITextContext

	// SetId sets the id rule contexts.
	SetId(ITextContext)

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	id         ITextContext
	quotedText antlr.Token
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) GetQuotedText() antlr.Token { return s.quotedText }

func (s *FieldContext) SetQuotedText(v antlr.Token) { s.quotedText = v }

func (s *FieldContext) GetId() ITextContext { return s.id }

func (s *FieldContext) SetId(v ITextContext) { s.id = v }

func (s *FieldContext) Text() ITextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *FieldContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterExpressionSTRING, 0)
}

func (s *FieldContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FilterExpressionRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionEXCLAIM, FilterExpressionDIGIT, FilterExpressionTEXT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)

			var _x = p.Text()

			localctx.(*FieldContext).id = _x
		}

	case FilterExpressionSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)

			var _m = p.Match(FilterExpressionSTRING)

			localctx.(*FieldContext).quotedText = _m
		}

	case FilterExpressionOR, FilterExpressionAND, FilterExpressionNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Keyword()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FloatVal() IFloatValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatValContext)
}

func (s *NumberContext) IntVal() IIntValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntValContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FilterExpressionRULE_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.FloatVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.IntVal()
		}

	}

	return localctx
}

// IIntValContext is an interface to support dynamic dispatch.
type IIntValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntValContext differentiates from other interfaces.
	IsIntValContext()
}

type IntValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntValContext() *IntValContext {
	var p = new(IntValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_intVal
	return p
}

func (*IntValContext) IsIntValContext() {}

func NewIntValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntValContext {
	var p = new(IntValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_intVal

	return p
}

func (s *IntValContext) GetParser() antlr.Parser { return s.parser }

func (s *IntValContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionMINUS, 0)
}

func (s *IntValContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionDIGIT)
}

func (s *IntValContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionDIGIT, i)
}

func (s *IntValContext) HEX_DIGIT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionHEX_DIGIT, 0)
}

func (s *IntValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitIntVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) IntVal() (localctx IIntValContext) {
	this := p
	_ = this

	localctx = NewIntValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FilterExpressionRULE_intVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FilterExpressionMINUS {
			{
				p.SetState(214)
				p.Match(FilterExpressionMINUS)
			}

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FilterExpressionDIGIT {
			{
				p.SetState(217)
				p.Match(FilterExpressionDIGIT)
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FilterExpressionMINUS {
			{
				p.SetState(222)
				p.Match(FilterExpressionMINUS)
			}

		}
		{
			p.SetState(225)
			p.Match(FilterExpressionHEX_DIGIT)
		}

	}

	return localctx
}

// IFloatValContext is an interface to support dynamic dispatch.
type IFloatValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatValContext differentiates from other interfaces.
	IsFloatValContext()
}

type FloatValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValContext() *FloatValContext {
	var p = new(FloatValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_floatVal
	return p
}

func (*FloatValContext) IsFloatValContext() {}

func NewFloatValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValContext {
	var p = new(FloatValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_floatVal

	return p
}

func (s *FloatValContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValContext) DOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionDOT, 0)
}

func (s *FloatValContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionMINUS, 0)
}

func (s *FloatValContext) EXPONENT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionEXPONENT, 0)
}

func (s *FloatValContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionDIGIT)
}

func (s *FloatValContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionDIGIT, i)
}

func (s *FloatValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitFloatVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) FloatVal() (localctx IFloatValContext) {
	this := p
	_ = this

	localctx = NewFloatValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FilterExpressionRULE_floatVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionMINUS {
		{
			p.SetState(228)
			p.Match(FilterExpressionMINUS)
		}

	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionDIGIT:
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FilterExpressionDIGIT {
			{
				p.SetState(231)
				p.Match(FilterExpressionDIGIT)
			}

			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(236)
			p.Match(FilterExpressionDOT)
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FilterExpressionDIGIT {
			{
				p.SetState(237)
				p.Match(FilterExpressionDIGIT)
			}

			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case FilterExpressionDOT:
		{
			p.SetState(243)
			p.Match(FilterExpressionDOT)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FilterExpressionDIGIT {
			{
				p.SetState(244)
				p.Match(FilterExpressionDIGIT)
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FilterExpressionEXPONENT {
		{
			p.SetState(251)
			p.Match(FilterExpressionEXPONENT)
		}

	}

	return localctx
}

// INotOpContext is an interface to support dynamic dispatch.
type INotOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOpContext differentiates from other interfaces.
	IsNotOpContext()
}

type NotOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOpContext() *NotOpContext {
	var p = new(NotOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_notOp
	return p
}

func (*NotOpContext) IsNotOpContext() {}

func NewNotOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOpContext {
	var p = new(NotOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_notOp

	return p
}

func (s *NotOpContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FilterExpressionMINUS, 0)
}

func (s *NotOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionNOT, 0)
}

func (s *NotOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *NotOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *NotOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitNotOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) NotOp() (localctx INotOpContext) {
	this := p
	_ = this

	localctx = NewNotOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FilterExpressionRULE_notOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterExpressionMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(FilterExpressionMINUS)
		}

	case FilterExpressionNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Match(FilterExpressionNOT)
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FilterExpressionWS {
			{
				p.SetState(256)
				p.Match(FilterExpressionWS)
			}

			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAndOpContext is an interface to support dynamic dispatch.
type IAndOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndOpContext differentiates from other interfaces.
	IsAndOpContext()
}

type AndOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndOpContext() *AndOpContext {
	var p = new(AndOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_andOp
	return p
}

func (*AndOpContext) IsAndOpContext() {}

func NewAndOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndOpContext {
	var p = new(AndOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_andOp

	return p
}

func (s *AndOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AndOpContext) AND() antlr.TerminalNode {
	return s.GetToken(FilterExpressionAND, 0)
}

func (s *AndOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *AndOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *AndOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitAndOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) AndOp() (localctx IAndOpContext) {
	this := p
	_ = this

	localctx = NewAndOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FilterExpressionRULE_andOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FilterExpressionWS {
		{
			p.SetState(263)
			p.Match(FilterExpressionWS)
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(268)
		p.Match(FilterExpressionAND)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FilterExpressionWS {
		{
			p.SetState(269)
			p.Match(FilterExpressionWS)
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrOpContext is an interface to support dynamic dispatch.
type IOrOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrOpContext differentiates from other interfaces.
	IsOrOpContext()
}

type OrOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrOpContext() *OrOpContext {
	var p = new(OrOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_orOp
	return p
}

func (*OrOpContext) IsOrOpContext() {}

func NewOrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrOpContext {
	var p = new(OrOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_orOp

	return p
}

func (s *OrOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OrOpContext) OR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionOR, 0)
}

func (s *OrOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *OrOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *OrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitOrOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) OrOp() (localctx IOrOpContext) {
	this := p
	_ = this

	localctx = NewOrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FilterExpressionRULE_orOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FilterExpressionWS {
		{
			p.SetState(274)
			p.Match(FilterExpressionWS)
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(FilterExpressionOR)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FilterExpressionWS {
		{
			p.SetState(280)
			p.Match(FilterExpressionWS)
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISepContext is an interface to support dynamic dispatch.
type ISepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSepContext differentiates from other interfaces.
	IsSepContext()
}

type SepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepContext() *SepContext {
	var p = new(SepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_sep
	return p
}

func (*SepContext) IsSepContext() {}

func NewSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepContext {
	var p = new(SepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_sep

	return p
}

func (s *SepContext) GetParser() antlr.Parser { return s.parser }

func (s *SepContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FilterExpressionCOMMA, 0)
}

func (s *SepContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(FilterExpressionWS)
}

func (s *SepContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(FilterExpressionWS, i)
}

func (s *SepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitSep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Sep() (localctx ISepContext) {
	this := p
	_ = this

	localctx = NewSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FilterExpressionRULE_sep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(285)
			p.Match(FilterExpressionWS)
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(FilterExpressionCOMMA)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterExpressionWS {
		{
			p.SetState(292)
			p.Match(FilterExpressionWS)
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterExpressionRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterExpressionRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) OR() antlr.TerminalNode {
	return s.GetToken(FilterExpressionOR, 0)
}

func (s *KeywordContext) AND() antlr.TerminalNode {
	return s.GetToken(FilterExpressionAND, 0)
}

func (s *KeywordContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterExpressionNOT, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterExpressionVisitor:
		return t.VisitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterExpression) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FilterExpressionRULE_keyword)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&56) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *FilterExpression) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ValueContext = nil
		if localctx != nil {
			t = localctx.(*ValueContext)
		}
		return p.Value_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FilterExpression) Value_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
