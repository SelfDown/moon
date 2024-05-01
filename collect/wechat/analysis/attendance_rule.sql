SELECT a.*,arg.`default`,arg.weight
FROM attendance_rule a
left join attendance_rule_group arg on a.rule_group_id  = arg.rule_group_id