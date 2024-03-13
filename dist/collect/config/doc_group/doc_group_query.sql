select a.*
from doc_group a
where a.is_delete = '0'
{{ if .name_list }}
and a.name in ({{.name_list}})
{{ end }}
order by a.type,a.order_index