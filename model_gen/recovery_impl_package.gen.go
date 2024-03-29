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

func newRecoveryImplPackage(db *gorm.DB, opts ...gen.DOOption) recoveryImplPackage {
	_recoveryImplPackage := recoveryImplPackage{}

	_recoveryImplPackage.recoveryImplPackageDo.UseDB(db, opts...)
	_recoveryImplPackage.recoveryImplPackageDo.UseModel(&model.RecoveryImplPackage{})

	tableName := _recoveryImplPackage.recoveryImplPackageDo.TableName()
	_recoveryImplPackage.ALL = field.NewAsterisk(tableName)
	_recoveryImplPackage.RecoveryImplPackageID = field.NewString(tableName, "recovery_impl_package_id")
	_recoveryImplPackage.RecoveryImplPackageName = field.NewString(tableName, "recovery_impl_package_name")
	_recoveryImplPackage.Note = field.NewString(tableName, "note")

	_recoveryImplPackage.fillFieldMap()

	return _recoveryImplPackage
}

type recoveryImplPackage struct {
	recoveryImplPackageDo

	ALL                     field.Asterisk
	RecoveryImplPackageID   field.String // ID
	RecoveryImplPackageName field.String // 恢复实现方法名称
	Note                    field.String // 备注

	fieldMap map[string]field.Expr
}

func (r recoveryImplPackage) Table(newTableName string) *recoveryImplPackage {
	r.recoveryImplPackageDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r recoveryImplPackage) As(alias string) *recoveryImplPackage {
	r.recoveryImplPackageDo.DO = *(r.recoveryImplPackageDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *recoveryImplPackage) updateTableName(table string) *recoveryImplPackage {
	r.ALL = field.NewAsterisk(table)
	r.RecoveryImplPackageID = field.NewString(table, "recovery_impl_package_id")
	r.RecoveryImplPackageName = field.NewString(table, "recovery_impl_package_name")
	r.Note = field.NewString(table, "note")

	r.fillFieldMap()

	return r
}

func (r *recoveryImplPackage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *recoveryImplPackage) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["recovery_impl_package_id"] = r.RecoveryImplPackageID
	r.fieldMap["recovery_impl_package_name"] = r.RecoveryImplPackageName
	r.fieldMap["note"] = r.Note
}

func (r recoveryImplPackage) clone(db *gorm.DB) recoveryImplPackage {
	r.recoveryImplPackageDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r recoveryImplPackage) replaceDB(db *gorm.DB) recoveryImplPackage {
	r.recoveryImplPackageDo.ReplaceDB(db)
	return r
}

type recoveryImplPackageDo struct{ gen.DO }

type IRecoveryImplPackageDo interface {
	gen.SubQuery
	Debug() IRecoveryImplPackageDo
	WithContext(ctx context.Context) IRecoveryImplPackageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecoveryImplPackageDo
	WriteDB() IRecoveryImplPackageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecoveryImplPackageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecoveryImplPackageDo
	Not(conds ...gen.Condition) IRecoveryImplPackageDo
	Or(conds ...gen.Condition) IRecoveryImplPackageDo
	Select(conds ...field.Expr) IRecoveryImplPackageDo
	Where(conds ...gen.Condition) IRecoveryImplPackageDo
	Order(conds ...field.Expr) IRecoveryImplPackageDo
	Distinct(cols ...field.Expr) IRecoveryImplPackageDo
	Omit(cols ...field.Expr) IRecoveryImplPackageDo
	Join(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo
	Group(cols ...field.Expr) IRecoveryImplPackageDo
	Having(conds ...gen.Condition) IRecoveryImplPackageDo
	Limit(limit int) IRecoveryImplPackageDo
	Offset(offset int) IRecoveryImplPackageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecoveryImplPackageDo
	Unscoped() IRecoveryImplPackageDo
	Create(values ...*model.RecoveryImplPackage) error
	CreateInBatches(values []*model.RecoveryImplPackage, batchSize int) error
	Save(values ...*model.RecoveryImplPackage) error
	First() (*model.RecoveryImplPackage, error)
	Take() (*model.RecoveryImplPackage, error)
	Last() (*model.RecoveryImplPackage, error)
	Find() ([]*model.RecoveryImplPackage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecoveryImplPackage, err error)
	FindInBatches(result *[]*model.RecoveryImplPackage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RecoveryImplPackage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecoveryImplPackageDo
	Assign(attrs ...field.AssignExpr) IRecoveryImplPackageDo
	Joins(fields ...field.RelationField) IRecoveryImplPackageDo
	Preload(fields ...field.RelationField) IRecoveryImplPackageDo
	FirstOrInit() (*model.RecoveryImplPackage, error)
	FirstOrCreate() (*model.RecoveryImplPackage, error)
	FindByPage(offset int, limit int) (result []*model.RecoveryImplPackage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecoveryImplPackageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recoveryImplPackageDo) Debug() IRecoveryImplPackageDo {
	return r.withDO(r.DO.Debug())
}

func (r recoveryImplPackageDo) WithContext(ctx context.Context) IRecoveryImplPackageDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recoveryImplPackageDo) ReadDB() IRecoveryImplPackageDo {
	return r.Clauses(dbresolver.Read)
}

func (r recoveryImplPackageDo) WriteDB() IRecoveryImplPackageDo {
	return r.Clauses(dbresolver.Write)
}

func (r recoveryImplPackageDo) Session(config *gorm.Session) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Session(config))
}

