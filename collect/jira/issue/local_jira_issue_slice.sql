select a.*
from jira_issue_slice a
WHERE
{{ if .issue_key_list }}
a.issue_key in ({{.issue_key_list}})
{{ else }}
a.updated > DATE_SUB(CURDATE(), INTERVAL 1 DAY)
{{ end }}
