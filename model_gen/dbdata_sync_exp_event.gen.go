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

func newDbdataSyncExpEvent(db *gorm.DB, opts ...gen.DOOption) dbdataSyncExpEvent {
	_dbdataSyncExpEvent := dbdataSyncExpEvent{}

	_dbdataSyncExpEvent.dbdataSyncExpEventDo.UseDB(db, opts...)
	_dbdataSyncExpEvent.dbdataSyncExpEventDo.UseModel(&model.DbdataSyncExpEvent{})

	tableName := _dbdataSyncExpEvent.dbdataSyncExpEventDo.TableName()
	_dbdataSyncExpEvent.ALL = field.NewAsterisk(tableName)
	_dbdataSyncExpEvent.DbsyncExpEventID = field.NewString(tableName, "dbsync_exp_event_id")
	_dbdataSyncExpEvent.SoftUserID = field.NewString(tableName, "soft_user_id")
	_dbdataSyncExpEvent.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_dbdataSyncExpEvent.TaskTitle = field.NewString(tableName, "task_title")
	_dbdataSyncExpEvent.SyncStartTime = field.NewTime(tableName, "sync_start_time")
	_dbdataSyncExpEvent.SyncEndTime = field.NewTime(tableName, "sync_end_time")
	_dbdataSyncExpEvent.StepIndex = field.NewString(tableName, "step_index")
	_dbdataSyncExpEvent.SyncCurrentStep = field.NewString(tableName, "sync_current_step")
	_dbdataSyncExpEvent.SyncLog = field.NewString(tableName, "sync_log")
	_dbdataSyncExpEvent.SyncScript = field.NewString(tableName, "sync_script")
	_dbdataSyncExpEvent.AddTime = field.NewTime(tableName, "add_time")
	_dbdataSyncExpEvent.OpUser = field.NewString(tableName, "op_user")
	_dbdataSyncExpEvent.Note = field.NewString(tableName, "note")
	_dbdataSyncExpEvent.SyncStatu = field.NewString(tableName, "sync_statu")
	_dbdataSyncExpEvent.UploadFileName = field.NewString(tableName, "upload_file_name")
	_dbdataSyncExpEvent.UploadFilePath = field.NewString(tableName, "upload_file_path")
	_dbdataSyncExpEvent.DirectoryPath = field.NewString(tableName, "directory_path")
	_dbdataSyncExpEvent.OpPercent = field.NewString(tableName, "op_percent")
	_dbdataSyncExpEvent.CanEditor = field.NewString(tableName, "can_editor")
	_dbdataSyncExpEvent.StransactionIdx = field.NewString(tableName, "stransaction_idx")
	_dbdataSyncExpEvent.ExpOraVersion = field.NewString(tableName, "exp_ora_version")

	_dbdataSyncExpEvent.fillFieldMap()

	return _dbdataSyncExpEvent
}

type dbdataSyncExpEvent struct {
	dbdataSyncExpEventDo

	ALL              field.Asterisk
	DbsyncExpEventID field.String
	SoftUserID       field.String
	InstallSoftID    field.String
	TaskTitle        field.String
	SyncStartTime    field.Time
	SyncEndTime      field.Time
	StepIndex        field.String
	SyncCurrentStep  field.String
	SyncLog          field.String
	SyncScript       field.String
	AddTime          field.Time
	OpUser           field.String
	Note             field.String
	SyncStatu        field.String
	UploadFileName   field.String
	UploadFilePath   field.String
	DirectoryPath    field.String
	OpPercent        field.String
	CanEditor        field.String
	StransactionIdx  field.String
	ExpOraVersion    field.String

	fieldMap map[string]field.Expr
}

