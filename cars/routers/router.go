package routers

import (
	"belajar-gin-cars/controllers"
	"belajar-gin-cars/models"

	"github.com/gin-gonic/gin"

	_ "belajar-gin-cars/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description This is a sample API for car management
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email koder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhos
// @BasePath /
func StartServer() *gin.Engine {
	models.StartDB()

	router := gin.Default()

	// Create
	router.POST("/cars", controllers.CreateCar)
	// Read All
	router.GET("/cars", controllers.GetAllCars)
	// Read
	router.GET("/cars/:carID", controllers.GetCar)
	// Update
	router.PUT("/cars/:carID", controllers.UpdateCar)
	// Delete
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
