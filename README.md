# GraphQL Implementation in Golang with MySQL
This is example of using GraphQL for API in Go with MySQL. This project is for my learning purpose.

# Dependencies

| **Tech**                                                               | **Description**                                                                       |
| ---------------------------------------------------------------------- | --------------------------------------------------------------------------------------|
| [graphql-go](https://github.com/graphql-go/graphql)                    | `graphql-go` An implementation of GraphQL for Go / Golang.                            |
| [graphql-go-handler](https://github.com/graphql-go/graphql-go-handler) | Middleware to handle GraphQL queries through HTTP requests.                           |
| [godotenv](https://github.com/joho/godotenv)                           | A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).|
| [go-mysql-driver](https://github.com/go-sql-driver/mysql)              | A MySQL-Driver for Go's database/sql package.                                         |

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
Query to get all users
```graphql
query {
  users {
    id
    name
    email
    created_at
    updated_at
  }
}
```
Query to get user by id
```graphql
query {
  user(id: 1) {
    id
    name
    email
    created_at
    updated_at
  }
}
```
Query to update user
```graphql
mutation {
  updateUser(id: 1, name: "Kharisma Januar", email: "kharisma.januar@gmail.com") {
    id
    name
    email
    updated_at
  }
}
```
Query to delete user (soft delete)
```graphql
mutation {
  deleteUser(id: 1) {
    id
  }
}
```
Query to create new task
```graphql
mutation {
  createTask(name: "Study GraphQL Go", description: "Implement GraphQL in Go", user_id:1) {
    id
    name
    created_at
    user{
        id
    }
  }
}
```
Query to get all tasks
```graphql
query{
    tasks{
        id
        name
        description
        is_complete
        created_at
        updated_at
        user{
            id
            name
            email
        }
    }
}
```
Query to get task by id
```graphql
query {
  task(id: 1) {
    id
    name
    description
    is_complete
    user{
        id
        email
        name
    }
    created_at
    updated_at
  }
}
```
Query to update task
```graphql
mutation {
  updateTask(id: 1, name: "Learn Git", description: "Learn Git for beginners", user_id:1, is_complete: true) {
    id
    name
    description
    user{
        id
        name
    }
    is_complete
    updated_at
  }
}
```
Query to delete task (soft delete)
```graphql
mutation{
    deleteTask(id:1) {
        id
    }
}
```
