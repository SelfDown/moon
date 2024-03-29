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

func newMessageSendTarget(db *gorm.DB, opts ...gen.DOOption) messageSendTarget {
	_messageSendTarget := messageSendTarget{}

	_messageSendTarget.messageSendTargetDo.UseDB(db, opts...)
	_messageSendTarget.messageSendTargetDo.UseModel(&model.MessageSendTarget{})

	tableName := _messageSendTarget.messageSendTargetDo.TableName()
	_messageSendTarget.ALL = field.NewAsterisk(tableName)
	_messageSendTarget.MessageTargetID = field.NewString(tableName, "message_target_id")
	_messageSendTarget.TargetID = field.NewString(tableName, "target_id")
	_messageSendTarget.TargetType = field.NewString(tableName, "target_type")
	_messageSendTarget.TargetDesc = field.NewString(tableName, "target_desc")
	_messageSendTarget.SysProjectID = field.NewString(tableName, "sys_project_id")
	_messageSendTarget.Statu = field.NewString(tableName, "statu")
	_messageSendTarget.CreateTime = field.NewTime(tableName, "create_time")
	_messageSendTarget.ModifyTime = field.NewTime(tableName, "modify_time")
	_messageSendTarget.Comments = field.NewString(tableName, "comments")

	_messageSendTarget.fillFieldMap()

	return _messageSendTarget
}

type messageSendTarget struct {
	messageSendTargetDo

	ALL             field.Asterisk
	MessageTargetID field.String
	TargetID        field.String // 发送对象ID（比如某个群的ID）
	/*
		对象类型
		0 邮件
		1 QQ
		2 微信

	*/
	TargetType   field.String
	TargetDesc   field.String // 对象描述（比如：某个群的名字）
	SysProjectID field.String // 所属项目
	Statu        field.String // 1 启用 0 禁用
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (m messageSendTarget) Table(newTableName string) *messageSendTarget {
	m.messageSendTargetDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messageSendTarget) As(alias string) *messageSendTarget {
	m.messageSendTargetDo.DO = *(m.messageSendTargetDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messageSendTarget) updateTableName(table string) *messageSendTarget {
	m.ALL = field.NewAsterisk(table)
	m.MessageTargetID = field.NewString(table, "message_target_id")
	m.TargetID = field.NewString(table, "target_id")
	m.TargetType = field.NewString(table, "target_type")
	m.TargetDesc = field.NewString(table, "target_desc")
	m.SysProjectID = field.NewString(table, "sys_project_id")
	m.Statu = field.NewString(table, "statu")
	m.CreateTime = field.NewTime(table, "create_time")
	m.ModifyTime = field.NewTime(table, "modify_time")
	m.Comments = field.NewString(table, "comments")

	m.fillFieldMap()

	return m
}

func (m *messageSendTarget) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messageSendTarget) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["message_target_id"] = m.MessageTargetID
	m.fieldMap["target_id"] = m.TargetID
	m.fieldMap["target_type"] = m.TargetType
	m.fieldMap["target_desc"] = m.TargetDesc
	m.fieldMap["sys_project_id"] = m.SysProjectID
	m.fieldMap["statu"] = m.Statu
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["modify_time"] = m.ModifyTime
	m.fieldMap["comments"] = m.Comments
}

func (m messageSendTarget) clone(db *gorm.DB) messageSendTarget {
	m.messageSendTargetDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messageSendTarget) replaceDB(db *gorm.DB) messageSendTarget {
	m.messageSendTargetDo.ReplaceDB(db)
	return m
}

type messageSendTargetDo struct{ gen.DO }

