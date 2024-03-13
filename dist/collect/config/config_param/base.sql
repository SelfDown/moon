select a.*
from config_detail_change_history a
where 1=1
{{ if .change_id_list}}
and a.change_id in ({{.change_id_list}})
{{ end }}
order by a.create_time desc
{{ if  .pagination  }}
limit {{.start}} , {{.size}}
{{ end }}