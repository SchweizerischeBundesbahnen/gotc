package flavors

import (
	"github.com/gophercloud/gophercloud"
)

// Get will retrieve information for a specified hardware flavor.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
