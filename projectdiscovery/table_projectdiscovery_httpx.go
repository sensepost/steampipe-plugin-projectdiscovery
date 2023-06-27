package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/httpx/runner"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryHttpx() *plugin.Table {
	return &plugin.Table{
		Name:        `projectdiscovery_httpx`,
		Description: `httpx is a fast and multi-purpose HTTP toolkit. <https://github.com/projectdiscovery/httpx>`,
		List: &plugin.ListConfig{
			Hydrate: listHttpxScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: `The target`},
			{Name: "url", Type: proto.ColumnType_STRING, Description: `Target URL`},
			{Name: "method", Type: proto.ColumnType_STRING, Description: `HTTP method`},
			{Name: "host", Type: proto.ColumnType_IPADDR, Description: `Target host ip`},
			{Name: "path", Type: proto.ColumnType_STRING, Description: `Target path`},
			{Name: "port", Type: proto.ColumnType_INT, Description: `Target port`},
			{Name: "title", Type: proto.ColumnType_STRING, Description: `HTML title tag value`},
			{Name: "status_code", Type: proto.ColumnType_INT, Description: `HTTP response status code`},
			{Name: "content_length", Type: proto.ColumnType_INT, Description: `HTTP response content length`},
			{Name: "web_server", Type: proto.ColumnType_STRING, Description: `Remote webserver according to the Server header`},
			{Name: "technologies", Type: proto.ColumnType_JSON, Description: `HTTP technologies in use`},
			{Name: "response_time", Type: proto.ColumnType_STRING, Description: `HTTP response time`},
			{Name: "a", Type: proto.ColumnType_JSON, Description: `Target A record(s)`},
			{Name: "cname", Type: proto.ColumnType_STRING, Description: `Target CNAME record(s)`},
			{Name: "hashes", Type: proto.ColumnType_JSON, Description: `HTTP response status code`},
			{Name: "websocket", Type: proto.ColumnType_BOOL, Description: `True if the remote endpoint want to upgrade to a websocket`},
			{Name: "failed", Type: proto.ColumnType_BOOL, Description: `True if the probe failed`},
		},
	}
}

func listHttpxScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)

	host := d.EqualsQualString("target")
	logger.Debug("target host", host)

	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: goflags.StringSlice{host},
		ExtractTitle:    true,
		Hashes:          "sha1",
		TechDetect:      true,
	}

	options.OnResult = func(r runner.Result) {
		logger.Debug("httpx result", r)
		if r.Err != nil {
			return
		}

		d.StreamListItem(ctx, r)
	}

	if err := options.ValidateOptions(); err != nil {
		return nil, err
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		return nil, err
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()

	return nil, nil
}
