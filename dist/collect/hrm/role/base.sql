select a.*
from role a
where 1=1
{{ if .search }}
and (
a.role_name like {{.search}}
or a.role_code like {{.search}}
)
{{ end }}
order by a.order_index desc
{{ if  .pagination  }}
limit {{.start}} , {{.size}}
{{ end }}