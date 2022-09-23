// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"github.com/sirupsen/logrus"
	"godemo/app"
	"godemo/model"
)

func main() {
	logrus.SetReportCaller(true)
	model.Init()
	app.InitWebFramework()
	app.StartServer()
}
