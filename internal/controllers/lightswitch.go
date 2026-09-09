package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

    "github.com/vaultnite/transcendent/internal/middleware"
	"github.com/vaultnite/transcendent/internal/models"
)

type LightswitchController struct {
    *http.ServeMux
}

func NewLightSwitchController() *LightswitchController {
	c := &LightswitchController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *LightswitchController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.lightswitch.public", "transcendent", "Transcendent-LightSwitch-Controller"))

	c.HandleFunc("GET /api/service/{serviceId}/status", func(w http.ResponseWriter, r *http.Request) {
		res := c.getServiceStatus(r.PathValue("serviceId"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	})

	c.HandleFunc("GET /api/service/bulk/status", func(w http.ResponseWriter, r *http.Request) {
		services := r.URL.Query()["serviceId"]
		res := make([]models.LightswitchServiceStatus, len(services))

		for i, serviceId := range r.URL.Query()["serviceId"] {
			res[i] = c.getServiceStatus(serviceId)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	})
}

func (c *LightswitchController) getServiceStatus(serviceId string) models.LightswitchServiceStatus {
	return models.LightswitchServiceStatus{
		ServiceInstanceID:  strings.ToLower(serviceId),
		Status:             "UP",
		Message:            fmt.Sprintf("%s is UP", serviceId),
		MaintenanceURI:     "https://status.vaultnite.com",
		OverrideCatalogIDs: []string{},
		AllowedActions:     []string{"PLAY", "DOWNLOAD"},
		Banned:             false,
		LauncherInfoDTO:    models.LightswitchServiceStatusLauncherInfoDTO{},
	}
}
