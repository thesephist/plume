package plume

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	Rooms []*Room
}

func (srv *Server) handleUpgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		log.Println(err)
		return
	}

	var client *Client
	for {
		var msg Message

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)

			if client != nil {
				client.Send("left")
			}

			break
		}

		switch msg.Type {
		case MsgHello:
			textParts := strings.Split(msg.Text, "\n")
			if len(textParts) != 2 {
				// malformed hello message
				break
			}

			u := User{
				Name:  textParts[0],
				Email: textParts[1],
			}
			client = srv.Rooms[0].Enter(u)

			client.OnMessage = func(msg Message) {
				conn.WriteJSON(msg)
			}

			log.Printf("User @%s entered", u.Name)
			client.Send("entered")
		case MsgText:
			if client == nil {
				break
			}

			client.Send(msg.Text)
		default:
			log.Printf("unknown message type: %v", msg)
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	indexFile, err := os.Open("./static/index.html")
	defer indexFile.Close()

	if err != nil {
		io.WriteString(w, "error reading index")
		return
	}

	io.Copy(w, indexFile)
}

func StartServer() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:4884",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	plumeSrv := Server{
		Rooms: []*Room{},
	}

	// For now, have one room
	plumeSrv.Rooms = append(plumeSrv.Rooms, NewRoom())

	r.HandleFunc("/", handleHome)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		plumeSrv.handleUpgrade(w, r)
	})

	log.Fatal(srv.ListenAndServe())
}
