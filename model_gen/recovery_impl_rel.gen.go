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

func newRecoveryImplRel(db *gorm.DB, opts ...gen.DOOption) recoveryImplRel {
	_recoveryImplRel := recoveryImplRel{}

	_recoveryImplRel.recoveryImplRelDo.UseDB(db, opts...)
	_recoveryImplRel.recoveryImplRelDo.UseModel(&model.RecoveryImplRel{})

	tableName := _recoveryImplRel.recoveryImplRelDo.TableName()
	_recoveryImplRel.ALL = field.NewAsterisk(tableName)
	_recoveryImplRel.RecoveryImplRelID = field.NewString(tableName, "recovery_impl_rel_id")
	_recoveryImplRel.RecoveryImplPackageID = field.NewString(tableName, "recovery_impl_package_id")
	_recoveryImplRel.RecoveryImplID = field.NewString(tableName, "recovery_impl_id")
	_recoveryImplRel.Sort = field.NewFloat64(tableName, "sort")

	_recoveryImplRel.fillFieldMap()

	return _recoveryImplRel
}

type recoveryImplRel struct {
	recoveryImplRelDo

	ALL                   field.Asterisk
	RecoveryImplRelID     field.String // ID
	RecoveryImplPackageID field.String // ID
	RecoveryImplID        field.String // ID
	Sort                  field.Float64

	fieldMap map[string]field.Expr
}

func (r recoveryImplRel) Table(newTableName string) *recoveryImplRel {
	r.recoveryImplRelDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r recoveryImplRel) As(alias string) *recoveryImplRel {
	r.recoveryImplRelDo.DO = *(r.recoveryImplRelDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *recoveryImplRel) updateTableName(table string) *recoveryImplRel {
	r.ALL = field.NewAsterisk(table)
	r.RecoveryImplRelID = field.NewString(table, "recovery_impl_rel_id")
	r.RecoveryImplPackageID = field.NewString(table, "recovery_impl_package_id")
	r.RecoveryImplID = field.NewString(table, "recovery_impl_id")
	r.Sort = field.NewFloat64(table, "sort")

	r.fillFieldMap()

	return r
}

func (r *recoveryImplRel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *recoveryImplRel) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 4)
	r.fieldMap["recovery_impl_rel_id"] = r.RecoveryImplRelID
	r.fieldMap["recovery_impl_package_id"] = r.RecoveryImplPackageID
	r.fieldMap["recovery_impl_id"] = r.RecoveryImplID
	r.fieldMap["sort"] = r.Sort
}

func (r recoveryImplRel) clone(db *gorm.DB) recoveryImplRel {
	r.recoveryImplRelDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r recoveryImplRel) replaceDB(db *gorm.DB) recoveryImplRel {
	r.recoveryImplRelDo.ReplaceDB(db)
	return r
}

type recoveryImplRelDo struct{ gen.DO }

