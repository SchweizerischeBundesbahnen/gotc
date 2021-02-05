package instances

import (
	"strings"

	"github.com/gophercloud/gophercloud"
)

func baseURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("instances")
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id)
}

func userRootURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "root")
}

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "action")
}

func tagURL(c *gophercloud.ServiceClient, id string) string {
	endpoint := c.Endpoint
	// map: https://rds.eu-ch.o13bb.otc.t-systems.com/rds/v1/$(tenant_id)s
	// to: https://rds.eu-ch.o13bb.otc.t-systems.com/v1/$(tenant_id)s/rds
	c.Endpoint = strings.Replace(endpoint, "rds/", "", 1)

	return c.ServiceURL("rds", id, "tags")
}
