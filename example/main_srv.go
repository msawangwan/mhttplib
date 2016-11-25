package main

import (
	"github.com/msawangwan/mhttplib/msrvd"
)

const (
	LISTEN_PORT = "127.0.0.1:9090"
)

func main() {
	server := msrvd.NewStaticContentHandler()
	server.ListenAndServeStaticContent(LISTEN_PORT)
}
