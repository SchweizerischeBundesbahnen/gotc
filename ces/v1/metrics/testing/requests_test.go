package testing

import (
	"testing"

	"github.com/SchweizerischeBundesbahnen/gotc/ces/v1/metrics"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestListMetrics(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListMetrics(t)

	allPages, err := metrics.List(fake.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	allMetrics, err := metrics.ExtractMetrics(allPages)
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expectedMetrics, allMetrics)
}

func TestGetMetric(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetMetrics(t)

	metric, err := metrics.Get(fake.ServiceClient(), nil).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedDatapoints, metric)
}
