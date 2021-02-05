package clients

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/writeameer/vmcli/config"
	"github.com/writeameer/vmcli/iam"
)

// GetVMClient gets the VM client
func GetVMClient() compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	vmClient.Authorizer = a
	return vmClient
}

// GetVMExtensionsClient gets the VM Extensions Client
func GetVMExtensionsClient() compute.VirtualMachineExtensionsClient {
	extClient := compute.NewVirtualMachineExtensionsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	extClient.Authorizer = a
	return extClient
}
