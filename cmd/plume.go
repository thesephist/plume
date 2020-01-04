package main

import (
	"fmt"

	"github.com/thesephist/plume"
)

func main() {
	fmt.Println("Starting Plume server...\n")

	room := plume.NewRoom()
	me := plume.User{
		Name:  "thesephist",
		Email: "linus@thesephist.com",
	}
	bot := plume.User{
		Name:  "bot",
		Email: "bot@plume.chat",
	}

	meClient := room.Enter(me)
	botClient := room.Enter(bot)

	meClient.Send("hello from me")
	botClient.Send("hi from bot")
}
