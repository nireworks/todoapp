package server

import (
	"log"
	"net/http"
	"time"
)

func (s *Server) mwLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

func (s *Server) mwTimer(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(time.Since(start))
	})
}
