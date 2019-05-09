package metrics

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List retrieves the status and information for all database instances.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return MetricPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves the status and information for a specified database instance.
func Get(client *gophercloud.ServiceClient, values map[string]string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, values), &r.Body, nil)
	return
}
