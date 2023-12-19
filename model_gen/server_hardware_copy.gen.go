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

func newServerHardwareCopy(db *gorm.DB, opts ...gen.DOOption) serverHardwareCopy {
	_serverHardwareCopy := serverHardwareCopy{}

	_serverHardwareCopy.serverHardwareCopyDo.UseDB(db, opts...)
	_serverHardwareCopy.serverHardwareCopyDo.UseModel(&model.ServerHardwareCopy{})

	tableName := _serverHardwareCopy.serverHardwareCopyDo.TableName()
	_serverHardwareCopy.ALL = field.NewAsterisk(tableName)
	_serverHardwareCopy.ServerHardwareID = field.NewString(tableName, "server_hardware_id")
	_serverHardwareCopy.CPUCores = field.NewInt32(tableName, "cpu_cores")
	_serverHardwareCopy.CPUFrequency = field.NewString(tableName, "cpu_frequency")
	_serverHardwareCopy.CPULogicCount = field.NewInt32(tableName, "cpu_logic_count")
	_serverHardwareCopy.CPUPhysicalCount = field.NewInt32(tableName, "cpu_physical_count")
	_serverHardwareCopy.CPUCache = field.NewInt32(tableName, "cpu_cache")
	_serverHardwareCopy.MemorySize = field.NewInt32(tableName, "memory_size")
	_serverHardwareCopy.DiskTotalSize = field.NewInt32(tableName, "disk_total_size")
	_serverHardwareCopy.DiskCount = field.NewInt32(tableName, "disk_count")
	_serverHardwareCopy.Comments = field.NewString(tableName, "comments")
	_serverHardwareCopy.ModifyTime = field.NewTime(tableName, "modify_time")
	_serverHardwareCopy.CreateTime = field.NewTime(tableName, "create_time")

	_serverHardwareCopy.fillFieldMap()

	return _serverHardwareCopy
}

type serverHardwareCopy struct {
	serverHardwareCopyDo

	ALL              field.Asterisk
	ServerHardwareID field.String // 主键UUID
	CPUCores         field.Int32  // cpu核心查看每颗CPU封装的是几核
	CPUFrequency     field.String // 如：Intel(R) Xeon(R) CPU E7-4830 v4 @ 2.00GHz
	CPULogicCount    field.Int32  // 逻辑核数
	CPUPhysicalCount field.Int32  // 物理核数
	CPUCache         field.Int32  // cpu缓存(单位KB)
	MemorySize       field.Int32  // 内存大小，单位MB
	DiskTotalSize    field.Int32  // 服务器磁盘总容量（单位GB）
	DiskCount        field.Int32  // 磁盘数量
	Comments         field.String // 备注说明
	ModifyTime       field.Time   // 记录修改时间（数据库自动写入）
	CreateTime       field.Time   // 记录创建时间（数据库自动写入）

	fieldMap map[string]field.Expr
}

func (s serverHardwareCopy) Table(newTableName string) *serverHardwareCopy {
	s.serverHardwareCopyDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverHardwareCopy) As(alias string) *serverHardwareCopy {
	s.serverHardwareCopyDo.DO = *(s.serverHardwareCopyDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverHardwareCopy) updateTableName(table string) *serverHardwareCopy {
	s.ALL = field.NewAsterisk(table)
	s.ServerHardwareID = field.NewString(table, "server_hardware_id")
	s.CPUCores = field.NewInt32(table, "cpu_cores")
	s.CPUFrequency = field.NewString(table, "cpu_frequency")
	s.CPULogicCount = field.NewInt32(table, "cpu_logic_count")
	s.CPUPhysicalCount = field.NewInt32(table, "cpu_physical_count")
	s.CPUCache = field.NewInt32(table, "cpu_cache")
	s.MemorySize = field.NewInt32(table, "memory_size")
	s.DiskTotalSize = field.NewInt32(table, "disk_total_size")
	s.DiskCount = field.NewInt32(table, "disk_count")
	s.Comments = field.NewString(table, "comments")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *serverHardwareCopy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverHardwareCopy) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["server_hardware_id"] = s.ServerHardwareID
	s.fieldMap["cpu_cores"] = s.CPUCores
	s.fieldMap["cpu_frequency"] = s.CPUFrequency
	s.fieldMap["cpu_logic_count"] = s.CPULogicCount
	s.fieldMap["cpu_physical_count"] = s.CPUPhysicalCount
	s.fieldMap["cpu_cache"] = s.CPUCache
	s.fieldMap["memory_size"] = s.MemorySize
	s.fieldMap["disk_total_size"] = s.DiskTotalSize
	s.fieldMap["disk_count"] = s.DiskCount
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["create_time"] = s.CreateTime
}

func (s serverHardwareCopy) clone(db *gorm.DB) serverHardwareCopy {
	s.serverHardwareCopyDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverHardwareCopy) replaceDB(db *gorm.DB) serverHardwareCopy {
	s.serverHardwareCopyDo.ReplaceDB(db)
	return s
}

type serverHardwareCopyDo struct{ gen.DO }

