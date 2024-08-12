package console

import (
	"kodinggo/db"
	"kodinggo/internal/config"
	handlerHttp "kodinggo/internal/delivery/http"
	"kodinggo/internal/repository"
	"kodinggo/internal/usecase"
	"kodinggo/internal/worker"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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

	mysql := db.NewMysql()
	defer mysql.Close()

	redis := db.NewRedis()

	// Init asynq client
	workerClient := worker.InitAsynqClient(config.GetRedisHost())

	storyRepo := repository.NewStoryRepository(mysql, redis)
	userRepo := repository.NewUserRepository(mysql)

	commentService := newCommentClientGRPC()
	categoryService := newCategoryClientGRPC()

	storyUsecase := usecase.NewStoryUsecase(storyRepo, commentService, categoryService, workerClient)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Create a new Echo instance
	e := echo.New()
	// e.Use(echoMiddleware.Recover())
	// e.Use(echoMiddleware.Logger())

	routeGroup := e.Group("/api/v1")

	handlerHttp.NewStoryHandler(routeGroup, storyUsecase)

	handlerHttp.NewUserHandler(routeGroup, userUsecase)

	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := e.Start(":3200")
		if err != nil {
			errCh <- err
		}
	}()

	go func() {
		defer wg.Done()
		err := worker.InitAsynqServer(config.GetRedisHost())
		if err != nil {
			errCh <- err
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			logrus.Error(err.Error())
		}
	}
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
