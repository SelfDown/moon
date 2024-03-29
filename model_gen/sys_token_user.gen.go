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

func newSysTokenUser(db *gorm.DB, opts ...gen.DOOption) sysTokenUser {
	_sysTokenUser := sysTokenUser{}

	_sysTokenUser.sysTokenUserDo.UseDB(db, opts...)
	_sysTokenUser.sysTokenUserDo.UseModel(&model.SysTokenUser{})

	tableName := _sysTokenUser.sysTokenUserDo.TableName()
	_sysTokenUser.ALL = field.NewAsterisk(tableName)
	_sysTokenUser.SysTokenUserID = field.NewString(tableName, "sys_token_user_id")
	_sysTokenUser.SysTokenID = field.NewString(tableName, "sys_token_id")
	_sysTokenUser.Userid = field.NewString(tableName, "userid")

	_sysTokenUser.fillFieldMap()

	return _sysTokenUser
}

type sysTokenUser struct {
	sysTokenUserDo

	ALL            field.Asterisk
	SysTokenUserID field.String // ID
	SysTokenID     field.String // ID
	Userid         field.String // 项目唯一标识符

	fieldMap map[string]field.Expr
}

func (s sysTokenUser) Table(newTableName string) *sysTokenUser {
	s.sysTokenUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTokenUser) As(alias string) *sysTokenUser {
	s.sysTokenUserDo.DO = *(s.sysTokenUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTokenUser) updateTableName(table string) *sysTokenUser {
	s.ALL = field.NewAsterisk(table)
	s.SysTokenUserID = field.NewString(table, "sys_token_user_id")
	s.SysTokenID = field.NewString(table, "sys_token_id")
	s.Userid = field.NewString(table, "userid")

	s.fillFieldMap()

	return s
}

func (s *sysTokenUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTokenUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["sys_token_user_id"] = s.SysTokenUserID
	s.fieldMap["sys_token_id"] = s.SysTokenID
	s.fieldMap["userid"] = s.Userid
}

func (s sysTokenUser) clone(db *gorm.DB) sysTokenUser {
	s.sysTokenUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTokenUser) replaceDB(db *gorm.DB) sysTokenUser {
	s.sysTokenUserDo.ReplaceDB(db)
	return s
}

type sysTokenUserDo struct{ gen.DO }

type ISysTokenUserDo interface {
	gen.SubQuery
	Debug() ISysTokenUserDo
	WithContext(ctx context.Context) ISysTokenUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTokenUserDo
	WriteDB() ISysTokenUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTokenUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTokenUserDo
	Not(conds ...gen.Condition) ISysTokenUserDo
	Or(conds ...gen.Condition) ISysTokenUserDo
	Select(conds ...field.Expr) ISysTokenUserDo
	Where(conds ...gen.Condition) ISysTokenUserDo
	Order(conds ...field.Expr) ISysTokenUserDo
	Distinct(cols ...field.Expr) ISysTokenUserDo
	Omit(cols ...field.Expr) ISysTokenUserDo
	Join(table schema.Tabler, on ...field.Expr) ISysTokenUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTokenUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTokenUserDo
	Group(cols ...field.Expr) ISysTokenUserDo
	Having(conds ...gen.Condition) ISysTokenUserDo
	Limit(limit int) ISysTokenUserDo
	Offset(offset int) ISysTokenUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTokenUserDo
	Unscoped() ISysTokenUserDo
	Create(values ...*model.SysTokenUser) error
	CreateInBatches(values []*model.SysTokenUser, batchSize int) error
	Save(values ...*model.SysTokenUser) error
	First() (*model.SysTokenUser, error)
	Take() (*model.SysTokenUser, error)
	Last() (*model.SysTokenUser, error)
	Find() ([]*model.SysTokenUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTokenUser, err error)
	FindInBatches(result *[]*model.SysTokenUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysTokenUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTokenUserDo
	Assign(attrs ...field.AssignExpr) ISysTokenUserDo
	Joins(fields ...field.RelationField) ISysTokenUserDo
	Preload(fields ...field.RelationField) ISysTokenUserDo
	FirstOrInit() (*model.SysTokenUser, error)
	FirstOrCreate() (*model.SysTokenUser, error)
	FindByPage(offset int, limit int) (result []*model.SysTokenUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTokenUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTokenUserDo) Debug() ISysTokenUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTokenUserDo) WithContext(ctx context.Context) ISysTokenUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTokenUserDo) ReadDB() ISysTokenUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTokenUserDo) WriteDB() ISysTokenUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTokenUserDo) Session(config *gorm.Session) ISysTokenUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTokenUserDo) Clauses(conds ...clause.Expression) ISysTokenUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTokenUserDo) Returning(value interface{}, columns ...string) ISysTokenUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTokenUserDo) Not(conds ...gen.Condition) ISysTokenUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTokenUserDo) Or(conds ...gen.Condition) ISysTokenUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTokenUserDo) Select(conds ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTokenUserDo) Where(conds ...gen.Condition) ISysTokenUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTokenUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysTokenUserDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysTokenUserDo) Order(conds ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTokenUserDo) Distinct(cols ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTokenUserDo) Omit(cols ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTokenUserDo) Join(table schema.Tabler, on ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTokenUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTokenUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTokenUserDo) Group(cols ...field.Expr) ISysTokenUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTokenUserDo) Having(conds ...gen.Condition) ISysTokenUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTokenUserDo) Limit(limit int) ISysTokenUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTokenUserDo) Offset(offset int) ISysTokenUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTokenUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTokenUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTokenUserDo) Unscoped() ISysTokenUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTokenUserDo) Create(values ...*model.SysTokenUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTokenUserDo) CreateInBatches(values []*model.SysTokenUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTokenUserDo) Save(values ...*model.SysTokenUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTokenUserDo) First() (*model.SysTokenUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenUser), nil
	}
}

func (s sysTokenUserDo) Take() (*model.SysTokenUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenUser), nil
	}
}

func (s sysTokenUserDo) Last() (*model.SysTokenUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenUser), nil
	}
}

func (s sysTokenUserDo) Find() ([]*model.SysTokenUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysTokenUser), err
}

func (s sysTokenUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTokenUser, err error) {
	buf := make([]*model.SysTokenUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTokenUserDo) FindInBatches(result *[]*model.SysTokenUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTokenUserDo) Attrs(attrs ...field.AssignExpr) ISysTokenUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTokenUserDo) Assign(attrs ...field.AssignExpr) ISysTokenUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTokenUserDo) Joins(fields ...field.RelationField) ISysTokenUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTokenUserDo) Preload(fields ...field.RelationField) ISysTokenUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTokenUserDo) FirstOrInit() (*model.SysTokenUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenUser), nil
	}
}

func (s sysTokenUserDo) FirstOrCreate() (*model.SysTokenUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTokenUser), nil
	}
}

func (s sysTokenUserDo) FindByPage(offset int, limit int) (result []*model.SysTokenUser, count int64, err error) {
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

func (s sysTokenUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTokenUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTokenUserDo) Delete(models ...*model.SysTokenUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTokenUserDo) withDO(do gen.Dao) *sysTokenUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
