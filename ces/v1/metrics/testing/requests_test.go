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

	actual, err := metrics.List(fake.ServiceClient()).Extract()

	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, expectedMetrics, actual)
}

//func TestGetMetric(t *testing.T) {
//	th.SetupHTTP()
//	defer th.TeardownHTTP()
//	HandleGetMetrics(t)
//
//	metric, err := metrics.Get(fake.ServiceClient(), instanceID).Extract()
//
//	th.AssertNoErr(t, err)
//	th.AssertDeepEquals(t, &expectedMetric, metric)
//}
