# Table: projectdiscovery_katana

[katana](https://github.com/projectdiscovery/katana) is a next-generation crawling and spidering framework.

Unfortunately, this table does not yet quite work as expected. The `depth` column does not correctly get interpreted by the tool. There is an open issue for it [here](https://github.com/projectdiscovery/katana/issues/503).

## Examples

### Crawl a single URL

```sql
select
    response
from
    projectdiscovery_katana
where
    target = 'https://www.google.com/'
    and depth = 1
```

### Get response headers from a response

```sql
select
    jsonb_pretty(response -> 'headers')
from
    projectdiscovery_katana
where
    target = 'https://www.google.com/'
    and depth = 1
```
