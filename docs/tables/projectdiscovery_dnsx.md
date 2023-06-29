# Table: projectdiscovery_dnsx

[dnsx](https://github.com/projectdiscovery/dnsx) dnsx is a fast and multi-purpose DNS toolkit allow to run multiple DNS queries of your choice with a list of user-supplied resolvers.

## Examples

### Lookup the A records for a domain

```sql
select
    address
from
    projectdiscovery_dnsx
where
    target = 'google.com';
```
