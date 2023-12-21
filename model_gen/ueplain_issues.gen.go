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

func newUeplainIssues(db *gorm.DB, opts ...gen.DOOption) ueplainIssues {
	_ueplainIssues := ueplainIssues{}

	_ueplainIssues.ueplainIssuesDo.UseDB(db, opts...)
	_ueplainIssues.ueplainIssuesDo.UseModel(&model.UeplainIssues{})

	tableName := _ueplainIssues.ueplainIssuesDo.TableName()
	_ueplainIssues.ALL = field.NewAsterisk(tableName)
	_ueplainIssues.UeplainIssuesID = field.NewString(tableName, "ueplain_issues_id")
	_ueplainIssues.IssueKey = field.NewString(tableName, "issue_key")
	_ueplainIssues.Summary = field.NewString(tableName, "summary")
	_ueplainIssues.Assignee = field.NewString(tableName, "assignee")
	_ueplainIssues.Priority = field.NewString(tableName, "priority")
	_ueplainIssues.Status = field.NewString(tableName, "status")
	_ueplainIssues.Tags = field.NewString(tableName, "tags")
	_ueplainIssues.IssueCreateTime = field.NewString(tableName, "issue_create_time")
	_ueplainIssues.Duedate = field.NewString(tableName, "duedate")
	_ueplainIssues.LinkIssues = field.NewString(tableName, "link_issues")
	_ueplainIssues.ConfluenceAddr = field.NewString(tableName, "confluence_addr")
	_ueplainIssues.CreateUser = field.NewString(tableName, "create_user")
	_ueplainIssues.CreateTime = field.NewString(tableName, "create_time")
	_ueplainIssues.Tag = field.NewString(tableName, "tag")
	_ueplainIssues.LinkStatus = field.NewString(tableName, "link_status")

	_ueplainIssues.fillFieldMap()

	return _ueplainIssues
}

type ueplainIssues struct {
	ueplainIssuesDo

	ALL             field.Asterisk
	UeplainIssuesID field.String
	IssueKey        field.String // 标题
	Summary         field.String
	Assignee        field.String
	Priority        field.String
	Status          field.String
	Tags            field.String
	IssueCreateTime field.String
	Duedate         field.String
	LinkIssues      field.String
	ConfluenceAddr  field.String
	CreateUser      field.String
	CreateTime      field.String
	Tag             field.String
	LinkStatus      field.String

	fieldMap map[string]field.Expr
}

func (u ueplainIssues) Table(newTableName string) *ueplainIssues {
	u.ueplainIssuesDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u ueplainIssues) As(alias string) *ueplainIssues {
	u.ueplainIssuesDo.DO = *(u.ueplainIssuesDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *ueplainIssues) updateTableName(table string) *ueplainIssues {
	u.ALL = field.NewAsterisk(table)
	u.UeplainIssuesID = field.NewString(table, "ueplain_issues_id")
	u.IssueKey = field.NewString(table, "issue_key")
	u.Summary = field.NewString(table, "summary")
	u.Assignee = field.NewString(table, "assignee")
	u.Priority = field.NewString(table, "priority")
	u.Status = field.NewString(table, "status")
	u.Tags = field.NewString(table, "tags")
	u.IssueCreateTime = field.NewString(table, "issue_create_time")
	u.Duedate = field.NewString(table, "duedate")
	u.LinkIssues = field.NewString(table, "link_issues")
	u.ConfluenceAddr = field.NewString(table, "confluence_addr")
	u.CreateUser = field.NewString(table, "create_user")
	u.CreateTime = field.NewString(table, "create_time")
	u.Tag = field.NewString(table, "tag")
	u.LinkStatus = field.NewString(table, "link_status")

	u.fillFieldMap()

	return u
}

func (u *ueplainIssues) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *ueplainIssues) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 15)
	u.fieldMap["ueplain_issues_id"] = u.UeplainIssuesID
	u.fieldMap["issue_key"] = u.IssueKey
	u.fieldMap["summary"] = u.Summary
	u.fieldMap["assignee"] = u.Assignee
	u.fieldMap["priority"] = u.Priority
	u.fieldMap["status"] = u.Status
	u.fieldMap["tags"] = u.Tags
	u.fieldMap["issue_create_time"] = u.IssueCreateTime
	u.fieldMap["duedate"] = u.Duedate
	u.fieldMap["link_issues"] = u.LinkIssues
	u.fieldMap["confluence_addr"] = u.ConfluenceAddr
	u.fieldMap["create_user"] = u.CreateUser
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["tag"] = u.Tag
	u.fieldMap["link_status"] = u.LinkStatus
}

func (u ueplainIssues) clone(db *gorm.DB) ueplainIssues {
	u.ueplainIssuesDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u ueplainIssues) replaceDB(db *gorm.DB) ueplainIssues {
	u.ueplainIssuesDo.ReplaceDB(db)
	return u
}

type ueplainIssuesDo struct{ gen.DO }

