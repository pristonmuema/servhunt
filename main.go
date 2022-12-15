package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"servhunt/config"
	httpdao "servhunt/infra/dao"
	"servhunt/infra/token"
	"servhunt/infra/utils"
	"servhunt/routing"
	"servhunt/servitorservices"
	svcdao "servhunt/servitorservices/dao"
	"servhunt/user"
	"servhunt/user/dao"
	"syscall"
	"time"
)

var (
	defaultPort = 9094
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	rootLogger := utils.GetRootLogger()
	gin.SetMode(gin.ReleaseMode)
	router := InitRouter()
	tokenMaker, tkn := token.NewJWTMaker(token.RandomString(32))
	if tkn != nil {
		rootLogger.Fatal("An error occurred when creating a token maker")
	}
	conf := config.InitViperConfig()
	initDB := httpdao.Connection(conf)
	initRepo := httpdao.InitRepository(initDB)
	userDao := dao.NewUserRepoImpl(initRepo)

	userService := user.NewUserServiceImpl(userDao, tokenMaker)
	userHandler := user.NewUsersHandlerImpl(userService)
	userRouter := routing.NewUserRouter(router, userHandler, tokenMaker)
	userRouter.InitUserRoutes()

	servDao := svcdao.NewServiceRepoImpl(initRepo)
	servitorSvc := servitorservices.NewServitorSvc(servDao)
	servitorHandler := servitorservices.NewServitorServicesHandlerImpl(servitorSvc)
	servitorRouter := routing.NewServitorServicesRouter(router, servitorHandler, tokenMaker)
	servitorRouter.InitServitorServicesRoutes()

	errA := initDB.AutoMigrate(&dao.User{}, &dao.Language{}, &svcdao.Service{}, &svcdao.Location{}, &svcdao.Category{})
	if errA != nil {
		rootLogger.Fatal("An error occurred when running db migrations")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: router,
	}
	// Initializing the server in a goroutine so that it won't block the graceful handling
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			rootLogger.Fatal("An error occurred while starting up the server", zap.NamedError("error", err))
		}
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	rootLogger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		rootLogger.Fatal("Server forced to shutdown: ", zap.NamedError("error", err))
	}

	rootLogger.Info("Server exiting")
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(utils.CORSMiddleware())
	return router
}
