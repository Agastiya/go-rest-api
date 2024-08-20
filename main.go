package main

import (
	"go-rest-api/App"
)

// @title           			Go Rest API
// @version         			1.0
// @description     			Go rest API Skeleton
// @host      					localhost:7000
// @BasePath  					/skeleton
// @securityDefinitions.apikey  Bearer
// @in 							Header
// @name 						Authorization
func main() {
	App.AppInitialization()
}
