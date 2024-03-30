package controllers

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func GetFirmware(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=firmware.bin")
	filePath := filepath.Join("firmware", c.Param("deviceId"), "firmware.bin")
	c.FileAttachment(filePath, "firmware.bin")
}
