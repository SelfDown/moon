select cg.name as group_name ,
       cg.description ,
       g.*
from config_detail g
left join config_group cg on g.group_id =cg.group_id
where 1=1
{{ if .group_id }}
and cg.group_id = {{.group_id}}
{{ end }}
{{ if .config_detail_id_list }}
and g.config_detail_id in ({{.config_detail_id_list}})
{{ end }}
{{ if .search }}
and (
    cg.description like {{.search}}
    or cg.group_id like {{.search}}
    or g.name like {{.search}}
    or g.value like {{.search}}
)
{{ end }}
{{ if .config_detail_list}}
and (g.group_id,g.name) in({{ range $index,$item := .config_detail_list }} {{if $index}},{{end}} ({{$item.group_id}},{{$item.name_copy}}){{ end}})
{{end}}
order by cg.name ,g.name