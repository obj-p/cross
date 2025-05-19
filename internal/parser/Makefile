antlr-4.13.1-complete.jar:
	curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar

schema/antlr/schema_parser.go: schema/antlr/SchemaLexer.g4 schema/antlr/SchemaParser.g4 antlr-4.13.1-complete.jar
	java -jar antlr-4.13.1-complete.jar -Dlanguage=Go -package antlr -o ./ schema/antlr/SchemaLexer.g4 schema/antlr/SchemaParser.g4 -visitor

.PHONY: schema-antlr
schema-antlr: antlr-4.13.1-complete.jar schema/antlr/schema_parser.go
