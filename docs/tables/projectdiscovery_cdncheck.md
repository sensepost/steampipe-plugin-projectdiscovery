# Table: projectdiscovery_cdncheck

[cdncheck](https://github.com/projectdiscovery/cdncheck) is a utility to detect various technology for a given IP address.

## Examples

### Get CDN, Cloud Provider and WAF information for an IP

```sql
select
  cdn,
  cloud,
  waf
from
  projectdiscovery_cdncheck
where
  target = '172.217.170.14';
```

### Get CDN, Cloud Provider and WAF information for multiple IP's

```sql
select
  first_ip,
  last_ip
from
  projectdiscovery_asnmap
where
  target in 
  (
    '104.16.132.229',
    '104.16.133.229'
  );
```

### Get CDN information for a domain (depends on the `net` plugin)

```sql
select
  cdn,
  cloud,
  waf
from
  projectdiscovery_cdncheck
where
  target in
  (
    select
      rtrim(ip::text, '/32')
    from
      net_dns_record
    where
      domain = 'google.com'
      and type = 'A' limit 1
  );
```
