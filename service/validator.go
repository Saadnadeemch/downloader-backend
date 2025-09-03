package service

import (
	"log"
	"net/url"
	"strings"

	"github.com/prodownlaoder/Direct/models"
)

// host → platform mapping
var hostToPlatform = map[string]string{
	// YouTube
	"youtube.com":       "YouTube",
	"youtu.be":          "YouTube",
	"m.youtube.com":     "YouTube",
	"music.youtube.com": "YouTube",

	// Vimeo
	"vimeo.com":        "Vimeo",
	"player.vimeo.com": "Vimeo",

	// Dailymotion
	"dailymotion.com": "Dailymotion",
	"dai.ly":          "Dailymotion",

	// Twitter / X
	"twitter.com":        "Twitter",
	"x.com":              "Twitter",
	"mobile.twitter.com": "Twitter",

	// Twitch
	"twitch.tv":       "Twitch",
	"clips.twitch.tv": "Twitch",

	// Reddit
	"reddit.com": "Reddit",
	"redd.it":    "Reddit",
	"v.redd.it":  "Reddit",

	// Pinterest
	"pinterest.com": "Pinterest",
	"pin.it":        "Pinterest",

	// Other
	"ok.ru":          "OK.ru",
	"metacafe.com":   "Metacafe",
	"facebook.com":   "Facebook",
	"fb.watch":       "Facebook",
	"instagram.com":  "Instagram",
	"threads.net":    "Threads",
	"tiktok.com":     "TikTok",
	"likee.video":    "Likee",
	"bilibili.tv":    "Bilibili",
	"streamable.com": "Streamable",
	"vk.com":         "VK",
}

// Platforms that usually provide separate audio & video (DASH, adaptive streaming)
var DirectPlatforms = map[string]bool{
	"Instagram": true,
}

// DetectPlatform checks if the given URL belongs to a supported platform.
func DetectPlatform(inputURL string) models.PlatformInfo {
	parsed, err := url.Parse(inputURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		log.Printf("Invalid URL: %s", inputURL)
		return models.PlatformInfo{IsValid: false}
	}

	host := strings.ToLower(strings.TrimPrefix(parsed.Host, "www."))
	platform, exists := hostToPlatform[host]

	if exists {
		Direct := "No"
		if DirectPlatforms[platform] {
			Direct = "Yes"
		}
		log.Printf("Platform detected: %s (Separate AV: %s)", platform, Direct)
		return models.PlatformInfo{
			IsValid:  true,
			Platform: platform,
			IsDirect: Direct,
		}
	}

	log.Printf("Unsupported platform for URL: %s", inputURL)
	return models.PlatformInfo{IsValid: false}
}
