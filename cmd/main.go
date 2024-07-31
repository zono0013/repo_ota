package main

import (
	"log"
	"repo/internal/application/services"
	"repo/internal/infrastructure/database"
	"repo/internal/infrastructure/repositories"
	"repo/internal/infrastructure/server"
	"repo/internal/interfaces/handlers"
)

func main() {
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repositories.NewMySQLUserRepository(db)
	reportRepo := repositories.NewMySQLReportRepository(db)

	userService := services.NewUserService(userRepo)
	reportService := services.NewReportService(reportRepo)

	userHandler := handlers.NewUserHandler(userService)
	reportHandler := handlers.NewReportHandler(reportService)

	server := server.NewServer(userHandler, reportHandler)

	log.Fatal(server.Run(":8080"))
}
