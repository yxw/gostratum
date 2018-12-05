package main

import (
	"antpool-mining-btm/session"
	"log"

	//"github.com/devktor/gostratum"
	"antpool-mining-btm/jedi/gostratum"
	"bufio"
	"fmt"
	"os"
)

type BtmLoginParams struct {
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: stratum_host address")
		os.Exit(1)
	}
	host := os.Args[1]
	address := os.Args[2]

	client, err := gostratum.Connect(host)

	if err != nil {
		fmt.Println("failed to connect: ", err)
		os.Exit(2)
	}

	fmt.Println("enter to exit")

	/*
			client.AddressSubscribe(address, func(address string, status string, err error) {
				fmt.Println("address update: ", address, " status=", status, " error=", err)
		    })
	*/

	//client.AddressSubscribe(address, func(address string, status string, err error) {
	//	fmt.Println("address update: ", address, " status=", status, " error=", err)
	//})
	//return c.Subscribe("blockchain.address.subscribe", WrapAddressHandler(callback, c.decoder), func(result *json.RawMessage) {
	//	status, err := c.decoder.DecodeString(result)
	//	callback(address, status, err)
	//}, address)
	//client.Subscribe("job", WrapJobHandler(handleJobNotification, client.decoder), nil, nil)

	loginParams := &session.LoginParams{
		Login: address,
		Pass:  "password",
		Agent: "xminer/1.0.0"}

	response := client.Request("login", loginParams)
	log.Printf("got response: %v", *response)

	bufio.NewReader(os.Stdin).ReadString('\n')

}
