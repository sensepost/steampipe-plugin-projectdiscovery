# Projectdiscovery.io Plugin for Steampipe

[![Twitter](https://img.shields.io/badge/twitter-%40leonjza-blue.svg)](https://twitter.com/leonjza)

Use SQL to query Projectdiscovery.io tools for footprinting information.

- **[Get started →](https://hub.steampipe.io/plugins/sensepost/projectdiscovery)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/sensepost/projectdiscovery/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/sensepost/steampipe-plugin-projectdiscovery/issues)

## Quick start

### Install

Download and install the latest ProjectDiscovery plugin:

```bash
steampipe plugin install sensepost/projectdiscovery
```

Configure your account details in `~/.steampipe/config/projectdiscovery.spc`:

```hcl
connection "projectdiscovery" {
  plugin = "sensepost/projectdiscovery"

  # Naabu

  # Top ports to scan for naabu.
  # Can be one of: full, 100, 1000
  naabu_top_ports = "100"

  # Chaos

  # Project chaos API key. Sign up for an API key at: https://chaos.projectdiscovery.io/
  # chaos_api_key = "enpg9i8k4uxl0jtzoutym44cpm6rbxskr6fqoz11mxxpkiqtn4l7oju66rlqqz8j"

  # Cloudlist

  # Digital Ocean API key. Get an API key post authentication for a team by browsing to:
  # API (bottom left) -> Personal access tokens
  # A read-only key would suffice
  # cloudlist_do_token = "dop_v1_y0jzo0bp8wl7f3t0px74jea5hnxiicz1sl58z5mso6ep6a544v3mq1jp2qj4ed6a"
}
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
with target as (
  select domain from (
    values ('tesla.com'), ('reddit.com')
  ) t(domain)
), chaos as (
  select
    distinct concat(subdomain, '.', domain) as domain
  from
    projectdiscovery_chaos
  where
    domain in (
      select domain from target
    )
), subfinder as (
  select
    distinct
      host as domain
  from
    projectdiscovery_subfinder
  where
    target in (
      select domain from target
    )
)
select count(*) from (
  select
    domain from chaos
  union select
    domain from subfinder
  where domain not like '%*%'
) domains;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/sensepost/steampipe-plugin-projectdiscovery.git
cd steampipe-plugin-projectdiscovery
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make install
```

Configure the plugin:

```bash
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/projectdiscovery.spc
```

Try it!

```text
steampipe query
> .inspect projectdiscovery
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/sensepost/steampipe-plugin-projectdiscovery/blob/master/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [ProjectDiscovery Plugin](https://github.com/turbot/steampipe-plugin-projectdiscovery/labels/help%20wanted)
