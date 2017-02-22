//websocket server
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

const (
	LISTENADDR = ":1234"
)

var(
	clientMap map[string]*websocket.Conn
)

func getHeartBeat(ws *websocket.Conn) {

	clientStr := ws.Request().RemoteAddr
	fmt.Println("[debug] client a:"+ clientStr)

	if _, ok := clientMap[clientStr]; ok {
		log.Println("client is in")
	} else {
		log.Println("client is no in")
		clientMap[clientStr] = ws
		go sendBroadcast(nil, fmt.Sprintf("Hello, %s join\n", clientStr))
		for ck, _ := range clientMap {
			log.Println("c in map: ", ck)
		}
	}

	defer func(){
		if err := ws.Close(); err != nil {
			log.Println("[err] Websocket could't bu closed,", err.Error())
		} else {
			if _, ok := clientMap[clientStr]; ok {
				delete(clientMap, clientStr)
			}
			log.Println("websocket closed.", clientStr)
		}
	}()

	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("[err] Can't recieve", err.Error())
			return
		}

		fmt.Println("[debug] client b:"+ ws.Request().RemoteAddr)

		clientMap[ws.Request().RemoteAddr] = ws

		log.Println("[info] Recieved back from client["+ws.Request().RemoteAddr +"]: " + reply)
		t := time.Now()
		msg := fmt.Sprintf("Recieved: %s, // %s", reply, t.Format(time.RFC3339))

		log.Println("[info] Sending to client:" + msg)

		go sendBroadcast(ws, msg)
	}
}

func sendBroadcast(conn *websocket.Conn, msg string){
	for ck, client := range clientMap {
		if conn == client {
			log.Println("Continue....")
//			continue
		}
		log.Printf("send => %s ", ck)

		websocket.Message.Send(client, msg)
	}
}

func getClientMap(w http.ResponseWriter, r *http.Request){
	for ck, _ := range clientMap {
		log.Println("Client" ,ck)
	}
}

func main() {
	clientMap =  make(map[string]*websocket.Conn)
	http.Handle("/", websocket.Handler(getHeartBeat))

	http.HandleFunc("/cm", getClientMap)

	if err := http.ListenAndServe(LISTENADDR, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
