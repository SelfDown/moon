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

func newDbscriptReqProject(db *gorm.DB, opts ...gen.DOOption) dbscriptReqProject {
	_dbscriptReqProject := dbscriptReqProject{}

	_dbscriptReqProject.dbscriptReqProjectDo.UseDB(db, opts...)
	_dbscriptReqProject.dbscriptReqProjectDo.UseModel(&model.DbscriptReqProject{})

	tableName := _dbscriptReqProject.dbscriptReqProjectDo.TableName()
	_dbscriptReqProject.ALL = field.NewAsterisk(tableName)
	_dbscriptReqProject.DbscriptHosID = field.NewString(tableName, "dbscript_hos_id")
	_dbscriptReqProject.DbscriptEventID = field.NewString(tableName, "dbscript_event_id")
	_dbscriptReqProject.SysProjectID = field.NewString(tableName, "sys_project_id")
	_dbscriptReqProject.CreateTime = field.NewTime(tableName, "create_time")
	_dbscriptReqProject.ModifyTime = field.NewTime(tableName, "modify_time")
	_dbscriptReqProject.Comments = field.NewString(tableName, "comments")

	_dbscriptReqProject.fillFieldMap()

	return _dbscriptReqProject
}

type dbscriptReqProject struct {
	dbscriptReqProjectDo

	ALL             field.Asterisk
	DbscriptHosID   field.String
	DbscriptEventID field.String
	SysProjectID    field.String // 所属项目
	CreateTime      field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time   // 记录修改时间（数据库自动写入）
	Comments        field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (d dbscriptReqProject) Table(newTableName string) *dbscriptReqProject {
	d.dbscriptReqProjectDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbscriptReqProject) As(alias string) *dbscriptReqProject {
	d.dbscriptReqProjectDo.DO = *(d.dbscriptReqProjectDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbscriptReqProject) updateTableName(table string) *dbscriptReqProject {
	d.ALL = field.NewAsterisk(table)
	d.DbscriptHosID = field.NewString(table, "dbscript_hos_id")
	d.DbscriptEventID = field.NewString(table, "dbscript_event_id")
	d.SysProjectID = field.NewString(table, "sys_project_id")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")

	d.fillFieldMap()

	return d
}

func (d *dbscriptReqProject) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbscriptReqProject) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["dbscript_hos_id"] = d.DbscriptHosID
	d.fieldMap["dbscript_event_id"] = d.DbscriptEventID
	d.fieldMap["sys_project_id"] = d.SysProjectID
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
}

func (d dbscriptReqProject) clone(db *gorm.DB) dbscriptReqProject {
	d.dbscriptReqProjectDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbscriptReqProject) replaceDB(db *gorm.DB) dbscriptReqProject {
	d.dbscriptReqProjectDo.ReplaceDB(db)
	return d
}

type dbscriptReqProjectDo struct{ gen.DO }

