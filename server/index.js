const express = require('express');
const axios = require('axios');
const path = require('path');
const cors = require('cors');
const dotenv = require('dotenv');
const { fileURLToPath } = require('url');

dotenv.config();

const app = express();
const port = process.env.PORT || 3000;

app.use(cors());
app.use(express.json());

app.use(express.static(path.join(__dirname, '..', 'client', 'dist')));

app.get('/download', async (req, res) => {
  const videoURL = req.query.url;
  if (!videoURL) {
    res.status(400).send('YouTube URL is required');
    return;
  }

  try {
    const response = await axios.get('https://youtube-video-and-shorts-downloader1.p.rapidapi.com/api/getYTVideo', {
      params: { url: videoURL },
      headers: {
        'x-rapidapi-key': process.env.API_KEY,
        'x-rapidapi-host': process.env.HOST,
      },
    });
    const downloadLink = response.data.links[8].link;
    const videoResponse = await axios.get(downloadLink, { responseType: 'arraybuffer' });

    const blob = Buffer.from(videoResponse.data);
    const blobUrl = `data:video/mp4;base64,${blob.toString('base64')}`;

    const videoData = {
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
  console.log(`Server is running on port ${port}`);
});
