SELECT a.*
FROM
webshell_token a
{{ if .server_id }}
left join server_os_users sou on sou.server_os_users_id = a.server_os_users_id
{{ end }}
{{ if .search }}
left join user_account ua  on ua.userid  = a.create_user
{{ end }}
where 1=1
{{ if .search}}
and ua.nick like {{.search}}
{{ end }}
{{ if .server_id }}
and sou.server_id = {{.server_id}}
{{ end }}
order by a.create_time desc
{{ if  .pagination  }}
limit {{.start}} , {{.size}}
{{ end }}