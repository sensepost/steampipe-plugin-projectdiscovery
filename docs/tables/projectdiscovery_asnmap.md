# Table: projectdiscovery_asnmap

[asnmap](https://github.com/projectdiscovery/asnmap) allows for quickly mapping organization network ranges using ASN information.

## Examples

### Get ASN's for an organisation by name

```sql
select distinct
  asn
from
  projectdiscovery_asnmap
where
  target = 'google'
```

### Get first and last IP of an IP block

```sql
select
  first_ip,
  last_ip
from
  projectdiscovery_asnmap
where
  target = '172.217.170.110'
```

### Get the organisation name that an ASN belongs to

```sql
select distinct
  org
from
  projectdiscovery_asnmap
where
  target = '15169'
```
