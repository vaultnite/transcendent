package controllers

import (
	"net/http"
)

type DataRouterController struct {
	*http.ServeMux
}

func NewDataRouterController() *DataRouterController {
	c := &DataRouterController{ServeMux: http.NewServeMux()}
	c.addRoutes()
	return c
}

func (c *DataRouterController) addRoutes() {
	c.HandleFunc("POST /api/v1/public/data", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
