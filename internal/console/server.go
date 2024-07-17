package console

import (
	"kodinggo/db"
	handlerHttp "kodinggo/internal/delivery/http"
	"kodinggo/internal/repository"
	"kodinggo/internal/usecase"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.NewMysql()
	defer db.Close()

	storyRepo := repository.NewStoryRepository(db)
	storyUsecase := usecase.NewStoryUsecase(storyRepo)

	// Create a new Echo instance
	e := echo.New()

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewStoryHandler(routeGroup, storyUsecase)

	e.Logger.Fatal(e.Start(":3200"))
}
