package Services

import (
	"errors"
	"go-rest-api/Constant"
	"go-rest-api/Dto"
	"go-rest-api/Helper/Response"
	"go-rest-api/Helper/Validator"
)

type AccountInterface interface {
	CreateAccount(params Dto.CreateAccount) (result Response.RespResultService)
	UpdateAccountPassword(params Dto.UpdateAccountPassword) (result Response.RespResultService)
	UpdateAccountStatus(params Dto.UpdateAccountStatus) (result Response.RespResultService)
}

func (c authService) CreateAccount(params Dto.CreateAccount) (result Response.RespResultService) {

	// Validate Email
	emailExists, err := repoAuth.ValidateAccountData("email", params.Email)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}

	if emailExists {
		return Response.ResponseService(true, errors.New("email already exists"), Constant.StatusBadRequestAlreadyExists, nil, nil)
	}

	// Validate Username
	usernameExists, err := repoAuth.ValidateAccountData("username", params.Username)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}

	if usernameExists {
		return Response.ResponseService(true, errors.New("username already exists"), Constant.StatusBadRequestAlreadyExists, nil, nil)
	}

	// Validate Phone
	phoneExists, err := repoAuth.ValidateAccountData("phone", params.Phone)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}

	if phoneExists {
		return Response.ResponseService(true, errors.New("phone already exists"), Constant.StatusBadRequestAlreadyExists, nil, nil)
	}

	if errValidator := Validator.Validate(params); errValidator != nil {
		return Response.ResponseService(true, errValidator, Constant.StatusBadRequestInvalidData, nil, nil)
	}

	err = repoAuth.CreateAccount(params)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}
	return Response.ResponseService(false, nil, Constant.StatusOKJson, nil, nil)
}

func (c authService) UpdateAccountPassword(params Dto.UpdateAccountPassword) (result Response.RespResultService) {

	if errValidator := Validator.Validate(params); errValidator != nil {
		return Response.ResponseService(true, errValidator, Constant.StatusBadRequestInvalidData, nil, nil)
	}

	err := repoAuth.UpdateAccountPassword(params)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}
	return Response.ResponseService(false, nil, Constant.StatusOKJson, nil, nil)
}

func (c authService) UpdateAccountStatus(params Dto.UpdateAccountStatus) (result Response.RespResultService) {

	if errValidator := Validator.Validate(params); errValidator != nil {
		return Response.ResponseService(true, errValidator, Constant.StatusBadRequestInvalidData, nil, nil)
	}

	err := repoAuth.UpdateAccountStatus(params)
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}
	return Response.ResponseService(false, nil, Constant.StatusOKJson, nil, nil)
}
