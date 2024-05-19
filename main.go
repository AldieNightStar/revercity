package main

import (
	"fmt"
	"time"

	"github.com/AldieNightStar/revercity/revercity"
)

func main() {
	var connector = revercity.NewTcpConnector("localhost:25565")
	var c, err = revercity.Serve(7777, connector)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println(c)
		time.Sleep(time.Second)
	}
}
