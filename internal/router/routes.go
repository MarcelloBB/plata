package router

import (
	"github.com/MarcelloBB/plata/docs"
	"github.com/MarcelloBB/plata/internal/controller"
	"github.com/MarcelloBB/plata/internal/repository"
	"github.com/MarcelloBB/plata/internal/usecase"
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
