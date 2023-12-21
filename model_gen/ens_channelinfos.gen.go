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

func newEnsChannelinfos(db *gorm.DB, opts ...gen.DOOption) ensChannelinfos {
	_ensChannelinfos := ensChannelinfos{}

	_ensChannelinfos.ensChannelinfosDo.UseDB(db, opts...)
	_ensChannelinfos.ensChannelinfosDo.UseModel(&model.EnsChannelinfos{})

	tableName := _ensChannelinfos.ensChannelinfosDo.TableName()
	_ensChannelinfos.ALL = field.NewAsterisk(tableName)
	_ensChannelinfos.ID = field.NewString(tableName, "id")
	_ensChannelinfos.Name = field.NewString(tableName, "name")
	_ensChannelinfos.Type = field.NewString(tableName, "type")
	_ensChannelinfos.Status = field.NewString(tableName, "status")
	_ensChannelinfos.Adapterstate = field.NewString(tableName, "adapterstate")
	_ensChannelinfos.Lastactivity = field.NewTime(tableName, "lastactivity")
	_ensChannelinfos.Elapsedtime = field.NewString(tableName, "elapsedtime")
	_ensChannelinfos.Queue = field.NewString(tableName, "queue")
	_ensChannelinfos.Count_ = field.NewString(tableName, "count")
	_ensChannelinfos.Addtime = field.NewTime(tableName, "addtime")
	_ensChannelinfos.Project = field.NewString(tableName, "project")

	_ensChannelinfos.fillFieldMap()

	return _ensChannelinfos
}

type ensChannelinfos struct {
	ensChannelinfosDo

	ALL          field.Asterisk
	ID           field.String
	Name         field.String
	Type         field.String
	Status       field.String
	Adapterstate field.String
	Lastactivity field.Time
	Elapsedtime  field.String
	Queue        field.String
	Count_       field.String
	Addtime      field.Time
	Project      field.String

	fieldMap map[string]field.Expr
}

func (e ensChannelinfos) Table(newTableName string) *ensChannelinfos {
	e.ensChannelinfosDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ensChannelinfos) As(alias string) *ensChannelinfos {
	e.ensChannelinfosDo.DO = *(e.ensChannelinfosDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ensChannelinfos) updateTableName(table string) *ensChannelinfos {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewString(table, "id")
	e.Name = field.NewString(table, "name")
	e.Type = field.NewString(table, "type")
	e.Status = field.NewString(table, "status")
	e.Adapterstate = field.NewString(table, "adapterstate")
	e.Lastactivity = field.NewTime(table, "lastactivity")
	e.Elapsedtime = field.NewString(table, "elapsedtime")
	e.Queue = field.NewString(table, "queue")
	e.Count_ = field.NewString(table, "count")
	e.Addtime = field.NewTime(table, "addtime")
	e.Project = field.NewString(table, "project")

	e.fillFieldMap()

	return e
}

func (e *ensChannelinfos) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ensChannelinfos) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 11)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["type"] = e.Type
	e.fieldMap["status"] = e.Status
	e.fieldMap["adapterstate"] = e.Adapterstate
	e.fieldMap["lastactivity"] = e.Lastactivity
	e.fieldMap["elapsedtime"] = e.Elapsedtime
	e.fieldMap["queue"] = e.Queue
	e.fieldMap["count"] = e.Count_
	e.fieldMap["addtime"] = e.Addtime
	e.fieldMap["project"] = e.Project
}

func (e ensChannelinfos) clone(db *gorm.DB) ensChannelinfos {
	e.ensChannelinfosDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ensChannelinfos) replaceDB(db *gorm.DB) ensChannelinfos {
	e.ensChannelinfosDo.ReplaceDB(db)
	return e
}

type ensChannelinfosDo struct{ gen.DO }

