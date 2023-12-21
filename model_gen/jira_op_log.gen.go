// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"moon/model"
)

func newJiraOpLog(db *gorm.DB, opts ...gen.DOOption) jiraOpLog {
	_jiraOpLog := jiraOpLog{}

	_jiraOpLog.jiraOpLogDo.UseDB(db, opts...)
	_jiraOpLog.jiraOpLogDo.UseModel(&model.JiraOpLog{})

	tableName := _jiraOpLog.jiraOpLogDo.TableName()
	_jiraOpLog.ALL = field.NewAsterisk(tableName)
	_jiraOpLog.JiraLogID = field.NewString(tableName, "jira_log_id")
	_jiraOpLog.CreateUser = field.NewString(tableName, "create_user")
	_jiraOpLog.CreateTime = field.NewString(tableName, "create_time")
	_jiraOpLog.Op = field.NewString(tableName, "op")
	_jiraOpLog.ProjectKey = field.NewString(tableName, "project_key")
	_jiraOpLog.Summary = field.NewString(tableName, "summary")
	_jiraOpLog.Description = field.NewString(tableName, "description")
	_jiraOpLog.Label = field.NewString(tableName, "label")
	_jiraOpLog.Assignee = field.NewString(tableName, "assignee")
	_jiraOpLog.ComponentName = field.NewString(tableName, "component_name")
	_jiraOpLog.FixVersion = field.NewString(tableName, "fix_version")
	_jiraOpLog.Reporter = field.NewString(tableName, "reporter")
	_jiraOpLog.Priority = field.NewString(tableName, "priority")
	_jiraOpLog.IssueType = field.NewString(tableName, "issue_type")
	_jiraOpLog.HasScript = field.NewString(tableName, "has_script")
	_jiraOpLog.Duedate = field.NewString(tableName, "duedate")
	_jiraOpLog.Beans = field.NewString(tableName, "beans")
	_jiraOpLog.Msg = field.NewString(tableName, "msg")
	_jiraOpLog.Success = field.NewString(tableName, "success")
	_jiraOpLog.IssueKey = field.NewString(tableName, "issue_key")

	_jiraOpLog.fillFieldMap()

	return _jiraOpLog
}

type jiraOpLog struct {
	jiraOpLogDo

	ALL           field.Asterisk
	JiraLogID     field.String
	CreateUser    field.String
	CreateTime    field.String
	Op            field.String
	ProjectKey    field.String
	Summary       field.String
	Description   field.String
	Label         field.String
	Assignee      field.String
	ComponentName field.String
	FixVersion    field.String
	Reporter      field.String
	Priority      field.String
	IssueType     field.String
	HasScript     field.String
	Duedate       field.String
	Beans         field.String
	Msg           field.String
	Success       field.String
	IssueKey      field.String

	fieldMap map[string]field.Expr
}

func (j jiraOpLog) Table(newTableName string) *jiraOpLog {
	j.jiraOpLogDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jiraOpLog) As(alias string) *jiraOpLog {
	j.jiraOpLogDo.DO = *(j.jiraOpLogDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jiraOpLog) updateTableName(table string) *jiraOpLog {
	j.ALL = field.NewAsterisk(table)
	j.JiraLogID = field.NewString(table, "jira_log_id")
	j.CreateUser = field.NewString(table, "create_user")
	j.CreateTime = field.NewString(table, "create_time")
	j.Op = field.NewString(table, "op")
	j.ProjectKey = field.NewString(table, "project_key")
	j.Summary = field.NewString(table, "summary")
	j.Description = field.NewString(table, "description")
	j.Label = field.NewString(table, "label")
	j.Assignee = field.NewString(table, "assignee")
	j.ComponentName = field.NewString(table, "component_name")
	j.FixVersion = field.NewString(table, "fix_version")
	j.Reporter = field.NewString(table, "reporter")
	j.Priority = field.NewString(table, "priority")
	j.IssueType = field.NewString(table, "issue_type")
	j.HasScript = field.NewString(table, "has_script")
	j.Duedate = field.NewString(table, "duedate")
	j.Beans = field.NewString(table, "beans")
	j.Msg = field.NewString(table, "msg")
	j.Success = field.NewString(table, "success")
	j.IssueKey = field.NewString(table, "issue_key")

	j.fillFieldMap()

	return j
}

func (j *jiraOpLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jiraOpLog) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 20)
	j.fieldMap["jira_log_id"] = j.JiraLogID
	j.fieldMap["create_user"] = j.CreateUser
	j.fieldMap["create_time"] = j.CreateTime
	j.fieldMap["op"] = j.Op
	j.fieldMap["project_key"] = j.ProjectKey
	j.fieldMap["summary"] = j.Summary
	j.fieldMap["description"] = j.Description
	j.fieldMap["label"] = j.Label
	j.fieldMap["assignee"] = j.Assignee
	j.fieldMap["component_name"] = j.ComponentName
	j.fieldMap["fix_version"] = j.FixVersion
	j.fieldMap["reporter"] = j.Reporter
	j.fieldMap["priority"] = j.Priority
	j.fieldMap["issue_type"] = j.IssueType
	j.fieldMap["has_script"] = j.HasScript
	j.fieldMap["duedate"] = j.Duedate
	j.fieldMap["beans"] = j.Beans
	j.fieldMap["msg"] = j.Msg
	j.fieldMap["success"] = j.Success
	j.fieldMap["issue_key"] = j.IssueKey
}

func (j jiraOpLog) clone(db *gorm.DB) jiraOpLog {
	j.jiraOpLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jiraOpLog) replaceDB(db *gorm.DB) jiraOpLog {
	j.jiraOpLogDo.ReplaceDB(db)
	return j
}

