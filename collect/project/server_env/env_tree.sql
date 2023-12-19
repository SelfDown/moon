select t.env_name as env_type_name,se.*
from server_env se
left join sys_env  t on se.env_code = t.env_code
where se.flag_del = '0'
{{ if .sys_project_id }}
and se.sys_project_id = {{.sys_project_id}}
{{ end }}
{{ if .server_env_code }}
and se.server_env_code = {{.server_env_code}}
{{ end }}
{{ if .search}}
and (se.server_env_name like {{.search}} or se.server_env_code like {{.search}})
{{ end }}
{{ if .exclude }}
and se.server_env_id != {{.exclude }}
{{ end }}
order by se.sys_project_id, se.order_id
