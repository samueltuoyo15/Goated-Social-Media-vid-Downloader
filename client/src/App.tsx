import { useState } from 'react'

const App = () => {
  const [videoURL, setVideoURL] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState(null)
  const [videoDescription, setVideoDescription] = useState('')
  const [videoDownloadLink, setVideoDownloadLink] = useState('')
  const [videoDownloadLinkMP3, setVideoDownloadLinkMP3] = useState('')

  const submitForm = async (e) => {
    e.preventDefault()
    setIsLoading(true)
    setError(null)

    if (!videoURL.trim()) {
      setIsLoading(false)
      return
    }

    try {
      const response = await fetch(`/download?url=${encodeURIComponent(videoURL)}`)
      if (!response.ok) throw new Error('Error fetching video')

      const data = await response.json()
      const description = data.description || 'No description available'
      const downloadLink = data.downloadLink || 'No download link available'
      const downloadLinkMP3 = data.downloadLinkMP3 || 'No Audio File'
      setVideoDescription(description)
      setVideoDownloadLink(downloadLink)
      setVideoDownloadLinkMP3(downloadLinkMP3)
    } catch (error) {
      alert('Failed to download video')
      console.error(error)
      setError('Failed to download video. Please try again.')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <>
    <section className="select-none flex justify-center items-center flex-col bg-neutral-900 h-screen w-screen text-white text-center mb-8">
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
      {videoDescription && (
        <div className="mt-5">
          <p className="text-white">{videoDescription}</p>
          {videoDownloadLink && (
            <a onClick={() => videoDownloadLink} className="mx-auto text-center text-blue-500 block mb-5 mt-4">
              Download Video + Audio
            </a>
          )}
          {videoDownloadLinkMP3 && (
            <a onClick={() => videoDownloadLinkMP3} className="mx-auto text-center text-blue-500 block">
              Download Audio Only
            </a>
          )}
        </div>
      )}
    </section>
   <footer className="fixed p-4 text-center bg-neutral-900 text-white text-center bottom-0 w-full">Developed By Samuel Tuoyo With ♥️</footer>
   </>
  )
}

export default App
