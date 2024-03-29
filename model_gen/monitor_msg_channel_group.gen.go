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

func newMonitorMsgChannelGroup(db *gorm.DB, opts ...gen.DOOption) monitorMsgChannelGroup {
	_monitorMsgChannelGroup := monitorMsgChannelGroup{}

	_monitorMsgChannelGroup.monitorMsgChannelGroupDo.UseDB(db, opts...)
	_monitorMsgChannelGroup.monitorMsgChannelGroupDo.UseModel(&model.MonitorMsgChannelGroup{})

	tableName := _monitorMsgChannelGroup.monitorMsgChannelGroupDo.TableName()
	_monitorMsgChannelGroup.ALL = field.NewAsterisk(tableName)
	_monitorMsgChannelGroup.ChannelGroupID = field.NewString(tableName, "channel_group_id")
	_monitorMsgChannelGroup.ChannelGroupCode = field.NewString(tableName, "channel_group_code")
	_monitorMsgChannelGroup.ChannelGroupName = field.NewString(tableName, "channel_group_name")

	_monitorMsgChannelGroup.fillFieldMap()

	return _monitorMsgChannelGroup
}

type monitorMsgChannelGroup struct {
	monitorMsgChannelGroupDo

	ALL              field.Asterisk
	ChannelGroupID   field.String
	ChannelGroupCode field.String
	ChannelGroupName field.String

	fieldMap map[string]field.Expr
}

func (m monitorMsgChannelGroup) Table(newTableName string) *monitorMsgChannelGroup {
	m.monitorMsgChannelGroupDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m monitorMsgChannelGroup) As(alias string) *monitorMsgChannelGroup {
	m.monitorMsgChannelGroupDo.DO = *(m.monitorMsgChannelGroupDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *monitorMsgChannelGroup) updateTableName(table string) *monitorMsgChannelGroup {
	m.ALL = field.NewAsterisk(table)
	m.ChannelGroupID = field.NewString(table, "channel_group_id")
	m.ChannelGroupCode = field.NewString(table, "channel_group_code")
	m.ChannelGroupName = field.NewString(table, "channel_group_name")

	m.fillFieldMap()

	return m
}

func (m *monitorMsgChannelGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *monitorMsgChannelGroup) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 3)
	m.fieldMap["channel_group_id"] = m.ChannelGroupID
	m.fieldMap["channel_group_code"] = m.ChannelGroupCode
	m.fieldMap["channel_group_name"] = m.ChannelGroupName
}

func (m monitorMsgChannelGroup) clone(db *gorm.DB) monitorMsgChannelGroup {
	m.monitorMsgChannelGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m monitorMsgChannelGroup) replaceDB(db *gorm.DB) monitorMsgChannelGroup {
	m.monitorMsgChannelGroupDo.ReplaceDB(db)
	return m
}

type monitorMsgChannelGroupDo struct{ gen.DO }

type IMonitorMsgChannelGroupDo interface {
	gen.SubQuery
	Debug() IMonitorMsgChannelGroupDo
	WithContext(ctx context.Context) IMonitorMsgChannelGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMonitorMsgChannelGroupDo
	WriteDB() IMonitorMsgChannelGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMonitorMsgChannelGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMonitorMsgChannelGroupDo
	Not(conds ...gen.Condition) IMonitorMsgChannelGroupDo
	Or(conds ...gen.Condition) IMonitorMsgChannelGroupDo
	Select(conds ...field.Expr) IMonitorMsgChannelGroupDo
	Where(conds ...gen.Condition) IMonitorMsgChannelGroupDo
	Order(conds ...field.Expr) IMonitorMsgChannelGroupDo
	Distinct(cols ...field.Expr) IMonitorMsgChannelGroupDo
	Omit(cols ...field.Expr) IMonitorMsgChannelGroupDo
	Join(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo
	Group(cols ...field.Expr) IMonitorMsgChannelGroupDo
	Having(conds ...gen.Condition) IMonitorMsgChannelGroupDo
	Limit(limit int) IMonitorMsgChannelGroupDo
	Offset(offset int) IMonitorMsgChannelGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitorMsgChannelGroupDo
	Unscoped() IMonitorMsgChannelGroupDo
	Create(values ...*model.MonitorMsgChannelGroup) error
	CreateInBatches(values []*model.MonitorMsgChannelGroup, batchSize int) error
	Save(values ...*model.MonitorMsgChannelGroup) error
	First() (*model.MonitorMsgChannelGroup, error)
	Take() (*model.MonitorMsgChannelGroup, error)
	Last() (*model.MonitorMsgChannelGroup, error)
	Find() ([]*model.MonitorMsgChannelGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MonitorMsgChannelGroup, err error)
	FindInBatches(result *[]*model.MonitorMsgChannelGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MonitorMsgChannelGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMonitorMsgChannelGroupDo
	Assign(attrs ...field.AssignExpr) IMonitorMsgChannelGroupDo
	Joins(fields ...field.RelationField) IMonitorMsgChannelGroupDo
	Preload(fields ...field.RelationField) IMonitorMsgChannelGroupDo
	FirstOrInit() (*model.MonitorMsgChannelGroup, error)
	FirstOrCreate() (*model.MonitorMsgChannelGroup, error)
	FindByPage(offset int, limit int) (result []*model.MonitorMsgChannelGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMonitorMsgChannelGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m monitorMsgChannelGroupDo) Debug() IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Debug())
}

func (m monitorMsgChannelGroupDo) WithContext(ctx context.Context) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m monitorMsgChannelGroupDo) ReadDB() IMonitorMsgChannelGroupDo {
	return m.Clauses(dbresolver.Read)
}

