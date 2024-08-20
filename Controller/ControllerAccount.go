package Controller

import (
	"encoding/json"
	"go-rest-api/Constant"
	"go-rest-api/Dto"
	"go-rest-api/Helper/Response"
	"go-rest-api/Helper/Utils"
	"go-rest-api/Services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Account interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	UpdateAccountPassword(w http.ResponseWriter, r *http.Request)
	UpdateAccountStatus(w http.ResponseWriter, r *http.Request)
}

func (c Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var params Dto.CreateAccount
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	ctxSetData, _, err := Utils.SetValueContext(r)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	params.CreatedBy = ctxSetData.Name
	params.CreatedTime = Utils.TimeNow()

	result := Services.Account.CreateAccount(params)
	if result.HasErr {
		Response.ResponseError(w, result.Err, Constant.InternalCode(result.InternalCode))
		return
	}

	Response.ResponseSuccess(w, nil, Constant.StatusOKJson)
}

func (c Controller) UpdateAccountPassword(w http.ResponseWriter, r *http.Request) {

	var params Dto.UpdateAccountPassword
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	id := chi.URLParam(r, "id")
	params.Id, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result := Services.Account.UpdateAccountPassword(params)
	if result.HasErr {
		Response.ResponseError(w, result.Err, Constant.InternalCode(result.InternalCode))
		return
	}

	Response.ResponseSuccess(w, nil, Constant.StatusOKJson)
}

func (c Controller) UpdateAccountStatus(w http.ResponseWriter, r *http.Request) {

	var params Dto.UpdateAccountStatus
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	id := chi.URLParam(r, "id")
	params.Id, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result := Services.Account.UpdateAccountStatus(params)
	if result.HasErr {
		Response.ResponseError(w, result.Err, Constant.InternalCode(result.InternalCode))
		return
	}

	Response.ResponseSuccess(w, nil, Constant.StatusOKJson)
}
