select a.*
from server_instance a
where 1=1
{{ if .server_env_id }}
and a.server_env_id = {{.server_env_id}}
{{ end }}
order by inet_aton(a.server_ip)