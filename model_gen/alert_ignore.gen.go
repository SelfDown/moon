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

func newAlertIgnore(db *gorm.DB, opts ...gen.DOOption) alertIgnore {
	_alertIgnore := alertIgnore{}

	_alertIgnore.alertIgnoreDo.UseDB(db, opts...)
	_alertIgnore.alertIgnoreDo.UseModel(&model.AlertIgnore{})

	tableName := _alertIgnore.alertIgnoreDo.TableName()
	_alertIgnore.ALL = field.NewAsterisk(tableName)
	_alertIgnore.AlertIgnoreID = field.NewString(tableName, "alert_ignore_id")
	_alertIgnore.Reg = field.NewString(tableName, "reg")

	_alertIgnore.fillFieldMap()

	return _alertIgnore
}

type alertIgnore struct {
	alertIgnoreDo

	ALL           field.Asterisk
	AlertIgnoreID field.String
	Reg           field.String

	fieldMap map[string]field.Expr
}

func (a alertIgnore) Table(newTableName string) *alertIgnore {
	a.alertIgnoreDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alertIgnore) As(alias string) *alertIgnore {
	a.alertIgnoreDo.DO = *(a.alertIgnoreDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alertIgnore) updateTableName(table string) *alertIgnore {
	a.ALL = field.NewAsterisk(table)
	a.AlertIgnoreID = field.NewString(table, "alert_ignore_id")
	a.Reg = field.NewString(table, "reg")

	a.fillFieldMap()

	return a
}

func (a *alertIgnore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alertIgnore) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["alert_ignore_id"] = a.AlertIgnoreID
	a.fieldMap["reg"] = a.Reg
}

func (a alertIgnore) clone(db *gorm.DB) alertIgnore {
	a.alertIgnoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alertIgnore) replaceDB(db *gorm.DB) alertIgnore {
	a.alertIgnoreDo.ReplaceDB(db)
	return a
}

type alertIgnoreDo struct{ gen.DO }

type IAlertIgnoreDo interface {
	gen.SubQuery
	Debug() IAlertIgnoreDo
	WithContext(ctx context.Context) IAlertIgnoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAlertIgnoreDo
	WriteDB() IAlertIgnoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAlertIgnoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAlertIgnoreDo
	Not(conds ...gen.Condition) IAlertIgnoreDo
	Or(conds ...gen.Condition) IAlertIgnoreDo
	Select(conds ...field.Expr) IAlertIgnoreDo
	Where(conds ...gen.Condition) IAlertIgnoreDo
	Order(conds ...field.Expr) IAlertIgnoreDo
	Distinct(cols ...field.Expr) IAlertIgnoreDo
	Omit(cols ...field.Expr) IAlertIgnoreDo
	Join(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo
	Group(cols ...field.Expr) IAlertIgnoreDo
	Having(conds ...gen.Condition) IAlertIgnoreDo
	Limit(limit int) IAlertIgnoreDo
	Offset(offset int) IAlertIgnoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertIgnoreDo
	Unscoped() IAlertIgnoreDo
	Create(values ...*model.AlertIgnore) error
	CreateInBatches(values []*model.AlertIgnore, batchSize int) error
	Save(values ...*model.AlertIgnore) error
	First() (*model.AlertIgnore, error)
	Take() (*model.AlertIgnore, error)
	Last() (*model.AlertIgnore, error)
	Find() ([]*model.AlertIgnore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlertIgnore, err error)
	FindInBatches(result *[]*model.AlertIgnore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AlertIgnore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAlertIgnoreDo
	Assign(attrs ...field.AssignExpr) IAlertIgnoreDo
	Joins(fields ...field.RelationField) IAlertIgnoreDo
	Preload(fields ...field.RelationField) IAlertIgnoreDo
	FirstOrInit() (*model.AlertIgnore, error)
	FirstOrCreate() (*model.AlertIgnore, error)
	FindByPage(offset int, limit int) (result []*model.AlertIgnore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAlertIgnoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a alertIgnoreDo) Debug() IAlertIgnoreDo {
	return a.withDO(a.DO.Debug())
}

func (a alertIgnoreDo) WithContext(ctx context.Context) IAlertIgnoreDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alertIgnoreDo) ReadDB() IAlertIgnoreDo {
	return a.Clauses(dbresolver.Read)
}

func (a alertIgnoreDo) WriteDB() IAlertIgnoreDo {
	return a.Clauses(dbresolver.Write)
}

func (a alertIgnoreDo) Session(config *gorm.Session) IAlertIgnoreDo {
	return a.withDO(a.DO.Session(config))
}

func (a alertIgnoreDo) Clauses(conds ...clause.Expression) IAlertIgnoreDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alertIgnoreDo) Returning(value interface{}, columns ...string) IAlertIgnoreDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alertIgnoreDo) Not(conds ...gen.Condition) IAlertIgnoreDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alertIgnoreDo) Or(conds ...gen.Condition) IAlertIgnoreDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alertIgnoreDo) Select(conds ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alertIgnoreDo) Where(conds ...gen.Condition) IAlertIgnoreDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alertIgnoreDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAlertIgnoreDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a alertIgnoreDo) Order(conds ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alertIgnoreDo) Distinct(cols ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alertIgnoreDo) Omit(cols ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alertIgnoreDo) Join(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alertIgnoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alertIgnoreDo) RightJoin(table schema.Tabler, on ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alertIgnoreDo) Group(cols ...field.Expr) IAlertIgnoreDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alertIgnoreDo) Having(conds ...gen.Condition) IAlertIgnoreDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alertIgnoreDo) Limit(limit int) IAlertIgnoreDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alertIgnoreDo) Offset(offset int) IAlertIgnoreDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alertIgnoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertIgnoreDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alertIgnoreDo) Unscoped() IAlertIgnoreDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alertIgnoreDo) Create(values ...*model.AlertIgnore) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alertIgnoreDo) CreateInBatches(values []*model.AlertIgnore, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alertIgnoreDo) Save(values ...*model.AlertIgnore) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alertIgnoreDo) First() (*model.AlertIgnore, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertIgnore), nil
	}
}

func (a alertIgnoreDo) Take() (*model.AlertIgnore, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertIgnore), nil
	}
}

func (a alertIgnoreDo) Last() (*model.AlertIgnore, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertIgnore), nil
	}
}

func (a alertIgnoreDo) Find() ([]*model.AlertIgnore, error) {
	result, err := a.DO.Find()
	return result.([]*model.AlertIgnore), err
}

func (a alertIgnoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlertIgnore, err error) {
	buf := make([]*model.AlertIgnore, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alertIgnoreDo) FindInBatches(result *[]*model.AlertIgnore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alertIgnoreDo) Attrs(attrs ...field.AssignExpr) IAlertIgnoreDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alertIgnoreDo) Assign(attrs ...field.AssignExpr) IAlertIgnoreDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alertIgnoreDo) Joins(fields ...field.RelationField) IAlertIgnoreDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alertIgnoreDo) Preload(fields ...field.RelationField) IAlertIgnoreDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alertIgnoreDo) FirstOrInit() (*model.AlertIgnore, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertIgnore), nil
	}
}

func (a alertIgnoreDo) FirstOrCreate() (*model.AlertIgnore, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertIgnore), nil
	}
}

func (a alertIgnoreDo) FindByPage(offset int, limit int) (result []*model.AlertIgnore, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a alertIgnoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alertIgnoreDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alertIgnoreDo) Delete(models ...*model.AlertIgnore) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alertIgnoreDo) withDO(do gen.Dao) *alertIgnoreDo {
	a.DO = *do.(*gen.DO)
	return a
}
