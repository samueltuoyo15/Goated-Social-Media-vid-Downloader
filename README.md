# Video Metadata Fetcher

A simple application to fetch metadata from video URLs using `yt-dlp`. It consists of a React frontend and an Express backend.

## Description

This project provides a way to retrieve video metadata, such as title, thumbnail, duration, category, and available download links, from a given video URL. The backend uses `yt-dlp` to extract the metadata, and the frontend provides a user interface to interact with the API.

## Installation

### Frontend (Client)

1.  Navigate to the `client` directory:

    ```bash
    cd client
    ```

2.  Install the dependencies:

    ```bash
    npm install
    ```

### Backend (Server)

1.  Navigate to the `server` directory:

    ```bash
    cd ../server
    ```

2.  Install the dependencies:

    ```bash
    npm install
    ```

## Usage

### Backend (Server)

1.  Start the server:

    ```bash
    cd server
    npm run start
    ```

    (or if you have a `start` script defined in `package.json`, use that.  If not, try `node index.js`)

    The server will start at `http://localhost:10000` (or the port specified in your environment or `index.ts`).

### Frontend (Client)

1.  Start the development server:

    ```bash
    cd client
    npm run dev
    ```

    This will usually start the app at `http://localhost:5173/` or similar.

2.  Open the application in your browser.

3.  Enter the video URL in the input field and submit.

4.  The metadata for the video will be displayed.

## API Endpoint

The backend exposes the following endpoint:

-   `GET /metadata?url={videoUrl}`: Fetches the metadata for the provided video URL.

    Example: `http://localhost:10000/metadata?url=https://www.youtube.com/watch?v=dQw4w9WgXcQ`

## Project Structure

-   `client/`: Contains the React frontend code.
    -   `src/`: Contains the React components, styles, and application logic.
    -   `public/`: Contains static assets like images.
    -   `vite.config.ts`: Vite configuration file.
    -   `tailwind.config.ts`: Tailwind CSS configuration file.
-   `server/`: Contains the Express backend code.
    -   `controller/`: Contains the logic for handling requests.
        -   `downloader.ts`: Contains the function to fetch metadata using `yt-dlp`.
    -   `routes/`: Contains the API routes.
        -   `downloadRoute.ts`: Defines the `/metadata` route.
    -   `index.ts`: The main server file.
    -   `vercel.json`: Configuration file for Vercel deployment.

## Dependencies

### Frontend

-   React
-   Vite
-   Tailwind CSS
-   autoprefixer

### Backend

-   Express
-   helmet

## Environment Variables

-   `PORT`: Specifies the port for the server to listen on (default: 10000).

## Contributing

Contributions are welcome!  If you want to contribute:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes.
4.  Submit a pull request.

## License

This project has no license.
