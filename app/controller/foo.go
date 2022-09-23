package controller

import (
	"godemo/app/dto"
	"godemo/app/response"
	"godemo/model"
	errno "godemo/utils"
	"time"

	"github.com/labstack/echo/v4"
	// "gopkg.in/guregu/null.v4"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func AddAffair(c echo.Context) error {
	var err error
	var AffairRequest dto.AffairReq
	var Affair model.Affair
	err = c.Bind(&AffairRequest)
	if err != nil {
		errno := errno.NewErrno(10001, "Bind Error in QueryAffair", err)
		return response.SendResponse(c, errno)
	}
	Affair.StartTime, _ = time.Parse(TIME_LAYOUT, AffairRequest.Start)
	if AffairRequest.End != "" {
		Affair.EndTime, _ = time.Parse(TIME_LAYOUT, AffairRequest.End)
	}
	Affair.Name = AffairRequest.Name
	Affair.Emergency = AffairRequest.Emergency
	err = model.DB.Create(&Affair).Error
	if err != nil {
		errno := errno.NewErrno(10003, "Database Create Error", err)
		return response.SendResponse(c, errno)
	}
	return response.SendResponse(c, nil)
}

// func DeleteAffair(c echo.Context) error {
// 	// just a demo
// 	var result model.Foo
// 	model.DB.Model(&model.Foo{}).First(&result)

// 	return c.JSON(http.StatusOK, response.FooResponse{
// 		Foo: c.QueryParam("foo"),
// 		Bar: 42,
// 	})

// }

// func QueryAffair(c echo.Context) error {
// 	// just a demo
// 	var AffairRequest dto.AffairReq
// 	var ResultAffairs []dto.AffairReq
// 	var err error
// 	err = c.Bind(&AffairRequest)
// 	if err != nil {
// 		errno := errno.NewErrno(10001, "Bind Error in QueryAffair", err)
// 		return response.SendResponse(c, errno)
// 	}
// 	err = model.DB.Find(&ResultAffairs, AffairRequest).Error
// 	if err != nil {
// 		errno := errno.NewErrno(10002, "Error Find in Database", err)
// 		return response.SendResponse(c, errno)
// 	}
// 	return response.SendResponse(c, nil, ResultAffairs)
// }

// func UpdateAffair(c echo.Context) error {
// 	// just a demo
// 	var result model.Foo
// 	model.DB.Model(&model.Foo{}).First(&result)

// 	return c.JSON(http.StatusOK, response.FooResponse{
// 		Foo: c.QueryParam("foo"),
// 		Bar: 42,
// 	})
// }
