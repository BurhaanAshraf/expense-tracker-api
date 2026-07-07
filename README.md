# Expense Tracker API (Go)

A production-inspired RESTful Expense Tracker API built with Go using the standard `net/http` package. The project follows a layered architecture with JWT authentication, MySQL, request validation, middleware, and analytics endpoints to demonstrate clean backend engineering practices.

---

## Features

### Authentication

- User Registration
- User Login
- Password Hashing using bcrypt
- JWT Authentication
- Protected Routes

### Expense Management

- Create Expense
- Get All Expenses
- Get Expense by ID
- Update Expense
- Delete Expense

### Analytics

- Dashboard Summary
- Category-wise Expense Summary

### Query Capabilities

- Filter expenses by category
- Pagination support
- Sort expenses by amount (ascending/descending)
- Sort expenses by date (ascending/descending)

### Production Features

- Layered Architecture
- Repository Pattern
- Service Layer
- DTOs (Request/Response Objects)
- Request Validation
- Standardized API Responses
- Request Logging Middleware
- Graceful Shutdown
- Environment-based Configuration

---

## Architecture

```
                Client
                   │
             HTTP Request
                   │
                   ▼
          Logging Middleware
                   │
                   ▼
            JWT Middleware
                   │
                   ▼
              HTTP Handler
                   │
                   ▼
             Service Layer
                   │
                   ▼
           Repository Layer
                   │
                   ▼
                 MySQL
```

---

## Project Structure

```
expense-tracker-api/
├── cmd/
│   └── api/
├── internal/
│   ├── auth/
│   ├── config/
│   ├── database/
│   ├── dto/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── response/
│   ├── service/
│   └── validation/
├── .env.example
├── README.md
├── go.mod
└── go.sum
```

---

## Tech Stack

- Go
- MySQL
- JWT
- bcrypt
- net/http
- database/sql

---

## Environment Variables

Create a `.env` file using the provided `.env.example`.

```env
APP_NAME=Finance API
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=finance_db

JWT_SECRET=your_secure_jwt_secret_here
```

---

## Running Locally

Clone the repository:

```bash
git clone https://github.com/BurhaanAshraf/expense-tracker-api.git
```

Navigate to the project:

```bash
cd expense-tracker-api
```

Install dependencies:

```bash
go mod tidy
```

Create the environment file:

```bash
cp .env.example .env
```

Update the values in `.env` with your MySQL credentials and JWT secret.

Run the application:

```bash
go run ./cmd/api
```

The server will start on:

```
http://localhost:8080
```

---

## API Endpoints

### Authentication

| Method | Endpoint    |
| ------ | ----------- |
| POST   | `/register` |
| POST   | `/login`    |

### Expenses

| Method | Endpoint         |
| ------ | ---------------- |
| POST   | `/expenses`      |
| GET    | `/expenses`      |
| GET    | `/expenses/{id}` |
| PUT    | `/expenses/{id}` |
| DELETE | `/expenses/{id}` |

### Analytics

| Method | Endpoint                |
| ------ | ----------------------- |
| GET    | `/dashboard`            |
| GET    | `/dashboard/categories` |

---

## Query Parameters

The `GET /expenses` endpoint supports filtering, pagination, and sorting.

### Pagination

```http
GET /expenses?page=1&limit=10
```

### Filter by Category

```http
GET /expenses?category=Food
```

### Sort by Amount

```http
GET /expenses?sort=amount_asc

GET /expenses?sort=amount_desc
```

### Sort by Date

```http
GET /expenses?sort=date_asc

GET /expenses?sort=date_desc
```

### Combined Example

```http
GET /expenses?category=Food&page=1&limit=5&sort=amount_desc
```

---

## Example API Response

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Dinner",
    "amount": 500,
    "category": "Food",
    "expense_date": "2026-07-07T00:00:00Z",
    "notes": "BBQ"
  }
}
```

---

## Live Demo

https://expense-tracker-api-g3c4.onrender.com

---

## Author

**Burhan Ashraf**

Backend Developer focused on Go, REST APIs, and backend systems.
