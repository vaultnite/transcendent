package controllers

import (
	"net/http"

	"github.com/vaultnite/transcendent/internal/middleware"
)

type EULATrackingController struct {
	*http.ServeMux
}

func NewEULATrackingController() *EULATrackingController {
	c := &EULATrackingController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *EULATrackingController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.eulatracking", "transcendent", "Transcendent-Account-Controller"))

	c.HandleFunc("GET /api/shared/agreements/{agreementId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
