# Table: projectdiscovery_subfinder

[subfinder](https://github.com/projectdiscovery/subfinder) is a dast passive subdomain enumeration tool.

## Examples

### Get subdomains for a domain

```sql
select
  host
from
  projectdiscovery_subfinder
where
  target = 'google.com'
```

### Get the sources of information for subdomain info

```sql
select distinct
  source
from
  projectdiscovery_subfinder
where
  target = 'google.com'
```

### Count subdomains

```sql
select
  count(*)
from
  projectdiscovery_subfinder
where
  target = 'google.com'
```
