package Services

import (
	"errors"
	"go-rest-api/Constant"
	"go-rest-api/Dto"
	"go-rest-api/Helper/Jwt"
	"go-rest-api/Helper/Response"
	"go-rest-api/Helper/Validator"
	"go-rest-api/Repository/auth"
	"strings"
)

type AuthInterface interface {
	Login(params Dto.Login) (result Response.RespResultService)
}

func (c authService) Login(params Dto.Login) (result Response.RespResultService) {
	if errValidator := Validator.Validate(params); errValidator != nil {
		return Response.ResponseService(true, errValidator, Constant.StatusBadRequestInvalidData, nil, nil)
	}

	resultData, err := repoAuth.Login(params)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return Response.ResponseService(true, errors.New("email or password incorrect"), Constant.StatusBadRequestNotExist, nil, nil)
		}
		return Response.ResponseService(true, err, Constant.StatusInternalServerErrorDB, nil, nil)
	}

	if !resultData.Active {
		return Response.ResponseService(true, errors.New("this account is inactive"), Constant.StatusBadRequestNotExist, nil, nil)
	}

	resultData.AccessToken, err = Jwt.JwtVar.CreateToken(auth.LoginData{
		Id:       resultData.Id,
		Name:     resultData.Name,
		Username: resultData.Username,
		Email:    resultData.Email,
		Active:   resultData.Active,
	})
	if err != nil {
		return Response.ResponseService(true, err, Constant.StatusInternalServerError, nil, nil)
	}

	return Response.ResponseService(false, nil, Constant.StatusOKJson, nil, resultData)
}
