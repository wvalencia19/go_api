package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/wvalencia19/go_api/models"
)

type ConsumerController struct {
	beego.Controller
}

// @Title GetIndexAction
// @Description create object
// @Param	latitude		query 	string	true		"Latitude from position"
// @Param	longitude		query 	string	true		"Longitude from position"
// @Success 200 {objet} models.Result
// @Failure 403 body is empty
// @router / [get]
func (context *ConsumerController) GetIndexAction() {
	latitude := context.GetString("latitude")
	longitude := context.GetString("longitude")
	coordinates := models.Coordinates{}
	coordinates.Latitude = latitude
	coordinates.Longitude = longitude
	dataBody, _ := json.Marshal(coordinates)

	requestURL := "https://sitidata-stdr.appspot.com/api/geoinverso"
	request := httplib.Post(requestURL)
	request.Header("Authorization", "Token")
	request.Body(string(dataBody))

	result := models.Result{}
	request.ToJSON(&result)
	logResponse, _ := request.String()
	beego.Info("Response Geocoder: " + logResponse)
	context.Data["json"] = result
	context.ServeJSON()

}
