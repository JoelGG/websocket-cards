package router

import (
	"fmt"
	"time"
)

type Message struct {
	content string
	client  *Client
	time    time.Time
}

func NewMessage(content string, client *Client) *Message {
	m := Message{}
	m.content = content
	m.client = client
	m.time = time.Now()
	return &m
}

func (m *Message) ToString() string {
	return fmt.Sprintf("%d at %s: %s", m.client.index, m.time.Local().Format(time.Kitchen), m.content)
}
