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

func newSysCode(db *gorm.DB, opts ...gen.DOOption) sysCode {
	_sysCode := sysCode{}

	_sysCode.sysCodeDo.UseDB(db, opts...)
	_sysCode.sysCodeDo.UseModel(&model.SysCode{})

	tableName := _sysCode.sysCodeDo.TableName()
	_sysCode.ALL = field.NewAsterisk(tableName)
	_sysCode.SysCode = field.NewString(tableName, "sys_code")
	_sysCode.SysCodeText = field.NewString(tableName, "sys_code_text")
	_sysCode.SysCodeType = field.NewString(tableName, "sys_code_type")
	_sysCode.PCode = field.NewString(tableName, "p_code")
	_sysCode.Status = field.NewString(tableName, "status")
	_sysCode.OrderIndex = field.NewInt32(tableName, "order_index")
	_sysCode.CreateTime = field.NewTime(tableName, "create_time")
	_sysCode.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysCode.Comments = field.NewString(tableName, "comments")
	_sysCode.Property1 = field.NewString(tableName, "property_1")
	_sysCode.SysCodeID = field.NewString(tableName, "sys_code_id")
	_sysCode.SysCodeTypeName = field.NewString(tableName, "sys_code_type_name")

	_sysCode.fillFieldMap()

	return _sysCode
}

type sysCode struct {
	sysCodeDo

	ALL         field.Asterisk
	SysCode     field.String // 编码
	SysCodeText field.String
	SysCodeType field.String
	PCode       field.String // 父sys_code，用于维护码表之间的子父基关系
	/*
		数据是否有效
		1、或空 有效
		0、无效
	*/
	Status          field.String
	OrderIndex      field.Int32  // 排序
	CreateTime      field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time   // 记录修改时间（数据库自动写入）
	Comments        field.String // 备注说明
	Property1       field.String
	SysCodeID       field.String
	SysCodeTypeName field.String

	fieldMap map[string]field.Expr
}

func (s sysCode) Table(newTableName string) *sysCode {
	s.sysCodeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysCode) As(alias string) *sysCode {
	s.sysCodeDo.DO = *(s.sysCodeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysCode) updateTableName(table string) *sysCode {
	s.ALL = field.NewAsterisk(table)
	s.SysCode = field.NewString(table, "sys_code")
	s.SysCodeText = field.NewString(table, "sys_code_text")
	s.SysCodeType = field.NewString(table, "sys_code_type")
	s.PCode = field.NewString(table, "p_code")
	s.Status = field.NewString(table, "status")
	s.OrderIndex = field.NewInt32(table, "order_index")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")
	s.Property1 = field.NewString(table, "property_1")
	s.SysCodeID = field.NewString(table, "sys_code_id")
	s.SysCodeTypeName = field.NewString(table, "sys_code_type_name")

	s.fillFieldMap()

	return s
}

func (s *sysCode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysCode) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["sys_code"] = s.SysCode
	s.fieldMap["sys_code_text"] = s.SysCodeText
	s.fieldMap["sys_code_type"] = s.SysCodeType
	s.fieldMap["p_code"] = s.PCode
	s.fieldMap["status"] = s.Status
	s.fieldMap["order_index"] = s.OrderIndex
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["property_1"] = s.Property1
	s.fieldMap["sys_code_id"] = s.SysCodeID
	s.fieldMap["sys_code_type_name"] = s.SysCodeTypeName
}

func (s sysCode) clone(db *gorm.DB) sysCode {
	s.sysCodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysCode) replaceDB(db *gorm.DB) sysCode {
	s.sysCodeDo.ReplaceDB(db)
	return s
}

type sysCodeDo struct{ gen.DO }

