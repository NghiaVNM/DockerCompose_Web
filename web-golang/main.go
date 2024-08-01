package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    r := gin.Default()

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017").SetMaxPoolSize(100))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    db := client.Database("database")
    usersCollection := db.Collection("users")

    r.GET("/", func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        var users []User
        cursor, err := usersCollection.Find(ctx, bson.M{})
        if err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }
        defer cursor.Close(ctx)

        for cursor.Next(ctx) {
            var user User
            cursor.Decode(&user)
            users = append(users, user)
        }

        if err := cursor.Err(); err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }

        userList := ""
        for _, user := range users {
            userList += fmt.Sprintf("<li>%s - %d</li>", user.Name, user.Age)
        }

        c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
            <html>
                <head>
                    <title>Golang Web App</title>
                </head>
                <body>
                    <h1>Hello World from Golang</h1>
                    <h2>User List</h2>
                    <ul>
                        %s
                    </ul>
                </body>
            </html>
        `, userList)))
    })

    r.Run(":3002") 
}