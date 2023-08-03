package projectdiscovery

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-projectdiscovery",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		// todo: uncover
		TableMap: map[string]*plugin.Table{
			"projectdiscovery_asnmap":    tableProjectdiscoveryAsnmap(),
			"projectdiscovery_cdncheck":  tableProjectdiscoveryCdncheck(),
			"projectdiscovery_chaos":     tableProjectdiscoveryChaos(),
			"projectdiscovery_cloudlist": tableProjectdiscoveryCloudlist(),
			"projectdiscovery_dnsx":      tableProjectdiscoveryDnsx(),
			"projectdiscovery_httpx":     tableProjectdiscoveryHttpx(),
			"projectdiscovery_naabu":     tableProjectdiscoveryNaabu(),
			"projectdiscovery_subfinder": tableProjectdiscoverySubfinder(),
			"projectdiscovery_tlsx":      tableProjectdiscoveryTlsx(),
		},
	}
	return p
}
