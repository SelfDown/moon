SELECT
    code.sys_code_text as menu_type_name,
    a.*
FROM sys_menu a
left join sys_code code on code.sys_code_type = "menu_type"  and a.menu_type = code.sys_code
where 1=1
{{ if .search }}
and ( a.name like {{.search}} or a.menu_url like {{.search}})
{{ end }}
order by a.order_index
