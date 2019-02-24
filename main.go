package main

import (
	"fmt"
	"net"
	"flag"
	"log"
	"strconv"
	"encoding/hex"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "h", "", "host")
	flag.IntVar(&port, "p", 3333, "port")
}

func main() {
	flag.Parse()
	fmt.Println(port)
	address := host + ":" + strconv.FormatInt(int64(port), 10)
	log.Printf("listen to %s", address)

	var l net.Listener
	l, err := net.Listen("tcp", address)

	if err != nil {
		log.Print(err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("accept failed, %v", err)
			break
		}
		log.Printf("accept a new connection %v", c.RemoteAddr())
		go handle(c)
	}
}

func handle(conn net.Conn) {
	defer func() {
		log.Printf("closed connection %v", conn.RemoteAddr())
		conn.Close()
	}()
	var buf = make([]byte, 65536)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read failed %v", err)
			return
		}
		if n > 0 {
			x := make([]byte, n)
			copy(x, buf[0:n])
			fmt.Print(hex.Dump(x))
			fmt.Printf("received %v \n", string(x))
		}
	}
}
