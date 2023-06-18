# Service Auth

Service-Auth is a powerful authentication service that offers a comprehensive set of endpoints for user registry creation, JSON Web Token (JWT) generation, and JWT validation. This service ensures secure and reliable authentication for your applications.

We take pride in implementing the best software development practices, including the Clean Architecture principles coined by the renowned software engineer, Robert C. Martin (also known as Uncle Bob). Our architecture ensures that our codebase is organized, testable, and maintainable, allowing us to continuously provide a high-quality service to our clients.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [MySQL](https://www.mysql.com/downloads/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/efishery-technical-test.git
$ cd efishery-technical-test/service-auth
```

- Install Dependencies

```
$ go mod tidy
```

### Running

```
$ go run cmd/main.go
```

or simply running this command

```
$ make run
```

### Running on Docker

```
$ docker compose up db
$ docker compose up service-auth

```

### Migration Database

```
$ Automatically Migrated The Database When Service is Up

```

### Features

This service has the following API endpoints:

- `` `POST /register` ``: endpoint for registering user
- `` `POST /login` ``: endpoint for user login
- `` `GET /validate/jwt` ``: endpoint for validate the jwt token

We can test the endpoint using the postman collection in `efishery-technical-test/service-auth/tools`.

### Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-ralali-technical-test/blob/master/LICENSE) file for details.