type jiraOpLogDo struct{ gen.DO }

type IJiraOpLogDo interface {
	gen.SubQuery
	Debug() IJiraOpLogDo
	WithContext(ctx context.Context) IJiraOpLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJiraOpLogDo
	WriteDB() IJiraOpLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJiraOpLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJiraOpLogDo
	Not(conds ...gen.Condition) IJiraOpLogDo
	Or(conds ...gen.Condition) IJiraOpLogDo
	Select(conds ...field.Expr) IJiraOpLogDo
	Where(conds ...gen.Condition) IJiraOpLogDo
	Order(conds ...field.Expr) IJiraOpLogDo
	Distinct(cols ...field.Expr) IJiraOpLogDo
	Omit(cols ...field.Expr) IJiraOpLogDo
	Join(table schema.Tabler, on ...field.Expr) IJiraOpLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJiraOpLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJiraOpLogDo
	Group(cols ...field.Expr) IJiraOpLogDo
	Having(conds ...gen.Condition) IJiraOpLogDo
	Limit(limit int) IJiraOpLogDo
	Offset(offset int) IJiraOpLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJiraOpLogDo
	Unscoped() IJiraOpLogDo
	Create(values ...*model.JiraOpLog) error
	CreateInBatches(values []*model.JiraOpLog, batchSize int) error
	Save(values ...*model.JiraOpLog) error
	First() (*model.JiraOpLog, error)
	Take() (*model.JiraOpLog, error)
	Last() (*model.JiraOpLog, error)
	Find() ([]*model.JiraOpLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JiraOpLog, err error)
	FindInBatches(result *[]*model.JiraOpLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.JiraOpLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJiraOpLogDo
	Assign(attrs ...field.AssignExpr) IJiraOpLogDo
	Joins(fields ...field.RelationField) IJiraOpLogDo
	Preload(fields ...field.RelationField) IJiraOpLogDo
	FirstOrInit() (*model.JiraOpLog, error)
	FirstOrCreate() (*model.JiraOpLog, error)
	FindByPage(offset int, limit int) (result []*model.JiraOpLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJiraOpLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jiraOpLogDo) Debug() IJiraOpLogDo {
	return j.withDO(j.DO.Debug())
}

func (j jiraOpLogDo) WithContext(ctx context.Context) IJiraOpLogDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jiraOpLogDo) ReadDB() IJiraOpLogDo {
	return j.Clauses(dbresolver.Read)
}

func (j jiraOpLogDo) WriteDB() IJiraOpLogDo {
	return j.Clauses(dbresolver.Write)
}

func (j jiraOpLogDo) Session(config *gorm.Session) IJiraOpLogDo {
	return j.withDO(j.DO.Session(config))
}

func (j jiraOpLogDo) Clauses(conds ...clause.Expression) IJiraOpLogDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jiraOpLogDo) Returning(value interface{}, columns ...string) IJiraOpLogDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jiraOpLogDo) Not(conds ...gen.Condition) IJiraOpLogDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jiraOpLogDo) Or(conds ...gen.Condition) IJiraOpLogDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jiraOpLogDo) Select(conds ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jiraOpLogDo) Where(conds ...gen.Condition) IJiraOpLogDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jiraOpLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IJiraOpLogDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jiraOpLogDo) Order(conds ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jiraOpLogDo) Distinct(cols ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jiraOpLogDo) Omit(cols ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jiraOpLogDo) Join(table schema.Tabler, on ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jiraOpLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jiraOpLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jiraOpLogDo) Group(cols ...field.Expr) IJiraOpLogDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jiraOpLogDo) Having(conds ...gen.Condition) IJiraOpLogDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jiraOpLogDo) Limit(limit int) IJiraOpLogDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jiraOpLogDo) Offset(offset int) IJiraOpLogDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jiraOpLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJiraOpLogDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jiraOpLogDo) Unscoped() IJiraOpLogDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jiraOpLogDo) Create(values ...*model.JiraOpLog) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jiraOpLogDo) CreateInBatches(values []*model.JiraOpLog, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jiraOpLogDo) Save(values ...*model.JiraOpLog) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jiraOpLogDo) First() (*model.JiraOpLog, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JiraOpLog), nil
	}
}

func (j jiraOpLogDo) Take() (*model.JiraOpLog, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JiraOpLog), nil
	}
}

func (j jiraOpLogDo) Last() (*model.JiraOpLog, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JiraOpLog), nil
	}
}

func (j jiraOpLogDo) Find() ([]*model.JiraOpLog, error) {
	result, err := j.DO.Find()
	return result.([]*model.JiraOpLog), err
}

func (j jiraOpLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JiraOpLog, err error) {
	buf := make([]*model.JiraOpLog, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jiraOpLogDo) FindInBatches(result *[]*model.JiraOpLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jiraOpLogDo) Attrs(attrs ...field.AssignExpr) IJiraOpLogDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jiraOpLogDo) Assign(attrs ...field.AssignExpr) IJiraOpLogDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jiraOpLogDo) Joins(fields ...field.RelationField) IJiraOpLogDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jiraOpLogDo) Preload(fields ...field.RelationField) IJiraOpLogDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jiraOpLogDo) FirstOrInit() (*model.JiraOpLog, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JiraOpLog), nil
	}
}

func (j jiraOpLogDo) FirstOrCreate() (*model.JiraOpLog, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JiraOpLog), nil
	}
}

func (j jiraOpLogDo) FindByPage(offset int, limit int) (result []*model.JiraOpLog, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jiraOpLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jiraOpLogDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jiraOpLogDo) Delete(models ...*model.JiraOpLog) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jiraOpLogDo) withDO(do gen.Dao) *jiraOpLogDo {
	j.DO = *do.(*gen.DO)
	return j
}
