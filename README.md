# userAuth

## Overview

**userAuth** is a baseline project designed to explore and learn the fundamentals of backend development and user authentication. This project serves as a practical learning tool to understand how to implement basic authentication mechanisms, manage user sessions, and interact with a database in a backend environment.

## Features that I wish to implement

- **User Registration**: Users can sign up with an email and password.
- **User Login**: Authenticated users can log in using their credentials.
- **JWT Authentication**: Secure API endpoints using JSON Web Tokens (JWT).
- **Password Hashing**: User passwords are securely hashed before storage.
- **Session Management**: Manage user sessions to ensure secure access to protected resources.

## Tech Stack

- **Backend**: Go
- **Database**: SQL (Postgres)
- **Authentication**: JWT for token-based authentication
- **Additional Tools**: sqlc for database operations, goose for database migrations

## Getting Started

### Prerequisites

- Go 1.20 or later
- A SQL database (PostgreSQL)
- make (non mandatory)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/YuanziX/userAuth.git
   cd userAuth
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Copy .env.sample to .env and fill in the required fields.
   ```bash
   cp .env.sample .env
   ```

4. Run the application:
   ```bash
   make run
   ```

## Learning Goals

- Understanding backend architecture and design.
- Implementing user authentication securely.
- Gaining hands-on experience with Go and SQL databases.
- Exploring JWT for secure API authentication.

## Contributing

This project is primarily for personal learning, but contributions are welcome. Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.