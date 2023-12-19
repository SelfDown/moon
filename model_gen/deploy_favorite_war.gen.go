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

func newDeployFavoriteWar(db *gorm.DB, opts ...gen.DOOption) deployFavoriteWar {
	_deployFavoriteWar := deployFavoriteWar{}

	_deployFavoriteWar.deployFavoriteWarDo.UseDB(db, opts...)
	_deployFavoriteWar.deployFavoriteWarDo.UseModel(&model.DeployFavoriteWar{})

	tableName := _deployFavoriteWar.deployFavoriteWarDo.TableName()
	_deployFavoriteWar.ALL = field.NewAsterisk(tableName)
	_deployFavoriteWar.FavoID = field.NewString(tableName, "favo_id")
	_deployFavoriteWar.Userid = field.NewString(tableName, "userid")
	_deployFavoriteWar.WarID = field.NewString(tableName, "war_id")
	_deployFavoriteWar.CreateTime = field.NewTime(tableName, "create_time")
	_deployFavoriteWar.ModifyTime = field.NewTime(tableName, "modify_time")
	_deployFavoriteWar.Comments = field.NewString(tableName, "comments")

	_deployFavoriteWar.fillFieldMap()

	return _deployFavoriteWar
}

type deployFavoriteWar struct {
	deployFavoriteWarDo

	ALL        field.Asterisk
	FavoID     field.String
	Userid     field.String
	WarID      field.String
	CreateTime field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime field.Time   // 记录修改时间（数据库自动写入）
	Comments   field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (d deployFavoriteWar) Table(newTableName string) *deployFavoriteWar {
	d.deployFavoriteWarDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deployFavoriteWar) As(alias string) *deployFavoriteWar {
	d.deployFavoriteWarDo.DO = *(d.deployFavoriteWarDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deployFavoriteWar) updateTableName(table string) *deployFavoriteWar {
	d.ALL = field.NewAsterisk(table)
	d.FavoID = field.NewString(table, "favo_id")
	d.Userid = field.NewString(table, "userid")
	d.WarID = field.NewString(table, "war_id")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")

	d.fillFieldMap()

	return d
}

func (d *deployFavoriteWar) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deployFavoriteWar) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["favo_id"] = d.FavoID
	d.fieldMap["userid"] = d.Userid
	d.fieldMap["war_id"] = d.WarID
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
}

func (d deployFavoriteWar) clone(db *gorm.DB) deployFavoriteWar {
	d.deployFavoriteWarDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deployFavoriteWar) replaceDB(db *gorm.DB) deployFavoriteWar {
	d.deployFavoriteWarDo.ReplaceDB(db)
	return d
}

type deployFavoriteWarDo struct{ gen.DO }

type IDeployFavoriteWarDo interface {
	gen.SubQuery
	Debug() IDeployFavoriteWarDo
	WithContext(ctx context.Context) IDeployFavoriteWarDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeployFavoriteWarDo
	WriteDB() IDeployFavoriteWarDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeployFavoriteWarDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeployFavoriteWarDo
	Not(conds ...gen.Condition) IDeployFavoriteWarDo
	Or(conds ...gen.Condition) IDeployFavoriteWarDo
	Select(conds ...field.Expr) IDeployFavoriteWarDo
	Where(conds ...gen.Condition) IDeployFavoriteWarDo
	Order(conds ...field.Expr) IDeployFavoriteWarDo
	Distinct(cols ...field.Expr) IDeployFavoriteWarDo
	Omit(cols ...field.Expr) IDeployFavoriteWarDo
	Join(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo
	Group(cols ...field.Expr) IDeployFavoriteWarDo
	Having(conds ...gen.Condition) IDeployFavoriteWarDo
	Limit(limit int) IDeployFavoriteWarDo
	Offset(offset int) IDeployFavoriteWarDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployFavoriteWarDo
	Unscoped() IDeployFavoriteWarDo
	Create(values ...*model.DeployFavoriteWar) error
	CreateInBatches(values []*model.DeployFavoriteWar, batchSize int) error
	Save(values ...*model.DeployFavoriteWar) error
	First() (*model.DeployFavoriteWar, error)
	Take() (*model.DeployFavoriteWar, error)
	Last() (*model.DeployFavoriteWar, error)
	Find() ([]*model.DeployFavoriteWar, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployFavoriteWar, err error)
	FindInBatches(result *[]*model.DeployFavoriteWar, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeployFavoriteWar) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeployFavoriteWarDo
	Assign(attrs ...field.AssignExpr) IDeployFavoriteWarDo
	Joins(fields ...field.RelationField) IDeployFavoriteWarDo
	Preload(fields ...field.RelationField) IDeployFavoriteWarDo
	FirstOrInit() (*model.DeployFavoriteWar, error)
	FirstOrCreate() (*model.DeployFavoriteWar, error)
	FindByPage(offset int, limit int) (result []*model.DeployFavoriteWar, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeployFavoriteWarDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deployFavoriteWarDo) Debug() IDeployFavoriteWarDo {
	return d.withDO(d.DO.Debug())
}

func (d deployFavoriteWarDo) WithContext(ctx context.Context) IDeployFavoriteWarDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deployFavoriteWarDo) ReadDB() IDeployFavoriteWarDo {
	return d.Clauses(dbresolver.Read)
}

func (d deployFavoriteWarDo) WriteDB() IDeployFavoriteWarDo {
	return d.Clauses(dbresolver.Write)
}

func (d deployFavoriteWarDo) Session(config *gorm.Session) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Session(config))
}

