package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vaultnite/transcendent/internal/models"
)

type CalderaController struct {
	*http.ServeMux
}

func NewCalderaController() *CalderaController {
	c := &CalderaController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *CalderaController) addRoutes() {
	c.HandleFunc("GET /api/v1/launcher/racp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.CalderaAnticheatProviderResponse{
			Provider:     "EasyAntiCheat",
			JWT:          "TODO",
			DeploymentID: "62a9473a2dca46b29ccf17577fcf42d7",
			SandboxID:    "fn",
		})
	})
}
