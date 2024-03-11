package testsocketserver

import (
	"fmt"
	"net"
)

// StartServer - запуск тестового сервера
func StartServer(typeConn string, address string) {
	listner, _ := net.Listen(typeConn, "localhost:8081")
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

// handleClient - хендлер обращений к серверу
func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		readLen, err := conn.Read(buf)
		fmt.Println("Readed data", string(buf))
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		conn.Write(append([]byte("You send "), buf[:readLen]...))
	}
}
