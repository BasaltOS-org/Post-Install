package main

import (
	"time"
	"BasaltPostInstallAssistant/internal/ipc"


)

func main(){

	go ipc.Server()
	go ipc.Client()
	time.Sleep(3000 * time.Second)
}