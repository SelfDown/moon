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

func newServerSSHEvent(db *gorm.DB, opts ...gen.DOOption) serverSSHEvent {
	_serverSSHEvent := serverSSHEvent{}

	_serverSSHEvent.serverSSHEventDo.UseDB(db, opts...)
	_serverSSHEvent.serverSSHEventDo.UseModel(&model.ServerSSHEvent{})

	tableName := _serverSSHEvent.serverSSHEventDo.TableName()
	_serverSSHEvent.ALL = field.NewAsterisk(tableName)
	_serverSSHEvent.ServerSSHEventID = field.NewString(tableName, "server_ssh_event_id")
	_serverSSHEvent.IP = field.NewString(tableName, "ip")
	_serverSSHEvent.SSHPort = field.NewString(tableName, "ssh_port")
	_serverSSHEvent.SSHUser = field.NewString(tableName, "ssh_user")
	_serverSSHEvent.Pwd = field.NewString(tableName, "pwd")
	_serverSSHEvent.Cmd = field.NewString(tableName, "cmd")
	_serverSSHEvent.Ret = field.NewString(tableName, "ret")
	_serverSSHEvent.CreateTime = field.NewTime(tableName, "create_time")
	_serverSSHEvent.Note = field.NewString(tableName, "note")
	_serverSSHEvent.TransactionID = field.NewString(tableName, "transaction_id")
	_serverSSHEvent.ModifyTime = field.NewTime(tableName, "modify_time")
	_serverSSHEvent.Comments = field.NewString(tableName, "comments")

	_serverSSHEvent.fillFieldMap()

	return _serverSSHEvent
}

type serverSSHEvent struct {
	serverSSHEventDo

	ALL              field.Asterisk
	ServerSSHEventID field.String
	IP               field.String
	SSHPort          field.String
	SSHUser          field.String
	Pwd              field.String
	Cmd              field.String
	Ret              field.String
	CreateTime       field.Time // 记录创建时间（数据库自动写入）
	Note             field.String
	TransactionID    field.String
	ModifyTime       field.Time   // 记录修改时间（数据库自动写入）
	Comments         field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s serverSSHEvent) Table(newTableName string) *serverSSHEvent {
	s.serverSSHEventDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverSSHEvent) As(alias string) *serverSSHEvent {
	s.serverSSHEventDo.DO = *(s.serverSSHEventDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverSSHEvent) updateTableName(table string) *serverSSHEvent {
	s.ALL = field.NewAsterisk(table)
	s.ServerSSHEventID = field.NewString(table, "server_ssh_event_id")
	s.IP = field.NewString(table, "ip")
	s.SSHPort = field.NewString(table, "ssh_port")
	s.SSHUser = field.NewString(table, "ssh_user")
	s.Pwd = field.NewString(table, "pwd")
	s.Cmd = field.NewString(table, "cmd")
	s.Ret = field.NewString(table, "ret")
	s.CreateTime = field.NewTime(table, "create_time")
	s.Note = field.NewString(table, "note")
	s.TransactionID = field.NewString(table, "transaction_id")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *serverSSHEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverSSHEvent) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["server_ssh_event_id"] = s.ServerSSHEventID
	s.fieldMap["ip"] = s.IP
	s.fieldMap["ssh_port"] = s.SSHPort
	s.fieldMap["ssh_user"] = s.SSHUser
	s.fieldMap["pwd"] = s.Pwd
	s.fieldMap["cmd"] = s.Cmd
	s.fieldMap["ret"] = s.Ret
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["note"] = s.Note
	s.fieldMap["transaction_id"] = s.TransactionID
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s serverSSHEvent) clone(db *gorm.DB) serverSSHEvent {
	s.serverSSHEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverSSHEvent) replaceDB(db *gorm.DB) serverSSHEvent {
	s.serverSSHEventDo.ReplaceDB(db)
	return s
}

type serverSSHEventDo struct{ gen.DO }

