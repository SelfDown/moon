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

func newSentryIssues(db *gorm.DB, opts ...gen.DOOption) sentryIssues {
	_sentryIssues := sentryIssues{}

	_sentryIssues.sentryIssuesDo.UseDB(db, opts...)
	_sentryIssues.sentryIssuesDo.UseModel(&model.SentryIssues{})

	tableName := _sentryIssues.sentryIssuesDo.TableName()
	_sentryIssues.ALL = field.NewAsterisk(tableName)
	_sentryIssues.SentryIssuesID = field.NewString(tableName, "sentry_issues_id")
	_sentryIssues.SysProjectCode = field.NewString(tableName, "sys_project_code")
	_sentryIssues.SentryProject = field.NewString(tableName, "sentry_project")
	_sentryIssues.SentryProjectID = field.NewString(tableName, "sentry_project_id")
	_sentryIssues.Title = field.NewString(tableName, "title")
	_sentryIssues.Count_ = field.NewString(tableName, "count")
	_sentryIssues.UserCount = field.NewString(tableName, "user_count")
	_sentryIssues.FirstSeen = field.NewString(tableName, "first_seen")
	_sentryIssues.LastSeen = field.NewString(tableName, "last_seen")
	_sentryIssues.Status = field.NewString(tableName, "status")
	_sentryIssues.ShortID = field.NewString(tableName, "short_id")
	_sentryIssues.DetailLink = field.NewString(tableName, "detail_link")
	_sentryIssues.IssueID = field.NewString(tableName, "issue_id")
	_sentryIssues.IssueType = field.NewString(tableName, "issue_type")
	_sentryIssues.Level = field.NewString(tableName, "level")
	_sentryIssues.IssueCategory = field.NewString(tableName, "issue_category")
	_sentryIssues.MetadataValue = field.NewString(tableName, "metadata_value")
	_sentryIssues.Culprit = field.NewString(tableName, "culprit")
	_sentryIssues.Filename = field.NewString(tableName, "filename")
	_sentryIssues.Event = field.NewString(tableName, "event")
	_sentryIssues.Assignee = field.NewString(tableName, "assignee")
	_sentryIssues.SentryURL = field.NewString(tableName, "sentry_url")
	_sentryIssues.CreateTime = field.NewString(tableName, "create_time")
	_sentryIssues.SentryHasClose = field.NewString(tableName, "sentry_has_close")
	_sentryIssues.RequestURL = field.NewString(tableName, "request_url")
	_sentryIssues.WorkStation = field.NewString(tableName, "work_station")
	_sentryIssues.Team = field.NewString(tableName, "team")

	_sentryIssues.fillFieldMap()

	return _sentryIssues
}

type sentryIssues struct {
	sentryIssuesDo

	ALL             field.Asterisk
	SentryIssuesID  field.String
	SysProjectCode  field.String
	SentryProject   field.String // 项目标识
	SentryProjectID field.String // 项目id
	Title           field.String // 问题的标题
	Count_          field.String // 问题的事件总计数
	UserCount       field.String // 用户的问题总计数
	FirstSeen       field.String // 首次出现的时间
	LastSeen        field.String // 最近出现的时间
	Status          field.String // issue状态
	ShortID         field.String // 短标题
	DetailLink      field.String // 详情链接
	IssueID         field.String // issue_id
	IssueType       field.String // issue类型
	Level           field.String // issue等级
	IssueCategory   field.String // 捕获类型
	MetadataValue   field.String // 报错关键语句
	Culprit         field.String // 子标题
	Filename        field.String // 报错的方法
	Event           field.String // 事件
	Assignee        field.String
	SentryURL       field.String
	CreateTime      field.String // 创建时间
	SentryHasClose  field.String // 是否关闭
	RequestURL      field.String
	WorkStation     field.String
	Team            field.String

	fieldMap map[string]field.Expr
}

