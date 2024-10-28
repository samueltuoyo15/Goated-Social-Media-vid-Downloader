import express from 'express'
import axios from 'axios'
import path from 'path'
import cors from 'cors'
import dotenv from 'dotenv'
import { fileURLToPath } from 'url'
const app = express()
dotenv.config()
const port = process.env.PORT
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
app.use(cors())
app.options('*', cors()) 
app.use(express.static(path.join(__dirname, '..', 'client', 'dist')))

app.get('/download', async (req, res) => {
  const videoURL = req.query.url
  if (!videoURL) return res.status(400).send('YouTube URL is required')
  
  try {
    const options = {
      method: 'GET',
      port: 'null',
      url: process.env.END_POINT,
      params: { url: videoURL },
      headers: {
        'x-rapidapi-key': process.env.API_KEY,
        'x-rapidapi-host': process.env.HOST
      }
    }
    const response = await axios.request(options)
    console.log('API Response:', response.data)
    const videoData = {
      description: response.data.description,
      picture: response.data.picture,
      downloadLinkMP3: response.data.links[0].link,
      downloadLink: response.data.links[8].link
    }
   res.json(videoData)
  } catch (error) {
    console.error('Error fetching video:', error)
    res.status(500).send('Error downloading the video')
  }
})
app.listen(port, () => {
  console.log(`Server is running on port ${port}`)
})
