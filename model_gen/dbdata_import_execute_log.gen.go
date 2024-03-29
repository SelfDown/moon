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

func newDbdataImportExecuteLog(db *gorm.DB, opts ...gen.DOOption) dbdataImportExecuteLog {
	_dbdataImportExecuteLog := dbdataImportExecuteLog{}

	_dbdataImportExecuteLog.dbdataImportExecuteLogDo.UseDB(db, opts...)
	_dbdataImportExecuteLog.dbdataImportExecuteLogDo.UseModel(&model.DbdataImportExecuteLog{})

	tableName := _dbdataImportExecuteLog.dbdataImportExecuteLogDo.TableName()
	_dbdataImportExecuteLog.ALL = field.NewAsterisk(tableName)
	_dbdataImportExecuteLog.DbdataImportExecuteLogID = field.NewString(tableName, "dbdata_import_execute_log_id")
	_dbdataImportExecuteLog.DbdataImportID = field.NewString(tableName, "dbdata_import_id")
	_dbdataImportExecuteLog.OpUser = field.NewString(tableName, "op_user")
	_dbdataImportExecuteLog.SyncStartTime = field.NewString(tableName, "sync_start_time")
	_dbdataImportExecuteLog.SyncEndTime = field.NewString(tableName, "sync_end_time")
	_dbdataImportExecuteLog.AddTime = field.NewString(tableName, "add_time")
	_dbdataImportExecuteLog.SyncStatu = field.NewString(tableName, "sync_statu")
	_dbdataImportExecuteLog.SyncScript = field.NewString(tableName, "sync_script")
	_dbdataImportExecuteLog.SyncLog = field.NewString(tableName, "sync_log")
	_dbdataImportExecuteLog.Size = field.NewInt64(tableName, "size")
	_dbdataImportExecuteLog.Msg = field.NewString(tableName, "msg")
	_dbdataImportExecuteLog.RemoteShellPath = field.NewString(tableName, "remote_shell_path")
	_dbdataImportExecuteLog.LogPath = field.NewString(tableName, "log_path")
	_dbdataImportExecuteLog.TarFilename = field.NewString(tableName, "tar_filename")
	_dbdataImportExecuteLog.TarDir = field.NewString(tableName, "tar_dir")
	_dbdataImportExecuteLog.RemoteDumpTarName = field.NewString(tableName, "remote_dump_tar_name")
	_dbdataImportExecuteLog.RemoteDumpTarDir = field.NewString(tableName, "remote_dump_tar_dir")
	_dbdataImportExecuteLog.LocalDumpPath = field.NewString(tableName, "local_dump_path")
	_dbdataImportExecuteLog.LocalDumpPathAbs = field.NewString(tableName, "local_dump_path_abs")
	_dbdataImportExecuteLog.SftpDumpPath = field.NewString(tableName, "sftp_dump_path")
	_dbdataImportExecuteLog.ErrorCount = field.NewString(tableName, "error_count")

	_dbdataImportExecuteLog.fillFieldMap()

	return _dbdataImportExecuteLog
}

type dbdataImportExecuteLog struct {
	dbdataImportExecuteLogDo

	ALL                      field.Asterisk
	DbdataImportExecuteLogID field.String
	DbdataImportID           field.String
	OpUser                   field.String
	SyncStartTime            field.String
	SyncEndTime              field.String
	AddTime                  field.String
	SyncStatu                field.String
	SyncScript               field.String
	SyncLog                  field.String
	Size                     field.Int64
	Msg                      field.String
	RemoteShellPath          field.String
	LogPath                  field.String
	TarFilename              field.String
	TarDir                   field.String
	RemoteDumpTarName        field.String
	RemoteDumpTarDir         field.String
	LocalDumpPath            field.String
	LocalDumpPathAbs         field.String
	SftpDumpPath             field.String
	ErrorCount               field.String

	fieldMap map[string]field.Expr
}

