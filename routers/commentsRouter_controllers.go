package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:CategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:ConvenioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EntidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:EstadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:NaturalezaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/convenios_crud/controllers:PaisCategoriaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
