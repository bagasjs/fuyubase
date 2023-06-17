package sqlex

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type SQLite3Engine struct {
}

func NewSQLite3Engine() *SQLite3Engine {
	return &SQLite3Engine{}
}

func (e *SQLite3Engine) ParseSchema(source string) (*Schema, error) {
	lex := NewSQLLexer(source)
	tokens := lex.LoadAllTokens()
	tokensCount := len(tokens)
	schema := &Schema{Name: "", Fields: make([]Field, 0)}

	for i := 0; i < tokensCount; i++ {
		if tokens[i].Literal == "CREATE" {
			i += 1
			if tokens[i].Literal != "TABLE" {
				return schema, errors.New("Invalid syntax after 'CREATE' expecting 'TABLE'")
			}
			i += 1
			if tokens[i].Type != TokenIdentifier {
				return schema, errors.New("Invalid syntax identifier expecting after 'CREATE TABLE'")
			}

			schema.Name = tokens[i].Literal
			i += 1
			if tokens[i].Type != TokenLParen {
				return schema, errors.New("Invalid syntax 'CREATE TABLE xxx (' give a parenthese in ")
			}
		} else if tokens[i].Literal == "CONSTRAINT" {
			// Parse constraint into the schema field
		} else if tokens[i].Type == TokenIdentifier {
			field := Field{Nullable: true, PrimaryKey: false, ForeignKey: false, Unique: false}
			field.Name = tokens[i].Literal
			if i+1 >= tokensCount {
				return schema, errors.New(fmt.Sprintf("Invalid end of statement after '%s' in line %d pos %d", tokens[i].Literal, tokens[i].Line, tokens[i].Pos))
			}
			// Check field type
			field.Type = tokens[i+1].Literal
			i += 2
			if i < tokensCount && tokens[i].Type == TokenLParen {
				if i+2 < tokensCount && tokens[i+1].Type == TokenNumber && tokens[i+2].Type == TokenRParen {
					field.Type += "(" + tokens[i+1].Literal + ")"
					i += 3
				} else if i+2 >= tokensCount {
					return schema, errors.New(fmt.Sprintf("Invalid end of statement after '%s' in line %d pos %d", tokens[i].Literal, tokens[i].Line, tokens[i].Pos))
				} else {
					return schema, errors.New(fmt.Sprintf("Invalid syntax '%s %s' after '%s' in line %d pos %d", tokens[i+1].Literal, tokens[i+2].Literal, tokens[i].Literal, tokens[i].Line, tokens[i].Pos))
				}
			}

			log.Println("TYPE", field.Type)

			if i < tokensCount && tokens[i].Literal == "NOT" {
				if i+1 < tokensCount && tokens[i+1].Literal == "NULL" {
					field.Nullable = false
				}
				i += 2
			}
			log.Println("NULLABLE", field.Nullable)
		} else {
			return schema, errors.New(fmt.Sprintf("Invalid syntax '%s' at line %d pos %d", tokens[i].Literal, tokens[i].Line, tokens[i].Pos))
		}
	}

	return schema, nil
}

func (e *SQLite3Engine) WriteSchema(schema *Schema) string {
	body := ""
	constraints := ""
	uniqueFields := make([]string, 0)
	pkFields := make([]string, 0)
	fkConstraints := make([]string, 0)

	for _, field := range schema.Fields {
		if field.Nullable {
			body = fmt.Sprintf("%s\n\t%s %s,", body, field.Name, field.Type)
		} else {
			body = fmt.Sprintf("%s\n\t%s %s NOT NULL,", body, field.Name, field.Type)
		}

		if field.Unique {
			uniqueFields = append(uniqueFields, field.Name)
		}
		if field.PrimaryKey {
			pkFields = append(pkFields, field.Name)
		}

		if field.ForeignKey {
			fkConstraints = append(fkConstraints, fmt.Sprintf("CONSTRAINT FK_%s_%s FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE %s\n",
				schema.Name, field.Related, field.Name, field.Related, field.Reference, field.OnDelete))
		}
	}
	if len(uniqueFields) > 0 {
		constraints = fmt.Sprintf("%s\n\tCONSTRAINT UC_%s UNIQUE(%s)", constraints,
			schema.Name, strings.Join(uniqueFields, ","))
	}
	if len(pkFields) > 0 {

		constraints = fmt.Sprintf("%s,\n\tCONSTRAINT PK_%s PRIMARY KEY(%s)", constraints,
			schema.Name, strings.Join(pkFields, ","))
	}
	if len(fkConstraints) > 0 {
		constraints = fmt.Sprintf("%s,\n\t%s", constraints, strings.Join(fkConstraints, ",\n"))
	}

	return fmt.Sprintf("CREATE TABLE %s (%s)", schema.Name,
		fmt.Sprintf("%s\n\t%s\n", body, constraints))
}
