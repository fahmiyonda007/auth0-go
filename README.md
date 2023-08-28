
# GO Starter

Auth0-go is a REST API project that uses auth0 as identity provider and controller book as an example


## Installation

```bash
  git clone https://github.com/fahmiyonda007/auth0-go.git
```

download package from go.mod
```bash
  go mod download
```

run 
```bash
  go run main.go
```
## API Reference

#### Login

```http
  POST /api/login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your username form auth0 user management |
| `password` | `string` | **Required**. Your password auth0 user management |

#### Get item

```http
  GET /api/books?page=1&length=10
  GET /api/books/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |
| `page`      | `int` | **Required** default is 1. page of items |
| `length`      | `int` | **Required** default is 10. size of items per page |

