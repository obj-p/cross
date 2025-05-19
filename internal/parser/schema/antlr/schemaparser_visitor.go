// Code generated from schema/antlr/SchemaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // SchemaParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SchemaParser.
type SchemaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SchemaParser#schema.
	VisitSchema(ctx *SchemaContext) interface{}

	// Visit a parse tree produced by SchemaParser#data.
	VisitData(ctx *DataContext) interface{}

	// Visit a parse tree produced by SchemaParser#data_body.
	VisitData_body(ctx *Data_bodyContext) interface{}

	// Visit a parse tree produced by SchemaParser#members.
	VisitMembers(ctx *MembersContext) interface{}

	// Visit a parse tree produced by SchemaParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by SchemaParser#member_identifier.
	VisitMember_identifier(ctx *Member_identifierContext) interface{}

	// Visit a parse tree produced by SchemaParser#type_identifier.
	VisitType_identifier(ctx *Type_identifierContext) interface{}

	// Visit a parse tree produced by SchemaParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
