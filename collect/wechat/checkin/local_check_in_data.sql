SELECT b.userid as user_id,a.*
FROM check_in_data a
left join user_account b on a.userid = b.wechat_userid
where a.sch_checkin_time  >= {{.start_time}}
