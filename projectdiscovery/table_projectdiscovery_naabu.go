package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableProjectdiscoveryNaabu() *plugin.Table {
	return &plugin.Table{
		Name:        `projectdiscovery_naabu`,
		Description: `Naabu is a fast port scanner written in Go. <https://github.com/projectdiscovery/naabu>`,
		List: &plugin.ListConfig{
			Hydrate: listScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required},
			},
		},

		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Description: `Original target that was scanned`},
			{Name: "host", Type: proto.ColumnType_STRING, Description: `Resolved hostname of the target`},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: `Resolved IP address of the target`},
			{Name: "port", Type: proto.ColumnType_INT, Description: `A port that is open`},
		},
	}
}

type naabuRow struct {
	Target string `json:"target"`
	Host   string `json:"host"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
}

func listScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	host := d.EqualsQualString("target")
	logger.Debug("target host", host)

	// configure naabu
	naabuOptions := runner.Options{
		Host:     goflags.StringSlice{host},
		ScanType: "c", // just assume we'll never run as root, so always connect-scan it
	}

	// configure ports
	config := GetConfig(d.Connection)
	naabuOptions.TopPorts = *config.NaabuTopPorts
	logger.Debug("top ports", *config.NaabuTopPorts)

	// configure the callback to stream results
	naabuOptions.OnResult = func(hr *result.HostResult) {
		for _, port := range hr.Ports {
			logger.Debug("naabu result", port)
			d.StreamListItem(ctx, naabuRow{Target: host, Host: hr.Host, Ip: hr.IP, Port: port.Port})
		}
	}

	// prepare and run the scan with our options
	naabuRunner, err := runner.NewRunner(&naabuOptions)
	if err != nil {
		return nil, err
	}
	defer naabuRunner.Close()
	naabuRunner.RunEnumeration()

	return nil, nil
}
