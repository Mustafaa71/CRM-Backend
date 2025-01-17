# CRM-Backend

## Overview

This project involves building the backend (i.e., server-side portion) of a CRM application using Go. Users can make HTTP requests to your server to perform CRUD operations. This project is designed to solidify your skills in building an HTTP server, designing RESTful endpoints, and handling business logic for API requests.



### Main Components

1. **HTTP Server**: The core of the application, handling incoming HTTP requests.
2. **Router**: Using `gorilla/mux` for method-based routing and handling URL path variables.
3. **Mock Database**: A map structure to store customer data.
4. **Handler Functions**: Functions to handle CRUD operations on customer data.

### Endpoints

- **GET /customers**: Retrieve all customers.
- **GET /customers/{id}**: Retrieve a single customer by ID.
- **POST /customers**: Add a new customer.
- **PATCH /customers/{id}**: Update an existing customer.
- **DELETE /customers/{id}**: Delete a customer by ID.

## Setup and Running

1. Install the necessary dependencies:
    ```sh
    go get -u github.com/gorilla/mux
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

## Contributing

Feel free to open issues and submit pull requests if you want to contribute to this project. Please ensure your contributions adhere to the existing code style and include appropriate tests.

