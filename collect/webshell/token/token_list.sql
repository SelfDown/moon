SELECT sou.user_name ,
       si.server_ip,
       si.server_busi_name,
       ua.nick as create_user_name,
       a.*
FROM ( require(./base.sql) )as a
left join server_os_users sou on sou.server_os_users_id = a.server_os_users_id
left join server_instance si  on si.server_id  = sou.server_id
left join user_account ua  on ua.userid  = a.create_user
