# Object Storage using MinIO

A simple Docker Compose project to run MinIO as an object storage server.

## Prerequisites

- Docker
- Docker Compose

## How to use

1. Clone this repository
2. Run `docker-compose up` to start the MinIO server
3. Access the MinIO server through `http://localhost:9010`
4. Use the default credentials: `minio` as username and `minio123` as password

## Features

- Run MinIO server as an object storage server
- Expose MinIO server to the host machine through port 9010
- Use default credentials: `minio` as username and `minio123` as password
- Store data in a persistent volume

## Configuration

You can customize the configuration by editing the `docker-compose.yml` file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

2024
