select a.*
from server_instance a
{{ if .project_code }}
left join server_env env on env.server_env_id = a.server_env_id
left join sys_projects p on env.sys_project_id = p.sys_project_id
{{ end }}
where 1=1 and ifnull(a.is_del,'0') = '0'
{{ if .project_code}}
and p.project_code = {{.project_code}}
{{ end }}
{{ if .server_env_id }}
and a.server_env_id = {{.server_env_id}}
{{ end }}
{{ if .search }}
and (
    a.server_ip like {{.search}} or a.server_name like {{.search}}
)
{{ end }}
{{ if .server_ip }}
and a.server_ip = {{.server_ip}}
{{ end }}
{{ if .server_id }}
and a.server_id = {{.server_id}}
{{ end }}
{{ if .exclude }}
and a.server_id != {{.exclude}}
{{ end }}
