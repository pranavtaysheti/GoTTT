package main

import (
	"log"
	"io"

	"golang.org/x/net/websocket"
)

const (
	wssDisconnected = iota
	wssConnecting
	wssConnected
)

const (
	wsmUpdateRoom      = "UPDATE_ROOM"
	wsmUpdateBoard     = "UPDATE_BOARD"
	wsmConnectionReady = "CONNECTION_READY"
)

func wsHandler(ws *websocket.Conn) {
	cl, err := getClientFromRequest(ws.Request())
	if err != nil {
		log.Println(err)
		return
	}

	err = cl.ConnectSocket(ws)
	if err != nil {
		log.Println(err)
		return
	}

	b := make([]byte, 1024)
	
	for err != io.EOF {
		if err != nil {
			log.Println("Error reading websocket:", err)
		}
		_, lerr := ws.Read(b)
		err = lerr
	}

	cl.socket = nil
}

