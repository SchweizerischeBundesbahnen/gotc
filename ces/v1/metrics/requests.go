package metrics

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)


// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
    ToListQuery() (string, error)
}

// ListOpts represents options used to list policies.
type ListOpts struct {
    Namespace string `q:"namespace"`
    Name string `q:"metric_name"`
    Dim string `q:"dim"`
    Start string `q:"start"`
    Limit int `q:"limit"`
    Order string `q:"order"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListQuery() (string, error) {
    q, err := gophercloud.BuildQueryString(opts)
    return q.String(), err
}


// List retrieves the status and information for all database instances.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
    url := listURL(client)
    if opts != nil {
        query, err := opts.ToListQuery()
        if err != nil {
            return pagination.Pager{Err: err}
        }
        url += query
    }

    return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
        p := MetricPage{pagination.MarkerPageBase{PageResult: r}}
        p.MarkerPageBase.Owner = p
        return p
    })
}

// Get retrieves the status and information for a specified database instance.
func Get(client *gophercloud.ServiceClient, values map[string]string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, values), &r.Body, nil)
	return
}
