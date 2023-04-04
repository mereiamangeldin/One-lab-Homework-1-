package http

import "github.com/gorilla/mux"

func (s *Server) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/users", s.handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users", s.handler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", s.handler.GetUserById).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", s.handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/v1/users/{id}", s.handler.UpdateUser).Methods("PUT")
	return r
}
