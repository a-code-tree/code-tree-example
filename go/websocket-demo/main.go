package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//go:embed index.html static/*
var content embed.FS

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}



func handleConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// 读取客户端发送的消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Received message: %s\n", p)

		// 向客户端发送相同的消息
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 提供嵌入的HTML页面
		data, err := content.ReadFile("index.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	http.HandleFunc("/ws", handleConnection)

	fmt.Println("Server is running on :18082")
	err := http.ListenAndServe(":18082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
