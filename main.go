package main

import (
	"fmt"
	"log"

	"github.com/bagasjs/fuyubase/sqlex"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	schema := sqlex.NewSchema("users")
	schema.BigIntFieldQ("id", false, false, false, false)
	schema.CharFieldQ("email", 255, false, false, false)
	schema.CharFieldQ("password", 255, false, false, false)
	schema.CharFieldQ("name", 255, true, false, false)
	engine := sqlex.NewSQLite3Engine()
	migration := engine.WriteSchema(schema)
	fmt.Println("SCHEMA 1:\n", migration)
	schema2, err := engine.ParseSchema(migration)
	if err != nil {
		log.Fatal(err.Error())
	}
	migration2 := engine.WriteSchema(schema2)
	fmt.Println("SCHEMA 2:\n", migration2)
}
