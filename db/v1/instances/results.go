package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// InstancePage represents a single page of a paginated instance collection.
type InstancePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks to see whether the collection is empty.
func (page InstancePage) IsEmpty() (bool, error) {
	instances, err := ExtractInstances(page)
	return len(instances) == 0, err
}

// NextPageURL will retrieve the next page URL.
func (page InstancePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"instances_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// ExtractInstances will convert a generic pagination struct into a more
// relevant slice of Instance structs.
func ExtractInstances(r pagination.Page) ([]Instance, error) {
	var s struct {
		Instances []Instance `json:"instances"`
	}
	err := (r.(InstancePage)).ExtractInto(&s)
	return s.Instances, err
}
