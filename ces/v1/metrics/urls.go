package metrics

import (
    "github.com/gophercloud/gophercloud"
    "net/url"
)

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}

func getURL(c *gophercloud.ServiceClient, values map[string]string) string {
    query := url.Values{}
    for k, v := range values {
        query.Add(k, v)
    }
	return c.ServiceURL("metric-data?"+ query.Encode())
}
