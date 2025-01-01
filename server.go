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
	transaction service.TransactionService
	customer    service.CustomerService
	product     service.ProductService
	employee    service.EmployeeService

	host   string
	engine *gin.Engine
}

func (s *Server) initRoute() {
	group := s.engine.Group("/")
	controller.ObjTransactionController(group, s.transaction).Route()
	controller.ObjCustomerController(group, s.customer).Route()
	controller.ObjProductController(group, s.product).Route()
	controller.ObjEmployeeController(group, s.employee).Route()

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
		host:        ":" + cfg.Api.Port,
		engine:      gin.Default(),
		transaction: service.ObjTransactionService(repo.ObjTransactionRepo(db)),
		customer:    service.ObjCustomerService(repo.ObjCustomerRepo(db)),
		product:     service.ObjProductService(repo.ObjProductRepo(db)),
		employee:    service.EmployeeService(repo.ObjEmployeeRepo(db)),
	}
}
