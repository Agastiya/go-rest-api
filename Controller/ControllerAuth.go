package Controller

import (
	"encoding/json"
	"go-rest-api/Constant"
	"go-rest-api/Dto"
	"go-rest-api/Helper/Response"
	"go-rest-api/Services"
	"net/http"
)

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func (c Controller) Login(w http.ResponseWriter, r *http.Request) {

	var params Dto.Login
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	result := Services.Auth.Login(params)
	if result.HasErr {
		Response.ResponseError(w, result.Err, Constant.InternalCode(result.InternalCode))
		return
	}

	Response.ResponseSuccess(w, result.Result, Constant.StatusOKJson)
}
