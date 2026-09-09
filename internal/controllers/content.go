package controllers

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "os"
    "path"

    "github.com/vaultnite/transcendent/internal/errs"
    "github.com/vaultnite/transcendent/internal/storage"
)

type ContentController struct {
    *http.ServeMux

    sp storage.Provider
}

func NewContentController(sp storage.Provider) *ContentController {
    c := &ContentController{
        ServeMux: http.NewServeMux(),
        sp:       sp,
    }
    c.addRoutes()
    return c
}

func (c *ContentController) addRoutes() {
    c.HandleFunc("GET /api/pages/{contentKey}", func(w http.ResponseWriter, r *http.Request) {
        contentKey := r.PathValue("contentKey")
        fileData, err := c.sp.ReadFile(path.Join("content", fmt.Sprintf("%s.json", contentKey)))
        if err != nil {
            if errors.Is(err, os.ErrNotExist) {
                res := errs.NewGenericHTTPError(http.StatusNotFound, "com.transcendent.common.404")
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusNotFound)
                json.NewEncoder(w).Encode(&res)
            } else {
                res := errs.NewInternalServerError("content-common")
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(&res)
            }
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(fileData) // idk if i should verify the content is json or just pass it along like this...
    })

    c.HandleFunc("GET /api/pages/{contentKey}/{contentSubKey}", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusTeapot)
    })
}
