package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickeasters/nobones-api/lookup"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const kind = "AdmissionReview"

func AdmissionWebhook(c echo.Context) error {
	var req v1.AdmissionReview
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.Request == nil {
		c.Logger().Info("No admission request present")
		return c.NoContent(http.StatusBadRequest)
	}

	out := v1.AdmissionReview{
		TypeMeta: req.TypeMeta,
		Response: &v1.AdmissionResponse{
			UID:    req.Request.UID,
			Result: &metav1.Status{},
		},
	}

	// what kind of day is it?
	bones, err := lookup.BonesDay()
	if err != nil {
		out.Response.Allowed = false
		out.Response.Result.Code = http.StatusInternalServerError
		out.Response.Result.Message = "Unable to determine if it's a bones day, so take it easy."
	} else if bones {
		out.Response.Allowed = true
		out.Response.Result.Code = http.StatusOK
		out.Response.Result.Message = "It's a bones day!"
	} else {
		out.Response.Allowed = false
		out.Response.Result.Code = http.StatusForbidden
		out.Response.Result.Message = "No bones today, so let's not deploy."
	}

	return c.JSON(http.StatusOK, out)
}
