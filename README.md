# Go Gin GORM Boilerplate

A production-ready Go boilerplate designed with **Clean Architecture** principles, scalability, and developer experience in mind. Built on top of Gin framework and GORM.

---

## Features

- **Standard Go Layout**: Follows the community-driven [project layout](https://github.com/golang-standards/project-layout).
- **Clean Architecture**: Separation of concerns using **Handler -> Service -> Repository** layers.
- **UUID Ready**: Uses `google/uuid` for secure, non-predictable primary keys.
- **Global Middleware**: Includes ready-to-use authentication and custom middleware examples.
- **Configuration Management**: Effortless environment management using `.env`.
- **GORM Integrated**: Pre-configured for PostgreSQL with auto-migration support.
- **Internal/External Pattern**: Strict usage of the `internal/` package pattern for high-quality encapsulation.

---

## Project Structure

```text
├── cmd/
│   ├── api/            # Web API entry point
│   └── worker/         # Background worker entry point
├── internal/
│   ├── app/            # Domain logic (Example: User, Product)
│   │   └── example/    # Layered logic (Model, Repo, Service, Handler)
│   ├── config/         # DB & Env configurations
│   ├── middleware/     # Custom HTTP middlewares
│   ├── route/          # Centralized route definitions
│   └── utils/          # Internal helper functions
├── pkg/                # Public libraries (sharable with other projects)
├── .env                # Environment variables
├── go.mod              # Dependency management
└── README.md
```

---

## Tech Stack

- **Language**: [Go](https://go.dev/) (1.25+)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL (Default)
- **ID Generation**: UUID v4

---

## Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/woroagung/boilerplate-go-gin-gorm.git
cd boilerplate-go-gin-gorm
```

### 2. Setup environment variables
Create a `.env` file in the root directory:
```env
APP_NAME=YOUR_APP_NAME
APP_ENV=local
APP_KEY=your_secret_key
APP_PORT=8080

DB_HOST=localhost
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=yourdb
DB_PORT=5432

JWT_SECRET=your_secret_key
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the application
```bash
go run cmd/api/main.go
```

---

## Contributing

Feel free to fork this project and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

---

## Author

**Woro A**
- GitHub: [@woroagung](https://github.com/woroagung)
