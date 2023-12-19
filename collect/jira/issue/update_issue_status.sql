update  publish_issue_list a,jira_issue_slice b
set a.issue_statu = b.issue_statu
where a.issue_key in ({{.issue_key_change_list}}) and a.issue_key = b.issue_key