func (m monitorMsgChannelGroupDo) WriteDB() IMonitorMsgChannelGroupDo {
	return m.Clauses(dbresolver.Write)
}

func (m monitorMsgChannelGroupDo) Session(config *gorm.Session) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Session(config))
}

func (m monitorMsgChannelGroupDo) Clauses(conds ...clause.Expression) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m monitorMsgChannelGroupDo) Returning(value interface{}, columns ...string) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m monitorMsgChannelGroupDo) Not(conds ...gen.Condition) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m monitorMsgChannelGroupDo) Or(conds ...gen.Condition) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m monitorMsgChannelGroupDo) Select(conds ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m monitorMsgChannelGroupDo) Where(conds ...gen.Condition) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m monitorMsgChannelGroupDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMonitorMsgChannelGroupDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m monitorMsgChannelGroupDo) Order(conds ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m monitorMsgChannelGroupDo) Distinct(cols ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m monitorMsgChannelGroupDo) Omit(cols ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m monitorMsgChannelGroupDo) Join(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m monitorMsgChannelGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m monitorMsgChannelGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m monitorMsgChannelGroupDo) Group(cols ...field.Expr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m monitorMsgChannelGroupDo) Having(conds ...gen.Condition) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m monitorMsgChannelGroupDo) Limit(limit int) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m monitorMsgChannelGroupDo) Offset(offset int) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m monitorMsgChannelGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m monitorMsgChannelGroupDo) Unscoped() IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Unscoped())
}

func (m monitorMsgChannelGroupDo) Create(values ...*model.MonitorMsgChannelGroup) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m monitorMsgChannelGroupDo) CreateInBatches(values []*model.MonitorMsgChannelGroup, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m monitorMsgChannelGroupDo) Save(values ...*model.MonitorMsgChannelGroup) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m monitorMsgChannelGroupDo) First() (*model.MonitorMsgChannelGroup, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroup), nil
	}
}

func (m monitorMsgChannelGroupDo) Take() (*model.MonitorMsgChannelGroup, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroup), nil
	}
}

func (m monitorMsgChannelGroupDo) Last() (*model.MonitorMsgChannelGroup, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroup), nil
	}
}

func (m monitorMsgChannelGroupDo) Find() ([]*model.MonitorMsgChannelGroup, error) {
	result, err := m.DO.Find()
	return result.([]*model.MonitorMsgChannelGroup), err
}

func (m monitorMsgChannelGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MonitorMsgChannelGroup, err error) {
	buf := make([]*model.MonitorMsgChannelGroup, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m monitorMsgChannelGroupDo) FindInBatches(result *[]*model.MonitorMsgChannelGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m monitorMsgChannelGroupDo) Attrs(attrs ...field.AssignExpr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m monitorMsgChannelGroupDo) Assign(attrs ...field.AssignExpr) IMonitorMsgChannelGroupDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m monitorMsgChannelGroupDo) Joins(fields ...field.RelationField) IMonitorMsgChannelGroupDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m monitorMsgChannelGroupDo) Preload(fields ...field.RelationField) IMonitorMsgChannelGroupDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m monitorMsgChannelGroupDo) FirstOrInit() (*model.MonitorMsgChannelGroup, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroup), nil
	}
}

func (m monitorMsgChannelGroupDo) FirstOrCreate() (*model.MonitorMsgChannelGroup, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MonitorMsgChannelGroup), nil
	}
}

func (m monitorMsgChannelGroupDo) FindByPage(offset int, limit int) (result []*model.MonitorMsgChannelGroup, count int64, err error) {
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

func (m monitorMsgChannelGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m monitorMsgChannelGroupDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m monitorMsgChannelGroupDo) Delete(models ...*model.MonitorMsgChannelGroup) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *monitorMsgChannelGroupDo) withDO(do gen.Dao) *monitorMsgChannelGroupDo {
	m.DO = *do.(*gen.DO)
	return m
}
