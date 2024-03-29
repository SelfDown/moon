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

func newDbdataExport(db *gorm.DB, opts ...gen.DOOption) dbdataExport {
	_dbdataExport := dbdataExport{}

	_dbdataExport.dbdataExportDo.UseDB(db, opts...)
	_dbdataExport.dbdataExportDo.UseModel(&model.DbdataExport{})

	tableName := _dbdataExport.dbdataExportDo.TableName()
	_dbdataExport.ALL = field.NewAsterisk(tableName)
	_dbdataExport.DbdataExportID = field.NewString(tableName, "dbdata_export_id")
	_dbdataExport.Title = field.NewString(tableName, "title")
	_dbdataExport.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_dbdataExport.SoftUserID = field.NewString(tableName, "soft_user_id")
	_dbdataExport.AddTime = field.NewString(tableName, "add_time")
	_dbdataExport.OpUser = field.NewString(tableName, "op_user")
	_dbdataExport.Note = field.NewString(tableName, "note")
	_dbdataExport.UploadFileName = field.NewString(tableName, "upload_file_name")
	_dbdataExport.UploadFilePath = field.NewString(tableName, "upload_file_path")
	_dbdataExport.ExportWay = field.NewString(tableName, "export_way")
	_dbdataExport.UserMetadata = field.NewString(tableName, "user_metadata")
	_dbdataExport.DirectoryPath = field.NewString(tableName, "directory_path")
	_dbdataExport.DirectoryName = field.NewString(tableName, "directory_name")
	_dbdataExport.OpPercent = field.NewString(tableName, "op_percent")
	_dbdataExport.CanEditor = field.NewString(tableName, "can_editor")
	_dbdataExport.StransactionIdx = field.NewString(tableName, "stransaction_idx")
	_dbdataExport.ExpOraVersion = field.NewString(tableName, "exp_ora_version")
	_dbdataExport.AllMetaDumpFilename = field.NewString(tableName, "all_meta_dump_filename")
	_dbdataExport.Sid = field.NewString(tableName, "sid")
	_dbdataExport.AllMetaLogFilename = field.NewString(tableName, "all_meta_log_filename")
	_dbdataExport.IsDelete = field.NewString(tableName, "is_delete")

	_dbdataExport.fillFieldMap()

	return _dbdataExport
}

type dbdataExport struct {
	dbdataExportDo

	ALL                 field.Asterisk
	DbdataExportID      field.String
	Title               field.String
	InstallSoftID       field.String
	SoftUserID          field.String
	AddTime             field.String
	OpUser              field.String
	Note                field.String
	UploadFileName      field.String
	UploadFilePath      field.String
	ExportWay           field.String // 导出方式
	UserMetadata        field.String // 用户表空间导出方式
	DirectoryPath       field.String
	DirectoryName       field.String
	OpPercent           field.String
	CanEditor           field.String
	StransactionIdx     field.String
	ExpOraVersion       field.String
	AllMetaDumpFilename field.String
	Sid                 field.String
	AllMetaLogFilename  field.String
	IsDelete            field.String

	fieldMap map[string]field.Expr
}

