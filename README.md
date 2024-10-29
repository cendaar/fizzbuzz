# Fizzbuzz REST Server

This is a simple Fizzbuzz REST server implemented in Golang using the Gin framework.

It provides an API endpoint for generating Fizzbuzz sequences and retrieving statistics about the most common requests.

## Features

- Accepts parameters for generating Fizzbuzz sequences.
- Exposes an API endpoint for generating Fizzbuzz sequences.
- Provides statistics about the most common requests.

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/cendaar/fizzbuzz.git
cd fizzbuzz
```

### 2. Build the fizzbuzz server:

```bash
make run
```

### 3. Run unit tests:

```bash
make test
```

### 4. Test endpoint fizzbuzz:

```bash
$ curl -X GET "http://localhost:8080/fizzbuzz" -d '{"int1": 2, "int2": 3, "limit": 10, "str1": "Hello", "str2": "World"}' | jq
{
  "result": "1,Hello,World,Hello,5,HelloWorld,7,Hello,World,Hello"
}
```

### 5. Test stats endpoint

```bash
$ curl -X GET "http://localhost:8080/stats" | jq
{
  "result": {
    "hits": 2,
    "most_common_request": {
      "int1": 2,
      "int2": 3,
      "limit": 10,
      "str1": "Hello",
      "str2": "World"
    }
  }
}
```

##  Personal remarks

It took me 6 hours to develop this project I don't consider it finished. Here is a list of things I would have enhanced/added:

- Implement persitent statistics data (Redis or PostgreSQL)
-	Security: Implement HTTPS, CORS config, input validation, and rate limiting.
-	Performance and Scalability: Add caching, enable concurrency, and ensure scalability in a containerized environment.
- Dockerization: Ensure the app is containerized for easy deployment, scaling, and testing.
-	Monitoring and Alerting: Set up Prometheus for metrics, add alerting, and implement distributed tracing.
-	Error Handling and Logging: Use structured logging and centralized error handling.
-	Documentation and Testing: Use Swagger, improve tests, and document API usage.
-	Automated CI/CD: Automate building, testing, and deployment in a CI/CD pipeline.
-	Health Checks and Graceful Shutdown: Add health checks and ensure the server shuts down gracefully.
- Caching: Cache the most frequent requests (using Redis) to avoid recomputing the FizzBuzz sequence on every request.