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

func newServerVpsEvent(db *gorm.DB, opts ...gen.DOOption) serverVpsEvent {
	_serverVpsEvent := serverVpsEvent{}

	_serverVpsEvent.serverVpsEventDo.UseDB(db, opts...)
	_serverVpsEvent.serverVpsEventDo.UseModel(&model.ServerVpsEvent{})

	tableName := _serverVpsEvent.serverVpsEventDo.TableName()
	_serverVpsEvent.ALL = field.NewAsterisk(tableName)
	_serverVpsEvent.ServerVpsEventID = field.NewInt32(tableName, "server_vps_event_id")
	_serverVpsEvent.CreateTime = field.NewTime(tableName, "create_time")
	_serverVpsEvent.ModifyTime = field.NewTime(tableName, "modify_time")
	_serverVpsEvent.UserID = field.NewString(tableName, "user_id")
	_serverVpsEvent.Comments = field.NewString(tableName, "comments")
	_serverVpsEvent.VpsEventTypeID = field.NewString(tableName, "vps_event_type_id")
	_serverVpsEvent.ServerID = field.NewString(tableName, "server_id")
	_serverVpsEvent.EventID = field.NewString(tableName, "event_id")

	_serverVpsEvent.fillFieldMap()

	return _serverVpsEvent
}

type serverVpsEvent struct {
	serverVpsEventDo

	ALL              field.Asterisk
	ServerVpsEventID field.Int32
	CreateTime       field.Time // 记录创建时间（数据库自动写入）
	ModifyTime       field.Time
	UserID           field.String // 操作用户
	/*
		事件操作细节
		比如 ：正在克隆虚拟机
	*/
	Comments field.String
	/*
		事件类型与 server_vps_event_type,主键关联

	*/
	VpsEventTypeID field.String
	ServerID       field.String
	/*
		事件ID。
		比如用户申请创建5个虚拟机，那么这个算一次事件，我们将创建5个虚拟机的后台细节全部写入到表中，
		并且事件ID一样，这样前台可以根据事件ID，获取事件状态。
	*/
	EventID field.String

	fieldMap map[string]field.Expr
}

