package ws

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsMessage struct {
	MsgType MsgType
	Message string
}

const (
	Clear MsgType = iota
	Message
)

type MsgType int

type WS struct {
	upgrader   websocket.Upgrader
	sessionMap map[uuid.UUID]*websocket.Conn
	mutex      sync.Mutex
}

func NewWs() *WS {
	return &WS{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sessionMap: make(map[uuid.UUID]*websocket.Conn),
		mutex:      sync.Mutex{},
	}
}

func (me *WS) Handle(c *gin.Context) {
	wsSession, err := me.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		wsSession.Close()
		log.Println(err)
	}
	me.sessionMap[uuid.New()] = wsSession

}

func (me *WS) BroadcastJSON(t MsgType, msg string) {
	me.mutex.Lock()
	println(msg)
	m := WsMessage{MsgType: t, Message: msg}
	for key, wsSession := range me.sessionMap {
		err := wsSession.WriteJSON(m)
		if err != nil {
			log.Println("error while writing to session")
			delete(me.sessionMap, key)
		}
	}
	time.Sleep(50 * time.Millisecond)
	me.mutex.Unlock()
}
