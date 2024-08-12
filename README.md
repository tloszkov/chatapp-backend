# ChatApp API

This is a sample server for a chat application built with Go and Gin framework. It provides endpoints for managing users, messages, group messages, and message boards.

## Prerequisites

- Go 1.16 or higher
- MongoDB
- Git

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/ChatApp.git
    cd ChatApp
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up MongoDB and update the connection details in `DBConnector.go`.

## Running the Project

1. Start the server:

    ```sh
    go run main.go
    ```

2. The server will run on `http://localhost:8090`.

## Swagger Documentation

Access the Swagger documentation at:
http://localhost:8090/swagger/index.html

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
