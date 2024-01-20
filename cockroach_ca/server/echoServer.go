package server

import (
	"fmt"

	cockroachHandlers "github.com/hank8451/cockroach_ca/cockroach/handlers"
	cockroachRepositories "github.com/hank8451/cockroach_ca/cockroach/repositories"
	cockroachUsecases "github.com/hank8451/cockroach_ca/cockroach/usecases"
	"github.com/hank8451/cockroach_ca/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.initializeCockroachHttpHandler()

	s.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgresRepository := cockroachRepositories.NewCockroachPostgresRepository(s.db)
	cockroachFCMMessaging := cockroachRepositories.NewCockroachFCMMessaging()

	cockroachUsecase := cockroachUsecases.NewCockroachUsecaseImpl(
		cockroachPostgresRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := cockroachHandlers.NewCockroachHttpHandler(cockroachUsecase)

	//Routers
	cockroachRouters := s.app.Group("/v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
