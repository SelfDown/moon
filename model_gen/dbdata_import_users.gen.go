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

func newDbdataImportUsers(db *gorm.DB, opts ...gen.DOOption) dbdataImportUsers {
	_dbdataImportUsers := dbdataImportUsers{}

	_dbdataImportUsers.dbdataImportUsersDo.UseDB(db, opts...)
	_dbdataImportUsers.dbdataImportUsersDo.UseModel(&model.DbdataImportUsers{})

	tableName := _dbdataImportUsers.dbdataImportUsersDo.TableName()
	_dbdataImportUsers.ALL = field.NewAsterisk(tableName)
	_dbdataImportUsers.DbdataImportUsersID = field.NewString(tableName, "dbdata_import_users_id")
	_dbdataImportUsers.DbdataImportID = field.NewString(tableName, "dbdata_import_id")
	_dbdataImportUsers.DbUser = field.NewString(tableName, "db_user")

	_dbdataImportUsers.fillFieldMap()

	return _dbdataImportUsers
}

type dbdataImportUsers struct {
	dbdataImportUsersDo

	ALL                 field.Asterisk
	DbdataImportUsersID field.String
	DbdataImportID      field.String
	DbUser              field.String

	fieldMap map[string]field.Expr
}

func (d dbdataImportUsers) Table(newTableName string) *dbdataImportUsers {
	d.dbdataImportUsersDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataImportUsers) As(alias string) *dbdataImportUsers {
	d.dbdataImportUsersDo.DO = *(d.dbdataImportUsersDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataImportUsers) updateTableName(table string) *dbdataImportUsers {
	d.ALL = field.NewAsterisk(table)
	d.DbdataImportUsersID = field.NewString(table, "dbdata_import_users_id")
	d.DbdataImportID = field.NewString(table, "dbdata_import_id")
	d.DbUser = field.NewString(table, "db_user")

	d.fillFieldMap()

	return d
}

func (d *dbdataImportUsers) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataImportUsers) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 3)
	d.fieldMap["dbdata_import_users_id"] = d.DbdataImportUsersID
	d.fieldMap["dbdata_import_id"] = d.DbdataImportID
	d.fieldMap["db_user"] = d.DbUser
}

func (d dbdataImportUsers) clone(db *gorm.DB) dbdataImportUsers {
	d.dbdataImportUsersDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataImportUsers) replaceDB(db *gorm.DB) dbdataImportUsers {
	d.dbdataImportUsersDo.ReplaceDB(db)
	return d
}

type dbdataImportUsersDo struct{ gen.DO }

