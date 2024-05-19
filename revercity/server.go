package revercity

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Serve(port int, con Connector) (*Control, error) {
	var address = fmt.Sprintf("0.0.0.0:%d", port)
	var serv, err = net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	var control = newControl()
	go serveAll(serv, con, control)
	return control, nil
}

func serveAll(serv net.Listener, con Connector, control *Control) {
	for {
		var sock, err = serv.Accept()
		if err != nil {
			log.Println("Had error during connect:", err)
			continue
		}
		go serveSocket(sock, con, control)
	}
}

func serveSocket(clientSock net.Conn, con Connector, control *Control) {
	var servSock, err = con.Connect()
	if err != nil {
		log.Println("Error during connection to reverse:", err)
		return
	}

	log.Println("Connected new Client")
	control.mut.Lock()
	control.connections += 1
	control.mut.Unlock()

	var toBreak = false
	for {
		// Break if needed
		if toBreak {
			break
		}

		if control.toStop {
			log.Println("Stopping due to Control Stop()")
			servSock.Close()
			clientSock.Close()

			control.mut.Lock()
			control.connections -= 1
			control.mut.Unlock()
			break
		}

		// Client -> Server
		go func() {
			if _, err = io.Copy(clientSock, servSock); err != nil {
				if err == io.EOF {
					log.Println("Client -> Server : Disconnect")
				} else {
					log.Println("Error during Client -> Server:", err)
				}

				servSock.Close()
				clientSock.Close()

				control.mut.Lock()
				control.fails += 1
				control.connections -= 1
				control.mut.Unlock()
			}
		}()

		// Server -> Client
		go func() {
			if _, err = io.Copy(servSock, clientSock); err != nil {
				if err == io.EOF {
					log.Println("Client -> Server : Disconnect")
				} else {
					log.Println("Error during Server -> Client:", err)
				}

				servSock.Close()
				clientSock.Close()

				control.mut.Lock()
				control.fails += 1
				control.connections -= 1
				control.mut.Unlock()
				toBreak = true
			}
		}()

		// Sleep for CPU safe usage
		time.Sleep(time.Millisecond)
	}
}
