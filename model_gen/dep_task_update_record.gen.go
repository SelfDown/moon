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

func newDepTaskUpdateRecord(db *gorm.DB, opts ...gen.DOOption) depTaskUpdateRecord {
	_depTaskUpdateRecord := depTaskUpdateRecord{}

	_depTaskUpdateRecord.depTaskUpdateRecordDo.UseDB(db, opts...)
	_depTaskUpdateRecord.depTaskUpdateRecordDo.UseModel(&model.DepTaskUpdateRecord{})

	tableName := _depTaskUpdateRecord.depTaskUpdateRecordDo.TableName()
	_depTaskUpdateRecord.ALL = field.NewAsterisk(tableName)
	_depTaskUpdateRecord.DepTaskUpdateRecordID = field.NewString(tableName, "dep_task_update_record_id")
	_depTaskUpdateRecord.DepTaskID = field.NewString(tableName, "dep_task_id")
	_depTaskUpdateRecord.Version = field.NewString(tableName, "version")
	_depTaskUpdateRecord.CreateUser = field.NewString(tableName, "create_user")
	_depTaskUpdateRecord.OpType = field.NewString(tableName, "op_type")
	_depTaskUpdateRecord.CreateTime = field.NewString(tableName, "create_time")

	_depTaskUpdateRecord.fillFieldMap()

	return _depTaskUpdateRecord
}

type depTaskUpdateRecord struct {
	depTaskUpdateRecordDo

	ALL                   field.Asterisk
	DepTaskUpdateRecordID field.String
	DepTaskID             field.String // 部署任务
	Version               field.String // 版本号
	CreateUser            field.String
	OpType                field.String
	CreateTime            field.String

	fieldMap map[string]field.Expr
}

func (d depTaskUpdateRecord) Table(newTableName string) *depTaskUpdateRecord {
	d.depTaskUpdateRecordDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d depTaskUpdateRecord) As(alias string) *depTaskUpdateRecord {
	d.depTaskUpdateRecordDo.DO = *(d.depTaskUpdateRecordDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *depTaskUpdateRecord) updateTableName(table string) *depTaskUpdateRecord {
	d.ALL = field.NewAsterisk(table)
	d.DepTaskUpdateRecordID = field.NewString(table, "dep_task_update_record_id")
	d.DepTaskID = field.NewString(table, "dep_task_id")
	d.Version = field.NewString(table, "version")
	d.CreateUser = field.NewString(table, "create_user")
	d.OpType = field.NewString(table, "op_type")
	d.CreateTime = field.NewString(table, "create_time")

	d.fillFieldMap()

	return d
}

func (d *depTaskUpdateRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *depTaskUpdateRecord) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["dep_task_update_record_id"] = d.DepTaskUpdateRecordID
	d.fieldMap["dep_task_id"] = d.DepTaskID
	d.fieldMap["version"] = d.Version
	d.fieldMap["create_user"] = d.CreateUser
	d.fieldMap["op_type"] = d.OpType
	d.fieldMap["create_time"] = d.CreateTime
}

func (d depTaskUpdateRecord) clone(db *gorm.DB) depTaskUpdateRecord {
	d.depTaskUpdateRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d depTaskUpdateRecord) replaceDB(db *gorm.DB) depTaskUpdateRecord {
	d.depTaskUpdateRecordDo.ReplaceDB(db)
	return d
}

type depTaskUpdateRecordDo struct{ gen.DO }

type IDepTaskUpdateRecordDo interface {
	gen.SubQuery
	Debug() IDepTaskUpdateRecordDo
	WithContext(ctx context.Context) IDepTaskUpdateRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDepTaskUpdateRecordDo
	WriteDB() IDepTaskUpdateRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDepTaskUpdateRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDepTaskUpdateRecordDo
	Not(conds ...gen.Condition) IDepTaskUpdateRecordDo
	Or(conds ...gen.Condition) IDepTaskUpdateRecordDo
	Select(conds ...field.Expr) IDepTaskUpdateRecordDo
	Where(conds ...gen.Condition) IDepTaskUpdateRecordDo
	Order(conds ...field.Expr) IDepTaskUpdateRecordDo
	Distinct(cols ...field.Expr) IDepTaskUpdateRecordDo
	Omit(cols ...field.Expr) IDepTaskUpdateRecordDo
	Join(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo
	Group(cols ...field.Expr) IDepTaskUpdateRecordDo
	Having(conds ...gen.Condition) IDepTaskUpdateRecordDo
	Limit(limit int) IDepTaskUpdateRecordDo
	Offset(offset int) IDepTaskUpdateRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDepTaskUpdateRecordDo
	Unscoped() IDepTaskUpdateRecordDo
	Create(values ...*model.DepTaskUpdateRecord) error
	CreateInBatches(values []*model.DepTaskUpdateRecord, batchSize int) error
	Save(values ...*model.DepTaskUpdateRecord) error
	First() (*model.DepTaskUpdateRecord, error)
	Take() (*model.DepTaskUpdateRecord, error)
	Last() (*model.DepTaskUpdateRecord, error)
	Find() ([]*model.DepTaskUpdateRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DepTaskUpdateRecord, err error)
	FindInBatches(result *[]*model.DepTaskUpdateRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DepTaskUpdateRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDepTaskUpdateRecordDo
	Assign(attrs ...field.AssignExpr) IDepTaskUpdateRecordDo
	Joins(fields ...field.RelationField) IDepTaskUpdateRecordDo
	Preload(fields ...field.RelationField) IDepTaskUpdateRecordDo
	FirstOrInit() (*model.DepTaskUpdateRecord, error)
	FirstOrCreate() (*model.DepTaskUpdateRecord, error)
	FindByPage(offset int, limit int) (result []*model.DepTaskUpdateRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDepTaskUpdateRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d depTaskUpdateRecordDo) Debug() IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Debug())
}

func (d depTaskUpdateRecordDo) WithContext(ctx context.Context) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d depTaskUpdateRecordDo) ReadDB() IDepTaskUpdateRecordDo {
	return d.Clauses(dbresolver.Read)
}

