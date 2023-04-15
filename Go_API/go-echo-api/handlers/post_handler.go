package handlers

import (
	"net/http"
	"strconv"

	"github.com/arjun/modules/go-echo-api/service"
	"github.com/labstack/echo/v4"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()

	if err != nil {
		c.String(http.StatusBadGateway, "unable to get data")
	}

	res := make(map[string]interface{})
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "unable to get data")
	}
	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "unable to get data")
	}
	res := make(map[string]interface{})
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
