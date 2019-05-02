package metrics

import (
	"github.com/gophercloud/gophercloud"
)

// List retrieves the status and information for all database instances.
func List(client *gophercloud.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(listURL(client), &r.Body, nil)
	return
}

type GetRequest struct {
    Namespace string
    Name string `json:"metric_name"`
}
// Get retrieves the status and information for a specified database instance.
func Get(client *gophercloud.ServiceClient, values map[string]string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, values), &r.Body, nil)
	return
}
