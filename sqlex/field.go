package sqlex

import "fmt"

type Field struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Nullable      bool   `json:"nullable"`
	Unique        bool   `json:"unique"`
	PrimaryKey    bool   `json:"primaryKey"`
	AutoIncrement bool   `json:"autoIncrement"`
	ForeignKey    bool   `json:"foreignKey"`
	Related       string `json:"related"`
	Reference     string `json:"reference"`
	OnDelete      string `json:"onDelete"`
}

func (self *Schema) CharField(name string, length uint32, nullable bool, unique bool,
	primaryKey bool, foreignKey bool, related string, reference string, onDelete string) {
	self.Fields = append(self.Fields, Field{
		Name:       name,
		Type:       fmt.Sprintf("VARCHAR(%d)", length),
		Nullable:   nullable,
		Unique:     unique,
		PrimaryKey: primaryKey,
		ForeignKey: foreignKey,
		Related:    related,
		Reference:  reference,
		OnDelete:   onDelete,
	})
}

func (self *Schema) CharFieldQ(name string, length uint32, nullable bool, unique bool,
	primaryKey bool) {
	self.Fields = append(self.Fields, Field{
		Name:       name,
		Type:       fmt.Sprintf("VARCHAR(%d)", length),
		Nullable:   nullable,
		Unique:     unique,
		PrimaryKey: primaryKey,
		ForeignKey: false,
		Related:    "",
		Reference:  "",
		OnDelete:   "",
	})
}

func (self *Schema) IntField(name string, nullable bool, unique bool, autoIncrement bool,
	primaryKey bool, foreignKey bool, related string, reference string, onDelete string) {
	self.Fields = append(self.Fields, Field{
		Name:          name,
		Nullable:      nullable,
		Type:          "INT",
		Unique:        unique,
		AutoIncrement: autoIncrement,
		PrimaryKey:    primaryKey,
		ForeignKey:    foreignKey,
		Related:       related,
		Reference:     reference,
		OnDelete:      onDelete,
	})
}

func (self *Schema) IntFieldQ(name string, nullable bool, unique bool, autoIncrement bool,
	primaryKey bool) {
	self.Fields = append(self.Fields, Field{
		Name:          name,
		Nullable:      nullable,
		Type:          "INT",
		Unique:        unique,
		AutoIncrement: autoIncrement,
		PrimaryKey:    primaryKey,
		ForeignKey:    false,
		Related:       "",
		Reference:     "",
		OnDelete:      "",
	})
}

func (self *Schema) BigIntField(name string, nullable bool, unique bool, autoIncrement bool,
	primaryKey bool, foreignKey bool, related string, reference string, onDelete string) {
	self.Fields = append(self.Fields, Field{
		Name:          name,
		Nullable:      nullable,
		Type:          "BIGINT",
		Unique:        unique,
		AutoIncrement: autoIncrement,
		PrimaryKey:    primaryKey,
		ForeignKey:    foreignKey,
		Related:       related,
		Reference:     reference,
		OnDelete:      onDelete,
	})
}

func (self *Schema) BigIntFieldQ(name string, nullable bool, unique bool, autoIncrement bool,
	primaryKey bool) {
	self.Fields = append(self.Fields, Field{
		Name:          name,
		Nullable:      nullable,
		Type:          "BIGINT",
		Unique:        unique,
		AutoIncrement: autoIncrement,
		PrimaryKey:    primaryKey,
		ForeignKey:    false,
		Related:       "",
		Reference:     "",
		OnDelete:      "",
	})
}

func (self *Schema) FloatField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "FLOAT",
		Unique:   unique,
	})
}

func (self *Schema) BooleanField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "BOOLEAN",
		Unique:   unique,
	})
}

func (self *Schema) TextField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "TEXT",
		Unique:   unique,
	})
}

func (self *Schema) DateField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "DATE",
		Unique:   unique,
	})
}

func (self *Schema) DateTimeField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "DATETIME",
		Unique:   unique,
	})
}

func (self *Schema) TimeField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "TIME",
		Unique:   unique,
	})
}

func (self *Schema) YearField(name string, nullable bool, unique bool) {
	self.Fields = append(self.Fields, Field{
		Name:     name,
		Nullable: nullable,
		Type:     "TIME",
		Unique:   unique,
	})
}
