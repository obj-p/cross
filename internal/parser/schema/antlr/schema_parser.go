// Code generated from schema/antlr/SchemaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // SchemaParser
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

type SchemaParser struct {
	*antlr.BaseParser
}

var SchemaParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func schemaparserParserInit() {
	staticData := &SchemaParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'data'", "'effect'", "'view'", "'('", "')'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "DATA", "EFFECT", "VIEW", "LPARANS", "RPARANS", "QUESTION", "IDENTIFIER",
		"WS",
	}
	staticData.RuleNames = []string{
		"schema", "data", "data_body", "members", "member", "member_identifier",
		"type_identifier", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 50, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12, 0,
		19, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 5,
		3, 33, 8, 3, 10, 3, 12, 3, 36, 9, 3, 1, 4, 1, 4, 3, 4, 40, 8, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8,
		10, 12, 14, 0, 0, 44, 0, 17, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 27, 1, 0,
		0, 0, 6, 34, 1, 0, 0, 0, 8, 37, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12, 45,
		1, 0, 0, 0, 14, 47, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0,
		18, 19, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1,
		0, 0, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 24, 5, 1, 0, 0, 24,
		25, 3, 14, 7, 0, 25, 26, 3, 4, 2, 0, 26, 3, 1, 0, 0, 0, 27, 28, 5, 4, 0,
		0, 28, 29, 3, 6, 3, 0, 29, 30, 5, 5, 0, 0, 30, 5, 1, 0, 0, 0, 31, 33, 3,
		8, 4, 0, 32, 31, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34,
		35, 1, 0, 0, 0, 35, 7, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 39, 3, 10, 5,
		0, 38, 40, 5, 6, 0, 0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41,
		1, 0, 0, 0, 41, 42, 3, 12, 6, 0, 42, 9, 1, 0, 0, 0, 43, 44, 3, 14, 7, 0,
		44, 11, 1, 0, 0, 0, 45, 46, 3, 14, 7, 0, 46, 13, 1, 0, 0, 0, 47, 48, 5,
		7, 0, 0, 48, 15, 1, 0, 0, 0, 3, 19, 34, 39,
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

// SchemaParserInit initializes any static state used to implement SchemaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSchemaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SchemaParserInit() {
	staticData := &SchemaParserParserStaticData
	staticData.once.Do(schemaparserParserInit)
}

// NewSchemaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSchemaParser(input antlr.TokenStream) *SchemaParser {
	SchemaParserInit()
	this := new(SchemaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SchemaParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SchemaParser.g4"

	return this
}

// SchemaParser tokens.
const (
	SchemaParserEOF        = antlr.TokenEOF
	SchemaParserDATA       = 1
	SchemaParserEFFECT     = 2
	SchemaParserVIEW       = 3
	SchemaParserLPARANS    = 4
	SchemaParserRPARANS    = 5
	SchemaParserQUESTION   = 6
	SchemaParserIDENTIFIER = 7
	SchemaParserWS         = 8
)

// SchemaParser rules.
const (
	SchemaParserRULE_schema            = 0
	SchemaParserRULE_data              = 1
	SchemaParserRULE_data_body         = 2
	SchemaParserRULE_members           = 3
	SchemaParserRULE_member            = 4
	SchemaParserRULE_member_identifier = 5
	SchemaParserRULE_type_identifier   = 6
	SchemaParserRULE_identifier        = 7
)

// ISchemaContext is an interface to support dynamic dispatch.
type ISchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllData() []IDataContext
	Data(i int) IDataContext

	// IsSchemaContext differentiates from other interfaces.
	IsSchemaContext()
}

type SchemaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaContext() *SchemaContext {
	var p = new(SchemaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_schema
	return p
}

func InitEmptySchemaContext(p *SchemaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_schema
}

func (*SchemaContext) IsSchemaContext() {}

func NewSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaContext {
	var p = new(SchemaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_schema

	return p
}

func (s *SchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaContext) EOF() antlr.TerminalNode {
	return s.GetToken(SchemaParserEOF, 0)
}

func (s *SchemaContext) AllData() []IDataContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataContext); ok {
			len++
		}
	}

	tst := make([]IDataContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataContext); ok {
			tst[i] = t.(IDataContext)
			i++
		}
	}

	return tst
}

func (s *SchemaContext) Data(i int) IDataContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataContext); ok {
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

	return t.(IDataContext)
}

func (s *SchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterSchema(s)
	}
}

func (s *SchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitSchema(s)
	}
}

