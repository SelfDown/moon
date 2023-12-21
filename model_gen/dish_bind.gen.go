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

func newDishBind(db *gorm.DB, opts ...gen.DOOption) dishBind {
	_dishBind := dishBind{}

	_dishBind.dishBindDo.UseDB(db, opts...)
	_dishBind.dishBindDo.UseModel(&model.DishBind{})

	tableName := _dishBind.dishBindDo.TableName()
	_dishBind.ALL = field.NewAsterisk(tableName)
	_dishBind.DishBindID = field.NewString(tableName, "dish_bind_id")
	_dishBind.DishTypeID = field.NewString(tableName, "dish_type_id")
	_dishBind.BindType = field.NewString(tableName, "bind_type")
	_dishBind.ServerID = field.NewString(tableName, "server_id")
	_dishBind.ProjectID = field.NewString(tableName, "project_id")

	_dishBind.fillFieldMap()

	return _dishBind
}

type dishBind struct {
	dishBindDo

	ALL        field.Asterisk
	DishBindID field.String
	DishTypeID field.String
	BindType   field.String
	ServerID   field.String
	ProjectID  field.String

	fieldMap map[string]field.Expr
}

func (d dishBind) Table(newTableName string) *dishBind {
	d.dishBindDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dishBind) As(alias string) *dishBind {
	d.dishBindDo.DO = *(d.dishBindDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dishBind) updateTableName(table string) *dishBind {
	d.ALL = field.NewAsterisk(table)
	d.DishBindID = field.NewString(table, "dish_bind_id")
	d.DishTypeID = field.NewString(table, "dish_type_id")
	d.BindType = field.NewString(table, "bind_type")
	d.ServerID = field.NewString(table, "server_id")
	d.ProjectID = field.NewString(table, "project_id")

	d.fillFieldMap()

	return d
}

func (d *dishBind) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dishBind) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 5)
	d.fieldMap["dish_bind_id"] = d.DishBindID
	d.fieldMap["dish_type_id"] = d.DishTypeID
	d.fieldMap["bind_type"] = d.BindType
	d.fieldMap["server_id"] = d.ServerID
	d.fieldMap["project_id"] = d.ProjectID
}

func (d dishBind) clone(db *gorm.DB) dishBind {
	d.dishBindDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dishBind) replaceDB(db *gorm.DB) dishBind {
	d.dishBindDo.ReplaceDB(db)
	return d
}

type dishBindDo struct{ gen.DO }

type IDishBindDo interface {
	gen.SubQuery
	Debug() IDishBindDo
	WithContext(ctx context.Context) IDishBindDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDishBindDo
	WriteDB() IDishBindDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDishBindDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDishBindDo
	Not(conds ...gen.Condition) IDishBindDo
	Or(conds ...gen.Condition) IDishBindDo
	Select(conds ...field.Expr) IDishBindDo
	Where(conds ...gen.Condition) IDishBindDo
	Order(conds ...field.Expr) IDishBindDo
	Distinct(cols ...field.Expr) IDishBindDo
	Omit(cols ...field.Expr) IDishBindDo
	Join(table schema.Tabler, on ...field.Expr) IDishBindDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDishBindDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDishBindDo
	Group(cols ...field.Expr) IDishBindDo
	Having(conds ...gen.Condition) IDishBindDo
	Limit(limit int) IDishBindDo
	Offset(offset int) IDishBindDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDishBindDo
	Unscoped() IDishBindDo
	Create(values ...*model.DishBind) error
	CreateInBatches(values []*model.DishBind, batchSize int) error
	Save(values ...*model.DishBind) error
	First() (*model.DishBind, error)
	Take() (*model.DishBind, error)
	Last() (*model.DishBind, error)
	Find() ([]*model.DishBind, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DishBind, err error)
	FindInBatches(result *[]*model.DishBind, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DishBind) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDishBindDo
	Assign(attrs ...field.AssignExpr) IDishBindDo
	Joins(fields ...field.RelationField) IDishBindDo
	Preload(fields ...field.RelationField) IDishBindDo
	FirstOrInit() (*model.DishBind, error)
	FirstOrCreate() (*model.DishBind, error)
	FindByPage(offset int, limit int) (result []*model.DishBind, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDishBindDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dishBindDo) Debug() IDishBindDo {
	return d.withDO(d.DO.Debug())
}

func (d dishBindDo) WithContext(ctx context.Context) IDishBindDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dishBindDo) ReadDB() IDishBindDo {
	return d.Clauses(dbresolver.Read)
}

func (d dishBindDo) WriteDB() IDishBindDo {
	return d.Clauses(dbresolver.Write)
}

func (d dishBindDo) Session(config *gorm.Session) IDishBindDo {
	return d.withDO(d.DO.Session(config))
}

func (d dishBindDo) Clauses(conds ...clause.Expression) IDishBindDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dishBindDo) Returning(value interface{}, columns ...string) IDishBindDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dishBindDo) Not(conds ...gen.Condition) IDishBindDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dishBindDo) Or(conds ...gen.Condition) IDishBindDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dishBindDo) Select(conds ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dishBindDo) Where(conds ...gen.Condition) IDishBindDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dishBindDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDishBindDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dishBindDo) Order(conds ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dishBindDo) Distinct(cols ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dishBindDo) Omit(cols ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dishBindDo) Join(table schema.Tabler, on ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dishBindDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dishBindDo) RightJoin(table schema.Tabler, on ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dishBindDo) Group(cols ...field.Expr) IDishBindDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dishBindDo) Having(conds ...gen.Condition) IDishBindDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dishBindDo) Limit(limit int) IDishBindDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dishBindDo) Offset(offset int) IDishBindDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dishBindDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDishBindDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dishBindDo) Unscoped() IDishBindDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dishBindDo) Create(values ...*model.DishBind) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dishBindDo) CreateInBatches(values []*model.DishBind, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dishBindDo) Save(values ...*model.DishBind) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dishBindDo) First() (*model.DishBind, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DishBind), nil
	}
}

func (d dishBindDo) Take() (*model.DishBind, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DishBind), nil
	}
}

func (d dishBindDo) Last() (*model.DishBind, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DishBind), nil
	}
}

func (d dishBindDo) Find() ([]*model.DishBind, error) {
	result, err := d.DO.Find()
	return result.([]*model.DishBind), err
}

func (d dishBindDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DishBind, err error) {
	buf := make([]*model.DishBind, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dishBindDo) FindInBatches(result *[]*model.DishBind, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dishBindDo) Attrs(attrs ...field.AssignExpr) IDishBindDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dishBindDo) Assign(attrs ...field.AssignExpr) IDishBindDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dishBindDo) Joins(fields ...field.RelationField) IDishBindDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dishBindDo) Preload(fields ...field.RelationField) IDishBindDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dishBindDo) FirstOrInit() (*model.DishBind, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DishBind), nil
	}
}

func (d dishBindDo) FirstOrCreate() (*model.DishBind, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DishBind), nil
	}
}

func (d dishBindDo) FindByPage(offset int, limit int) (result []*model.DishBind, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dishBindDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dishBindDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dishBindDo) Delete(models ...*model.DishBind) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dishBindDo) withDO(do gen.Dao) *dishBindDo {
	d.DO = *do.(*gen.DO)
	return d
}
