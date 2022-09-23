package app

import (
	"godemo/app/controller"

	// echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	affair := e.Group("api/affair")
	// affair.GET("/query", controller.QueryAffair)
	affair.POST("/add", controller.AddAffair)
	// affair.POST("/delete", controller.DeleteAffair)
	// affair.POST("/update", controller.UpdateAffair)
}
