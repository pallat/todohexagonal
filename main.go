package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/pallat/todoapi/router"
	"github.com/pallat/todoapi/store"
	"github.com/pallat/todoapi/todo"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("myapp").Collection("todos")

	r := router.NewFiberRouter()

	// gormStore := store.NewGormStore(db)
	mongoStore := store.NewMongoDBStore(collection)

	handler := todo.NewTodoHandler(mongoStore)
	r.POST("/todos", handler.NewTask)

	if err := r.Listen(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