func (s *SchemaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitSchema(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Schema() (localctx ISchemaContext) {
	localctx = NewSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SchemaParserRULE_schema)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SchemaParserDATA {
		{
			p.SetState(16)
			p.Data()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(SchemaParserEOF)
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

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATA() antlr.TerminalNode
	Identifier() IIdentifierContext
	Data_body() IData_bodyContext

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_data
	return p
}

func InitEmptyDataContext(p *DataContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_data
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) DATA() antlr.TerminalNode {
	return s.GetToken(SchemaParserDATA, 0)
}

func (s *DataContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DataContext) Data_body() IData_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_bodyContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitData(s)
	}
}

func (s *DataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SchemaParserRULE_data)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(SchemaParserDATA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Identifier()
	}
	{
		p.SetState(25)
		p.Data_body()
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

// IData_bodyContext is an interface to support dynamic dispatch.
type IData_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPARANS() antlr.TerminalNode
	Members() IMembersContext
	RPARANS() antlr.TerminalNode

	// IsData_bodyContext differentiates from other interfaces.
	IsData_bodyContext()
}

type Data_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_bodyContext() *Data_bodyContext {
	var p = new(Data_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_data_body
	return p
}

func InitEmptyData_bodyContext(p *Data_bodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_data_body
}

func (*Data_bodyContext) IsData_bodyContext() {}

func NewData_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_bodyContext {
	var p = new(Data_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_data_body

	return p
}

func (s *Data_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_bodyContext) LPARANS() antlr.TerminalNode {
	return s.GetToken(SchemaParserLPARANS, 0)
}

func (s *Data_bodyContext) Members() IMembersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMembersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *Data_bodyContext) RPARANS() antlr.TerminalNode {
	return s.GetToken(SchemaParserRPARANS, 0)
}

func (s *Data_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterData_body(s)
	}
}

func (s *Data_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitData_body(s)
	}
}

func (s *Data_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitData_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Data_body() (localctx IData_bodyContext) {
	localctx = NewData_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SchemaParserRULE_data_body)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(SchemaParserLPARANS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Members()
	}
	{
		p.SetState(29)
		p.Match(SchemaParserRPARANS)
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

// IMembersContext is an interface to support dynamic dispatch.
type IMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMember() []IMemberContext
	Member(i int) IMemberContext

	// IsMembersContext differentiates from other interfaces.
	IsMembersContext()
}

type MembersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembersContext() *MembersContext {
	var p = new(MembersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_members
	return p
}

func InitEmptyMembersContext(p *MembersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_members
}

func (*MembersContext) IsMembersContext() {}

func NewMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembersContext {
	var p = new(MembersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_members

	return p
}

func (s *MembersContext) GetParser() antlr.Parser { return s.parser }

func (s *MembersContext) AllMember() []IMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberContext); ok {
			len++
		}
	}

	tst := make([]IMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberContext); ok {
			tst[i] = t.(IMemberContext)
			i++
		}
	}

	return tst
}

func (s *MembersContext) Member(i int) IMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberContext); ok {
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

	return t.(IMemberContext)
}

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterMembers(s)
	}
}

func (s *MembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitMembers(s)
	}
}

func (s *MembersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitMembers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Members() (localctx IMembersContext) {
	localctx = NewMembersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SchemaParserRULE_members)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SchemaParserIDENTIFIER {
		{
			p.SetState(31)
			p.Member()
		}

		p.SetState(36)
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

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Member_identifier() IMember_identifierContext
	Type_identifier() IType_identifierContext
	QUESTION() antlr.TerminalNode

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_member
	return p
}

func InitEmptyMemberContext(p *MemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_member
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) Member_identifier() IMember_identifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMember_identifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMember_identifierContext)
}

func (s *MemberContext) Type_identifier() IType_identifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_identifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_identifierContext)
}

func (s *MemberContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(SchemaParserQUESTION, 0)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitMember(s)
	}
}

func (s *MemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SchemaParserRULE_member)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Member_identifier()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SchemaParserQUESTION {
		{
			p.SetState(38)
			p.Match(SchemaParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(41)
		p.Type_identifier()
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

// IMember_identifierContext is an interface to support dynamic dispatch.
type IMember_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsMember_identifierContext differentiates from other interfaces.
	IsMember_identifierContext()
}

type Member_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMember_identifierContext() *Member_identifierContext {
	var p = new(Member_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_member_identifier
	return p
}

func InitEmptyMember_identifierContext(p *Member_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_member_identifier
}

func (*Member_identifierContext) IsMember_identifierContext() {}

func NewMember_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Member_identifierContext {
	var p = new(Member_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_member_identifier

	return p
}

func (s *Member_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Member_identifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Member_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Member_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Member_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterMember_identifier(s)
	}
}

func (s *Member_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitMember_identifier(s)
	}
}

func (s *Member_identifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitMember_identifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Member_identifier() (localctx IMember_identifierContext) {
	localctx = NewMember_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SchemaParserRULE_member_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Identifier()
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

// IType_identifierContext is an interface to support dynamic dispatch.
type IType_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsType_identifierContext differentiates from other interfaces.
	IsType_identifierContext()
}

type Type_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_identifierContext() *Type_identifierContext {
	var p = new(Type_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_type_identifier
	return p
}

func InitEmptyType_identifierContext(p *Type_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_type_identifier
}

func (*Type_identifierContext) IsType_identifierContext() {}

func NewType_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_identifierContext {
	var p = new(Type_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_type_identifier

	return p
}

func (s *Type_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_identifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Type_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterType_identifier(s)
	}
}

func (s *Type_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitType_identifier(s)
	}
}

func (s *Type_identifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitType_identifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Type_identifier() (localctx IType_identifierContext) {
	localctx = NewType_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SchemaParserRULE_type_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Identifier()
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SchemaParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SchemaParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SchemaParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SchemaParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SchemaParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SchemaParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SchemaParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(SchemaParserIDENTIFIER)
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
