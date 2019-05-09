package metrics

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"log"
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
	Unit      string
	Average   float64
	Timestamp int64
}

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a Get operation.
type GetResult struct {
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

type MetricPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks to see whether the collection is empty.
func (page MetricPage) IsEmpty() (bool, error) {
	metrics, err := ExtractMetrics(page)
	return len(metrics) == 0, err
}

// NextPageURL will retrieve the next page URL.
func (page MetricPage) NextPageURL() (string, error) {
	var s struct {
		Metadata *Metadata `json:"meta_data"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	log.Printf("%+v", s.Metadata)
	return "", nil
}

// ExtractMetrics will convert a generic pagination struct into a more
// relevant slice of Metric structs.
func ExtractMetrics(r pagination.Page) ([]Metric, error) {
	var s struct {
		Metrics  []Metric  `json:"metrics"`
		Metadata *Metadata `json:"meta_data"`
	}
	err := (r.(MetricPage)).ExtractInto(&s)
	return s.Metrics, err
}
