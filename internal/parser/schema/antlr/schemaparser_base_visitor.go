// Code generated from schema/antlr/SchemaParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // SchemaParser
import "github.com/antlr4-go/antlr/v4"

type BaseSchemaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSchemaParserVisitor) VisitSchema(ctx *SchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitData(ctx *DataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitData_body(ctx *Data_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitMembers(ctx *MembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitMember_identifier(ctx *Member_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitType_identifier(ctx *Type_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSchemaParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
