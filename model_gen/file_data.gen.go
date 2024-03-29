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

func newFileData(db *gorm.DB, opts ...gen.DOOption) fileData {
	_fileData := fileData{}

	_fileData.fileDataDo.UseDB(db, opts...)
	_fileData.fileDataDo.UseModel(&model.FileData{})

	tableName := _fileData.fileDataDo.TableName()
	_fileData.ALL = field.NewAsterisk(tableName)
	_fileData.FileID = field.NewString(tableName, "file_id")
	_fileData.TargetID = field.NewString(tableName, "target_id")
	_fileData.FileServiceType = field.NewString(tableName, "file_service_type")
	_fileData.Size = field.NewInt64(tableName, "size")
	_fileData.Data = field.NewBytes(tableName, "data")
	_fileData.Name = field.NewString(tableName, "name")
	_fileData.CreateTime = field.NewTime(tableName, "create_time")
	_fileData.ModifyTime = field.NewTime(tableName, "modify_time")
	_fileData.OpUser = field.NewString(tableName, "op_user")
	_fileData.FileType = field.NewString(tableName, "file_type")

	_fileData.fillFieldMap()

	return _fileData
}

type fileData struct {
	fileDataDo

	ALL             field.Asterisk
	FileID          field.String
	TargetID        field.String
	FileServiceType field.String
	Size            field.Int64
	Data            field.Bytes
	Name            field.String
	CreateTime      field.Time
	ModifyTime      field.Time
	OpUser          field.String
	FileType        field.String

	fieldMap map[string]field.Expr
}

func (f fileData) Table(newTableName string) *fileData {
	f.fileDataDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fileData) As(alias string) *fileData {
	f.fileDataDo.DO = *(f.fileDataDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fileData) updateTableName(table string) *fileData {
	f.ALL = field.NewAsterisk(table)
	f.FileID = field.NewString(table, "file_id")
	f.TargetID = field.NewString(table, "target_id")
	f.FileServiceType = field.NewString(table, "file_service_type")
	f.Size = field.NewInt64(table, "size")
	f.Data = field.NewBytes(table, "data")
	f.Name = field.NewString(table, "name")
	f.CreateTime = field.NewTime(table, "create_time")
	f.ModifyTime = field.NewTime(table, "modify_time")
	f.OpUser = field.NewString(table, "op_user")
	f.FileType = field.NewString(table, "file_type")

	f.fillFieldMap()

	return f
}

func (f *fileData) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fileData) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 10)
	f.fieldMap["file_id"] = f.FileID
	f.fieldMap["target_id"] = f.TargetID
	f.fieldMap["file_service_type"] = f.FileServiceType
	f.fieldMap["size"] = f.Size
	f.fieldMap["data"] = f.Data
	f.fieldMap["name"] = f.Name
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["modify_time"] = f.ModifyTime
	f.fieldMap["op_user"] = f.OpUser
	f.fieldMap["file_type"] = f.FileType
}

func (f fileData) clone(db *gorm.DB) fileData {
	f.fileDataDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fileData) replaceDB(db *gorm.DB) fileData {
	f.fileDataDo.ReplaceDB(db)
	return f
}

type fileDataDo struct{ gen.DO }

type IFileDataDo interface {
	gen.SubQuery
	Debug() IFileDataDo
	WithContext(ctx context.Context) IFileDataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFileDataDo
	WriteDB() IFileDataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFileDataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFileDataDo
	Not(conds ...gen.Condition) IFileDataDo
	Or(conds ...gen.Condition) IFileDataDo
	Select(conds ...field.Expr) IFileDataDo
	Where(conds ...gen.Condition) IFileDataDo
	Order(conds ...field.Expr) IFileDataDo
	Distinct(cols ...field.Expr) IFileDataDo
	Omit(cols ...field.Expr) IFileDataDo
	Join(table schema.Tabler, on ...field.Expr) IFileDataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFileDataDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFileDataDo
	Group(cols ...field.Expr) IFileDataDo
	Having(conds ...gen.Condition) IFileDataDo
	Limit(limit int) IFileDataDo
	Offset(offset int) IFileDataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFileDataDo
	Unscoped() IFileDataDo
	Create(values ...*model.FileData) error
	CreateInBatches(values []*model.FileData, batchSize int) error
	Save(values ...*model.FileData) error
	First() (*model.FileData, error)
	Take() (*model.FileData, error)
	Last() (*model.FileData, error)
	Find() ([]*model.FileData, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileData, err error)
	FindInBatches(result *[]*model.FileData, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FileData) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFileDataDo
	Assign(attrs ...field.AssignExpr) IFileDataDo
	Joins(fields ...field.RelationField) IFileDataDo
	Preload(fields ...field.RelationField) IFileDataDo
	FirstOrInit() (*model.FileData, error)
	FirstOrCreate() (*model.FileData, error)
	FindByPage(offset int, limit int) (result []*model.FileData, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFileDataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fileDataDo) Debug() IFileDataDo {
	return f.withDO(f.DO.Debug())
}

func (f fileDataDo) WithContext(ctx context.Context) IFileDataDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileDataDo) ReadDB() IFileDataDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileDataDo) WriteDB() IFileDataDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileDataDo) Session(config *gorm.Session) IFileDataDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileDataDo) Clauses(conds ...clause.Expression) IFileDataDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileDataDo) Returning(value interface{}, columns ...string) IFileDataDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileDataDo) Not(conds ...gen.Condition) IFileDataDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileDataDo) Or(conds ...gen.Condition) IFileDataDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileDataDo) Select(conds ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileDataDo) Where(conds ...gen.Condition) IFileDataDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileDataDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFileDataDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f fileDataDo) Order(conds ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileDataDo) Distinct(cols ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileDataDo) Omit(cols ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileDataDo) Join(table schema.Tabler, on ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileDataDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileDataDo) RightJoin(table schema.Tabler, on ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileDataDo) Group(cols ...field.Expr) IFileDataDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileDataDo) Having(conds ...gen.Condition) IFileDataDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileDataDo) Limit(limit int) IFileDataDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileDataDo) Offset(offset int) IFileDataDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileDataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFileDataDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileDataDo) Unscoped() IFileDataDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileDataDo) Create(values ...*model.FileData) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileDataDo) CreateInBatches(values []*model.FileData, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileDataDo) Save(values ...*model.FileData) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileDataDo) First() (*model.FileData, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileData), nil
	}
}

func (f fileDataDo) Take() (*model.FileData, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileData), nil
	}
}

func (f fileDataDo) Last() (*model.FileData, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileData), nil
	}
}

func (f fileDataDo) Find() ([]*model.FileData, error) {
	result, err := f.DO.Find()
	return result.([]*model.FileData), err
}

func (f fileDataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileData, err error) {
	buf := make([]*model.FileData, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileDataDo) FindInBatches(result *[]*model.FileData, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileDataDo) Attrs(attrs ...field.AssignExpr) IFileDataDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileDataDo) Assign(attrs ...field.AssignExpr) IFileDataDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileDataDo) Joins(fields ...field.RelationField) IFileDataDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileDataDo) Preload(fields ...field.RelationField) IFileDataDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileDataDo) FirstOrInit() (*model.FileData, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileData), nil
	}
}

func (f fileDataDo) FirstOrCreate() (*model.FileData, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileData), nil
	}
}

func (f fileDataDo) FindByPage(offset int, limit int) (result []*model.FileData, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileDataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileDataDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileDataDo) Delete(models ...*model.FileData) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileDataDo) withDO(do gen.Dao) *fileDataDo {
	f.DO = *do.(*gen.DO)
	return f
}
