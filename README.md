# Tempestas - Cloud-Based Weather Service

Tempestas is a cloud-based weather service built to provide efficient weather data access with Kubernetes support. It utilizes Redis as a caching mechanism to minimize API calls to the main weather API, ensuring smooth and efficient service operation.

## Features

- **Real-Time Weather Data**: Fetches real-time weather information from a primary weather API.
- **Caching Mechanism**: Utilizes Redis to cache weather data, reducing redundant API calls and improving response times.
- **Microservice Architecture**: Designed to operate efficiently in a microservices environment.
- **Kubernetes Support**: Full compatibility with Kubernetes for deployment, scaling, and management.
- **High Availability**: Ensures service availability through robust architecture and caching.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) (for containerization)
- [Kubernetes](https://kubernetes.io/) (for deployment)
- [Redis](https://redis.io/) (for caching)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Manni-MinM/tempestas.git
   cd tempestas
   ```

2. Build the Docker image:

   ```bash
   docker build -t tempestas .
   ```

3. Deploy to Kubernetes using Helm or kubectl:

   ```bash
   kubectl apply -f k8s/deployment.yaml
   ```

### Configuration

Edit the configuration files in the `config` directory to set up your weather API keys, Redis settings, and other parameters.

### Usage

Once deployed, the service will be accessible via the configured endpoints. You can request weather data, and the service will utilize Redis for caching to enhance performance.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
