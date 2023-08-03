# Table: projectdiscovery_chaos

[chaos-client](https://github.com/projectdiscovery/chaos-client) is a Go client to communicate with Chaos DB API.

**Note:** This service requires a valid API key to access the Chaos dataset. Request a key [here](https://chaos.projectdiscovery.io/).

## Examples

### Get subdomains for a domain

```sql
select
  subdomain
from
  projectdiscovery_chaos
where
  domain = 'google.com';
```

### Get FQDN's for a domain, excluding wildcards

```sql
select distinct
  concat(subdomain, '.', domain)
from
  projectdiscovery_chaos
where
  domain = 'google.com'
  and subdomain not like '%*%';
```
