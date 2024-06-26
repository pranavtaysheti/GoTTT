//ABBREVATIONS:
//wsm = WebSocket message
//aep = API Endpoint

const wsmUpdateRoom = "UPDATE_ROOM";
const wsmUpdateBoard = "UPDATE_BOARD";
const wsmConnectionReady = "CONNECTION_READY";

const socket = new WebSocket("ws://localhost:3000/ws");

const updateBoard = () => {
  //TODO
}

const updateRoom = () => {
  //TODO
}

const handleMessage = (message) => {
  switch (message) {
    case wsmUpdateRoom:
      updateRoom();
      break;
    case wsmUpdateBoard:
      updateBoard();
      break;
    default:
      console.error("Recieved unrecognized message from socket", message);
  }
}

socket.onopen = (_ev) => {
  socket.send(wsmConnectionReady)
}

socket.onmessage = (ev) => {
  handleMessage(ev.data)
}