import express, {Application} from 'express';
import { fileURLToPath } from 'url';
import path, {dirname} from 'path';
import helmet from "helmet";
import downloadRoute from "./routes/downloadRoute";

const app: Application = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(helmet());
app.use(downloadRoute);

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)
app.use(express.static(path.join(__dirname, '..', 'client', 'dist')));


const PORT = process.env.PORT || 10000; 
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
