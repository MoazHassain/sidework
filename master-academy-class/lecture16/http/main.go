package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}

	bs := make([]byte, 1024)

	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println(n)

	reqstr := string(bs)
	fmt.Println(reqstr)

	body := `<!DOCTYPE html>
		<html lang = "en">
		<head>
			<meta charset="UTF-8">
			<title>Document</title>
		</head>
		<body>
			<h1>hello world</h1>
		</body>
		</html>`

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-length: %d\r\n" +
		"Content-type: text/html\r\n" +
		"\r\n%s"

	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))
}
