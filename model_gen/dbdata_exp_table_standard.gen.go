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

func newDbdataExpTableStandard(db *gorm.DB, opts ...gen.DOOption) dbdataExpTableStandard {
	_dbdataExpTableStandard := dbdataExpTableStandard{}

	_dbdataExpTableStandard.dbdataExpTableStandardDo.UseDB(db, opts...)
	_dbdataExpTableStandard.dbdataExpTableStandardDo.UseModel(&model.DbdataExpTableStandard{})

	tableName := _dbdataExpTableStandard.dbdataExpTableStandardDo.TableName()
	_dbdataExpTableStandard.ALL = field.NewAsterisk(tableName)
	_dbdataExpTableStandard.TableStandardID = field.NewString(tableName, "table_standard_id")
	_dbdataExpTableStandard.TableName_ = field.NewString(tableName, "table_name")
	_dbdataExpTableStandard.FilterWhereSql = field.NewString(tableName, "filter_where_sql")
	_dbdataExpTableStandard.ExportWay = field.NewString(tableName, "export_way")
	_dbdataExpTableStandard.ModifyTime = field.NewTime(tableName, "modify_time")
	_dbdataExpTableStandard.Comments = field.NewString(tableName, "comments")
	_dbdataExpTableStandard.IsBlacklist = field.NewString(tableName, "is_blacklist")

	_dbdataExpTableStandard.fillFieldMap()

	return _dbdataExpTableStandard
}

type dbdataExpTableStandard struct {
	dbdataExpTableStandardDo

	ALL             field.Asterisk
	TableStandardID field.String
	TableName_      field.String // 表名，包含owner，比如 EMR.CPOE_ORD
	FilterWhereSql  field.String // 数据过滤条件SQL
	ExportWay       field.String // all、数据和结构  meta仅结构，query 条件,backlist
	ModifyTime      field.Time
	Comments        field.String
	IsBlacklist     field.String

	fieldMap map[string]field.Expr
}

func (d dbdataExpTableStandard) Table(newTableName string) *dbdataExpTableStandard {
	d.dbdataExpTableStandardDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataExpTableStandard) As(alias string) *dbdataExpTableStandard {
	d.dbdataExpTableStandardDo.DO = *(d.dbdataExpTableStandardDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataExpTableStandard) updateTableName(table string) *dbdataExpTableStandard {
	d.ALL = field.NewAsterisk(table)
	d.TableStandardID = field.NewString(table, "table_standard_id")
	d.TableName_ = field.NewString(table, "table_name")
	d.FilterWhereSql = field.NewString(table, "filter_where_sql")
	d.ExportWay = field.NewString(table, "export_way")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")
	d.IsBlacklist = field.NewString(table, "is_blacklist")

	d.fillFieldMap()

	return d
}

func (d *dbdataExpTableStandard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataExpTableStandard) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 7)
	d.fieldMap["table_standard_id"] = d.TableStandardID
	d.fieldMap["table_name"] = d.TableName_
	d.fieldMap["filter_where_sql"] = d.FilterWhereSql
	d.fieldMap["export_way"] = d.ExportWay
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
	d.fieldMap["is_blacklist"] = d.IsBlacklist
}

func (d dbdataExpTableStandard) clone(db *gorm.DB) dbdataExpTableStandard {
	d.dbdataExpTableStandardDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataExpTableStandard) replaceDB(db *gorm.DB) dbdataExpTableStandard {
	d.dbdataExpTableStandardDo.ReplaceDB(db)
	return d
}

type dbdataExpTableStandardDo struct{ gen.DO }

