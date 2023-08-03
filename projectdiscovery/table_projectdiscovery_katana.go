package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/katana/pkg/engine/standard"
	"github.com/projectdiscovery/katana/pkg/output"
	"github.com/projectdiscovery/katana/pkg/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryKatana() *plugin.Table {
	return &plugin.Table{
		Name:        "projectdiscovery_katana",
		Description: "A next-generation crawling and spidering framework. <https://github.com/projectdiscovery/katana>",
		List: &plugin.ListConfig{
			Hydrate: listKatanaScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required, CacheMatch: "exact"},
				{Name: `depth`, Require: plugin.Optional, CacheMatch: "exact"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: "Target to lookup."},
			{Name: "depth", Type: proto.ColumnType_INT, Transform: transform.FromQual("depth"), Description: "Depth to scan."},
			{Name: "request", Type: proto.ColumnType_JSON, Description: "The HTTP request."},
			{Name: "response", Type: proto.ColumnType_JSON, Description: "The HTTP response."},
			{Name: "error", Type: proto.ColumnType_STRING, Description: "An error, if any."},
		},
	}
}

func listKatanaScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	target := d.EqualsQualString("target")
	logger.Debug("target host", target)

	depth := d.EqualsQuals["depth"].GetInt64Value()
	if depth == 0 {
		depth = 1
	}
	logger.Debug("depth for crawl", depth)

	options := &types.Options{
		MaxDepth:     int(depth),
		FieldScope:   "rdn", // root domain name?
		BodyReadSize: 2 * 1024 * 1024,
		RateLimit:    150,
		Verbose:      true,
		Strategy:     "depth-first",
	}

	options.OnResult = func(r output.Result) {
		logger.Debug("crawl result", r)
		d.StreamListItem(ctx, r)
	}

	crawlerOptions, err := types.NewCrawlerOptions(options)
	if err != nil {
		logger.Error("projectdiscovery_katana.listKatanaScan", "new_crawler_options_api_err", err)
		return nil, err
	}
	defer crawlerOptions.Close()

	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		logger.Error("projectdiscovery_katana.listKatanaScan", "new_crawler_instance_api_err", err)
		return nil, err
	}
	defer crawler.Close()

	if err := crawler.Crawl(target); err != nil {
		return nil, err
	}

	return nil, nil
}
