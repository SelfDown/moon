SELECT a.*
FROM collect_doc_important a
where 1=1
{{ if .collect_doc_id }}
and a.collect_doc_id ={{.collect_doc_id}}
{{ end }}
order by a.order_index asc