func (s serverVpsEvent) Table(newTableName string) *serverVpsEvent {
	s.serverVpsEventDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverVpsEvent) As(alias string) *serverVpsEvent {
	s.serverVpsEventDo.DO = *(s.serverVpsEventDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverVpsEvent) updateTableName(table string) *serverVpsEvent {
	s.ALL = field.NewAsterisk(table)
	s.ServerVpsEventID = field.NewInt32(table, "server_vps_event_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.UserID = field.NewString(table, "user_id")
	s.Comments = field.NewString(table, "comments")
	s.VpsEventTypeID = field.NewString(table, "vps_event_type_id")
	s.ServerID = field.NewString(table, "server_id")
	s.EventID = field.NewString(table, "event_id")

	s.fillFieldMap()

	return s
}

func (s *serverVpsEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverVpsEvent) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["server_vps_event_id"] = s.ServerVpsEventID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["vps_event_type_id"] = s.VpsEventTypeID
	s.fieldMap["server_id"] = s.ServerID
	s.fieldMap["event_id"] = s.EventID
}

func (s serverVpsEvent) clone(db *gorm.DB) serverVpsEvent {
	s.serverVpsEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverVpsEvent) replaceDB(db *gorm.DB) serverVpsEvent {
	s.serverVpsEventDo.ReplaceDB(db)
	return s
}

type serverVpsEventDo struct{ gen.DO }

type IServerVpsEventDo interface {
	gen.SubQuery
	Debug() IServerVpsEventDo
	WithContext(ctx context.Context) IServerVpsEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerVpsEventDo
	WriteDB() IServerVpsEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerVpsEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerVpsEventDo
	Not(conds ...gen.Condition) IServerVpsEventDo
	Or(conds ...gen.Condition) IServerVpsEventDo
	Select(conds ...field.Expr) IServerVpsEventDo
	Where(conds ...gen.Condition) IServerVpsEventDo
	Order(conds ...field.Expr) IServerVpsEventDo
	Distinct(cols ...field.Expr) IServerVpsEventDo
	Omit(cols ...field.Expr) IServerVpsEventDo
	Join(table schema.Tabler, on ...field.Expr) IServerVpsEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerVpsEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerVpsEventDo
	Group(cols ...field.Expr) IServerVpsEventDo
	Having(conds ...gen.Condition) IServerVpsEventDo
	Limit(limit int) IServerVpsEventDo
	Offset(offset int) IServerVpsEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerVpsEventDo
	Unscoped() IServerVpsEventDo
	Create(values ...*model.ServerVpsEvent) error
	CreateInBatches(values []*model.ServerVpsEvent, batchSize int) error
	Save(values ...*model.ServerVpsEvent) error
	First() (*model.ServerVpsEvent, error)
	Take() (*model.ServerVpsEvent, error)
	Last() (*model.ServerVpsEvent, error)
	Find() ([]*model.ServerVpsEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerVpsEvent, err error)
	FindInBatches(result *[]*model.ServerVpsEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerVpsEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerVpsEventDo
	Assign(attrs ...field.AssignExpr) IServerVpsEventDo
	Joins(fields ...field.RelationField) IServerVpsEventDo
	Preload(fields ...field.RelationField) IServerVpsEventDo
	FirstOrInit() (*model.ServerVpsEvent, error)
	FirstOrCreate() (*model.ServerVpsEvent, error)
	FindByPage(offset int, limit int) (result []*model.ServerVpsEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerVpsEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverVpsEventDo) Debug() IServerVpsEventDo {
	return s.withDO(s.DO.Debug())
}

func (s serverVpsEventDo) WithContext(ctx context.Context) IServerVpsEventDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverVpsEventDo) ReadDB() IServerVpsEventDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverVpsEventDo) WriteDB() IServerVpsEventDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverVpsEventDo) Session(config *gorm.Session) IServerVpsEventDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverVpsEventDo) Clauses(conds ...clause.Expression) IServerVpsEventDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverVpsEventDo) Returning(value interface{}, columns ...string) IServerVpsEventDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverVpsEventDo) Not(conds ...gen.Condition) IServerVpsEventDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverVpsEventDo) Or(conds ...gen.Condition) IServerVpsEventDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverVpsEventDo) Select(conds ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverVpsEventDo) Where(conds ...gen.Condition) IServerVpsEventDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverVpsEventDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerVpsEventDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverVpsEventDo) Order(conds ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverVpsEventDo) Distinct(cols ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverVpsEventDo) Omit(cols ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverVpsEventDo) Join(table schema.Tabler, on ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverVpsEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverVpsEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverVpsEventDo) Group(cols ...field.Expr) IServerVpsEventDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverVpsEventDo) Having(conds ...gen.Condition) IServerVpsEventDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverVpsEventDo) Limit(limit int) IServerVpsEventDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverVpsEventDo) Offset(offset int) IServerVpsEventDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverVpsEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerVpsEventDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverVpsEventDo) Unscoped() IServerVpsEventDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverVpsEventDo) Create(values ...*model.ServerVpsEvent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverVpsEventDo) CreateInBatches(values []*model.ServerVpsEvent, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverVpsEventDo) Save(values ...*model.ServerVpsEvent) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverVpsEventDo) First() (*model.ServerVpsEvent, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerVpsEvent), nil
	}
}

func (s serverVpsEventDo) Take() (*model.ServerVpsEvent, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerVpsEvent), nil
	}
}

func (s serverVpsEventDo) Last() (*model.ServerVpsEvent, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerVpsEvent), nil
	}
}

func (s serverVpsEventDo) Find() ([]*model.ServerVpsEvent, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerVpsEvent), err
}

func (s serverVpsEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerVpsEvent, err error) {
	buf := make([]*model.ServerVpsEvent, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverVpsEventDo) FindInBatches(result *[]*model.ServerVpsEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverVpsEventDo) Attrs(attrs ...field.AssignExpr) IServerVpsEventDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverVpsEventDo) Assign(attrs ...field.AssignExpr) IServerVpsEventDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverVpsEventDo) Joins(fields ...field.RelationField) IServerVpsEventDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverVpsEventDo) Preload(fields ...field.RelationField) IServerVpsEventDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverVpsEventDo) FirstOrInit() (*model.ServerVpsEvent, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerVpsEvent), nil
	}
}

func (s serverVpsEventDo) FirstOrCreate() (*model.ServerVpsEvent, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerVpsEvent), nil
	}
}

func (s serverVpsEventDo) FindByPage(offset int, limit int) (result []*model.ServerVpsEvent, count int64, err error) {
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

func (s serverVpsEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverVpsEventDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverVpsEventDo) Delete(models ...*model.ServerVpsEvent) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverVpsEventDo) withDO(do gen.Dao) *serverVpsEventDo {
	s.DO = *do.(*gen.DO)
	return s
}
