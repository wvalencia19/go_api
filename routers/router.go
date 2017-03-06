// @APIVersion 1.0.0
// @Title First API on Golang
// @Description practice for create golang API on Beego framework
// @Contact wvalencia@tappsi.co
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/wvalencia19/go_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/consumer",
			beego.NSInclude(
				&controllers.ConsumerController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
