SELECT a.*
FROM collect_doc_demo a
{{ if .doc_type }}
join collect_doc b on a.collect_doc_id = b.collect_doc_id and b.is_delete = '0'
{{ end }}
where 1=1
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}
{{ if .doc_type }}
and b.type  = {{.doc_type}}
{{ end }}
order by a.order_index asc