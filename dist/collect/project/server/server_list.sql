select
    dd.sys_code_text AS server_service_type_text,
    f.env_code,
    f.server_env_name,
    (select count(1) from server_install_soft soft  where s.server_id = soft.server_id) as soft_count,
    s.*
from (require(base.sql)) s
LEFT JOIN sys_code dd ON dd.sys_code =  s.server_service_type AND dd.sys_code_type = "server_service_type"
LEFT JOIN server_env f ON f.server_env_id = s.server_env_id
order by inet_aton(s.server_ip)