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

func newSysTokenBusiness(db *gorm.DB, opts ...gen.DOOption) sysTokenBusiness {
	_sysTokenBusiness := sysTokenBusiness{}

	_sysTokenBusiness.sysTokenBusinessDo.UseDB(db, opts...)
	_sysTokenBusiness.sysTokenBusinessDo.UseModel(&model.SysTokenBusiness{})

	tableName := _sysTokenBusiness.sysTokenBusinessDo.TableName()
	_sysTokenBusiness.ALL = field.NewAsterisk(tableName)
	_sysTokenBusiness.SysTokenBusinessID = field.NewString(tableName, "sys_token_business_id")
	_sysTokenBusiness.SysTokenID = field.NewString(tableName, "sys_token_id")
	_sysTokenBusiness.BusinessName = field.NewString(tableName, "business_name")

	_sysTokenBusiness.fillFieldMap()

	return _sysTokenBusiness
}

type sysTokenBusiness struct {
	sysTokenBusinessDo

	ALL                field.Asterisk
	SysTokenBusinessID field.String // ID
	SysTokenID         field.String // ID
	BusinessName       field.String // 码表值business_name

	fieldMap map[string]field.Expr
}

func (s sysTokenBusiness) Table(newTableName string) *sysTokenBusiness {
	s.sysTokenBusinessDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTokenBusiness) As(alias string) *sysTokenBusiness {
	s.sysTokenBusinessDo.DO = *(s.sysTokenBusinessDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTokenBusiness) updateTableName(table string) *sysTokenBusiness {
	s.ALL = field.NewAsterisk(table)
	s.SysTokenBusinessID = field.NewString(table, "sys_token_business_id")
	s.SysTokenID = field.NewString(table, "sys_token_id")
	s.BusinessName = field.NewString(table, "business_name")

	s.fillFieldMap()

	return s
}

func (s *sysTokenBusiness) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTokenBusiness) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["sys_token_business_id"] = s.SysTokenBusinessID
	s.fieldMap["sys_token_id"] = s.SysTokenID
	s.fieldMap["business_name"] = s.BusinessName
}

func (s sysTokenBusiness) clone(db *gorm.DB) sysTokenBusiness {
	s.sysTokenBusinessDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTokenBusiness) replaceDB(db *gorm.DB) sysTokenBusiness {
	s.sysTokenBusinessDo.ReplaceDB(db)
	return s
}

type sysTokenBusinessDo struct{ gen.DO }

