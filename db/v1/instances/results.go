package instances

import (
	"encoding/json"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/db/v1/users"
	"github.com/gophercloud/gophercloud/pagination"
	"strings"
	"time"
)

type Volume struct {
	Type string
	Size int
}

type Flavor struct {
	ID string
}

type OTCInstanceList struct {
	Instances []Instance
}

type DatastorePartial struct {
	Type    string
	Version string
}

type Instance struct {
	// Indicates the datetime that the instance was created
	Created time.Time `json:"-"`

	// Indicates the most recent datetime that the instance was updated.
	Updated time.Time `json:"-"`

	// Indicates the hardware flavor the instance uses.
	Flavor   Flavor
	Hostname string

	// Indicates the unique identifier for the instance resource.
	ID string

	// The human-readable name of the instance.
	Name string

	// The build status of the instance.
	Status string

	// OTC: Indicates the DB instance type, which can be master, slave, or readreplica.
	Type string

	// OTC: Indicates the region where the DB instance is deployed.
	Region string

	// OTC: Indicates the AZ ID.
	AvailabilityZone string

	// OTC: Indicates the VPC ID.
	VPC string

	// OTC: Indicates the nics information.
	NICs struct {
		SubnetID string
	}

	// OTC: Indicates the security group information.
	SecurityGroup struct {
		ID string
	}

	// OTC: Indicates the database port number.
	DBPort int

	BackupStrategy struct {
		StartTime backupStrategyTime
		KeepDays  int
	}

	// OTC: Returned only when you create primary/standby DB instances.
	SlaveID string

	// OTC: Indicates the primary/standby DB instance information. Returned only when you obtain a primary/standby DB
	HA struct {
		ReplicationMode string
	}

	// OTC: Returned only when you obtain the read replica information.
	ReplicaOf string

	// Information about the attached volume of the instance.
	Volume Volume

	// Indicates how the instance stores data.
	Datastore DatastorePartial
}

const RFC3339 = "2006-01-02T15:04:05+0000"

type JSONRFC3339 time.Time

func (jt *JSONRFC3339) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339(t)
	return nil
}

func (r *Instance) UnmarshalJSON(b []byte) error {
	type tmp Instance
	var s struct {
		tmp
		Created JSONRFC3339 `json:"created"`
		Updated JSONRFC3339 `json:"updated"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Instance(s.tmp)

	r.Created = time.Time(s.Created)
	r.Updated = time.Time(s.Updated)

	return nil
}

type backupStrategyTime struct {
	time.Time
}

func (ct *backupStrategyTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	timeLayout := "15:04:05"
	ct.Time, err = time.Parse(timeLayout, s)
	return
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a Create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ConfigurationResult represents the result of a AttachConfigurationGroup/DetachConfigurationGroup operation.
type ConfigurationResult struct {
	gophercloud.ErrResult
}

// Extract will extract an Instance from various result structs.
func (r commonResult) Extract() (*Instance, error) {
	var s struct {
		Instance *Instance `json:"instance"`
	}
	err := r.ExtractInto(&s)
	return s.Instance, err
}

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

// EnableRootUserResult represents the result of an operation to enable the root user.
type EnableRootUserResult struct {
	gophercloud.Result
}

// Extract will extract root user information from a UserRootResult.
func (r EnableRootUserResult) Extract() (*users.User, error) {
	var s struct {
		User *users.User `json:"user"`
	}
	err := r.ExtractInto(&s)
	return s.User, err
}

// ActionResult represents the result of action requests, such as: restarting
// an instance service, resizing its memory allocation, and resizing its
// attached volume size.
type ActionResult struct {
	gophercloud.ErrResult
}

// IsRootEnabledResult is the result of a call to IsRootEnabled. To see if
// root is enabled, call the type's Extract method.
type IsRootEnabledResult struct {
	gophercloud.Result
}

// Extract is used to extract the data from a IsRootEnabledResult.
func (r IsRootEnabledResult) Extract() (bool, error) {
	return r.Body.(map[string]interface{})["rootEnabled"] == true, r.Err
}