func (d dbdataImportExecuteLog) Table(newTableName string) *dbdataImportExecuteLog {
	d.dbdataImportExecuteLogDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataImportExecuteLog) As(alias string) *dbdataImportExecuteLog {
	d.dbdataImportExecuteLogDo.DO = *(d.dbdataImportExecuteLogDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataImportExecuteLog) updateTableName(table string) *dbdataImportExecuteLog {
	d.ALL = field.NewAsterisk(table)
	d.DbdataImportExecuteLogID = field.NewString(table, "dbdata_import_execute_log_id")
	d.DbdataImportID = field.NewString(table, "dbdata_import_id")
	d.OpUser = field.NewString(table, "op_user")
	d.SyncStartTime = field.NewString(table, "sync_start_time")
	d.SyncEndTime = field.NewString(table, "sync_end_time")
	d.AddTime = field.NewString(table, "add_time")
	d.SyncStatu = field.NewString(table, "sync_statu")
	d.SyncScript = field.NewString(table, "sync_script")
	d.SyncLog = field.NewString(table, "sync_log")
	d.Size = field.NewInt64(table, "size")
	d.Msg = field.NewString(table, "msg")
	d.RemoteShellPath = field.NewString(table, "remote_shell_path")
	d.LogPath = field.NewString(table, "log_path")
	d.TarFilename = field.NewString(table, "tar_filename")
	d.TarDir = field.NewString(table, "tar_dir")
	d.RemoteDumpTarName = field.NewString(table, "remote_dump_tar_name")
	d.RemoteDumpTarDir = field.NewString(table, "remote_dump_tar_dir")
	d.LocalDumpPath = field.NewString(table, "local_dump_path")
	d.LocalDumpPathAbs = field.NewString(table, "local_dump_path_abs")
	d.SftpDumpPath = field.NewString(table, "sftp_dump_path")
	d.ErrorCount = field.NewString(table, "error_count")

	d.fillFieldMap()

	return d
}

func (d *dbdataImportExecuteLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataImportExecuteLog) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 21)
	d.fieldMap["dbdata_import_execute_log_id"] = d.DbdataImportExecuteLogID
	d.fieldMap["dbdata_import_id"] = d.DbdataImportID
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["sync_start_time"] = d.SyncStartTime
	d.fieldMap["sync_end_time"] = d.SyncEndTime
	d.fieldMap["add_time"] = d.AddTime
	d.fieldMap["sync_statu"] = d.SyncStatu
	d.fieldMap["sync_script"] = d.SyncScript
	d.fieldMap["sync_log"] = d.SyncLog
	d.fieldMap["size"] = d.Size
	d.fieldMap["msg"] = d.Msg
	d.fieldMap["remote_shell_path"] = d.RemoteShellPath
	d.fieldMap["log_path"] = d.LogPath
	d.fieldMap["tar_filename"] = d.TarFilename
	d.fieldMap["tar_dir"] = d.TarDir
	d.fieldMap["remote_dump_tar_name"] = d.RemoteDumpTarName
	d.fieldMap["remote_dump_tar_dir"] = d.RemoteDumpTarDir
	d.fieldMap["local_dump_path"] = d.LocalDumpPath
	d.fieldMap["local_dump_path_abs"] = d.LocalDumpPathAbs
	d.fieldMap["sftp_dump_path"] = d.SftpDumpPath
	d.fieldMap["error_count"] = d.ErrorCount
}

func (d dbdataImportExecuteLog) clone(db *gorm.DB) dbdataImportExecuteLog {
	d.dbdataImportExecuteLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataImportExecuteLog) replaceDB(db *gorm.DB) dbdataImportExecuteLog {
	d.dbdataImportExecuteLogDo.ReplaceDB(db)
	return d
}

type dbdataImportExecuteLogDo struct{ gen.DO }

