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

func newFailureReportChangeList(db *gorm.DB, opts ...gen.DOOption) failureReportChangeList {
	_failureReportChangeList := failureReportChangeList{}

	_failureReportChangeList.failureReportChangeListDo.UseDB(db, opts...)
	_failureReportChangeList.failureReportChangeListDo.UseModel(&model.FailureReportChangeList{})

	tableName := _failureReportChangeList.failureReportChangeListDo.TableName()
	_failureReportChangeList.ALL = field.NewAsterisk(tableName)
	_failureReportChangeList.FailureReportChangeListID = field.NewString(tableName, "failure_report_change_list_id")
	_failureReportChangeList.FailureReportID = field.NewString(tableName, "failure_report_id")
	_failureReportChangeList.Name = field.NewString(tableName, "name")
	_failureReportChangeList.Operation = field.NewString(tableName, "operation")
	_failureReportChangeList.Before = field.NewString(tableName, "before")
	_failureReportChangeList.After = field.NewString(tableName, "after")
	_failureReportChangeList.CreateUser = field.NewString(tableName, "create_user")
	_failureReportChangeList.CreateTime = field.NewString(tableName, "create_time")

	_failureReportChangeList.fillFieldMap()

	return _failureReportChangeList
}

type failureReportChangeList struct {
	failureReportChangeListDo

	ALL                       field.Asterisk
	FailureReportChangeListID field.String
	FailureReportID           field.String
	Name                      field.String
	Operation                 field.String
	Before                    field.String
	After                     field.String
	CreateUser                field.String
	CreateTime                field.String

	fieldMap map[string]field.Expr
}

func (f failureReportChangeList) Table(newTableName string) *failureReportChangeList {
	f.failureReportChangeListDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f failureReportChangeList) As(alias string) *failureReportChangeList {
	f.failureReportChangeListDo.DO = *(f.failureReportChangeListDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *failureReportChangeList) updateTableName(table string) *failureReportChangeList {
	f.ALL = field.NewAsterisk(table)
	f.FailureReportChangeListID = field.NewString(table, "failure_report_change_list_id")
	f.FailureReportID = field.NewString(table, "failure_report_id")
	f.Name = field.NewString(table, "name")
	f.Operation = field.NewString(table, "operation")
	f.Before = field.NewString(table, "before")
	f.After = field.NewString(table, "after")
	f.CreateUser = field.NewString(table, "create_user")
	f.CreateTime = field.NewString(table, "create_time")

	f.fillFieldMap()

	return f
}

func (f *failureReportChangeList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *failureReportChangeList) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["failure_report_change_list_id"] = f.FailureReportChangeListID
	f.fieldMap["failure_report_id"] = f.FailureReportID
	f.fieldMap["name"] = f.Name
	f.fieldMap["operation"] = f.Operation
	f.fieldMap["before"] = f.Before
	f.fieldMap["after"] = f.After
	f.fieldMap["create_user"] = f.CreateUser
	f.fieldMap["create_time"] = f.CreateTime
}

func (f failureReportChangeList) clone(db *gorm.DB) failureReportChangeList {
	f.failureReportChangeListDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f failureReportChangeList) replaceDB(db *gorm.DB) failureReportChangeList {
	f.failureReportChangeListDo.ReplaceDB(db)
	return f
}

type failureReportChangeListDo struct{ gen.DO }

