package controller

import (
	"encoding/json"
	"fmt"
	"os/exec"
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


func ExtractMetaData(url string) (*VideoMetaData, error) {
	cmd := exec.Command("yt-dlp", "-j", "--no-playlist", url)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %v", err)
	}

	var metaData map[string]interface{}
	if err := json.Unmarshal(output, &metaData); err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}


	var links []VideoQuality
	formats := metaData["formats"].([]interface{})
	for _, format := range formats {
		f := format.(map[string]interface{})
		if f["url"] != nil {
			links = append(links, VideoQuality{
				Link:    f["url"].(string),
				Quality: f["format"].(string),
			})
		}
	}


	videoData := &VideoMetaData{
		Title: metaData["title"].(string),
		Thumbnail: metaData["thumbnail"].(string),
		Duration: int(metaData["duration"].(float64)),
		Category: "Uncharacterized",
		Links: links,
	}

	if categories, ok := metaData["categories"].([]interface{}); ok && len(categories) > 0 {
		videoData.Category = categories[0].(string)
	}

	return videoData, nil
}