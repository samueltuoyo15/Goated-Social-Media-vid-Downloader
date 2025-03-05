# Palm Download 🌴⬇️

A modern and secure web application for effortlessly downloading your favorite videos from social media platforms. Enjoy high-quality downloads with a simple and intuitive user interface.

## ✨ Features

-   **Effortless Downloads**: Easily download videos by pasting the URL.
-   **Quality Selection**: Choose from various available video qualities to suit your needs.
-   **Secure**: Built with security in mind to ensure a safe downloading experience.
-   **Sleek UI**: A modern and user-friendly interface.
-   **Free**: Completely free to use.

## 🛠️ Technologies Used

| Category        | Technology                               | Description                                                                 |
| :-------------- | :--------------------------------------- | :-------------------------------------------------------------------------- |
| Frontend        | React 19.0.0                             | Building the user interface.                                                |
| Frontend        | TypeScript                               | Ensuring type safety and improving code quality.                            |
| Frontend        | Vite 6.2.0                               | Fast build tool and development server.                                     |
| Frontend        | Tailwind CSS                             | Utility-first CSS framework for styling.                                     |
| Backend         | Go                                       | Backend logic and API handling.                                             |
| Backend         | Gin Gonic                                | Web framework for Go.                                                        |
| Containerization| Docker                                   | For containerizing the application.                                          |
| Other           | yt-dlp                                 | Command-line program to download videos from YouTube and other sites          |

## 📦 Installation

Follow these steps to get the project up and running on your local machine:

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
    docker run -p 3000:3000 palm-download
    ```

    *Note: The application will be running on port 3000.*

## 🚀 Usage

1.  Open your browser and navigate to `http://localhost:3000`.
2.  Paste the URL of the video you wish to download into the input field.
3.  Click the "Download" button.
4.  If available, select your desired video quality from the dropdown menu.
5.  Click the "Download Video" button to start the download.

## 🤝 Contributing

Contributions are welcome! Here's how you can contribute to the project:

1.  **Fork the repository.**
2.  **Create a new branch** for your feature or bug fix:

    ```bash
    git checkout -b feature/your-feature-name
    ```

3.  **Make your changes and commit them:**

    ```bash
    git add .
    git commit -m "Add your descriptive commit message"
    ```

4.  **Push your changes to your forked repository:**

    ```bash
    git push origin feature/your-feature-name
    ```

5.  **Submit a pull request** to the main repository.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

[![Built with Dokugen](https://img.shields.io/badge/Built%20with-Dokugen-brightgreen)](https://github.com/samueltuoyo15/Dokugen)
