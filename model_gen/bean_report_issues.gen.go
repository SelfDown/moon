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

func newBeanReportIssues(db *gorm.DB, opts ...gen.DOOption) beanReportIssues {
	_beanReportIssues := beanReportIssues{}

	_beanReportIssues.beanReportIssuesDo.UseDB(db, opts...)
	_beanReportIssues.beanReportIssuesDo.UseModel(&model.BeanReportIssues{})

	tableName := _beanReportIssues.beanReportIssuesDo.TableName()
	_beanReportIssues.ALL = field.NewAsterisk(tableName)
	_beanReportIssues.BeanReportIssuesID = field.NewString(tableName, "bean_report_issues_id")
	_beanReportIssues.BeanReportID = field.NewString(tableName, "bean_report_id")
	_beanReportIssues.Team = field.NewString(tableName, "team")
	_beanReportIssues.JqlVersion = field.NewString(tableName, "jql_version")
	_beanReportIssues.JqlNoVersion = field.NewString(tableName, "jql_no_version")
	_beanReportIssues.IssueVersionData = field.NewString(tableName, "issue_version_data")
	_beanReportIssues.IssueNoVersionData = field.NewString(tableName, "issue_no_version_data")
	_beanReportIssues.IssueReqData = field.NewString(tableName, "issue_req_data")

	_beanReportIssues.fillFieldMap()

	return _beanReportIssues
}

type beanReportIssues struct {
	beanReportIssuesDo

	ALL                field.Asterisk
	BeanReportIssuesID field.String
	BeanReportID       field.String
	Team               field.String
	JqlVersion         field.String
	JqlNoVersion       field.String
	IssueVersionData   field.String
	IssueNoVersionData field.String
	IssueReqData       field.String

	fieldMap map[string]field.Expr
}

func (b beanReportIssues) Table(newTableName string) *beanReportIssues {
	b.beanReportIssuesDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b beanReportIssues) As(alias string) *beanReportIssues {
	b.beanReportIssuesDo.DO = *(b.beanReportIssuesDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *beanReportIssues) updateTableName(table string) *beanReportIssues {
	b.ALL = field.NewAsterisk(table)
	b.BeanReportIssuesID = field.NewString(table, "bean_report_issues_id")
	b.BeanReportID = field.NewString(table, "bean_report_id")
	b.Team = field.NewString(table, "team")
	b.JqlVersion = field.NewString(table, "jql_version")
	b.JqlNoVersion = field.NewString(table, "jql_no_version")
	b.IssueVersionData = field.NewString(table, "issue_version_data")
	b.IssueNoVersionData = field.NewString(table, "issue_no_version_data")
	b.IssueReqData = field.NewString(table, "issue_req_data")

	b.fillFieldMap()

	return b
}

func (b *beanReportIssues) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *beanReportIssues) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 8)
	b.fieldMap["bean_report_issues_id"] = b.BeanReportIssuesID
	b.fieldMap["bean_report_id"] = b.BeanReportID
	b.fieldMap["team"] = b.Team
	b.fieldMap["jql_version"] = b.JqlVersion
	b.fieldMap["jql_no_version"] = b.JqlNoVersion
	b.fieldMap["issue_version_data"] = b.IssueVersionData
	b.fieldMap["issue_no_version_data"] = b.IssueNoVersionData
	b.fieldMap["issue_req_data"] = b.IssueReqData
}

func (b beanReportIssues) clone(db *gorm.DB) beanReportIssues {
	b.beanReportIssuesDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b beanReportIssues) replaceDB(db *gorm.DB) beanReportIssues {
	b.beanReportIssuesDo.ReplaceDB(db)
	return b
}

type beanReportIssuesDo struct{ gen.DO }