type IUeplainIssuesDo interface {
	gen.SubQuery
	Debug() IUeplainIssuesDo
	WithContext(ctx context.Context) IUeplainIssuesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUeplainIssuesDo
	WriteDB() IUeplainIssuesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUeplainIssuesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUeplainIssuesDo
	Not(conds ...gen.Condition) IUeplainIssuesDo
	Or(conds ...gen.Condition) IUeplainIssuesDo
	Select(conds ...field.Expr) IUeplainIssuesDo
	Where(conds ...gen.Condition) IUeplainIssuesDo
	Order(conds ...field.Expr) IUeplainIssuesDo
	Distinct(cols ...field.Expr) IUeplainIssuesDo
	Omit(cols ...field.Expr) IUeplainIssuesDo
	Join(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo
	Group(cols ...field.Expr) IUeplainIssuesDo
	Having(conds ...gen.Condition) IUeplainIssuesDo
	Limit(limit int) IUeplainIssuesDo
	Offset(offset int) IUeplainIssuesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUeplainIssuesDo
	Unscoped() IUeplainIssuesDo
	Create(values ...*model.UeplainIssues) error
	CreateInBatches(values []*model.UeplainIssues, batchSize int) error
	Save(values ...*model.UeplainIssues) error
	First() (*model.UeplainIssues, error)
	Take() (*model.UeplainIssues, error)
	Last() (*model.UeplainIssues, error)
	Find() ([]*model.UeplainIssues, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UeplainIssues, err error)
	FindInBatches(result *[]*model.UeplainIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UeplainIssues) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUeplainIssuesDo
	Assign(attrs ...field.AssignExpr) IUeplainIssuesDo
	Joins(fields ...field.RelationField) IUeplainIssuesDo
	Preload(fields ...field.RelationField) IUeplainIssuesDo
	FirstOrInit() (*model.UeplainIssues, error)
	FirstOrCreate() (*model.UeplainIssues, error)
	FindByPage(offset int, limit int) (result []*model.UeplainIssues, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUeplainIssuesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u ueplainIssuesDo) Debug() IUeplainIssuesDo {
	return u.withDO(u.DO.Debug())
}

func (u ueplainIssuesDo) WithContext(ctx context.Context) IUeplainIssuesDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u ueplainIssuesDo) ReadDB() IUeplainIssuesDo {
	return u.Clauses(dbresolver.Read)
}

func (u ueplainIssuesDo) WriteDB() IUeplainIssuesDo {
	return u.Clauses(dbresolver.Write)
}

func (u ueplainIssuesDo) Session(config *gorm.Session) IUeplainIssuesDo {
	return u.withDO(u.DO.Session(config))
}

func (u ueplainIssuesDo) Clauses(conds ...clause.Expression) IUeplainIssuesDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u ueplainIssuesDo) Returning(value interface{}, columns ...string) IUeplainIssuesDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u ueplainIssuesDo) Not(conds ...gen.Condition) IUeplainIssuesDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u ueplainIssuesDo) Or(conds ...gen.Condition) IUeplainIssuesDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u ueplainIssuesDo) Select(conds ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u ueplainIssuesDo) Where(conds ...gen.Condition) IUeplainIssuesDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u ueplainIssuesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUeplainIssuesDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u ueplainIssuesDo) Order(conds ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u ueplainIssuesDo) Distinct(cols ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u ueplainIssuesDo) Omit(cols ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u ueplainIssuesDo) Join(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u ueplainIssuesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u ueplainIssuesDo) RightJoin(table schema.Tabler, on ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u ueplainIssuesDo) Group(cols ...field.Expr) IUeplainIssuesDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u ueplainIssuesDo) Having(conds ...gen.Condition) IUeplainIssuesDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u ueplainIssuesDo) Limit(limit int) IUeplainIssuesDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u ueplainIssuesDo) Offset(offset int) IUeplainIssuesDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u ueplainIssuesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUeplainIssuesDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u ueplainIssuesDo) Unscoped() IUeplainIssuesDo {
	return u.withDO(u.DO.Unscoped())
}

func (u ueplainIssuesDo) Create(values ...*model.UeplainIssues) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u ueplainIssuesDo) CreateInBatches(values []*model.UeplainIssues, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u ueplainIssuesDo) Save(values ...*model.UeplainIssues) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u ueplainIssuesDo) First() (*model.UeplainIssues, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UeplainIssues), nil
	}
}

func (u ueplainIssuesDo) Take() (*model.UeplainIssues, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UeplainIssues), nil
	}
}

func (u ueplainIssuesDo) Last() (*model.UeplainIssues, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UeplainIssues), nil
	}
}

func (u ueplainIssuesDo) Find() ([]*model.UeplainIssues, error) {
	result, err := u.DO.Find()
	return result.([]*model.UeplainIssues), err
}

func (u ueplainIssuesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UeplainIssues, err error) {
	buf := make([]*model.UeplainIssues, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u ueplainIssuesDo) FindInBatches(result *[]*model.UeplainIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u ueplainIssuesDo) Attrs(attrs ...field.AssignExpr) IUeplainIssuesDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u ueplainIssuesDo) Assign(attrs ...field.AssignExpr) IUeplainIssuesDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u ueplainIssuesDo) Joins(fields ...field.RelationField) IUeplainIssuesDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u ueplainIssuesDo) Preload(fields ...field.RelationField) IUeplainIssuesDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u ueplainIssuesDo) FirstOrInit() (*model.UeplainIssues, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UeplainIssues), nil
	}
}

func (u ueplainIssuesDo) FirstOrCreate() (*model.UeplainIssues, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UeplainIssues), nil
	}
}

func (u ueplainIssuesDo) FindByPage(offset int, limit int) (result []*model.UeplainIssues, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u ueplainIssuesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u ueplainIssuesDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u ueplainIssuesDo) Delete(models ...*model.UeplainIssues) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *ueplainIssuesDo) withDO(do gen.Dao) *ueplainIssuesDo {
	u.DO = *do.(*gen.DO)
	return u
}