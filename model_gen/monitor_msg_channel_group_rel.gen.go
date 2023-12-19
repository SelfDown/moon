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

func newMonitorMsgChannelGroupRel(db *gorm.DB, opts ...gen.DOOption) monitorMsgChannelGroupRel {
	_monitorMsgChannelGroupRel := monitorMsgChannelGroupRel{}

	_monitorMsgChannelGroupRel.monitorMsgChannelGroupRelDo.UseDB(db, opts...)
	_monitorMsgChannelGroupRel.monitorMsgChannelGroupRelDo.UseModel(&model.MonitorMsgChannelGroupRel{})

	tableName := _monitorMsgChannelGroupRel.monitorMsgChannelGroupRelDo.TableName()
	_monitorMsgChannelGroupRel.ALL = field.NewAsterisk(tableName)
	_monitorMsgChannelGroupRel.ChannelGroupRelID = field.NewString(tableName, "channel_group_rel_id")
	_monitorMsgChannelGroupRel.ChannelGroupID = field.NewString(tableName, "channel_group_id")
	_monitorMsgChannelGroupRel.ChannelID = field.NewString(tableName, "channel_id")

	_monitorMsgChannelGroupRel.fillFieldMap()

	return _monitorMsgChannelGroupRel
}

type monitorMsgChannelGroupRel struct {
	monitorMsgChannelGroupRelDo

	ALL               field.Asterisk
	ChannelGroupRelID field.String
	ChannelGroupID    field.String
	ChannelID         field.String

	fieldMap map[string]field.Expr
}

func (m monitorMsgChannelGroupRel) Table(newTableName string) *monitorMsgChannelGroupRel {
	m.monitorMsgChannelGroupRelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m monitorMsgChannelGroupRel) As(alias string) *monitorMsgChannelGroupRel {
	m.monitorMsgChannelGroupRelDo.DO = *(m.monitorMsgChannelGroupRelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *monitorMsgChannelGroupRel) updateTableName(table string) *monitorMsgChannelGroupRel {
	m.ALL = field.NewAsterisk(table)
	m.ChannelGroupRelID = field.NewString(table, "channel_group_rel_id")
	m.ChannelGroupID = field.NewString(table, "channel_group_id")
	m.ChannelID = field.NewString(table, "channel_id")

	m.fillFieldMap()

	return m
}

func (m *monitorMsgChannelGroupRel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *monitorMsgChannelGroupRel) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 3)
	m.fieldMap["channel_group_rel_id"] = m.ChannelGroupRelID
	m.fieldMap["channel_group_id"] = m.ChannelGroupID
	m.fieldMap["channel_id"] = m.ChannelID
}

func (m monitorMsgChannelGroupRel) clone(db *gorm.DB) monitorMsgChannelGroupRel {
	m.monitorMsgChannelGroupRelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m monitorMsgChannelGroupRel) replaceDB(db *gorm.DB) monitorMsgChannelGroupRel {
	m.monitorMsgChannelGroupRelDo.ReplaceDB(db)
	return m
}

type monitorMsgChannelGroupRelDo struct{ gen.DO }