func (s sentryIssues) Table(newTableName string) *sentryIssues {
	s.sentryIssuesDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sentryIssues) As(alias string) *sentryIssues {
	s.sentryIssuesDo.DO = *(s.sentryIssuesDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sentryIssues) updateTableName(table string) *sentryIssues {
	s.ALL = field.NewAsterisk(table)
	s.SentryIssuesID = field.NewString(table, "sentry_issues_id")
	s.SysProjectCode = field.NewString(table, "sys_project_code")
	s.SentryProject = field.NewString(table, "sentry_project")
	s.SentryProjectID = field.NewString(table, "sentry_project_id")
	s.Title = field.NewString(table, "title")
	s.Count_ = field.NewString(table, "count")
	s.UserCount = field.NewString(table, "user_count")
	s.FirstSeen = field.NewString(table, "first_seen")
	s.LastSeen = field.NewString(table, "last_seen")
	s.Status = field.NewString(table, "status")
	s.ShortID = field.NewString(table, "short_id")
	s.DetailLink = field.NewString(table, "detail_link")
	s.IssueID = field.NewString(table, "issue_id")
	s.IssueType = field.NewString(table, "issue_type")
	s.Level = field.NewString(table, "level")
	s.IssueCategory = field.NewString(table, "issue_category")
	s.MetadataValue = field.NewString(table, "metadata_value")
	s.Culprit = field.NewString(table, "culprit")
	s.Filename = field.NewString(table, "filename")
	s.Event = field.NewString(table, "event")
	s.Assignee = field.NewString(table, "assignee")
	s.SentryURL = field.NewString(table, "sentry_url")
	s.CreateTime = field.NewString(table, "create_time")
	s.SentryHasClose = field.NewString(table, "sentry_has_close")
	s.RequestURL = field.NewString(table, "request_url")
	s.WorkStation = field.NewString(table, "work_station")
	s.Team = field.NewString(table, "team")

	s.fillFieldMap()

	return s
}

func (s *sentryIssues) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sentryIssues) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 27)
	s.fieldMap["sentry_issues_id"] = s.SentryIssuesID
	s.fieldMap["sys_project_code"] = s.SysProjectCode
	s.fieldMap["sentry_project"] = s.SentryProject
	s.fieldMap["sentry_project_id"] = s.SentryProjectID
	s.fieldMap["title"] = s.Title
	s.fieldMap["count"] = s.Count_
	s.fieldMap["user_count"] = s.UserCount
	s.fieldMap["first_seen"] = s.FirstSeen
	s.fieldMap["last_seen"] = s.LastSeen
	s.fieldMap["status"] = s.Status
	s.fieldMap["short_id"] = s.ShortID
	s.fieldMap["detail_link"] = s.DetailLink
	s.fieldMap["issue_id"] = s.IssueID
	s.fieldMap["issue_type"] = s.IssueType
	s.fieldMap["level"] = s.Level
	s.fieldMap["issue_category"] = s.IssueCategory
	s.fieldMap["metadata_value"] = s.MetadataValue
	s.fieldMap["culprit"] = s.Culprit
	s.fieldMap["filename"] = s.Filename
	s.fieldMap["event"] = s.Event
	s.fieldMap["assignee"] = s.Assignee
	s.fieldMap["sentry_url"] = s.SentryURL
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["sentry_has_close"] = s.SentryHasClose
	s.fieldMap["request_url"] = s.RequestURL
	s.fieldMap["work_station"] = s.WorkStation
	s.fieldMap["team"] = s.Team
}

func (s sentryIssues) clone(db *gorm.DB) sentryIssues {
	s.sentryIssuesDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sentryIssues) replaceDB(db *gorm.DB) sentryIssues {
	s.sentryIssuesDo.ReplaceDB(db)
	return s
}

type sentryIssuesDo struct{ gen.DO }

type ISentryIssuesDo interface {
	gen.SubQuery
	Debug() ISentryIssuesDo
	WithContext(ctx context.Context) ISentryIssuesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISentryIssuesDo
	WriteDB() ISentryIssuesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISentryIssuesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISentryIssuesDo
	Not(conds ...gen.Condition) ISentryIssuesDo
	Or(conds ...gen.Condition) ISentryIssuesDo
	Select(conds ...field.Expr) ISentryIssuesDo
	Where(conds ...gen.Condition) ISentryIssuesDo
	Order(conds ...field.Expr) ISentryIssuesDo
	Distinct(cols ...field.Expr) ISentryIssuesDo
	Omit(cols ...field.Expr) ISentryIssuesDo
	Join(table schema.Tabler, on ...field.Expr) ISentryIssuesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISentryIssuesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISentryIssuesDo
	Group(cols ...field.Expr) ISentryIssuesDo
	Having(conds ...gen.Condition) ISentryIssuesDo
	Limit(limit int) ISentryIssuesDo
	Offset(offset int) ISentryIssuesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISentryIssuesDo
	Unscoped() ISentryIssuesDo
	Create(values ...*model.SentryIssues) error
	CreateInBatches(values []*model.SentryIssues, batchSize int) error
	Save(values ...*model.SentryIssues) error
	First() (*model.SentryIssues, error)
	Take() (*model.SentryIssues, error)
	Last() (*model.SentryIssues, error)
	Find() ([]*model.SentryIssues, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SentryIssues, err error)
	FindInBatches(result *[]*model.SentryIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SentryIssues) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISentryIssuesDo
	Assign(attrs ...field.AssignExpr) ISentryIssuesDo
	Joins(fields ...field.RelationField) ISentryIssuesDo
	Preload(fields ...field.RelationField) ISentryIssuesDo
	FirstOrInit() (*model.SentryIssues, error)
	FirstOrCreate() (*model.SentryIssues, error)
	FindByPage(offset int, limit int) (result []*model.SentryIssues, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISentryIssuesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sentryIssuesDo) Debug() ISentryIssuesDo {
	return s.withDO(s.DO.Debug())
}

func (s sentryIssuesDo) WithContext(ctx context.Context) ISentryIssuesDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sentryIssuesDo) ReadDB() ISentryIssuesDo {
	return s.Clauses(dbresolver.Read)
}

