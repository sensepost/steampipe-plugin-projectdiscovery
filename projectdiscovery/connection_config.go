package projectdiscovery

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type projectdiscoveryConfig struct {
	NaabuTopPorts *string `cty:"naabu_top_ports"`
	ChaosAPIKey   *string `cty:"chaos_api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"naabu_top_ports": {Type: schema.TypeString},
	"chaos_api_key":   {Type: schema.TypeString},
}

func ConfigInstance() interface{} {
	return &projectdiscoveryConfig{}
}

// GetConfig will retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) projectdiscoveryConfig {
	if connection == nil || connection.Config == nil {
		return projectdiscoveryConfig{}
	}

	config, _ := connection.Config.(projectdiscoveryConfig)
	return config
}
