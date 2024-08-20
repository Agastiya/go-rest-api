package App

import (
	"flag"
	"go-rest-api/Config"
	"go-rest-api/Constant"
	"go-rest-api/Controller"
	"go-rest-api/Helper/Jwt"
	"go-rest-api/Routes"
	"go-rest-api/Routes/Middleware"
	"strings"

	"github.com/go-chi/chi/v5"
)

var environment = flag.String("tag", "", "define tag")

func init() {
	flag.Parse()

	if strings.Contains(*environment, "DEV-") {
		*environment = Constant.Development
	} else if strings.Contains(*environment, "PROD-") {
		*environment = Constant.Production
	} else {
		*environment = Constant.Local
	}
}

func ServiceInit(env *Config.Environment) {
	Jwt.JwtVar = &Jwt.JwtService{ConfigJwt: env.Jwt}
}

func AppInitialization() {
	environment := Config.GetEnvironment(*environment)
	ServiceInit(&environment.Environment)

	environment.Database.BuildConnection()
	environment.Routes = &Routes.Routes{
		Chi:        chi.NewRouter(),
		Env:        environment.Environment.App.Environment,
		Controller: &Controller.Controller{},
		Middleware: &Middleware.Middleware{
			Jwt:            Jwt.JwtVar,
			SwaggerSetting: environment.Environment.Swagger,
		},
	}
	environment.Engine.ServeHTTP(environment.Routes.CollectRoutes())
}
