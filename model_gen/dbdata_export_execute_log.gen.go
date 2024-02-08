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

func newDbdataExportExecuteLog(db *gorm.DB, opts ...gen.DOOption) dbdataExportExecuteLog {
	_dbdataExportExecuteLog := dbdataExportExecuteLog{}

	_dbdataExportExecuteLog.dbdataExportExecuteLogDo.UseDB(db, opts...)
	_dbdataExportExecuteLog.dbdataExportExecuteLogDo.UseModel(&model.DbdataExportExecuteLog{})

	tableName := _dbdataExportExecuteLog.dbdataExportExecuteLogDo.TableName()
	_dbdataExportExecuteLog.ALL = field.NewAsterisk(tableName)
	_dbdataExportExecuteLog.DbdataExportExecuteLogID = field.NewString(tableName, "dbdata_export_execute_log_id")
	_dbdataExportExecuteLog.DbdataExportID = field.NewString(tableName, "dbdata_export_id")
	_dbdataExportExecuteLog.OpUser = field.NewString(tableName, "op_user")
	_dbdataExportExecuteLog.SyncStartTime = field.NewString(tableName, "sync_start_time")
	_dbdataExportExecuteLog.SyncEndTime = field.NewString(tableName, "sync_end_time")
	_dbdataExportExecuteLog.AddTime = field.NewString(tableName, "add_time")
	_dbdataExportExecuteLog.SyncStatu = field.NewString(tableName, "sync_statu")
	_dbdataExportExecuteLog.SyncScript = field.NewString(tableName, "sync_script")
	_dbdataExportExecuteLog.SyncLog = field.NewString(tableName, "sync_log")
	_dbdataExportExecuteLog.Size = field.NewInt64(tableName, "size")
	_dbdataExportExecuteLog.Msg = field.NewString(tableName, "msg")
	_dbdataExportExecuteLog.RemoteShellPath = field.NewString(tableName, "remote_shell_path")
	_dbdataExportExecuteLog.LogPath = field.NewString(tableName, "log_path")
	_dbdataExportExecuteLog.TarFilename = field.NewString(tableName, "tar_filename")
	_dbdataExportExecuteLog.TarDir = field.NewString(tableName, "tar_dir")
	_dbdataExportExecuteLog.RemoteDumpTarName = field.NewString(tableName, "remote_dump_tar_name")
	_dbdataExportExecuteLog.RemoteDumpTarDir = field.NewString(tableName, "remote_dump_tar_dir")
	_dbdataExportExecuteLog.LocalDumpPath = field.NewString(tableName, "local_dump_path")
	_dbdataExportExecuteLog.LocalDumpPathAbs = field.NewString(tableName, "local_dump_path_abs")
	_dbdataExportExecuteLog.SftpDumpPath = field.NewString(tableName, "sftp_dump_path")

	_dbdataExportExecuteLog.fillFieldMap()

	return _dbdataExportExecuteLog
}

type dbdataExportExecuteLog struct {
	dbdataExportExecuteLogDo

	ALL                      field.Asterisk
	DbdataExportExecuteLogID field.String
	DbdataExportID           field.String
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

	fieldMap map[string]field.Expr
}

func (d dbdataExportExecuteLog) Table(newTableName string) *dbdataExportExecuteLog {
	d.dbdataExportExecuteLogDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataExportExecuteLog) As(alias string) *dbdataExportExecuteLog {
	d.dbdataExportExecuteLogDo.DO = *(d.dbdataExportExecuteLogDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataExportExecuteLog) updateTableName(table string) *dbdataExportExecuteLog {
	d.ALL = field.NewAsterisk(table)
	d.DbdataExportExecuteLogID = field.NewString(table, "dbdata_export_execute_log_id")
	d.DbdataExportID = field.NewString(table, "dbdata_export_id")
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

	d.fillFieldMap()

	return d
}

func (d *dbdataExportExecuteLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataExportExecuteLog) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 20)
	d.fieldMap["dbdata_export_execute_log_id"] = d.DbdataExportExecuteLogID
	d.fieldMap["dbdata_export_id"] = d.DbdataExportID
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
}

func (d dbdataExportExecuteLog) clone(db *gorm.DB) dbdataExportExecuteLog {
	d.dbdataExportExecuteLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataExportExecuteLog) replaceDB(db *gorm.DB) dbdataExportExecuteLog {
	d.dbdataExportExecuteLogDo.ReplaceDB(db)
	return d
}

type dbdataExportExecuteLogDo struct{ gen.DO }