type IMonitorMsgChannelGroupRelDo interface {
	gen.SubQuery
	Debug() IMonitorMsgChannelGroupRelDo
	WithContext(ctx context.Context) IMonitorMsgChannelGroupRelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMonitorMsgChannelGroupRelDo
	WriteDB() IMonitorMsgChannelGroupRelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMonitorMsgChannelGroupRelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMonitorMsgChannelGroupRelDo
	Not(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo
	Or(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo
	Select(conds ...field.Expr) IMonitorMsgChannelGroupRelDo
	Where(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo
	Order(conds ...field.Expr) IMonitorMsgChannelGroupRelDo
	Distinct(cols ...field.Expr) IMonitorMsgChannelGroupRelDo
	Omit(cols ...field.Expr) IMonitorMsgChannelGroupRelDo
	Join(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo
	Group(cols ...field.Expr) IMonitorMsgChannelGroupRelDo
	Having(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo
	Limit(limit int) IMonitorMsgChannelGroupRelDo
	Offset(offset int) IMonitorMsgChannelGroupRelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitorMsgChannelGroupRelDo
	Unscoped() IMonitorMsgChannelGroupRelDo
	Create(values ...*model.MonitorMsgChannelGroupRel) error
	CreateInBatches(values []*model.MonitorMsgChannelGroupRel, batchSize int) error
	Save(values ...*model.MonitorMsgChannelGroupRel) error
	First() (*model.MonitorMsgChannelGroupRel, error)
	Take() (*model.MonitorMsgChannelGroupRel, error)
	Last() (*model.MonitorMsgChannelGroupRel, error)
	Find() ([]*model.MonitorMsgChannelGroupRel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MonitorMsgChannelGroupRel, err error)
	FindInBatches(result *[]*model.MonitorMsgChannelGroupRel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MonitorMsgChannelGroupRel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMonitorMsgChannelGroupRelDo
	Assign(attrs ...field.AssignExpr) IMonitorMsgChannelGroupRelDo
	Joins(fields ...field.RelationField) IMonitorMsgChannelGroupRelDo
	Preload(fields ...field.RelationField) IMonitorMsgChannelGroupRelDo
	FirstOrInit() (*model.MonitorMsgChannelGroupRel, error)
	FirstOrCreate() (*model.MonitorMsgChannelGroupRel, error)
	FindByPage(offset int, limit int) (result []*model.MonitorMsgChannelGroupRel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMonitorMsgChannelGroupRelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m monitorMsgChannelGroupRelDo) Debug() IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Debug())
}

func (m monitorMsgChannelGroupRelDo) WithContext(ctx context.Context) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m monitorMsgChannelGroupRelDo) ReadDB() IMonitorMsgChannelGroupRelDo {
	return m.Clauses(dbresolver.Read)
}

func (m monitorMsgChannelGroupRelDo) WriteDB() IMonitorMsgChannelGroupRelDo {
	return m.Clauses(dbresolver.Write)
}

func (m monitorMsgChannelGroupRelDo) Session(config *gorm.Session) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Session(config))
}

func (m monitorMsgChannelGroupRelDo) Clauses(conds ...clause.Expression) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m monitorMsgChannelGroupRelDo) Returning(value interface{}, columns ...string) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m monitorMsgChannelGroupRelDo) Not(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m monitorMsgChannelGroupRelDo) Or(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m monitorMsgChannelGroupRelDo) Select(conds ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m monitorMsgChannelGroupRelDo) Where(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m monitorMsgChannelGroupRelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMonitorMsgChannelGroupRelDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m monitorMsgChannelGroupRelDo) Order(conds ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m monitorMsgChannelGroupRelDo) Distinct(cols ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m monitorMsgChannelGroupRelDo) Omit(cols ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m monitorMsgChannelGroupRelDo) Join(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m monitorMsgChannelGroupRelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m monitorMsgChannelGroupRelDo) RightJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m monitorMsgChannelGroupRelDo) Group(cols ...field.Expr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m monitorMsgChannelGroupRelDo) Having(conds ...gen.Condition) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m monitorMsgChannelGroupRelDo) Limit(limit int) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m monitorMsgChannelGroupRelDo) Offset(offset int) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m monitorMsgChannelGroupRelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m monitorMsgChannelGroupRelDo) Unscoped() IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m monitorMsgChannelGroupRelDo) Create(values ...*model.MonitorMsgChannelGroupRel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m monitorMsgChannelGroupRelDo) CreateInBatches(values []*model.MonitorMsgChannelGroupRel, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m monitorMsgChannelGroupRelDo) Save(values ...*model.MonitorMsgChannelGroupRel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m monitorMsgChannelGroupRelDo) First() (*model.MonitorMsgChannelGroupRel, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroupRel), nil
	}
}

func (m monitorMsgChannelGroupRelDo) Take() (*model.MonitorMsgChannelGroupRel, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroupRel), nil
	}
}

func (m monitorMsgChannelGroupRelDo) Last() (*model.MonitorMsgChannelGroupRel, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroupRel), nil
	}
}

func (m monitorMsgChannelGroupRelDo) Find() ([]*model.MonitorMsgChannelGroupRel, error) {
	result, err := m.DO.Find()
	return result.([]*model.MonitorMsgChannelGroupRel), err
}

func (m monitorMsgChannelGroupRelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MonitorMsgChannelGroupRel, err error) {
	buf := make([]*model.MonitorMsgChannelGroupRel, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m monitorMsgChannelGroupRelDo) FindInBatches(result *[]*model.MonitorMsgChannelGroupRel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m monitorMsgChannelGroupRelDo) Attrs(attrs ...field.AssignExpr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m monitorMsgChannelGroupRelDo) Assign(attrs ...field.AssignExpr) IMonitorMsgChannelGroupRelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m monitorMsgChannelGroupRelDo) Joins(fields ...field.RelationField) IMonitorMsgChannelGroupRelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m monitorMsgChannelGroupRelDo) Preload(fields ...field.RelationField) IMonitorMsgChannelGroupRelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m monitorMsgChannelGroupRelDo) FirstOrInit() (*model.MonitorMsgChannelGroupRel, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroupRel), nil
	}
}

func (m monitorMsgChannelGroupRelDo) FirstOrCreate() (*model.MonitorMsgChannelGroupRel, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroupRel), nil
	}
}

func (m monitorMsgChannelGroupRelDo) FindByPage(offset int, limit int) (result []*model.MonitorMsgChannelGroupRel, count int64, err error) {
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

func (m monitorMsgChannelGroupRelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m monitorMsgChannelGroupRelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m monitorMsgChannelGroupRelDo) Delete(models ...*model.MonitorMsgChannelGroupRel) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *monitorMsgChannelGroupRelDo) withDO(do gen.Dao) *monitorMsgChannelGroupRelDo {
	m.DO = *do.(*gen.DO)
	return m
}