func (d dbdataSyncExpEvent) Table(newTableName string) *dbdataSyncExpEvent {
	d.dbdataSyncExpEventDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataSyncExpEvent) As(alias string) *dbdataSyncExpEvent {
	d.dbdataSyncExpEventDo.DO = *(d.dbdataSyncExpEventDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataSyncExpEvent) updateTableName(table string) *dbdataSyncExpEvent {
	d.ALL = field.NewAsterisk(table)
	d.DbsyncExpEventID = field.NewString(table, "dbsync_exp_event_id")
	d.SoftUserID = field.NewString(table, "soft_user_id")
	d.InstallSoftID = field.NewString(table, "install_soft_id")
	d.TaskTitle = field.NewString(table, "task_title")
	d.SyncStartTime = field.NewTime(table, "sync_start_time")
	d.SyncEndTime = field.NewTime(table, "sync_end_time")
	d.StepIndex = field.NewString(table, "step_index")
	d.SyncCurrentStep = field.NewString(table, "sync_current_step")
	d.SyncLog = field.NewString(table, "sync_log")
	d.SyncScript = field.NewString(table, "sync_script")
	d.AddTime = field.NewTime(table, "add_time")
	d.OpUser = field.NewString(table, "op_user")
	d.Note = field.NewString(table, "note")
	d.SyncStatu = field.NewString(table, "sync_statu")
	d.UploadFileName = field.NewString(table, "upload_file_name")
	d.UploadFilePath = field.NewString(table, "upload_file_path")
	d.DirectoryPath = field.NewString(table, "directory_path")
	d.OpPercent = field.NewString(table, "op_percent")
	d.CanEditor = field.NewString(table, "can_editor")
	d.StransactionIdx = field.NewString(table, "stransaction_idx")
	d.ExpOraVersion = field.NewString(table, "exp_ora_version")

	d.fillFieldMap()

	return d
}

func (d *dbdataSyncExpEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataSyncExpEvent) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 21)
	d.fieldMap["dbsync_exp_event_id"] = d.DbsyncExpEventID
	d.fieldMap["soft_user_id"] = d.SoftUserID
	d.fieldMap["install_soft_id"] = d.InstallSoftID
	d.fieldMap["task_title"] = d.TaskTitle
	d.fieldMap["sync_start_time"] = d.SyncStartTime
	d.fieldMap["sync_end_time"] = d.SyncEndTime
	d.fieldMap["step_index"] = d.StepIndex
	d.fieldMap["sync_current_step"] = d.SyncCurrentStep
	d.fieldMap["sync_log"] = d.SyncLog
	d.fieldMap["sync_script"] = d.SyncScript
	d.fieldMap["add_time"] = d.AddTime
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["note"] = d.Note
	d.fieldMap["sync_statu"] = d.SyncStatu
	d.fieldMap["upload_file_name"] = d.UploadFileName
	d.fieldMap["upload_file_path"] = d.UploadFilePath
	d.fieldMap["directory_path"] = d.DirectoryPath
	d.fieldMap["op_percent"] = d.OpPercent
	d.fieldMap["can_editor"] = d.CanEditor
	d.fieldMap["stransaction_idx"] = d.StransactionIdx
	d.fieldMap["exp_ora_version"] = d.ExpOraVersion
}

func (d dbdataSyncExpEvent) clone(db *gorm.DB) dbdataSyncExpEvent {
	d.dbdataSyncExpEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataSyncExpEvent) replaceDB(db *gorm.DB) dbdataSyncExpEvent {
	d.dbdataSyncExpEventDo.ReplaceDB(db)
	return d
}

type dbdataSyncExpEventDo struct{ gen.DO }

