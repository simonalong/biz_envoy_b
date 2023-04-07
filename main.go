package main

import (
	"biz-b-service/router"
	"github.com/isyscore/isc-gobase/server"
)

func main() {
	router.Register()
	server.Run()
}
