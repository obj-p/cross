// Code generated from schema/antlr/SchemaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // SchemaParser
import "github.com/antlr4-go/antlr/v4"

// BaseSchemaParserListener is a complete listener for a parse tree produced by SchemaParser.
type BaseSchemaParserListener struct{}

var _ SchemaParserListener = &BaseSchemaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSchemaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSchemaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSchemaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSchemaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseSchemaParserListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseSchemaParserListener) ExitSchema(ctx *SchemaContext) {}

// EnterData is called when production data is entered.
func (s *BaseSchemaParserListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseSchemaParserListener) ExitData(ctx *DataContext) {}

// EnterData_body is called when production data_body is entered.
func (s *BaseSchemaParserListener) EnterData_body(ctx *Data_bodyContext) {}

// ExitData_body is called when production data_body is exited.
func (s *BaseSchemaParserListener) ExitData_body(ctx *Data_bodyContext) {}

// EnterMembers is called when production members is entered.
func (s *BaseSchemaParserListener) EnterMembers(ctx *MembersContext) {}

// ExitMembers is called when production members is exited.
func (s *BaseSchemaParserListener) ExitMembers(ctx *MembersContext) {}

// EnterMember is called when production member is entered.
func (s *BaseSchemaParserListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseSchemaParserListener) ExitMember(ctx *MemberContext) {}

// EnterMember_identifier is called when production member_identifier is entered.
func (s *BaseSchemaParserListener) EnterMember_identifier(ctx *Member_identifierContext) {}

// ExitMember_identifier is called when production member_identifier is exited.
func (s *BaseSchemaParserListener) ExitMember_identifier(ctx *Member_identifierContext) {}

// EnterType_identifier is called when production type_identifier is entered.
func (s *BaseSchemaParserListener) EnterType_identifier(ctx *Type_identifierContext) {}

// ExitType_identifier is called when production type_identifier is exited.
func (s *BaseSchemaParserListener) ExitType_identifier(ctx *Type_identifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSchemaParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSchemaParserListener) ExitIdentifier(ctx *IdentifierContext) {}
