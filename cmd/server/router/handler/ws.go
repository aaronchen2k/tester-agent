package handler

import (
	_logUtils "github.com/aaronchen2k/tester/internal/pkg/libs/log"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

type WsCtrl struct {
	BaseCtrl
	*neffos.NSConn `stateless:"true"`
	Namespace      string
}

func NewWsCtrl() *WsCtrl {
	return &WsCtrl{Namespace: "default"}
}

func (c *WsCtrl) OnNamespaceConnected(msg neffos.Message) error {
	_logUtils.Infof("%s connected", c.Conn.ID())
	return nil
}

func (c *WsCtrl) OnNamespaceDisconnect(msg neffos.Message) error {
	_logUtils.Infof("%s disconnected", c.Conn.ID())
	return nil
}

func (c *WsCtrl) OnChat(msg websocket.Message) error {
	_logUtils.Infof("%s", msg)
	return nil
}

func (c *WsCtrl) Chat(msg websocket.Message, ns *neffos.NSConn) (ret interface{}, err error) {
	_logUtils.Infof("%s", msg)
	ret = msg
	return
}
