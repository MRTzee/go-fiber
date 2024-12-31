# 🚀 GO API FIBER GORM

A robust RESTful API built with Fiber framework, GORM, and PostgreSQL, featuring JWT authentication and comprehensive file upload capabilities (both single and multiple files).

## ✨ Features

- **JWT Authentication** - Secure endpoint protection
- **File Management** - Support for single and multiple file uploads
- **Input Validation** - Comprehensive request validation
- **PostgreSQL Database** - Reliable data persistence
- **RESTful Architecture** - Clean and standardized API design

## 🛠 Technology Stack

- **[Fiber](https://gofiber.io/)** - Fast HTTP web framework
- **[GORM](https://gorm.io/)** - Powerful ORM for Golang
- **[PostgreSQL](https://www.postgresql.org/)** - Advanced open-source database
- **[Viper](https://github.com/spf13/viper)** - Complete configuration solution
- **[Docker](https://www.docker.com/)** - Containerization platform

## 🚀 Quick Start

1. **Prerequisites**
   - Docker and Docker Compose installed
   - Git installed

2. **Installation**
   ```bash
   # Clone the repository
   git clone [your-repo-url]

   # Navigate to project directory
   cd [project-directory]

   # Start the application
   docker-compose up --build
   ```

## 📌 API Endpoints

### Authentication
- **Create User** `POST /users`
  ```json
  {
    "name": "muksal",
    "email": "muksal@mail.com",
    "password": "123456",
    "address": "aceh",
    "phone": "+6285322420875"
  }
  ```

- **Login** `POST /login`
  ```json
  {
    "email": "muksal@mail.com",
    "password": "123456"
  }
  ```

### User Management
- **Get All Users** `GET /users`
- **Get User by ID** `GET /users/:id`
- **Update User** `PUT /users/:id`
- **Update Email** `PUT /users/:id/update-email`
- **Delete User** `DELETE /users/:id`

### Book Management
- **Get All Books** `GET /books`
- **Create Book** `POST /books` (Multipart Form)
  - `title`: Book title
  - `author`: Author name
  - `cover`: Image file

### Category & Gallery Management
- **Get Categories** `GET /category`
- **Create Category** `POST /category`
- **Get Gallery** `GET /gallery`
- **Upload Photos** `POST /gallery` (Multipart Form)
  - `photos`: Multiple image files (JPG/PNG)
  - `category_id`: Category reference
- **Delete Photo** `DELETE /gallery/:id`

## 🔒 Authentication

All protected endpoints require a JWT token in the `x-token` header. The token is obtained upon successful login and is valid for 2 minutes.

## 📝 Notes

- Supported image formats for gallery: JPG, PNG
- One category can have multiple photos (One-to-Many relationship)
- All endpoints except login and user creation require authentication

## 🤝 Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.

## 📄 License

[MRTzee]

---
⭐️ If you find this project helpful, please give it a star!
