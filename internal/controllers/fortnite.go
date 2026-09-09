package controllers

import (
	"github.com/vaultnite/transcendent/internal/middleware"
	"github.com/vaultnite/transcendent/internal/storage"
)

type FortniteController struct {
	*mcpController
}

func NewFortniteController(sp storage.Provider) *FortniteController {
	c := &FortniteController{mcpController: newMCPController(sp, "fortnite")}
	c.addRoutes()
	return c
}

func (c *FortniteController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.fortnite", "transcendent", "Transcendent-Fortnite-Controller"))
}