func (d deployFavoriteWarDo) Clauses(conds ...clause.Expression) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deployFavoriteWarDo) Returning(value interface{}, columns ...string) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deployFavoriteWarDo) Not(conds ...gen.Condition) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deployFavoriteWarDo) Or(conds ...gen.Condition) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deployFavoriteWarDo) Select(conds ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deployFavoriteWarDo) Where(conds ...gen.Condition) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deployFavoriteWarDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeployFavoriteWarDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deployFavoriteWarDo) Order(conds ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deployFavoriteWarDo) Distinct(cols ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deployFavoriteWarDo) Omit(cols ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deployFavoriteWarDo) Join(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deployFavoriteWarDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deployFavoriteWarDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deployFavoriteWarDo) Group(cols ...field.Expr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deployFavoriteWarDo) Having(conds ...gen.Condition) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deployFavoriteWarDo) Limit(limit int) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deployFavoriteWarDo) Offset(offset int) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deployFavoriteWarDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deployFavoriteWarDo) Unscoped() IDeployFavoriteWarDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deployFavoriteWarDo) Create(values ...*model.DeployFavoriteWar) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deployFavoriteWarDo) CreateInBatches(values []*model.DeployFavoriteWar, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deployFavoriteWarDo) Save(values ...*model.DeployFavoriteWar) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deployFavoriteWarDo) First() (*model.DeployFavoriteWar, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFavoriteWar), nil
	}
}

func (d deployFavoriteWarDo) Take() (*model.DeployFavoriteWar, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFavoriteWar), nil
	}
}

func (d deployFavoriteWarDo) Last() (*model.DeployFavoriteWar, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFavoriteWar), nil
	}
}

func (d deployFavoriteWarDo) Find() ([]*model.DeployFavoriteWar, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeployFavoriteWar), err
}

func (d deployFavoriteWarDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployFavoriteWar, err error) {
	buf := make([]*model.DeployFavoriteWar, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deployFavoriteWarDo) FindInBatches(result *[]*model.DeployFavoriteWar, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deployFavoriteWarDo) Attrs(attrs ...field.AssignExpr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deployFavoriteWarDo) Assign(attrs ...field.AssignExpr) IDeployFavoriteWarDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deployFavoriteWarDo) Joins(fields ...field.RelationField) IDeployFavoriteWarDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deployFavoriteWarDo) Preload(fields ...field.RelationField) IDeployFavoriteWarDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deployFavoriteWarDo) FirstOrInit() (*model.DeployFavoriteWar, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFavoriteWar), nil
	}
}

func (d deployFavoriteWarDo) FirstOrCreate() (*model.DeployFavoriteWar, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFavoriteWar), nil
	}
}

func (d deployFavoriteWarDo) FindByPage(offset int, limit int) (result []*model.DeployFavoriteWar, count int64, err error) {
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

func (d deployFavoriteWarDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deployFavoriteWarDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deployFavoriteWarDo) Delete(models ...*model.DeployFavoriteWar) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deployFavoriteWarDo) withDO(do gen.Dao) *deployFavoriteWarDo {
	d.DO = *do.(*gen.DO)
	return d
}
