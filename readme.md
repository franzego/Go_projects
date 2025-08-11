# Server Monitor

Server Monitor API
A lightweight Go-based server monitoring API that provides health checks and performance metrics for applications. Built with scalability and observability in mind.

## Features
Health Check Endpoint – Quickly verify that the service is up and running.

Logging Middleware – Tracks request method, path, status code, and response time.

Structured Routing – Clean route handling using net/http.

Production-Ready Structure – Organized for maintainability and easy scaling.

(Optional) Prometheus Integration – Ready for monitoring and alerting.

## Project Structure
.
├── api.go              # Entry point of the API
├── middleware/         # Logging and other middleware
├── handlers/           # Request handlers
├── go.mod              # Go modules file
├── go.sum
└── .gitignore

##Installation
1. Clone the repository
   git clone https://github.com/<your-username>/server-monitor.git
cd server-monitor
2. Install dependencies
   go mod tidy
3. Run the application
   go run api.go
   
## Middleware
This project includes a logging middleware that:

Wraps the HTTP ResponseWriter to capture status codes

Logs method, path, response code, and duration for each request

## Future Enhancements
Integrate Prometheus metrics

Add authentication middleware

Implement additional system monitoring endpoints


