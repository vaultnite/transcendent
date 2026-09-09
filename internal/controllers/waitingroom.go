package controllers

import (
	"net/http"

    "github.com/vaultnite/transcendent/internal/middleware"
)

type WaitingRoomController struct {
	*http.ServeMux
}

func NewWaitingRoomController() *WaitingRoomController {
	c := &WaitingRoomController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *WaitingRoomController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.waitingroom.public", "transcendent", "Transcendent-Waitingroom-Controller"))

	c.HandleFunc("GET /api/waitingroom", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
