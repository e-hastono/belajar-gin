package routers

import (
	"belajar-gin-cars/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars", controllers.GetAllCars)
	router.GET("/cars/:carID", controllers.GetCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
