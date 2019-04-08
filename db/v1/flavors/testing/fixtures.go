package testing

import (
	"fmt"
	"testing"

	"github.com/gophercloud/gophercloud/testhelper/fixture"
)

const flavor = `
{
	"id": %s,
	"name": "%s",
	"ram": %d,
	"specCode": "%s"
}
`

var (
	flavorID = "{flavorID}"
	_baseURL = "/flavors"
	resURL   = "/flavors/" + flavorID
)

var (
	flavor1 = fmt.Sprintf(flavor, "1", "m1.tiny", 512, "rds.mysql.m1.xlarge")

	getFlavorResp = fmt.Sprintf(`{"flavor": %s}`, flavor1)
)

func HandleGet(t *testing.T) {
	fixture.SetupHandler(t, resURL, "GET", "", getFlavorResp, 200)
}
