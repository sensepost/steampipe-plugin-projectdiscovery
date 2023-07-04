# Table: projectdiscovery_httpx

[httpx](https://github.com/projectdiscovery/httpx) httpx is a fast and multi-purpose HTTP toolkit that allows running multiple probes using the retryablehttp library.

## Examples

### Get HTTP status codes and web servers in used by two target URL's

```sql
select
  status_code,
  web_server
from
  projectdiscovery_httpx
where
  target in 
  (
    'https://www.google.com/',
    'https://twitter.com/'
  )
```

### Get webserver technologes in used as per Wapalyzer

```sql
select distinct
  jsonb_array_elements_text(technologies)
from
  projectdiscovery_httpx
where
  target in 
  (
    'https://www.google.com/',
    'https://twitter.com/',
    'https://facebook.com'
  )
```

### Get a SHA1 hash of the response headers

```sql
select
  hashes ->> 'header_sha1'
from
  projectdiscovery_httpx
where
  target = 'https://www.google.com/'
```

### Get HTTP requests to a set of targets as well as the IP's they resolve to

```sql
with targets as
(
  select
    *
  from
    projectdiscovery_httpx
  where
    target in
    (
      'https://www.google.com/',
      'https://twitter.com/',
      'https://facebook.com'
    )
)
,
ips as
(
  select
    *
  from
    projectdiscovery_httpx
  where
    target in
    (
      select
        jsonb_array_elements_text(a)
      from
        targets
    )
)
select
  *
from
  targets
union
select
  *
from
  ips
```
