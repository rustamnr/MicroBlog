package main

import (
	"log"

	microblog "github.com/lsmltesting/MicroBlog"
)

func main() {
	server := new(microblog.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
