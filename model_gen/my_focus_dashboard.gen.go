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

func newMyFocusDashboard(db *gorm.DB, opts ...gen.DOOption) myFocusDashboard {
	_myFocusDashboard := myFocusDashboard{}

	_myFocusDashboard.myFocusDashboardDo.UseDB(db, opts...)
	_myFocusDashboard.myFocusDashboardDo.UseModel(&model.MyFocusDashboard{})

	tableName := _myFocusDashboard.myFocusDashboardDo.TableName()
	_myFocusDashboard.ALL = field.NewAsterisk(tableName)
	_myFocusDashboard.MyDashboardID = field.NewString(tableName, "my_dashboard_id")
	_myFocusDashboard.DashboardID = field.NewString(tableName, "dashboard_id")
	_myFocusDashboard.AddTime = field.NewTime(tableName, "add_time")
	_myFocusDashboard.Userid = field.NewString(tableName, "userid")

	_myFocusDashboard.fillFieldMap()

	return _myFocusDashboard
}

type myFocusDashboard struct {
	myFocusDashboardDo

	ALL           field.Asterisk
	MyDashboardID field.String // ID
	DashboardID   field.String // 面板ID
	AddTime       field.Time
	Userid        field.String // 用户主键ID

	fieldMap map[string]field.Expr
}

func (m myFocusDashboard) Table(newTableName string) *myFocusDashboard {
	m.myFocusDashboardDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m myFocusDashboard) As(alias string) *myFocusDashboard {
	m.myFocusDashboardDo.DO = *(m.myFocusDashboardDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *myFocusDashboard) updateTableName(table string) *myFocusDashboard {
	m.ALL = field.NewAsterisk(table)
	m.MyDashboardID = field.NewString(table, "my_dashboard_id")
	m.DashboardID = field.NewString(table, "dashboard_id")
	m.AddTime = field.NewTime(table, "add_time")
	m.Userid = field.NewString(table, "userid")

	m.fillFieldMap()

	return m
}

func (m *myFocusDashboard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *myFocusDashboard) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["my_dashboard_id"] = m.MyDashboardID
	m.fieldMap["dashboard_id"] = m.DashboardID
	m.fieldMap["add_time"] = m.AddTime
	m.fieldMap["userid"] = m.Userid
}

func (m myFocusDashboard) clone(db *gorm.DB) myFocusDashboard {
	m.myFocusDashboardDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m myFocusDashboard) replaceDB(db *gorm.DB) myFocusDashboard {
	m.myFocusDashboardDo.ReplaceDB(db)
	return m
}

type myFocusDashboardDo struct{ gen.DO }

