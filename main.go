package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AldieNightStar/revercity/revercity"
)

func main() {
	var args = os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Enter [connect address] [new port]")
		return
	}

	address := args[0]
	newPort, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Wrong new port")
		return
	}

	connector := revercity.NewTcpConnector(address)
	c, err := revercity.Serve(newPort, connector)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Logging
	for {
		fmt.Println(c)
		time.Sleep(time.Second * 5)
	}
}
