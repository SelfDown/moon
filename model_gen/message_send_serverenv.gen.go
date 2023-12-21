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

func newMessageSendServerEnv(db *gorm.DB, opts ...gen.DOOption) messageSendServerEnv {
	_messageSendServerEnv := messageSendServerEnv{}

	_messageSendServerEnv.messageSendServerEnvDo.UseDB(db, opts...)
	_messageSendServerEnv.messageSendServerEnvDo.UseModel(&model.MessageSendServerEnv{})

	tableName := _messageSendServerEnv.messageSendServerEnvDo.TableName()
	_messageSendServerEnv.ALL = field.NewAsterisk(tableName)
	_messageSendServerEnv.MessageSendServerEnvID = field.NewString(tableName, "message_send_serverEnv_id")
	_messageSendServerEnv.MessageTargetID = field.NewString(tableName, "message_target_id")
	_messageSendServerEnv.ServerEnvID = field.NewString(tableName, "server_env_id")
	_messageSendServerEnv.CreateTime = field.NewTime(tableName, "create_time")
	_messageSendServerEnv.ModifyTime = field.NewTime(tableName, "modify_time")
	_messageSendServerEnv.Comments = field.NewString(tableName, "comments")

	_messageSendServerEnv.fillFieldMap()

	return _messageSendServerEnv
}

type messageSendServerEnv struct {
	messageSendServerEnvDo

	ALL                    field.Asterisk
	MessageSendServerEnvID field.String
	MessageTargetID        field.String // 发送对象ID（比如某个群的ID）
	/*
		对象类型
		0 邮件
		1 QQ
		2 微信

	*/
	ServerEnvID field.String
	CreateTime  field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime  field.Time   // 记录修改时间（数据库自动写入）
	Comments    field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (m messageSendServerEnv) Table(newTableName string) *messageSendServerEnv {
	m.messageSendServerEnvDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messageSendServerEnv) As(alias string) *messageSendServerEnv {
	m.messageSendServerEnvDo.DO = *(m.messageSendServerEnvDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messageSendServerEnv) updateTableName(table string) *messageSendServerEnv {
	m.ALL = field.NewAsterisk(table)
	m.MessageSendServerEnvID = field.NewString(table, "message_send_serverEnv_id")
	m.MessageTargetID = field.NewString(table, "message_target_id")
	m.ServerEnvID = field.NewString(table, "server_env_id")
	m.CreateTime = field.NewTime(table, "create_time")
	m.ModifyTime = field.NewTime(table, "modify_time")
	m.Comments = field.NewString(table, "comments")

	m.fillFieldMap()

	return m
}

func (m *messageSendServerEnv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messageSendServerEnv) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 6)
	m.fieldMap["message_send_serverEnv_id"] = m.MessageSendServerEnvID
	m.fieldMap["message_target_id"] = m.MessageTargetID
	m.fieldMap["server_env_id"] = m.ServerEnvID
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["modify_time"] = m.ModifyTime
	m.fieldMap["comments"] = m.Comments
}

func (m messageSendServerEnv) clone(db *gorm.DB) messageSendServerEnv {
	m.messageSendServerEnvDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messageSendServerEnv) replaceDB(db *gorm.DB) messageSendServerEnv {
	m.messageSendServerEnvDo.ReplaceDB(db)
	return m
}

type messageSendServerEnvDo struct{ gen.DO }

