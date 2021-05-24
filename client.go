package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	index int
	conn  *websocket.Conn
}

type WsMsg struct {
	mt  int
	msg []byte
	idx int
}

func NewClient(index int, conn *websocket.Conn) *Client {
	return &Client{index: index, conn: conn}
}

func (c *Client) Start(op chan WsMsg) {
	fmt.Println("starting client")
	for {
		mt, msg, _ := c.conn.ReadMessage()
		op <- WsMsg{mt: mt, msg: msg, idx: c.index}
	}
}

func (c *Client) Close() {
	c.conn.Close()
}
