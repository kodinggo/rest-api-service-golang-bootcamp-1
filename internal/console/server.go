package console

import (
	"kodinggo/db"
	"kodinggo/internal/config"
	handlerHttp "kodinggo/internal/delivery/http"
	"kodinggo/internal/repository"
	"kodinggo/internal/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbCategory "github.com/kodinggo/category-service-gb1/pb/category"
	pbComment "github.com/kodinggo/comment-service-gb1/pb/comment"
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

	db := db.NewMysql()
	defer db.Close()

	storyRepo := repository.NewStoryRepository(db)
	userRepo := repository.NewUserRepository(db)

	commentService := newCommentClientGRPC()
	categoryService := newCategoryClientGRPC()

	storyUsecase := usecase.NewStoryUsecase(storyRepo, commentService, categoryService)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Create a new Echo instance
	e := echo.New()
	// e.Use(echoMiddleware.Recover())
	// e.Use(echoMiddleware.Logger())

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewStoryHandler(routeGroup, storyUsecase)

	handlerHttp.NewUserHandler(routeGroup, userUsecase)

	e.Logger.Fatal(e.Start(":3200"))
}

func newCommentClientGRPC() pbComment.CommentServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:3100", opts...)
	if err != nil {
		log.Fatal(err)
	}

	return pbComment.NewCommentServiceClient(conn)
}

func newCategoryClientGRPC() pbCategory.CategoryServiceClient {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:3300", opts...)
	if err != nil {
		log.Fatal(err)
	}

	return pbCategory.NewCategoryServiceClient(conn)
}