func (s sentryIssuesDo) WriteDB() ISentryIssuesDo {
	return s.Clauses(dbresolver.Write)
}

func (s sentryIssuesDo) Session(config *gorm.Session) ISentryIssuesDo {
	return s.withDO(s.DO.Session(config))
}

func (s sentryIssuesDo) Clauses(conds ...clause.Expression) ISentryIssuesDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sentryIssuesDo) Returning(value interface{}, columns ...string) ISentryIssuesDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sentryIssuesDo) Not(conds ...gen.Condition) ISentryIssuesDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sentryIssuesDo) Or(conds ...gen.Condition) ISentryIssuesDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sentryIssuesDo) Select(conds ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sentryIssuesDo) Where(conds ...gen.Condition) ISentryIssuesDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sentryIssuesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISentryIssuesDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sentryIssuesDo) Order(conds ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sentryIssuesDo) Distinct(cols ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sentryIssuesDo) Omit(cols ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sentryIssuesDo) Join(table schema.Tabler, on ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sentryIssuesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sentryIssuesDo) RightJoin(table schema.Tabler, on ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sentryIssuesDo) Group(cols ...field.Expr) ISentryIssuesDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sentryIssuesDo) Having(conds ...gen.Condition) ISentryIssuesDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sentryIssuesDo) Limit(limit int) ISentryIssuesDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sentryIssuesDo) Offset(offset int) ISentryIssuesDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sentryIssuesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISentryIssuesDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sentryIssuesDo) Unscoped() ISentryIssuesDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sentryIssuesDo) Create(values ...*model.SentryIssues) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sentryIssuesDo) CreateInBatches(values []*model.SentryIssues, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sentryIssuesDo) Save(values ...*model.SentryIssues) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sentryIssuesDo) First() (*model.SentryIssues, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SentryIssues), nil
	}
}

func (s sentryIssuesDo) Take() (*model.SentryIssues, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SentryIssues), nil
	}
}

func (s sentryIssuesDo) Last() (*model.SentryIssues, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SentryIssues), nil
	}
}

func (s sentryIssuesDo) Find() ([]*model.SentryIssues, error) {
	result, err := s.DO.Find()
	return result.([]*model.SentryIssues), err
}

func (s sentryIssuesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SentryIssues, err error) {
	buf := make([]*model.SentryIssues, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sentryIssuesDo) FindInBatches(result *[]*model.SentryIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sentryIssuesDo) Attrs(attrs ...field.AssignExpr) ISentryIssuesDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sentryIssuesDo) Assign(attrs ...field.AssignExpr) ISentryIssuesDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sentryIssuesDo) Joins(fields ...field.RelationField) ISentryIssuesDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sentryIssuesDo) Preload(fields ...field.RelationField) ISentryIssuesDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sentryIssuesDo) FirstOrInit() (*model.SentryIssues, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SentryIssues), nil
	}
}

func (s sentryIssuesDo) FirstOrCreate() (*model.SentryIssues, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SentryIssues), nil
	}
}

func (s sentryIssuesDo) FindByPage(offset int, limit int) (result []*model.SentryIssues, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sentryIssuesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sentryIssuesDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sentryIssuesDo) Delete(models ...*model.SentryIssues) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sentryIssuesDo) withDO(do gen.Dao) *sentryIssuesDo {
	s.DO = *do.(*gen.DO)
	return s
}