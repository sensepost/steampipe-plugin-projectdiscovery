# Table: projectdiscovery_nabuu

[nabuu](https://github.com/projectdiscovery/nabuu) is a fast port scanner written in go with a focus on reliability and simplicity.

## Examples

### Perform a top 1000 open ports scan

```sql
select
  port
from
  projectdiscovery_naabu
where
  target = 'scanme.sh'
```

### Perform a top 1000 open ports scan of a network range

```sql
select
  host,
  port
from
  projectdiscovery_naabu
where
  target = '192.168.0.0/29'
```
