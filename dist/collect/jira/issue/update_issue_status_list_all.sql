update publish_issue_list
set issue_statu = (
    select issue_statu
    from jira_issue_slice b
    where b.issue_key = publish_issue_list.issue_key
    limit 1
    )
where EXISTS (
    select issue_statu
    from jira_issue_slice b
    where b.issue_key = publish_issue_list.issue_key and b.issue_statu != publish_issue_list.issue_statu
    )