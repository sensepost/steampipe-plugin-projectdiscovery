# Table: projectdiscovery_tlsx

[tlsx](https://github.com/projectdiscovery/tlsx) is a fast and configurable TLS grabber focused on TLS based data collection.

## Examples

### Get a JARM hash for a remote host

```sql
select
  jarm_hash
from
  projectdiscovery_tlsx
where
  target = 'google.com';
```

### Get Subject Alternate Names from a remote TLS certificate

```sql
select
  jsonb_array_elements_text(certificate_response -> 'subject_an')
from
  projectdiscovery_tlsx
where
  target = 'google.com';
```

### Get serial numbers of certificates in a remote certificate chain

```sql
select
  jsonb_path_query(chain, '$.serial')
from
  projectdiscovery_tlsx
where
  target = 'google.com';
```

### Check certificate expiryt for remote services

```sql
select
  jsonb_pretty(certificate_response -> 'not_after')::timestamp
from
  projectdiscovery_tlsx
where
  target in
  (
    'google.com',
    'twitter.com',
    'facebook.com'
  );
```
