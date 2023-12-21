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

func newEnsActivityDataHours(db *gorm.DB, opts ...gen.DOOption) ensActivityDataHours {
	_ensActivityDataHours := ensActivityDataHours{}

	_ensActivityDataHours.ensActivityDataHoursDo.UseDB(db, opts...)
	_ensActivityDataHours.ensActivityDataHoursDo.UseModel(&model.EnsActivityDataHours{})

	tableName := _ensActivityDataHours.ensActivityDataHoursDo.TableName()
	_ensActivityDataHours.ALL = field.NewAsterisk(tableName)
	_ensActivityDataHours.ID = field.NewInt32(tableName, "id")
	_ensActivityDataHours.Hostname = field.NewString(tableName, "hostname")
	_ensActivityDataHours.Hosttype = field.NewString(tableName, "hosttype")
	_ensActivityDataHours.Instance = field.NewString(tableName, "instance")
	_ensActivityDataHours.Namespace = field.NewString(tableName, "namespace")
	_ensActivityDataHours.Period = field.NewInt32(tableName, "period")
	_ensActivityDataHours.Sitedimension = field.NewString(tableName, "sitedimension")
	_ensActivityDataHours.Timeslot = field.NewTime(tableName, "timeslot")
	_ensActivityDataHours.Timeslotutc = field.NewTime(tableName, "timeslotutc")
	_ensActivityDataHours.Totalcount = field.NewInt32(tableName, "totalcount")
	_ensActivityDataHours.Totalduration = field.NewFloat64(tableName, "totalduration")
	_ensActivityDataHours.Totaldurationsquare = field.NewFloat64(tableName, "totaldurationsquare")
	_ensActivityDataHours.Totalqueueduration = field.NewFloat64(tableName, "totalqueueduration")
	_ensActivityDataHours.Project = field.NewString(tableName, "project")

	_ensActivityDataHours.fillFieldMap()

	return _ensActivityDataHours
}

type ensActivityDataHours struct {
	ensActivityDataHoursDo

	ALL                 field.Asterisk
	ID                  field.Int32
	Hostname            field.String
	Hosttype            field.String
	Instance            field.String
	Namespace           field.String
	Period              field.Int32
	Sitedimension       field.String
	Timeslot            field.Time
	Timeslotutc         field.Time
	Totalcount          field.Int32
	Totalduration       field.Float64
	Totaldurationsquare field.Float64
	Totalqueueduration  field.Float64
	Project             field.String // 项目编码

	fieldMap map[string]field.Expr
}

func (e ensActivityDataHours) Table(newTableName string) *ensActivityDataHours {
	e.ensActivityDataHoursDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ensActivityDataHours) As(alias string) *ensActivityDataHours {
	e.ensActivityDataHoursDo.DO = *(e.ensActivityDataHoursDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ensActivityDataHours) updateTableName(table string) *ensActivityDataHours {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Hostname = field.NewString(table, "hostname")
	e.Hosttype = field.NewString(table, "hosttype")
	e.Instance = field.NewString(table, "instance")
	e.Namespace = field.NewString(table, "namespace")
	e.Period = field.NewInt32(table, "period")
	e.Sitedimension = field.NewString(table, "sitedimension")
	e.Timeslot = field.NewTime(table, "timeslot")
	e.Timeslotutc = field.NewTime(table, "timeslotutc")
	e.Totalcount = field.NewInt32(table, "totalcount")
	e.Totalduration = field.NewFloat64(table, "totalduration")
	e.Totaldurationsquare = field.NewFloat64(table, "totaldurationsquare")
	e.Totalqueueduration = field.NewFloat64(table, "totalqueueduration")
	e.Project = field.NewString(table, "project")

	e.fillFieldMap()

	return e
}

func (e *ensActivityDataHours) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ensActivityDataHours) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 14)
	e.fieldMap["id"] = e.ID
	e.fieldMap["hostname"] = e.Hostname
	e.fieldMap["hosttype"] = e.Hosttype
	e.fieldMap["instance"] = e.Instance
	e.fieldMap["namespace"] = e.Namespace
	e.fieldMap["period"] = e.Period
	e.fieldMap["sitedimension"] = e.Sitedimension
	e.fieldMap["timeslot"] = e.Timeslot
	e.fieldMap["timeslotutc"] = e.Timeslotutc
	e.fieldMap["totalcount"] = e.Totalcount
	e.fieldMap["totalduration"] = e.Totalduration
	e.fieldMap["totaldurationsquare"] = e.Totaldurationsquare
	e.fieldMap["totalqueueduration"] = e.Totalqueueduration
	e.fieldMap["project"] = e.Project
}

func (e ensActivityDataHours) clone(db *gorm.DB) ensActivityDataHours {
	e.ensActivityDataHoursDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ensActivityDataHours) replaceDB(db *gorm.DB) ensActivityDataHours {
	e.ensActivityDataHoursDo.ReplaceDB(db)
	return e
}

type ensActivityDataHoursDo struct{ gen.DO }