type IMessageSendTargetDo interface {
	gen.SubQuery
	Debug() IMessageSendTargetDo
	WithContext(ctx context.Context) IMessageSendTargetDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessageSendTargetDo
	WriteDB() IMessageSendTargetDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessageSendTargetDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessageSendTargetDo
	Not(conds ...gen.Condition) IMessageSendTargetDo
	Or(conds ...gen.Condition) IMessageSendTargetDo
	Select(conds ...field.Expr) IMessageSendTargetDo
	Where(conds ...gen.Condition) IMessageSendTargetDo
	Order(conds ...field.Expr) IMessageSendTargetDo
	Distinct(cols ...field.Expr) IMessageSendTargetDo
	Omit(cols ...field.Expr) IMessageSendTargetDo
	Join(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo
	Group(cols ...field.Expr) IMessageSendTargetDo
	Having(conds ...gen.Condition) IMessageSendTargetDo
	Limit(limit int) IMessageSendTargetDo
	Offset(offset int) IMessageSendTargetDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageSendTargetDo
	Unscoped() IMessageSendTargetDo
	Create(values ...*model.MessageSendTarget) error
	CreateInBatches(values []*model.MessageSendTarget, batchSize int) error
	Save(values ...*model.MessageSendTarget) error
	First() (*model.MessageSendTarget, error)
	Take() (*model.MessageSendTarget, error)
	Last() (*model.MessageSendTarget, error)
	Find() ([]*model.MessageSendTarget, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessageSendTarget, err error)
	FindInBatches(result *[]*model.MessageSendTarget, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MessageSendTarget) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessageSendTargetDo
	Assign(attrs ...field.AssignExpr) IMessageSendTargetDo
	Joins(fields ...field.RelationField) IMessageSendTargetDo
	Preload(fields ...field.RelationField) IMessageSendTargetDo
	FirstOrInit() (*model.MessageSendTarget, error)
	FirstOrCreate() (*model.MessageSendTarget, error)
	FindByPage(offset int, limit int) (result []*model.MessageSendTarget, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessageSendTargetDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messageSendTargetDo) Debug() IMessageSendTargetDo {
	return m.withDO(m.DO.Debug())
}

func (m messageSendTargetDo) WithContext(ctx context.Context) IMessageSendTargetDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messageSendTargetDo) ReadDB() IMessageSendTargetDo {
	return m.Clauses(dbresolver.Read)
}

func (m messageSendTargetDo) WriteDB() IMessageSendTargetDo {
	return m.Clauses(dbresolver.Write)
}

func (m messageSendTargetDo) Session(config *gorm.Session) IMessageSendTargetDo {
	return m.withDO(m.DO.Session(config))
}

func (m messageSendTargetDo) Clauses(conds ...clause.Expression) IMessageSendTargetDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messageSendTargetDo) Returning(value interface{}, columns ...string) IMessageSendTargetDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messageSendTargetDo) Not(conds ...gen.Condition) IMessageSendTargetDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messageSendTargetDo) Or(conds ...gen.Condition) IMessageSendTargetDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messageSendTargetDo) Select(conds ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messageSendTargetDo) Where(conds ...gen.Condition) IMessageSendTargetDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messageSendTargetDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMessageSendTargetDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m messageSendTargetDo) Order(conds ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messageSendTargetDo) Distinct(cols ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messageSendTargetDo) Omit(cols ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messageSendTargetDo) Join(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messageSendTargetDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messageSendTargetDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messageSendTargetDo) Group(cols ...field.Expr) IMessageSendTargetDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messageSendTargetDo) Having(conds ...gen.Condition) IMessageSendTargetDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messageSendTargetDo) Limit(limit int) IMessageSendTargetDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messageSendTargetDo) Offset(offset int) IMessageSendTargetDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messageSendTargetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageSendTargetDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messageSendTargetDo) Unscoped() IMessageSendTargetDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messageSendTargetDo) Create(values ...*model.MessageSendTarget) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messageSendTargetDo) CreateInBatches(values []*model.MessageSendTarget, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messageSendTargetDo) Save(values ...*model.MessageSendTarget) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messageSendTargetDo) First() (*model.MessageSendTarget, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendTarget), nil
	}
}

func (m messageSendTargetDo) Take() (*model.MessageSendTarget, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendTarget), nil
	}
}

func (m messageSendTargetDo) Last() (*model.MessageSendTarget, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendTarget), nil
	}
}

func (m messageSendTargetDo) Find() ([]*model.MessageSendTarget, error) {
	result, err := m.DO.Find()
	return result.([]*model.MessageSendTarget), err
}

func (m messageSendTargetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessageSendTarget, err error) {
	buf := make([]*model.MessageSendTarget, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messageSendTargetDo) FindInBatches(result *[]*model.MessageSendTarget, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messageSendTargetDo) Attrs(attrs ...field.AssignExpr) IMessageSendTargetDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messageSendTargetDo) Assign(attrs ...field.AssignExpr) IMessageSendTargetDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messageSendTargetDo) Joins(fields ...field.RelationField) IMessageSendTargetDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messageSendTargetDo) Preload(fields ...field.RelationField) IMessageSendTargetDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messageSendTargetDo) FirstOrInit() (*model.MessageSendTarget, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendTarget), nil
	}
}

func (m messageSendTargetDo) FirstOrCreate() (*model.MessageSendTarget, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessageSendTarget), nil
	}
}

func (m messageSendTargetDo) FindByPage(offset int, limit int) (result []*model.MessageSendTarget, count int64, err error) {
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

func (m messageSendTargetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messageSendTargetDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messageSendTargetDo) Delete(models ...*model.MessageSendTarget) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messageSendTargetDo) withDO(do gen.Dao) *messageSendTargetDo {
	m.DO = *do.(*gen.DO)
	return m
}
