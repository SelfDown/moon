select b.server_ip,a.*
from server_os_users a
left join server_instance  b on a.server_id  = b.server_id
where 1=1
{{ if .server_os_users_id}}
and a.server_os_users_id = {{.server_os_users_id}}
{{ end }}
{{ if .server_id_list}}
and a.server_id in ({{.server_id_list}})
{{ end }}
{{ if .server_id }}
and a.server_id = {{ .server_id }}
{{ end }}
{{ if .server_user_names}}
and a.user_name in ({{.server_user_names}})
{{ end }}