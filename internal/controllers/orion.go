// orion (Paragon) controller

package controllers

import (
	"github.com/vaultnite/transcendent/internal/middleware"
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
}
