package flavors

import (
	"github.com/gophercloud/gophercloud"
)

// GetResult temporarily holds the response from a Get call.
type GetResult struct {
	gophercloud.Result
}

// Extract provides access to the individual Flavor returned by the Get function.
func (r GetResult) Extract() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"flavor"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

// Flavor records represent (virtual) hardware configurations for server resources in a region.
type Flavor struct {
	// The flavor's unique identifier.
	// Contains 0 if the ID is not an integer.
	ID int `json:"id"`

	// The RAM capacity for the flavor.
	RAM int `json:"ram"`

	// The Name field provides a human-readable moniker for the flavor.
	Name string `json:"name"`

	// The flavor's unique identifier as a string
	SpecCode string `json:"specCode"`
}
