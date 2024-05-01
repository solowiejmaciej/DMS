package main

import (
	"dms/controllers"
	"dms/db"
	"dms/initializers"
	"dms/middleware"
	"github.com/gin-contrib/cors"
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
	initializers.ConnectToRabbitMq()
}

func main() {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                            // Allow all origins
	config.AllowCredentials = true                                                                           // Allow credentials
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Api-Key"} // Allow specific headers

	r.Use(cors.New(config))
	r.Use(ginlogrus.LoggerMiddleware(ginlogrus.LoggerMiddlewareParams{}))

	r.GET("/api/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	r.POST("/api/token", controllers.GenerateToken)
	r.POST("/api/token/refresh", controllers.RefreshToken)
	r.POST("/api/token/apikey/regenerate", middleware.RequireToken, controllers.RegenerateApiKey)
	r.POST("/api/me", middleware.RequireToken, controllers.GetMe)

	r.POST("/api/user", controllers.AddUser)
	r.GET("/api/user/:userId", controllers.GetUserById)
	r.PUT("/api/user/:userId", middleware.RequireToken, controllers.UpdateUser)
	r.POST("/api/user/send-confirm-phone", middleware.RequireToken, controllers.PublishSendConfirmPhoneNumber)
	r.POST("/api/user/confirm-phone", middleware.RequireToken, controllers.ConfirmPhoneNumber)

	r.POST("/api/activate/:deviceId", controllers.ActivateDevice)
	r.DELETE("/api/device/:deviceId", middleware.RequireDeviceOwnership, controllers.DeleteDevice)

	r.GET("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, middleware.CachePage(time.Hour), controllers.GetConfiguration)
	r.POST("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, controllers.AddConfiguration)
	r.PUT("/api/configuration/:deviceId", middleware.RequireDeviceOwnership, middleware.DeleteCache(), controllers.UpdateOrCreateConfiguration)

	r.GET("/api/device", controllers.GetDevices)
	r.GET("/api/device/:deviceId/status", middleware.RequireDeviceOwnership, controllers.GetDeviceStatus)
	r.GET("/api/device/:deviceId", middleware.RequireDeviceOwnership, controllers.GetDevice)
	r.POST("/api/device/:deviceId/reboot", middleware.RequireDeviceOwnership, controllers.RebootDevice)
	r.POST("/api/device/:deviceId/update", middleware.RequireDeviceOwnership, controllers.UpdateDeviceFirmware)
	r.POST("/api/device/:deviceId/test-mqtt-connection", middleware.RequireDeviceOwnership, controllers.TestMqttConnection)

	r.GET("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.GetFirmware)
	r.POST("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.UploadFirmware)
	r.DELETE("/api/device/:deviceId/firmware", middleware.RequireDeviceOwnership, controllers.DeleteFirmware)

	r.POST("/api/bug", middleware.RequireToken, controllers.ReportBug)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}
