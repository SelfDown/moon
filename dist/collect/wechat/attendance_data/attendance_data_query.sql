select
    at1.name as attendance_type_name,
    at1.short_name as attendance_type_short_name,
    at2.name as attendance_morning_type_name,
    at2.short_name as attendance_morning_type_short_name,
    at3.name as attendance_afternoon_type_name,
    at3.short_name as attendance_afternoon_type_short_name,
    a.*
from attendance_data a
left join attendance_type at1 on a.attendance_type = at1.attendance_type
left join attendance_type at2 on a.attendance_morning_type = at2.attendance_type
left join attendance_type at3 on a.attendance_afternoon_type = at3.attendance_type
where
DATE_FORMAT(attendance_day, '%Y-%m')={{.month}}
{{ if .user_id}}
and a.user_id = {{.user_id}}
{{ end}}
order by a.attendance_day