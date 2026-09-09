package main

import (
    "log"
    "log/slog"
    "net/http"
    "os"
    "path/filepath"

    "github.com/vaultnite/transcendent/internal/buildinfo"
    "github.com/vaultnite/transcendent/internal/controllers"
    "github.com/vaultnite/transcendent/internal/database"
    "github.com/vaultnite/transcendent/internal/middleware"
    "github.com/vaultnite/transcendent/internal/router"
    "github.com/vaultnite/transcendent/internal/storage"
)

// unset by default so dev run can override to localappdata & in build will be set to exPath
var (
    rootDir = ""
)

func main() {
    slog.Info(buildinfo.String())

    if rootDir == "" {
        ex, err := os.Executable()
        if err != nil {
            panic(err)
        }
        rootDir = filepath.Dir(ex)
    }

    dbPath := filepath.Join(rootDir, "data.sqlite")

    db, err := database.Open(dbPath)
    if err != nil {
        log.Fatal("Error opening database:", err)
    }
    defer db.Close()

    if err = database.Migrate(db); err != nil {
        log.Fatal(err)
    }

    sp := storage.NewStorageProviderLocal(filepath.Join(rootDir))

    mux := router.NewRouter()

    mux.Use(middleware.Logger)
    mux.HandleGroup("/account", controllers.NewAccountController(db))
    mux.HandleGroup("/caldera", controllers.NewCalderaController())
    mux.HandleGroup("/content", controllers.NewContentController(sp))
    mux.HandleGroup("/datarouter", controllers.NewDataRouterController())
    mux.HandleGroup("/eulatracking", controllers.NewEULATrackingController())
    mux.HandleGroup("/fortnite", controllers.NewFortniteController(sp))
    mux.HandleGroup("/friends", controllers.NewFriendsController())
    mux.HandleGroup("/lightswitch", controllers.NewLightSwitchController())
    mux.HandleGroup("/orion", controllers.NewOrionController(sp)) // paragon
    mux.HandleGroup("/waitingroom", controllers.NewWaitingRoomController())
    mux.HandleGroup("/wex", controllers.NewWexController(sp)) // battle breakers

    slog.Info("Transcendent starting on localhost:3551")
    log.Fatal(http.ListenAndServe("localhost:3551", mux))
}
