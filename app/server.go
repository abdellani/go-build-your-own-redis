package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/abdellani/go-build-your-own-redis/app/config"
	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
	"github.com/abdellani/go-build-your-own-redis/app/handler"
	"github.com/abdellani/go-build-your-own-redis/app/rdb"
	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	configurations := LoadConfigurations()
	storage := storage.New()
	LoadRDB(configurations, storage)
	startServer(
		handler.NewHandler(storage, configurations))
}

func LoadRDB(configurations *config.Config, storage *storage.Storage) {
	filePath := filepath.Join(configurations.RDB.Dir, configurations.RDB.FileName)
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("RDB not loaded")
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("RDB not loaded")
		return
	}
	decoder := rdb.New(data)
	decoder.Decode()
	decoder.WriteEntries(storage)

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
