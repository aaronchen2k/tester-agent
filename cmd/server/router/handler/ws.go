package handler

import (
	"fmt"
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/kataras/iris/v12/websocket"
)

type WsCtrl struct {
	BaseCtrl
	*websocket.NSConn `stateless:"true"`
	Namespace         string
}

func NewWsCtrl() *WsCtrl {
	return &WsCtrl{Namespace: "default"}
}

func (c *WsCtrl) OnNamespaceConnected(msg websocket.Message) error {
	_logUtils.Infof("%s connected", c.Conn.ID())
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%d", 1)),
	})
	return nil
}

// This will call the "OnVisit" event on all clients, except the current one,
// (it can't because it's left but for any case use this type of design)
func (c *WsCtrl) OnNamespaceDisconnect(msg websocket.Message) error {
	_logUtils.Infof("%s disconnected", c.Conn.ID())
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%d", 2)),
	})
	return nil
}

// This will call the "OnVisit" event on all clients, including the current one,
// with the 'newCount' variable.
func (c *WsCtrl) OnChat(msg websocket.Message) (err error) {
	ctx := websocket.GetContext(c.Conn)

	str := ctx.RemoteAddr()
	_logUtils.Info(str + ", " + string(msg.Body))

	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnChat",
		Body:      []byte(fmt.Sprintf("%d", 2)),
	})

	return
}
