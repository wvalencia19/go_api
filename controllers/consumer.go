package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/wvalencia19/go_api/models"
)

type ConsumerController struct {
	beego.Controller
}

// @Title GetIndexAction
// @Description create object
// @Param	latitude		query 	string	true		"Latitude from position"
// @Param	longitude		query 	string	true		"Longitude from position"
// @Success 200 {objet} models.Coordinates
// @Failure 403 body is empty
// @router / [get]
func (context *ConsumerController) GetIndexAction() {
	latitude := context.GetString("latitude")
	longitude := context.GetString("longitude")
	destination := models.Coordinates{}
	destination.Latitude, _ = strconv.ParseFloat(latitude, models.CoordinatesBitSize)
	destination.Longitude, _ = strconv.ParseFloat(longitude, models.CoordinatesBitSize)
	context.Data["json"] = destination
	context.ServeJSON()
}
