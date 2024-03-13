SELECT a.*
FROM collect_doc a
join doc_group dg on a.parent_dir  = dg.doc_group_id and dg.is_delete  = '0'
where a.is_delete = '0'
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}
{{ if .type }}
and a.type ={{.type}}
{{ end }}
order by a.order_index asc