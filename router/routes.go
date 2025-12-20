package router

import (
	"github.com/MarcelloBB/gin-boilerplate/controller"
	"github.com/MarcelloBB/gin-boilerplate/docs"
	"github.com/MarcelloBB/gin-boilerplate/repository"
	"github.com/MarcelloBB/gin-boilerplate/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, dbConnection *gorm.DB) {
	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	// @BasePath /api/
	api := r.Group(basePath)
	{
		api.GET("/user", UserController.GetUsers)
	}

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
