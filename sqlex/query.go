package sqlex

import (
	"fmt"
	"strings"
)

type Statement interface {
	ToSQLQueryString() string
}

type WhereStatement struct {
	conditions []string
}

func (self *WhereStatement) ToSQLQueryString() string {
	return "WHERE " + strings.Join(self.conditions, " ")
}

func (self *WhereStatement) Where(left string, op string, right string) *WhereStatement {
	if len(self.conditions) > 0 {
		self.conditions = append(self.conditions, "AND", left, op, right)
	} else {
		self.conditions = append(self.conditions, left, op, right)
	}
	return self
}

func (self *WhereStatement) AndWhere(left string, op string, right string) *WhereStatement {
	self.conditions = append(self.conditions, "AND", left, op, right)
	return self
}

func (self *WhereStatement) OrWhere(left string, op string, right string) *WhereStatement {
	self.conditions = append(self.conditions, "OR", left, op, right)
	return self
}

type SelectStatement struct {
	*WhereStatement
	columns []string
	table   string
}

func Select(columns ...string) *SelectStatement {
	return &SelectStatement{
		columns:        columns,
		table:          "",
		WhereStatement: &WhereStatement{conditions: make([]string, 0)},
	}
}

func (self *SelectStatement) From(table string) *SelectStatement {
	self.table = table
	return self
}

func (self *SelectStatement) ToSQLQueryString() string {
	stmt := fmt.Sprintf("SELECT (%s) FROM %s", strings.Join(self.columns, ","), self.table)
	if len(self.WhereStatement.conditions) > 0 {
		stmt = fmt.Sprintf("%s %s", stmt, self.WhereStatement.ToSQLQueryString())
	}
	return stmt
}

// func (self *SelectStatement) Where(left string, op string, right string) *SelectStatement {
// 	self.Where(left, op, right)
// 	return self
// }

// func (self *SelectStatement) AndWhere(left string, op string, right string) *SelectStatement {
// 	self.AndWhere(left, op, right)
// 	return self
// }

// func (self *SelectStatement) OrWhere(left string, op string, right string) *SelectStatement {
// 	self.OrWhere(left, op, right)
// 	return self
// }
