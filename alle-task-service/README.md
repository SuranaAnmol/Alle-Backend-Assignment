# Task Management API

This API provides CRUD operations for managing tasks using Go, Gin, and MongoDB. It supports task creation, retrieval, updating, deletion, pagination, and filtering.

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/SuranaAnmol/Alle-Backend-Assignment-Go.git
   ```
2. **Navigate to the project directory:**
   ```sh
   cd Alle-Backend-Assignment-Go
   ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```
4. **Ensure MongoDB is running.**
5. **Run the application:**
   ```sh
   go run main.go
   ```

## API Endpoints

### Create a Task
**POST** `/tasks`

**Request Body:**
```json
{
  "title": "Sample Task",
  "description": "This is a sample task description",
  "status": "pending"
}
```
**Response:**
```json
{
  "message": "Task created",
  "id": "<task_id>"
}
```

### Get All Tasks (Pagination & Filtering)
**GET** `/tasks?page=1&limit=10&status=pending`

**Response:**
```json
{
  "page": 1,
  "limit": 10,
  "tasks": [
    {
      "id": "<task_id>",
      "title": "Sample Task",
      "description": "This is a sample task description",
      "status": "pending"
    }
  ]
}
```

### Get a Task by ID
**GET** `/tasks/{id}`

**Response:**
```json
{
  "id": "<task_id>",
  "title": "Sample Task",
  "description": "This is a sample task description",
  "status": "pending"
}
```

### Update a Task
**PUT** `/tasks/{id}`

**Request Body:**
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "status": "completed"
}
```
**Response:**
```json
{
  "message": "Task updated"
}
```

### Delete a Task
**DELETE** `/tasks/{id}`

**Response:**
```json
{
  "message": "Task deleted"
}
```

## Notes
- Uses MongoDB as the database.
- Tasks are stored in the `tasks` collection.
- Supports filtering by `status`.
- Pagination parameters: `page` and `limit`.

For issues, raise a GitHub issue or reach out.

