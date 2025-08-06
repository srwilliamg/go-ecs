# Go ECS - Clean Architecture REST API

A modern Go REST API application built with Clean Architecture principles, providing a robust and maintainable foundation for enterprise applications.

## 🏗️ Clean Architecture Advantages

This project implements **Clean Architecture** (also known as Hexagonal Architecture), offering several key benefits:

### 🎯 **Separation of Concerns**
- **Domain Layer**: Contains pure business logic, isolated from external dependencies
- **Application Layer**: Coordinates between domain and infrastructure layers
- **Infrastructure Layer**: Handles external concerns (databases, web frameworks, logging)
- **Interface Layer**: Adapts external dependencies to internal interfaces

### 🔄 **Dependency Inversion**
- Dependencies point inward toward the domain layer
- External frameworks and tools are plugins to the business logic
- Easy to swap implementations without affecting core business rules

### 🧪 **Enhanced Testability**
- Each layer can be tested in isolation
- Business logic can be tested without databases or web frameworks
- Mock implementations can easily replace external dependencies

### 🔧 **Maintainability & Flexibility**
- Changes to external services don't affect business logic
- Easy to add new features without breaking existing functionality
- Framework-agnostic core business logic
- Clear boundaries between different concerns

### 📈 **Scalability**
- Well-defined interfaces enable team collaboration
- Independent development of different layers
- Easy to extend with new use cases and features

## 🚀 Technologies Used

### **Core Technologies**
- **Go 1.24.3** - Modern, fast, and efficient programming language
- **Chi Router** - Lightweight HTTP router for building REST APIs
- **PostgreSQL** - Robust relational database for data persistence

### **Database & ORM**
- **lib/pq** - Pure Go PostgreSQL driver
- **SQLX** - Extensions for Go's database/sql package
- **Squirrel** - SQL query builder for dynamic query construction

### **Logging & Validation**
- **Zap** - Blazing fast, structured logging
- **Validator** - Input validation with struct tags

### **Security**
- **bcrypt** - Password hashing for secure authentication

### **DevOps & Deployment**
- **Docker** - Containerization for consistent deployment
- **Docker Compose** - Multi-container application orchestration
- **Makefile** - Build automation and development workflow

## 📋 Prerequisites

- **Go 1.24.3+**
- **Docker & Docker Compose**
- **PostgreSQL** (if running locally without Docker)

## 🛠️ Installation & Setup

### **1. Clone the Repository**
```bash
git clone https://github.com/srwilliamg/go-ecs.git
cd go-ecs
```

### **2. Using Docker Compose (Recommended)**
```bash
# Start the entire application stack
make start-compose

# Or manually with docker-compose
docker-compose up -d
```

### **3. Local Development Setup**
```bash
# Start only the database
make start-db

# Install dependencies
go mod tidy

# Run the application locally
make run
```

### **4. Build Options**
```bash
# Build binary
make build

# Build Docker image
make build-docker

# Clean build artifacts
make clean
```

## 🎯 Available Make Commands

```bash
make help              # Display all available commands
make fmt               # Format Go code
make tidy              # Update dependencies
make lint              # Run linter (requires ./bin/lint.sh)
make test              # Run unit tests
make run               # Run application locally
make build             # Build binary
make build-docker      # Build Docker image
make start-compose     # Start with Docker Compose
make stop-compose      # Stop Docker Compose
make start-db          # Start only database
```

## 🌐 API Endpoints

### **Users API**
```http
GET /users          # Get all users
POST /users         # Create a new user
```

### **Health Check**
```http
GET /              # Application health check
```

### **Example Usage**

**Create a User:**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure_password"
  }'
```

**Get All Users:**
```bash
curl http://localhost:8080/users
```

## 📁 Project Structure

```
go-ecs/
├── cmd/                          # Application entry points
├── internal/                     # Private application code
│   ├── domain/                   # 🔵 Domain Layer (Business Logic)
│   │   ├── entities/             # Business entities
│   │   └── use-case/             # Business use cases
│   ├── application/              # 🟢 Application Layer (Orchestration)
│   │   ├── controller/           # HTTP controllers
│   │   ├── dto/                  # Data Transfer Objects
│   │   ├── middleware/           # HTTP middleware
│   │   ├── routes/               # Route definitions
│   │   ├── request/              # Request/response utilities
│   │   └── validator/            # Input validation
│   ├── infrastructure/           # 🟡 Infrastructure Layer (External Concerns)
│   │   ├── config/               # Configuration management
│   │   ├── db/                   # Database connections
│   │   └── logger/               # Logging implementation
│   ├── interfaces/               # 🟣 Interface Adapters
│   │   ├── db/                   # Database interfaces
│   │   ├── logger/               # Logger interfaces
│   │   └── repository/           # Repository interfaces
│   └── repositories/             # 🔴 Repository Implementations
│       └── users/                # User repository
├── config/                       # Configuration files
│   └── init.sql                  # Database initialization
├── bin/                          # Binary files and scripts
├── docker-compose.yml            # Docker Compose configuration
├── Dockerfile                    # Docker image definition
├── Makefile                      # Build automation
├── go.mod                        # Go module definition
└── main.go                       # Application entry point
```

### **Architecture Layers Explained**

1. **🔵 Domain Layer** (`internal/domain/`)
   - Contains pure business logic and entities
   - No dependencies on external frameworks
   - Defines interfaces for external dependencies

2. **🟢 Application Layer** (`internal/application/`)
   - Orchestrates between domain and infrastructure
   - Contains controllers, DTOs, and application services
   - Handles HTTP concerns and request/response formatting

3. **🟡 Infrastructure Layer** (`internal/infrastructure/`)
   - Implements external concerns (database, logging, config)
   - Contains framework-specific implementations
   - Provides concrete implementations of domain interfaces

4. **🟣 Interface Adapters** (`internal/interfaces/`)
   - Defines contracts between layers
   - Enables dependency inversion
   - Facilitates testing with mock implementations

5. **🔴 Repository Layer** (`internal/repositories/`)
   - Data access implementations
   - Encapsulates database operations
   - Implements repository interfaces from domain layer

## 🔧 Development Guidelines

### **Code Organization**
- Keep business logic in the domain layer
- Use dependency injection for loose coupling
- Define interfaces in the domain layer, implement in infrastructure
- Keep controllers thin - delegate to use cases

### **Database Migrations**
- Database schema is defined in `config/init.sql`
- Modify schema changes through proper migration scripts
- Use transactions for data consistency

### **Error Handling**
- Use custom error types for business logic errors
- Handle errors at appropriate layers
- Provide meaningful error messages to API consumers

### **Testing Strategy**
- Unit tests for domain layer (business logic)
- Integration tests for repositories
- End-to-end tests for API endpoints
- Use mocks for external dependencies

## 🌍 Environment Configuration

The application uses environment variables for configuration:

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASS=pass
DB_NAME=gopi-db

# Application Configuration
APP_ENV=development
PORT=8080
LOG_LEVEL=info
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Follow the Clean Architecture principles
4. Add tests for new functionality
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 📞 Support

For questions or support, please open an issue in the GitHub repository.

---

**Built with ❤️ using Clean Architecture principles in Go**