# blog

This is a simple blog application built with Golang, Gin, GORM, and MySQL. The application allows users to sign up, log in, create blog posts, and like/dislike blog posts.

## Features

- User Signup and Login
- Create, Read, Update, and Delete (CRUD) Blog Posts
- Like and Dislike Blog Posts
- User authentication and authorization using JWT

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/armanfarokhi/blog.git
    cd blog
    ```

2. Create a `.env` file in the root directory with the following content:

    ```env
    DB_DIALECT=mysql
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    JWT_SECRET=your_jwt_secret
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:8080`.


