package response

import (
	"net/http"
	"godemo/utils"
	"github.com/labstack/echo/v4"
)

// Response models are defined in utils/errno

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c echo.Context, err error, data ...interface{}) error {
	code, msg := errno.DecodeErr(err)
	if len(data) > 0 {
		return c.JSON(http.StatusOK, Response{
			Code: code,
			Msg:  msg,
			Data: data,
		})
	}
	return c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
	})
}
