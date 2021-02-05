package clients

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/writeameer/vmcli/config"
	"github.com/writeameer/vmcli/iam"
)

func GetSubnetsClient() network.SubnetsClient {
	subnetsClient := network.NewSubnetsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	subnetsClient.Authorizer = auth
	return subnetsClient
}
