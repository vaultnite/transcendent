// wex (Battle Breakers) controller
// all implementations are based on the documentation gathered by dippyishere (https://github.com/Breakers-Revived/battle-breakers-documentation) and have not been tested by me on a running client.
// any PRs/issues with findings are welcomed, even if to confirm everything works

package controllers

import (
	"github.com/vaultnite/transcendent/internal/middleware"
	"github.com/vaultnite/transcendent/internal/storage"
)

type WexController struct {
	*mcpController
}

func NewWexController(sp storage.Provider) *WexController {
	c := &WexController{mcpController: newMCPController(sp, "wex")}
	c.addRoutes()
	return c
}

func (c *WexController) addRoutes() {
	c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.wex", "transcendent", "Transcendent-Wex-Controller"))
}
