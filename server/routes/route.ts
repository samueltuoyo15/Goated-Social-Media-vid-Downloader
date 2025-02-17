import { fetchMetaData, download } from "../controller/downloader"
import { Router } from "express"

const router = Router()

router.get("/metadata", fetchMetaData)
router.get("/downlaod", download)

export default router 