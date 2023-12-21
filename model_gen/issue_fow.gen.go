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

func newIssueFow(db *gorm.DB, opts ...gen.DOOption) issueFow {
	_issueFow := issueFow{}

	_issueFow.issueFowDo.UseDB(db, opts...)
	_issueFow.issueFowDo.UseModel(&model.IssueFow{})

	tableName := _issueFow.issueFowDo.TableName()
	_issueFow.ALL = field.NewAsterisk(tableName)
	_issueFow.IssueFowID = field.NewString(tableName, "issue_fow_id")
	_issueFow.IssueStatu = field.NewString(tableName, "issue_statu")
	_issueFow.Weight = field.NewInt32(tableName, "weight")

	_issueFow.fillFieldMap()

	return _issueFow
}

type issueFow struct {
	issueFowDo

	ALL        field.Asterisk
	IssueFowID field.String
	IssueStatu field.String
	Weight     field.Int32

	fieldMap map[string]field.Expr
}

func (i issueFow) Table(newTableName string) *issueFow {
	i.issueFowDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i issueFow) As(alias string) *issueFow {
	i.issueFowDo.DO = *(i.issueFowDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *issueFow) updateTableName(table string) *issueFow {
	i.ALL = field.NewAsterisk(table)
	i.IssueFowID = field.NewString(table, "issue_fow_id")
	i.IssueStatu = field.NewString(table, "issue_statu")
	i.Weight = field.NewInt32(table, "weight")

	i.fillFieldMap()

	return i
}

func (i *issueFow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *issueFow) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 3)
	i.fieldMap["issue_fow_id"] = i.IssueFowID
	i.fieldMap["issue_statu"] = i.IssueStatu
	i.fieldMap["weight"] = i.Weight
}

func (i issueFow) clone(db *gorm.DB) issueFow {
	i.issueFowDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i issueFow) replaceDB(db *gorm.DB) issueFow {
	i.issueFowDo.ReplaceDB(db)
	return i
}

type issueFowDo struct{ gen.DO }

type IIssueFowDo interface {
	gen.SubQuery
	Debug() IIssueFowDo
	WithContext(ctx context.Context) IIssueFowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIssueFowDo
	WriteDB() IIssueFowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIssueFowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIssueFowDo
	Not(conds ...gen.Condition) IIssueFowDo
	Or(conds ...gen.Condition) IIssueFowDo
	Select(conds ...field.Expr) IIssueFowDo
	Where(conds ...gen.Condition) IIssueFowDo
	Order(conds ...field.Expr) IIssueFowDo
	Distinct(cols ...field.Expr) IIssueFowDo
	Omit(cols ...field.Expr) IIssueFowDo
	Join(table schema.Tabler, on ...field.Expr) IIssueFowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIssueFowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIssueFowDo
	Group(cols ...field.Expr) IIssueFowDo
	Having(conds ...gen.Condition) IIssueFowDo
	Limit(limit int) IIssueFowDo
	Offset(offset int) IIssueFowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueFowDo
	Unscoped() IIssueFowDo
	Create(values ...*model.IssueFow) error
	CreateInBatches(values []*model.IssueFow, batchSize int) error
	Save(values ...*model.IssueFow) error
	First() (*model.IssueFow, error)
	Take() (*model.IssueFow, error)
	Last() (*model.IssueFow, error)
	Find() ([]*model.IssueFow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueFow, err error)
	FindInBatches(result *[]*model.IssueFow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IssueFow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIssueFowDo
	Assign(attrs ...field.AssignExpr) IIssueFowDo
	Joins(fields ...field.RelationField) IIssueFowDo
	Preload(fields ...field.RelationField) IIssueFowDo
	FirstOrInit() (*model.IssueFow, error)
	FirstOrCreate() (*model.IssueFow, error)
	FindByPage(offset int, limit int) (result []*model.IssueFow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIssueFowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i issueFowDo) Debug() IIssueFowDo {
	return i.withDO(i.DO.Debug())
}

func (i issueFowDo) WithContext(ctx context.Context) IIssueFowDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i issueFowDo) ReadDB() IIssueFowDo {
	return i.Clauses(dbresolver.Read)
}

func (i issueFowDo) WriteDB() IIssueFowDo {
	return i.Clauses(dbresolver.Write)
}

func (i issueFowDo) Session(config *gorm.Session) IIssueFowDo {
	return i.withDO(i.DO.Session(config))
}

func (i issueFowDo) Clauses(conds ...clause.Expression) IIssueFowDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i issueFowDo) Returning(value interface{}, columns ...string) IIssueFowDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i issueFowDo) Not(conds ...gen.Condition) IIssueFowDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i issueFowDo) Or(conds ...gen.Condition) IIssueFowDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i issueFowDo) Select(conds ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i issueFowDo) Where(conds ...gen.Condition) IIssueFowDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i issueFowDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IIssueFowDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i issueFowDo) Order(conds ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i issueFowDo) Distinct(cols ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i issueFowDo) Omit(cols ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i issueFowDo) Join(table schema.Tabler, on ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i issueFowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i issueFowDo) RightJoin(table schema.Tabler, on ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i issueFowDo) Group(cols ...field.Expr) IIssueFowDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i issueFowDo) Having(conds ...gen.Condition) IIssueFowDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i issueFowDo) Limit(limit int) IIssueFowDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i issueFowDo) Offset(offset int) IIssueFowDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i issueFowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueFowDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i issueFowDo) Unscoped() IIssueFowDo {
	return i.withDO(i.DO.Unscoped())
}

func (i issueFowDo) Create(values ...*model.IssueFow) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i issueFowDo) CreateInBatches(values []*model.IssueFow, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i issueFowDo) Save(values ...*model.IssueFow) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i issueFowDo) First() (*model.IssueFow, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFow), nil
	}
}

func (i issueFowDo) Take() (*model.IssueFow, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFow), nil
	}
}

func (i issueFowDo) Last() (*model.IssueFow, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFow), nil
	}
}

func (i issueFowDo) Find() ([]*model.IssueFow, error) {
	result, err := i.DO.Find()
	return result.([]*model.IssueFow), err
}

func (i issueFowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueFow, err error) {
	buf := make([]*model.IssueFow, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i issueFowDo) FindInBatches(result *[]*model.IssueFow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i issueFowDo) Attrs(attrs ...field.AssignExpr) IIssueFowDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i issueFowDo) Assign(attrs ...field.AssignExpr) IIssueFowDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i issueFowDo) Joins(fields ...field.RelationField) IIssueFowDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i issueFowDo) Preload(fields ...field.RelationField) IIssueFowDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i issueFowDo) FirstOrInit() (*model.IssueFow, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFow), nil
	}
}

func (i issueFowDo) FirstOrCreate() (*model.IssueFow, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFow), nil
	}
}

func (i issueFowDo) FindByPage(offset int, limit int) (result []*model.IssueFow, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i issueFowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i issueFowDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i issueFowDo) Delete(models ...*model.IssueFow) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *issueFowDo) withDO(do gen.Dao) *issueFowDo {
	i.DO = *do.(*gen.DO)
	return i
}
