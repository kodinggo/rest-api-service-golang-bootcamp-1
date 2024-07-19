package console

import (
	"kodinggo/db"
	"kodinggo/internal/config"
	handlerHttp "kodinggo/internal/delivery/http"
	"kodinggo/internal/repository"
	"kodinggo/internal/usecase"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start the HTTP server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	// Get env variables from .env file
	config.LoadWithViper()
	log := config.SetupLogger()

	db := db.NewMysql()
	defer db.Close()

	storyRepo := repository.NewStoryRepository(db)
	userRepo := repository.NewUserRepository(db)

	storyUsecase := usecase.NewStoryUsecase(storyRepo, log)
	userUsecase := usecase.NewUserUsecase(userRepo, log)

	// Create a new Echo instance
	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewStoryHandler(routeGroup, storyUsecase)

	handlerHttp.NewUserHandler(routeGroup, userUsecase)

	e.Logger.Fatal(e.Start(":3200"))
}
