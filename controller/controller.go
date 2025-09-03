package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prodownlaoder/Direct/models"
	services "github.com/prodownlaoder/Direct/service"
)

func VideoHandler(c *gin.Context) {
	var req models.Request

	// Parse request JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate required fields
	if req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL and Quality are required"})
		return
	}

	// Detect platform (YouTube, TikTok, etc.)
	platformInfo := services.DetectPlatform(req.URL)
	log.Printf("[INFO] Detected Platform: %s | ", platformInfo.Platform)

	// If unsupported → return error
	if !platformInfo.IsValid {
		log.Printf("[ERROR] Unsupported platform ")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Unsupported or invalid platform",
			"platform": platformInfo.Platform,
		})
		return
	}

	GetVideoInfo(
		c,
		platformInfo.Platform,
		req.URL,
	)
}

func GetVideoInfo(c *gin.Context, platform string, url string) {

	videoInfo, err := services.GetVideoInfoService(url, platform)
	if err != nil {
		log.Printf("Error fetching video info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch video info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type":       "direct-download",
		"video_info": videoInfo,
	})
}
