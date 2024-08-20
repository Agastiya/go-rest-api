package Constant

import (
	"net/http"
)

type (
	InternalCode int

	Response struct {
		HttpCode    int
		HttpTitle   string
		Description string
	}
)

const (

	// 2xx
	StatusOKJson            InternalCode = 2000
	StatusOKAccountApproved InternalCode = 2001
	StatusOKValidationOK    InternalCode = 2002
	StatusOKNoRowAffected   InternalCode = 2003

	// 4xx
	StatusBadRequestJson             InternalCode = 4001
	StatusBadRequestNotExist         InternalCode = 4002
	StatusBadRequestAlreadyExists    InternalCode = 4003
	StatusBadRequestInvalidData      InternalCode = 4004
	StatusBadRequestInvalidParameter InternalCode = 4005
	// > 401
	StatusUnauthorizedInvalidToken   InternalCode = 4011
	StatusUnauthorizedTokenNotExists InternalCode = 4012
	StatusUnauthorizedBearerNotFound InternalCode = 4013
	StatusUnauthorizedErrorVerifying InternalCode = 4014

	//5xx
	StatusInternalServerError   InternalCode = 5000
	StatusInternalServerErrorDB InternalCode = 5002
)

func (i InternalCode) is2xx() bool {
	return i >= 2000 && i < 3000
}

func (i InternalCode) is4xx() bool {
	return i >= 4000 && i < 5000
}

func (i InternalCode) is5xx() bool {
	return i >= 5000 && i < 6000
}

func (i InternalCode) Response() Response {

	if i.is2xx() {
		return response2xx[i]
	} else if i.is4xx() {
		return response4xx[i]
	} else if i.is5xx() {
		return response5xx[i]
	} else {
		return Response{}
	}
}

var response2xx = map[InternalCode]Response{
	StatusOKJson: {
		HttpCode:    http.StatusOK,
		HttpTitle:   http.StatusText(http.StatusOK),
		Description: "Success",
	},
	StatusOKAccountApproved: {
		HttpCode:    http.StatusOK,
		HttpTitle:   http.StatusText(http.StatusOK),
		Description: "Account Approved",
	},
	StatusOKValidationOK: {
		HttpCode:    http.StatusOK,
		HttpTitle:   http.StatusText(http.StatusOK),
		Description: "Validation Successful",
	},
	StatusOKNoRowAffected: {
		HttpCode:    http.StatusOK,
		HttpTitle:   http.StatusText(http.StatusOK),
		Description: "No Row Affected",
	},
}

var response4xx = map[InternalCode]Response{
	StatusBadRequestJson: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "json error",
	},
	StatusBadRequestNotExist: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "data not exists",
	},
	StatusBadRequestAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "data already exists",
	},
	StatusBadRequestInvalidData: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "invalid parameter",
	},
	StatusBadRequestInvalidParameter: {
		HttpCode:    http.StatusBadRequest,
		HttpTitle:   http.StatusText(http.StatusBadRequest),
		Description: "invalid parameter",
	},
	StatusUnauthorizedInvalidToken: {
		HttpCode:    http.StatusUnauthorized,
		HttpTitle:   http.StatusText(http.StatusUnauthorized),
		Description: "invalid bearer token",
	},
	StatusUnauthorizedTokenNotExists: {
		HttpCode:    http.StatusUnauthorized,
		HttpTitle:   http.StatusText(http.StatusUnauthorized),
		Description: "token not exists",
	},
	StatusUnauthorizedBearerNotFound: {
		HttpCode:    http.StatusUnauthorized,
		HttpTitle:   http.StatusText(http.StatusUnauthorized),
		Description: "Error verifying JWT token: 'Bearer ' Not Found",
	},
	StatusUnauthorizedErrorVerifying: {
		HttpCode:    http.StatusUnauthorized,
		HttpTitle:   http.StatusText(http.StatusUnauthorized),
		Description: "Error verifying JWT token",
	},
}

var response5xx = map[InternalCode]Response{
	StatusInternalServerError: {
		HttpCode:    http.StatusInternalServerError,
		HttpTitle:   http.StatusText(http.StatusInternalServerError),
		Description: "internal server error (System)",
	},
	StatusInternalServerErrorDB: {
		HttpCode:    http.StatusBadGateway,
		HttpTitle:   http.StatusText(http.StatusBadGateway),
		Description: "internal server error (DB)",
	},
}
