import { Request, Response } from "express"
import { exec } from "child_process"

interface VideoMetaData {
	title: string;
	thumbnail: string;  
	duration: number;      
	category?: string;    
	links: {link: string, quality: string}[];
}

const extractMetaData = (url: string): Promise<VideoMetaData> => {
  return new Promise((resolve, reject) => {
    exec(`yt-dlp -j --no-playlist ${url}`, (error, stdout) => {
      if(error) return reject(`Failed to fetch meta data ${error.message}`)
      
      try{
        const metaData = JSON.parse(stdout)
        const videoData: VideoMetaData = {
          title: metaData.title || "No Title",
          thumbnail: metaData.thumbnail || "",
          duration: metaData.duration || 0,
          category: metaData.categories?.[0] || "Uncharacterized",
          links: metaData.formats.filter((format: any) => format.url).map((format: any) => ({
            link: format.url,
            quality: format.format,
          })),
        }
        resolve(videoData)
      } catch(error){
        reject(`Error parsing Data ${error}`)
      }
    })
  })
}

export const fetchMetaData = async (req: Request, res: Response): Promise<any> => {
  const videourl = req.query.url as string 
  if(!videourl) return res.status(400).json({error: "Video Url is Required"})
  
  try{
    const metaData = await extractMetaData(videourl)
    res.status(200).json(metaData)
  } catch(error){
    res.status(500).json({error: "Failed to get metadata"})
    console.error(error)
  }
}

