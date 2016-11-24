package main

import (
	"github.com/msawangwan/mhttplib/msrvd"
)

func main() {
	s := msrvd.NewStaticContentHandler()
	s.ServeStaticContent(":8000")
}
