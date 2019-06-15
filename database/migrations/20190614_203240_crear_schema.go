package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20190614_203240 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20190614_203240{}
	m.Created = "20190614_203240"

	migration.Register("CrearSchema_20190614_203240", m)
}

// Run the migrations
func (m *CrearSchema_20190614_203240) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE SCHEMA convenios AUTHORIZATION desarrollooas;")
}

// Reverse the migrations
func (m *CrearSchema_20190614_203240) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA convenios;")
}
