package sources

import (
	"github.com/hashicorp/go-multierror"
	driversbase "github.com/jitsucom/jitsu/server/drivers/base"
)

type Unit struct {
	SourceType               string
	DriverPerCollection      map[string]driversbase.Driver
	DestinationIDs           []string
	PostHandleDestinationIDs []string

	hash uint64
}

//Close all drivers
func (u *Unit) Close() (multiErr error) {
	for _, driver := range u.DriverPerCollection {
		if err := driver.Close(); err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	return
}