func (d dbdataExport) Table(newTableName string) *dbdataExport {
	d.dbdataExportDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dbdataExport) As(alias string) *dbdataExport {
	d.dbdataExportDo.DO = *(d.dbdataExportDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dbdataExport) updateTableName(table string) *dbdataExport {
	d.ALL = field.NewAsterisk(table)
	d.DbdataExportID = field.NewString(table, "dbdata_export_id")
	d.Title = field.NewString(table, "title")
	d.InstallSoftID = field.NewString(table, "install_soft_id")
	d.SoftUserID = field.NewString(table, "soft_user_id")
	d.AddTime = field.NewString(table, "add_time")
	d.OpUser = field.NewString(table, "op_user")
	d.Note = field.NewString(table, "note")
	d.UploadFileName = field.NewString(table, "upload_file_name")
	d.UploadFilePath = field.NewString(table, "upload_file_path")
	d.ExportWay = field.NewString(table, "export_way")
	d.UserMetadata = field.NewString(table, "user_metadata")
	d.DirectoryPath = field.NewString(table, "directory_path")
	d.DirectoryName = field.NewString(table, "directory_name")
	d.OpPercent = field.NewString(table, "op_percent")
	d.CanEditor = field.NewString(table, "can_editor")
	d.StransactionIdx = field.NewString(table, "stransaction_idx")
	d.ExpOraVersion = field.NewString(table, "exp_ora_version")
	d.AllMetaDumpFilename = field.NewString(table, "all_meta_dump_filename")
	d.Sid = field.NewString(table, "sid")
	d.AllMetaLogFilename = field.NewString(table, "all_meta_log_filename")
	d.IsDelete = field.NewString(table, "is_delete")

	d.fillFieldMap()

	return d
}

func (d *dbdataExport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dbdataExport) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 21)
	d.fieldMap["dbdata_export_id"] = d.DbdataExportID
	d.fieldMap["title"] = d.Title
	d.fieldMap["install_soft_id"] = d.InstallSoftID
	d.fieldMap["soft_user_id"] = d.SoftUserID
	d.fieldMap["add_time"] = d.AddTime
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["note"] = d.Note
	d.fieldMap["upload_file_name"] = d.UploadFileName
	d.fieldMap["upload_file_path"] = d.UploadFilePath
	d.fieldMap["export_way"] = d.ExportWay
	d.fieldMap["user_metadata"] = d.UserMetadata
	d.fieldMap["directory_path"] = d.DirectoryPath
	d.fieldMap["directory_name"] = d.DirectoryName
	d.fieldMap["op_percent"] = d.OpPercent
	d.fieldMap["can_editor"] = d.CanEditor
	d.fieldMap["stransaction_idx"] = d.StransactionIdx
	d.fieldMap["exp_ora_version"] = d.ExpOraVersion
	d.fieldMap["all_meta_dump_filename"] = d.AllMetaDumpFilename
	d.fieldMap["sid"] = d.Sid
	d.fieldMap["all_meta_log_filename"] = d.AllMetaLogFilename
	d.fieldMap["is_delete"] = d.IsDelete
}

func (d dbdataExport) clone(db *gorm.DB) dbdataExport {
	d.dbdataExportDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dbdataExport) replaceDB(db *gorm.DB) dbdataExport {
	d.dbdataExportDo.ReplaceDB(db)
	return d
}

type dbdataExportDo struct{ gen.DO }

type IDbdataExportDo interface {
	gen.SubQuery
	Debug() IDbdataExportDo
	WithContext(ctx context.Context) IDbdataExportDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDbdataExportDo
	WriteDB() IDbdataExportDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDbdataExportDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDbdataExportDo
	Not(conds ...gen.Condition) IDbdataExportDo
	Or(conds ...gen.Condition) IDbdataExportDo
	Select(conds ...field.Expr) IDbdataExportDo
	Where(conds ...gen.Condition) IDbdataExportDo
	Order(conds ...field.Expr) IDbdataExportDo
	Distinct(cols ...field.Expr) IDbdataExportDo
	Omit(cols ...field.Expr) IDbdataExportDo
	Join(table schema.Tabler, on ...field.Expr) IDbdataExportDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportDo
	Group(cols ...field.Expr) IDbdataExportDo
	Having(conds ...gen.Condition) IDbdataExportDo
	Limit(limit int) IDbdataExportDo
	Offset(offset int) IDbdataExportDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportDo
	Unscoped() IDbdataExportDo
	Create(values ...*model.DbdataExport) error
	CreateInBatches(values []*model.DbdataExport, batchSize int) error
	Save(values ...*model.DbdataExport) error
	First() (*model.DbdataExport, error)
	Take() (*model.DbdataExport, error)
	Last() (*model.DbdataExport, error)
	Find() ([]*model.DbdataExport, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExport, err error)
	FindInBatches(result *[]*model.DbdataExport, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DbdataExport) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDbdataExportDo
	Assign(attrs ...field.AssignExpr) IDbdataExportDo
	Joins(fields ...field.RelationField) IDbdataExportDo
	Preload(fields ...field.RelationField) IDbdataExportDo
	FirstOrInit() (*model.DbdataExport, error)
	FirstOrCreate() (*model.DbdataExport, error)
	FindByPage(offset int, limit int) (result []*model.DbdataExport, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDbdataExportDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dbdataExportDo) Debug() IDbdataExportDo {
	return d.withDO(d.DO.Debug())
}

func (d dbdataExportDo) WithContext(ctx context.Context) IDbdataExportDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dbdataExportDo) ReadDB() IDbdataExportDo {
	return d.Clauses(dbresolver.Read)
}

