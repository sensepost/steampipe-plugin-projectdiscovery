package main

import (
	"github.com/sensepost/steampipe-plugin-projectdiscovery/projectdiscovery"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: projectdiscovery.Plugin})
}
