package Controller

import (
	"go-rest-api/Constant"
	"go-rest-api/Helper/Response"
	"net/http"
)

type ControllerService interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

// @Tags     Ping
// @Accept   json
// @Produce  json
// @Router   /ping [get]
func (c Controller) Ping(w http.ResponseWriter, r *http.Request) {
	Response.ResponseSuccess(w, nil, Constant.StatusOKJson)
}
