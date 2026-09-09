package controllers

import (
	"encoding/json"
	"github.com/vaultnite/transcendent/internal/models"
	"net/http"

	"github.com/vaultnite/transcendent/internal/middleware"
)

type FriendsController struct {
	*http.ServeMux
}

func NewFriendsController() *FriendsController {
	c := &FriendsController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *FriendsController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.friends", "transcendent", "Transcendent-Waitingroom-Controller"))

	c.HandleFunc("GET /api/public/friends/{accountId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	})

	c.HandleFunc("GET /api/public/blocklist/{accountId}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.BlockListResponse{
			BlockedUsers: []string{},
		})
	})

	c.HandleFunc("GET /api/public/list/{namespace}/{accountId}/recentPlayers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.RecentPlayersResponse{
			RecentPlayers: []string{},
		})
	})
}
