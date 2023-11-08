package main

import (
	"gintraining/webapi"
)

func main() {
	srv := webapi.InitServer()
	srv.ServerRun()
}
