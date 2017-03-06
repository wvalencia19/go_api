package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/wvalencia19/go_api/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["github.com/wvalencia19/go_api/controllers:ConsumerController"],
		beego.ControllerComments{
			Method: "GetIndexAction",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
