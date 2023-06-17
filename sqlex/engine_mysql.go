package sqlex

import (
	"fmt"
	"strings"
)

type MySQLEngine struct {
}

func NewMySQLEngine() *MySQLEngine {
	return &MySQLEngine{}
}

func (e *MySQLEngine) ParseSchema(source string) Schema {
	return Schema{}
}

func (e *MySQLEngine) WriteSchema(schema *Schema) string {
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
