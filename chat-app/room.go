package main

import (
  "github.com/gorilla/websocket"
  "net/http"
  "log"
  "chat/trace"
  "github.com/stretchr/objx"
)

const (
  socketBufferSize = 1024
  messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize:socketBufferSize,
  WriteBufferSize:socketBufferSize}

type room struct {
  forward chan *message
  join chan *client
  leave chan *client

  clients map[*client]bool
  tracer trace.Tracer
}

func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      r.clients[client] = true
      r.tracer.Trace("New client joined")
    case client := <-r.leave:
      delete(r.clients, client)
      close(client.send)
      r.tracer.Trace("Client is leaving")
    case msg := <- r.forward:
      for client := range r.clients {
        client.send <- msg
      }
      r.tracer.Trace("All messages are sent")
    }
  }
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  socket, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
    log.Fatal("ServeHttp", err)
    return
  }
  authCookie, err := req.Cookie("auth")
  if err != nil {

  }
  client := &client{
    socket:socket,
    send: make(chan *message, messageBufferSize),
    room:r,
    userData: objx.MustFromBase64(authCookie.Value),
  }
  r.join <- client
  defer func() { r.leave <- client}()
  go client.write()
  client.read()
}

func newRoom() *room {
  return &room{
    forward: make(chan *message),
    join: make(chan *client),
    leave: make(chan *client),
    clients: make(map[*client]bool),
    tracer: trace.Off(),
  }
}
