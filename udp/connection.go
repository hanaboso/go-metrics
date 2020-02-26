package udp

import (
	"fmt"
	"log"
	"net"
	"time"
)

const errorFormat = "[UDP] %+v"

// Connection represents UDP connection
type Connection struct {
	UDP *net.UDPConn
}

// Connect creates UDP connection
func (connection *Connection) Connect(dsn string) {
	address, err := net.ResolveUDPAddr("udp", dsn)

	if err != nil {
		log.Println(fmt.Sprintf(errorFormat, err))
		time.Sleep(time.Second)
		connection.Connect(dsn)

		return
	}

	connection.UDP, err = net.DialUDP("udp", nil, address)

	if err != nil {
		log.Println(fmt.Sprintf(errorFormat, err))
		time.Sleep(time.Second)
		connection.Connect(dsn)

		return
	}

	ticker := time.NewTicker(time.Minute)

	go func() {
		for range ticker.C {
			connection.Connect(dsn)
		}
	}()
}

// Disconnect from UDP
func (connection *Connection) Disconnect() {
	if err := connection.UDP.Close(); err != nil {
		log.Println(fmt.Sprintf(errorFormat, err))
		time.Sleep(time.Second)
		connection.Disconnect()

		return
	}

	connection.UDP = nil
}

// IsConnected checks connection status
func (connection *Connection) IsConnected() bool {
	return connection.UDP != nil
}