type IRecoveryImplRelDo interface {
	gen.SubQuery
	Debug() IRecoveryImplRelDo
	WithContext(ctx context.Context) IRecoveryImplRelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecoveryImplRelDo
	WriteDB() IRecoveryImplRelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecoveryImplRelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecoveryImplRelDo
	Not(conds ...gen.Condition) IRecoveryImplRelDo
	Or(conds ...gen.Condition) IRecoveryImplRelDo
	Select(conds ...field.Expr) IRecoveryImplRelDo
	Where(conds ...gen.Condition) IRecoveryImplRelDo
	Order(conds ...field.Expr) IRecoveryImplRelDo
	Distinct(cols ...field.Expr) IRecoveryImplRelDo
	Omit(cols ...field.Expr) IRecoveryImplRelDo
	Join(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo
	Group(cols ...field.Expr) IRecoveryImplRelDo
	Having(conds ...gen.Condition) IRecoveryImplRelDo
	Limit(limit int) IRecoveryImplRelDo
	Offset(offset int) IRecoveryImplRelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecoveryImplRelDo
	Unscoped() IRecoveryImplRelDo
	Create(values ...*model.RecoveryImplRel) error
	CreateInBatches(values []*model.RecoveryImplRel, batchSize int) error
	Save(values ...*model.RecoveryImplRel) error
	First() (*model.RecoveryImplRel, error)
	Take() (*model.RecoveryImplRel, error)
	Last() (*model.RecoveryImplRel, error)
	Find() ([]*model.RecoveryImplRel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecoveryImplRel, err error)
	FindInBatches(result *[]*model.RecoveryImplRel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RecoveryImplRel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecoveryImplRelDo
	Assign(attrs ...field.AssignExpr) IRecoveryImplRelDo
	Joins(fields ...field.RelationField) IRecoveryImplRelDo
	Preload(fields ...field.RelationField) IRecoveryImplRelDo
	FirstOrInit() (*model.RecoveryImplRel, error)
	FirstOrCreate() (*model.RecoveryImplRel, error)
	FindByPage(offset int, limit int) (result []*model.RecoveryImplRel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecoveryImplRelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recoveryImplRelDo) Debug() IRecoveryImplRelDo {
	return r.withDO(r.DO.Debug())
}

func (r recoveryImplRelDo) WithContext(ctx context.Context) IRecoveryImplRelDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recoveryImplRelDo) ReadDB() IRecoveryImplRelDo {
	return r.Clauses(dbresolver.Read)
}

func (r recoveryImplRelDo) WriteDB() IRecoveryImplRelDo {
	return r.Clauses(dbresolver.Write)
}

func (r recoveryImplRelDo) Session(config *gorm.Session) IRecoveryImplRelDo {
	return r.withDO(r.DO.Session(config))
}

func (r recoveryImplRelDo) Clauses(conds ...clause.Expression) IRecoveryImplRelDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recoveryImplRelDo) Returning(value interface{}, columns ...string) IRecoveryImplRelDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recoveryImplRelDo) Not(conds ...gen.Condition) IRecoveryImplRelDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recoveryImplRelDo) Or(conds ...gen.Condition) IRecoveryImplRelDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recoveryImplRelDo) Select(conds ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recoveryImplRelDo) Where(conds ...gen.Condition) IRecoveryImplRelDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recoveryImplRelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRecoveryImplRelDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r recoveryImplRelDo) Order(conds ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recoveryImplRelDo) Distinct(cols ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recoveryImplRelDo) Omit(cols ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recoveryImplRelDo) Join(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recoveryImplRelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recoveryImplRelDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recoveryImplRelDo) Group(cols ...field.Expr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recoveryImplRelDo) Having(conds ...gen.Condition) IRecoveryImplRelDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recoveryImplRelDo) Limit(limit int) IRecoveryImplRelDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recoveryImplRelDo) Offset(offset int) IRecoveryImplRelDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recoveryImplRelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecoveryImplRelDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recoveryImplRelDo) Unscoped() IRecoveryImplRelDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recoveryImplRelDo) Create(values ...*model.RecoveryImplRel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recoveryImplRelDo) CreateInBatches(values []*model.RecoveryImplRel, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recoveryImplRelDo) Save(values ...*model.RecoveryImplRel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recoveryImplRelDo) First() (*model.RecoveryImplRel, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplRel), nil
	}
}

func (r recoveryImplRelDo) Take() (*model.RecoveryImplRel, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplRel), nil
	}
}

func (r recoveryImplRelDo) Last() (*model.RecoveryImplRel, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplRel), nil
	}
}

func (r recoveryImplRelDo) Find() ([]*model.RecoveryImplRel, error) {
	result, err := r.DO.Find()
	return result.([]*model.RecoveryImplRel), err
}

func (r recoveryImplRelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecoveryImplRel, err error) {
	buf := make([]*model.RecoveryImplRel, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recoveryImplRelDo) FindInBatches(result *[]*model.RecoveryImplRel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recoveryImplRelDo) Attrs(attrs ...field.AssignExpr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recoveryImplRelDo) Assign(attrs ...field.AssignExpr) IRecoveryImplRelDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recoveryImplRelDo) Joins(fields ...field.RelationField) IRecoveryImplRelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recoveryImplRelDo) Preload(fields ...field.RelationField) IRecoveryImplRelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recoveryImplRelDo) FirstOrInit() (*model.RecoveryImplRel, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplRel), nil
	}
}

func (r recoveryImplRelDo) FirstOrCreate() (*model.RecoveryImplRel, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplRel), nil
	}
}

func (r recoveryImplRelDo) FindByPage(offset int, limit int) (result []*model.RecoveryImplRel, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r recoveryImplRelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recoveryImplRelDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recoveryImplRelDo) Delete(models ...*model.RecoveryImplRel) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recoveryImplRelDo) withDO(do gen.Dao) *recoveryImplRelDo {
	r.DO = *do.(*gen.DO)
	return r
}