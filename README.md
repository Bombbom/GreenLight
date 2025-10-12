# README 

## Setup
- Prerequisites: 
    - Go
    - curl 
    - hey tool
    - git
    - browser


```
go install github.com/rakyll/hey@latest
go get github.com/julienschmidt/httprouter@v1
```

## Greenlight API

| Method | URL Pattern | Action |
|--------|-------------|--------|
| GET | /v1/healthcheck | Show application health and version information |
| GET | /v1/movies | Show the details of all movies |
| POST | /v1/movies | Create a new movie |
| GET | /v1/movies/:id | Show the details of a specific movie |
| PATCH | /v1/movies/:id | Update the details of a specific movie |
| DELETE | /v1/movies/:id | Delete a specific movie |
| POST | /v1/users | Register a new user |
| PUT | /v1/users/activated | Activate a specific user |
| PUT | /v1/users/password | Update the password for a specific user |
| POST | /v1/tokens/authentication | Generate a new authentication token |
| POST | /v1/tokens/password-reset | Generate a new password-reset token |
| GET | /debug/vars | Display application metrics |

## Project structure 

```
.
├── Makefile
├── README.md
├── bin
├── cmd
│   └── api
│       ├── healthcheck.go
│       └── main.go
├── go.mod
├── internal
├── migrations
└── remote
```

- The bin directory will contain our compiled applicatio binaries, ready for deployment to a production server.
- The cmd/api
- The internal 
- The migrations 
- The remote 
- The go.mod 
- The Makefile 

## Check
- Endpoint: /v1/healthcheck

```bash 
go run ./cmd/api/
```


```bash 
$ curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Sun, 12 Oct 2025 13:55:02 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

status: available
environment: development
version: 1.0.0
```

