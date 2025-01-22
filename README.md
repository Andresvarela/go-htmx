# README

## Project Description

This project is a web application built using the Go programming language and the Gin web framework. It provides a simple user management system with functionalities to create, read, and delete users. The application interacts with a PostgreSQL database to store user information.

## Using

### Prerequisites

- Go 1.20 or later
- PostgreSQL
- Git

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-htmx.git
    cd go-htmx
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the PostgreSQL database:
    ```sh
    psql -U postgres -c "CREATE DATABASE users_operations;"
    psql -U postgres -d users_operations -c "CREATE TABLE usuarios (nombre VARCHAR(50), email VARCHAR(50));"
    ```

4. Run the application:
    ```sh
    go run core/cmd/main.go
    ```

5. Open your browser and navigate to `http://localhost:8080`.

## Implementation

### Project Structure

- `core/cmd/main.go`: Entry point of the application.
- `infra/primary/controllers/users.go`: Contains the `UsersController` which handles HTTP requests.
- `infra/secondary/persistence/dao/users.go`: Contains the `UserDAO` which interacts with the PostgreSQL database.
- `infra/templates/`: Contains HTML templates for rendering user data.

### Main Components

#### `UsersController`

Handles HTTP requests for user operations:
- `GET /`: Fetches and displays all users.
- `POST /users`: Creates a new user.
- `DELETE /users/:name`: Deletes a user by name.

#### `UserDAO`

Interacts with the PostgreSQL database to perform CRUD operations:
- `FindAllUsers()`: Retrieves all users from the database.
- `CreateUser(user User)`: Inserts a new user into the database.
- `DeleteUser(name string)`: Deletes a user from the database by name.

### Example Usage

1. **Fetch all users**:
    - Navigate to `http://localhost:8080` to see a list of all users.

2. **Create a new user**:
    - Send a `POST` request to `http://localhost:8080/users` with form data `name` and `email`.

3. **Delete a user**:
    - Send a `DELETE` request to `http://localhost:8080/users/{name}` where `{name}` is the name of the user to be deleted.