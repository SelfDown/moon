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

func newSysLabels(db *gorm.DB, opts ...gen.DOOption) sysLabels {
	_sysLabels := sysLabels{}

	_sysLabels.sysLabelsDo.UseDB(db, opts...)
	_sysLabels.sysLabelsDo.UseModel(&model.SysLabels{})

	tableName := _sysLabels.sysLabelsDo.TableName()
	_sysLabels.ALL = field.NewAsterisk(tableName)
	_sysLabels.LabelID = field.NewString(tableName, "label_id")
	_sysLabels.LabelGroup = field.NewString(tableName, "label_group")
	_sysLabels.LabelName = field.NewString(tableName, "label_name")
	_sysLabels.OperatorUser = field.NewString(tableName, "operator_user")
	_sysLabels.OperatorTime = field.NewTime(tableName, "operator_time")

	_sysLabels.fillFieldMap()

	return _sysLabels
}

type sysLabels struct {
	sysLabelsDo

	ALL          field.Asterisk
	LabelID      field.String
	LabelGroup   field.String
	LabelName    field.String
	OperatorUser field.String
	OperatorTime field.Time

	fieldMap map[string]field.Expr
}

func (s sysLabels) Table(newTableName string) *sysLabels {
	s.sysLabelsDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysLabels) As(alias string) *sysLabels {
	s.sysLabelsDo.DO = *(s.sysLabelsDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysLabels) updateTableName(table string) *sysLabels {
	s.ALL = field.NewAsterisk(table)
	s.LabelID = field.NewString(table, "label_id")
	s.LabelGroup = field.NewString(table, "label_group")
	s.LabelName = field.NewString(table, "label_name")
	s.OperatorUser = field.NewString(table, "operator_user")
	s.OperatorTime = field.NewTime(table, "operator_time")

	s.fillFieldMap()

	return s
}

func (s *sysLabels) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysLabels) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["label_id"] = s.LabelID
	s.fieldMap["label_group"] = s.LabelGroup
	s.fieldMap["label_name"] = s.LabelName
	s.fieldMap["operator_user"] = s.OperatorUser
	s.fieldMap["operator_time"] = s.OperatorTime
}

func (s sysLabels) clone(db *gorm.DB) sysLabels {
	s.sysLabelsDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysLabels) replaceDB(db *gorm.DB) sysLabels {
	s.sysLabelsDo.ReplaceDB(db)
	return s
}

type sysLabelsDo struct{ gen.DO }

type ISysLabelsDo interface {
	gen.SubQuery
	Debug() ISysLabelsDo
	WithContext(ctx context.Context) ISysLabelsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysLabelsDo
	WriteDB() ISysLabelsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysLabelsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysLabelsDo
	Not(conds ...gen.Condition) ISysLabelsDo
	Or(conds ...gen.Condition) ISysLabelsDo
	Select(conds ...field.Expr) ISysLabelsDo
	Where(conds ...gen.Condition) ISysLabelsDo
	Order(conds ...field.Expr) ISysLabelsDo
	Distinct(cols ...field.Expr) ISysLabelsDo
	Omit(cols ...field.Expr) ISysLabelsDo
	Join(table schema.Tabler, on ...field.Expr) ISysLabelsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysLabelsDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysLabelsDo
	Group(cols ...field.Expr) ISysLabelsDo
	Having(conds ...gen.Condition) ISysLabelsDo
	Limit(limit int) ISysLabelsDo
	Offset(offset int) ISysLabelsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLabelsDo
	Unscoped() ISysLabelsDo
	Create(values ...*model.SysLabels) error
	CreateInBatches(values []*model.SysLabels, batchSize int) error
	Save(values ...*model.SysLabels) error
	First() (*model.SysLabels, error)
	Take() (*model.SysLabels, error)
	Last() (*model.SysLabels, error)
	Find() ([]*model.SysLabels, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLabels, err error)
	FindInBatches(result *[]*model.SysLabels, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysLabels) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysLabelsDo
	Assign(attrs ...field.AssignExpr) ISysLabelsDo
	Joins(fields ...field.RelationField) ISysLabelsDo
	Preload(fields ...field.RelationField) ISysLabelsDo
	FirstOrInit() (*model.SysLabels, error)
	FirstOrCreate() (*model.SysLabels, error)
	FindByPage(offset int, limit int) (result []*model.SysLabels, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysLabelsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysLabelsDo) Debug() ISysLabelsDo {
	return s.withDO(s.DO.Debug())
}

func (s sysLabelsDo) WithContext(ctx context.Context) ISysLabelsDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysLabelsDo) ReadDB() ISysLabelsDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysLabelsDo) WriteDB() ISysLabelsDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysLabelsDo) Session(config *gorm.Session) ISysLabelsDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysLabelsDo) Clauses(conds ...clause.Expression) ISysLabelsDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysLabelsDo) Returning(value interface{}, columns ...string) ISysLabelsDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysLabelsDo) Not(conds ...gen.Condition) ISysLabelsDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysLabelsDo) Or(conds ...gen.Condition) ISysLabelsDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysLabelsDo) Select(conds ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysLabelsDo) Where(conds ...gen.Condition) ISysLabelsDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysLabelsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysLabelsDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysLabelsDo) Order(conds ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysLabelsDo) Distinct(cols ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysLabelsDo) Omit(cols ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysLabelsDo) Join(table schema.Tabler, on ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysLabelsDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysLabelsDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysLabelsDo) Group(cols ...field.Expr) ISysLabelsDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysLabelsDo) Having(conds ...gen.Condition) ISysLabelsDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysLabelsDo) Limit(limit int) ISysLabelsDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysLabelsDo) Offset(offset int) ISysLabelsDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysLabelsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysLabelsDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysLabelsDo) Unscoped() ISysLabelsDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysLabelsDo) Create(values ...*model.SysLabels) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysLabelsDo) CreateInBatches(values []*model.SysLabels, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysLabelsDo) Save(values ...*model.SysLabels) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysLabelsDo) First() (*model.SysLabels, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLabels), nil
	}
}

func (s sysLabelsDo) Take() (*model.SysLabels, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLabels), nil
	}
}

func (s sysLabelsDo) Last() (*model.SysLabels, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLabels), nil
	}
}

func (s sysLabelsDo) Find() ([]*model.SysLabels, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysLabels), err
}

func (s sysLabelsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLabels, err error) {
	buf := make([]*model.SysLabels, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysLabelsDo) FindInBatches(result *[]*model.SysLabels, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysLabelsDo) Attrs(attrs ...field.AssignExpr) ISysLabelsDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysLabelsDo) Assign(attrs ...field.AssignExpr) ISysLabelsDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysLabelsDo) Joins(fields ...field.RelationField) ISysLabelsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysLabelsDo) Preload(fields ...field.RelationField) ISysLabelsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysLabelsDo) FirstOrInit() (*model.SysLabels, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLabels), nil
	}
}

func (s sysLabelsDo) FirstOrCreate() (*model.SysLabels, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLabels), nil
	}
}

func (s sysLabelsDo) FindByPage(offset int, limit int) (result []*model.SysLabels, count int64, err error) {
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

func (s sysLabelsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysLabelsDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysLabelsDo) Delete(models ...*model.SysLabels) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysLabelsDo) withDO(do gen.Dao) *sysLabelsDo {
	s.DO = *do.(*gen.DO)
	return s
}
