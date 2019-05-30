// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/convenios_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/naturaleza",
			beego.NSInclude(
				&controllers.NaturalezaController{},
			),
		),

		beego.NSNamespace("/categoria",
			beego.NSInclude(
				&controllers.CategoriaController{},
			),
		),

		beego.NSNamespace("/pais_categoria",
			beego.NSInclude(
				&controllers.PaisCategoriaController{},
			),
		),

		beego.NSNamespace("/entidad",
			beego.NSInclude(
				&controllers.EntidadController{},
			),
		),

		beego.NSNamespace("/estados",
			beego.NSInclude(
				&controllers.EstadosController{},
			),
		),

		beego.NSNamespace("/convenio",
			beego.NSInclude(
				&controllers.ConvenioController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
