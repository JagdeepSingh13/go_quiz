package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var quizCollection *mongo.Collection

func main() {
	setupDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/quizzes", getQuizzes)

	log.Fatal(app.Listen(":3000"))
}

func setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	quizCollection = client.Database("quiz").Collection("quizzes")
}

func getQuizzes(c *fiber.Ctx) error {
	cur, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	// a slice of map[string]any
	quizzes := []map[string]any{}
	err = cur.All(context.Background(), &quizzes)
	if err != nil {
		panic(err)
	}

	return c.JSON(quizzes)
}
