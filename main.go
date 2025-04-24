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
	"context"
	"time"
	"crypto/tls"
	"github.com/redis/go-redis/v9"
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


var ctx = context.Background()
var rdb *redis.Client

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

  rdb = redis.NewClient(&redis.Options{
	  Addr: os.Getenv("REDIS_URL"),
  	Password: os.Getenv("REDIS_PASSWORD"),             
  	DB: 0,  
  	TLSConfig: &tls.Config{},
})
 
 
 pong, err := rdb.Ping(ctx).Result()
   if err != nil {
  	log.Fatalf("Redis connection failed: %v", err)
   }
  fmt.Println("Redis connected:", pong)

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)

	http.HandleFunc("/submit", rateLimit(func(w http.ResponseWriter, r *http.Request) {
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
    remaining := getRemainingLimit(r)
		w.Header().Set("Content-Type", "text/html")
		renderVideoData(w, videoData, remaining)
	}))

	http.HandleFunc("/download", rateLimit(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "URL parameter is required", http.StatusBadRequest)
			return
		}
		
		
		fileName := r.URL.Query().Get("filename")
		if fileName == ""{
		  fileName = "video.mp4"
		}
		
		resp, err := http.Get(url)
		if err != nil{
		  http.Error(w, "Failed to fetch video", http.StatusInternalServerError)
		  return 
		}
		
		defer resp.Body.Close()
		
   w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
   w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
   w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
   
   _, err = io.Copy(w, resp.Body)
   if err != nil {
     log.Printf("Error Streaming Video %v", err)
   }
	}))
 
 
 http.HandleFunc("/rate-info", func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		key := fmt.Sprintf("rate_limit:%s:%s", ip, time.Now().Format("2006-01-02"))

		count, _ := rdb.Get(ctx, key).Int()
		remaining := 4 - count
		if remaining < 0 {
			remaining = 0
		}
    
    w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"remaining": remaining})
	})
	
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func fetchVideoMetaData(videoURL, apiKey string) (*VideoResponse, error) {
  cacheKey := fmt.Sprintf("video_meta:%s", videoURL)
  cacheData, err := rdb.Get(ctx , cacheKey).Result()
		if err == nil {
		var v VideoResponse
		if json.Unmarshal([]byte(cachedData), &v) == nil {
			log.Println("Cache hit")
			return &v, nil
		}
	}
	
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
	body, _ := io.ReadAll(resp.Body)
	var result VideoResponse
	if json.Unmarshal(body, &result) != nil || result.Error {
		return &result, nil
	}

	cachedData, _ := json.Marshal(result)
	rdb.Set(ctx, cacheKey, cachedData, 8*time.Minute)
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
		<p class="text-white mb-2"><strong>Duration:</strong>%d</p>
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

	fmt.Fprintf(w, `
  </select>
</div>
<script>document.querySelector('[x-data]').__x.$data.remaining = %d</script>
  <a 
  x-bind:href="'/download?url=' + encodeURIComponent(selectedUrl) + '&filename=%s.mp4'" 
  class="block mb-32 w-full mt-4 bg-red-900 text-center text-white p-3 rounded-md hover:bg-blue-600"
 download
 >
  Download Video
 </a>
</div>`, sanitizedTitle, remaining)
}


func rateLimit(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        
        ip := r.RemoteAddr
        if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
            ip = strings.Split(forwarded, ",")[0]
        } else {
            ip = strings.Split(ip, ":")[0] 
        }

        key := fmt.Sprintf("rate_limit:%s:%s", ip, time.Now().Format("2006-01-02"))
        
        count, err := rdb.Incr(ctx, key).Result()
        if err != nil {
            log.Printf("Redis error: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        if count == 1 {
            _, err := rdb.Expire(ctx, key, 24*time.Hour).Result()
            if err != nil {
                log.Printf("Redis expiration error: %v", err)
            }
        }
       
       w.Header().Set("X-RateLimit-Limit", "4")
       w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", 4-count))
   
        if count > 4 {
            http.Error(w, "Rate limit exceeded (4 requests/day)", http.StatusTooManyRequests)
            return
        }

        next(w, r)
    }
}


func getIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}

func getRemainingLimit(r *http.Request) int {
	ip := getIP(r)
	key := fmt.Sprintf("rate_limit:%s:%s", ip, time.Now().Format("2006-01-02"))
	count, _ := rdb.Get(ctx, key).Int()
	remaining := 4 - count
	if remaining < 0 {
		remaining = 0
	}
	return remaining
}






