select a.*,
(
   select GROUP_CONCAT(name )
   from ldap_group lg
   join role_ldap_group rlg on rlg.ldap_group_id =lg.ldap_group_id
   where rlg.role_id =a.role_id
) as ldap_names,
(
    select GROUP_CONCAT(lg.ldap_group_id)
    from ldap_group lg
    join role_ldap_group rlg   on rlg.ldap_group_id =lg.ldap_group_id
    where rlg.role_id =a.role_id
) as ldap_group_ids
from (require(base.sql)) as a