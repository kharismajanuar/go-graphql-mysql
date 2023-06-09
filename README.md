# GraphQL Implementation in Golang with MySQL
Using GraphQL for API in Go with MySQL. This project is for learning purpose.

# Dependencies

| **Tech**                                                               | **Description**                                                                                                              |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| [graphql-go](https://github.com/graphql-go/graphql)                    | `graphql-go` An implementation of GraphQL for Go / Golang.                                                                   |
| [graphql-go-handler](https://github.com/graphql-go/graphql-go-handler) | Middleware to handle GraphQL queries through HTTP requests.                                                                  |
| [GoDotEnv](https://github.com/joho/godotenv)                           | A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).                                       |
| [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)              | A MySQL-Driver for Go's database/sql package                                                                                 |

# Setup

1. Install Go
2. Install MySQL
3. Create database
4. Create local.env
```
export DB_CONN="<username>:<password>@tcp(<hostname>:<port>)/<db_name>?parseTime=true"
```

# Sample Queries
 Query to create new user
```graphql
mutation {
  createUser(name: "Kharisma Januar", email: "kharisma.januar@gmail.com") {
    id
    name
    email
    created_at
  }
}
```
