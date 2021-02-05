package main

import (
	"context"
	"log"

	"github.com/writeameer/vmcli/config"
)

var (
	ctx = context.Background()
)

func main() {

	config.ParseEnvironment()

	CreateVM("rg-ea-net-prod", "test01", "azureuser", "P@ssw0rd!@#", "EastAsia", "MicrosoftSQLServer", "SQL2016SP2-WS2016", "Enterprise")

}

func getSubnet() {

	config.ParseEnvironment()
	subnet, err := GetVirtualNetworkSubnet(ctx, "vnet-prod", "snet-tier3-dbvm", "rg-ea-net-prod")

	if err != nil {
		log.Println("Could not get subnet")
		panic(err)
	}

	log.Println(*subnet.AddressPrefix)

}

func nicStuff() {
	nic, err := CreateNIC(ctx, "vnet-prod", "snet-tier3-dbvm", "", "", "ameer-vm-nic", "rg-ea-net-prod")

	if err != nil {
		panic(err)
	}

	log.Println(*nic.Name)
}
