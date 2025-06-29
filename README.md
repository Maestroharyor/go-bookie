# Go Bookie

A simple REST API for book management built with Go and the Fiber web framework. This project demonstrates basic CRUD operations with in-memory data storage.

## Features

- ✅ Get all books
- ✅ Get a single book by ID
- ✅ Create a new book
- ✅ Update an existing book
- ✅ Delete a book
- ✅ JSON response format with consistent structure
- ✅ Error handling and validation
- ✅ Hot reload support with Air

## Tech Stack

- **Language**: Go 1.24.4
- **Web Framework**: Fiber v2
- **Data Storage**: In-memory (slice)
- **Development Tools**: Air (hot reloading)

## Project Structure

```
go-bookie/
├── data/
│   └── data.go          # Book data model and sample data
├── handlers/
│   └── handlers.go      # HTTP request handlers
├── response/
│   └── response.go      # Response utility functions
├── .air.toml           # Air configuration for hot reloading
├── main.go             # Application entry point
├── go.mod              # Go module dependencies
└── go.sum              # Dependency checksums
```

## API Endpoints

| Method | Endpoint   | Description     | Request Body |
| ------ | ---------- | --------------- | ------------ |
| GET    | `/`        | Get all books   | None         |
| GET    | `/:bookId` | Get single book | None         |
| POST   | `/`        | Create new book | Book JSON    |
| PUT    | `/:bookId` | Update book     | Book JSON    |
| DELETE | `/:bookId` | Delete book     | Book JSON    |

## Book Model

```json
{
  "id": 1,
  "name": "Book Title",
  "author": "Author Name",
  "quantity": 10
}
```

## Response Format

### Success Response

```json
{
  "success": true,
  "message": "Success message",
  "data": {
    /* response data */
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "data": null
}
```

## Getting Started

### Prerequisites

- Go 1.24.4 or higher
- Air (for hot reloading, optional)

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd go-bookie
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Run the application**

   ```bash
   go run main.go
   ```

   Or with hot reloading using Air:

   ```bash
   air
   ```

4. **The server will start on port 3000**
   ```
   http://localhost:3000
   ```

### Development with Air

This project includes Air configuration for hot reloading during development. Air will automatically rebuild and restart the server when you make changes to `.go` files.

Install Air globally:

```bash
go install github.com/cosmtrek/air@latest
```

Then run:

```bash
air
```

## API Usage Examples

### Get All Books

```bash
curl -X GET http://localhost:3000/
```

### Get Single Book

```bash
curl -X GET http://localhost:3000/1
```

### Create New Book

```bash
curl -X POST http://localhost:3000/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New Book",
    "author": "Author Name",
    "quantity": 5
  }'
```

### Update Book

```bash
curl -X PUT http://localhost:3000/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Book",
    "author": "Updated Author",
    "quantity": 15
  }'
```

### Delete Book

```bash
curl -X DELETE http://localhost:3000/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Book to Delete",
    "author": "Author Name"
  }'
```

## Sample Data

The application comes with sample books:

1. **The Great Gatsby** by F. Scott Fitzgerald (Quantity: 10)
2. **To Kill a Mockingbird** by Harper Lee (Quantity: 5)
3. **1984** by George Orwell (Quantity: 8)

## Error Handling

The API includes comprehensive error handling for:

- Invalid book IDs
- Missing required fields (name, author)
- Book not found scenarios
- Invalid JSON data

## Middleware

- **Logger**: Logs HTTP requests
- **Recover**: Recovers from panics
- **404 Handler**: Custom not found response

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).
