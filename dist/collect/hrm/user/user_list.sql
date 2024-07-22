select
    a.userid as user_id,
    {{ if eq (get_key "system_model") "company" }}
    a.userpwd as password,
    {{else }}
    a.password as password,
    {{ end }}
    (
        SELECT GROUP_CONCAT(ur.role_des)
        FROM user_role ur
                 left join user_role_id_list r  on  r.role_id  = ur.role_id
        where  r.user_id  = a.userid
    ) as role_names,
    (
        SELECT GROUP_CONCAT(ur.role_code)
        FROM user_role ur
                 left join user_role_id_list r  on  r.role_id  = ur.role_id
        where  r.user_id  = a.userid
    ) as roles,
    a.*
from user_account a
where
ifnull(a.statu,'1')!='0'
require('./base_where.common')


