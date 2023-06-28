package projectdiscovery

import (
	"context"
	"fmt"

	"github.com/projectdiscovery/cloudlist/pkg/inventory"
	"github.com/projectdiscovery/cloudlist/pkg/schema"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryCLoudlist() *plugin.Table {
	return &plugin.Table{
		Name:        `projectdiscovery_cloudlist`,
		Description: `Cloudlist is a tool for listing Assets from multiple Cloud Providers. <https://github.com/projectdiscovery/cloudlist>`,
		List: &plugin.ListConfig{
			Hydrate: listCloudlistScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `provider`, Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "provider", Type: proto.ColumnType_STRING, Transform: transform.FromQual("provider"), Description: `Target provider under query`},
			{Name: "id", Type: proto.ColumnType_BOOL, Description: `The id name of the resource provider`},
			{Name: "public", Type: proto.ColumnType_BOOL, Description: `True if the resource is public`},
			{Name: "public_ipv4", Type: proto.ColumnType_STRING, Description: `The public ipv4 address of the resource`},
			{Name: "private_ipv4", Type: proto.ColumnType_STRING, Description: `The private ipv4 address of the resource`},
			{Name: "dns_name", Type: proto.ColumnType_STRING, Description: `The DNS name of the resource`},
		},
	}
}

type cloudListScanRow struct {
	Public      bool
	Provider    string
	Id          string
	PublicIpv4  string
	PrivateIpv4 string
	DnsName     string
}

func listCloudlistScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	provider := d.EqualsQualString("provider")
	logger.Debug("cloudlist provider", provider)

	opts, err := cloudListGetProviderConfig(d.Connection, provider)
	if err != nil {
		return nil, err
	}

	inventory, err := inventory.New(opts)
	if err != nil {
		return nil, err
	}

	for _, provider := range inventory.Providers {
		resources, err := provider.Resources(context.Background())
		if err != nil {
			return nil, err
		}
		for _, resource := range resources.Items {
			d.StreamListItem(ctx, cloudListScanRow{
				Public:      resource.Public,
				Provider:    resource.Provider,
				Id:          resource.ID,
				PublicIpv4:  resource.PublicIPv4,
				PrivateIpv4: resource.PrivateIpv4,
				DnsName:     resource.DNSName,
			})
		}
	}

	return nil, nil
}

// cloudListGetProviderConfig grabs the configuration for a provider from the steampipe config
func cloudListGetProviderConfig(conn *plugin.Connection, provider string) (schema.Options, error) {
	config := GetConfig(conn)

	// todo: implement more providers.
	// the most work would be to map config -> schema.OptioonBlock's

	switch provider {
	case "do":
		if config.CloudListDoToken == nil || *config.CloudListDoToken == "" {
			return nil, fmt.Errorf("digital ocean token not configured")
		}
		return schema.Options{
			schema.OptionBlock{
				"provider":           "do",
				"digitalocean_token": *config.CloudListDoToken,
			},
		}, nil
	case "gcp":
	case "scw":
	case "azure":
	case "cloudflare":
	case "heroku":
	case "linode":
	case "fastly":
	case "alibaba":
	case "namecheap":
	case "terraform":
	case "consul":
	case "nomad":
	case "hetzner":
	case "openstack":
	case "kubernetes":
	case "aws":
	default:
		return nil, fmt.Errorf("invalid provider name, or provider not implemented yet: %s", provider)
	}

	return nil, nil
}
