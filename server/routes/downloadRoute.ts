const {fetchMetaData} = require("../controller/downloader")
const {Router} = require("express")

const router = Router()

router.get("/metadata", fetchMetaData)

module.exports = router