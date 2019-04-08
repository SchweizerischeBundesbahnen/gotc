package testing

import (
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/lbischof/gotc/db/v1/flavors"
)

func TestGetFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGet(t)

	actual, err := flavors.Get(fake.ServiceClient(), flavorID).Extract()
	th.AssertNoErr(t, err)

	expected := &flavors.Flavor{
		ID:       1,
		Name:     "m1.tiny",
		RAM:      512,
		SpecCode: "rds.mysql.m1.xlarge",
	}

	th.AssertDeepEquals(t, expected, actual)
}
