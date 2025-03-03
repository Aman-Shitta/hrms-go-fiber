# HRMS API with Fiber and MongoDB (Dockerized)

This project is a simple **Human Resource Management System (HRMS) API** built using **Fiber (Go Web Framework)** and **MongoDB** as the database. The MongoDB instance is managed using **Docker**.

## Features
- Employee CRUD operations (Create, Read, Update, Delete)
- MongoDB integration using the `mongo-driver`
- Structured Fiber routing
- Docker support for MongoDB

## Prerequisites
Make sure you have the following installed:
- [Go](https://go.dev/dl/)
- [Docker](https://www.docker.com/get-started)
- [MongoDB (optional, if not using Docker)](https://www.mongodb.com/try/download/community)

## Getting Started

### 1. Clone the repository
```sh
git clone https://github.com/Aman-Shitta/hrms-go-fiber.git
cd hrms-go-fiber
```

### 2. Run MongoDB in Docker
If you don't have a local MongoDB instance, you can use Docker:
```sh
docker run --name mongo-hrms -d -p 27017:27017 mongo:latest
```
This will start a MongoDB container on port `27017`.

### 3. Install dependencies
```sh
go mod tidy
```

### 4. Run the API
```sh
go run main.go
```

## API Endpoints
| Method | Endpoint | Description |
|--------|------------|-------------|
| GET    | `/api/v1/employee` | Get all employees |
| POST   | `/api/v1/employees` | Create a new employee |
| GET    | `/api/v1/employee/:id` | Get an employee by ID |
| PUT    | `/api/v1/employee/:id` | Update an employee by ID |
| DELETE | `/api/v1/employee/:id` | Delete an employee by ID |

## Project Structure
```
hrms-go-fiber/
├── database/
│   ├── database.go  # MongoDB connection setup
│
├── employee/
│   ├── employee.go  # Employee model & CRUD handlers
│
├── main.go          # Application entry point
├── go.mod           # Module definition
├── go.sum           # Dependency checksum
└── README.md        # Project documentation
```

## Example Usage

### Create an Employee
```sh
curl -X POST "http://localhost:3000/api/v1/employees" \
     -H "Content-Type: application/json" \
     -d '{"name": "John Doe", "salary": 50000, "age": 30}'
```

### Get All Employees
```sh
curl -X GET "http://localhost:3000/api/v1/employee"
```

### Update an Employee
```sh
curl -X PUT "http://localhost:3000/api/v1/employee/{id}" \
     -H "Content-Type: application/json" \
     -d '{"name": "Jane Doe", "salary": 60000, "age": 32}'
```

### Delete an Employee
```sh
curl -X DELETE "http://localhost:3000/api/v1/employee/{id}"
```

## Stopping and Removing MongoDB Container
If you used Docker for MongoDB, you can stop and remove the container with:
```sh
docker stop mongo-hrms

docker rm mongo-hrms
```

## License
This project is open-source and available under the [MIT License](LICENSE).

---