type IServerHardwareCopyDo interface {
	gen.SubQuery
	Debug() IServerHardwareCopyDo
	WithContext(ctx context.Context) IServerHardwareCopyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerHardwareCopyDo
	WriteDB() IServerHardwareCopyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerHardwareCopyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerHardwareCopyDo
	Not(conds ...gen.Condition) IServerHardwareCopyDo
	Or(conds ...gen.Condition) IServerHardwareCopyDo
	Select(conds ...field.Expr) IServerHardwareCopyDo
	Where(conds ...gen.Condition) IServerHardwareCopyDo
	Order(conds ...field.Expr) IServerHardwareCopyDo
	Distinct(cols ...field.Expr) IServerHardwareCopyDo
	Omit(cols ...field.Expr) IServerHardwareCopyDo
	Join(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo
	Group(cols ...field.Expr) IServerHardwareCopyDo
	Having(conds ...gen.Condition) IServerHardwareCopyDo
	Limit(limit int) IServerHardwareCopyDo
	Offset(offset int) IServerHardwareCopyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerHardwareCopyDo
	Unscoped() IServerHardwareCopyDo
	Create(values ...*model.ServerHardwareCopy) error
	CreateInBatches(values []*model.ServerHardwareCopy, batchSize int) error
	Save(values ...*model.ServerHardwareCopy) error
	First() (*model.ServerHardwareCopy, error)
	Take() (*model.ServerHardwareCopy, error)
	Last() (*model.ServerHardwareCopy, error)
	Find() ([]*model.ServerHardwareCopy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerHardwareCopy, err error)
	FindInBatches(result *[]*model.ServerHardwareCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerHardwareCopy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerHardwareCopyDo
	Assign(attrs ...field.AssignExpr) IServerHardwareCopyDo
	Joins(fields ...field.RelationField) IServerHardwareCopyDo
	Preload(fields ...field.RelationField) IServerHardwareCopyDo
	FirstOrInit() (*model.ServerHardwareCopy, error)
	FirstOrCreate() (*model.ServerHardwareCopy, error)
	FindByPage(offset int, limit int) (result []*model.ServerHardwareCopy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerHardwareCopyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverHardwareCopyDo) Debug() IServerHardwareCopyDo {
	return s.withDO(s.DO.Debug())
}

func (s serverHardwareCopyDo) WithContext(ctx context.Context) IServerHardwareCopyDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverHardwareCopyDo) ReadDB() IServerHardwareCopyDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverHardwareCopyDo) WriteDB() IServerHardwareCopyDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverHardwareCopyDo) Session(config *gorm.Session) IServerHardwareCopyDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverHardwareCopyDo) Clauses(conds ...clause.Expression) IServerHardwareCopyDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverHardwareCopyDo) Returning(value interface{}, columns ...string) IServerHardwareCopyDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverHardwareCopyDo) Not(conds ...gen.Condition) IServerHardwareCopyDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverHardwareCopyDo) Or(conds ...gen.Condition) IServerHardwareCopyDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverHardwareCopyDo) Select(conds ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverHardwareCopyDo) Where(conds ...gen.Condition) IServerHardwareCopyDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverHardwareCopyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerHardwareCopyDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverHardwareCopyDo) Order(conds ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverHardwareCopyDo) Distinct(cols ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverHardwareCopyDo) Omit(cols ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverHardwareCopyDo) Join(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverHardwareCopyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverHardwareCopyDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverHardwareCopyDo) Group(cols ...field.Expr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverHardwareCopyDo) Having(conds ...gen.Condition) IServerHardwareCopyDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverHardwareCopyDo) Limit(limit int) IServerHardwareCopyDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverHardwareCopyDo) Offset(offset int) IServerHardwareCopyDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverHardwareCopyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerHardwareCopyDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverHardwareCopyDo) Unscoped() IServerHardwareCopyDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverHardwareCopyDo) Create(values ...*model.ServerHardwareCopy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverHardwareCopyDo) CreateInBatches(values []*model.ServerHardwareCopy, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverHardwareCopyDo) Save(values ...*model.ServerHardwareCopy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverHardwareCopyDo) First() (*model.ServerHardwareCopy, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerHardwareCopy), nil
	}
}

func (s serverHardwareCopyDo) Take() (*model.ServerHardwareCopy, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerHardwareCopy), nil
	}
}

func (s serverHardwareCopyDo) Last() (*model.ServerHardwareCopy, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerHardwareCopy), nil
	}
}

func (s serverHardwareCopyDo) Find() ([]*model.ServerHardwareCopy, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerHardwareCopy), err
}

func (s serverHardwareCopyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerHardwareCopy, err error) {
	buf := make([]*model.ServerHardwareCopy, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverHardwareCopyDo) FindInBatches(result *[]*model.ServerHardwareCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverHardwareCopyDo) Attrs(attrs ...field.AssignExpr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverHardwareCopyDo) Assign(attrs ...field.AssignExpr) IServerHardwareCopyDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverHardwareCopyDo) Joins(fields ...field.RelationField) IServerHardwareCopyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverHardwareCopyDo) Preload(fields ...field.RelationField) IServerHardwareCopyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverHardwareCopyDo) FirstOrInit() (*model.ServerHardwareCopy, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerHardwareCopy), nil
	}
}

func (s serverHardwareCopyDo) FirstOrCreate() (*model.ServerHardwareCopy, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerHardwareCopy), nil
	}
}

func (s serverHardwareCopyDo) FindByPage(offset int, limit int) (result []*model.ServerHardwareCopy, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s serverHardwareCopyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverHardwareCopyDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverHardwareCopyDo) Delete(models ...*model.ServerHardwareCopy) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverHardwareCopyDo) withDO(do gen.Dao) *serverHardwareCopyDo {
	s.DO = *do.(*gen.DO)
	return s
}
