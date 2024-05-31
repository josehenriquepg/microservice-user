# User Management API

This is a RESTful API for user management, including authentication and authorization. The API is built with Go, uses PostgreSQL as the database, and JSON Web Tokens (JWT) for secure authentication.

## 📚 Features

- User registration
- User login with JWT token generation
- JWT-based authentication and authorization
- User profile retrieval (protected route)

## :gear: How to use

#### Prerequisites

- Go installed on your machine
- PostgreSQL installed and running

#### Steps

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/user-management-api.git
cd user-management-api
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Create a .env file in the root directory with the following content:**

```env
JWT_SECRET_KEY=your_generated_secret_key
DB_HOST=localhost
DB_USER=youruser
DB_NAME=yourdb
DB_PASSWORD=yourpassword
DB_PORT=5432
```

4. **Run PostgreSQL and create the database:**

```sql
CREATE DATABASE yourdb;
```

5. **Update the config/config.go file with your PostgreSQL credentials if needed:**

```go
dsn := "host=localhost user=youruser dbname=yourdb sslmode=disable password=yourpassword"
```

6. **Run the application:**

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

#### API Endpoints

**Register User**
  
  - URL: `/register`
  - Method: `POST`
  - Request Body:

    ```json
      {
        "username": "yourusername",
        "password": "yourpassword"
      }
    ```
  - Response:
    - `201 Created`

      ```json
        {
          "message": "User registered successfully"
        }
      ```

**Login User**
  
  - URL: `/login`
  - Method: `POST`
  - Request Body:

    ```json
      {
        "username": "yourusername",
        "password": "yourpassword"
      }
    ```
  - Response:
    
    - `200 OK`

      ```json
        {
          "token": "your.jwt.token"
        }
      ```

    - `401 Unauthorized`

      ```json

        {
          "error": "Invalid username or password"
        }
      ```
#### Get User Profile (Protected Route)
  - URL: `/auth/profile`
  - Method: `GET`
  - Headers:
    - `Authorization: Bearer your.jwt.token`
  - Response:
    - `200 OK`
      ```json
        {
          "message": "Authenticated",
          "username": "yourusername"
        }
      ```
    - `401 Unauthorized`
      ```json
        {
          "error": "Unauthorized"
        }
      ```

## 💼 License

This project is licensed under the MIT License.

## 💻 Technologies used
[![techs](https://skillicons.dev/icons?i=go,postgres,git&theme=dark)](https://skillicons.dev)
- Gin (web framework)
- GORM (ORM)
- JWT (JSON Web Tokens)
- godotenv (environment variables)

## 🤝 Collaborators 
<table>
  <tr>
    <td align="center">
      <a href="http://github.com/josehenriquepg">
        <img src="https://avatars.githubusercontent.com/josehenriquepg" width="100px;" alt="Foto de José Henrique no GitHub"/><br>
        <sub>
          <b>José Henrique</b>
        </sub>
      </a>
    </td>
  </tr>
</table>