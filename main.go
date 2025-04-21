package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/joho/godotenv"
) 

type VideoRequest struct {
	URL string `json:"url"`
}

type VideoResponse struct {
	URL string `json:"url"`
	Source string `json:"source"`
	ID string `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Duration int `json:"duration"`
	Medias []struct {
		URL string `json:"url"`
		Quality string `json:"quality"`
		Width int `json:"width"`
		Height int `json:"height"`
		Ext string `json:"ext"`
	} `json:"medias"`
	Error bool `json:"error"`
}


func main() {
 if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found.")
		}
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable not set")
	}

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error Parsing Form", http.StatusBadRequest)
			return
		}

		videoURL := r.FormValue("videoURL")
		if videoURL == "" {
			http.Error(w, "videoURL is required", http.StatusBadRequest)
			return
		}

		videoData, err := fetchVideoMetaData(videoURL, secretKey)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching video meta data: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		renderVideoData(w, videoData)
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func fetchVideoMetaData(videoURL, apiKey string) (*VideoResponse, error) {
	payload := VideoRequest{URL: videoURL}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://social-download-all-in-one.p.rapidapi.com/v1/social/autolink",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "social-download-all-in-one.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result VideoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func renderVideoData(w io.Writer, data *VideoResponse) {
	if len(data.Medias) == 0 {
		fmt.Fprint(w, `<div class="text-red-500">No download links available</div>`)
		return
	}
  
  sanitizedTitle := strings.ReplaceAll(data.Title, "/", "-")
	fmt.Fprintf(w, `
	<div class="mt-6 mb-20 p-4 rounded-lg shadow-2xl" x-data="{ selectedUrl: '%s' }">
		<h3 class="text-lg font-bold mb-4">Video Details</h3>
		<img src="%s" alt="Video Thumbnail" class="w-full rounded-md mb-4" />
		<p class="text-white mb-2"><strong>Title:</strong> %s</p>
		<p class="text-white mb-2"><strong>Source:</strong> %s</p>
		<p class="text-white mb-2"><strong>Author:</strong> %s</p>
		<p class="text-white mb-2"><strong>Duration:</strong> %d seconds</p>
		<div class="mt-4">
			<label for="qualitySelect" class="block mb-2">Select Quality</label>
			<select id="qualitySelect" x-model="selectedUrl" class="w-full p-2 bg-neutral-800 text-white rounded-md border">`,
		data.Medias[0].URL,
		data.Thumbnail,
		data.Title,
		data.Source,
		data.Author,
		data.Duration,
	)

	for _, media := range data.Medias {
	  qualityLabel := strings.TrimSpace(media.Quality)
	  if qualityLabel == ""{
	    qualityLabel = fmt.Sprintf("%dx%d %s", media.Width, media.Height, media.Ext) 
	  }
		fmt.Fprintf(w, `<option value="%s">%s</option>`, media.URL, qualityLabel)
	}

	fmt.Fprint(w, `
			</select>
		</div>
   <button @click="fetch(selectedUrl)
	.then(res => res.blob())
	.then(blob => {
		const url = URL.createObjectURL(blob)
		const a = document.createElement('a')
		a.href = url
		a.download = '%s.mp4'
  	a.click()
		URL.revokeObjectURL(url)
	})"
	class="cursor-pointer block mb-32 w-full mt-4 bg-red-900 text-center text-white p-3 rounded-md hover:bg-blue-600">
	Download Video
</button>
	</div>`, sanitizedTitle)
}


