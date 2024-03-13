select a.*
from sys_code a
where a.sys_code_type = {{.sys_code_type}}
{{ if .sys_code_list }}
and a.sys_code in ({{.sys_code_list}})
{{ end }}