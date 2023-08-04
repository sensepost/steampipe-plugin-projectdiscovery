package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoverySubfinder() *plugin.Table {
	return &plugin.Table{
		Name:        "projectdiscovery_subfinder",
		Description: "Fast passive subdomain enumeration tool. <https://github.com/projectdiscovery/subfinder>",
		List: &plugin.ListConfig{
			Hydrate: listSubfinderScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "target", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: "The target domain."},
			{Name: "host", Type: proto.ColumnType_STRING, Description: "Host of the discovered domain."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "The data source."},
		},
	}
}

type subfinderRow struct {
	Domain string
	Host   string
	Source string
}

func listSubfinderScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	target := d.EqualsQualString("target")
	logger.Debug("target host", target)

	opts := &runner.Options{
		// All:                true,
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 5,
		RemoveWildcard:     true,
		// todo: figure how to make this file requirement something we can feed in from
		// this plugins configuration
		// ProviderConfig: "your_provider_config.yaml",
	}

	opts.ResultCallback = func(result *resolve.HostEntry) {
		logger.Debug("got subfinder result", result)
		d.StreamListItem(ctx, subfinderRow{
			Domain: result.Domain,
			Host:   result.Host,
			Source: result.Source,
		})
	}

	subfinder, err := runner.NewRunner(opts)
	if err != nil {
		logger.Error("projectdiscovery_subfinder.listSubfinderScan", "new_runner_api_error", err)
		return nil, err
	}

	if err = subfinder.EnumerateSingleDomainWithCtx(ctx, target, nil); err != nil {
		// just log the error
		logger.Warn("subfinder enumeration had an error", err)
	}

	return nil, nil
}
