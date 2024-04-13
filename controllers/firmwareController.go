package controllers

import (
	"context"
	"dms/initializers"
	"dms/repositories"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func GetFirmware(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")
	softwareVersion := c.Query("software_version")

	if softwareVersion == "" {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	path := fmt.Sprintf("userId_%s/deviceId_%s/version_%s/firmware.bin", strconv.Itoa(int(userId)), deviceId, softwareVersion)
	blob := initializers.AzContainer.NewBlockBlobURL(path)
	ctx := context.Background()
	response, err := blob.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		c.JSON(404, gin.H{"error": "Firmware not found"})
		return
	}

	// Read the downloaded data into a []byte
	bodyStream := response.Body(azblob.RetryReaderOptions{})
	data, err := io.ReadAll(bodyStream)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while reading firmware"})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=firmware.bin")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Data(http.StatusOK, "application/octet-stream", data)

}

func UploadFirmware(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")

	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	softwareVersion := c.Query("software_version")
	if softwareVersion == "" {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return

	}

	file, err := c.FormFile("firmware")
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	path := fmt.Sprintf("userId_%s/deviceId_%s/version_%s/firmware.bin", strconv.Itoa(int(userId)), deviceId, softwareVersion)
	blob := initializers.AzContainer.NewBlockBlobURL(path)
	ctx := context.Background()
	fileStream, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while reading firmware"})
		return
	}
	_, err = azblob.UploadStreamToBlockBlob(ctx, fileStream, blob, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while uploading firmware"})
		return
	}

	c.JSON(201, gin.H{"message": "Firmware uploaded successfully"})
}

func DeleteFirmware(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")

	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	softwareVersion := c.Query("software_version")
	if softwareVersion == "" {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return

	}

	path := fmt.Sprintf("userId_%s/deviceId_%s/version_%s/firmware.bin", strconv.Itoa(int(userId)), deviceId, softwareVersion)
	blob := initializers.AzContainer.NewBlockBlobURL(path)
	ctx := context.Background()

	// Check if the blob exists
	_, err = blob.GetProperties(ctx, azblob.BlobAccessConditions{}, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		c.JSON(404, gin.H{"error": "Firmware not found"})
		return
	}

	_, err = blob.Delete(ctx, azblob.DeleteSnapshotsOptionInclude, azblob.BlobAccessConditions{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while deleting firmware"})
		return
	}
	initializers.RedisClient.Del(c.Request.URL.RequestURI())
	c.JSON(200, gin.H{"message": "Firmware deleted successfully"})
}
