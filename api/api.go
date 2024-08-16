package api

import (
	"github.com/gin-gonic/gin"
	"github.com/katsuokaisao/gin-play/api/handler"
	"github.com/katsuokaisao/gin-play/api/middleware"
	"github.com/katsuokaisao/gin-play/domain"
)

type Server struct {
	addr        string
	router      *gin.Engine
	jwtParser   domain.JWTParser
	todoHandler *handler.TodoHandler
}

func NewServer(
	apiEnv *domain.APIEnv,
	jwtParser domain.JWTParser,
	todoHandler *handler.TodoHandler,
) *Server {
	router := gin.Default()
	server := &Server{
		addr:        apiEnv.Addr,
		router:      router,
		jwtParser:   jwtParser,
		todoHandler: todoHandler,
	}
	return server
}

func (s *Server) RegisterRoutes() {
	todoR := s.router.Group("/todos")
	todoRW := s.router.Group("/todos")
	todoR.Use(middleware.JWTMiddleware(s.jwtParser, domain.ScopeTodoRead))
	todoRW.Use(middleware.JWTMiddleware(s.jwtParser, domain.ScopeTodoReadWrite))
	todoR.GET("", s.todoHandler.List)
	todoR.GET("/:id", s.todoHandler.Get)
	todoRW.POST("", s.todoHandler.Create)
	todoRW.PUT("/:id", s.todoHandler.Update)
	todoRW.DELETE("/:id", s.todoHandler.Delete)
}

func (s *Server) Run() error {
	return s.router.Run(s.addr)
}
