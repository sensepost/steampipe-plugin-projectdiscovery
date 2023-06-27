package projectdiscovery

import (
	"context"
	"errors"

	"github.com/projectdiscovery/chaos-client/pkg/chaos"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableProjectdiscoveryChaos() *plugin.Table {
	return &plugin.Table{
		Name:        `projectdiscovery_chaos`,
		Description: `Choas is an Internet-wide assets data project. <https://chaos.projectdiscovery.io/>`,
		List: &plugin.ListConfig{
			Hydrate: listChaos,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `domain`, Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: `Domain under query`},
			{Name: "subdomain", Type: proto.ColumnType_STRING, Description: `A subdomain`},
		},
	}
}

type chaosRow struct {
	Domain    string `json:"domain"`
	Subdomain string `json:"subomain"`
}

func listChaos(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	domain := d.EqualsQualString("domain")

	config := GetConfig(d.Connection)
	if *config.ChaosAPIKey == "" {
		return nil, errors.New("this table requires a configured chaos api key")
	}

	chaosClient := chaos.New(*config.ChaosAPIKey)

	for entry := range chaosClient.GetSubdomains(&chaos.SubdomainsRequest{Domain: domain}) {
		if entry.Error != nil {
			return nil, entry.Error
		}

		d.StreamListItem(ctx, chaosRow{Domain: domain, Subdomain: entry.Subdomain})
	}

	return nil, nil
}
