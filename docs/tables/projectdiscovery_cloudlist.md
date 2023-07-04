# Table: projectdiscovery_cloudlist

[cloudlist](https://github.com/projectdiscovery/cloudlist) Cloudlist is a tool for listing Assets from multiple Cloud Providers.

**Note:** You need valid API credentials for supported cloud providers.

## Examples

### Get public IPv4 adresses & DNS names for assets in your Digital Ocean account

```sql
select
  public_ipv4,
  dns_name
from
  projectdiscovery_cloudlist
where
  provider = 'do'
  and public_ipv4 != ''
```

### Get private IPv4 adresses for assets in your Digital Ocean account

```sql
select
  private_ipv4
from
  projectdiscovery_cloudlist
where
  provider = 'do'
  and public = false
```
