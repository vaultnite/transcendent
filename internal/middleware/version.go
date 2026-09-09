package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vaultnite/transcendent/internal/buildinfo"
	"github.com/vaultnite/transcendent/internal/models"
)

func NewVersionHandler(app, cln, moduleName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.APIVersionResponse{
			App:                       app,
			ServerDate:                time.Now().Format("2006-01-02T15:04:05.000Z"),
			OverridePropertiesVersion: "unknown",
			CLN:                       cln,
			Build:                     buildinfo.Commit,
			ModuleName:                moduleName,
			BuildDate:                 buildinfo.BuildDate().Format("2006-01-02T15:04:05.000Z"),
			Version:                   "unknown",
			Branch:                    buildinfo.Branch,
			Modules:                   nil,
		})
	}
}
