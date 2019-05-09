package testing

import (
	"fmt"
	"testing"

	"github.com/SchweizerischeBundesbahnen/gotc/ces/v1/metrics"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
	"net/http"
)

// escape percent character
var metric = `
{
    "namespace": "SYS.ECS",
    "dimensions": [
        {
            "name": "instance_id",
            "value": "d9112af5-6913-4f3b-bd0a-3f96711e004d"
        }
    ],
    "metric_name": "cpu_util",
    "unit": "%%"
}
`

var metadata = `
{
    "count": 1,
    "marker": "SYS.ECS.cpu_util.instance_id:d9112af5-6913-4f3b-bd0a-3f96711e004d",
    "total": 1
}
`

var datapoint = `
{
    "average": 0,
    "timestamp": 1442341200000,
    "unit": "Count"
}
`

var (
	listURL = "/metrics"
	getURL  = "/metric-data"
)

var (
	getMetricResp   = fmt.Sprintf(`{"datapoints": [%s], "metric_name":"cpu_util"}`, datapoint)
	listMetricsResp = fmt.Sprintf(`{"metrics":[%s], "meta_data":%s}`, metric, metadata)
)

var metricsListBody1 = `
{
    "metrics": [
        {
            "namespace": "SYS.ECS",
            "dimensions": [
                {
                    "name": "instance_id",
                    "value": "id1"
                }
            ],
            "metric_name": "cpu_util",
            "unit": "%%"
        }
    ],
    "meta_data": {
        "count": 1,
        "marker": "metricsListMarker1",
        "total": 3
    }
}
`
var metricsListBody2 = `
{
    "metrics": [
        {
            "namespace": "SYS.ECS",
            "dimensions": [
                {
                    "name": "instance_id",
                    "value": "id2"
                }
            ],
            "metric_name": "cpu_util",
            "unit": "%%"
        }
    ],
    "meta_data": {
        "count": 1,
        "marker": "metricsListMarker2",
        "total": 3
    }
}
`

var expectedDatapoints = []metrics.Datapoint{
	{
		Unit:      "Count",
		Average:   0,
		Timestamp: 1442341200000,
	},
}

var expectedMetrics = []metrics.Metric{
	{
		Namespace: "SYS.ECS",
		Dimensions: []metrics.Dimension{
			{
				Name:  "instance_id",
				Value: "id1",
			},
		},
		Name: "cpu_util",
		Unit: "%",
	},
	{
		Namespace: "SYS.ECS",
		Dimensions: []metrics.Dimension{
			{
				Name:  "instance_id",
				Value: "id2",
			},
		},
		Name: "cpu_util",
		Unit: "%",
	},
}

func HandleListMetrics(t *testing.T) {
	th.Mux.HandleFunc(listURL, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		r.ParseForm()
		marker := r.Form.Get("start")
		switch marker {
		case "":
			fmt.Fprintf(w, metricsListBody1)
		case "metricsListMarker1":
			fmt.Fprintf(w, metricsListBody2)
		case "metricsListMarker2":
			fmt.Fprintf(w, `{"metrics":[], "meta_data": {}}`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

func HandleGetMetrics(t *testing.T) {
	fixture.SetupHandler(t, getURL, "GET", "", getMetricResp, 200)
}
