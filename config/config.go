// Package config manages loading configuration from environment and command-line params
package config

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
)

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.
	clientID               string
	clientSecret           string
	tenantID               string
	subscriptionID         string
	locationDefault        string
	authorizationServerURL string
	cloudName              string = "AzurePublicCloud"
	useDeviceFlow                 = false
	keepResources                 = false
	groupName              string // deprecated, use baseGroupName instead
	environment            *azure.Environment
)

// ClientID is the OAuth client ID.
func ClientID() string {
	return clientID
}

// ClientSecret is the OAuth client secret.
func ClientSecret() string {
	return clientSecret
}

// TenantID is the AAD tenant to which this client belongs.
func TenantID() string {
	return tenantID
}

// SubscriptionID is a target subscription for Azure resources.
func SubscriptionID() string {
	return subscriptionID
}

// DefaultLocation returns the default location wherein to create new resources.
// Some resource types are not available in all locations so another location might need
// to be chosen.
func DefaultLocation() string {
	return locationDefault
}

// AuthorizationServerURL is the OAuth authorization server URL.
// Q: Can this be gotten from the `azure.Environment` in `Environment()`?
func AuthorizationServerURL() string {
	return authorizationServerURL
}

// Environment returns an `azure.Environment{...}` for the current cloud.
func Environment() *azure.Environment {
	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		// TODO: move to initialization of var
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}