func (d dbdataExportDo) WriteDB() IDbdataExportDo {
	return d.Clauses(dbresolver.Write)
}

func (d dbdataExportDo) Session(config *gorm.Session) IDbdataExportDo {
	return d.withDO(d.DO.Session(config))
}

func (d dbdataExportDo) Clauses(conds ...clause.Expression) IDbdataExportDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dbdataExportDo) Returning(value interface{}, columns ...string) IDbdataExportDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dbdataExportDo) Not(conds ...gen.Condition) IDbdataExportDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dbdataExportDo) Or(conds ...gen.Condition) IDbdataExportDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dbdataExportDo) Select(conds ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dbdataExportDo) Where(conds ...gen.Condition) IDbdataExportDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dbdataExportDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDbdataExportDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dbdataExportDo) Order(conds ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dbdataExportDo) Distinct(cols ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dbdataExportDo) Omit(cols ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dbdataExportDo) Join(table schema.Tabler, on ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dbdataExportDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dbdataExportDo) RightJoin(table schema.Tabler, on ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dbdataExportDo) Group(cols ...field.Expr) IDbdataExportDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dbdataExportDo) Having(conds ...gen.Condition) IDbdataExportDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dbdataExportDo) Limit(limit int) IDbdataExportDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dbdataExportDo) Offset(offset int) IDbdataExportDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dbdataExportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDbdataExportDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dbdataExportDo) Unscoped() IDbdataExportDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dbdataExportDo) Create(values ...*model.DbdataExport) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dbdataExportDo) CreateInBatches(values []*model.DbdataExport, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dbdataExportDo) Save(values ...*model.DbdataExport) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dbdataExportDo) First() (*model.DbdataExport, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExport), nil
	}
}

func (d dbdataExportDo) Take() (*model.DbdataExport, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExport), nil
	}
}

func (d dbdataExportDo) Last() (*model.DbdataExport, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExport), nil
	}
}

func (d dbdataExportDo) Find() ([]*model.DbdataExport, error) {
	result, err := d.DO.Find()
	return result.([]*model.DbdataExport), err
}

func (d dbdataExportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DbdataExport, err error) {
	buf := make([]*model.DbdataExport, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dbdataExportDo) FindInBatches(result *[]*model.DbdataExport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dbdataExportDo) Attrs(attrs ...field.AssignExpr) IDbdataExportDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dbdataExportDo) Assign(attrs ...field.AssignExpr) IDbdataExportDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dbdataExportDo) Joins(fields ...field.RelationField) IDbdataExportDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dbdataExportDo) Preload(fields ...field.RelationField) IDbdataExportDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dbdataExportDo) FirstOrInit() (*model.DbdataExport, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExport), nil
	}
}

func (d dbdataExportDo) FirstOrCreate() (*model.DbdataExport, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DbdataExport), nil
	}
}

func (d dbdataExportDo) FindByPage(offset int, limit int) (result []*model.DbdataExport, count int64, err error) {
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

func (d dbdataExportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dbdataExportDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dbdataExportDo) Delete(models ...*model.DbdataExport) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dbdataExportDo) withDO(do gen.Dao) *dbdataExportDo {
	d.DO = *do.(*gen.DO)
	return d
}
