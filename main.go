package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"p1/router"
	"p1/setting"
	"syscall"
)

func main() {
	r := router.InitRouter()

	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
