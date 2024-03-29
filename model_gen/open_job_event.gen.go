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

func newOpenJobEvent(db *gorm.DB, opts ...gen.DOOption) openJobEvent {
	_openJobEvent := openJobEvent{}

	_openJobEvent.openJobEventDo.UseDB(db, opts...)
	_openJobEvent.openJobEventDo.UseModel(&model.OpenJobEvent{})

	tableName := _openJobEvent.openJobEventDo.TableName()
	_openJobEvent.ALL = field.NewAsterisk(tableName)
	_openJobEvent.ID = field.NewInt32(tableName, "id")
	_openJobEvent.OpenJobEventID = field.NewString(tableName, "open_job_event_id")
	_openJobEvent.OpenJobEventDetail = field.NewString(tableName, "open_job_event_detail")
	_openJobEvent.ActionCodeID = field.NewString(tableName, "action_code_id")
	_openJobEvent.EventStatusCode = field.NewString(tableName, "event_status_code")
	_openJobEvent.OpUserID = field.NewString(tableName, "op_user_id")
	_openJobEvent.CreateTime = field.NewTime(tableName, "create_time")
	_openJobEvent.ModifyTime = field.NewTime(tableName, "modify_time")
	_openJobEvent.Comments = field.NewString(tableName, "comments")

	_openJobEvent.fillFieldMap()

	return _openJobEvent
}

type openJobEvent struct {
	openJobEventDo

	ALL                field.Asterisk
	ID                 field.Int32
	OpenJobEventID     field.String // 事件ID,uuid
	OpenJobEventDetail field.String // 事件细节
	ActionCodeID       field.String // sys_code.sys_type="open_job_action'
	/*
		sys_code.sys_code_type='open_job_action_event'n
		action_code_id+event_status_code 最终形成状态
	*/
	EventStatusCode field.String
	OpUserID        field.String // 执行人
	CreateTime      field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time   // 记录修改时间（数据库自动写入）
	Comments        field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (o openJobEvent) Table(newTableName string) *openJobEvent {
	o.openJobEventDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o openJobEvent) As(alias string) *openJobEvent {
	o.openJobEventDo.DO = *(o.openJobEventDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *openJobEvent) updateTableName(table string) *openJobEvent {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt32(table, "id")
	o.OpenJobEventID = field.NewString(table, "open_job_event_id")
	o.OpenJobEventDetail = field.NewString(table, "open_job_event_detail")
	o.ActionCodeID = field.NewString(table, "action_code_id")
	o.EventStatusCode = field.NewString(table, "event_status_code")
	o.OpUserID = field.NewString(table, "op_user_id")
	o.CreateTime = field.NewTime(table, "create_time")
	o.ModifyTime = field.NewTime(table, "modify_time")
	o.Comments = field.NewString(table, "comments")

	o.fillFieldMap()

	return o
}

func (o *openJobEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *openJobEvent) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 9)
	o.fieldMap["id"] = o.ID
	o.fieldMap["open_job_event_id"] = o.OpenJobEventID
	o.fieldMap["open_job_event_detail"] = o.OpenJobEventDetail
	o.fieldMap["action_code_id"] = o.ActionCodeID
	o.fieldMap["event_status_code"] = o.EventStatusCode
	o.fieldMap["op_user_id"] = o.OpUserID
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["modify_time"] = o.ModifyTime
	o.fieldMap["comments"] = o.Comments
}

func (o openJobEvent) clone(db *gorm.DB) openJobEvent {
	o.openJobEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o openJobEvent) replaceDB(db *gorm.DB) openJobEvent {
	o.openJobEventDo.ReplaceDB(db)
	return o
}

type openJobEventDo struct{ gen.DO }

