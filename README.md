# Golang JWT Project

![Golang JWT](https://miro.medium.com/v2/resize:fit:1000/1*Z98o3l6mg8WGBiPuFd2t2A.jpeg)

This project demonstrates how to build a secure authentication system in a Golang application. It includes examples of login, registration, and forgot password functionalities, as well as implementing access and refresh tokens. The project uses MongoDB for data storage, Redis for caching and session management, and follows clean architecture principles.

## 🚀 Features

- **Secure authentication system**
- **Login and registration endpoints**
- **Forgot password functionality**
- **Access and refresh tokens**
- **Middleware for JWT authentication**
- **Example endpoints for protected routes**
- **MongoDB for data storage**
- **Redis for caching and session management**
- **Clean architecture**

## 📋 Prerequisites

- **Go 1.16 or higher**
- **A basic understanding of JWTs**
- **MongoDB**
- **Redis**

## 📦 Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/rohankarn35/golang-jwt.git
    cd golang-jwt
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Create a `.env` file with the following fields:**
    ```env
    REDIS_ADDR=<redis address>
    REDIS_PASSWORD=           # Leave empty if Redis has no password
    REDIS_DB=0                # Redis database number (default is 0)

    MONGODB_URI=mongodb+srv://<name>:<password>@cluster0.a4ecxsw.mongodb.net/?retryWrites=true&w=majority&appName=<clustername>

    JWTSECRET= your-secret-key
    ```

## 🚀 Usage

1. **Run the application:**
    ```sh
    go run main.go
    ```

2. **Use a tool like Postman to test the endpoints.**

## 📚 Endpoints

- `POST /login`: Authenticate a user and receive a JWT.
- `POST /register`: Register a new user.
- `POST /forgot-password`: Initiate password reset process.
- `POST /auth/logout`: Logout a user.
- `POST /auth/reset-password`: Reset a user's password.
- `POST /auth/refresh`: Generate a new refresh token.

## 💡 Example

### Generating a JWT

```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user_id": "user_id",
    "exp":      time.Now().Add(time.Hour * 72).Unix(),
})

tokenString, err := token.SignedString([]byte("your-256-bit-secret"))
if err != nil {
    log.Fatal(err)
}

fmt.Println(tokenString)
```

### Verifying a JWT

```go
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("your-256-bit-secret"), nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["user_id"])
} else {
    fmt.Println(err)
}
```

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## 🙏 Acknowledgements

- [jwt-go](https://github.com/dgrijalva/jwt-go) library for JWT implementation in Go.
- [Redis](https://redis.io/) for caching and session management.
- [Gin](https://github.com/gin-gonic/gin) for the HTTP web framework.
- [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.

