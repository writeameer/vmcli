package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest"
	"github.com/writeameer/vmcli/clients"
	"github.com/writeameer/vmcli/config"
)

var (
	subscriptionID     = "367c3873-5d48-4c62-aec6-e23d28d181e4"
	armAuthorizer      autorest.Authorizer
	batchAuthorizer    autorest.Authorizer
	graphAuthorizer    autorest.Authorizer
	keyvaultAuthorizer autorest.Authorizer

	userAgent string
)

func main() {

	config.ParseEnvironment()

	fmt.Println("Hello World")
	resourceClient := clients.GetGroupsClient()

	result, err := resourceClient.List(context.Background(), "", nil)

	if err != nil {
		panic(err)
	}

	for _, item := range result.Values() {
		log.Println(*item.Name)
	}
}
