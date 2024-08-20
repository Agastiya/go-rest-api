package Middleware

import (
	"go-rest-api/Config"
	"go-rest-api/Helper/Jwt"
)

type Middleware struct {
	MiddlewareInterface
	Jwt            Jwt.JwtInterface
	SwaggerSetting Config.SwaggerSetting
}
