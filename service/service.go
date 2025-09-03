package service

import (
	"fmt"

	"github.com/prodownlaoder/Direct/models"
	"github.com/prodownlaoder/Direct/runner"
)

// GetVideoInfoService tries to fetch video info using yt-dlp
func GetVideoInfoService(videoURL string, platform string) (*models.VideoInfo, error) {
	info, err := runner.GetDirectInfoFromYTDLP(videoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get video info from runner: %w", err)
	}

	// Set the source (you can also include platform if needed)
	info.Source = "yt-dlp"
	return info, nil
}
