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

func newDbdataExportTables(db *gorm.DB, opts ...gen.DOOption) dbdataExportTables {
	_dbdataExportTables := dbdataExportTables{}

	_dbdataExportTables.dbdataExportTablesDo.UseDB(db, opts...)
	_dbdataExportTables.dbdataExportTablesDo.UseModel(&model.DbdataExportTables{})

	tableName := _dbdataExportTables.dbdataExportTablesDo.TableName()
	_dbdataExportTables.ALL = field.NewAsterisk(tableName)
	_dbdataExportTables.DbdataExportTablesID = field.NewString(tableName, "dbdata_export_tables_id")
	_dbdataExportTables.DbdataExportID = field.NewString(tableName, "dbdata_export_id")
	_dbdataExportTables.DbUser = field.NewString(tableName, "db_user")
	_dbdataExportTables.DbTable = field.NewString(tableName, "db_table")
	_dbdataExportTables.DbQuery = field.NewString(tableName, "db_query")
	_dbdataExportTables.ExportWay = field.NewString(tableName, "export_way")

	_dbdataExportTables.fillFieldMap()

	return _dbdataExportTables
}

type dbdataExportTables struct {
	dbdataExportTablesDo

	ALL                  field.Asterisk
	DbdataExportTablesID field.String
	DbdataExportID       field.String
	DbUser               field.String
	DbTable              field.String
	DbQuery              field.String
	ExportWay            field.String // 导出方式

	fieldMap map[string]field.Expr
}

func (d dbdataExportTables) Table(newTableName string) *dbdataExportTables {
	d.dbdataExportTablesDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataExportTables) As(alias string) *dbdataExportTables {
	d.dbdataExportTablesDo.DO = *(d.dbdataExportTablesDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataExportTables) updateTableName(table string) *dbdataExportTables {
	d.ALL = field.NewAsterisk(table)
	d.DbdataExportTablesID = field.NewString(table, "dbdata_export_tables_id")
	d.DbdataExportID = field.NewString(table, "dbdata_export_id")
	d.DbUser = field.NewString(table, "db_user")
	d.DbTable = field.NewString(table, "db_table")
	d.DbQuery = field.NewString(table, "db_query")
	d.ExportWay = field.NewString(table, "export_way")

	d.fillFieldMap()

	return d
}

func (d *dbdataExportTables) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataExportTables) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["dbdata_export_tables_id"] = d.DbdataExportTablesID
	d.fieldMap["dbdata_export_id"] = d.DbdataExportID
	d.fieldMap["db_user"] = d.DbUser
	d.fieldMap["db_table"] = d.DbTable
	d.fieldMap["db_query"] = d.DbQuery
	d.fieldMap["export_way"] = d.ExportWay
}

func (d dbdataExportTables) clone(db *gorm.DB) dbdataExportTables {
	d.dbdataExportTablesDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataExportTables) replaceDB(db *gorm.DB) dbdataExportTables {
	d.dbdataExportTablesDo.ReplaceDB(db)
	return d
}

type dbdataExportTablesDo struct{ gen.DO }

