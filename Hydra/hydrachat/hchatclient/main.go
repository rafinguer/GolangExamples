package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	//"Hydra/HydraConfigurator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonumous%d", rand.Intn(500))
	fmt.Println("Starting HydraChatClient...")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	/*
		confStruct := struct {
			Name string `name:"name"`
			RemoteAddr string `name:"remoteip"`
			TCP bool `name:"tcp"`
		}

		HydraConfigurator.GetConfiguration(HydraConfigurator.CUSTOM, &confStruct, "chat.conf")
		name = confStruct.Name
		proto := "tcp"
		if !confStruct.TCP {
			proto:= "udp"
		}
	*/

	fmt.Println("Hello %s, connecting to the Hydra chat system...\n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to Hydra chat system: ", err)
	}
	fmt.Println("Connected to Hydra chat system")
	name += ":"

	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	/*
		for err == nil {
			msg := ""
			fmt.Print(name)
			fmt.Scan(&msg)
			msg = name+msg+"\n"
			fmt.Prinln("Duplicate: " + msg)
			_, err = fmt.Fprintf(conn,msg)
		}
	*/

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}
}
