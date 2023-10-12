# HTTP Middleware Passing Context

This is a simple example project that demonstrates how to use middleware in a Go web application to pass context between middleware functions.

## Overview

The project showcases the use of middleware in a web application built with Go. It uses the [gorilla/mux](https://github.com/gorilla/mux) router for routing and a basic custom middleware to pass context between handlers.

## Prerequisites

Before running this project, make sure you have the following installed:

- Go (at least Go 1.21)
- [gorilla/mux](https://github.com/gorilla/mux)

## Installation

Clone the repository:

```bash
git clone https://github.com/muhadif/whynottrygolang.git
cd whynottrygolang/http-middleware-passing-context
````

# Test
```
curl --request GET \
  --url 'http://localhost:8000/hello?isNeedError=true&tes=a' \
  --header 'User-Agent: insomnia/2023.5.8' \
  --header 'userID: 2' \
  --header 'username: adifa'
  ```

