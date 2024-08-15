package api

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuokaisao/gin-play/api/handler"
	"github.com/katsuokaisao/gin-play/domain"
)

type Server struct {
	addr        string
	router      *gin.Engine
	todoHandler *handler.TodoHandler
}

func NewServer(
	apiEnv *domain.APIEnv,
	todoHandler *handler.TodoHandler,
) *Server {
	router := gin.Default()
	server := &Server{
		addr:        apiEnv.Addr,
		router:      router,
		todoHandler: todoHandler,
	}
	return server
}

func (s *Server) RegisterRoutes() {
	s.router.POST("/todos", s.todoHandler.Create)
	s.router.GET("/todos", s.todoHandler.List)
	s.router.GET("/todos/:id", s.todoHandler.Get)
	s.router.PUT("/todos/:id", s.todoHandler.Update)
	s.router.DELETE("/todos/:id", s.todoHandler.Delete)
}

func (s *Server) Run() error {
	return s.router.Run(s.addr)
}
