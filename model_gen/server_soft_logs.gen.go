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

func newServerSoftLogs(db *gorm.DB, opts ...gen.DOOption) serverSoftLogs {
	_serverSoftLogs := serverSoftLogs{}

	_serverSoftLogs.serverSoftLogsDo.UseDB(db, opts...)
	_serverSoftLogs.serverSoftLogsDo.UseModel(&model.ServerSoftLogs{})

	tableName := _serverSoftLogs.serverSoftLogsDo.TableName()
	_serverSoftLogs.ALL = field.NewAsterisk(tableName)
	_serverSoftLogs.ServerSoftLogsID = field.NewString(tableName, "server_soft_logs_id")
	_serverSoftLogs.ServerSoftLogID = field.NewInt32(tableName, "server_soft_log_id")
	_serverSoftLogs.ServerSoftLogPath = field.NewString(tableName, "server_soft_log_path")
	_serverSoftLogs.Name = field.NewString(tableName, "name")
	_serverSoftLogs.InstallSoftID = field.NewString(tableName, "install_soft_id")

	_serverSoftLogs.fillFieldMap()

	return _serverSoftLogs
}

type serverSoftLogs struct {
	serverSoftLogsDo

	ALL               field.Asterisk
	ServerSoftLogsID  field.String // 日志组ID
	ServerSoftLogID   field.Int32  // 日志ID
	ServerSoftLogPath field.String // 日志完整路径
	Name              field.String
	InstallSoftID     field.String

	fieldMap map[string]field.Expr
}

func (s serverSoftLogs) Table(newTableName string) *serverSoftLogs {
	s.serverSoftLogsDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverSoftLogs) As(alias string) *serverSoftLogs {
	s.serverSoftLogsDo.DO = *(s.serverSoftLogsDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverSoftLogs) updateTableName(table string) *serverSoftLogs {
	s.ALL = field.NewAsterisk(table)
	s.ServerSoftLogsID = field.NewString(table, "server_soft_logs_id")
	s.ServerSoftLogID = field.NewInt32(table, "server_soft_log_id")
	s.ServerSoftLogPath = field.NewString(table, "server_soft_log_path")
	s.Name = field.NewString(table, "name")
	s.InstallSoftID = field.NewString(table, "install_soft_id")

	s.fillFieldMap()

	return s
}

func (s *serverSoftLogs) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverSoftLogs) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["server_soft_logs_id"] = s.ServerSoftLogsID
	s.fieldMap["server_soft_log_id"] = s.ServerSoftLogID
	s.fieldMap["server_soft_log_path"] = s.ServerSoftLogPath
	s.fieldMap["name"] = s.Name
	s.fieldMap["install_soft_id"] = s.InstallSoftID
}

func (s serverSoftLogs) clone(db *gorm.DB) serverSoftLogs {
	s.serverSoftLogsDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverSoftLogs) replaceDB(db *gorm.DB) serverSoftLogs {
	s.serverSoftLogsDo.ReplaceDB(db)
	return s
}

type serverSoftLogsDo struct{ gen.DO }

