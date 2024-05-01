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

func newAttendanceRule(db *gorm.DB, opts ...gen.DOOption) attendanceRule {
	_attendanceRule := attendanceRule{}

	_attendanceRule.attendanceRuleDo.UseDB(db, opts...)
	_attendanceRule.attendanceRuleDo.UseModel(&model.AttendanceRule{})

	tableName := _attendanceRule.attendanceRuleDo.TableName()
	_attendanceRule.ALL = field.NewAsterisk(tableName)
	_attendanceRule.RuleID = field.NewString(tableName, "rule_id")
	_attendanceRule.RuleGroupID = field.NewString(tableName, "rule_group_id")
	_attendanceRule.Name = field.NewString(tableName, "name")
	_attendanceRule.RuleCode = field.NewString(tableName, "rule_code")
	_attendanceRule.RuleType = field.NewString(tableName, "rule_type")
	_attendanceRule.RuleTime = field.NewString(tableName, "rule_time")

	_attendanceRule.fillFieldMap()

	return _attendanceRule
}

type attendanceRule struct {
	attendanceRuleDo

	ALL         field.Asterisk
	RuleID      field.String
	RuleGroupID field.String
	Name        field.String
	RuleCode    field.String
	RuleType    field.String
	RuleTime    field.String

	fieldMap map[string]field.Expr
}

func (a attendanceRule) Table(newTableName string) *attendanceRule {
	a.attendanceRuleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a attendanceRule) As(alias string) *attendanceRule {
	a.attendanceRuleDo.DO = *(a.attendanceRuleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *attendanceRule) updateTableName(table string) *attendanceRule {
	a.ALL = field.NewAsterisk(table)
	a.RuleID = field.NewString(table, "rule_id")
	a.RuleGroupID = field.NewString(table, "rule_group_id")
	a.Name = field.NewString(table, "name")
	a.RuleCode = field.NewString(table, "rule_code")
	a.RuleType = field.NewString(table, "rule_type")
	a.RuleTime = field.NewString(table, "rule_time")

	a.fillFieldMap()

	return a
}

func (a *attendanceRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *attendanceRule) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["rule_id"] = a.RuleID
	a.fieldMap["rule_group_id"] = a.RuleGroupID
	a.fieldMap["name"] = a.Name
	a.fieldMap["rule_code"] = a.RuleCode
	a.fieldMap["rule_type"] = a.RuleType
	a.fieldMap["rule_time"] = a.RuleTime
}

func (a attendanceRule) clone(db *gorm.DB) attendanceRule {
	a.attendanceRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a attendanceRule) replaceDB(db *gorm.DB) attendanceRule {
	a.attendanceRuleDo.ReplaceDB(db)
	return a
}

type attendanceRuleDo struct{ gen.DO }

