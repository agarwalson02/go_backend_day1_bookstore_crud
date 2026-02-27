# Bookstore Management System 📚

A professional RESTful API for managing a bookstore, built with **Go (Golang)**, **Gorilla Mux**, **GORM**, and **MySQL**.

## 🚀 Features

- **Full CRUD Operations**: Create, Read, Update, and Delete books.
- **RESTful API Design**: Clean and structured endpoints.
- **ORM Integration**: Powered by GORM for seamless database interactions.
- **Auto-Migrate**: Automatically creates tables in MySQL based on Go structs.

## 🛠️ Tech Stack

- **Language**: Go 1.2x
- **Router**: `github.com/gorilla/mux`
- **ORM**: `github.com/jinzhu/gorm`
- **Database**: MySQL

## 📋 Prerequisites

Before running this project, ensure you have:
- [Go](https://golang.org/doc/install) installed.
- [MySQL](https://www.mysql.com/downloads/) server running.
- A database named `simplerest` created.

## ⚙️ Configuration

The database connection is configured in `pkg/config/app.go`. Ensure your MySQL credentials match or update the following line:

```go
d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
```

## 🚦 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| **POST** | `/book/` | Create a new book |
| **GET** | `/book/` | Retrieve all books |
| **GET** | `/book/{bookId}` | Retrieve a specific book by ID |
| **PUT** | `/book/{bookId}` | Update book details |
| **DELETE** | `/book/{bookId}` | Remove a book from the store |

## 📦 Data Structure

Each book record contains:
- `name`: String
- `author`: String
- `publication`: String

## 🏃 How to Run

1.  **Clone the project**:
    ```bash
    git clone https://github.com/agarwalson02/go-bookstore
    ```
2.  **Navigate to the project directory**:
    ```bash
    cd 3_crud_msql
    ```
3.  **Install dependencies**:
    ```bash
    go mod download
    ```
4.  **Run the application**:
    ```bash
    go run cmd/main/main.go
    ```

The server will start on `localhost:9000` (or the configured port).

---
Developed with ❤️ by [Agarwalson02](https://github.com/agarwalson02)
