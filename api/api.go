package api

import (
	"errors"
	"fmt"
	"net/http"
)

type ApiConfig struct {
	FileserverHits int
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (a *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.FileserverHits++
		next.ServeHTTP(w, r)
	})
}

func (a *ApiConfig) MiddlewareMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits: " + fmt.Sprint(a.FileserverHits)))
}

func (a *ApiConfig) MiddlewareMetricsReset(w http.ResponseWriter, r *http.Request) {
	if a == nil {
		panic(errors.New("ApiConfig is nil"))
	}
	a.FileserverHits = 0
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