type IDbdataImportUsersDo interface {
	gen.SubQuery
	Debug() IDbdataImportUsersDo
	WithContext(ctx context.Context) IDbdataImportUsersDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataImportUsersDo
	WriteDB() IDbdataImportUsersDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataImportUsersDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataImportUsersDo
	Not(conds ...gen.Condition) IDbdataImportUsersDo
	Or(conds ...gen.Condition) IDbdataImportUsersDo
	Select(conds ...field.Expr) IDbdataImportUsersDo
	Where(conds ...gen.Condition) IDbdataImportUsersDo
	Order(conds ...field.Expr) IDbdataImportUsersDo
	Distinct(cols ...field.Expr) IDbdataImportUsersDo
	Omit(cols ...field.Expr) IDbdataImportUsersDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo
	Group(cols ...field.Expr) IDbdataImportUsersDo
	Having(conds ...gen.Condition) IDbdataImportUsersDo
	Limit(limit int) IDbdataImportUsersDo
	Offset(offset int) IDbdataImportUsersDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataImportUsersDo
	Unscoped() IDbdataImportUsersDo
	Create(values ...*model.DbdataImportUsers) error
	CreateInBatches(values []*model.DbdataImportUsers, batchSize int) error
	Save(values ...*model.DbdataImportUsers) error
	First() (*model.DbdataImportUsers, error)
	Take() (*model.DbdataImportUsers, error)
	Last() (*model.DbdataImportUsers, error)
	Find() ([]*model.DbdataImportUsers, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataImportUsers, err error)
	FindInBatches(result *[]*model.DbdataImportUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataImportUsers) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataImportUsersDo
	Assign(attrs ...field.AssignExpr) IDbdataImportUsersDo
	Joins(fields ...field.RelationField) IDbdataImportUsersDo
	Preload(fields ...field.RelationField) IDbdataImportUsersDo
	FirstOrInit() (*model.DbdataImportUsers, error)
	FirstOrCreate() (*model.DbdataImportUsers, error)
	FindByPage(offset int, limit int) (result []*model.DbdataImportUsers, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataImportUsersDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataImportUsersDo) Debug() IDbdataImportUsersDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataImportUsersDo) WithContext(ctx context.Context) IDbdataImportUsersDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataImportUsersDo) ReadDB() IDbdataImportUsersDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataImportUsersDo) WriteDB() IDbdataImportUsersDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataImportUsersDo) Session(config *gorm.Session) IDbdataImportUsersDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataImportUsersDo) Clauses(conds ...clause.Expression) IDbdataImportUsersDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataImportUsersDo) Returning(value interface{}, columns ...string) IDbdataImportUsersDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataImportUsersDo) Not(conds ...gen.Condition) IDbdataImportUsersDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataImportUsersDo) Or(conds ...gen.Condition) IDbdataImportUsersDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataImportUsersDo) Select(conds ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataImportUsersDo) Where(conds ...gen.Condition) IDbdataImportUsersDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataImportUsersDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataImportUsersDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataImportUsersDo) Order(conds ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataImportUsersDo) Distinct(cols ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataImportUsersDo) Omit(cols ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataImportUsersDo) Join(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataImportUsersDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataImportUsersDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataImportUsersDo) Group(cols ...field.Expr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataImportUsersDo) Having(conds ...gen.Condition) IDbdataImportUsersDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataImportUsersDo) Limit(limit int) IDbdataImportUsersDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataImportUsersDo) Offset(offset int) IDbdataImportUsersDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataImportUsersDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataImportUsersDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataImportUsersDo) Unscoped() IDbdataImportUsersDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataImportUsersDo) Create(values ...*model.DbdataImportUsers) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataImportUsersDo) CreateInBatches(values []*model.DbdataImportUsers, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataImportUsersDo) Save(values ...*model.DbdataImportUsers) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataImportUsersDo) First() (*model.DbdataImportUsers, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportUsers), nil
	}
}

func (d dbdataImportUsersDo) Take() (*model.DbdataImportUsers, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportUsers), nil
	}
}

func (d dbdataImportUsersDo) Last() (*model.DbdataImportUsers, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportUsers), nil
	}
}

func (d dbdataImportUsersDo) Find() ([]*model.DbdataImportUsers, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataImportUsers), err
}

func (d dbdataImportUsersDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataImportUsers, err error) {
	buf := make([]*model.DbdataImportUsers, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataImportUsersDo) FindInBatches(result *[]*model.DbdataImportUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataImportUsersDo) Attrs(attrs ...field.AssignExpr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataImportUsersDo) Assign(attrs ...field.AssignExpr) IDbdataImportUsersDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataImportUsersDo) Joins(fields ...field.RelationField) IDbdataImportUsersDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataImportUsersDo) Preload(fields ...field.RelationField) IDbdataImportUsersDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataImportUsersDo) FirstOrInit() (*model.DbdataImportUsers, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportUsers), nil
	}
}

func (d dbdataImportUsersDo) FirstOrCreate() (*model.DbdataImportUsers, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportUsers), nil
	}
}

func (d dbdataImportUsersDo) FindByPage(offset int, limit int) (result []*model.DbdataImportUsers, count int64, err error) {
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

func (d dbdataImportUsersDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataImportUsersDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataImportUsersDo) Delete(models ...*model.DbdataImportUsers) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataImportUsersDo) withDO(do gen.Dao) *dbdataImportUsersDo {
	d.DO = *do.(*gen.DO)
	return d
}
