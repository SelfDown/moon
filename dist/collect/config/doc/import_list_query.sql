SELECT a.*
FROM collect_doc_important a
{{ if .doc_type }}
left join collect_doc b on a.collect_doc_id = b.collect_doc_id
{{ end }}
where 1=1
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}
{{ if .doc_type }}
and b.type  = {{.doc_type}}
{{ end }}
order by a.order_index asc

