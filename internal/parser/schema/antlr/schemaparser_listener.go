// Code generated from schema/antlr/SchemaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // SchemaParser
import "github.com/antlr4-go/antlr/v4"

// SchemaParserListener is a complete listener for a parse tree produced by SchemaParser.
type SchemaParserListener interface {
	antlr.ParseTreeListener

	// EnterSchema is called when entering the schema production.
	EnterSchema(c *SchemaContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterData_body is called when entering the data_body production.
	EnterData_body(c *Data_bodyContext)

	// EnterMembers is called when entering the members production.
	EnterMembers(c *MembersContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterMember_identifier is called when entering the member_identifier production.
	EnterMember_identifier(c *Member_identifierContext)

	// EnterType_identifier is called when entering the type_identifier production.
	EnterType_identifier(c *Type_identifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitSchema is called when exiting the schema production.
	ExitSchema(c *SchemaContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitData_body is called when exiting the data_body production.
	ExitData_body(c *Data_bodyContext)

	// ExitMembers is called when exiting the members production.
	ExitMembers(c *MembersContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitMember_identifier is called when exiting the member_identifier production.
	ExitMember_identifier(c *Member_identifierContext)

	// ExitType_identifier is called when exiting the type_identifier production.
	ExitType_identifier(c *Type_identifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
