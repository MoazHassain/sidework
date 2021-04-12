package main

import (
	"fmt"
	"net"
)

func main() {

	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", ":8888")
	// conn, err = net.Dial("tcp", "91.205.173.170:8888")

	if err != nil {
		fmt.Println(err.Error())
	}

	// bs := make([]byte, 100)
	// conn.Write()

	conn.Write([]byte("Hello from Moaz Hossain"))

	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(string(bs[:n]))

	conn.Close()
}
