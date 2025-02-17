import { Router } from "express"
import { fetchMetaData, download } from "../controller/downloader"

const router = Router()

router.get("/metadata", fetchMetaData)
router.get("/downlaod", download)

export default router 