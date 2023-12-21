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

func newSysFunc(db *gorm.DB, opts ...gen.DOOption) sysFunc {
	_sysFunc := sysFunc{}

	_sysFunc.sysFuncDo.UseDB(db, opts...)
	_sysFunc.sysFuncDo.UseModel(&model.SysFunc{})

	tableName := _sysFunc.sysFuncDo.TableName()
	_sysFunc.ALL = field.NewAsterisk(tableName)
	_sysFunc.FuncID = field.NewString(tableName, "func_id")
	_sysFunc.FuncCode = field.NewString(tableName, "func_code")
	_sysFunc.MenuID = field.NewString(tableName, "menu_id")
	_sysFunc.MenuCode = field.NewString(tableName, "menu_code")
	_sysFunc.FuncText = field.NewString(tableName, "func_text")
	_sysFunc.CreateTime = field.NewTime(tableName, "create_time")
	_sysFunc.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysFunc.Comments = field.NewString(tableName, "comments")

	_sysFunc.fillFieldMap()

	return _sysFunc
}

type sysFunc struct {
	sysFuncDo

	ALL        field.Asterisk
	FuncID     field.String
	FuncCode   field.String
	MenuID     field.String
	MenuCode   field.String
	FuncText   field.String
	CreateTime field.Time
	ModifyTime field.Time
	Comments   field.String

	fieldMap map[string]field.Expr
}

func (s sysFunc) Table(newTableName string) *sysFunc {
	s.sysFuncDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysFunc) As(alias string) *sysFunc {
	s.sysFuncDo.DO = *(s.sysFuncDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysFunc) updateTableName(table string) *sysFunc {
	s.ALL = field.NewAsterisk(table)
	s.FuncID = field.NewString(table, "func_id")
	s.FuncCode = field.NewString(table, "func_code")
	s.MenuID = field.NewString(table, "menu_id")
	s.MenuCode = field.NewString(table, "menu_code")
	s.FuncText = field.NewString(table, "func_text")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysFunc) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysFunc) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["func_id"] = s.FuncID
	s.fieldMap["func_code"] = s.FuncCode
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["menu_code"] = s.MenuCode
	s.fieldMap["func_text"] = s.FuncText
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysFunc) clone(db *gorm.DB) sysFunc {
	s.sysFuncDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysFunc) replaceDB(db *gorm.DB) sysFunc {
	s.sysFuncDo.ReplaceDB(db)
	return s
}

type sysFuncDo struct{ gen.DO }

type ISysFuncDo interface {
	gen.SubQuery
	Debug() ISysFuncDo
	WithContext(ctx context.Context) ISysFuncDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysFuncDo
	WriteDB() ISysFuncDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysFuncDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysFuncDo
	Not(conds ...gen.Condition) ISysFuncDo
	Or(conds ...gen.Condition) ISysFuncDo
	Select(conds ...field.Expr) ISysFuncDo
	Where(conds ...gen.Condition) ISysFuncDo
	Order(conds ...field.Expr) ISysFuncDo
	Distinct(cols ...field.Expr) ISysFuncDo
	Omit(cols ...field.Expr) ISysFuncDo
	Join(table schema.Tabler, on ...field.Expr) ISysFuncDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysFuncDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysFuncDo
	Group(cols ...field.Expr) ISysFuncDo
	Having(conds ...gen.Condition) ISysFuncDo
	Limit(limit int) ISysFuncDo
	Offset(offset int) ISysFuncDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysFuncDo
	Unscoped() ISysFuncDo
	Create(values ...*model.SysFunc) error
	CreateInBatches(values []*model.SysFunc, batchSize int) error
	Save(values ...*model.SysFunc) error
	First() (*model.SysFunc, error)
	Take() (*model.SysFunc, error)
	Last() (*model.SysFunc, error)
	Find() ([]*model.SysFunc, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysFunc, err error)
	FindInBatches(result *[]*model.SysFunc, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysFunc) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysFuncDo
	Assign(attrs ...field.AssignExpr) ISysFuncDo
	Joins(fields ...field.RelationField) ISysFuncDo
	Preload(fields ...field.RelationField) ISysFuncDo
	FirstOrInit() (*model.SysFunc, error)
	FirstOrCreate() (*model.SysFunc, error)
	FindByPage(offset int, limit int) (result []*model.SysFunc, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysFuncDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysFuncDo) Debug() ISysFuncDo {
	return s.withDO(s.DO.Debug())
}

func (s sysFuncDo) WithContext(ctx context.Context) ISysFuncDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysFuncDo) ReadDB() ISysFuncDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysFuncDo) WriteDB() ISysFuncDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysFuncDo) Session(config *gorm.Session) ISysFuncDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysFuncDo) Clauses(conds ...clause.Expression) ISysFuncDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysFuncDo) Returning(value interface{}, columns ...string) ISysFuncDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysFuncDo) Not(conds ...gen.Condition) ISysFuncDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysFuncDo) Or(conds ...gen.Condition) ISysFuncDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysFuncDo) Select(conds ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysFuncDo) Where(conds ...gen.Condition) ISysFuncDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysFuncDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysFuncDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysFuncDo) Order(conds ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysFuncDo) Distinct(cols ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysFuncDo) Omit(cols ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysFuncDo) Join(table schema.Tabler, on ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysFuncDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysFuncDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysFuncDo) Group(cols ...field.Expr) ISysFuncDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysFuncDo) Having(conds ...gen.Condition) ISysFuncDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysFuncDo) Limit(limit int) ISysFuncDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysFuncDo) Offset(offset int) ISysFuncDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysFuncDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysFuncDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysFuncDo) Unscoped() ISysFuncDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysFuncDo) Create(values ...*model.SysFunc) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysFuncDo) CreateInBatches(values []*model.SysFunc, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysFuncDo) Save(values ...*model.SysFunc) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysFuncDo) First() (*model.SysFunc, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFunc), nil
	}
}

func (s sysFuncDo) Take() (*model.SysFunc, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFunc), nil
	}
}

func (s sysFuncDo) Last() (*model.SysFunc, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFunc), nil
	}
}

func (s sysFuncDo) Find() ([]*model.SysFunc, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysFunc), err
}

func (s sysFuncDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysFunc, err error) {
	buf := make([]*model.SysFunc, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysFuncDo) FindInBatches(result *[]*model.SysFunc, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysFuncDo) Attrs(attrs ...field.AssignExpr) ISysFuncDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysFuncDo) Assign(attrs ...field.AssignExpr) ISysFuncDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysFuncDo) Joins(fields ...field.RelationField) ISysFuncDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysFuncDo) Preload(fields ...field.RelationField) ISysFuncDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysFuncDo) FirstOrInit() (*model.SysFunc, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFunc), nil
	}
}

func (s sysFuncDo) FirstOrCreate() (*model.SysFunc, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFunc), nil
	}
}

func (s sysFuncDo) FindByPage(offset int, limit int) (result []*model.SysFunc, count int64, err error) {
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

func (s sysFuncDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysFuncDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysFuncDo) Delete(models ...*model.SysFunc) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysFuncDo) withDO(do gen.Dao) *sysFuncDo {
	s.DO = *do.(*gen.DO)
	return s
}