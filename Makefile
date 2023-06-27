local:
	go build -o ~/.steampipe/plugins/local/projectdiscovery/projectdiscovery.plugin
install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/projectdiscovery@latest/steampipe-plugin-projectdiscovery.plugin