type IAttendanceRuleDo interface {
	gen.SubQuery
	Debug() IAttendanceRuleDo
	WithContext(ctx context.Context) IAttendanceRuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAttendanceRuleDo
	WriteDB() IAttendanceRuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAttendanceRuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAttendanceRuleDo
	Not(conds ...gen.Condition) IAttendanceRuleDo
	Or(conds ...gen.Condition) IAttendanceRuleDo
	Select(conds ...field.Expr) IAttendanceRuleDo
	Where(conds ...gen.Condition) IAttendanceRuleDo
	Order(conds ...field.Expr) IAttendanceRuleDo
	Distinct(cols ...field.Expr) IAttendanceRuleDo
	Omit(cols ...field.Expr) IAttendanceRuleDo
	Join(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo
	Group(cols ...field.Expr) IAttendanceRuleDo
	Having(conds ...gen.Condition) IAttendanceRuleDo
	Limit(limit int) IAttendanceRuleDo
	Offset(offset int) IAttendanceRuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAttendanceRuleDo
	Unscoped() IAttendanceRuleDo
	Create(values ...*model.AttendanceRule) error
	CreateInBatches(values []*model.AttendanceRule, batchSize int) error
	Save(values ...*model.AttendanceRule) error
	First() (*model.AttendanceRule, error)
	Take() (*model.AttendanceRule, error)
	Last() (*model.AttendanceRule, error)
	Find() ([]*model.AttendanceRule, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AttendanceRule, err error)
	FindInBatches(result *[]*model.AttendanceRule, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AttendanceRule) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAttendanceRuleDo
	Assign(attrs ...field.AssignExpr) IAttendanceRuleDo
	Joins(fields ...field.RelationField) IAttendanceRuleDo
	Preload(fields ...field.RelationField) IAttendanceRuleDo
	FirstOrInit() (*model.AttendanceRule, error)
	FirstOrCreate() (*model.AttendanceRule, error)
	FindByPage(offset int, limit int) (result []*model.AttendanceRule, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAttendanceRuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a attendanceRuleDo) Debug() IAttendanceRuleDo {
	return a.withDO(a.DO.Debug())
}

func (a attendanceRuleDo) WithContext(ctx context.Context) IAttendanceRuleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a attendanceRuleDo) ReadDB() IAttendanceRuleDo {
	return a.Clauses(dbresolver.Read)
}

func (a attendanceRuleDo) WriteDB() IAttendanceRuleDo {
	return a.Clauses(dbresolver.Write)
}

func (a attendanceRuleDo) Session(config *gorm.Session) IAttendanceRuleDo {
	return a.withDO(a.DO.Session(config))
}

func (a attendanceRuleDo) Clauses(conds ...clause.Expression) IAttendanceRuleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a attendanceRuleDo) Returning(value interface{}, columns ...string) IAttendanceRuleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a attendanceRuleDo) Not(conds ...gen.Condition) IAttendanceRuleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a attendanceRuleDo) Or(conds ...gen.Condition) IAttendanceRuleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a attendanceRuleDo) Select(conds ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a attendanceRuleDo) Where(conds ...gen.Condition) IAttendanceRuleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a attendanceRuleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAttendanceRuleDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a attendanceRuleDo) Order(conds ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a attendanceRuleDo) Distinct(cols ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a attendanceRuleDo) Omit(cols ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a attendanceRuleDo) Join(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a attendanceRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a attendanceRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a attendanceRuleDo) Group(cols ...field.Expr) IAttendanceRuleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a attendanceRuleDo) Having(conds ...gen.Condition) IAttendanceRuleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a attendanceRuleDo) Limit(limit int) IAttendanceRuleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a attendanceRuleDo) Offset(offset int) IAttendanceRuleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a attendanceRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAttendanceRuleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a attendanceRuleDo) Unscoped() IAttendanceRuleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a attendanceRuleDo) Create(values ...*model.AttendanceRule) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a attendanceRuleDo) CreateInBatches(values []*model.AttendanceRule, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a attendanceRuleDo) Save(values ...*model.AttendanceRule) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a attendanceRuleDo) First() (*model.AttendanceRule, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttendanceRule), nil
	}
}

func (a attendanceRuleDo) Take() (*model.AttendanceRule, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttendanceRule), nil
	}
}

func (a attendanceRuleDo) Last() (*model.AttendanceRule, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttendanceRule), nil
	}
}

func (a attendanceRuleDo) Find() ([]*model.AttendanceRule, error) {
	result, err := a.DO.Find()
	return result.([]*model.AttendanceRule), err
}

func (a attendanceRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AttendanceRule, err error) {
	buf := make([]*model.AttendanceRule, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a attendanceRuleDo) FindInBatches(result *[]*model.AttendanceRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a attendanceRuleDo) Attrs(attrs ...field.AssignExpr) IAttendanceRuleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a attendanceRuleDo) Assign(attrs ...field.AssignExpr) IAttendanceRuleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a attendanceRuleDo) Joins(fields ...field.RelationField) IAttendanceRuleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a attendanceRuleDo) Preload(fields ...field.RelationField) IAttendanceRuleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a attendanceRuleDo) FirstOrInit() (*model.AttendanceRule, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttendanceRule), nil
	}
}

func (a attendanceRuleDo) FirstOrCreate() (*model.AttendanceRule, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttendanceRule), nil
	}
}

func (a attendanceRuleDo) FindByPage(offset int, limit int) (result []*model.AttendanceRule, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a attendanceRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a attendanceRuleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a attendanceRuleDo) Delete(models ...*model.AttendanceRule) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *attendanceRuleDo) withDO(do gen.Dao) *attendanceRuleDo {
	a.DO = *do.(*gen.DO)
	return a
}