type IFailureReportChangeListDo interface {
	gen.SubQuery
	Debug() IFailureReportChangeListDo
	WithContext(ctx context.Context) IFailureReportChangeListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFailureReportChangeListDo
	WriteDB() IFailureReportChangeListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFailureReportChangeListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFailureReportChangeListDo
	Not(conds ...gen.Condition) IFailureReportChangeListDo
	Or(conds ...gen.Condition) IFailureReportChangeListDo
	Select(conds ...field.Expr) IFailureReportChangeListDo
	Where(conds ...gen.Condition) IFailureReportChangeListDo
	Order(conds ...field.Expr) IFailureReportChangeListDo
	Distinct(cols ...field.Expr) IFailureReportChangeListDo
	Omit(cols ...field.Expr) IFailureReportChangeListDo
	Join(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo
	Group(cols ...field.Expr) IFailureReportChangeListDo
	Having(conds ...gen.Condition) IFailureReportChangeListDo
	Limit(limit int) IFailureReportChangeListDo
	Offset(offset int) IFailureReportChangeListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFailureReportChangeListDo
	Unscoped() IFailureReportChangeListDo
	Create(values ...*model.FailureReportChangeList) error
	CreateInBatches(values []*model.FailureReportChangeList, batchSize int) error
	Save(values ...*model.FailureReportChangeList) error
	First() (*model.FailureReportChangeList, error)
	Take() (*model.FailureReportChangeList, error)
	Last() (*model.FailureReportChangeList, error)
	Find() ([]*model.FailureReportChangeList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FailureReportChangeList, err error)
	FindInBatches(result *[]*model.FailureReportChangeList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FailureReportChangeList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFailureReportChangeListDo
	Assign(attrs ...field.AssignExpr) IFailureReportChangeListDo
	Joins(fields ...field.RelationField) IFailureReportChangeListDo
	Preload(fields ...field.RelationField) IFailureReportChangeListDo
	FirstOrInit() (*model.FailureReportChangeList, error)
	FirstOrCreate() (*model.FailureReportChangeList, error)
	FindByPage(offset int, limit int) (result []*model.FailureReportChangeList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFailureReportChangeListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f failureReportChangeListDo) Debug() IFailureReportChangeListDo {
	return f.withDO(f.DO.Debug())
}

func (f failureReportChangeListDo) WithContext(ctx context.Context) IFailureReportChangeListDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f failureReportChangeListDo) ReadDB() IFailureReportChangeListDo {
	return f.Clauses(dbresolver.Read)
}

func (f failureReportChangeListDo) WriteDB() IFailureReportChangeListDo {
	return f.Clauses(dbresolver.Write)
}

func (f failureReportChangeListDo) Session(config *gorm.Session) IFailureReportChangeListDo {
	return f.withDO(f.DO.Session(config))
}

func (f failureReportChangeListDo) Clauses(conds ...clause.Expression) IFailureReportChangeListDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f failureReportChangeListDo) Returning(value interface{}, columns ...string) IFailureReportChangeListDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f failureReportChangeListDo) Not(conds ...gen.Condition) IFailureReportChangeListDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f failureReportChangeListDo) Or(conds ...gen.Condition) IFailureReportChangeListDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f failureReportChangeListDo) Select(conds ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f failureReportChangeListDo) Where(conds ...gen.Condition) IFailureReportChangeListDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f failureReportChangeListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFailureReportChangeListDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f failureReportChangeListDo) Order(conds ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f failureReportChangeListDo) Distinct(cols ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f failureReportChangeListDo) Omit(cols ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f failureReportChangeListDo) Join(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f failureReportChangeListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f failureReportChangeListDo) RightJoin(table schema.Tabler, on ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f failureReportChangeListDo) Group(cols ...field.Expr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f failureReportChangeListDo) Having(conds ...gen.Condition) IFailureReportChangeListDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f failureReportChangeListDo) Limit(limit int) IFailureReportChangeListDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f failureReportChangeListDo) Offset(offset int) IFailureReportChangeListDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f failureReportChangeListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFailureReportChangeListDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f failureReportChangeListDo) Unscoped() IFailureReportChangeListDo {
	return f.withDO(f.DO.Unscoped())
}

func (f failureReportChangeListDo) Create(values ...*model.FailureReportChangeList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f failureReportChangeListDo) CreateInBatches(values []*model.FailureReportChangeList, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f failureReportChangeListDo) Save(values ...*model.FailureReportChangeList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f failureReportChangeListDo) First() (*model.FailureReportChangeList, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FailureReportChangeList), nil
	}
}

func (f failureReportChangeListDo) Take() (*model.FailureReportChangeList, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FailureReportChangeList), nil
	}
}

func (f failureReportChangeListDo) Last() (*model.FailureReportChangeList, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FailureReportChangeList), nil
	}
}

func (f failureReportChangeListDo) Find() ([]*model.FailureReportChangeList, error) {
	result, err := f.DO.Find()
	return result.([]*model.FailureReportChangeList), err
}

func (f failureReportChangeListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FailureReportChangeList, err error) {
	buf := make([]*model.FailureReportChangeList, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f failureReportChangeListDo) FindInBatches(result *[]*model.FailureReportChangeList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f failureReportChangeListDo) Attrs(attrs ...field.AssignExpr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f failureReportChangeListDo) Assign(attrs ...field.AssignExpr) IFailureReportChangeListDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f failureReportChangeListDo) Joins(fields ...field.RelationField) IFailureReportChangeListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f failureReportChangeListDo) Preload(fields ...field.RelationField) IFailureReportChangeListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f failureReportChangeListDo) FirstOrInit() (*model.FailureReportChangeList, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FailureReportChangeList), nil
	}
}

func (f failureReportChangeListDo) FirstOrCreate() (*model.FailureReportChangeList, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FailureReportChangeList), nil
	}
}

func (f failureReportChangeListDo) FindByPage(offset int, limit int) (result []*model.FailureReportChangeList, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f failureReportChangeListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f failureReportChangeListDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f failureReportChangeListDo) Delete(models ...*model.FailureReportChangeList) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *failureReportChangeListDo) withDO(do gen.Dao) *failureReportChangeListDo {
	f.DO = *do.(*gen.DO)
	return f
}
