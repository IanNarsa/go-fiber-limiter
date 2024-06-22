# Romeo-Lima API

Romeo-Lima API is a boilerplate RESTful API built using Go and the Fiber framework.

## Features

- **Customer Management:** Perform CRUD operations on customer data.
- **Rate Limiting:** Implements IP-based rate limiting for enhanced security.
- **PostgreSQL Support:** Utilizes PostgreSQL as the database backend.

## Folder Structure
```bash
Romeo-Lima:.
│   .env
│   .gitignore
│   go.mod
│   go.sum
│   Readme.md
│
├───cmd
│       main.go
│
├───config
│       config.go
│
├───internal
│   ├───customers
│   │   ├───delivery
│   │   │       handler.go
│   │   │
│   │   ├───repository
│   │   │       customer_repository.go
│   │   │       repository.go
│   │   │
│   │   └───usecase
│   │           customer_usecase.go
│   │
│   └───models
│           customer.go
│
└───pkg
    └───logger
            logger.go

``` 

## Getting Started

### Prerequisites

- Go 1.17 or later
- PostgreSQL 13 or later
- Fiber 2.x

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/romeo-lima.git
   ```

2. Create a `.env` file with your database credentials and other configuration settings.

3. Start the API server:
   ```bash
   go run main.go
   ```

### API Endpoints

- **GET /customers:** Retrieve a list of all customers.
- **GET /customers/:id:** Retrieve a single customer by ID.
- **POST /customers:** Create a new customer.

### Configuration

The API uses environment variables for configuration:

- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name

You can set these variables in a `.env` file or through system environment variables.

### Rate Limiting Middleware

The API uses rate limiting middleware to protect against abuse:

```go
app.Use(limiter.New(limiter.Config{
    Max:        1,
    Expiration: 30 * time.Second,
    KeyGenerator: func(c *fiber.Ctx) string {
        return c.IP()
    },
    LimitReached: func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
            "error": "Too Many Requests",
            "message": "You have exceeded the request limit. Please try again later.",
        })
    },
}))
```

- **Max:** Maximum number of requests allowed within the expiration time frame (1 request).
- **Expiration:** Time frame after which the request count resets (30 seconds).
- **KeyGenerator:** Function to generate a unique key based on the client's IP address.
- **LimitReached:** Function called when the rate limit is exceeded, returning a 429 error with a JSON response.

### Sample results
**Success** </br>
![success-sample](/assets/success.png)


**Failed Over Limit**
![over-limit-sample](/assets/failed.png)
