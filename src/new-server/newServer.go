package new_server

import (
	"net/http"

	"github.com/AKiranB/linkedin-scraper-microservice/src/routes"
	"github.com/AKiranB/linkedin-scraper-microservice/src/utils"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.AddHeaders(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", routes.HealthHandler())
	mux.HandleFunc("/test", nil)
	mux.HandleFunc("/jobs", nil)
}

func NewServer() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	return corsMiddleware(mux)
}
