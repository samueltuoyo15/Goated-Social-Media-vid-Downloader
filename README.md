# EverDownload: Secure Social Media Video Downloader 🚀

## Description

EverDownload is a modern web application that allows you to securely download your favorite social media videos. It provides a simple and intuitive interface for fetching video metadata and downloading videos in various qualities. Built with React, TypeScript, Go, and Docker, EverDownload offers a seamless experience for users looking to save online videos for offline viewing.

## Features

- **Secure Downloads**: Safely download videos without compromising your privacy.
- **Multi-Quality Support**: Choose from various video qualities to suit your needs.
- **User-Friendly Interface**: A clean and intuitive design for a seamless user experience.
- **Real-time Metadata Fetching**: Quickly retrieve video details like title, duration, and thumbnail.
- **Cross-Platform Compatibility**: Works on any device with a modern web browser.

## Technologies Used

| Category | Technology   | Description                               |
| :------- | :----------- | :---------------------------------------- |
| Frontend | React        | JavaScript library for building user interfaces |
|          | TypeScript   | Adds static typing to JavaScript           |
|          | Tailwind CSS | CSS framework for styling                 |
|          | Vite         | Build tool for modern web development      |
| Backend  | Go           | Programming language for server-side logic |
| Other    | Docker       | Platform for containerization             |

## Installation

Follow these steps to set up EverDownload locally using Docker:

1.  **Clone the Repository**:

    ```bash
    git clone <repository-url>
    cd EverDownload
    ```

2.  **Build the Docker Image**:

    ```bash
    docker build -t everdownload .
    ```

3.  **Run the Docker Container**:

    ```bash
    docker run -p 10000:10000 everdownload
    ```

## Usage

1.  **Open the Application**:

    Open your web browser and navigate to `http://localhost:10000`.

2.  **Paste the Video URL**:

    Enter the URL of the social media video you want to download into the input field.

    ![Input Field](https://i.imgur.com/your-input-field-screenshot.png)

3.  **Click the "Download" Button**:

    Click the "Download" button to fetch the video metadata.

4.  **Select the Desired Quality**:

    Choose the video quality from the dropdown menu.

    ![Quality Selection](https://i.imgur.com/your-quality-selection-screenshot.png)

5.  **Download the Video**:

    Click the "Download Video" button to start the download.

## Contributing

We welcome contributions from the community! Here's how you can contribute:

1.  **Fork the Repository**:

    Fork the repository on GitHub.

2.  **Set Up Your Development Environment**:

    *   Install Go: [https://go.dev/doc/install](https://go.dev/doc/install)
    *   Install Node.js and npm: [https://nodejs.org/](https://nodejs.org/)

3.  **Make Changes**:

    Create a new branch for your feature or bug fix.

    ```bash
    git checkout -b feature/your-feature
    ```

4.  **Test Your Changes**:

    *   Run the client tests: `cd client && npm install && npm run lint`
    *   Run the server: `cd server && go run main.go`

5.  **Commit Your Changes**:

    ```bash
    git add .
    git commit -m "Add: your feature or fix"
    ```

6.  **Push to Your Fork**:

    ```bash
    git push origin feature/your-feature
    ```

7.  **Submit a Pull Request**:

    Submit a pull request to the main repository.

## License

This project is open source and available under the [MIT License](LICENSE).

[![Built with Dokugen](https://img.shields.io/badge/Built%20with-Dokugen-brightgreen)](https://github.com/samueltuoyo15/Dokugen)
