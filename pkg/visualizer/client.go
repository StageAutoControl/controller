package visualizer

import (
	"github.com/gorilla/websocket"

	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

var (
	newline = []byte{'\n'}
)

type client struct {
	logger logging.Logger
	server *Server
	conn   *websocket.Conn
	send   chan []byte
}

func (c *client) read() {
	defer func() {
		c.server.unregister <- c
		if err := c.conn.Close(); err != nil {
			c.logger.Error(err)
		}
	}()
	for {
		if _, _, err := c.conn.NextReader(); err != nil {
			if err := c.conn.Close(); err != nil {
				c.logger.Error(err)
			}
			break
		}
	}
}

func (c *client) stop() {
	close(c.send)
}

func (c *client) write() {
	for message := range c.send {
		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}
		if _, err := w.Write(message); err != nil {
			c.logger.Error(err)
		}

		n := len(c.send)
		for i := 0; i < n; i++ {
			if _, err := w.Write(newline); err != nil {
				c.logger.Error(err)
			}
			if _, err := w.Write(<-c.send); err != nil {
				c.logger.Error(err)
			}
		}

		if err := w.Close(); err != nil {
			return
		}
	}
}
