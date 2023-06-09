package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

// Table Creation
const TABLE_USER = `CREATE TABLE IF NOT EXISTS users
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    email varchar(150) NOT NULL,
    created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	deleted_at datetime
);`

const TABLE_TASK = `CREATE TABLE IF NOT EXISTS tasks
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    description text NOT NULL,
    user_id int NOT NULL,
	is_complete BOOLEAN NOT NULL DEFAULT 0,
    created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	deleted_at datetime,
	constraint fk_users_tasks FOREIGN KEY (user_id) REFERENCES users(id)
);`

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      int       `json:"user_id"`
	IsComplete  bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func DBConn() (*sql.DB, error) {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("Failed to load env file. Error: %v \n", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DB_CONN"))
	if err != nil {
		log.Printf("Failed open connection to MySQL. Error: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Connection failed to MySQL. Error: %v \n", err)
	}

	sliceTable := []string{TABLE_USER, TABLE_TASK}

	for _, query := range sliceTable {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Failed table creation. Error: %v \n", err)
		}
	}

	log.Println("Connection to MySQL success")

	return db, nil
}

func main() {
	db, err := DBConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create new object type
	userType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.Name, nil
					}

					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The email address of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.Email, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The created_at date of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.CreatedAt, nil
					}

					return nil, nil
				},
			},
			"updated_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The updated_at date of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.UpdatedAt, nil
					}

					return nil, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The deleted_at date of the user.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok {
						return user.DeletedAt, nil
					}

					return nil, nil
				},
			},
		},
	})

	taskType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Task",
		Description: "A Task",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.ID, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if post, ok := p.Source.(*Task); ok {
						return post.Name, nil
					}

					return nil, nil
				},
			},
			"description": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The description of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.Description, nil
					}

					return nil, nil
				},
			},
			"is_complete": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Boolean),
				Description: "The status of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.IsComplete, nil
					}

					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The created_at date of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.CreatedAt, nil
					}

					return nil, nil
				},
			},
			"updated_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The updated_at date of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.UpdatedAt, nil
					}

					return nil, nil
				},
			},
			"deleted_at": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.DateTime),
				Description: "The deleted_at date of the task.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						return task.CreatedAt, nil
					}

					return nil, nil
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if task, ok := p.Source.(*Task); ok {
						user := &User{}
						err = db.QueryRow("select id, name, email, created_at, updated_at from users where id = ? AND deleted_at IS NULL;", task.UserID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
						if err != nil {
							log.Printf("Failed query select user by id. Error: %v \n", err)
						}

						return user, nil
					}

					return nil, nil
				},
			},
		},
	})

	// Create GraphQL schema
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get a user.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					user := &User{}
					err = db.QueryRow("select id, name, email, created_at, updated_at from users where id = ? AND deleted_at IS NULL;", id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
					if err != nil {
						log.Printf("Failed query select user by id. Error: %v \n", err)
					}

					return user, nil
				},
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "List of users.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL;")
					if err != nil {
						log.Printf("Failed query list of users. Error: %v \n", err)
					}
					var users []*User

					for rows.Next() {
						user := &User{}

						err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
						if err != nil {
							log.Printf("Failed scan row list of users. Error: %v \n", err)
						}
						users = append(users, user)
					}

					return users, nil
				},
			},
			"task": &graphql.Field{
				Type:        taskType,
				Description: "Get a task.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					task := &Task{}
					err = db.QueryRow("select id, name, description, user_id, is_complete, created_at, updated_at from tasks where id = ? AND deleted_at IS NULL;", id).Scan(&task.ID, &task.Name, &task.Description, &task.UserID, &task.IsComplete, &task.CreatedAt, &task.UpdatedAt)
					if err != nil {
						log.Printf("Failed query select a task. Error: %v \n", err)
					}

					return task, nil
				},
			},
			"tasks": &graphql.Field{
				Type:        graphql.NewList(taskType),
				Description: "List of tasks.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT id, name, description, user_id, is_complete, created_at, updated_at from tasks WHERE deleted_at IS NULL;")
					if err != nil {
						log.Printf("Failed query select list of tasks. Error: %v \n", err)
					}
					var tasks []*Task

					for rows.Next() {
						task := &Task{}

						err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.UserID, &task.IsComplete, &task.CreatedAt, &task.UpdatedAt)
						if err != nil {
							log.Printf("Failed scan rows list of tasks. Error: %v \n", err)
						}
						tasks = append(tasks, task)
					}

					return tasks, nil
				},
			},
		},
	})

	// Create GraphQL mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			// User
			"createUser": &graphql.Field{
				Type:        userType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					email, _ := params.Args["email"].(string)
					createdAt := time.Now()
					updatedAt := time.Now()

					stmt, err := db.Prepare("INSERT INTO users(name, email, created_at, updated_at) VALUES(?,?,?,?);")
					if err != nil {
						log.Printf("Failed prepare query insert new user. Error: %v \n", err)
					}

					res, err := stmt.Exec(name, email, createdAt, updatedAt)
					if err != nil {
						log.Printf("Failed exec query insert new user. Error: %v \n", err)
					}

					lastInsertId, _ := res.LastInsertId()

					newUser := &User{
						ID:        int(lastInsertId),
						Name:      name,
						Email:     email,
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					}

					return newUser, nil
				},
			},
			"updateUser": &graphql.Field{
				Type:        userType,
				Description: "Update a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					email, _ := params.Args["email"].(string)
					updatedAt := time.Now()

					stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?;")
					if err != nil {
						log.Printf("Failed prepare query update user. Error: %v \n", err)
					}

					_, err = stmt.Exec(name, email, updatedAt, id)
					if err != nil {
						log.Printf("Failed exec query update user. Error: %v \n", err)
					}

					updateUser := &User{
						ID:        id,
						Name:      name,
						Email:     email,
						UpdatedAt: updatedAt,
					}

					return updateUser, nil
				},
			},
			"deleteUser": &graphql.Field{
				Type:        userType,
				Description: "Delete a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					deletedAt := time.Now()

					stmt, err := db.Prepare("UPDATE users SET deleted_at = ? WHERE id = ?;")
					if err != nil {
						log.Printf("Failed prepare query delete user. Error: %v \n", err)
					}

					_, err = stmt.Exec(deletedAt, id)
					if err != nil {
						log.Printf("Failed exec query delete user. Error: %v \n", err)
					}

					return nil, nil
				},
			},
			// Task
			"createTask": &graphql.Field{
				Type:        taskType,
				Description: "Create new task",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, _ := params.Args["name"].(string)
					description, _ := params.Args["description"].(string)
					userId, _ := params.Args["user_id"].(int)
					createdAt := time.Now()
					updatedAt := time.Now()

					stmt, err := db.Prepare("INSERT INTO tasks(name, description, user_id, created_at, updated_at) VALUES(?,?,?,?,?);")
					if err != nil {
						log.Printf("Failed prepare query insert new task. Error: %v \n", err)
					}

					res, err := stmt.Exec(name, description, userId, createdAt, updatedAt)
					if err != nil {
						log.Printf("Failed exec query insert new task. Error: %v \n", err)
					}

					lastInsertId, _ := res.LastInsertId()

					newTask := &Task{
						ID:          int(lastInsertId),
						Name:        name,
						Description: description,
						UserID:      userId,
						CreatedAt:   createdAt,
						UpdatedAt:   updatedAt,
					}

					return newTask, nil
				},
			},
			"updateTask": &graphql.Field{
				Type:        taskType,
				Description: "Update a task",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"is_complete": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					description, _ := params.Args["description"].(string)
					userId, _ := params.Args["user_id"].(int)
					isComplete, _ := params.Args["is_complete"].(bool)
					updatedAt := time.Now()

					stmt, err := db.Prepare("UPDATE tasks SET name = ?, description = ?, user_id = ?, is_complete = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL;")
					if err != nil {
						log.Printf("Failed prepare query update task. Error: %v \n", err)
					}

					_, err = stmt.Exec(name, description, userId, isComplete, updatedAt, id)
					if err != nil {
						log.Printf("Failed exec query update task. Error: %v \n", err)
					}

					updateTask := &Task{
						ID:          id,
						Name:        name,
						Description: description,
						UserID:      userId,
						IsComplete:  isComplete,
						UpdatedAt:   time.Time{},
					}

					return updateTask, nil
				},
			},
			"deleteTask": &graphql.Field{
				Type:        taskType,
				Description: "Delete a task",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)
					deletedAt := time.Now()

					stmt, err := db.Prepare("UPDATE tasks SET deleted_at = ? WHERE id = ?;")
					if err != nil {
						log.Printf("Failed prepare query delete task. Error: %v \n", err)
					}

					_, err = stmt.Exec(deletedAt, id)
					if err != nil {
						log.Printf("Failed exec query delete task. Error: %v \n", err)
					}

					return nil, nil
				},
			},
		},
	})
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Start HTTP web server
	http.Handle("/graphql", handler)
	log.Println("Starting Web Server at http://localhost:8000/")
	http.ListenAndServe(":8000", nil)
}
