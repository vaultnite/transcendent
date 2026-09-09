// orion (Paragon) controller

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vaultnite/transcendent/internal/middleware"
	"github.com/vaultnite/transcendent/internal/models"
	"github.com/vaultnite/transcendent/internal/storage"
)

type OrionController struct {
	*mcpController
}

func NewOrionController(sp storage.Provider) *OrionController {
	c := &OrionController{mcpController: newMCPController(sp, "orion")}
	c.addRoutes()
	return c
}

func (c *OrionController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.orion", "transcendent", "Transcendent-Orion-Controller"))

	c.HandleFunc("GET /api/game/v2/ratings/account/{accountId}/mmr/{ratingType}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// TODO: proper resp body
		json.NewEncoder(w).Encode(&models.GetAccountMMRResponse{
			Rating:         9000,
			NumGamesPlayed: 1,
		})
	})

	// orion & ut; additional elo endpoints are also present, none present in dippy's wex dump nor fortnite
	c.HandleFunc("POST /api/game/v2/ratings/team/elo/{ratingType}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.GetTeamEloResponse{
			Rating: 9000,
		})
	})

	// base path is present in orion, ut only calls to 2 subpaths (/estimate & /report/{ratingType}/{timeWaited})
	c.HandleFunc("GET /api/game/v2/wait_times", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&[]models.WaitTimeEstimateOrion{})
	})

	c.HandleFunc("GET /api/game/v2/wait_times/hero/{heroId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
