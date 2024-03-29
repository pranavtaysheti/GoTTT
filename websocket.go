package main

import (
	"log"
	"io"

	"golang.org/x/net/websocket"
)

func wsHandler(ws *websocket.Conn) {
	connections[ws] = wssConnecting

	buffer := make([]byte, 1024)
	var err error = nil

	for err != io.EOF {
		_, wsErr := ws.Read(buffer)
		err = wsErr
	}

	connections[ws] = wssDisconnected
	log.Println(err)
	return
}