type IOpenJobEventDo interface {
	gen.SubQuery
	Debug() IOpenJobEventDo
	WithContext(ctx context.Context) IOpenJobEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOpenJobEventDo
	WriteDB() IOpenJobEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOpenJobEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOpenJobEventDo
	Not(conds ...gen.Condition) IOpenJobEventDo
	Or(conds ...gen.Condition) IOpenJobEventDo
	Select(conds ...field.Expr) IOpenJobEventDo
	Where(conds ...gen.Condition) IOpenJobEventDo
	Order(conds ...field.Expr) IOpenJobEventDo
	Distinct(cols ...field.Expr) IOpenJobEventDo
	Omit(cols ...field.Expr) IOpenJobEventDo
	Join(table schema.Tabler, on ...field.Expr) IOpenJobEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOpenJobEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOpenJobEventDo
	Group(cols ...field.Expr) IOpenJobEventDo
	Having(conds ...gen.Condition) IOpenJobEventDo
	Limit(limit int) IOpenJobEventDo
	Offset(offset int) IOpenJobEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOpenJobEventDo
	Unscoped() IOpenJobEventDo
	Create(values ...*model.OpenJobEvent) error
	CreateInBatches(values []*model.OpenJobEvent, batchSize int) error
	Save(values ...*model.OpenJobEvent) error
	First() (*model.OpenJobEvent, error)
	Take() (*model.OpenJobEvent, error)
	Last() (*model.OpenJobEvent, error)
	Find() ([]*model.OpenJobEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpenJobEvent, err error)
	FindInBatches(result *[]*model.OpenJobEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OpenJobEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOpenJobEventDo
	Assign(attrs ...field.AssignExpr) IOpenJobEventDo
	Joins(fields ...field.RelationField) IOpenJobEventDo
	Preload(fields ...field.RelationField) IOpenJobEventDo
	FirstOrInit() (*model.OpenJobEvent, error)
	FirstOrCreate() (*model.OpenJobEvent, error)
	FindByPage(offset int, limit int) (result []*model.OpenJobEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOpenJobEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o openJobEventDo) Debug() IOpenJobEventDo {
	return o.withDO(o.DO.Debug())
}

func (o openJobEventDo) WithContext(ctx context.Context) IOpenJobEventDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o openJobEventDo) ReadDB() IOpenJobEventDo {
	return o.Clauses(dbresolver.Read)
}

func (o openJobEventDo) WriteDB() IOpenJobEventDo {
	return o.Clauses(dbresolver.Write)
}

func (o openJobEventDo) Session(config *gorm.Session) IOpenJobEventDo {
	return o.withDO(o.DO.Session(config))
}

func (o openJobEventDo) Clauses(conds ...clause.Expression) IOpenJobEventDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o openJobEventDo) Returning(value interface{}, columns ...string) IOpenJobEventDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o openJobEventDo) Not(conds ...gen.Condition) IOpenJobEventDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o openJobEventDo) Or(conds ...gen.Condition) IOpenJobEventDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o openJobEventDo) Select(conds ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o openJobEventDo) Where(conds ...gen.Condition) IOpenJobEventDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o openJobEventDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOpenJobEventDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o openJobEventDo) Order(conds ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o openJobEventDo) Distinct(cols ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o openJobEventDo) Omit(cols ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o openJobEventDo) Join(table schema.Tabler, on ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o openJobEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o openJobEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o openJobEventDo) Group(cols ...field.Expr) IOpenJobEventDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o openJobEventDo) Having(conds ...gen.Condition) IOpenJobEventDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o openJobEventDo) Limit(limit int) IOpenJobEventDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o openJobEventDo) Offset(offset int) IOpenJobEventDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o openJobEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOpenJobEventDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o openJobEventDo) Unscoped() IOpenJobEventDo {
	return o.withDO(o.DO.Unscoped())
}

func (o openJobEventDo) Create(values ...*model.OpenJobEvent) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o openJobEventDo) CreateInBatches(values []*model.OpenJobEvent, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o openJobEventDo) Save(values ...*model.OpenJobEvent) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o openJobEventDo) First() (*model.OpenJobEvent, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenJobEvent), nil
	}
}

func (o openJobEventDo) Take() (*model.OpenJobEvent, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenJobEvent), nil
	}
}

func (o openJobEventDo) Last() (*model.OpenJobEvent, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenJobEvent), nil
	}
}

func (o openJobEventDo) Find() ([]*model.OpenJobEvent, error) {
	result, err := o.DO.Find()
	return result.([]*model.OpenJobEvent), err
}

func (o openJobEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpenJobEvent, err error) {
	buf := make([]*model.OpenJobEvent, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o openJobEventDo) FindInBatches(result *[]*model.OpenJobEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o openJobEventDo) Attrs(attrs ...field.AssignExpr) IOpenJobEventDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o openJobEventDo) Assign(attrs ...field.AssignExpr) IOpenJobEventDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o openJobEventDo) Joins(fields ...field.RelationField) IOpenJobEventDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o openJobEventDo) Preload(fields ...field.RelationField) IOpenJobEventDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o openJobEventDo) FirstOrInit() (*model.OpenJobEvent, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenJobEvent), nil
	}
}

func (o openJobEventDo) FirstOrCreate() (*model.OpenJobEvent, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenJobEvent), nil
	}
}

func (o openJobEventDo) FindByPage(offset int, limit int) (result []*model.OpenJobEvent, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o openJobEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o openJobEventDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o openJobEventDo) Delete(models ...*model.OpenJobEvent) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *openJobEventDo) withDO(do gen.Dao) *openJobEventDo {
	o.DO = *do.(*gen.DO)
	return o
}
