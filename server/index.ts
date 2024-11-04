import express, { Application, Request, Response } from 'express';
import axios from 'axios'
import path from 'path'
import cors from 'cors'
import dotenv from 'dotenv'
import { fileURLToPath } from 'url'
const app = express()

dotenv.config();

const app: Application = express();
const port: string | number = process.env.PORT || 3000;


app.use(cors());

app.use(express.json());

interface VideoData {
  description: string;
  picture: string;
  downloadLinkMP3: string;
  downloadLink: string;
}
const port = process.env.PORT
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

app.use(express.static(path.join(__dirname, '..', 'client', 'dist')))

app.get('/yt_downloader', async (req: Request, res: Response) => {
  const videoURL = req.query.url as string;
  if (!videoURL) {
    res.status(400).send('YouTube URL is required');
    return;
  }

  try {
    const response = await axios.get<any>(process.env.END_POINT!, {
      params: { url: videoURL },
      headers: {
        'x-rapidapi-key': process.env.API_KEY!,
        'x-rapidapi-host': process.env.HOST!,
      },
    });
    const downloadLink = response.data.links[8].link;
    const videoResponse = await axios.get(downloadLink, { responseType: 'arraybuffer' });

    const blob = Buffer.from(videoResponse.data as ArrayBuffer);
    const blobUrl = `data:video/mp4;base64,${blob.toString('base64')}`; 

    const videoData: VideoData = {
      description: response.data.description,
      picture: response.data.picture,
      downloadLinkMP3: response.data.links[0].link,
      downloadLink: blobUrl, 
    };
    res.set({
      'Content-Disposition': `attachment; filename="${videoData.description}.mp4"`, 
      'Content-Type': 'video/mp4', 
      'Access-Control-Allow-Origin': '*', 
    });
    
    res.json(videoData);
  } catch (error) {
    console.error('Error fetching video:', error);
    res.status(500).send('Error downloading the video');
  }
});
app.listen(port, () => {
  console.log(`Server is running on port ${port}`)
})
