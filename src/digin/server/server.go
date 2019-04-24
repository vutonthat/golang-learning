package server

import (
	"digin/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ginEngine       *gin.Engine
	publicIPService services.PublicIPService
}

func New(publicIPService services.PublicIPService) *Server {
	s := Server{
		ginEngine:       gin.New(),
		publicIPService: publicIPService,
	}
	defer s.router()
	return &s
}

func (s *Server) Run() error {
	return s.ginEngine.Run()
}

func (s *Server) Close() {}

func (s *Server) router() {
	s.publicIPHandler()
}

func (s *Server) publicIPHandler() {
	h := publicIPHandler{
		ginEngine:       s.ginEngine,
		publicIPService: s.publicIPService,
	}
	h.router()
}
