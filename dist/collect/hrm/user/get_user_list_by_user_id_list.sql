select a.user_name as username,
       a.*
from user_account a
where 1=1
and a.user_id in ({{.user_id_list}})