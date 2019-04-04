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