type IEnsChannelinfosDo interface {
	gen.SubQuery
	Debug() IEnsChannelinfosDo
	WithContext(ctx context.Context) IEnsChannelinfosDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEnsChannelinfosDo
	WriteDB() IEnsChannelinfosDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEnsChannelinfosDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEnsChannelinfosDo
	Not(conds ...gen.Condition) IEnsChannelinfosDo
	Or(conds ...gen.Condition) IEnsChannelinfosDo
	Select(conds ...field.Expr) IEnsChannelinfosDo
	Where(conds ...gen.Condition) IEnsChannelinfosDo
	Order(conds ...field.Expr) IEnsChannelinfosDo
	Distinct(cols ...field.Expr) IEnsChannelinfosDo
	Omit(cols ...field.Expr) IEnsChannelinfosDo
	Join(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo
	Group(cols ...field.Expr) IEnsChannelinfosDo
	Having(conds ...gen.Condition) IEnsChannelinfosDo
	Limit(limit int) IEnsChannelinfosDo
	Offset(offset int) IEnsChannelinfosDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEnsChannelinfosDo
	Unscoped() IEnsChannelinfosDo
	Create(values ...*model.EnsChannelinfos) error
	CreateInBatches(values []*model.EnsChannelinfos, batchSize int) error
	Save(values ...*model.EnsChannelinfos) error
	First() (*model.EnsChannelinfos, error)
	Take() (*model.EnsChannelinfos, error)
	Last() (*model.EnsChannelinfos, error)
	Find() ([]*model.EnsChannelinfos, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnsChannelinfos, err error)
	FindInBatches(result *[]*model.EnsChannelinfos, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EnsChannelinfos) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEnsChannelinfosDo
	Assign(attrs ...field.AssignExpr) IEnsChannelinfosDo
	Joins(fields ...field.RelationField) IEnsChannelinfosDo
	Preload(fields ...field.RelationField) IEnsChannelinfosDo
	FirstOrInit() (*model.EnsChannelinfos, error)
	FirstOrCreate() (*model.EnsChannelinfos, error)
	FindByPage(offset int, limit int) (result []*model.EnsChannelinfos, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEnsChannelinfosDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ensChannelinfosDo) Debug() IEnsChannelinfosDo {
	return e.withDO(e.DO.Debug())
}

func (e ensChannelinfosDo) WithContext(ctx context.Context) IEnsChannelinfosDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ensChannelinfosDo) ReadDB() IEnsChannelinfosDo {
	return e.Clauses(dbresolver.Read)
}

func (e ensChannelinfosDo) WriteDB() IEnsChannelinfosDo {
	return e.Clauses(dbresolver.Write)
}

func (e ensChannelinfosDo) Session(config *gorm.Session) IEnsChannelinfosDo {
	return e.withDO(e.DO.Session(config))
}

func (e ensChannelinfosDo) Clauses(conds ...clause.Expression) IEnsChannelinfosDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ensChannelinfosDo) Returning(value interface{}, columns ...string) IEnsChannelinfosDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ensChannelinfosDo) Not(conds ...gen.Condition) IEnsChannelinfosDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ensChannelinfosDo) Or(conds ...gen.Condition) IEnsChannelinfosDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ensChannelinfosDo) Select(conds ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ensChannelinfosDo) Where(conds ...gen.Condition) IEnsChannelinfosDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ensChannelinfosDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEnsChannelinfosDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e ensChannelinfosDo) Order(conds ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ensChannelinfosDo) Distinct(cols ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ensChannelinfosDo) Omit(cols ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ensChannelinfosDo) Join(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ensChannelinfosDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ensChannelinfosDo) RightJoin(table schema.Tabler, on ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ensChannelinfosDo) Group(cols ...field.Expr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ensChannelinfosDo) Having(conds ...gen.Condition) IEnsChannelinfosDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ensChannelinfosDo) Limit(limit int) IEnsChannelinfosDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ensChannelinfosDo) Offset(offset int) IEnsChannelinfosDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ensChannelinfosDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEnsChannelinfosDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ensChannelinfosDo) Unscoped() IEnsChannelinfosDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ensChannelinfosDo) Create(values ...*model.EnsChannelinfos) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ensChannelinfosDo) CreateInBatches(values []*model.EnsChannelinfos, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ensChannelinfosDo) Save(values ...*model.EnsChannelinfos) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ensChannelinfosDo) First() (*model.EnsChannelinfos, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsChannelinfos), nil
	}
}

func (e ensChannelinfosDo) Take() (*model.EnsChannelinfos, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsChannelinfos), nil
	}
}

func (e ensChannelinfosDo) Last() (*model.EnsChannelinfos, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsChannelinfos), nil
	}
}

func (e ensChannelinfosDo) Find() ([]*model.EnsChannelinfos, error) {
	result, err := e.DO.Find()
	return result.([]*model.EnsChannelinfos), err
}

func (e ensChannelinfosDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnsChannelinfos, err error) {
	buf := make([]*model.EnsChannelinfos, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ensChannelinfosDo) FindInBatches(result *[]*model.EnsChannelinfos, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ensChannelinfosDo) Attrs(attrs ...field.AssignExpr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ensChannelinfosDo) Assign(attrs ...field.AssignExpr) IEnsChannelinfosDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ensChannelinfosDo) Joins(fields ...field.RelationField) IEnsChannelinfosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ensChannelinfosDo) Preload(fields ...field.RelationField) IEnsChannelinfosDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ensChannelinfosDo) FirstOrInit() (*model.EnsChannelinfos, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsChannelinfos), nil
	}
}

func (e ensChannelinfosDo) FirstOrCreate() (*model.EnsChannelinfos, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsChannelinfos), nil
	}
}

func (e ensChannelinfosDo) FindByPage(offset int, limit int) (result []*model.EnsChannelinfos, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e ensChannelinfosDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ensChannelinfosDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ensChannelinfosDo) Delete(models ...*model.EnsChannelinfos) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ensChannelinfosDo) withDO(do gen.Dao) *ensChannelinfosDo {
	e.DO = *do.(*gen.DO)
	return e
}