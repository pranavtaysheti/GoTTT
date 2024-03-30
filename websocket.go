package main

import (
	"log"
	"io"

	"golang.org/x/net/websocket"
)

func wsHandler(ws *websocket.Conn) {
	c, err := getClientCookie(ws.Request())
	if err != nil {
		log.Println("Cookie not found. Could not associate websocket with a client.")
		return
	}

	cl, err := getClientFromCookie(c)
	if err != nil {
		log.Println("Client Object not found. Could not associate websocket with client")
		return
	}

	cl.socket = ws
	wsLoop(ws)
}

func wsLoop(ws *websocket.Conn) {
	b := make([]byte, 1024)

	var err error
	for err != io.EOF {
		if err != nil {
			log.Println("Error reading websocket:", err)
		}
		_, lerr := ws.Read(b)
		err = lerr
	} 
}
