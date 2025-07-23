package internal

import (
	"context"
	"log"
	"time"

	"github.com/JagdeepSingh13/go_quiz/internal/collection"
	"github.com/JagdeepSingh13/go_quiz/internal/controller"
	"github.com/JagdeepSingh13/go_quiz/internal/service"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

type App struct {
	httpServer  *fiber.App
	database    *mongo.Database
	quizService *service.QuizService
}

func (a *App) Init() {
	a.setupDB()
	a.setupServices()
	a.setupHttp()

	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.Ws()
	app.Get("/ws", websocket.New(wsController.Ws))

	log.Fatal(app.Listen(":3000"))

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
}

func (a *App) setupDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("quiz")
}
