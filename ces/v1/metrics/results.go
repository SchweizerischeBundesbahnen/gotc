package metrics

import (
	"github.com/gophercloud/gophercloud"
)

type Metadata struct {
	Count  int
	Marker string
	Total  int
}

type Dimension struct {
	Name  string
	Value string
}

type Metric struct {
	Namespace  string
	Dimensions []Dimension
	Name       string `json:"metric_name"`
	Unit       string
}

type Datapoint struct {
    Unit string
    Average int
}

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	commonResult
}

type ListResult struct {
	commonResult
}

// Extract will extract an Instance from various result structs.
func (r GetResult) Extract() ([]Datapoint, error) {
	var s struct {
		Datapoints []Datapoint `json:"datapoints"`
	}
	err := r.ExtractInto(&s)
	return s.Datapoints, err
}

// WARNING: this isn't paged. Will cut off everything after 1000 results!
func (r ListResult) Extract() ([]Metric, error) {
	var s struct {
		Metrics  []Metric  `json:"metrics"`
		Metadata *Metadata `json:"meta_data"`
	}
	err := r.ExtractInto(&s)
	return s.Metrics, err
}
