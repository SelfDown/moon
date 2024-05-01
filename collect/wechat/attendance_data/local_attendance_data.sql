select a.*
from attendance_data a
{{ if .username}}
left join user_account ua  on a.user_id  = ua.userid
{{ end }}
where a.attendance_day >= DATE_SUB(CURDATE(), INTERVAL 3 MONTH)
{{ if .username }}
and ua.username  = {{.username}}
{{ end }}