import { fetchMetaData } from "../controller/downloader"
import { Router } from "express"

const router = Router()

router.get("/metadata", fetchMetaData)

export default router 