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

func newDeployPrepareEvent(db *gorm.DB, opts ...gen.DOOption) deployPrepareEvent {
	_deployPrepareEvent := deployPrepareEvent{}

	_deployPrepareEvent.deployPrepareEventDo.UseDB(db, opts...)
	_deployPrepareEvent.deployPrepareEventDo.UseModel(&model.DeployPrepareEvent{})

	tableName := _deployPrepareEvent.deployPrepareEventDo.TableName()
	_deployPrepareEvent.ALL = field.NewAsterisk(tableName)
	_deployPrepareEvent.DepPreID = field.NewString(tableName, "dep_pre_id")
	_deployPrepareEvent.OpUser = field.NewString(tableName, "op_user")
	_deployPrepareEvent.OpTime = field.NewTime(tableName, "op_time")
	_deployPrepareEvent.HospitalArea = field.NewString(tableName, "hospital_area")
	_deployPrepareEvent.DepRepID = field.NewString(tableName, "dep_rep_id")
	_deployPrepareEvent.CreateTime = field.NewTime(tableName, "create_time")
	_deployPrepareEvent.ModifyTime = field.NewTime(tableName, "modify_time")
	_deployPrepareEvent.Comments = field.NewString(tableName, "comments")

	_deployPrepareEvent.fillFieldMap()

	return _deployPrepareEvent
}

type deployPrepareEvent struct {
	deployPrepareEventDo

	ALL          field.Asterisk
	DepPreID     field.String
	OpUser       field.String // 操作员
	OpTime       field.Time   // 操作时间
	HospitalArea field.String // 院区
	DepRepID     field.String
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (d deployPrepareEvent) Table(newTableName string) *deployPrepareEvent {
	d.deployPrepareEventDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deployPrepareEvent) As(alias string) *deployPrepareEvent {
	d.deployPrepareEventDo.DO = *(d.deployPrepareEventDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deployPrepareEvent) updateTableName(table string) *deployPrepareEvent {
	d.ALL = field.NewAsterisk(table)
	d.DepPreID = field.NewString(table, "dep_pre_id")
	d.OpUser = field.NewString(table, "op_user")
	d.OpTime = field.NewTime(table, "op_time")
	d.HospitalArea = field.NewString(table, "hospital_area")
	d.DepRepID = field.NewString(table, "dep_rep_id")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")

	d.fillFieldMap()

	return d
}

func (d *deployPrepareEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deployPrepareEvent) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["dep_pre_id"] = d.DepPreID
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["op_time"] = d.OpTime
	d.fieldMap["hospital_area"] = d.HospitalArea
	d.fieldMap["dep_rep_id"] = d.DepRepID
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
}

func (d deployPrepareEvent) clone(db *gorm.DB) deployPrepareEvent {
	d.deployPrepareEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deployPrepareEvent) replaceDB(db *gorm.DB) deployPrepareEvent {
	d.deployPrepareEventDo.ReplaceDB(db)
	return d
}

type deployPrepareEventDo struct{ gen.DO }

type IDeployPrepareEventDo interface {
	gen.SubQuery
	Debug() IDeployPrepareEventDo
	WithContext(ctx context.Context) IDeployPrepareEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeployPrepareEventDo
	WriteDB() IDeployPrepareEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeployPrepareEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeployPrepareEventDo
	Not(conds ...gen.Condition) IDeployPrepareEventDo
	Or(conds ...gen.Condition) IDeployPrepareEventDo
	Select(conds ...field.Expr) IDeployPrepareEventDo
	Where(conds ...gen.Condition) IDeployPrepareEventDo
	Order(conds ...field.Expr) IDeployPrepareEventDo
	Distinct(cols ...field.Expr) IDeployPrepareEventDo
	Omit(cols ...field.Expr) IDeployPrepareEventDo
	Join(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo
	Group(cols ...field.Expr) IDeployPrepareEventDo
	Having(conds ...gen.Condition) IDeployPrepareEventDo
	Limit(limit int) IDeployPrepareEventDo
	Offset(offset int) IDeployPrepareEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployPrepareEventDo
	Unscoped() IDeployPrepareEventDo
	Create(values ...*model.DeployPrepareEvent) error
	CreateInBatches(values []*model.DeployPrepareEvent, batchSize int) error
	Save(values ...*model.DeployPrepareEvent) error
	First() (*model.DeployPrepareEvent, error)
	Take() (*model.DeployPrepareEvent, error)
	Last() (*model.DeployPrepareEvent, error)
	Find() ([]*model.DeployPrepareEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployPrepareEvent, err error)
	FindInBatches(result *[]*model.DeployPrepareEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeployPrepareEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeployPrepareEventDo
	Assign(attrs ...field.AssignExpr) IDeployPrepareEventDo
	Joins(fields ...field.RelationField) IDeployPrepareEventDo
	Preload(fields ...field.RelationField) IDeployPrepareEventDo
	FirstOrInit() (*model.DeployPrepareEvent, error)
	FirstOrCreate() (*model.DeployPrepareEvent, error)
	FindByPage(offset int, limit int) (result []*model.DeployPrepareEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeployPrepareEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deployPrepareEventDo) Debug() IDeployPrepareEventDo {
	return d.withDO(d.DO.Debug())
}

func (d deployPrepareEventDo) WithContext(ctx context.Context) IDeployPrepareEventDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deployPrepareEventDo) ReadDB() IDeployPrepareEventDo {
	return d.Clauses(dbresolver.Read)
}

func (d deployPrepareEventDo) WriteDB() IDeployPrepareEventDo {
	return d.Clauses(dbresolver.Write)
}

func (d deployPrepareEventDo) Session(config *gorm.Session) IDeployPrepareEventDo {
	return d.withDO(d.DO.Session(config))
}

func (d deployPrepareEventDo) Clauses(conds ...clause.Expression) IDeployPrepareEventDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deployPrepareEventDo) Returning(value interface{}, columns ...string) IDeployPrepareEventDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deployPrepareEventDo) Not(conds ...gen.Condition) IDeployPrepareEventDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deployPrepareEventDo) Or(conds ...gen.Condition) IDeployPrepareEventDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deployPrepareEventDo) Select(conds ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deployPrepareEventDo) Where(conds ...gen.Condition) IDeployPrepareEventDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deployPrepareEventDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeployPrepareEventDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deployPrepareEventDo) Order(conds ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deployPrepareEventDo) Distinct(cols ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deployPrepareEventDo) Omit(cols ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deployPrepareEventDo) Join(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deployPrepareEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deployPrepareEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deployPrepareEventDo) Group(cols ...field.Expr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deployPrepareEventDo) Having(conds ...gen.Condition) IDeployPrepareEventDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deployPrepareEventDo) Limit(limit int) IDeployPrepareEventDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deployPrepareEventDo) Offset(offset int) IDeployPrepareEventDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deployPrepareEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployPrepareEventDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deployPrepareEventDo) Unscoped() IDeployPrepareEventDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deployPrepareEventDo) Create(values ...*model.DeployPrepareEvent) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deployPrepareEventDo) CreateInBatches(values []*model.DeployPrepareEvent, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deployPrepareEventDo) Save(values ...*model.DeployPrepareEvent) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deployPrepareEventDo) First() (*model.DeployPrepareEvent, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployPrepareEvent), nil
	}
}

