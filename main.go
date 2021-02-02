package main

import (
	"context"
	"log"
)

var (
	ctx = context.Background()
)

func main() {

	// subnet, err := GetVirtualNetworkSubnet(ctx, "vnet-prod", "snet-tier3-dbvm", "rg-ea-vm-prod")

	// if err != nil {
	// 	log.Println("Could not get subnet")
	// 	panic(err)
	// }
	// log.Println(subnet.AddressPrefix)
	// os.Exit(0)

	nic, err := CreateNIC(ctx, "vnet-prod", "snet-tier3-dbvm", "", "", "ameer-vm-nic", "rg-ea-vm-prod")

	if err != nil {
		panic(err)
	}

	log.Println(nic.Name)
}
