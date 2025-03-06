# Palm Download: Securely Download Social Media Videos 🌴

Tired of losing your favorite social media videos to the ever-changing internet landscape? Palm Download is here to help! This project provides a simple and secure way to download videos from platforms like YouTube, ensuring you can enjoy your favorite content offline, anytime, anywhere.

## ✨ Key Features

-   **Simple URL Input**: Just paste the video URL, and let Palm Download do the rest.
-   **Quality Selection**: Choose the video quality that suits your needs.
-   **Secure Downloads**: Downloads are direct and secure, ensuring your privacy.
-   **Modern UI**: A clean and intuitive interface built with React and Tailwind CSS.
-   **Cross-Platform**: Works seamlessly across different operating systems thanks to Docker.

## 🛠️ Technologies Used

| Category    | Technology          | Description                                                              |
| :---------- | :------------------ | :----------------------------------------------------------------------- |
| Frontend    | React               | A JavaScript library for building user interfaces.                      |
|             | TypeScript          | Adds static typing to JavaScript for improved code quality.            |
|             | Tailwind CSS        | A utility-first CSS framework for rapid UI development.                |
|             | Vite                | A fast build tool for modern web development.                            |
| Backend     | Go                  | A fast and efficient programming language for backend logic.             |
|             | Gin Gonic           | A high-performance HTTP web framework for Go.                           |
| Other       | Docker              | A platform for developing, shipping, and running applications in containers. |
|             | yt-dlp              | A command-line program to download videos from YouTube and other sites. |

## 📦 Installation

Follow these steps to get Palm Download up and running on your local machine:

1.  **Clone the repository:**

    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```

2.  **Build the Docker image:**

    ```bash
    docker build -t palm-download .
    ```

3.  **Run the Docker container:**

    ```bash
    docker run -p 10000:10000 palm-download
    ```

4.  **Access the application:**

    Open your web browser and navigate to `http://localhost:10000`.

## 🚀 Usage

1.  **Enter the Video URL:** Paste the URL of the video you want to download into the input field.
2.  **Click "Download":** Submit the form. The application will fetch the video metadata.
3.  **Select Quality:** Choose your desired video quality from the dropdown menu.
4.  **Download Video:** Click the "Download Video" button to start the download.

Here's a quick example:

![Palm Download Interface](https://i.imgur.com/placeholder.png)

*Replace the above image URL with an actual screenshot of the application interface.*

## 🤝 Contributing

We welcome contributions to Palm Download! Here's how you can help:

1.  **Fork the repository.**
2.  **Create a new branch** for your feature or bug fix:

    ```bash
    git checkout -b feature/your-feature-name
    ```

3.  **Make your changes** and commit them with clear, concise messages.
4.  **Push your branch** to your forked repository.
5.  **Submit a pull request** to the main repository.

## 📜 License

This project is licensed under the [MIT License](LICENSE) (if applicable).

[![Built with Dokugen](https://img.shields.io/badge/Built%20with-Dokugen-brightgreen)](https://github.com/samueltuoyo15/Dokugen)
