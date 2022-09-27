package main

import (
	"esb-worker/internal/core"
	"esb-worker/registry"
)

func ESBInit() {
	core.Start(registry.Registry())
}

func main() {

	//Start ESB
	forever := make(chan bool)
	go ESBInit()
	<-forever

}
