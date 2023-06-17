package sqlex

type SchemaEngine interface {
	ParseSchema(string) Schema
	WriteSchema(Schema) string
}

type Schema struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

func NewSchema(name string) *Schema {
	return &Schema{Name: name, Fields: make([]Field, 0)}
}
