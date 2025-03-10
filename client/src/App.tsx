import { useState } from "react"
import CustomSlider from "./Components/CustomSlider"

interface VideoMetaData {
  title: string;
  thumbnail: string;
  duration: number;
  category?: string;
  links: { link: string; quality: string }[];
}

const App = () => {
  const [videoURL, setVideoURL] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [videoData, setVideoData] = useState<VideoMetaData | null>(null)
  const [selectedUrl, setSelectedUrl] = useState("")

  const downloadVideo = async (link: string) => {
    if (!link) return;
    try {
      const response = await fetch(link)
      if (!response.ok) throw new Error("Failed to download video.")
      
      const blob = await response.blob()
      const newLink = URL.createObjectURL(blob)
      const downloadLink = document.createElement("a")
      downloadLink.download = "video.mp4"
      downloadLink.href = newLink
      downloadLink.click()
      URL.revokeObjectURL(newLink)
    } catch (error) {
      console.error("Error downloading video:", error)
      alert("Download failed. Please try again.")
    }
  };

  const submitForm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    setIsLoading(true)
    setError(null)
    setVideoData(null)

    if (!videoURL.trim() || !/^https?:\/\//.test(videoURL)) {
      setIsLoading(false);
      setError("Please enter a valid video URL.")
      return
    }
    
    if (videoURL.includes("x.com/") || videoURL.includes("twitter.com/")) {
    setIsLoading(false)
    setError("X (Twitter) downloads are currently under development.")
    return
  }
    try {
      const response = await fetch(`/metadata?url=${encodeURIComponent(videoURL)}`)
      const data = await response.json()

      if (!response.ok) {
        setError(data.error || "Failed to fetch metadata.")
        return
      }

      setVideoData(data)
      if (data.links?.length) {
        setSelectedUrl(data.links[0].link || "")
      }
    } catch (error) {
      setError("An error occurred while fetching metadata.")
      console.error(error)
    } finally {
      setIsLoading(false)
    }
  };

  return (
    <>
      <section
        onContextMenu={(e) => e.preventDefault()}
        className="select-none flex justify-center items-center flex-col bg-neutral-900 min-h-screen text-white text-center mb-8"
      >
        <CustomSlider />
        <div>
          <h2 className="text-2xl">Securely Download Your Favourite Social Media Videos</h2>
          <h2 className="text-1xl">for free</h2>
        </div>

        <form onSubmit={submitForm} className="flex justify-center items-center mt-10" id="download-form">
          <input
            value={videoURL}
            required
            onChange={(e) => setVideoURL(e.target.value)}
            type="url"
            placeholder="Paste your URL here..."
            className="w-52 text-black rounded p-3"
          />
          <input type="submit" value="Download" className="bg-neutral-800 text-white rounded ml-3 p-3" />
        </form>

        {isLoading && (
          <div className="absolute inset-0 flex flex-col justify-center items-center bg-neutral-900 bg-opacity-70">
            <img src="/spinning-dots.svg" alt="Loading..." className="w-72" />
            <p className="block text-lg">Fetching Video Info...</p>
          </div>
        )}

        {error && <p className="text-red-500 mt-5">{error}</p>}

        {videoData && (
          <div className="mt-6 mb-20 p-4 rounded-lg shadow-2xl">
            <h3 className="text-lg font-bold mb-4">Video Details</h3>
            {videoData.thumbnail && (
              <img src={videoData.thumbnail} alt="Video Thumbnail" className="w-full rounded-md mb-4" />
            )}
            <p className="text-white mb-2">
              <strong>Title:</strong> {videoData.title}
            </p>
            <p className="text-white mb-2">
              <strong>Category:</strong> {videoData.category}
            </p>
            <p className="text-white mb-2">
              <strong>Duration:</strong> {videoData.duration} seconds
            </p>

            {videoData.links?.length > 0 && (
              <div className="mt-4">
                <label htmlFor="qualitySelect" className="block mb-2">
                  Select Quality
                </label>
                <select
                  id="qualitySelect"
                  onChange={(e) => setSelectedUrl(e.target.value)}
                  className="w-full p-2 bg-neutral-800 text-white rounded-md border"
                >
                  {videoData.links.map((media, index) => (
                    <option key={index} value={media.link}>
                      {media.quality}
                    </option>
                  ))}
                </select>
              </div>
            )}

            <button
              type="button"
              onClick={() => downloadVideo(selectedUrl)}
              disabled={!selectedUrl}
              className="cursor-pointer block mb-32 w-full mt-4 bg-red-900 text-center text-white p-3 rounded-md hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Download Video
            </button>
          </div>
        )}
      </section>
      <footer className="fixed p-4 text-center bg-neutral-900 text-white text-center bottom-0 w-full">
        Developed by Samuel Tuoyo with ❤️
      </footer>
    </>
  )
}

export default App