package main

import (
	"dms/controllers"
	"dms/db"
	"dms/initializers"
	"dms/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/takt-corp/gin-logrus"
	"time"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConfigureLogs()
	db.ConnectToDB()
	initializers.Migrate()
	initializers.ConfigureMqtt()
	initializers.ConnectToRedis()
	initializers.ConnectToAzureStorage()
}

func main() {
	r := gin.New()
	r.Use(ginlogrus.LoggerMiddleware(ginlogrus.LoggerMiddlewareParams{}))
	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	r.POST("/api/user", controllers.AddUser)
	r.GET("/api/user/:userId", controllers.GetUserById)

	r.POST("/api/activate/:deviceId", controllers.ActivateDevice)
	r.DELETE("/api/device/:deviceId", middleware.RequireDeviceOwnership, controllers.DeleteDevice)

	r.GET("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, middleware.CachePage(time.Hour), controllers.GetConfiguration)
	r.POST("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, controllers.AddConfiguration)
	r.PUT("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, middleware.DeleteCache(), controllers.UpdateConfiguration)

	r.GET("/api/device/", controllers.GetDevices)
	r.GET("/api/device/:deviceId", middleware.RequireDeviceOwnership, controllers.GetDevice)
	r.POST("/api/device/:deviceId/reboot", middleware.RequireDeviceOwnership, controllers.RebootDevice)
	r.POST("/api/device/:deviceId/update", middleware.RequireDeviceOwnership, controllers.UpdateDeviceFirmware)

	r.GET("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.GetFirmware)
	r.POST("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.UploadFirmware)
	r.DELETE("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.DeleteFirmware)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
