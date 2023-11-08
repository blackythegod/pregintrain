package webapi

import (
	"gintraining/handler"
	"log"
	"net/http"
)

type Server struct {
	Handler *handler.Handler
}

func (s *Server) ServerRun() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: s.Handler.Router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func InitServer() *Server {
	h := handler.InitHandler()
	return &Server{
		Handler: h,
	}
}
