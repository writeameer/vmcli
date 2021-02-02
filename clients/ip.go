package clients

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/writeameer/vmcli/config"
	"github.com/writeameer/vmcli/iam"
)

// GetIPClient Gets the IP client
func GetIPClient() network.PublicIPAddressesClient {

	ipClient := network.NewPublicIPAddressesClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	ipClient.Authorizer = auth

	return ipClient
}
