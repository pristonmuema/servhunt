package routing

import (
	"github.com/gin-gonic/gin"
	"servhunt/infra/token"
	"servhunt/infra/utils"
	"servhunt/servitorservices"
	"servhunt/user"
)

type UsersRouter struct {
	engine *gin.Engine
	user.UsersHandler
	token.Maker
}

func NewUserRouter(engine *gin.Engine, handler user.UsersHandler, tm token.Maker) *UsersRouter {
	return &UsersRouter{
		engine:       engine,
		UsersHandler: handler,
		Maker:        tm,
	}
}

func (router UsersRouter) InitUserRoutes() {

	unauthenticated := router.engine.Group("/")
	{
		unauthenticated.POST("users", router.CreateUserAccount)
		unauthenticated.POST("login", router.Login)
		unauthenticated.POST("logout", router.Logout)
		unauthenticated.POST("change-password", router.ChangePassword)
	}

	v1 := router.engine.Group("/users").Use(utils.AuthMiddleware(router.Maker))
	{
		v1.PUT("/:user_id/update", router.UpdateUserAccount)
		v1.GET("", router.GetAllUsers)
		v1.GET("/:user_id", router.GetUserById)
		v1.GET("/phone/:phone_no", router.GetUserByPhoneNo)
		v1.GET("/email/:email", router.GetUserByEmail)
	}
}

type ServitorServicesRouter struct {
	engine *gin.Engine
	servitorservices.ServitorServicesHandler
	token.Maker
}

func NewServitorServicesRouter(engine *gin.Engine, handler servitorservices.ServitorServicesHandler,
	tm token.Maker) *ServitorServicesRouter {
	return &ServitorServicesRouter{
		engine:                  engine,
		ServitorServicesHandler: handler,
		Maker:                   tm,
	}
}

func (router ServitorServicesRouter) InitServitorServicesRoutes() {
	v1 := router.engine.Group("/services").Use(utils.AuthMiddleware(router.Maker))
	{
		v1.POST("", router.CreateService)
		v1.PUT("/:service_id/update", router.UpdateService)
		v1.GET("", router.GetAllServices)
		v1.GET("/:service_id", router.GetServiceByID)
		v1.GET("/servitors/:user_id", router.ServitorsService)
		v1.POST("/locations", router.CreateLocationInfo)
		v1.PUT("/locations/:id", router.UpdateLocationInfo)
		v1.GET("/locations", router.GetAllLocations)
		v1.GET("/locations/:service_id", router.GetServiceLocations)
		v1.POST("/categories", router.CreateCategory)
		v1.PUT("/categories/:id/update", router.UpdateCategory)
		v1.GET("/categories", router.GetAllCategories)
		v1.GET("/categories/:service_id", router.GetServiceCategories)
	}
}
