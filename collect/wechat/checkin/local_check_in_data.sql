SELECT a.*
FROM check_in_data a
where a.sch_checkin_time  >= {{.start_time}}
