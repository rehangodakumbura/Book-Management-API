# 📘 Book Management API (GoLang Intern Test)

A simple RESTful API built with [Go](https://golang.org/), [Fiber](https://gofiber.io/), and [GORM](https://gorm.io/) with SQLite for managing a collection of books.

---

## 🚀 Features

- ✅ Create, Read, Update, Delete (CRUD) operations on books
- ✅ Well-structured project layout with separation of concerns
- ✅ Uses SQLite for lightweight database storage
- ✅ Unit-test ready structure
- ✅ Proper error handling and HTTP status codes

---

## 🛠 Tech Stack

- [Go 1.18+](https://go.dev/dl/)
- [Fiber v2](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/index.html)

---

## 📁 Project Structure

```

.
├── main.go
├── go.mod
├── go.sum
├── README.md
├── models
│   └── book.go
├── handlers
│   └── book\_handler.go
├── services
│   └── book\_service.go
├── database
│   └── database.go

````

## 📦 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/rehangodakumbura/Book-Management-API
cd GoInternTest
```

### 2. Initialize Go Module (if needed)

```bash
go mod init book-api
```

### 3. Install Dependencies

```bash
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### 4. Run the App

```bash
# On Windows (PowerShell)
$env:CGO_ENABLED="1"
go run main.go
```

Visit: [http://localhost:3000/books](http://localhost:3000/books)

---

## 📡 API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| POST   | `/books`     | Create a new book |
| GET    | `/books`     | Get all books     |
| GET    | `/books/:id` | Get a book by ID  |
| PUT    | `/books/:id` | Update a book     |
| DELETE | `/books/:id` | Delete a book     |

### 📥 Sample Request Body (POST/PUT)

```json
{
  "title": "1984",
  "author": "George Orwell",
  "year": 1949
}
```

