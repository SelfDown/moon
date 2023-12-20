select ua.nick as create_username,
       ua2.nick as target_username,
       a.*
from (require(base.sql)) as a
left join user_account ua on a.create_user = ua.user_id
left join user_account ua2 on a.user_id = ua2.user_id
order by a.create_time desc