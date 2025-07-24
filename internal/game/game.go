package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/JagdeepSingh13/go_quiz/internal/entity"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Player struct {
	Name       string
	Connection *websocket.Conn
}

type Game struct {
	Id     uuid.UUID
	Quiz   entity.Quiz
	Code   string
	Player []Player

	Host *websocket.Conn
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func New(quiz entity.Quiz, host *websocket.Conn) Game {
	return Game{
		Id:     uuid.New(),
		Quiz:   quiz,
		Code:   generateCode(),
		Player: []Player{},
		Host:   host,
	}
}

func (g *Game) Start() {
	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {

}

func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	g.Player = append(g.Player, Player{
		Name:       name,
		Connection: connection,
	})
}