type IMyFocusDashboardDo interface {
	gen.SubQuery
	Debug() IMyFocusDashboardDo
	WithContext(ctx context.Context) IMyFocusDashboardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMyFocusDashboardDo
	WriteDB() IMyFocusDashboardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMyFocusDashboardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMyFocusDashboardDo
	Not(conds ...gen.Condition) IMyFocusDashboardDo
	Or(conds ...gen.Condition) IMyFocusDashboardDo
	Select(conds ...field.Expr) IMyFocusDashboardDo
	Where(conds ...gen.Condition) IMyFocusDashboardDo
	Order(conds ...field.Expr) IMyFocusDashboardDo
	Distinct(cols ...field.Expr) IMyFocusDashboardDo
	Omit(cols ...field.Expr) IMyFocusDashboardDo
	Join(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo
	Group(cols ...field.Expr) IMyFocusDashboardDo
	Having(conds ...gen.Condition) IMyFocusDashboardDo
	Limit(limit int) IMyFocusDashboardDo
	Offset(offset int) IMyFocusDashboardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMyFocusDashboardDo
	Unscoped() IMyFocusDashboardDo
	Create(values ...*model.MyFocusDashboard) error
	CreateInBatches(values []*model.MyFocusDashboard, batchSize int) error
	Save(values ...*model.MyFocusDashboard) error
	First() (*model.MyFocusDashboard, error)
	Take() (*model.MyFocusDashboard, error)
	Last() (*model.MyFocusDashboard, error)
	Find() ([]*model.MyFocusDashboard, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MyFocusDashboard, err error)
	FindInBatches(result *[]*model.MyFocusDashboard, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MyFocusDashboard) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMyFocusDashboardDo
	Assign(attrs ...field.AssignExpr) IMyFocusDashboardDo
	Joins(fields ...field.RelationField) IMyFocusDashboardDo
	Preload(fields ...field.RelationField) IMyFocusDashboardDo
	FirstOrInit() (*model.MyFocusDashboard, error)
	FirstOrCreate() (*model.MyFocusDashboard, error)
	FindByPage(offset int, limit int) (result []*model.MyFocusDashboard, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMyFocusDashboardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m myFocusDashboardDo) Debug() IMyFocusDashboardDo {
	return m.withDO(m.DO.Debug())
}

func (m myFocusDashboardDo) WithContext(ctx context.Context) IMyFocusDashboardDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m myFocusDashboardDo) ReadDB() IMyFocusDashboardDo {
	return m.Clauses(dbresolver.Read)
}

func (m myFocusDashboardDo) WriteDB() IMyFocusDashboardDo {
	return m.Clauses(dbresolver.Write)
}

func (m myFocusDashboardDo) Session(config *gorm.Session) IMyFocusDashboardDo {
	return m.withDO(m.DO.Session(config))
}

func (m myFocusDashboardDo) Clauses(conds ...clause.Expression) IMyFocusDashboardDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m myFocusDashboardDo) Returning(value interface{}, columns ...string) IMyFocusDashboardDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m myFocusDashboardDo) Not(conds ...gen.Condition) IMyFocusDashboardDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m myFocusDashboardDo) Or(conds ...gen.Condition) IMyFocusDashboardDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m myFocusDashboardDo) Select(conds ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m myFocusDashboardDo) Where(conds ...gen.Condition) IMyFocusDashboardDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m myFocusDashboardDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMyFocusDashboardDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m myFocusDashboardDo) Order(conds ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m myFocusDashboardDo) Distinct(cols ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m myFocusDashboardDo) Omit(cols ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m myFocusDashboardDo) Join(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m myFocusDashboardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m myFocusDashboardDo) RightJoin(table schema.Tabler, on ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m myFocusDashboardDo) Group(cols ...field.Expr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m myFocusDashboardDo) Having(conds ...gen.Condition) IMyFocusDashboardDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m myFocusDashboardDo) Limit(limit int) IMyFocusDashboardDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m myFocusDashboardDo) Offset(offset int) IMyFocusDashboardDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m myFocusDashboardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMyFocusDashboardDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m myFocusDashboardDo) Unscoped() IMyFocusDashboardDo {
	return m.withDO(m.DO.Unscoped())
}

func (m myFocusDashboardDo) Create(values ...*model.MyFocusDashboard) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m myFocusDashboardDo) CreateInBatches(values []*model.MyFocusDashboard, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m myFocusDashboardDo) Save(values ...*model.MyFocusDashboard) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m myFocusDashboardDo) First() (*model.MyFocusDashboard, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyFocusDashboard), nil
	}
}

func (m myFocusDashboardDo) Take() (*model.MyFocusDashboard, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyFocusDashboard), nil
	}
}

func (m myFocusDashboardDo) Last() (*model.MyFocusDashboard, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyFocusDashboard), nil
	}
}

func (m myFocusDashboardDo) Find() ([]*model.MyFocusDashboard, error) {
	result, err := m.DO.Find()
	return result.([]*model.MyFocusDashboard), err
}

func (m myFocusDashboardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MyFocusDashboard, err error) {
	buf := make([]*model.MyFocusDashboard, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m myFocusDashboardDo) FindInBatches(result *[]*model.MyFocusDashboard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m myFocusDashboardDo) Attrs(attrs ...field.AssignExpr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m myFocusDashboardDo) Assign(attrs ...field.AssignExpr) IMyFocusDashboardDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m myFocusDashboardDo) Joins(fields ...field.RelationField) IMyFocusDashboardDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m myFocusDashboardDo) Preload(fields ...field.RelationField) IMyFocusDashboardDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m myFocusDashboardDo) FirstOrInit() (*model.MyFocusDashboard, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyFocusDashboard), nil
	}
}

func (m myFocusDashboardDo) FirstOrCreate() (*model.MyFocusDashboard, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MyFocusDashboard), nil
	}
}

func (m myFocusDashboardDo) FindByPage(offset int, limit int) (result []*model.MyFocusDashboard, count int64, err error) {
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

func (m myFocusDashboardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m myFocusDashboardDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m myFocusDashboardDo) Delete(models ...*model.MyFocusDashboard) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *myFocusDashboardDo) withDO(do gen.Dao) *myFocusDashboardDo {
	m.DO = *do.(*gen.DO)
	return m
}