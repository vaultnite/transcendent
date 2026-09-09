package controllers

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/vaultnite/transcendent/internal/errs"
    "github.com/vaultnite/transcendent/internal/middleware"
    "github.com/vaultnite/transcendent/internal/models"
)

type AccountController struct {
    *http.ServeMux

    db *sql.DB
}

func NewAccountController(db *sql.DB) *AccountController {
    c := &AccountController{
        ServeMux: http.NewServeMux(),
        db:       db,
    }
    c.addRoutes()
    return c
}

func (c *AccountController) addRoutes() {
    c.HandleFunc("GET /api/version", middleware.NewVersionHandler("com.transcendent.account.public", "transcendent", "Transcendent-Account-Controller"))

    c.HandleFunc("POST /api/oauth/token", func(w http.ResponseWriter, r *http.Request) {
        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        now := time.Now().UTC()
        grantType := r.FormValue("grant_type")

        switch grantType {
        case "client_credentials":
            accessExpiresIn := time.Hour * 4

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(&models.ClientTokenResponse{
                AccessToken:    "{{vaultnite-access}}",
                ExpiresIn:      int64(accessExpiresIn.Seconds()),
                ExpiresAt:      now.Add(accessExpiresIn).Format("2006-01-02T15:04:05.000Z"),
                TokenType:      "bearer",
                ClientID:       "{{vaultnite-client}}",
                InternalClient: true,
                ClientService:  "fortnite",
                ProductID:      "",
                ApplicationID:  "",
            })
            return
        case "authorization_code":
        case "device_auth":
        case "device_code":
        case "exchange_code":
        case "external_auth":
        case "otp":
        case "password":
        case "refresh_token":
        case "token_to_token":
        default:
            res := errs.NewUnsupportedGrantTypeError("account-public", grantType)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(&res)
            return
        }

        accessExpiresIn := time.Hour * 8
        refreshExpiresIn := time.Hour * 32

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.TokenResponse{
            AccessToken:      "{{vaultnite-access}}",
            ExpiresIn:        int64(accessExpiresIn.Seconds()),
            ExpiresAt:        now.Add(accessExpiresIn).Format("2006-01-02T15:04:05.000Z"),
            TokenType:        "bearer",
            RefreshToken:     "{{vaultnite-refresh}}",
            RefreshExpires:   int64(refreshExpiresIn.Seconds()),
            RefreshExpiresAt: now.Add(refreshExpiresIn).Format("2006-01-02T15:04:05.000Z"),
            AccountID:        "{{vaultnite-user}}",
            ClientID:         "{{vaultnite-client}}",
            InternalClient:   true,
            ClientService:    "fortnite",
            DisplayName:      "{{vaultnite-user}}",
            App:              "fortnite",
            InAppID:          "{{vaultnite-user}}",
            ProductID:        "",
            ApplicationID:    "",
        })
    })

    c.HandleFunc("GET /api/oauth/verify", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
    })

    c.HandleFunc("GET /api/public/account/{accountId}", func(w http.ResponseWriter, r *http.Request) {
        accountId := r.PathValue("accountId")

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&models.PublicAccountLookupResponse{
            ID:          accountId,
            DisplayName: "nick",
        })
    })

    c.HandleFunc("GET /api/public/account/{accountId}/externalAuths", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(&[]models.AccountExternalAuth{})
    })

    c.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
        log.Println("reserve")
    })
}
