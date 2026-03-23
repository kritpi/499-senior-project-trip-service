# Keep in Trip - Backend Service

This is the backend service for the **Keep in Trip** senior project. It is built using **Go** and utilizes **Redis** for data storage and caching. The service is fully containerized for easy development and deployment.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Installation & Running (using Docker)

The project includes a `docker-compose.yml` file that sets up all the required services seamlessly, including the Go application server, a Redis instance, and RedisInsight.

1. **Clone the repository** (if you haven't already) and navigate to the project directory:
   ```bash
   cd 499-senior-project-trip-service
   ```

2. **Configure the application** by copying the example configuration file:
   ```bash
   cp config/config.example.json config/config.json
   ```
   *(Edit `config/config.json` if necessary to match your local environment).*

3. **Build and start the services** using Docker Compose:
   ```bash
   docker-compose up --build
   ```
   *(To run the containers in the background, append `-d` to the command).*

4. **Verify the services are running**:
   - **Backend API**: Accessible at `http://localhost:8080`
   - **RedisInsight** (GUI for managing Redis): Accessible at `http://localhost:5540`
   - **Redis Store**: Running on `localhost:6379`

## Stopping the Services

To stop the running containers, simply press `Ctrl+C` in the terminal where it's running, or run the following command if you started it in detached mode:
```bash
docker-compose down
```

## Configuration

The application loads its configuration from the `./config` directory. This directory is mounted as a volume in the Docker container, meaning you can update configuration files locally and restart the service to apply changes.
