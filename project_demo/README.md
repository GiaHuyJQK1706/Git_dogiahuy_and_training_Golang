# Project demo
## Lưu ý
- Các bài tập là các commit trong một pull request, có dạng là `GOLANG project demo bai x`
## Cấu trúc thư mục
```
project_demo/
│
├── dto/           # Data Transfer Objects
│   ├── project_dto.go
│   ├── user_dto.go
│
├── entities/      # Database Entities
│   ├── project.go
│   ├── project_user.go
│
├── middleware/    # Middleware (JWT validation, etc.)
│   ├── jwt_middleware.go
│
├── handler/       # HTTP Handlers (controllers)
│   ├── project_handler.go
│   ├── user_handler.go
│
├── repository/    # Data Access Layer
│   ├── project_repository.go
│   ├── user_repository.go
│
├── usecase/       # Business Logic Layer
│   ├── project_usecase.go
│   ├── user_usecase.go
│
├── validators/    # Custom Validators
│   ├── custom_validator.go
│
├── utils/         # Helper functions (e.g., CSV generation, JWT utils)
│   ├── csv_utils.go
│   ├── jwt_utils.go
│
├── config/        # Database and App Configurations
│   ├── database.go
│
├── routes/        # API Routes
│   ├── routes.go
│
├── tests/         # Unit Tests
│   ├── project_usecase_test.go
│   ├── project_repository_test.go
│
└── main.go        # Entry point for the application
```
