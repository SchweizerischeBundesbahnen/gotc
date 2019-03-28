package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
    "time"
    "strings"
)

type OTCInstanceList struct {
	Instances []Instance
}

type Instance struct {
	// Indicates the datetime that the instance was created
	Created time.Time `json:"-"`

	// Indicates the most recent datetime that the instance was updated.
	Updated time.Time `json:"-"`

	// Indicates the hardware flavor the instance uses.
	Flavor struct {
		ID string
	}

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
	Volume struct {
		Type string
		Size int
	}

	// Indicates how the instance stores data.
	Datastore struct {
		Type    string
		Version string
	}
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