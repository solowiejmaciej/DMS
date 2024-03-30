package main

import (
	"dms/controllers"
	"dms/initializers"
	"dms/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/takt-corp/gin-logrus"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConfigureLogs()
	initializers.ConnectToDB()
	initializers.Migrate()
	initializers.ConfigureMqtt()
}

func main() {
	r := gin.New()
	r.Use(ginlogrus.LoggerMiddleware(ginlogrus.LoggerMiddlewareParams{}))
	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	r.POST("/api/activate", controllers.ActivateDevice)
	r.GET("/api/configuration", middleware.RequireDeviceToken, controllers.GetConfiguration)
	r.POST("/api/configuration", middleware.RequireDeviceToken, controllers.AddConfiguration)
	r.POST("/api/device/:deviceId/reboot", controllers.RebootDevice)
	r.GET("/api/device/:deviceId", controllers.GetDevice)
	r.GET("/api/device/", controllers.GetDevices)
	r.GET("/api/device/:deviceId/firmware", middleware.RequireDeviceToken, controllers.GetFirmware)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
