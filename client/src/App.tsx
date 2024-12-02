import { useState } from 'react'

const App = () => {
  const [videoURL, setVideoURL] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState(null)
  const [videoData, setVideoData] = useState(null);
  const [selectedUrl, setSelectedUrl] = useState('');

  const submitForm = async (e) => {
    e.preventDefault()
    setIsLoading(true)
    setError(null)
    setVideoData(null);
    
    if (!videoURL.trim() || !/^https?:\/\/(www\.)?(youtube\.com|youtu\.be)\/.+$/.test(videoURL)) {
     setIsLoading(false)
     return
    }

    
    try {
      const response = await fetch(`/download?url=${encodeURIComponent(videoURL)}`);

      if (!response.ok) {
        throw new Error('Failed to fetch video data.');
      }

      const data = await response.json();
      setVideoData(data);
      setIsLoading(false)
      if (data.links && data.links.length > 0) {
        setSelectedUrl(data.links[0].link);
      }
    } catch (err) {
      console.error(err);
      setError('Something went wrong. Please check the URL or try again.');
    } finally {
      setIsDownloading(false);
    }
  };
  

  return (
    <>
    <section className="select-none flex justify-center items-center flex-col bg-neutral-900 min-h-screen text-white text-center mb-8">
      <img src="/animate.png" className="w-72" alt="Animation" />
      <div>
        <h2 className="text-2xl">Download Your Favourite YouTube Video</h2>
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
          <p className="block text-lg">Fetching Video Info.....</p>
        </div>
      )}

      {error && <p className="text-red-500 mt-5">{error}</p>}
        {videoData && (
          <div className="mt-6 mb-20 p-4 rounded-lg shadow-2xl">
            <h3 className="text-lg font-bold mb-4">Video Details</h3>
            {videoData.picture && <img src={videoData.picture} alt="Video Thumbnail" className="w-full rounded-md mb-4" />}
            <p className="text-white mb-2">
              <strong>Title:</strong> {videoData.title}
            </p>
            <p className="text-white mb-2">
              <strong>Category:</strong> {videoData.stats.category}
            </p>
             <p className="text-white mb-2">
              <strong>Duration:</strong> {videoData.links[0].approxDurationMs} milleseconds
            </p>

            {videoData.links && videoData.links.length > 0 && (
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
                      Quality {media.quality}
                    </option>
                  ))}
                </select>
              </div>
            )}

            <a
              download
              href={selectedUrl}
              style={{
                pointerEvents: !selectedUrl ? 'none' : 'auto',
                opacity: !selectedUrl ? 0.5 : 1,
              }}
              className="block mb-32 w-full mt-4 bg-red-900 text-center text-white p-3 rounded-md hover:bg-blue-600"
            >
              Download Video
            </a>
          </div>
        )}
    </section>
   <footer className="fixed p-4 text-center bg-neutral-900 text-white text-center bottom-0 w-full">Developed By Samuel Tuoyo With ♥️</footer>
   </>
  )
}

export default App