type IBeanReportIssuesDo interface {
	gen.SubQuery
	Debug() IBeanReportIssuesDo
	WithContext(ctx context.Context) IBeanReportIssuesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBeanReportIssuesDo
	WriteDB() IBeanReportIssuesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBeanReportIssuesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBeanReportIssuesDo
	Not(conds ...gen.Condition) IBeanReportIssuesDo
	Or(conds ...gen.Condition) IBeanReportIssuesDo
	Select(conds ...field.Expr) IBeanReportIssuesDo
	Where(conds ...gen.Condition) IBeanReportIssuesDo
	Order(conds ...field.Expr) IBeanReportIssuesDo
	Distinct(cols ...field.Expr) IBeanReportIssuesDo
	Omit(cols ...field.Expr) IBeanReportIssuesDo
	Join(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo
	Group(cols ...field.Expr) IBeanReportIssuesDo
	Having(conds ...gen.Condition) IBeanReportIssuesDo
	Limit(limit int) IBeanReportIssuesDo
	Offset(offset int) IBeanReportIssuesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBeanReportIssuesDo
	Unscoped() IBeanReportIssuesDo
	Create(values ...*model.BeanReportIssues) error
	CreateInBatches(values []*model.BeanReportIssues, batchSize int) error
	Save(values ...*model.BeanReportIssues) error
	First() (*model.BeanReportIssues, error)
	Take() (*model.BeanReportIssues, error)
	Last() (*model.BeanReportIssues, error)
	Find() ([]*model.BeanReportIssues, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeanReportIssues, err error)
	FindInBatches(result *[]*model.BeanReportIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BeanReportIssues) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBeanReportIssuesDo
	Assign(attrs ...field.AssignExpr) IBeanReportIssuesDo
	Joins(fields ...field.RelationField) IBeanReportIssuesDo
	Preload(fields ...field.RelationField) IBeanReportIssuesDo
	FirstOrInit() (*model.BeanReportIssues, error)
	FirstOrCreate() (*model.BeanReportIssues, error)
	FindByPage(offset int, limit int) (result []*model.BeanReportIssues, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBeanReportIssuesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b beanReportIssuesDo) Debug() IBeanReportIssuesDo {
	return b.withDO(b.DO.Debug())
}

func (b beanReportIssuesDo) WithContext(ctx context.Context) IBeanReportIssuesDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b beanReportIssuesDo) ReadDB() IBeanReportIssuesDo {
	return b.Clauses(dbresolver.Read)
}

func (b beanReportIssuesDo) WriteDB() IBeanReportIssuesDo {
	return b.Clauses(dbresolver.Write)
}

func (b beanReportIssuesDo) Session(config *gorm.Session) IBeanReportIssuesDo {
	return b.withDO(b.DO.Session(config))
}

func (b beanReportIssuesDo) Clauses(conds ...clause.Expression) IBeanReportIssuesDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b beanReportIssuesDo) Returning(value interface{}, columns ...string) IBeanReportIssuesDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b beanReportIssuesDo) Not(conds ...gen.Condition) IBeanReportIssuesDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b beanReportIssuesDo) Or(conds ...gen.Condition) IBeanReportIssuesDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b beanReportIssuesDo) Select(conds ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b beanReportIssuesDo) Where(conds ...gen.Condition) IBeanReportIssuesDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b beanReportIssuesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBeanReportIssuesDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b beanReportIssuesDo) Order(conds ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b beanReportIssuesDo) Distinct(cols ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b beanReportIssuesDo) Omit(cols ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b beanReportIssuesDo) Join(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b beanReportIssuesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b beanReportIssuesDo) RightJoin(table schema.Tabler, on ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b beanReportIssuesDo) Group(cols ...field.Expr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b beanReportIssuesDo) Having(conds ...gen.Condition) IBeanReportIssuesDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b beanReportIssuesDo) Limit(limit int) IBeanReportIssuesDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b beanReportIssuesDo) Offset(offset int) IBeanReportIssuesDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b beanReportIssuesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBeanReportIssuesDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b beanReportIssuesDo) Unscoped() IBeanReportIssuesDo {
	return b.withDO(b.DO.Unscoped())
}

func (b beanReportIssuesDo) Create(values ...*model.BeanReportIssues) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b beanReportIssuesDo) CreateInBatches(values []*model.BeanReportIssues, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b beanReportIssuesDo) Save(values ...*model.BeanReportIssues) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b beanReportIssuesDo) First() (*model.BeanReportIssues, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportIssues), nil
	}
}

func (b beanReportIssuesDo) Take() (*model.BeanReportIssues, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportIssues), nil
	}
}

func (b beanReportIssuesDo) Last() (*model.BeanReportIssues, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportIssues), nil
	}
}

func (b beanReportIssuesDo) Find() ([]*model.BeanReportIssues, error) {
	result, err := b.DO.Find()
	return result.([]*model.BeanReportIssues), err
}

func (b beanReportIssuesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeanReportIssues, err error) {
	buf := make([]*model.BeanReportIssues, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b beanReportIssuesDo) FindInBatches(result *[]*model.BeanReportIssues, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b beanReportIssuesDo) Attrs(attrs ...field.AssignExpr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b beanReportIssuesDo) Assign(attrs ...field.AssignExpr) IBeanReportIssuesDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b beanReportIssuesDo) Joins(fields ...field.RelationField) IBeanReportIssuesDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b beanReportIssuesDo) Preload(fields ...field.RelationField) IBeanReportIssuesDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b beanReportIssuesDo) FirstOrInit() (*model.BeanReportIssues, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportIssues), nil
	}
}

func (b beanReportIssuesDo) FirstOrCreate() (*model.BeanReportIssues, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportIssues), nil
	}
}

func (b beanReportIssuesDo) FindByPage(offset int, limit int) (result []*model.BeanReportIssues, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b beanReportIssuesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b beanReportIssuesDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b beanReportIssuesDo) Delete(models ...*model.BeanReportIssues) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *beanReportIssuesDo) withDO(do gen.Dao) *beanReportIssuesDo {
	b.DO = *do.(*gen.DO)
	return b
}