type IDbdataImportExecuteLogDo interface {
	gen.SubQuery
	Debug() IDbdataImportExecuteLogDo
	WithContext(ctx context.Context) IDbdataImportExecuteLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataImportExecuteLogDo
	WriteDB() IDbdataImportExecuteLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataImportExecuteLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataImportExecuteLogDo
	Not(conds ...gen.Condition) IDbdataImportExecuteLogDo
	Or(conds ...gen.Condition) IDbdataImportExecuteLogDo
	Select(conds ...field.Expr) IDbdataImportExecuteLogDo
	Where(conds ...gen.Condition) IDbdataImportExecuteLogDo
	Order(conds ...field.Expr) IDbdataImportExecuteLogDo
	Distinct(cols ...field.Expr) IDbdataImportExecuteLogDo
	Omit(cols ...field.Expr) IDbdataImportExecuteLogDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo
	Group(cols ...field.Expr) IDbdataImportExecuteLogDo
	Having(conds ...gen.Condition) IDbdataImportExecuteLogDo
	Limit(limit int) IDbdataImportExecuteLogDo
	Offset(offset int) IDbdataImportExecuteLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataImportExecuteLogDo
	Unscoped() IDbdataImportExecuteLogDo
	Create(values ...*model.DbdataImportExecuteLog) error
	CreateInBatches(values []*model.DbdataImportExecuteLog, batchSize int) error
	Save(values ...*model.DbdataImportExecuteLog) error
	First() (*model.DbdataImportExecuteLog, error)
	Take() (*model.DbdataImportExecuteLog, error)
	Last() (*model.DbdataImportExecuteLog, error)
	Find() ([]*model.DbdataImportExecuteLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataImportExecuteLog, err error)
	FindInBatches(result *[]*model.DbdataImportExecuteLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataImportExecuteLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataImportExecuteLogDo
	Assign(attrs ...field.AssignExpr) IDbdataImportExecuteLogDo
	Joins(fields ...field.RelationField) IDbdataImportExecuteLogDo
	Preload(fields ...field.RelationField) IDbdataImportExecuteLogDo
	FirstOrInit() (*model.DbdataImportExecuteLog, error)
	FirstOrCreate() (*model.DbdataImportExecuteLog, error)
	FindByPage(offset int, limit int) (result []*model.DbdataImportExecuteLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataImportExecuteLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataImportExecuteLogDo) Debug() IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataImportExecuteLogDo) WithContext(ctx context.Context) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataImportExecuteLogDo) ReadDB() IDbdataImportExecuteLogDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataImportExecuteLogDo) WriteDB() IDbdataImportExecuteLogDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataImportExecuteLogDo) Session(config *gorm.Session) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataImportExecuteLogDo) Clauses(conds ...clause.Expression) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataImportExecuteLogDo) Returning(value interface{}, columns ...string) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataImportExecuteLogDo) Not(conds ...gen.Condition) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataImportExecuteLogDo) Or(conds ...gen.Condition) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataImportExecuteLogDo) Select(conds ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataImportExecuteLogDo) Where(conds ...gen.Condition) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataImportExecuteLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataImportExecuteLogDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataImportExecuteLogDo) Order(conds ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataImportExecuteLogDo) Distinct(cols ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataImportExecuteLogDo) Omit(cols ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataImportExecuteLogDo) Join(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataImportExecuteLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataImportExecuteLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataImportExecuteLogDo) Group(cols ...field.Expr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataImportExecuteLogDo) Having(conds ...gen.Condition) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataImportExecuteLogDo) Limit(limit int) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataImportExecuteLogDo) Offset(offset int) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataImportExecuteLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataImportExecuteLogDo) Unscoped() IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataImportExecuteLogDo) Create(values ...*model.DbdataImportExecuteLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataImportExecuteLogDo) CreateInBatches(values []*model.DbdataImportExecuteLog, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataImportExecuteLogDo) Save(values ...*model.DbdataImportExecuteLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataImportExecuteLogDo) First() (*model.DbdataImportExecuteLog, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportExecuteLog), nil
	}
}

func (d dbdataImportExecuteLogDo) Take() (*model.DbdataImportExecuteLog, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportExecuteLog), nil
	}
}

func (d dbdataImportExecuteLogDo) Last() (*model.DbdataImportExecuteLog, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportExecuteLog), nil
	}
}

func (d dbdataImportExecuteLogDo) Find() ([]*model.DbdataImportExecuteLog, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataImportExecuteLog), err
}

func (d dbdataImportExecuteLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataImportExecuteLog, err error) {
	buf := make([]*model.DbdataImportExecuteLog, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataImportExecuteLogDo) FindInBatches(result *[]*model.DbdataImportExecuteLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataImportExecuteLogDo) Attrs(attrs ...field.AssignExpr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataImportExecuteLogDo) Assign(attrs ...field.AssignExpr) IDbdataImportExecuteLogDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataImportExecuteLogDo) Joins(fields ...field.RelationField) IDbdataImportExecuteLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataImportExecuteLogDo) Preload(fields ...field.RelationField) IDbdataImportExecuteLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataImportExecuteLogDo) FirstOrInit() (*model.DbdataImportExecuteLog, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportExecuteLog), nil
	}
}

func (d dbdataImportExecuteLogDo) FirstOrCreate() (*model.DbdataImportExecuteLog, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataImportExecuteLog), nil
	}
}

func (d dbdataImportExecuteLogDo) FindByPage(offset int, limit int) (result []*model.DbdataImportExecuteLog, count int64, err error) {
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

func (d dbdataImportExecuteLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataImportExecuteLogDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataImportExecuteLogDo) Delete(models ...*model.DbdataImportExecuteLog) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataImportExecuteLogDo) withDO(do gen.Dao) *dbdataImportExecuteLogDo {
	d.DO = *do.(*gen.DO)
	return d
}