type IDbdataExportTablesDo interface {
	gen.SubQuery
	Debug() IDbdataExportTablesDo
	WithContext(ctx context.Context) IDbdataExportTablesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataExportTablesDo
	WriteDB() IDbdataExportTablesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataExportTablesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataExportTablesDo
	Not(conds ...gen.Condition) IDbdataExportTablesDo
	Or(conds ...gen.Condition) IDbdataExportTablesDo
	Select(conds ...field.Expr) IDbdataExportTablesDo
	Where(conds ...gen.Condition) IDbdataExportTablesDo
	Order(conds ...field.Expr) IDbdataExportTablesDo
	Distinct(cols ...field.Expr) IDbdataExportTablesDo
	Omit(cols ...field.Expr) IDbdataExportTablesDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo
	Group(cols ...field.Expr) IDbdataExportTablesDo
	Having(conds ...gen.Condition) IDbdataExportTablesDo
	Limit(limit int) IDbdataExportTablesDo
	Offset(offset int) IDbdataExportTablesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportTablesDo
	Unscoped() IDbdataExportTablesDo
	Create(values ...*model.DbdataExportTables) error
	CreateInBatches(values []*model.DbdataExportTables, batchSize int) error
	Save(values ...*model.DbdataExportTables) error
	First() (*model.DbdataExportTables, error)
	Take() (*model.DbdataExportTables, error)
	Last() (*model.DbdataExportTables, error)
	Find() ([]*model.DbdataExportTables, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExportTables, err error)
	FindInBatches(result *[]*model.DbdataExportTables, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataExportTables) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataExportTablesDo
	Assign(attrs ...field.AssignExpr) IDbdataExportTablesDo
	Joins(fields ...field.RelationField) IDbdataExportTablesDo
	Preload(fields ...field.RelationField) IDbdataExportTablesDo
	FirstOrInit() (*model.DbdataExportTables, error)
	FirstOrCreate() (*model.DbdataExportTables, error)
	FindByPage(offset int, limit int) (result []*model.DbdataExportTables, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataExportTablesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataExportTablesDo) Debug() IDbdataExportTablesDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataExportTablesDo) WithContext(ctx context.Context) IDbdataExportTablesDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataExportTablesDo) ReadDB() IDbdataExportTablesDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataExportTablesDo) WriteDB() IDbdataExportTablesDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataExportTablesDo) Session(config *gorm.Session) IDbdataExportTablesDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataExportTablesDo) Clauses(conds ...clause.Expression) IDbdataExportTablesDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataExportTablesDo) Returning(value interface{}, columns ...string) IDbdataExportTablesDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataExportTablesDo) Not(conds ...gen.Condition) IDbdataExportTablesDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataExportTablesDo) Or(conds ...gen.Condition) IDbdataExportTablesDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataExportTablesDo) Select(conds ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataExportTablesDo) Where(conds ...gen.Condition) IDbdataExportTablesDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataExportTablesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataExportTablesDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataExportTablesDo) Order(conds ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataExportTablesDo) Distinct(cols ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataExportTablesDo) Omit(cols ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataExportTablesDo) Join(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataExportTablesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataExportTablesDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataExportTablesDo) Group(cols ...field.Expr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataExportTablesDo) Having(conds ...gen.Condition) IDbdataExportTablesDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataExportTablesDo) Limit(limit int) IDbdataExportTablesDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataExportTablesDo) Offset(offset int) IDbdataExportTablesDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataExportTablesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportTablesDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataExportTablesDo) Unscoped() IDbdataExportTablesDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataExportTablesDo) Create(values ...*model.DbdataExportTables) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataExportTablesDo) CreateInBatches(values []*model.DbdataExportTables, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataExportTablesDo) Save(values ...*model.DbdataExportTables) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataExportTablesDo) First() (*model.DbdataExportTables, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportTables), nil
	}
}

func (d dbdataExportTablesDo) Take() (*model.DbdataExportTables, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportTables), nil
	}
}

func (d dbdataExportTablesDo) Last() (*model.DbdataExportTables, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportTables), nil
	}
}

func (d dbdataExportTablesDo) Find() ([]*model.DbdataExportTables, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataExportTables), err
}

func (d dbdataExportTablesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExportTables, err error) {
	buf := make([]*model.DbdataExportTables, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataExportTablesDo) FindInBatches(result *[]*model.DbdataExportTables, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataExportTablesDo) Attrs(attrs ...field.AssignExpr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataExportTablesDo) Assign(attrs ...field.AssignExpr) IDbdataExportTablesDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataExportTablesDo) Joins(fields ...field.RelationField) IDbdataExportTablesDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataExportTablesDo) Preload(fields ...field.RelationField) IDbdataExportTablesDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataExportTablesDo) FirstOrInit() (*model.DbdataExportTables, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportTables), nil
	}
}

func (d dbdataExportTablesDo) FirstOrCreate() (*model.DbdataExportTables, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportTables), nil
	}
}

func (d dbdataExportTablesDo) FindByPage(offset int, limit int) (result []*model.DbdataExportTables, count int64, err error) {
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

func (d dbdataExportTablesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataExportTablesDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataExportTablesDo) Delete(models ...*model.DbdataExportTables) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataExportTablesDo) withDO(do gen.Dao) *dbdataExportTablesDo {
	d.DO = *do.(*gen.DO)
	return d
}
