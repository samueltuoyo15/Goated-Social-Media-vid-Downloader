const express = require("express");
const helmet = require("helmet");
const path = require("path");

const downloadRoute = require("./routes/downloadRoute");

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(helmet());
app.use(downloadRoute);

app.use(express.static(path.join(__dirname, "..", "client", "dist")));

const PORT = process.env.PORT || 10000;
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});