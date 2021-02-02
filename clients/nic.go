package clients

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/writeameer/vmcli/config"
	"github.com/writeameer/vmcli/iam"
)

// GetNicClient gets the NIC client
func GetNicClient() network.InterfacesClient {
	nicClient := network.NewInterfacesClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	nicClient.Authorizer = auth
	return nicClient
}
