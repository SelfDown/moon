SELECT a.summary AS issue_title, a.updated,
d.pname as issue_type,
 concat(b.pkey, '-', a.issuenum) AS issue_key , c.pname AS issue_statu
FROM jiraissue a
JOIN project b ON a.project = b.ID
JOIN issuestatus c ON c.ID = a.issuestatus
join issuetype d on a.issuetype = d.ID
WHERE a.updated > DATE_SUB(CURDATE(), INTERVAL 1 DAY)
