# Projectdiscovery.io Plugin for Steampipe

[![Twitter](https://img.shields.io/badge/twitter-%40leonjza-blue.svg)](https://twitter.com/leonjza)

Use SQL to query Projectdiscovery.io tools for footprinting information

- Documentation: [Table definitions & examples](/docs/tables/)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io) (well, not yet, waiting to publish on the SteamPipe plugin hub. For now, use the instructions in the [developing](#developing) section below):

```shell
steampipe plugin install sensepost/projectdiscovery
```

Run a query:

```sql
with target as (
  select domain from (
    values ('tesla.com')
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
) domains
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

```bash
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
