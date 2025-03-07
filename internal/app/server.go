package app

import "github.com/labstack/echo/v4"

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	//представляет собой веб-сервер
	e := echo.New()
	return &Server{
		echo: e,
	}
}

func (s *Server) Start() error {
	return s.echo.Start("localhost:8080")
}

//  Регистрируем обработчик
func (s *Server) AddHandler(path string, handler echo.HandlerFunc) {
	s.echo.GET(path, handler)
}

// Регистрируем обработчик
func (s *Server) PostHandler(path string, handler echo.HandlerFunc) {
	s.echo.POST(path, handler)
}
