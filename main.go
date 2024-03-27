package main

import (
	"github.com/Taliker/BootDevHTTP/api"
	"log"
	"net/http"
)

const (
	// PORT is the port number
	PORT = ":8080"
	// DIR is the directory of the static files
	DIR = "./static"
)

func main() {
	apiCfg := &api.ApiConfig{}
	mux := http.NewServeMux()
	mux.Handle("GET /app/", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(DIR)))))
	mux.HandleFunc("GET /metrics", apiCfg.MiddlewareMetrics)
	mux.HandleFunc("GET /reset", apiCfg.MiddlewareMetricsReset)
	mux.HandleFunc("GET /healthz", api.HealthHandler)
	corsMux := middlewareCORS(mux)
	srv := &http.Server{
		Addr:    PORT,
		Handler: corsMux,
	}
	log.Fatal(srv.ListenAndServe())
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
