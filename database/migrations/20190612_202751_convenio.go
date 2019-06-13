package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Convenio_20190612_202751 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Convenio_20190612_202751{}
	m.Created = "20190612_202751"

	migration.Register("Convenio_20190612_202751", m)
}

// Run the migrations
func (m *Convenio_20190612_202751) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE convenios.convenio(id serial NOT NULL,descripcion varchar NOT NULL,responsable varchar NOT NULL,correo_responsable varchar NOT NULL,enlace varchar NOT NULL,id_pais_categoria integer,id_entidad integer,id_estados integer,CONSTRAINT pk_convenio PRIMARY KEY (id));")
m.SQL("ALTER TABLE convenios.convenio OWNER TO desarrollooas;")

}

// Reverse the migrations
func (m *Convenio_20190612_202751) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenios.convenio CASCADE");
	
}
