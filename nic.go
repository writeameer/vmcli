package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/writeameer/vmcli/clients"
	"github.com/writeameer/vmcli/config"
)

// GetVirtualNetworkSubnet returns an existing subnet from a virtual network
func GetVirtualNetworkSubnet(ctx context.Context, vnetName string, subnetName string, groupName string) (network.Subnet, error) {
	config.ParseEnvironment()
	subnetsClient := clients.GetSubnetsClient()
	return subnetsClient.Get(ctx, groupName, vnetName, subnetName, "")
}

// CreateNIC Creates a NIC
func CreateNIC(ctx context.Context, vnetName string, subnetName string, nsgName string, ipName string, nicName string, groupName string) (nic network.Interface, err error) {

	subnet, err := GetVirtualNetworkSubnet(ctx, vnetName, subnetName, groupName)

	//GetVirtualNetworkSubnet(ctx, "vnet-prod", "snet-tier3-dbvm", "rg-ea-net-prod")

	if err != nil {
		log.Println("failed to get subnet:")
		log.Println("vnet name: " + vnetName)
		log.Println("subnet name: " + subnetName)
		log.Println("groupName name: " + groupName)
		log.Fatal(err)
	}

	nicParams := network.Interface{
		Name:     to.StringPtr(nicName),
		Location: to.StringPtr(config.DefaultLocation()),
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			IPConfigurations: &[]network.InterfaceIPConfiguration{
				{
					Name: to.StringPtr("ipConfig1"),
					InterfaceIPConfigurationPropertiesFormat: &network.InterfaceIPConfigurationPropertiesFormat{
						Subnet:                    &subnet,
						PrivateIPAllocationMethod: network.Dynamic,
						PublicIPAddress:           nil,
					},
				},
			},
		},
	}

	nicClient := clients.GetNicClient()
	future, err := nicClient.CreateOrUpdate(ctx, groupName, nicName, nicParams)
	if err != nil {
		return nic, fmt.Errorf("cannot create nic: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, nicClient.Client)
	if err != nil {
		return nic, fmt.Errorf("cannot get nic create or update future response: %v", err)
	}

	return future.Result(nicClient)
}
