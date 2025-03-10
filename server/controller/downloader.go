package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"github.com/gin-gonic/gin"
)

type VideoMetaData struct {
	Title string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Duration int `json:"duration"`
	Category string `json:"category"`
	Links []VideoQuality `json:"links"`
}


type VideoQuality struct {
	Link    string `json:"link"`
	Quality string `json:"quality"`
}

func ExtractMetaData(videoURL string) (*VideoMetaData, error) {
	cmd := exec.Command("yt-dlp", "-j", "--no-playlist", videoURL)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %v", err)
	}

	var metaData map[string]interface{}
	if err := json.Unmarshal(output, &metaData); err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}

	var links []VideoQuality
	formats, ok := metaData["formats"].([]interface{})
	if ok {
		for _, format := range formats {
			f := format.(map[string]interface{})
			if f["url"] != nil {
				links = append(links, VideoQuality{
					Link:    f["url"].(string),
					Quality: f["format"].(string),
				})
			}
		}
	}

	videoData := &VideoMetaData{
		Title:     metaData["title"].(string),
		Thumbnail: metaData["thumbnail"].(string),
		Duration:  int(metaData["duration"].(float64)),
		Category:  "Uncategorized",
		Links:     links,
	}

	if categories, ok := metaData["categories"].([]interface{}); ok && len(categories) > 0 {
		videoData.Category = categories[0].(string)
	}

	return videoData, nil
}

func ProxyRequest(c *gin.Context) {
	targetURL := c.Query("url")
	if targetURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing URL parameter"})
		return
	}

	parsedURL, err := url.Parse(targetURL)
	if err != nil || !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	resp, err := http.Get(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resource"})
		return
	}
	defer resp.Body.Close()

	c.Header("Content-Type", resp.Header.Get("Content-Type"))
	c.Header("Cache-Control", "public, max-age=86400")
	io.Copy(c.Writer, resp.Body)
}