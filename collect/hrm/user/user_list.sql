select a.user_name as username,
       j.sys_code_text as user_status_name,
       (
           SELECT GROUP_CONCAT(r.role_name)
           FROM user_role ur
           left join role r  on  r.role_code  = ur.role_id
           where  ur.user_id  = a.user_id
        ) as role_names,
       (
           SELECT GROUP_CONCAT(r.role_code)
           FROM user_role ur
           left join role r  on  r.role_code  = ur.role_id
           where  ur.user_id  = a.user_id
        ) as roles,
       a.*
from user_account a
left join sys_code j on a.user_status = j.sys_code and j.sys_code_type='user_job_status'
where
a.is_delete='0'
require('./base_where.common')


