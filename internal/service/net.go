package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService

	host *websocket.Conn
	tick int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]

	argument := parts[1]
	switch cmd {
	case "host":
		{
			fmt.Println("host quiz", argument)
			c.host = con
			c.tick = 100
			go func() {
				c.tick--
				c.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(c.tick)))
				time.Sleep(time.Second)
			}()
			break
		}
	case "join":
		{
			fmt.Println("join quiz", argument)
			c.host.WriteMessage(websocket.TextMessage, []byte("A player joined"))
			break
		}
	}
}
