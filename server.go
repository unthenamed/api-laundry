package main

import (
	"api-laundry/config"
	"api-laundry/controller"
	"api-laundry/repo"
	"api-laundry/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	service service.LaundryService
	host    string
	engine  *gin.Engine
}

func (s *Server) initRoute() {
	group := s.engine.Group("/")
	controller.NewHandlersController(group, s.service).Route()

}
func (s *Server) Run() {
	s.initRoute()
	s.engine.Run(s.host)
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	sql.Open(cfg.Database.Driver, cfg.Database.DSN)
	db, err := sql.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Server{
		host:   ":" + cfg.Api.Port,
		engine: gin.Default(),
		service: service.NewLaundryService(
			repo.ObjTransactionRepo(db),
			repo.ObjProductRepo(db),
			repo.ObjCustomerRepo(db),
			repo.ObjEmployeeRepo(db),
		),
	}
}
