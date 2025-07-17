# Task Management REST API Documentation

## Base URL
`http://localhost:8080/tasks`

## Endpoints

### 1. Get All Tasks
**Endpoint**: `GET /tasks`

**Description**: Retrieves all tasks from the system

**Request**:
```http
GET /tasks HTTP/1.1
Host: localhost:8080
```

**Response**:
- Status: `200 OK`
- Body:
```json
[
    {
        "id": 1,
        "title": "Implement user authentication",
        "description": "Create login and registration endpoints",
        "due_date": "2025-07-18T18:00:00Z",
        "status": "pending",
    }
]
```

### 2. Get Single Task
**Endpoint**: `GET /tasks/:id`

**Description**: Retrieves a specific task by ID

**Path Parameters**:
- `id` (required): Task ID (integer)

**Request**:
```http
GET /tasks/1 HTTP/1.1
Host: localhost:8080
```

**Response**:
- Success: `200 OK`
```json
{
    "id": 1,
    "title": "Implement user authentication",
    "description": "Create login and registration endpoints with JWT support",
    "due_date": "2025-07-18T18:00:00Z",
    "status": "pending"
}
```
- Not Found: `404 Not Found`
```json
{
    "error": "no task found with this id to see"
}
```

### 3. Create Task
**Endpoint**: `POST /tasks`

**Description**: Creates a new task

**Request**:
```http
POST /tasks HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
  "title": "Implement user authentication",
  "description": "Create login and registration endpoints with JWT support",
  "due_date": "2025-07-18T18:00:00Z",
  "status": "pending"
}
```

**Validation Rules**:
- `due_date`: ISO 8601 format
- `status`: must be `pending|in_progress|completed`

**Response**:
- Success: `201 Created`
```json
{
    "id": 1,
    "title": "Implement user authentication",
    "description": "Create login and registration endpoints with JWT support",
    "due_date": "2025-07-18T18:00:00Z",
    "status": "pending"
}
```
- Error: `400 Bad Request`
```json
{
    "error": "task title can not be empty"
}
```

### 4. Update Task
**Endpoint**: `PUT /tasks/:id`

**Description**: Updates an existing task (full or partial update)

**Path Parameters**:
- `id` (required): Task ID (integer)

**Request**:
```http
PUT /tasks/1 HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "status": "in_progress",
    "description": "Updated description"
}
```

**Response**:
- Success: `200 OK`
```json
{
    "message": "task updated successfully",
    "updated task": {
        "id": 1,
        "title": "Implement user authentication",
        "description": "Create login and registration endpoints with JWT support",
        "due_date": "2025-07-18T18:00:00Z",
        "status": "in_progress"
    }
}
```
- Not Found: `404 Not Found`
```json
{
    "error": "no task found with this id to update"
}
```

### 5. Delete Task
**Endpoint**: `DELETE /tasks/:id`

**Description**: Deletes a task by ID

**Path Parameters**:
- `id` (required): Task ID (integer)

**Request**:
```http
DELETE /tasks/1 HTTP/1.1
Host: localhost:8080
```

**Response**:
- Success: `200 Ok` (empty body)
```json
{
    "message": "task deleted successfully"
}
```
- Not Found: `404 Not Found`
```json
{
    "error": "no task found with this id to delete"
}
```

## Status Codes
| Code | Description |
|------|-------------|
| 200 | OK - Successful request, deletion |
| 201 | Created - Resource created |
| 400 | Bad Request - Invalid input |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error |

## Task Status Values
- `pending` 
- `in_progress`
- `completed`

## Date Format
All dates must be in ISO 8601 format:  
`YYYY-MM-DDTHH:MM:SSZ`  
Example: `"2025-07-14T18:00:00Z"`

## Example Requests using curl

### Get All Tasks
```bash
curl -X GET http://localhost:8080/tasks
```

### Get Single Task
```bash
curl -X GET http://localhost:8080/tasks/1
```

### Create Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Implement user authentication",
    "description": "Create login and registration endpoints with JWT support",
    "due_date": "2025-07-18T18:00:00Z",
    "status": "pending"
}'
```

### Update Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status": "in_progress"}'
```

### Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```