type ISysTokenBusinessDo interface {
	gen.SubQuery
	Debug() ISysTokenBusinessDo
	WithContext(ctx context.Context) ISysTokenBusinessDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTokenBusinessDo
	WriteDB() ISysTokenBusinessDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTokenBusinessDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTokenBusinessDo
	Not(conds ...gen.Condition) ISysTokenBusinessDo
	Or(conds ...gen.Condition) ISysTokenBusinessDo
	Select(conds ...field.Expr) ISysTokenBusinessDo
	Where(conds ...gen.Condition) ISysTokenBusinessDo
	Order(conds ...field.Expr) ISysTokenBusinessDo
	Distinct(cols ...field.Expr) ISysTokenBusinessDo
	Omit(cols ...field.Expr) ISysTokenBusinessDo
	Join(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo
	Group(cols ...field.Expr) ISysTokenBusinessDo
	Having(conds ...gen.Condition) ISysTokenBusinessDo
	Limit(limit int) ISysTokenBusinessDo
	Offset(offset int) ISysTokenBusinessDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTokenBusinessDo
	Unscoped() ISysTokenBusinessDo
	Create(values ...*model.SysTokenBusiness) error
	CreateInBatches(values []*model.SysTokenBusiness, batchSize int) error
	Save(values ...*model.SysTokenBusiness) error
	First() (*model.SysTokenBusiness, error)
	Take() (*model.SysTokenBusiness, error)
	Last() (*model.SysTokenBusiness, error)
	Find() ([]*model.SysTokenBusiness, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTokenBusiness, err error)
	FindInBatches(result *[]*model.SysTokenBusiness, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysTokenBusiness) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTokenBusinessDo
	Assign(attrs ...field.AssignExpr) ISysTokenBusinessDo
	Joins(fields ...field.RelationField) ISysTokenBusinessDo
	Preload(fields ...field.RelationField) ISysTokenBusinessDo
	FirstOrInit() (*model.SysTokenBusiness, error)
	FirstOrCreate() (*model.SysTokenBusiness, error)
	FindByPage(offset int, limit int) (result []*model.SysTokenBusiness, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTokenBusinessDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTokenBusinessDo) Debug() ISysTokenBusinessDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTokenBusinessDo) WithContext(ctx context.Context) ISysTokenBusinessDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTokenBusinessDo) ReadDB() ISysTokenBusinessDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTokenBusinessDo) WriteDB() ISysTokenBusinessDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTokenBusinessDo) Session(config *gorm.Session) ISysTokenBusinessDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTokenBusinessDo) Clauses(conds ...clause.Expression) ISysTokenBusinessDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTokenBusinessDo) Returning(value interface{}, columns ...string) ISysTokenBusinessDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTokenBusinessDo) Not(conds ...gen.Condition) ISysTokenBusinessDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTokenBusinessDo) Or(conds ...gen.Condition) ISysTokenBusinessDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTokenBusinessDo) Select(conds ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTokenBusinessDo) Where(conds ...gen.Condition) ISysTokenBusinessDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTokenBusinessDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysTokenBusinessDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysTokenBusinessDo) Order(conds ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTokenBusinessDo) Distinct(cols ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTokenBusinessDo) Omit(cols ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTokenBusinessDo) Join(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTokenBusinessDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTokenBusinessDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTokenBusinessDo) Group(cols ...field.Expr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTokenBusinessDo) Having(conds ...gen.Condition) ISysTokenBusinessDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTokenBusinessDo) Limit(limit int) ISysTokenBusinessDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTokenBusinessDo) Offset(offset int) ISysTokenBusinessDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTokenBusinessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTokenBusinessDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTokenBusinessDo) Unscoped() ISysTokenBusinessDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTokenBusinessDo) Create(values ...*model.SysTokenBusiness) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTokenBusinessDo) CreateInBatches(values []*model.SysTokenBusiness, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTokenBusinessDo) Save(values ...*model.SysTokenBusiness) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTokenBusinessDo) First() (*model.SysTokenBusiness, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenBusiness), nil
	}
}

func (s sysTokenBusinessDo) Take() (*model.SysTokenBusiness, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenBusiness), nil
	}
}

func (s sysTokenBusinessDo) Last() (*model.SysTokenBusiness, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenBusiness), nil
	}
}

func (s sysTokenBusinessDo) Find() ([]*model.SysTokenBusiness, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysTokenBusiness), err
}

func (s sysTokenBusinessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTokenBusiness, err error) {
	buf := make([]*model.SysTokenBusiness, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTokenBusinessDo) FindInBatches(result *[]*model.SysTokenBusiness, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTokenBusinessDo) Attrs(attrs ...field.AssignExpr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTokenBusinessDo) Assign(attrs ...field.AssignExpr) ISysTokenBusinessDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTokenBusinessDo) Joins(fields ...field.RelationField) ISysTokenBusinessDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTokenBusinessDo) Preload(fields ...field.RelationField) ISysTokenBusinessDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTokenBusinessDo) FirstOrInit() (*model.SysTokenBusiness, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenBusiness), nil
	}
}

func (s sysTokenBusinessDo) FirstOrCreate() (*model.SysTokenBusiness, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenBusiness), nil
	}
}

func (s sysTokenBusinessDo) FindByPage(offset int, limit int) (result []*model.SysTokenBusiness, count int64, err error) {
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

func (s sysTokenBusinessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTokenBusinessDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTokenBusinessDo) Delete(models ...*model.SysTokenBusiness) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTokenBusinessDo) withDO(do gen.Dao) *sysTokenBusinessDo {
	s.DO = *do.(*gen.DO)
	return s
}
