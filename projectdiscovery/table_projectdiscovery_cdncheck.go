package projectdiscovery

import (
	"context"
	"fmt"
	"net"

	"github.com/projectdiscovery/cdncheck"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryCdncheck() *plugin.Table {
	return &plugin.Table{
		Name:        "projectdiscovery_cdncheck",
		Description: "A utility to detect various technology for a given IP address. <https://github.com/projectdiscovery/cdncheck>",
		List: &plugin.ListConfig{
			Hydrate: listCdnCheck,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: "Target IP to lookup."},
			{Name: "cdn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Cdn"), Description: "CDN information."},
			{Name: "cloud", Type: proto.ColumnType_STRING, Transform: transform.FromField("Cloud"), Description: "Cloud information."},
			{Name: "waf", Type: proto.ColumnType_STRING, Transform: transform.FromField("Waf"), Description: "WAF information."},
		},
	}
}

type cdnCheckRow struct {
	Cdn   string
	Cloud string
	Waf   string
}

func listCdnCheck(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	target := d.EqualsQualString("target")
	logger.Debug("target host", target)

	c := cdnCheckRow{}

	client := cdncheck.New()
	ip := net.ParseIP(target)

	if ip == nil {
		return nil, fmt.Errorf("invalid ip: %s", target)
	}

	// checks if an IP is contained in the cdn denylist
	matched, val, err := client.CheckCDN(ip)
	if err != nil {
		return nil, fmt.Errorf("CheckCDN() failed: %s", err)
	}

	if matched {
		c.Cdn = val
	}

	// checks if an IP is contained in the cloud denylist
	matched, val, err = client.CheckCloud(ip)
	if err != nil {
		return nil, fmt.Errorf("CheckCloud() failed: %s", err)
	}

	if matched {
		c.Cloud = val
	}

	// checks if an IP is contained in the waf denylist
	matched, val, err = client.CheckWAF(ip)
	if err != nil {
		return nil, fmt.Errorf("CheckWAF() failed: %s", err)
	}

	if matched {
		c.Waf = val
	}

	d.StreamListItem(ctx, c)

	return nil, nil
}
