package controllers

import (
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
		w.WriteHeader(http.StatusNoContent)
	})

	c.HandleFunc("GET /api/public/blocklist/{accountId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