func (d deployPrepareEventDo) Take() (*model.DeployPrepareEvent, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployPrepareEvent), nil
	}
}

func (d deployPrepareEventDo) Last() (*model.DeployPrepareEvent, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployPrepareEvent), nil
	}
}

func (d deployPrepareEventDo) Find() ([]*model.DeployPrepareEvent, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeployPrepareEvent), err
}

func (d deployPrepareEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployPrepareEvent, err error) {
	buf := make([]*model.DeployPrepareEvent, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deployPrepareEventDo) FindInBatches(result *[]*model.DeployPrepareEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deployPrepareEventDo) Attrs(attrs ...field.AssignExpr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deployPrepareEventDo) Assign(attrs ...field.AssignExpr) IDeployPrepareEventDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deployPrepareEventDo) Joins(fields ...field.RelationField) IDeployPrepareEventDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deployPrepareEventDo) Preload(fields ...field.RelationField) IDeployPrepareEventDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deployPrepareEventDo) FirstOrInit() (*model.DeployPrepareEvent, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployPrepareEvent), nil
	}
}

func (d deployPrepareEventDo) FirstOrCreate() (*model.DeployPrepareEvent, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployPrepareEvent), nil
	}
}

func (d deployPrepareEventDo) FindByPage(offset int, limit int) (result []*model.DeployPrepareEvent, count int64, err error) {
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

func (d deployPrepareEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deployPrepareEventDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deployPrepareEventDo) Delete(models ...*model.DeployPrepareEvent) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deployPrepareEventDo) withDO(do gen.Dao) *deployPrepareEventDo {
	d.DO = *do.(*gen.DO)
	return d
}
