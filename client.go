package gotc

import (
	"github.com/gophercloud/gophercloud"
)

func NewDBV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "rdsv1")
	if err != nil {
		return nil, err
	}
	sc.MoreHeaders = map[string]string{
		"X-Language":   "en-us",
		"Content-Type": "application/json",
	}
	return sc, nil
}

func NewCloudeyeV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "cesv1")
	if err != nil {
		return nil, err
	}
	sc.MoreHeaders = map[string]string{
		"X-Language":   "en-us",
		"Content-Type": "application/json",
	}
	return sc, nil
}

// NewRDSV3 creates a ServiceClient that may be used with the v3 rds
// package.
func NewRDSV3(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "rdsv3")
	return sc, err
}

func initClientOpts(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	sc := new(gophercloud.ServiceClient)
	eo.ApplyDefaults(clientType)
	url, err := client.EndpointLocator(eo)
	if err != nil {
		return sc, err
	}
	sc.ProviderClient = client
	sc.Endpoint = url
	sc.Type = clientType
	return sc, nil
}
