// Code generated from schema/antlr/SchemaLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SchemaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SchemaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func schemalexerLexerInit() {
	staticData := &SchemaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'data'", "'effect'", "'view'", "'('", "')'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "DATA", "EFFECT", "VIEW", "LPARANS", "RPARANS", "QUESTION", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"DATA", "EFFECT", "VIEW", "LPARANS", "RPARANS", "QUESTION", "IDENTIFIER",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 54, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 43, 8, 6, 10,
		6, 12, 6, 46, 9, 6, 1, 7, 4, 7, 49, 8, 7, 11, 7, 12, 7, 50, 1, 7, 1, 7,
		0, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 3, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0,
		9, 10, 13, 13, 32, 32, 55, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 22, 1, 0, 0, 0, 5,
		29, 1, 0, 0, 0, 7, 34, 1, 0, 0, 0, 9, 36, 1, 0, 0, 0, 11, 38, 1, 0, 0,
		0, 13, 40, 1, 0, 0, 0, 15, 48, 1, 0, 0, 0, 17, 18, 5, 100, 0, 0, 18, 19,
		5, 97, 0, 0, 19, 20, 5, 116, 0, 0, 20, 21, 5, 97, 0, 0, 21, 2, 1, 0, 0,
		0, 22, 23, 5, 101, 0, 0, 23, 24, 5, 102, 0, 0, 24, 25, 5, 102, 0, 0, 25,
		26, 5, 101, 0, 0, 26, 27, 5, 99, 0, 0, 27, 28, 5, 116, 0, 0, 28, 4, 1,
		0, 0, 0, 29, 30, 5, 118, 0, 0, 30, 31, 5, 105, 0, 0, 31, 32, 5, 101, 0,
		0, 32, 33, 5, 119, 0, 0, 33, 6, 1, 0, 0, 0, 34, 35, 5, 40, 0, 0, 35, 8,
		1, 0, 0, 0, 36, 37, 5, 41, 0, 0, 37, 10, 1, 0, 0, 0, 38, 39, 5, 63, 0,
		0, 39, 12, 1, 0, 0, 0, 40, 44, 7, 0, 0, 0, 41, 43, 7, 1, 0, 0, 42, 41,
		1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0,
		45, 14, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 49, 7, 2, 0, 0, 48, 47, 1,
		0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		52, 1, 0, 0, 0, 52, 53, 6, 7, 0, 0, 53, 16, 1, 0, 0, 0, 3, 0, 44, 50, 1,
		6, 0, 0,
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

// SchemaLexerInit initializes any static state used to implement SchemaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSchemaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SchemaLexerInit() {
	staticData := &SchemaLexerLexerStaticData
	staticData.once.Do(schemalexerLexerInit)
}

// NewSchemaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSchemaLexer(input antlr.CharStream) *SchemaLexer {
	SchemaLexerInit()
	l := new(SchemaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SchemaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SchemaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SchemaLexer tokens.
const (
	SchemaLexerDATA       = 1
	SchemaLexerEFFECT     = 2
	SchemaLexerVIEW       = 3
	SchemaLexerLPARANS    = 4
	SchemaLexerRPARANS    = 5
	SchemaLexerQUESTION   = 6
	SchemaLexerIDENTIFIER = 7
	SchemaLexerWS         = 8
)