type IDbdataExpTableStandardDo interface {
	gen.SubQuery
	Debug() IDbdataExpTableStandardDo
	WithContext(ctx context.Context) IDbdataExpTableStandardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataExpTableStandardDo
	WriteDB() IDbdataExpTableStandardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataExpTableStandardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataExpTableStandardDo
	Not(conds ...gen.Condition) IDbdataExpTableStandardDo
	Or(conds ...gen.Condition) IDbdataExpTableStandardDo
	Select(conds ...field.Expr) IDbdataExpTableStandardDo
	Where(conds ...gen.Condition) IDbdataExpTableStandardDo
	Order(conds ...field.Expr) IDbdataExpTableStandardDo
	Distinct(cols ...field.Expr) IDbdataExpTableStandardDo
	Omit(cols ...field.Expr) IDbdataExpTableStandardDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo
	Group(cols ...field.Expr) IDbdataExpTableStandardDo
	Having(conds ...gen.Condition) IDbdataExpTableStandardDo
	Limit(limit int) IDbdataExpTableStandardDo
	Offset(offset int) IDbdataExpTableStandardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExpTableStandardDo
	Unscoped() IDbdataExpTableStandardDo
	Create(values ...*model.DbdataExpTableStandard) error
	CreateInBatches(values []*model.DbdataExpTableStandard, batchSize int) error
	Save(values ...*model.DbdataExpTableStandard) error
	First() (*model.DbdataExpTableStandard, error)
	Take() (*model.DbdataExpTableStandard, error)
	Last() (*model.DbdataExpTableStandard, error)
	Find() ([]*model.DbdataExpTableStandard, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExpTableStandard, err error)
	FindInBatches(result *[]*model.DbdataExpTableStandard, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataExpTableStandard) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataExpTableStandardDo
	Assign(attrs ...field.AssignExpr) IDbdataExpTableStandardDo
	Joins(fields ...field.RelationField) IDbdataExpTableStandardDo
	Preload(fields ...field.RelationField) IDbdataExpTableStandardDo
	FirstOrInit() (*model.DbdataExpTableStandard, error)
	FirstOrCreate() (*model.DbdataExpTableStandard, error)
	FindByPage(offset int, limit int) (result []*model.DbdataExpTableStandard, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataExpTableStandardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataExpTableStandardDo) Debug() IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataExpTableStandardDo) WithContext(ctx context.Context) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataExpTableStandardDo) ReadDB() IDbdataExpTableStandardDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataExpTableStandardDo) WriteDB() IDbdataExpTableStandardDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataExpTableStandardDo) Session(config *gorm.Session) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataExpTableStandardDo) Clauses(conds ...clause.Expression) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataExpTableStandardDo) Returning(value interface{}, columns ...string) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataExpTableStandardDo) Not(conds ...gen.Condition) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataExpTableStandardDo) Or(conds ...gen.Condition) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataExpTableStandardDo) Select(conds ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataExpTableStandardDo) Where(conds ...gen.Condition) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataExpTableStandardDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataExpTableStandardDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataExpTableStandardDo) Order(conds ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataExpTableStandardDo) Distinct(cols ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataExpTableStandardDo) Omit(cols ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataExpTableStandardDo) Join(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataExpTableStandardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataExpTableStandardDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataExpTableStandardDo) Group(cols ...field.Expr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataExpTableStandardDo) Having(conds ...gen.Condition) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataExpTableStandardDo) Limit(limit int) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataExpTableStandardDo) Offset(offset int) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataExpTableStandardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataExpTableStandardDo) Unscoped() IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataExpTableStandardDo) Create(values ...*model.DbdataExpTableStandard) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataExpTableStandardDo) CreateInBatches(values []*model.DbdataExpTableStandard, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataExpTableStandardDo) Save(values ...*model.DbdataExpTableStandard) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataExpTableStandardDo) First() (*model.DbdataExpTableStandard, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExpTableStandard), nil
	}
}

func (d dbdataExpTableStandardDo) Take() (*model.DbdataExpTableStandard, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExpTableStandard), nil
	}
}

func (d dbdataExpTableStandardDo) Last() (*model.DbdataExpTableStandard, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExpTableStandard), nil
	}
}

func (d dbdataExpTableStandardDo) Find() ([]*model.DbdataExpTableStandard, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataExpTableStandard), err
}

func (d dbdataExpTableStandardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExpTableStandard, err error) {
	buf := make([]*model.DbdataExpTableStandard, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataExpTableStandardDo) FindInBatches(result *[]*model.DbdataExpTableStandard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataExpTableStandardDo) Attrs(attrs ...field.AssignExpr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataExpTableStandardDo) Assign(attrs ...field.AssignExpr) IDbdataExpTableStandardDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataExpTableStandardDo) Joins(fields ...field.RelationField) IDbdataExpTableStandardDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataExpTableStandardDo) Preload(fields ...field.RelationField) IDbdataExpTableStandardDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataExpTableStandardDo) FirstOrInit() (*model.DbdataExpTableStandard, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExpTableStandard), nil
	}
}

func (d dbdataExpTableStandardDo) FirstOrCreate() (*model.DbdataExpTableStandard, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExpTableStandard), nil
	}
}

func (d dbdataExpTableStandardDo) FindByPage(offset int, limit int) (result []*model.DbdataExpTableStandard, count int64, err error) {
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

func (d dbdataExpTableStandardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataExpTableStandardDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataExpTableStandardDo) Delete(models ...*model.DbdataExpTableStandard) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataExpTableStandardDo) withDO(do gen.Dao) *dbdataExpTableStandardDo {
	d.DO = *do.(*gen.DO)
	return d
}