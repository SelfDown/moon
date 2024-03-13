select sp.*
from sys_projects sp
where sp.flag_del = '0'
{{ if .project_code }}
and sp.project_code = {{.project_code}}
{{ end }}
{{ if .exclude }}
and sp.sys_project_id != {{.exclude}}
{{ end }}
order by sp.order_id