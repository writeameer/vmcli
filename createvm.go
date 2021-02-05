package main

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/writeameer/vmcli/clients"
	"github.com/writeameer/vmcli/vendor/github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

func blah() {
	// CreateVM(
	// 	"mygroup",
	// 	"test01",
	// 	"EastAsia",
	// 	"MicrosoftSQLServer",
	// 	"SQL2016SP2-WS2016",
	// 	"Enterprise",
	// )
}

// CreateVM creates a VM
func CreateVM(groupName string, vmName string, username string, password string, location string, publisher string, offer string, sku string) {

	nicClient := clients.GetNicClient()
	nic, err := nicClient.Get(ctx, groupName, "ameer-vm-nic", "")

	if err != nil {
		panic(err)
	}

	vmClient := clients.GetVMClient()

	//var vm compute.VirtualMachine
	vm := compute.VirtualMachine{
		Location: to.StringPtr(location),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.VirtualMachineSizeTypesBasicA0,
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &compute.ImageReference{
					Publisher: to.StringPtr(publisher),
					Offer:     to.StringPtr(offer),
					Sku:       to.StringPtr(sku),
					Version:   to.StringPtr("latest"),
				},
			},
			OsProfile: &compute.OSProfile{
				ComputerName:  to.StringPtr(vmName),
				AdminUsername: to.StringPtr(username),
				AdminPassword: to.StringPtr(password),
				LinuxConfiguration: &compute.LinuxConfiguration{
					SSH: &compute.SSHConfiguration{
						PublicKeys: &[]compute.SSHPublicKey{
							{
								Path: to.StringPtr(
									fmt.Sprintf("/home/%s/.ssh/authorized_keys",
										username)),
								KeyData: to.StringPtr("blah"),
							},
						},
					},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: nic.ID,
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							Primary: to.BoolPtr(true),
						},
					},
				},
			},
		},
	}

	future, err := vmClient.CreateOrUpdate(
		ctx,
		groupName,
		vmName,
		myvm,
	)

	if err != nil {
		return vm, fmt.Errorf("cannot create vm: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, vmClient.Client)
	if err != nil {
		return vm, fmt.Errorf("cannot get the vm create or update future response: %v", err)
	}

	// return future.Result(vmClient)
}
