package socketclient

import (
	"context"
	"fmt"
	"net"
	"time"
)

// Socket - интерфейс сокет клиента
type Socket interface {
	Connect() error
	Disconnect(context.Context) error
	Send([]byte) error
	Read() ([]byte, error)
}

// TypeConnection -  тип сокет соединения
type TypeConnection string

const (
	//TCP -Тип соединения tcp
	TCP = "tcp"
	//UDP -Тип соединения udp
	UDP = "udp"
)

// SocketClient - реализация сокет клиента
type SocketClient struct {
	connection TypeConnection
	socket     net.Conn
	timeout    uint
	addres     string
}

// NewClient - конструктор SocketClient
func NewClient(connection TypeConnection, host string, port string, time uint) Socket {
	return &SocketClient{
		connection: connection,
		timeout:    time,
		addres:     host + ":" + port,
	}
}

// Connect - установить соединение с сокет сервером
func (s *SocketClient) Connect() error {
	var err error
	s.socket, err = net.DialTimeout(string(s.connection), s.addres, time.Second*time.Duration(s.timeout))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = s.socket.SetDeadline(time.Now().Add(time.Second * time.Duration(s.timeout)))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Succesed connect to", s.addres)

	return nil
}

// Disconnect - разорвать соект соединение с сервером
func (s *SocketClient) Disconnect(context context.Context) error {
	if err := s.socket.Close(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Succesed disconect from", s.addres)
	return nil
}

// Send - отправить данные на сокет сервер
func (s *SocketClient) Send(data []byte) error {
	_, err := s.socket.Write(data)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Success send data")
	return nil
}

// Read - читать данные с сокет сервера
func (s *SocketClient) Read() ([]byte, error) {
	data := make([]byte, 1024)
	_, err := s.socket.Read(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Success read data")

	return data, nil
}
