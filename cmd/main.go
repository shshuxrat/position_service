package main

import (
	"fmt"
	"net"
	"position_service/config"
	"position_service/genproto/position_service"
	"position_service/pkg/logger"
	"position_service/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Environment, "example_service")

	defer logger.Cleanup(log)

	pgStore := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := sqlx.Connect("postgres", pgStore)

	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)

	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	professionService := service.NewProfessionService(log, db)
	companyService := service.NewCompanyService(log, db)
	attributeService := service.NewAttributeService(log, db)

	s := grpc.NewServer()
	reflection.Register(s)

	position_service.RegisterProfessionServiseServer(s, professionService)
	position_service.RegisterCompanyServiceServer(s, companyService)
	position_service.RegisterAttributeServiceServer(s, attributeService)

	log.Info("main server running", logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Error("error while listening", logger.Error(err))
		return
	}

}
