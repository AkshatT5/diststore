# DistStore

DistStore is a robust, distributed file storage system implemented in Go. It provides a simple web interface and RESTful API for storing, retrieving, and managing files across a network of nodes.

## Features

- Distributed file storage across multiple nodes
- Content-addressable storage (CAS) for efficient file management and deduplication
- Simple web interface for file operations
- RESTful API for programmatic access
- File encryption for enhanced security
- Peer-to-peer communication for distributed storage and retrieval

## How It Works

DistStore uses a content-addressable storage (CAS) system to manage files:

1. When a file is stored, its content is hashed to generate a unique key.
2. The key determines where the file is stored in the system.
3. The file is encrypted and distributed across the network.

This approach ensures efficient storage, retrieval, and deduplication of content.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/AkshatT5/diststore.git
   cd diststore
   ```

2. Build the project:
   ```
   make build
   ```

3. Run the application:
   ```
   make run
   ```

4. Access the web interface at `http://localhost:8080`

## Usage

### Web Interface

The web interface provides a user-friendly way to interact with DistStore. You can:

- Upload files
- Download files using their keys
- Delete files
- View a list of stored files

### API

DistStore provides a RESTful API for programmatic access. For detailed API documentation, refer to the [API Documentation](api_docs.md).

Key endpoints:
- `POST /api/store`: Store a file
- `GET /api/get?key=<file_key>`: Retrieve a file
- `DELETE /api/delete?key=<file_key>`: Delete a file
- `GET /api/list`: List all stored files

## Configuration

The application can be configured using environment variables:

- `PORT`: The port on which the web server will listen (default: 8080)

## Architecture

DistStore is built with a modular architecture:

- `main.go`: Entry point of the application, sets up the network of nodes
- `server.go`: Implements the core FileServer functionality
- `store.go`: Handles the local file storage operations
- `web.go`: Implements the web server and API endpoints

The system uses a peer-to-peer network for distributed storage and retrieval, with each node capable of storing files and forwarding requests to other nodes.

## Development

To set up the development environment:

1. Install dependencies:
   ```
   go mod download
   ```

2. Run tests:
   ```
   make test
   ```

3. For local development, you can run multiple nodes by modifying the `main.go` file.

## Deployment

DistStore can be easily deployed using Docker:

1. Build the Docker image:
   ```
   docker build -t diststore .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 diststore
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all contributors who have helped shape DistStore
- Inspired by distributed systems and content-addressable storage concepts

## Contact

For any queries or support, please open an issue in the GitHub repository.