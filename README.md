# Service Discovery in Go Using RPC

This project demonstrates a Service Discovery mechanism in Golang using the built-in RPC (Remote Procedure Call) protocol. It consists of multiple services, including `Order Service`, `Product Service`, and `Service Discovery`. Each service must register with the `Service Discovery` service, monitor its health, and dynamically discover other services for interaction.

## Project Structure

- **Order Service**: Handles customer orders and interacts with the `Product Service`.
- **Product Service**: Manages product information and inventory.
- **Service Discovery**: Registers, monitors, and deregisters services dynamically.
- **Additional Services**: More services can be added, following the same procedure for registration and interaction.

### Key Features

- **Service Registration**: Each service (e.g., `Order`, `Product`, or any new service) must register itself with the `Service Discovery` service at startup by providing its name, address, and health check endpoint.
- **Health Check**: Every 5 seconds, the `Service Discovery` service pings the `/health` endpoint of registered services. If a service becomes unhealthy, it is removed from the registry.
- **Dynamic Service Lookup**: Any registered service can dynamically discover other services by asking the `Service Discovery` service for the address of another service using RPC. Once the address is retrieved, the services can interact using HTTP endpoints.

### Adding New Services

To add a new service:
1. **Register with Service Discovery**: The new service must send its name, address, and health check endpoint to the `Service Discovery` service during startup.
2. **Health Check**: The service should expose a `/health` endpoint that can be periodically checked by the `Service Discovery` service.
3. **Interaction with Other Services**: The new service can dynamically request the address of other registered services using RPC, then interact with those services via HTTP endpoints.

## Redis Integration

The `Service Discovery` service uses Redis to store registered service information in a hash set (HSET) with the following details:
- **Service Name**: Identifies the service (e.g., `product`, `order`, `inventory`, etc.).
- **Address**: The service's address (IP and port).
- **Health Check Endpoint**: The endpoint used to verify the service's health.

## Running the Project

### Prerequisites

- **Go**: Version 1.22+ is required.
- **Docker**: Ensure Docker is installed to run each service in its own container.
- **Redis**: Redis is used for storing service registration data.

## Docker Setup

Each service is containerized using Docker. Every service has:
- **A Dockerfile**: Defines how the container is built for each service.
- **A Bash Script**: Automates the build and run process of the container for each service.

### Steps to Run

1. Clone the repository:
    ```powershell
    git clone https://github.com/rohanyh101/service-discovery/ service-discovery
    cd service-discovery
    ```

2. Start Redis:
    ```powershell
    ./service-discovery/database/build.sh
    ```

3. Run the `Service Discovery` service:
    ```powershell
    ./service-discovery/build.sh
    ```

4. Run any other services (e.g., `Product Service`, `Order Service`):
    ```powershell
    ./service-discovery/product_service/build.sh
    ./service-discovery/order_service/build.sh
    ```
    
Each `build.sh` script in the respective service directories will:
- Build the Docker image.
- Remove the previous container, if exists.
- Start the container.
- Ensure the service registers itself with the `Service Discovery` service.

### Service Registration

At startup, each service (e.g., `Product`, `Order`, or any new service) registers itself with the `Service Discovery` service by providing the following details:
- **Service Name**: A unique identifier for the service (e.g., `product`, `order`).
- **Address**: The IP address and port the service is running on.
- **Health Check Endpoint**: The HTTP endpoint to check if the service is healthy (e.g., `/health`).

### Health Checks

The `Service Discovery` service checks each registered service's `/health` endpoint every 5 seconds. If a service becomes unresponsive or unhealthy, it is removed from Redis and is no longer discoverable.

### Dynamic Service Lookup Example

Any service can request the address of another service by calling the `Discovery.Get` method:

```powershell
var addr net.TCPAddr
if err := client.Call("Discovery.Get", "product", &addr); err != nil {}
```

Once the service address is obtained, HTTP requests can be made to interact with the target service.

### Example Endpoint Usage
Purchasing a Product (via Order Service):
```powershell
GET /purchase/{product_id}: process the purchase (decrement the product quantity by 1)
```

### Product Service Endpoints:
```powershell
GET /product/{product_id}: Retrieve product details.
PUT /product/{product_id}: Update product information (decrement the product quantity by 1).
```

### Extending the System
To add a new service (e.g., `Inventory Service`):

1. **Register the Service:** The new service must register itself with the `Service Discovery` by providing its name, address, and `/health` endpoint.
2. **Health Monitoring:** Ensure the service exposes a `/health` endpoint for periodic health checks by the Service Discovery.
3. **Service Discovery:** When the new service needs to interact with other services, it can use the `Discovery.Get` RPC call to obtain the address of the required service.
4. **Inter-Service Communication:** Once the address is retrieved, HTTP requests can be made to interact with other services' endpoints.

### Technologies Used
 - **Golang:** For building all services.
 - **RPC (Remote Procedure Call):** For communication between services and Service Discovery.
 - **Redis:** For storing service registration data.
 - **Docker:** For containerizing each service.
 - **HTTP:** For interactions between services after address resolution.

### Contributing
Contributions are welcome! If you have any ideas for improvements or new features, feel free to fork the repository and submit a pull request.
