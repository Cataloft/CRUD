package server

import (
	"awesomeProject/internal/handler"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e       *echo.Echo
	handler *handler.Handler
}

func New(handler *handler.Handler) *Server {
	var e *echo.Echo
	e = echo.New()

	return &Server{
		e:       e,
		handler: handler,
	}
}

func (s *Server) initHandlers() {
	s.e.POST("/register", s.handler.Register)
	s.e.GET("/getUser", s.handler.GetUser)
	s.e.GET("/updateUser", s.handler.UpdateUser)
	s.e.GET("/deleteUser", s.handler.DeleteUser)
	s.e.POST("/createBlog", s.handler.CreateBlog)
	s.e.GET("/getBlog", s.handler.GetBlog)
	s.e.GET("/getBlogOwner", s.handler.GetBlogOwner)
}

func (s *Server) Start() error {
	s.initHandlers()

	return s.e.Start(":8080")
}
