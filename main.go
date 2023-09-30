package main

import (
	"github.com/peacecwz/modem-sdk/modems/sagemcom"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	// Test codes
	client := sagemcom.NewSagemcomClient("http://192.168.178.1/rest/v1/")
	_, err := client.Login("")
	if err != nil {
		log.Panic(err)
	}

	resp, err := client.PortForwarding.GetPortForwardings()
	if err != nil {
		log.Panic(err)
	}

	logrus.Infof("PortForwarding: %+v", resp)

	defer client.Logout()
}
