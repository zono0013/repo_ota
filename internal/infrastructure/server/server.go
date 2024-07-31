package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"repo/internal/interfaces/handlers"
)

type Server struct {
	router *mux.Router
}

func NewServer(userHandler *handlers.UserHandler, reportHandler *handlers.ReportHandler) *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.setupRoutes(userHandler, reportHandler)
	return s
}

func (s *Server) setupRoutes(userHandler *handlers.UserHandler, reportHandler *handlers.ReportHandler) {
	// User routes
	s.router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	s.router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	s.router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")

	// Report routes
	s.router.HandleFunc("/reports", reportHandler.GetAllReports).Methods("GET")
	s.router.HandleFunc("/reports", reportHandler.CreateReport).Methods("POST")
	s.router.HandleFunc("/reports/{id}", reportHandler.GetReport).Methods("GET")
	s.router.HandleFunc("/reports/{id}", reportHandler.UpdateReport).Methods("PUT")
	s.router.HandleFunc("/reports/{id}", reportHandler.DeleteReport).Methods("DELETE")

	// ミドルウェアの追加
	s.router.Use(loggingMiddleware)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
