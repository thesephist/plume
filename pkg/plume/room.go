package plume

type Room struct {
	Sender          chan<- Message
	clientReceivers map[*Client]chan Message
}

func NewRoom() *Room {
	return &Room{
		Sender:          make(chan Message),
		clientReceivers: make(map[*Client]chan Message),
	}
}

func (rm *Room) Enter(u User) *Client {
	receiver := make(chan Message)
	client := Client{
		User:     u,
		Room:     rm,
		Receiver: receiver,
	}

	rm.clientReceivers[&client] = receiver
	go client.StartListening()

	return &client
}

func (rm *Room) Broadcast(msg Message) {
	for _, receiver := range rm.clientReceivers {
		receiver <- msg
	}
}
