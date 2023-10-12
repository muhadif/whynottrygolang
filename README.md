# HTTP Middleware and Error Handling Example

## Overview

This project demonstrates the use of middleware in a Go web application. It includes a core middleware (```CoreMiddleWare```) for extracting user-related information from request headers and an HTTP handler (```GetHello```) for the "/hello" route that showcases error handling.

## Core Middleware (```CoreMiddleWare```)

The ```CoreMiddleWare``` function is a middleware that:

- Extracts user-related information (username and userID) from request headers.
- Incorporates the extracted information into the request context.
- Passes the modified request to the next handler in the chain.
- Logs information about the request after processing.

### Example Usage:

```go
func CoreMiddleWare(next http.Handler) http.Handler {
// Implementation details...
}
```
## HTTP Handler (```GetHello```)

The ```GetHello``` function is an HTTP handler for the "/hello" route that:

- Checks the query parameter ```isNeedError``` to determine error generation.
- Creates and returns a custom error if requested.
- Generates a JSON response with a greeting message otherwise.

### Example Usage:

```go
func GetHello(w http.ResponseWriter, req *http.Request) error {
// Implementation details...
}
```
## Usage

1. Clone the repository:

```bash
git clone git@github.com:muhadif/whynottrygolang.git && cd http-middleware-passing-context
```
2. Run the project:

```bash
go run main.go
```
3. hit endpoint
```bash
curl --request GET \
  --url 'http://localhost:8000/hello?isNeedError=true&tes=a' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --header 'userID: 2' \
  --header 'username: adifa'
```