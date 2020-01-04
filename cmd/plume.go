package main

import (
	"fmt"

	"github.com/thesephist/plume/pkg/plume"
)

func main() {
	fmt.Println("Starting Plume server...\n")

	plume.StartServer()
}
