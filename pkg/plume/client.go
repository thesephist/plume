package plume

type Client struct {
	User     User
	Room     *Room
	Receiver <-chan Message
}

func (cl *Client) Send(text string) error {
	if cl.Room == nil {
		return PlumeError{"client is not in a room yet"}
	}

	cl.Room.Broadcast(Message{
		User: cl.User,
		Text: text,
	})

	return nil
}

func (cl *Client) Leave() error {
	if cl.Room == nil {
		return PlumeError{"client is not in a room yet"}
	}

	delete(cl.Room.clientReceivers, cl)
	cl.Room = nil

	return nil
}

func (cl *Client) StartListening() {
	for {
		msg := <-cl.Receiver
		// TODO: need to determine what API we need to hook into when
		// a client receives a message
		fmt.Printf("@%s: %s\n", msg.User.Name, msg.Text)
	}
}