type IDbdataSyncExpEventDo interface {
	gen.SubQuery
	Debug() IDbdataSyncExpEventDo
	WithContext(ctx context.Context) IDbdataSyncExpEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataSyncExpEventDo
	WriteDB() IDbdataSyncExpEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataSyncExpEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataSyncExpEventDo
	Not(conds ...gen.Condition) IDbdataSyncExpEventDo
	Or(conds ...gen.Condition) IDbdataSyncExpEventDo
	Select(conds ...field.Expr) IDbdataSyncExpEventDo
	Where(conds ...gen.Condition) IDbdataSyncExpEventDo
	Order(conds ...field.Expr) IDbdataSyncExpEventDo
	Distinct(cols ...field.Expr) IDbdataSyncExpEventDo
	Omit(cols ...field.Expr) IDbdataSyncExpEventDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo
	Group(cols ...field.Expr) IDbdataSyncExpEventDo
	Having(conds ...gen.Condition) IDbdataSyncExpEventDo
	Limit(limit int) IDbdataSyncExpEventDo
	Offset(offset int) IDbdataSyncExpEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataSyncExpEventDo
	Unscoped() IDbdataSyncExpEventDo
	Create(values ...*model.DbdataSyncExpEvent) error
	CreateInBatches(values []*model.DbdataSyncExpEvent, batchSize int) error
	Save(values ...*model.DbdataSyncExpEvent) error
	First() (*model.DbdataSyncExpEvent, error)
	Take() (*model.DbdataSyncExpEvent, error)
	Last() (*model.DbdataSyncExpEvent, error)
	Find() ([]*model.DbdataSyncExpEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataSyncExpEvent, err error)
	FindInBatches(result *[]*model.DbdataSyncExpEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataSyncExpEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataSyncExpEventDo
	Assign(attrs ...field.AssignExpr) IDbdataSyncExpEventDo
	Joins(fields ...field.RelationField) IDbdataSyncExpEventDo
	Preload(fields ...field.RelationField) IDbdataSyncExpEventDo
	FirstOrInit() (*model.DbdataSyncExpEvent, error)
	FirstOrCreate() (*model.DbdataSyncExpEvent, error)
	FindByPage(offset int, limit int) (result []*model.DbdataSyncExpEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataSyncExpEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataSyncExpEventDo) Debug() IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataSyncExpEventDo) WithContext(ctx context.Context) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataSyncExpEventDo) ReadDB() IDbdataSyncExpEventDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataSyncExpEventDo) WriteDB() IDbdataSyncExpEventDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataSyncExpEventDo) Session(config *gorm.Session) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataSyncExpEventDo) Clauses(conds ...clause.Expression) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataSyncExpEventDo) Returning(value interface{}, columns ...string) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataSyncExpEventDo) Not(conds ...gen.Condition) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataSyncExpEventDo) Or(conds ...gen.Condition) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataSyncExpEventDo) Select(conds ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataSyncExpEventDo) Where(conds ...gen.Condition) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataSyncExpEventDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataSyncExpEventDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataSyncExpEventDo) Order(conds ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataSyncExpEventDo) Distinct(cols ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataSyncExpEventDo) Omit(cols ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataSyncExpEventDo) Join(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataSyncExpEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataSyncExpEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataSyncExpEventDo) Group(cols ...field.Expr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataSyncExpEventDo) Having(conds ...gen.Condition) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataSyncExpEventDo) Limit(limit int) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataSyncExpEventDo) Offset(offset int) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataSyncExpEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataSyncExpEventDo) Unscoped() IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataSyncExpEventDo) Create(values ...*model.DbdataSyncExpEvent) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataSyncExpEventDo) CreateInBatches(values []*model.DbdataSyncExpEvent, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataSyncExpEventDo) Save(values ...*model.DbdataSyncExpEvent) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataSyncExpEventDo) First() (*model.DbdataSyncExpEvent, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataSyncExpEvent), nil
	}
}

func (d dbdataSyncExpEventDo) Take() (*model.DbdataSyncExpEvent, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataSyncExpEvent), nil
	}
}

func (d dbdataSyncExpEventDo) Last() (*model.DbdataSyncExpEvent, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataSyncExpEvent), nil
	}
}

func (d dbdataSyncExpEventDo) Find() ([]*model.DbdataSyncExpEvent, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataSyncExpEvent), err
}

func (d dbdataSyncExpEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataSyncExpEvent, err error) {
	buf := make([]*model.DbdataSyncExpEvent, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataSyncExpEventDo) FindInBatches(result *[]*model.DbdataSyncExpEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataSyncExpEventDo) Attrs(attrs ...field.AssignExpr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataSyncExpEventDo) Assign(attrs ...field.AssignExpr) IDbdataSyncExpEventDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataSyncExpEventDo) Joins(fields ...field.RelationField) IDbdataSyncExpEventDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataSyncExpEventDo) Preload(fields ...field.RelationField) IDbdataSyncExpEventDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataSyncExpEventDo) FirstOrInit() (*model.DbdataSyncExpEvent, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataSyncExpEvent), nil
	}
}

func (d dbdataSyncExpEventDo) FirstOrCreate() (*model.DbdataSyncExpEvent, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataSyncExpEvent), nil
	}
}

func (d dbdataSyncExpEventDo) FindByPage(offset int, limit int) (result []*model.DbdataSyncExpEvent, count int64, err error) {
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

func (d dbdataSyncExpEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataSyncExpEventDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataSyncExpEventDo) Delete(models ...*model.DbdataSyncExpEvent) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataSyncExpEventDo) withDO(do gen.Dao) *dbdataSyncExpEventDo {
	d.DO = *do.(*gen.DO)
	return d
}