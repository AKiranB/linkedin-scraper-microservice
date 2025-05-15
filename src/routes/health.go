package routes

import (
	"net/http"

	"github.com/AKiranB/linkedin-scraper-microservice/src/utils"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Hitting Health Handler Success")
		utils.Encode(w, http.StatusOK, "Healthy")
	}
}
