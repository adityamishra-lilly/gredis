package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/adityamishra-lilly/gredis/config"
)

func readCommand(conn net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := conn.Read(buf[:])
	if err != nil{
		return "", err
	}

	return string(buf[:n]), nil
}

func writeCommand(conn net.Conn, cmd string) error {
	if _, err := conn.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil

}

func StartServer() {
	log.Println("Starting the synchronous TCP Server on", config.Host, config.Port)
	
	var clients int = 0
 
	listener, err := net.Listen("tcp", config.Host + ":" + strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for{

		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		clients++
		log.Println("client connected with address:", conn.RemoteAddr(), "concurrent clients: ", clients)

		for{
			cmd, err := readCommand(conn)
			if err != nil {
				conn.Close()
				clients--
				log.Println("client disconnected", conn.RemoteAddr(), "concurrent clients", clients)
				if err == io.EOF{
					break
				}
				log.Println("err", err)
			}
			log.Println("command", cmd)
			if err = writeCommand(conn, cmd); err != nil {
				log.Print("err write:", err)
			}
		}
	} 	
}
