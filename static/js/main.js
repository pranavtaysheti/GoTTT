const socket = new WebSocket("ws://localhost:3000/ws")

const handleMessage = (message) => {
  if (message == "UPDATE_ROOM") {

  }
  else if (message == "UPDATE_BOARD") {

  }
}

socket.onopen = (_ev) => {
  socket.send("CONNECTION_READY")
}

socket.onmessage = (ev) => {
  handleMessage(ev.data)
}

