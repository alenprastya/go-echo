package controllers

import (
	"net/http"
	"strconv"

	"github.com/alen/echo-framework/models"
	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func StorePegwai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")

	result, err := models.StorePegawai(nama, alamat, telephone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdatePegwai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegwai(conv_id, nama, alamat, telephone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
