package visualizer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/StageAutoControl/controller/pkg/cntl"
	"github.com/StageAutoControl/controller/pkg/internal/logging"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

type Server struct {
	logger     logging.Logger
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func NewServer(logger logging.Logger) *Server {
	return &Server{
		logger:     logger,
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
	}
}

// Write the given command to the sockets
func (s *Server) Write(cmd cntl.Command) error {
	b, err := json.Marshal(&cmd)
	if err != nil {
		return err
	}

	s.send(b)
	return nil
}

// send the given byte slide to all sockets
func (s *Server) send(msg []byte) {
	s.broadcast <- msg
}

// Run the server
func (s *Server) Run(ctx context.Context) {
	go func() {
		for message := range s.broadcast {
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			s.stop()
			return
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
			}

		}
	}
}

func (s *Server) stop() {
	close(s.broadcast)
	close(s.register)
	close(s.unregister)

	for c := range s.clients {
		c.stop()
	}
}

// ServeRequest handles incoming http connections and upgrades them
func (s *Server) ServeRequest(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &client{
		logger: s.logger,
		server: s,
		conn:   conn,
		send:   make(chan []byte, 1024),
	}
	s.register <- client

	// we don't care for what the client tells us
	go client.read()
	client.write()
}
