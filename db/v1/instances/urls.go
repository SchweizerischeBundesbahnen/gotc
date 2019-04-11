package instances

import (
	"github.com/gophercloud/gophercloud"
	"log"
	"strings"
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
	log.Println(endpoint)
	// map: https://rds.eu-ch.o13bb.otc.t-systems.com/rds/v1/$(tenant_id)s
	// to: https://rds.eu-ch.o13bb.otc.t-systems.com/v1/$(tenant_id)s/rds
	c.Endpoint = strings.Replace(endpoint, "rds/", "", 1)
	log.Println(c.Endpoint)

	return c.ServiceURL("rds", id, "tags")
}