type IDbscriptReqProjectDo interface {
	gen.SubQuery
	Debug() IDbscriptReqProjectDo
	WithContext(ctx context.Context) IDbscriptReqProjectDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbscriptReqProjectDo
	WriteDB() IDbscriptReqProjectDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbscriptReqProjectDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbscriptReqProjectDo
	Not(conds ...gen.Condition) IDbscriptReqProjectDo
	Or(conds ...gen.Condition) IDbscriptReqProjectDo
	Select(conds ...field.Expr) IDbscriptReqProjectDo
	Where(conds ...gen.Condition) IDbscriptReqProjectDo
	Order(conds ...field.Expr) IDbscriptReqProjectDo
	Distinct(cols ...field.Expr) IDbscriptReqProjectDo
	Omit(cols ...field.Expr) IDbscriptReqProjectDo
	Join(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo
	Group(cols ...field.Expr) IDbscriptReqProjectDo
	Having(conds ...gen.Condition) IDbscriptReqProjectDo
	Limit(limit int) IDbscriptReqProjectDo
	Offset(offset int) IDbscriptReqProjectDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbscriptReqProjectDo
	Unscoped() IDbscriptReqProjectDo
	Create(values ...*model.DbscriptReqProject) error
	CreateInBatches(values []*model.DbscriptReqProject, batchSize int) error
	Save(values ...*model.DbscriptReqProject) error
	First() (*model.DbscriptReqProject, error)
	Take() (*model.DbscriptReqProject, error)
	Last() (*model.DbscriptReqProject, error)
	Find() ([]*model.DbscriptReqProject, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbscriptReqProject, err error)
	FindInBatches(result *[]*model.DbscriptReqProject, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbscriptReqProject) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbscriptReqProjectDo
	Assign(attrs ...field.AssignExpr) IDbscriptReqProjectDo
	Joins(fields ...field.RelationField) IDbscriptReqProjectDo
	Preload(fields ...field.RelationField) IDbscriptReqProjectDo
	FirstOrInit() (*model.DbscriptReqProject, error)
	FirstOrCreate() (*model.DbscriptReqProject, error)
	FindByPage(offset int, limit int) (result []*model.DbscriptReqProject, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbscriptReqProjectDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbscriptReqProjectDo) Debug() IDbscriptReqProjectDo {
	return d.withDO(d.DO.Debug())
}

func (d dbscriptReqProjectDo) WithContext(ctx context.Context) IDbscriptReqProjectDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbscriptReqProjectDo) ReadDB() IDbscriptReqProjectDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbscriptReqProjectDo) WriteDB() IDbscriptReqProjectDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbscriptReqProjectDo) Session(config *gorm.Session) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbscriptReqProjectDo) Clauses(conds ...clause.Expression) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbscriptReqProjectDo) Returning(value interface{}, columns ...string) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbscriptReqProjectDo) Not(conds ...gen.Condition) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbscriptReqProjectDo) Or(conds ...gen.Condition) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbscriptReqProjectDo) Select(conds ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbscriptReqProjectDo) Where(conds ...gen.Condition) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbscriptReqProjectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbscriptReqProjectDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbscriptReqProjectDo) Order(conds ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbscriptReqProjectDo) Distinct(cols ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbscriptReqProjectDo) Omit(cols ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbscriptReqProjectDo) Join(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbscriptReqProjectDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbscriptReqProjectDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbscriptReqProjectDo) Group(cols ...field.Expr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbscriptReqProjectDo) Having(conds ...gen.Condition) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbscriptReqProjectDo) Limit(limit int) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbscriptReqProjectDo) Offset(offset int) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbscriptReqProjectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbscriptReqProjectDo) Unscoped() IDbscriptReqProjectDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbscriptReqProjectDo) Create(values ...*model.DbscriptReqProject) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbscriptReqProjectDo) CreateInBatches(values []*model.DbscriptReqProject, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbscriptReqProjectDo) Save(values ...*model.DbscriptReqProject) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbscriptReqProjectDo) First() (*model.DbscriptReqProject, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbscriptReqProject), nil
	}
}

func (d dbscriptReqProjectDo) Take() (*model.DbscriptReqProject, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbscriptReqProject), nil
	}
}

func (d dbscriptReqProjectDo) Last() (*model.DbscriptReqProject, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbscriptReqProject), nil
	}
}

func (d dbscriptReqProjectDo) Find() ([]*model.DbscriptReqProject, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbscriptReqProject), err
}

func (d dbscriptReqProjectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbscriptReqProject, err error) {
	buf := make([]*model.DbscriptReqProject, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbscriptReqProjectDo) FindInBatches(result *[]*model.DbscriptReqProject, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbscriptReqProjectDo) Attrs(attrs ...field.AssignExpr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbscriptReqProjectDo) Assign(attrs ...field.AssignExpr) IDbscriptReqProjectDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbscriptReqProjectDo) Joins(fields ...field.RelationField) IDbscriptReqProjectDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbscriptReqProjectDo) Preload(fields ...field.RelationField) IDbscriptReqProjectDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbscriptReqProjectDo) FirstOrInit() (*model.DbscriptReqProject, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbscriptReqProject), nil
	}
}

func (d dbscriptReqProjectDo) FirstOrCreate() (*model.DbscriptReqProject, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbscriptReqProject), nil
	}
}

func (d dbscriptReqProjectDo) FindByPage(offset int, limit int) (result []*model.DbscriptReqProject, count int64, err error) {
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

func (d dbscriptReqProjectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbscriptReqProjectDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbscriptReqProjectDo) Delete(models ...*model.DbscriptReqProject) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbscriptReqProjectDo) withDO(do gen.Dao) *dbscriptReqProjectDo {
	d.DO = *do.(*gen.DO)
	return d
}
