package main

import "github.com/Perazzojoao/WebSocket-ChatApp/src/server"

func main() {
	app := server.NewServer()
	app.Start()
}