func (d depTaskUpdateRecordDo) WriteDB() IDepTaskUpdateRecordDo {
	return d.Clauses(dbresolver.Write)
}

func (d depTaskUpdateRecordDo) Session(config *gorm.Session) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Session(config))
}

func (d depTaskUpdateRecordDo) Clauses(conds ...clause.Expression) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d depTaskUpdateRecordDo) Returning(value interface{}, columns ...string) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d depTaskUpdateRecordDo) Not(conds ...gen.Condition) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d depTaskUpdateRecordDo) Or(conds ...gen.Condition) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d depTaskUpdateRecordDo) Select(conds ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d depTaskUpdateRecordDo) Where(conds ...gen.Condition) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d depTaskUpdateRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDepTaskUpdateRecordDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d depTaskUpdateRecordDo) Order(conds ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d depTaskUpdateRecordDo) Distinct(cols ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d depTaskUpdateRecordDo) Omit(cols ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d depTaskUpdateRecordDo) Join(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d depTaskUpdateRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d depTaskUpdateRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d depTaskUpdateRecordDo) Group(cols ...field.Expr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d depTaskUpdateRecordDo) Having(conds ...gen.Condition) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d depTaskUpdateRecordDo) Limit(limit int) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d depTaskUpdateRecordDo) Offset(offset int) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d depTaskUpdateRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d depTaskUpdateRecordDo) Unscoped() IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Unscoped())
}

func (d depTaskUpdateRecordDo) Create(values ...*model.DepTaskUpdateRecord) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d depTaskUpdateRecordDo) CreateInBatches(values []*model.DepTaskUpdateRecord, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d depTaskUpdateRecordDo) Save(values ...*model.DepTaskUpdateRecord) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d depTaskUpdateRecordDo) First() (*model.DepTaskUpdateRecord, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DepTaskUpdateRecord), nil
	}
}

func (d depTaskUpdateRecordDo) Take() (*model.DepTaskUpdateRecord, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DepTaskUpdateRecord), nil
	}
}

func (d depTaskUpdateRecordDo) Last() (*model.DepTaskUpdateRecord, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DepTaskUpdateRecord), nil
	}
}

func (d depTaskUpdateRecordDo) Find() ([]*model.DepTaskUpdateRecord, error) {
	result, err := d.DO.Find()
	return result.([]*model.DepTaskUpdateRecord), err
}

func (d depTaskUpdateRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DepTaskUpdateRecord, err error) {
	buf := make([]*model.DepTaskUpdateRecord, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d depTaskUpdateRecordDo) FindInBatches(result *[]*model.DepTaskUpdateRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d depTaskUpdateRecordDo) Attrs(attrs ...field.AssignExpr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d depTaskUpdateRecordDo) Assign(attrs ...field.AssignExpr) IDepTaskUpdateRecordDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d depTaskUpdateRecordDo) Joins(fields ...field.RelationField) IDepTaskUpdateRecordDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d depTaskUpdateRecordDo) Preload(fields ...field.RelationField) IDepTaskUpdateRecordDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d depTaskUpdateRecordDo) FirstOrInit() (*model.DepTaskUpdateRecord, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DepTaskUpdateRecord), nil
	}
}

func (d depTaskUpdateRecordDo) FirstOrCreate() (*model.DepTaskUpdateRecord, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DepTaskUpdateRecord), nil
	}
}

func (d depTaskUpdateRecordDo) FindByPage(offset int, limit int) (result []*model.DepTaskUpdateRecord, count int64, err error) {
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

func (d depTaskUpdateRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d depTaskUpdateRecordDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d depTaskUpdateRecordDo) Delete(models ...*model.DepTaskUpdateRecord) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *depTaskUpdateRecordDo) withDO(do gen.Dao) *depTaskUpdateRecordDo {
	d.DO = *do.(*gen.DO)
	return d
}