// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package clients

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/writeameer/vmcli/config"
	"github.com/writeameer/vmcli/iam"
)

// GetGroupsClient Get Resource Groups Client
func GetGroupsClient() resources.GroupsClient {
	config.ParseEnvironment()
	groupsClient := resources.NewGroupsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	groupsClient.Authorizer = a
	return groupsClient
}
