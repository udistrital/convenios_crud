package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPaisCategoria_20190614_203530 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPaisCategoria_20190614_203530{}
	m.Created = "20190614_203530"

	migration.Register("CrearTablaPaisCategoria_20190614_203530", m)
}

// Run the migrations
func (m *CrearTablaPaisCategoria_20190614_203530) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE convenios.pais_categoria(id serial NOT NULL,	pais integer,	CONSTRAINT pk_pais_categoria PRIMARY KEY (id));")
m.SQL("ALTER TABLE convenios.pais_categoria OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaPaisCategoria_20190614_203530) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP TABLE IF EXISTS convenios.pais_categoria CASCADE");
}
