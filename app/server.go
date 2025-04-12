package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/abdellani/go-build-your-own-redis/app/config"
	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
	"github.com/abdellani/go-build-your-own-redis/app/handler"
	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	startServer(
		handler.NewHandler(storage.New(), LoadConfigurations()))
}

func LoadConfigurations() *config.Config {
	configurations := config.New()
	flag.StringVar(&configurations.RDB.Dir, "dir", "", "the directory where RDB will be stored")
	flag.StringVar(&configurations.RDB.FileName, "dbfilename", "", "this RDB file name")
	flag.Parse()
	return configurations
}
func startServer(handler *handler.Handler) {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn, handler)
	}
}

func handleConnection(conn net.Conn, handler *handler.Handler) {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		received := bytes.Buffer{}
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		received.Write(buf[:n])
		des := deserializer.New(received.String())
		command := des.Deserialize()
		resp := handler.Handle(command)
		conn.Write([]byte(resp))
	}
}
