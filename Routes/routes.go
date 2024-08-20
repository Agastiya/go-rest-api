package Routes

import (
	"go-rest-api/Controller"
	"go-rest-api/Routes/Middleware"

	_ "go-rest-api/Library/Swagger/docs"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routes struct {
	Chi        *chi.Mux
	Env        string
	Middleware Middleware.MiddlewareInterface
	Controller Controller.ControllerInterface
}

func (app *Routes) CollectRoutes() *chi.Mux {

	appRoute := app.Chi
	appRoute.Use(chiMiddleware.RequestID)
	appRoute.Use(chiMiddleware.RealIP)
	appRoute.Use(chiMiddleware.RedirectSlashes)
	appRoute.Use(chiMiddleware.Recoverer)

	appRoute.Route("/skeleton", func(appRoute chi.Router) {
		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Route("/auth", func(appRoute chi.Router) {
				appRoute.Post("/login", app.Controller.Login)
			})
		})

		appRoute.Group(func(appRoute chi.Router) {
			appRoute.Use(app.Middleware.UserAuth())
			appRoute.Route("/account", func(appRoute chi.Router) {
				appRoute.Post("/", app.Controller.CreateAccount)
				appRoute.Route("/{id}", func(appRoute chi.Router) {
					appRoute.Put("/", app.Controller.UpdateAccountPassword)
					appRoute.Put("/active", app.Controller.UpdateAccountStatus)
				})
			})
		})

		// use basic auth or mount swagger based on specific environment
		if app.Env == "development" {
			appRoute.Group(func(appRoute chi.Router) {
				appRoute.Use(app.Middleware.BasicAuthSwagger())
				appRoute.Mount("/swagger", httpSwagger.WrapHandler)
			})
		} else if app.Env == "local" {
			appRoute.Mount("/swagger", httpSwagger.WrapHandler)
		}

		//sample controller function
		appRoute.Get("/ping", app.Controller.Ping)
	})

	return appRoute
}
