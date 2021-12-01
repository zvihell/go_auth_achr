package main

import (
	"go-auth/database"
	"go-auth/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	database.InitDB()

	r := routes.SetupRouter()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	}))

	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//config.AllowCredentials = true

	//r.Use(cors.New(cors.Config{
	//AllowAllOrigins:  true,
	//AllowCredentials: true,
	//}))

	r.Run(":8000")
}
