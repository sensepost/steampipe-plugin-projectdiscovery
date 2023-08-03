package projectdiscovery

import (
	"context"

	"github.com/projectdiscovery/tlsx/pkg/tlsx"
	"github.com/projectdiscovery/tlsx/pkg/tlsx/clients"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableProjectdiscoveryTlsx() *plugin.Table {
	return &plugin.Table{
		Name:        "projectdiscovery_tlsx",
		Description: "Fast and configurable TLS grabber. <https://github.com/projectdiscovery/tlsx>",
		List: &plugin.ListConfig{
			Hydrate: listTlsxScan,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: `target`, Require: plugin.Required},
				{Name: `port`, Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "target", Type: proto.ColumnType_STRING, Transform: transform.FromQual("target"), Description: "Original target that was scanned."},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the target was probed."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "The IP the request was made to."},
			{Name: "port", Type: proto.ColumnType_STRING, Description: "The port the request was made to."},
			{Name: "client", Type: proto.ColumnType_STRING, Transform: transform.FromField("TLSConnection"), Description: "The TLS client used."},
			{Name: "success", Type: proto.ColumnType_BOOL, Transform: transform.FromField("ProbeStatus"), Description: "False if the probe failed."},
			{Name: "error", Type: proto.ColumnType_STRING, Description: "The error that occured, if any."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "The TLS version the server responded with."},
			{Name: "cipher", Type: proto.ColumnType_STRING, Description: "The cipher used for the probe."},
			{Name: "certificate_response", Type: proto.ColumnType_JSON, Description: "The leaf certificate presented by the server."},
			{Name: "chain", Type: proto.ColumnType_JSON, Description: "The chain of certificates."},
			{Name: "jarm_hash", Type: proto.ColumnType_STRING, Description: "The calculated jarm hash."},
			{Name: "sni", Type: proto.ColumnType_STRING, Transform: transform.FromField("ServerName"), Description: "Server Name Indicator."},
		},
	}
}

func listTlsxScan(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	host := d.EqualsQualString("target")
	logger.Debug("target host", host)

	opts := &clients.Options{
		TLSVersion:    true,
		TLSChain:      true,
		Retries:       3,
		ReversePtrSNI: true,
		Jarm:          true,
	}

	// todo: handle port as an int.
	// special care should be taken here as tlsx expects a string :(
	port := d.EqualsQualString("port")
	if port == "" {
		port = "443"
	}

	service, err := tlsx.New(opts)
	if err != nil {
		return nil, err
	}

	// connect to any host either with hostname or ip
	// service.Connect(hostname, ip , port string)
	resp, err := service.Connect(host, "", port)
	if err != nil {
		logger.Warn("failed to connect to host and port", host, port, err)
		// dont bubble up the error. sometimes there simply isnt tls on the other side.
		return nil, nil
	}

	d.StreamListItem(ctx, resp)

	return nil, nil
}
