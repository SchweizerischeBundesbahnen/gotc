package testing

import (
	"fmt"
	"testing"

	"github.com/SchweizerischeBundesbahnen/gotc/ces/v1/metrics"
	"github.com/gophercloud/gophercloud/testhelper/fixture"
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

var datapoint = `
{
    "average": 0,
    "timestamp": 1442341200000,
    "unit": "Count"
}
`

var (
	listURL = "/metrics"
	// TODO
	getURL = "/metric-data"
)

var (
	// TODO
	getMetricResp   = fmt.Sprintf(`{"datapoints": [%s]}`, datapoint)
	listMetricsResp = fmt.Sprintf(`{"metrics":[%s], "meta_data":{}}`, metric)
)

var expectedDatapoints = []metrics.Datapoint{
    metrics.Datapoint{
       Unit: "Count",
       Average: 0,
       Timestamp: 1442341200000,
    },
	// TODO
}

var expectedMetrics = []metrics.Metric{
    metrics.Metric{
        Namespace: "SYS.ECS",
        Dimensions: []metrics.Dimension{

        },
        Name: "cpu_util",
        Unit: "%",
    },
}

func HandleListMetrics(t *testing.T) {
	fixture.SetupHandler(t, listURL, "GET", "", listMetricsResp, 200)
}

func HandleGetMetrics(t *testing.T) {
	fixture.SetupHandler(t, getURL, "GET", "", getMetricResp, 200)
}
