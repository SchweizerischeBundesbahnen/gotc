package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)


// List retrieves the status and information for all database instances.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, baseURL(client), func(r pagination.PageResult) pagination.Page {
		return InstancePage{pagination.LinkedPageBase{PageResult: r}}
	})
}
