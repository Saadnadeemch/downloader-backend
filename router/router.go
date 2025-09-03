package router

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	controllers "github.com/prodownlaoder/Direct/controller"
)

func sanitizeFileName(name string) string {
	if name == "" {
		return "video"
	}
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(`<>:"/\|?*`, r) {
			return -1
		}
		return r
	}, name)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/getvideo", controllers.VideoHandler)

	r.GET("/proxy-download", func(c *gin.Context) {
		url := c.Query("url")
		title := sanitizeFileName(c.Query("title"))
		platform := c.Query("platform")

		if url == "" || title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing required parameters"})
			return
		}

		filename := fmt.Sprintf("%s-%s.mp4", title, platform)

		// Fetch the actual video
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("❌ Proxy fetch failed: %v", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch video"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "remote server error"})
			return
		}

		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Length", fmt.Sprintf("%d", resp.ContentLength))

		if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			log.Printf("❌ Proxy stream error: %v", err)
		}
	})

	return r
}