type IServerSSHEventDo interface {
	gen.SubQuery
	Debug() IServerSSHEventDo
	WithContext(ctx context.Context) IServerSSHEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerSSHEventDo
	WriteDB() IServerSSHEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerSSHEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerSSHEventDo
	Not(conds ...gen.Condition) IServerSSHEventDo
	Or(conds ...gen.Condition) IServerSSHEventDo
	Select(conds ...field.Expr) IServerSSHEventDo
	Where(conds ...gen.Condition) IServerSSHEventDo
	Order(conds ...field.Expr) IServerSSHEventDo
	Distinct(cols ...field.Expr) IServerSSHEventDo
	Omit(cols ...field.Expr) IServerSSHEventDo
	Join(table schema.Tabler, on ...field.Expr) IServerSSHEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerSSHEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerSSHEventDo
	Group(cols ...field.Expr) IServerSSHEventDo
	Having(conds ...gen.Condition) IServerSSHEventDo
	Limit(limit int) IServerSSHEventDo
	Offset(offset int) IServerSSHEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerSSHEventDo
	Unscoped() IServerSSHEventDo
	Create(values ...*model.ServerSSHEvent) error
	CreateInBatches(values []*model.ServerSSHEvent, batchSize int) error
	Save(values ...*model.ServerSSHEvent) error
	First() (*model.ServerSSHEvent, error)
	Take() (*model.ServerSSHEvent, error)
	Last() (*model.ServerSSHEvent, error)
	Find() ([]*model.ServerSSHEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerSSHEvent, err error)
	FindInBatches(result *[]*model.ServerSSHEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerSSHEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerSSHEventDo
	Assign(attrs ...field.AssignExpr) IServerSSHEventDo
	Joins(fields ...field.RelationField) IServerSSHEventDo
	Preload(fields ...field.RelationField) IServerSSHEventDo
	FirstOrInit() (*model.ServerSSHEvent, error)
	FirstOrCreate() (*model.ServerSSHEvent, error)
	FindByPage(offset int, limit int) (result []*model.ServerSSHEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerSSHEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverSSHEventDo) Debug() IServerSSHEventDo {
	return s.withDO(s.DO.Debug())
}

func (s serverSSHEventDo) WithContext(ctx context.Context) IServerSSHEventDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverSSHEventDo) ReadDB() IServerSSHEventDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverSSHEventDo) WriteDB() IServerSSHEventDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverSSHEventDo) Session(config *gorm.Session) IServerSSHEventDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverSSHEventDo) Clauses(conds ...clause.Expression) IServerSSHEventDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverSSHEventDo) Returning(value interface{}, columns ...string) IServerSSHEventDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverSSHEventDo) Not(conds ...gen.Condition) IServerSSHEventDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverSSHEventDo) Or(conds ...gen.Condition) IServerSSHEventDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverSSHEventDo) Select(conds ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverSSHEventDo) Where(conds ...gen.Condition) IServerSSHEventDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverSSHEventDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerSSHEventDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverSSHEventDo) Order(conds ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverSSHEventDo) Distinct(cols ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverSSHEventDo) Omit(cols ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverSSHEventDo) Join(table schema.Tabler, on ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverSSHEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverSSHEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverSSHEventDo) Group(cols ...field.Expr) IServerSSHEventDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverSSHEventDo) Having(conds ...gen.Condition) IServerSSHEventDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverSSHEventDo) Limit(limit int) IServerSSHEventDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverSSHEventDo) Offset(offset int) IServerSSHEventDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverSSHEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerSSHEventDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverSSHEventDo) Unscoped() IServerSSHEventDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverSSHEventDo) Create(values ...*model.ServerSSHEvent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverSSHEventDo) CreateInBatches(values []*model.ServerSSHEvent, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverSSHEventDo) Save(values ...*model.ServerSSHEvent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverSSHEventDo) First() (*model.ServerSSHEvent, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSSHEvent), nil
	}
}

func (s serverSSHEventDo) Take() (*model.ServerSSHEvent, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSSHEvent), nil
	}
}

func (s serverSSHEventDo) Last() (*model.ServerSSHEvent, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSSHEvent), nil
	}
}

func (s serverSSHEventDo) Find() ([]*model.ServerSSHEvent, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerSSHEvent), err
}

func (s serverSSHEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerSSHEvent, err error) {
	buf := make([]*model.ServerSSHEvent, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverSSHEventDo) FindInBatches(result *[]*model.ServerSSHEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverSSHEventDo) Attrs(attrs ...field.AssignExpr) IServerSSHEventDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverSSHEventDo) Assign(attrs ...field.AssignExpr) IServerSSHEventDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverSSHEventDo) Joins(fields ...field.RelationField) IServerSSHEventDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverSSHEventDo) Preload(fields ...field.RelationField) IServerSSHEventDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverSSHEventDo) FirstOrInit() (*model.ServerSSHEvent, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSSHEvent), nil
	}
}

func (s serverSSHEventDo) FirstOrCreate() (*model.ServerSSHEvent, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerSSHEvent), nil
	}
}

func (s serverSSHEventDo) FindByPage(offset int, limit int) (result []*model.ServerSSHEvent, count int64, err error) {
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

func (s serverSSHEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverSSHEventDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverSSHEventDo) Delete(models ...*model.ServerSSHEvent) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverSSHEventDo) withDO(do gen.Dao) *serverSSHEventDo {
	s.DO = *do.(*gen.DO)
	return s
}