type IServerSoftLogsDo interface {
	gen.SubQuery
	Debug() IServerSoftLogsDo
	WithContext(ctx context.Context) IServerSoftLogsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerSoftLogsDo
	WriteDB() IServerSoftLogsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerSoftLogsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerSoftLogsDo
	Not(conds ...gen.Condition) IServerSoftLogsDo
	Or(conds ...gen.Condition) IServerSoftLogsDo
	Select(conds ...field.Expr) IServerSoftLogsDo
	Where(conds ...gen.Condition) IServerSoftLogsDo
	Order(conds ...field.Expr) IServerSoftLogsDo
	Distinct(cols ...field.Expr) IServerSoftLogsDo
	Omit(cols ...field.Expr) IServerSoftLogsDo
	Join(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo
	Group(cols ...field.Expr) IServerSoftLogsDo
	Having(conds ...gen.Condition) IServerSoftLogsDo
	Limit(limit int) IServerSoftLogsDo
	Offset(offset int) IServerSoftLogsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerSoftLogsDo
	Unscoped() IServerSoftLogsDo
	Create(values ...*model.ServerSoftLogs) error
	CreateInBatches(values []*model.ServerSoftLogs, batchSize int) error
	Save(values ...*model.ServerSoftLogs) error
	First() (*model.ServerSoftLogs, error)
	Take() (*model.ServerSoftLogs, error)
	Last() (*model.ServerSoftLogs, error)
	Find() ([]*model.ServerSoftLogs, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerSoftLogs, err error)
	FindInBatches(result *[]*model.ServerSoftLogs, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerSoftLogs) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerSoftLogsDo
	Assign(attrs ...field.AssignExpr) IServerSoftLogsDo
	Joins(fields ...field.RelationField) IServerSoftLogsDo
	Preload(fields ...field.RelationField) IServerSoftLogsDo
	FirstOrInit() (*model.ServerSoftLogs, error)
	FirstOrCreate() (*model.ServerSoftLogs, error)
	FindByPage(offset int, limit int) (result []*model.ServerSoftLogs, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerSoftLogsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverSoftLogsDo) Debug() IServerSoftLogsDo {
	return s.withDO(s.DO.Debug())
}

func (s serverSoftLogsDo) WithContext(ctx context.Context) IServerSoftLogsDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverSoftLogsDo) ReadDB() IServerSoftLogsDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverSoftLogsDo) WriteDB() IServerSoftLogsDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverSoftLogsDo) Session(config *gorm.Session) IServerSoftLogsDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverSoftLogsDo) Clauses(conds ...clause.Expression) IServerSoftLogsDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverSoftLogsDo) Returning(value interface{}, columns ...string) IServerSoftLogsDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverSoftLogsDo) Not(conds ...gen.Condition) IServerSoftLogsDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverSoftLogsDo) Or(conds ...gen.Condition) IServerSoftLogsDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverSoftLogsDo) Select(conds ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverSoftLogsDo) Where(conds ...gen.Condition) IServerSoftLogsDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverSoftLogsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerSoftLogsDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverSoftLogsDo) Order(conds ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverSoftLogsDo) Distinct(cols ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverSoftLogsDo) Omit(cols ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverSoftLogsDo) Join(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverSoftLogsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverSoftLogsDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverSoftLogsDo) Group(cols ...field.Expr) IServerSoftLogsDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverSoftLogsDo) Having(conds ...gen.Condition) IServerSoftLogsDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverSoftLogsDo) Limit(limit int) IServerSoftLogsDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverSoftLogsDo) Offset(offset int) IServerSoftLogsDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverSoftLogsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerSoftLogsDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverSoftLogsDo) Unscoped() IServerSoftLogsDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverSoftLogsDo) Create(values ...*model.ServerSoftLogs) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverSoftLogsDo) CreateInBatches(values []*model.ServerSoftLogs, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverSoftLogsDo) Save(values ...*model.ServerSoftLogs) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverSoftLogsDo) First() (*model.ServerSoftLogs, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSoftLogs), nil
	}
}

func (s serverSoftLogsDo) Take() (*model.ServerSoftLogs, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSoftLogs), nil
	}
}

func (s serverSoftLogsDo) Last() (*model.ServerSoftLogs, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSoftLogs), nil
	}
}

func (s serverSoftLogsDo) Find() ([]*model.ServerSoftLogs, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerSoftLogs), err
}

func (s serverSoftLogsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerSoftLogs, err error) {
	buf := make([]*model.ServerSoftLogs, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverSoftLogsDo) FindInBatches(result *[]*model.ServerSoftLogs, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverSoftLogsDo) Attrs(attrs ...field.AssignExpr) IServerSoftLogsDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverSoftLogsDo) Assign(attrs ...field.AssignExpr) IServerSoftLogsDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverSoftLogsDo) Joins(fields ...field.RelationField) IServerSoftLogsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverSoftLogsDo) Preload(fields ...field.RelationField) IServerSoftLogsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverSoftLogsDo) FirstOrInit() (*model.ServerSoftLogs, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSoftLogs), nil
	}
}

func (s serverSoftLogsDo) FirstOrCreate() (*model.ServerSoftLogs, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSoftLogs), nil
	}
}

func (s serverSoftLogsDo) FindByPage(offset int, limit int) (result []*model.ServerSoftLogs, count int64, err error) {
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

func (s serverSoftLogsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverSoftLogsDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverSoftLogsDo) Delete(models ...*model.ServerSoftLogs) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverSoftLogsDo) withDO(do gen.Dao) *serverSoftLogsDo {
	s.DO = *do.(*gen.DO)
	return s
}