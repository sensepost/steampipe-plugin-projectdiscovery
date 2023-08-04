package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/dnsx/libs/dnsx"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryDnsx() *plugin.Table {
	return &plugin.Table{
		Name:        "projectdiscovery_dnsx",
		Description: "dnsx is a fast and multi-purpose DNS toolkit. <https://github.com/projectdiscovery/dnsx>",
		List: &plugin.ListConfig{
			Hydrate: listDnsxScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "target", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: "Target to lookup."},
			{Name: "address", Type: proto.ColumnType_STRING, Description: "DNS A record response."},
		},
	}
}

type dnsxRow struct {
	Address string `json:"ip"`
}

func listDnsxScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	target := d.EqualsQualString("target")
	logger.Debug("target host", target)

	// Create DNS Resolver with default options
	dnsClient, err := dnsx.New(dnsx.DefaultOptions)
	if err != nil {
		logger.Error("projectdiscovery_dnsx.listDnsxScan", "connection_error", err)
		return nil, err
	}

	// DNS A question and returns corresponding IPs
	result, err := dnsClient.Lookup(target)
	if err != nil {
		logger.Warn("dnsx failed to lookup target", target, err)
		return nil, nil
	}

	for _, msg := range result {
		d.StreamListItem(ctx, dnsxRow{Address: msg})
	}

	return nil, nil
}
