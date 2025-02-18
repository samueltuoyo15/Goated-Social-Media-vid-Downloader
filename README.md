# Real Time Social Media Video Downloader

A simple application to fetch metadata from video URLs using `yt-dlp`. This project consists of a React frontend and an Express.js backend.

## Description

This application allows users to input a video URL and retrieve metadata such as title, thumbnail, duration, category, and available download links with their qualities. It leverages `yt-dlp` on the backend to extract the video information.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Docker Setup](#docker-setup)
- [Contributing](#contributing)
- [License](#license)

## Installation

Follow these steps to set up the project locally:

1.  **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd <project-directory>
    ```

2.  **Install server dependencies:**

    ```bash
    cd server
    npm install
    cd ..
    ```

3.  **Install client dependencies:**

    ```bash
    cd client
    npm install
    cd ..
    ```

4.  **Install `yt-dlp`**:

    Make sure you have `yt-dlp` installed on your system. You can find installation instructions on the [yt-dlp GitHub page](https://github.com/yt-dlp/yt-dlp).

    On macOS, you can use Homebrew:

    ```bash
    brew install yt-dlp
    ```

    On Linux:

    ```bash
    sudo apt install yt-dlp
    ```

## Usage

1.  **Start the server:**

    ```bash
    cd server
    npm run start
    cd ..
    ```

    The server will run on `http://localhost:10000` by default.

2.  **Start the client:**

    ```bash
    cd client
    npm run dev
    cd ..
    ```

    The client will typically run on `http://localhost:5173`.  Check the console output from `npm run dev` for the exact URL.

3.  **Open the application in your browser:**

    Navigate to the client URL (e.g., `http://localhost:5173`).

4.  **Enter a video URL:**

    Enter the URL of the video you want to fetch metadata from in the input field.

5.  **Fetch Metadata:**

    The application will display the video's metadata, including its title, thumbnail, duration, category, and available download links.

## Docker Setup

A `Dockerfile` is not provided but you can containerize the application using Docker with a few steps. Here's a basic outline:

1.  **Create a `Dockerfile` in the root directory:**

    ```dockerfile
    # Use a Node.js runtime as a base image
    FROM node:18

    # Set the working directory in the container
    WORKDIR /app

    # Copy package.json and package-lock.json to the working directory
    COPY package*.json ./

    # Install server dependencies
    RUN npm install

    # Copy the server code to the working directory
    COPY server ./server

    # Install client dependencies
    COPY client/package*.json ./client/
    WORKDIR /app/client
    RUN npm install
    WORKDIR /app

    # Copy the client code to the working directory
    COPY client ./client

    # Build the client
    RUN npm run build --prefix client

    # Expose the server port
    EXPOSE 10000

    # Command to start the server
    CMD ["npm", "start", "--prefix", "server"]
    ```

2.  **Build the Docker image:**

    ```bash
    docker build -t video-metadata-fetcher .
    ```

3.  **Run the Docker container:**

    ```bash
    docker run -p 10000:10000 video-metadata-fetcher
    ```

    Now, the application should be running in a Docker container, accessible at `http://localhost:10000`.  Remember that you'll still need `yt-dlp` available inside the container.  You might want to add `RUN apt-get update && apt-get install -y yt-dlp` to your Dockerfile.

## Contributing

Contributions are welcome!  Here's how you can contribute:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Make your changes and commit them with descriptive commit messages.
4.  Test your changes thoroughly.
5.  Submit a pull request.

## License

This project does not include a license file. As such, all rights are reserved by default.
