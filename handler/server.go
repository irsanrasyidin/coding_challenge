package handler

import (
	"coding_challenge/config"
	"coding_challenge/manager"
	"coding_challenge/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	s.srv.Use(middleware.LoggerMiddleware())

	NewRentHandler(s.srv, s.usecaseManager.GetRentUsecase())
	NewCarsHandler(s.srv, s.usecaseManager.GetCarsUsecase())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.Newconfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	if c.DbConfig.Host == "" || c.AppPort == "" {
		panic("No Host or port define")
	}

	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