type ISysCodeDo interface {
	gen.SubQuery
	Debug() ISysCodeDo
	WithContext(ctx context.Context) ISysCodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysCodeDo
	WriteDB() ISysCodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysCodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysCodeDo
	Not(conds ...gen.Condition) ISysCodeDo
	Or(conds ...gen.Condition) ISysCodeDo
	Select(conds ...field.Expr) ISysCodeDo
	Where(conds ...gen.Condition) ISysCodeDo
	Order(conds ...field.Expr) ISysCodeDo
	Distinct(cols ...field.Expr) ISysCodeDo
	Omit(cols ...field.Expr) ISysCodeDo
	Join(table schema.Tabler, on ...field.Expr) ISysCodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysCodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysCodeDo
	Group(cols ...field.Expr) ISysCodeDo
	Having(conds ...gen.Condition) ISysCodeDo
	Limit(limit int) ISysCodeDo
	Offset(offset int) ISysCodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysCodeDo
	Unscoped() ISysCodeDo
	Create(values ...*model.SysCode) error
	CreateInBatches(values []*model.SysCode, batchSize int) error
	Save(values ...*model.SysCode) error
	First() (*model.SysCode, error)
	Take() (*model.SysCode, error)
	Last() (*model.SysCode, error)
	Find() ([]*model.SysCode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysCode, err error)
	FindInBatches(result *[]*model.SysCode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysCode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysCodeDo
	Assign(attrs ...field.AssignExpr) ISysCodeDo
	Joins(fields ...field.RelationField) ISysCodeDo
	Preload(fields ...field.RelationField) ISysCodeDo
	FirstOrInit() (*model.SysCode, error)
	FirstOrCreate() (*model.SysCode, error)
	FindByPage(offset int, limit int) (result []*model.SysCode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysCodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysCodeDo) Debug() ISysCodeDo {
	return s.withDO(s.DO.Debug())
}

func (s sysCodeDo) WithContext(ctx context.Context) ISysCodeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysCodeDo) ReadDB() ISysCodeDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysCodeDo) WriteDB() ISysCodeDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysCodeDo) Session(config *gorm.Session) ISysCodeDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysCodeDo) Clauses(conds ...clause.Expression) ISysCodeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysCodeDo) Returning(value interface{}, columns ...string) ISysCodeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysCodeDo) Not(conds ...gen.Condition) ISysCodeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysCodeDo) Or(conds ...gen.Condition) ISysCodeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysCodeDo) Select(conds ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysCodeDo) Where(conds ...gen.Condition) ISysCodeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysCodeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysCodeDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysCodeDo) Order(conds ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysCodeDo) Distinct(cols ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysCodeDo) Omit(cols ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysCodeDo) Join(table schema.Tabler, on ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysCodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysCodeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysCodeDo) Group(cols ...field.Expr) ISysCodeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysCodeDo) Having(conds ...gen.Condition) ISysCodeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysCodeDo) Limit(limit int) ISysCodeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysCodeDo) Offset(offset int) ISysCodeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysCodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysCodeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysCodeDo) Unscoped() ISysCodeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysCodeDo) Create(values ...*model.SysCode) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysCodeDo) CreateInBatches(values []*model.SysCode, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysCodeDo) Save(values ...*model.SysCode) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysCodeDo) First() (*model.SysCode, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCode), nil
	}
}

func (s sysCodeDo) Take() (*model.SysCode, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCode), nil
	}
}

func (s sysCodeDo) Last() (*model.SysCode, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCode), nil
	}
}

func (s sysCodeDo) Find() ([]*model.SysCode, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysCode), err
}

func (s sysCodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysCode, err error) {
	buf := make([]*model.SysCode, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysCodeDo) FindInBatches(result *[]*model.SysCode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysCodeDo) Attrs(attrs ...field.AssignExpr) ISysCodeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysCodeDo) Assign(attrs ...field.AssignExpr) ISysCodeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysCodeDo) Joins(fields ...field.RelationField) ISysCodeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysCodeDo) Preload(fields ...field.RelationField) ISysCodeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysCodeDo) FirstOrInit() (*model.SysCode, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCode), nil
	}
}

func (s sysCodeDo) FirstOrCreate() (*model.SysCode, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCode), nil
	}
}

func (s sysCodeDo) FindByPage(offset int, limit int) (result []*model.SysCode, count int64, err error) {
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

func (s sysCodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysCodeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysCodeDo) Delete(models ...*model.SysCode) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysCodeDo) withDO(do gen.Dao) *sysCodeDo {
	s.DO = *do.(*gen.DO)
	return s
}