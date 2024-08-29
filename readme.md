# Cache Service

## Overview

This repository provides a simple, thread-safe, in-memory caching service written in Go. The service supports basic cache operations like setting, getting, and deleting key-value pairs with configurable expiration times.

### Repository Structure

- **`main.go`**: The entry point of the application, demonstrating the usage of the in-memory cache.
- **`cache/`**: Directory containing cache-related modules.
  - **`cache.go`**: Defines the `Cache` interface that includes the methods `Set`, `Get`, and `Delete`.
  - **`inmemorycache/`**: Implements the `Cache` interface using an in-memory data structure. It provides thread safety and expiration handling. It also includes mechanism to clean the expired data but it is commented for simplicity for now.



## How to Run the Code

1. **Clone the repository**:
    ```bash
    git clone https://github.com/SouravChoudhary/cache-service.git
    cd cache-service
    ```

2. **Run the application**:
    ```bash
    go run main.go
    ```

3. **Expected Output**:
   - The program will demonstrate setting and getting cache values, handling key expiration, and deleting keys.
   - You will see the output related to successful retrieval, expiration errors, and key deletion.

