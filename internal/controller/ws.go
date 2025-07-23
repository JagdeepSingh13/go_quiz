package controller

import (
	"log"

	"github.com/JagdeepSingh13/go_quiz/internal/service"
	"github.com/gofiber/contrib/websocket"
)

type WebSocketController struct {
	netService *service.NetService
}

func Ws(netService *service.NetService) WebSocketController {
	return WebSocketController{
		netService: netService,
	}
}

func (c *WebSocketController) Ws(con *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = con.ReadMessage(); err != nil {
			log.Println("read", err)
			break
		}

		c.netService.OnIncomingMessage(con, mt, msg)
	}
}
