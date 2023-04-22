package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	log.Printf("server listening %s \n", listen.Addr())
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Accept err: %s \n", err.Error())
			continue
		}
		HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("client close err: %s", err.Error())
		}
	}()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05 \n"))
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "close")
			return
		}
		time.Sleep(time.Second * 1)
	}
}
