package model

import (
	"log"
)

type hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan *Envelope

	// Register requests from the connections.
	Register chan *Connection

	// Unregister requests from connections.
	Unregister chan *Connection
}

type msgContainer struct {
}

var Hub = hub{
	broadcast:   make(chan *Envelope),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func (h *hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.registerConnection(c)
		case c := <-h.Unregister:
			h.closeConnection(c)
		case m := <-h.broadcast:
			h.handleMessage(m)
		}
	}
}

func (h *hub) registerConnection(c *Connection) {
	h.connections[c] = true

	log.Println("Connection registered")

	// Send info message to everybody
	h.handleMessage(createInfoMessage(h))
}

func (h *hub) closeConnection(c *Connection) {
	delete(h.connections, c)
	close(c.Send)

	// Send info message to everybody
	h.handleMessage(createInfoMessage(h))
}

func (h *hub) handleMessage(msg *Envelope) {

	for c := range h.connections {
		select {
		case c.Send <- msg:
		default:
			delete(h.connections, c)
			close(c.Send)
			go c.Ws.Close()
		}
	}
}

func createInfoMessage(hub *hub) *Envelope {
	// Create info payload
	info := &Info{UserCount: len(hub.connections)}

	// Create message envelope
	return &Envelope{Action: INFO, Info: info}
}
