package main

import (
	"merchants/service"
	"os"
)

func main() {

	// os.Setenv("DATABASE_URL", "mongodb://lordgift:gotraining@database:27017/merchants")
	os.Setenv("PORT", "8000")

	// service.CreateService()

	s := service.CreateService()
	r := service.SetupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
