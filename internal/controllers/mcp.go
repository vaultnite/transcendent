package controllers

import (
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io"
    "net/http"
    "os"
    "path"
    "time"

    "github.com/vaultnite/transcendent/internal/errs"
    "github.com/vaultnite/transcendent/internal/models"
    "github.com/vaultnite/transcendent/internal/storage"
)

type mcpController struct {
    *http.ServeMux

    sp        storage.Provider
    serviceId string
}

func newMCPController(sp storage.Provider, serviceId string) *mcpController {
    c := &mcpController{
        ServeMux:  http.NewServeMux(),
        sp:        sp,
        serviceId: serviceId,
    }
    c.addRoutes()
    return c
}

func (c *mcpController) addRoutes() {
    c.HandleFunc("GET /api/calendar/v1/timeline", func(w http.ResponseWriter, r *http.Request) {
        data, err := c.sp.ReadFile(path.Join("mcp", c.serviceId, "timeline.json"))
        if err != nil {
            res := errs.NewInternalServerError(c.serviceId)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(&res)
            return
        }

        w.Header().Add("Content-Type", "application/json")
        w.Write(data)
    })

    c.HandleFunc("GET /api/cloudstorage/system", func(w http.ResponseWriter, r *http.Request) {
        files, err := c.sp.ReadDir(path.Join("mcp", c.serviceId, "hotfixes"))
        if err != nil {
            res := errs.NewInternalServerError(c.serviceId)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(&res)
            return
        }

        result := make([]models.CloudstorageFilePointerSystem, len(files), len(files))

        for i, f := range files {
            file, err := c.sp.Open(path.Join("mcp", c.serviceId, "hotfixes", f.Name()))
            if err != nil {
                res := errs.NewInternalServerError(c.serviceId)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(&res)
                return
            }
            fileInfo, err := f.Info()
            if err != nil {
                res := errs.NewInternalServerError(c.serviceId)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(&res)
                return
            }

            hash1 := sha1.New()
            _, err = io.Copy(hash1, file)
            if err != nil {
                res := errs.NewInternalServerError(c.serviceId)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(&res)
                return
            }

            hash256 := sha256.New()
            _, err = io.Copy(hash256, file)
            if err != nil {
                res := errs.NewInternalServerError(c.serviceId)
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(&res)
                return
            }

            result[i] = models.CloudstorageFilePointerSystem{
                CloudstorageFilePointer: models.CloudstorageFilePointer{
                    UniqueFilename: f.Name(),
                    Filename:       f.Name(),
                    Hash:           hex.EncodeToString(hash1.Sum(nil)),
                    Hash256:        hex.EncodeToString(hash256.Sum(nil)),
                    Length:         int(fileInfo.Size()),
                    ContentType:    "application/octet-stream",
                    Uploaded:       fileInfo.ModTime().UTC().Format("2006-01-02T15:04:05.999Z"),
                    StorageType:    "S3",
                    StorageIDs:     nil,
                },
                DoNotCache: false,
            }
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&result)
    })

    c.HandleFunc("GET /api/cloudstorage/system/config", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.CloudstorageConfig{
            LastUpdated:        time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
            DisableV2:          false,
            IsAuthenticated:    true,
            EnumerateFilesPath: "/api/cloudstorage/system",
            EnableMigration:    false,
            EnableWrites:       false,
            EpicAppName:        "Live",
        })
    })

    c.HandleFunc("GET /api/cloudstorage/system/{uniqueFilename}", func(w http.ResponseWriter, r *http.Request) {
        uniqueFilename := r.PathValue("uniqueFilename")
        fileData, err := c.sp.ReadFile(path.Join("mcp", c.serviceId, "hotfixes", uniqueFilename))
        if err != nil {
            res := errs.NewInternalServerError(c.serviceId)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(&res)
            return
        }

        w.Header().Set("Content-Type", "application/octet-stream")
        w.Write(fileData)
    })

    c.HandleFunc("GET /api/cloudstorage/user/config", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.CloudstorageConfig{
            LastUpdated:        time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
            DisableV2:          false,
            IsAuthenticated:    true,
            EnumerateFilesPath: "/api/cloudstorage/user",
            EnableMigration:    false,
            EnableWrites:       false,
            EpicAppName:        "Live",
        })
    })

    c.HandleFunc("GET /api/cloudstorage/user/{accountId}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("[]"))
    })

    c.HandleFunc("GET /api/game/v2/enabled_features", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("[]"))
    })

    c.HandleFunc("GET /api/game/v2/grant_access/{accountId}", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    })

    c.HandleFunc("POST /api/game/v2/profile/{accountId}/client/{command}", func(w http.ResponseWriter, r *http.Request) {
        //        accountId := r.PathValue("accountId")

        w.Header().Set("Content-Type", "application/json")
        //        profile := bootstrapProfile("profile0", accountId)
        //        profile.Touch()
        //        json.NewEncoder(w).Encode(profile)

        data, err := os.ReadFile("E:\\FORTNITE MODDING ARCHIVE\\LawinServer\\profiles\\profile0.json")
        if err != nil {
            panic(err)
        }
        w.Write(data)
    })

    c.HandleFunc("POST /api/game/v2/tryPlayOnPlatform/account/{accountId}", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("true"))
    })

    c.HandleFunc("GET /api/matchmaking/session/findPlayer/{accountId}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("[]"))
    })

    c.HandleFunc("POST /api/matchmaking/session/matchMakingRequest", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        res := make([]models.MatchmakingRequestResponseWIP, 1, 1)
        res[0] = models.MatchmakingRequestResponseWIP{
            ServerAddress: "127.0.0.1",
            ServerPort:    80,
            Attributes:    []byte("{}"),
        }
        json.NewEncoder(w).Encode(&res)
    })

    c.HandleFunc("GET /api/receipts/v1/account/{accountId}/receipts", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&[]models.Receipt{})
    })

    // older versioncheck, superseeded by at least fn 4.00, present in 2017 orion and possibly ot fn builds (unsure the the response body is the same shape, but using for now)
    c.HandleFunc("GET /api/versioncheck", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.VersionCheckResponse{
            Type: models.VersionCheckResponseTypeNoUpdate,
        })
    })

    c.HandleFunc("GET /api/v2/versioncheck/{platform}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.VersionCheckResponse{
            Type: models.VersionCheckResponseTypeNoUpdate,
        })
    })
}

func bootstrapProfile(profileId, accountId string) *models.Profile {
    nowStr := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")

    changeProfile := models.ProfileChangeProfileEntry{
        Created:   nowStr,
        AccountID: accountId,
        ProfileID: profileId,
        Items:     map[string]models.ProfileItem{},
        Stats: models.ProfileStats{
            Attributes: []byte("{}"),
        },
    }

    changes := make([]models.ProfileChange, 1)
    changes[0] = models.ProfileChange{
        EnableConstructDelta: true,
        Profile:              changeProfile,
    }

    profile := &models.Profile{
        ProfileID:       profileId,
        ProfileChanges:  changes,
        CreationTime:    nowStr,
        ResponseVersion: 1,
    }

    return profile
}