type IEnsActivityDataHoursDo interface {
	gen.SubQuery
	Debug() IEnsActivityDataHoursDo
	WithContext(ctx context.Context) IEnsActivityDataHoursDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEnsActivityDataHoursDo
	WriteDB() IEnsActivityDataHoursDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEnsActivityDataHoursDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEnsActivityDataHoursDo
	Not(conds ...gen.Condition) IEnsActivityDataHoursDo
	Or(conds ...gen.Condition) IEnsActivityDataHoursDo
	Select(conds ...field.Expr) IEnsActivityDataHoursDo
	Where(conds ...gen.Condition) IEnsActivityDataHoursDo
	Order(conds ...field.Expr) IEnsActivityDataHoursDo
	Distinct(cols ...field.Expr) IEnsActivityDataHoursDo
	Omit(cols ...field.Expr) IEnsActivityDataHoursDo
	Join(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo
	Group(cols ...field.Expr) IEnsActivityDataHoursDo
	Having(conds ...gen.Condition) IEnsActivityDataHoursDo
	Limit(limit int) IEnsActivityDataHoursDo
	Offset(offset int) IEnsActivityDataHoursDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEnsActivityDataHoursDo
	Unscoped() IEnsActivityDataHoursDo
	Create(values ...*model.EnsActivityDataHours) error
	CreateInBatches(values []*model.EnsActivityDataHours, batchSize int) error
	Save(values ...*model.EnsActivityDataHours) error
	First() (*model.EnsActivityDataHours, error)
	Take() (*model.EnsActivityDataHours, error)
	Last() (*model.EnsActivityDataHours, error)
	Find() ([]*model.EnsActivityDataHours, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnsActivityDataHours, err error)
	FindInBatches(result *[]*model.EnsActivityDataHours, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EnsActivityDataHours) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEnsActivityDataHoursDo
	Assign(attrs ...field.AssignExpr) IEnsActivityDataHoursDo
	Joins(fields ...field.RelationField) IEnsActivityDataHoursDo
	Preload(fields ...field.RelationField) IEnsActivityDataHoursDo
	FirstOrInit() (*model.EnsActivityDataHours, error)
	FirstOrCreate() (*model.EnsActivityDataHours, error)
	FindByPage(offset int, limit int) (result []*model.EnsActivityDataHours, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEnsActivityDataHoursDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ensActivityDataHoursDo) Debug() IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Debug())
}

func (e ensActivityDataHoursDo) WithContext(ctx context.Context) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ensActivityDataHoursDo) ReadDB() IEnsActivityDataHoursDo {
	return e.Clauses(dbresolver.Read)
}

func (e ensActivityDataHoursDo) WriteDB() IEnsActivityDataHoursDo {
	return e.Clauses(dbresolver.Write)
}

func (e ensActivityDataHoursDo) Session(config *gorm.Session) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Session(config))
}

func (e ensActivityDataHoursDo) Clauses(conds ...clause.Expression) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ensActivityDataHoursDo) Returning(value interface{}, columns ...string) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ensActivityDataHoursDo) Not(conds ...gen.Condition) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ensActivityDataHoursDo) Or(conds ...gen.Condition) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ensActivityDataHoursDo) Select(conds ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ensActivityDataHoursDo) Where(conds ...gen.Condition) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ensActivityDataHoursDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEnsActivityDataHoursDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e ensActivityDataHoursDo) Order(conds ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ensActivityDataHoursDo) Distinct(cols ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ensActivityDataHoursDo) Omit(cols ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ensActivityDataHoursDo) Join(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ensActivityDataHoursDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ensActivityDataHoursDo) RightJoin(table schema.Tabler, on ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ensActivityDataHoursDo) Group(cols ...field.Expr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ensActivityDataHoursDo) Having(conds ...gen.Condition) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ensActivityDataHoursDo) Limit(limit int) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ensActivityDataHoursDo) Offset(offset int) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ensActivityDataHoursDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ensActivityDataHoursDo) Unscoped() IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ensActivityDataHoursDo) Create(values ...*model.EnsActivityDataHours) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ensActivityDataHoursDo) CreateInBatches(values []*model.EnsActivityDataHours, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ensActivityDataHoursDo) Save(values ...*model.EnsActivityDataHours) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ensActivityDataHoursDo) First() (*model.EnsActivityDataHours, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsActivityDataHours), nil
	}
}

func (e ensActivityDataHoursDo) Take() (*model.EnsActivityDataHours, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsActivityDataHours), nil
	}
}

func (e ensActivityDataHoursDo) Last() (*model.EnsActivityDataHours, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsActivityDataHours), nil
	}
}

func (e ensActivityDataHoursDo) Find() ([]*model.EnsActivityDataHours, error) {
	result, err := e.DO.Find()
	return result.([]*model.EnsActivityDataHours), err
}

func (e ensActivityDataHoursDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnsActivityDataHours, err error) {
	buf := make([]*model.EnsActivityDataHours, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ensActivityDataHoursDo) FindInBatches(result *[]*model.EnsActivityDataHours, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ensActivityDataHoursDo) Attrs(attrs ...field.AssignExpr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ensActivityDataHoursDo) Assign(attrs ...field.AssignExpr) IEnsActivityDataHoursDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ensActivityDataHoursDo) Joins(fields ...field.RelationField) IEnsActivityDataHoursDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ensActivityDataHoursDo) Preload(fields ...field.RelationField) IEnsActivityDataHoursDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ensActivityDataHoursDo) FirstOrInit() (*model.EnsActivityDataHours, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsActivityDataHours), nil
	}
}

func (e ensActivityDataHoursDo) FirstOrCreate() (*model.EnsActivityDataHours, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnsActivityDataHours), nil
	}
}

func (e ensActivityDataHoursDo) FindByPage(offset int, limit int) (result []*model.EnsActivityDataHours, count int64, err error) {
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

func (e ensActivityDataHoursDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ensActivityDataHoursDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ensActivityDataHoursDo) Delete(models ...*model.EnsActivityDataHours) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ensActivityDataHoursDo) withDO(do gen.Dao) *ensActivityDataHoursDo {
	e.DO = *do.(*gen.DO)
	return e
}