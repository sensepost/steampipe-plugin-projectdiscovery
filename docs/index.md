---
organization: Sensepost
category: ["osint"]
display_name: "Project Discovery"
short_name: "projectdiscovery"
description: "Steampipe plugin for interacting with projectdiscovery.io toolsets."
og_description: "Query ProjectDiscovery with SQL! Open source CLI. No DB required."
---

# Project Dicovery + Steampipe

[Project Discovery](https://projectdiscovery.io/#/) is an open-source software company that builds tools to detect and remediate vulnerabilities across your modern tech stack.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example, to get ASN related information:

```sql
select
  asn,
  country,
  first_ip,
  last_ip
from
  projectdiscovery_asnmap
where
  target = 'google' limit 5
```

```text
+-------+---------+--------------+----------------+
| asn   | country | first_ip     | last_ip        |
+-------+---------+--------------+----------------+
| 15169 | US      | 34.0.128.0   | 34.0.225.255   |
| 15169 | US      | 34.160.0.0   | 34.160.255.255 |
| 15169 | US      | 35.219.128.0 | 35.219.191.255 |
| 15169 | US      | 35.230.232.0 | 35.230.239.255 |
| 15169 | US      | 72.14.192.0  | 72.14.255.255  |
+-------+---------+--------------+----------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/sensepost/projectdiscovery/docs/tables)**

## Get Started

### Install

Download and install the latest Project Discovery plugin:

```bash
steampipe plugin install sensepost/projectdiscovery
```

### Configuration

Installing the latest projectdiscovery plugin will create a config file (`~/.steampipe/config/projectdiscovery.spc`). Some services require credentials such as the [choas dataset](https://chaos.projectdiscovery.io/#/) for the chaos table and various cloud providers for the cloudlist table.

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

- `chaos_api_key` - The API key to access the chaos dataset.
- `cloudlist_do_token` - The Digital Ocean API key.

## Get involved

- Open source: <https://github.com/sensepost/steampipe-plugin-projectdiscovery>
- Community: [Slack Channel](https://steampipe.io/community/join)
