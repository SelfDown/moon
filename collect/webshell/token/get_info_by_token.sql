SELECT  a.token, a.server_os_users_id,si.server_ip ,sou.user_name ,sou.user_pwd
FROM webshell_token a
left join  server_os_users sou on concat(sou.server_os_users_id,'#',sou.server_os_user_id)  = a.server_os_users_id
left join server_instance si  on si.server_id  = sou.server_id
where a.token = {{.token}} and a.close_time = ""
{{ if .once_token }}
and a.is_valid  = '1'
{{ end }}