type IDbdataExportExecuteLogDo interface {
	gen.SubQuery
	Debug() IDbdataExportExecuteLogDo
	WithContext(ctx context.Context) IDbdataExportExecuteLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataExportExecuteLogDo
	WriteDB() IDbdataExportExecuteLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataExportExecuteLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataExportExecuteLogDo
	Not(conds ...gen.Condition) IDbdataExportExecuteLogDo
	Or(conds ...gen.Condition) IDbdataExportExecuteLogDo
	Select(conds ...field.Expr) IDbdataExportExecuteLogDo
	Where(conds ...gen.Condition) IDbdataExportExecuteLogDo
	Order(conds ...field.Expr) IDbdataExportExecuteLogDo
	Distinct(cols ...field.Expr) IDbdataExportExecuteLogDo
	Omit(cols ...field.Expr) IDbdataExportExecuteLogDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo
	Group(cols ...field.Expr) IDbdataExportExecuteLogDo
	Having(conds ...gen.Condition) IDbdataExportExecuteLogDo
	Limit(limit int) IDbdataExportExecuteLogDo
	Offset(offset int) IDbdataExportExecuteLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportExecuteLogDo
	Unscoped() IDbdataExportExecuteLogDo
	Create(values ...*model.DbdataExportExecuteLog) error
	CreateInBatches(values []*model.DbdataExportExecuteLog, batchSize int) error
	Save(values ...*model.DbdataExportExecuteLog) error
	First() (*model.DbdataExportExecuteLog, error)
	Take() (*model.DbdataExportExecuteLog, error)
	Last() (*model.DbdataExportExecuteLog, error)
	Find() ([]*model.DbdataExportExecuteLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExportExecuteLog, err error)
	FindInBatches(result *[]*model.DbdataExportExecuteLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataExportExecuteLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataExportExecuteLogDo
	Assign(attrs ...field.AssignExpr) IDbdataExportExecuteLogDo
	Joins(fields ...field.RelationField) IDbdataExportExecuteLogDo
	Preload(fields ...field.RelationField) IDbdataExportExecuteLogDo
	FirstOrInit() (*model.DbdataExportExecuteLog, error)
	FirstOrCreate() (*model.DbdataExportExecuteLog, error)
	FindByPage(offset int, limit int) (result []*model.DbdataExportExecuteLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataExportExecuteLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataExportExecuteLogDo) Debug() IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataExportExecuteLogDo) WithContext(ctx context.Context) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataExportExecuteLogDo) ReadDB() IDbdataExportExecuteLogDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataExportExecuteLogDo) WriteDB() IDbdataExportExecuteLogDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataExportExecuteLogDo) Session(config *gorm.Session) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataExportExecuteLogDo) Clauses(conds ...clause.Expression) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataExportExecuteLogDo) Returning(value interface{}, columns ...string) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataExportExecuteLogDo) Not(conds ...gen.Condition) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataExportExecuteLogDo) Or(conds ...gen.Condition) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataExportExecuteLogDo) Select(conds ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataExportExecuteLogDo) Where(conds ...gen.Condition) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataExportExecuteLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataExportExecuteLogDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataExportExecuteLogDo) Order(conds ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataExportExecuteLogDo) Distinct(cols ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataExportExecuteLogDo) Omit(cols ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataExportExecuteLogDo) Join(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataExportExecuteLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataExportExecuteLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataExportExecuteLogDo) Group(cols ...field.Expr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataExportExecuteLogDo) Having(conds ...gen.Condition) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataExportExecuteLogDo) Limit(limit int) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataExportExecuteLogDo) Offset(offset int) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataExportExecuteLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataExportExecuteLogDo) Unscoped() IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataExportExecuteLogDo) Create(values ...*model.DbdataExportExecuteLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataExportExecuteLogDo) CreateInBatches(values []*model.DbdataExportExecuteLog, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataExportExecuteLogDo) Save(values ...*model.DbdataExportExecuteLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataExportExecuteLogDo) First() (*model.DbdataExportExecuteLog, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportExecuteLog), nil
	}
}

func (d dbdataExportExecuteLogDo) Take() (*model.DbdataExportExecuteLog, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportExecuteLog), nil
	}
}

func (d dbdataExportExecuteLogDo) Last() (*model.DbdataExportExecuteLog, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportExecuteLog), nil
	}
}

func (d dbdataExportExecuteLogDo) Find() ([]*model.DbdataExportExecuteLog, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataExportExecuteLog), err
}

func (d dbdataExportExecuteLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExportExecuteLog, err error) {
	buf := make([]*model.DbdataExportExecuteLog, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataExportExecuteLogDo) FindInBatches(result *[]*model.DbdataExportExecuteLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataExportExecuteLogDo) Attrs(attrs ...field.AssignExpr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataExportExecuteLogDo) Assign(attrs ...field.AssignExpr) IDbdataExportExecuteLogDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataExportExecuteLogDo) Joins(fields ...field.RelationField) IDbdataExportExecuteLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataExportExecuteLogDo) Preload(fields ...field.RelationField) IDbdataExportExecuteLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataExportExecuteLogDo) FirstOrInit() (*model.DbdataExportExecuteLog, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportExecuteLog), nil
	}
}

func (d dbdataExportExecuteLogDo) FirstOrCreate() (*model.DbdataExportExecuteLog, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExportExecuteLog), nil
	}
}

func (d dbdataExportExecuteLogDo) FindByPage(offset int, limit int) (result []*model.DbdataExportExecuteLog, count int64, err error) {
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

func (d dbdataExportExecuteLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataExportExecuteLogDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataExportExecuteLogDo) Delete(models ...*model.DbdataExportExecuteLog) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataExportExecuteLogDo) withDO(do gen.Dao) *dbdataExportExecuteLogDo {
	d.DO = *do.(*gen.DO)
	return d
}
