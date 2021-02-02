package config

import (
	"log"
	"os"
)

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {

	// AZURE_GROUP_NAME and `config.GroupName()` are deprecated.
	// Use AZURE_BASE_GROUP_NAME and `config.GenerateGroupName()` instead.
	// groupName = os.Getenv("AZURE_GROUP_NAME")
	// baseGroupName = os.Getenv("AZURE_BASE_GROUP_NAME")
	locationDefault = os.Getenv("AZURE_LOCATION_DEFAULT")

	// these must be provided by environment
	// clientID
	clientID = os.Getenv("AZURE_CLIENT_ID")
	if clientID == "" {
		log.Fatal("Environment variable was empty: AZURE_CLIENT_ID")
	}
	// clientSecret
	clientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("Environment variable was empty: AZURE_CLIENT_SECRET")
	}

	// tenantID (AAD)
	tenantID = os.Getenv("AZURE_TENANT_ID")
	if tenantID == "" {
		log.Fatal("Environment variable was empty: AZURE_TENANT_ID")
	}

	// subscriptionID (ARM)
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		log.Fatal("Environment variable was empty: AZURE_SUBSCRIPTION_ID")
	}

	return nil
}
