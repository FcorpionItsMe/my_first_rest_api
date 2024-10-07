package my_first_rest_api

import (
	"github.com/FcorpionItsMe/my_first_rest_api/pkg"
	"log"
)

func Run() {
	srv := pkg.Server{}
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occurred while start server: %s", err.Error())
	}
}
