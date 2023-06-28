package projectdiscovery

import (
	"context"

	asnmap "github.com/projectdiscovery/asnmap/libs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryAsnmap() *plugin.Table {
	return &plugin.Table{
		Name:        `projectdiscovery_asnmap`,
		Description: `Library for quickly mapping organization network ranges using ASN information. <https://github.com/projectdiscovery/asnmap>`,
		List: &plugin.ListConfig{
			Hydrate: listAsnmap,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: `The ASN, IP or Org name to lookup`},
			{Name: "asn", Type: proto.ColumnType_INT, Transform: transform.FromField("ASN"), Description: `The ASN`},
			{Name: "country", Type: proto.ColumnType_STRING, Transform: transform.FromField("Country"), Description: `The country`},
			{Name: "org", Type: proto.ColumnType_STRING, Transform: transform.FromField("Org"), Description: `The organisation`},
			{Name: "first_ip", Type: proto.ColumnType_INET, Transform: transform.FromField("FirstIp"), Description: `First IP for the ASN`},
			{Name: "last_ip", Type: proto.ColumnType_INET, Transform: transform.FromField("LastIp"), Description: `First IP for the ASN`},
		},
	}
}

func listAsnmap(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	target := d.EqualsQualString("target")
	logger.Debug("target", target)

	client, err := asnmap.NewClient()
	if err != nil {
		return nil, err
	}

	results, err := client.GetData(target)
	if err != nil {
		logger.Warn("asnmap failed to get data for domain", target)
	}
	logger.Debug("asnmap results", results)

	for _, asn := range results {
		d.StreamListItem(ctx, asn)
	}

	return nil, nil
}