type IMessageSendServerEnvDo interface {
	gen.SubQuery
	Debug() IMessageSendServerEnvDo
	WithContext(ctx context.Context) IMessageSendServerEnvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessageSendServerEnvDo
	WriteDB() IMessageSendServerEnvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessageSendServerEnvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessageSendServerEnvDo
	Not(conds ...gen.Condition) IMessageSendServerEnvDo
	Or(conds ...gen.Condition) IMessageSendServerEnvDo
	Select(conds ...field.Expr) IMessageSendServerEnvDo
	Where(conds ...gen.Condition) IMessageSendServerEnvDo
	Order(conds ...field.Expr) IMessageSendServerEnvDo
	Distinct(cols ...field.Expr) IMessageSendServerEnvDo
	Omit(cols ...field.Expr) IMessageSendServerEnvDo
	Join(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo
	Group(cols ...field.Expr) IMessageSendServerEnvDo
	Having(conds ...gen.Condition) IMessageSendServerEnvDo
	Limit(limit int) IMessageSendServerEnvDo
	Offset(offset int) IMessageSendServerEnvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageSendServerEnvDo
	Unscoped() IMessageSendServerEnvDo
	Create(values ...*model.MessageSendServerEnv) error
	CreateInBatches(values []*model.MessageSendServerEnv, batchSize int) error
	Save(values ...*model.MessageSendServerEnv) error
	First() (*model.MessageSendServerEnv, error)
	Take() (*model.MessageSendServerEnv, error)
	Last() (*model.MessageSendServerEnv, error)
	Find() ([]*model.MessageSendServerEnv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessageSendServerEnv, err error)
	FindInBatches(result *[]*model.MessageSendServerEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MessageSendServerEnv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessageSendServerEnvDo
	Assign(attrs ...field.AssignExpr) IMessageSendServerEnvDo
	Joins(fields ...field.RelationField) IMessageSendServerEnvDo
	Preload(fields ...field.RelationField) IMessageSendServerEnvDo
	FirstOrInit() (*model.MessageSendServerEnv, error)
	FirstOrCreate() (*model.MessageSendServerEnv, error)
	FindByPage(offset int, limit int) (result []*model.MessageSendServerEnv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessageSendServerEnvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messageSendServerEnvDo) Debug() IMessageSendServerEnvDo {
	return m.withDO(m.DO.Debug())
}

func (m messageSendServerEnvDo) WithContext(ctx context.Context) IMessageSendServerEnvDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messageSendServerEnvDo) ReadDB() IMessageSendServerEnvDo {
	return m.Clauses(dbresolver.Read)
}

func (m messageSendServerEnvDo) WriteDB() IMessageSendServerEnvDo {
	return m.Clauses(dbresolver.Write)
}

func (m messageSendServerEnvDo) Session(config *gorm.Session) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Session(config))
}

func (m messageSendServerEnvDo) Clauses(conds ...clause.Expression) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messageSendServerEnvDo) Returning(value interface{}, columns ...string) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messageSendServerEnvDo) Not(conds ...gen.Condition) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messageSendServerEnvDo) Or(conds ...gen.Condition) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messageSendServerEnvDo) Select(conds ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messageSendServerEnvDo) Where(conds ...gen.Condition) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messageSendServerEnvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMessageSendServerEnvDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m messageSendServerEnvDo) Order(conds ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messageSendServerEnvDo) Distinct(cols ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messageSendServerEnvDo) Omit(cols ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messageSendServerEnvDo) Join(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messageSendServerEnvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messageSendServerEnvDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messageSendServerEnvDo) Group(cols ...field.Expr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messageSendServerEnvDo) Having(conds ...gen.Condition) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messageSendServerEnvDo) Limit(limit int) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messageSendServerEnvDo) Offset(offset int) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messageSendServerEnvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messageSendServerEnvDo) Unscoped() IMessageSendServerEnvDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messageSendServerEnvDo) Create(values ...*model.MessageSendServerEnv) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messageSendServerEnvDo) CreateInBatches(values []*model.MessageSendServerEnv, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messageSendServerEnvDo) Save(values ...*model.MessageSendServerEnv) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messageSendServerEnvDo) First() (*model.MessageSendServerEnv, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendServerEnv), nil
	}
}

func (m messageSendServerEnvDo) Take() (*model.MessageSendServerEnv, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendServerEnv), nil
	}
}

func (m messageSendServerEnvDo) Last() (*model.MessageSendServerEnv, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendServerEnv), nil
	}
}

func (m messageSendServerEnvDo) Find() ([]*model.MessageSendServerEnv, error) {
	result, err := m.DO.Find()
	return result.([]*model.MessageSendServerEnv), err
}

func (m messageSendServerEnvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessageSendServerEnv, err error) {
	buf := make([]*model.MessageSendServerEnv, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messageSendServerEnvDo) FindInBatches(result *[]*model.MessageSendServerEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messageSendServerEnvDo) Attrs(attrs ...field.AssignExpr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messageSendServerEnvDo) Assign(attrs ...field.AssignExpr) IMessageSendServerEnvDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messageSendServerEnvDo) Joins(fields ...field.RelationField) IMessageSendServerEnvDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messageSendServerEnvDo) Preload(fields ...field.RelationField) IMessageSendServerEnvDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messageSendServerEnvDo) FirstOrInit() (*model.MessageSendServerEnv, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendServerEnv), nil
	}
}

func (m messageSendServerEnvDo) FirstOrCreate() (*model.MessageSendServerEnv, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendServerEnv), nil
	}
}

func (m messageSendServerEnvDo) FindByPage(offset int, limit int) (result []*model.MessageSendServerEnv, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messageSendServerEnvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messageSendServerEnvDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messageSendServerEnvDo) Delete(models ...*model.MessageSendServerEnv) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messageSendServerEnvDo) withDO(do gen.Dao) *messageSendServerEnvDo {
	m.DO = *do.(*gen.DO)
	return m
}