func (r recoveryImplPackageDo) Clauses(conds ...clause.Expression) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recoveryImplPackageDo) Returning(value interface{}, columns ...string) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recoveryImplPackageDo) Not(conds ...gen.Condition) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recoveryImplPackageDo) Or(conds ...gen.Condition) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recoveryImplPackageDo) Select(conds ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recoveryImplPackageDo) Where(conds ...gen.Condition) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recoveryImplPackageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRecoveryImplPackageDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r recoveryImplPackageDo) Order(conds ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recoveryImplPackageDo) Distinct(cols ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recoveryImplPackageDo) Omit(cols ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recoveryImplPackageDo) Join(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recoveryImplPackageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recoveryImplPackageDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recoveryImplPackageDo) Group(cols ...field.Expr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recoveryImplPackageDo) Having(conds ...gen.Condition) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recoveryImplPackageDo) Limit(limit int) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recoveryImplPackageDo) Offset(offset int) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recoveryImplPackageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recoveryImplPackageDo) Unscoped() IRecoveryImplPackageDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recoveryImplPackageDo) Create(values ...*model.RecoveryImplPackage) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recoveryImplPackageDo) CreateInBatches(values []*model.RecoveryImplPackage, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recoveryImplPackageDo) Save(values ...*model.RecoveryImplPackage) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recoveryImplPackageDo) First() (*model.RecoveryImplPackage, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplPackage), nil
	}
}

func (r recoveryImplPackageDo) Take() (*model.RecoveryImplPackage, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplPackage), nil
	}
}

func (r recoveryImplPackageDo) Last() (*model.RecoveryImplPackage, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplPackage), nil
	}
}

func (r recoveryImplPackageDo) Find() ([]*model.RecoveryImplPackage, error) {
	result, err := r.DO.Find()
	return result.([]*model.RecoveryImplPackage), err
}

func (r recoveryImplPackageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecoveryImplPackage, err error) {
	buf := make([]*model.RecoveryImplPackage, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recoveryImplPackageDo) FindInBatches(result *[]*model.RecoveryImplPackage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recoveryImplPackageDo) Attrs(attrs ...field.AssignExpr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recoveryImplPackageDo) Assign(attrs ...field.AssignExpr) IRecoveryImplPackageDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recoveryImplPackageDo) Joins(fields ...field.RelationField) IRecoveryImplPackageDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recoveryImplPackageDo) Preload(fields ...field.RelationField) IRecoveryImplPackageDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recoveryImplPackageDo) FirstOrInit() (*model.RecoveryImplPackage, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplPackage), nil
	}
}

func (r recoveryImplPackageDo) FirstOrCreate() (*model.RecoveryImplPackage, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecoveryImplPackage), nil
	}
}

func (r recoveryImplPackageDo) FindByPage(offset int, limit int) (result []*model.RecoveryImplPackage, count int64, err error) {
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

func (r recoveryImplPackageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recoveryImplPackageDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recoveryImplPackageDo) Delete(models ...*model.RecoveryImplPackage) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recoveryImplPackageDo) withDO(do gen.Dao) *recoveryImplPackageDo {
	r.DO = *do.(*gen.DO)
	return r
}
