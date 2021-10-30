package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickeasters/nobones-api/lookup"
)

type BonesDay struct {
	Bones bool `json:"bones"`
}

func GetBonesJSON(c echo.Context) error {
	bones, err := lookup.BonesDay()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, BonesDay{bones})
}

func GetBones(c echo.Context) error {
	bones, err := lookup.BonesDay()
	if err != nil {
		return err
	}
	if bones {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusNotFound)
}
