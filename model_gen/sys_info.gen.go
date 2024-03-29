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

func newSysInfo(db *gorm.DB, opts ...gen.DOOption) sysInfo {
	_sysInfo := sysInfo{}

	_sysInfo.sysInfoDo.UseDB(db, opts...)
	_sysInfo.sysInfoDo.UseModel(&model.SysInfo{})

	tableName := _sysInfo.sysInfoDo.TableName()
	_sysInfo.ALL = field.NewAsterisk(tableName)
	_sysInfo.SysInfoID = field.NewInt32(tableName, "sys_info_id")
	_sysInfo.SysInfoKey = field.NewString(tableName, "sys_info_key")
	_sysInfo.SysInfoValue = field.NewString(tableName, "sys_info_value")
	_sysInfo.CreateTime = field.NewTime(tableName, "create_time")
	_sysInfo.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysInfo.Comments = field.NewString(tableName, "comments")

	_sysInfo.fillFieldMap()

	return _sysInfo
}

type sysInfo struct {
	sysInfoDo

	ALL          field.Asterisk
	SysInfoID    field.Int32
	SysInfoKey   field.String
	SysInfoValue field.String
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysInfo) Table(newTableName string) *sysInfo {
	s.sysInfoDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysInfo) As(alias string) *sysInfo {
	s.sysInfoDo.DO = *(s.sysInfoDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysInfo) updateTableName(table string) *sysInfo {
	s.ALL = field.NewAsterisk(table)
	s.SysInfoID = field.NewInt32(table, "sys_info_id")
	s.SysInfoKey = field.NewString(table, "sys_info_key")
	s.SysInfoValue = field.NewString(table, "sys_info_value")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysInfo) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["sys_info_id"] = s.SysInfoID
	s.fieldMap["sys_info_key"] = s.SysInfoKey
	s.fieldMap["sys_info_value"] = s.SysInfoValue
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysInfo) clone(db *gorm.DB) sysInfo {
	s.sysInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysInfo) replaceDB(db *gorm.DB) sysInfo {
	s.sysInfoDo.ReplaceDB(db)
	return s
}

type sysInfoDo struct{ gen.DO }

type ISysInfoDo interface {
	gen.SubQuery
	Debug() ISysInfoDo
	WithContext(ctx context.Context) ISysInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysInfoDo
	WriteDB() ISysInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysInfoDo
	Not(conds ...gen.Condition) ISysInfoDo
	Or(conds ...gen.Condition) ISysInfoDo
	Select(conds ...field.Expr) ISysInfoDo
	Where(conds ...gen.Condition) ISysInfoDo
	Order(conds ...field.Expr) ISysInfoDo
	Distinct(cols ...field.Expr) ISysInfoDo
	Omit(cols ...field.Expr) ISysInfoDo
	Join(table schema.Tabler, on ...field.Expr) ISysInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysInfoDo
	Group(cols ...field.Expr) ISysInfoDo
	Having(conds ...gen.Condition) ISysInfoDo
	Limit(limit int) ISysInfoDo
	Offset(offset int) ISysInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysInfoDo
	Unscoped() ISysInfoDo
	Create(values ...*model.SysInfo) error
	CreateInBatches(values []*model.SysInfo, batchSize int) error
	Save(values ...*model.SysInfo) error
	First() (*model.SysInfo, error)
	Take() (*model.SysInfo, error)
	Last() (*model.SysInfo, error)
	Find() ([]*model.SysInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysInfo, err error)
	FindInBatches(result *[]*model.SysInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysInfoDo
	Assign(attrs ...field.AssignExpr) ISysInfoDo
	Joins(fields ...field.RelationField) ISysInfoDo
	Preload(fields ...field.RelationField) ISysInfoDo
	FirstOrInit() (*model.SysInfo, error)
	FirstOrCreate() (*model.SysInfo, error)
	FindByPage(offset int, limit int) (result []*model.SysInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysInfoDo) Debug() ISysInfoDo {
	return s.withDO(s.DO.Debug())
}

func (s sysInfoDo) WithContext(ctx context.Context) ISysInfoDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysInfoDo) ReadDB() ISysInfoDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysInfoDo) WriteDB() ISysInfoDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysInfoDo) Session(config *gorm.Session) ISysInfoDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysInfoDo) Clauses(conds ...clause.Expression) ISysInfoDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysInfoDo) Returning(value interface{}, columns ...string) ISysInfoDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysInfoDo) Not(conds ...gen.Condition) ISysInfoDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysInfoDo) Or(conds ...gen.Condition) ISysInfoDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysInfoDo) Select(conds ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysInfoDo) Where(conds ...gen.Condition) ISysInfoDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysInfoDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysInfoDo) Order(conds ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysInfoDo) Distinct(cols ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysInfoDo) Omit(cols ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysInfoDo) Join(table schema.Tabler, on ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysInfoDo) Group(cols ...field.Expr) ISysInfoDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysInfoDo) Having(conds ...gen.Condition) ISysInfoDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysInfoDo) Limit(limit int) ISysInfoDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysInfoDo) Offset(offset int) ISysInfoDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysInfoDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysInfoDo) Unscoped() ISysInfoDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysInfoDo) Create(values ...*model.SysInfo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysInfoDo) CreateInBatches(values []*model.SysInfo, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysInfoDo) Save(values ...*model.SysInfo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysInfoDo) First() (*model.SysInfo, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysInfo), nil
	}
}

func (s sysInfoDo) Take() (*model.SysInfo, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysInfo), nil
	}
}

func (s sysInfoDo) Last() (*model.SysInfo, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysInfo), nil
	}
}

func (s sysInfoDo) Find() ([]*model.SysInfo, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysInfo), err
}

func (s sysInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysInfo, err error) {
	buf := make([]*model.SysInfo, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysInfoDo) FindInBatches(result *[]*model.SysInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysInfoDo) Attrs(attrs ...field.AssignExpr) ISysInfoDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysInfoDo) Assign(attrs ...field.AssignExpr) ISysInfoDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysInfoDo) Joins(fields ...field.RelationField) ISysInfoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysInfoDo) Preload(fields ...field.RelationField) ISysInfoDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysInfoDo) FirstOrInit() (*model.SysInfo, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysInfo), nil
	}
}

func (s sysInfoDo) FirstOrCreate() (*model.SysInfo, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysInfo), nil
	}
}

func (s sysInfoDo) FindByPage(offset int, limit int) (result []*model.SysInfo, count int64, err error) {
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

func (s sysInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysInfoDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysInfoDo) Delete(models ...*model.SysInfo) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysInfoDo) withDO(do gen.Dao) *sysInfoDo {
	s.DO = *do.(*gen.DO)
	return s
}
