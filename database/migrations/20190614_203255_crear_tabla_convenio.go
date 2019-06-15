package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaConvenio_20190614_203255 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaConvenio_20190614_203255{}
	m.Created = "20190614_203255"

	migration.Register("CrearTablaConvenio_20190614_203255", m)
}

// Run the migrations
func (m *CrearTablaConvenio_20190614_203255) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE convenios.convenio(id serial NOT NULL,descripcion varchar NOT NULL,responsable varchar NOT NULL,correo_responsable varchar NOT NULL,enlace varchar NOT NULL,id_pais_categoria integer,id_entidad integer,id_estados integer,CONSTRAINT pk_convenio PRIMARY KEY (id));")
m.SQL("ALTER TABLE convenios.convenio OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaConvenio_20190614_203255) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP TABLE IF EXISTS convenios.convenio CASCADE;")

}
