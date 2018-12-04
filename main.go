package main

import (
	"merchants/service"
	"os"
)

func main() {
	//some server need environment variable for listen&serve.
	os.Setenv("PORT", "8000")

	s := service.CreateService()
	r := service.SetupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
