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

func newSysParamCopy1(db *gorm.DB, opts ...gen.DOOption) sysParamCopy1 {
	_sysParamCopy1 := sysParamCopy1{}

	_sysParamCopy1.sysParamCopy1Do.UseDB(db, opts...)
	_sysParamCopy1.sysParamCopy1Do.UseModel(&model.SysParamCopy1{})

	tableName := _sysParamCopy1.sysParamCopy1Do.TableName()
	_sysParamCopy1.ALL = field.NewAsterisk(tableName)
	_sysParamCopy1.SysParamID = field.NewString(tableName, "sys_param_id")
	_sysParamCopy1.ParamName = field.NewString(tableName, "param_name")
	_sysParamCopy1.ParamValue = field.NewString(tableName, "param_value")
	_sysParamCopy1.Notes = field.NewString(tableName, "notes")
	_sysParamCopy1.HospitalCode = field.NewString(tableName, "hospital_code")
	_sysParamCopy1.SysProjectID = field.NewInt32(tableName, "sys_project_id")
	_sysParamCopy1.CreateTime = field.NewTime(tableName, "create_time")
	_sysParamCopy1.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysParamCopy1.Comments = field.NewString(tableName, "comments")

	_sysParamCopy1.fillFieldMap()

	return _sysParamCopy1
}

type sysParamCopy1 struct {
	sysParamCopy1Do

	ALL          field.Asterisk
	SysParamID   field.String // 主键ID
	ParamName    field.String // 参数名称
	ParamValue   field.String // 参数值
	Notes        field.String // 备注
	HospitalCode field.String // 院区编码
	SysProjectID field.Int32  // 合作项目
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysParamCopy1) Table(newTableName string) *sysParamCopy1 {
	s.sysParamCopy1Do.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysParamCopy1) As(alias string) *sysParamCopy1 {
	s.sysParamCopy1Do.DO = *(s.sysParamCopy1Do.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysParamCopy1) updateTableName(table string) *sysParamCopy1 {
	s.ALL = field.NewAsterisk(table)
	s.SysParamID = field.NewString(table, "sys_param_id")
	s.ParamName = field.NewString(table, "param_name")
	s.ParamValue = field.NewString(table, "param_value")
	s.Notes = field.NewString(table, "notes")
	s.HospitalCode = field.NewString(table, "hospital_code")
	s.SysProjectID = field.NewInt32(table, "sys_project_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysParamCopy1) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysParamCopy1) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["sys_param_id"] = s.SysParamID
	s.fieldMap["param_name"] = s.ParamName
	s.fieldMap["param_value"] = s.ParamValue
	s.fieldMap["notes"] = s.Notes
	s.fieldMap["hospital_code"] = s.HospitalCode
	s.fieldMap["sys_project_id"] = s.SysProjectID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysParamCopy1) clone(db *gorm.DB) sysParamCopy1 {
	s.sysParamCopy1Do.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysParamCopy1) replaceDB(db *gorm.DB) sysParamCopy1 {
	s.sysParamCopy1Do.ReplaceDB(db)
	return s
}

type sysParamCopy1Do struct{ gen.DO }

type ISysParamCopy1Do interface {
	gen.SubQuery
	Debug() ISysParamCopy1Do
	WithContext(ctx context.Context) ISysParamCopy1Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysParamCopy1Do
	WriteDB() ISysParamCopy1Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysParamCopy1Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysParamCopy1Do
	Not(conds ...gen.Condition) ISysParamCopy1Do
	Or(conds ...gen.Condition) ISysParamCopy1Do
	Select(conds ...field.Expr) ISysParamCopy1Do
	Where(conds ...gen.Condition) ISysParamCopy1Do
	Order(conds ...field.Expr) ISysParamCopy1Do
	Distinct(cols ...field.Expr) ISysParamCopy1Do
	Omit(cols ...field.Expr) ISysParamCopy1Do
	Join(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do
	RightJoin(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do
	Group(cols ...field.Expr) ISysParamCopy1Do
	Having(conds ...gen.Condition) ISysParamCopy1Do
	Limit(limit int) ISysParamCopy1Do
	Offset(offset int) ISysParamCopy1Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysParamCopy1Do
	Unscoped() ISysParamCopy1Do
	Create(values ...*model.SysParamCopy1) error
	CreateInBatches(values []*model.SysParamCopy1, batchSize int) error
	Save(values ...*model.SysParamCopy1) error
	First() (*model.SysParamCopy1, error)
	Take() (*model.SysParamCopy1, error)
	Last() (*model.SysParamCopy1, error)
	Find() ([]*model.SysParamCopy1, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysParamCopy1, err error)
	FindInBatches(result *[]*model.SysParamCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysParamCopy1) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysParamCopy1Do
	Assign(attrs ...field.AssignExpr) ISysParamCopy1Do
	Joins(fields ...field.RelationField) ISysParamCopy1Do
	Preload(fields ...field.RelationField) ISysParamCopy1Do
	FirstOrInit() (*model.SysParamCopy1, error)
	FirstOrCreate() (*model.SysParamCopy1, error)
	FindByPage(offset int, limit int) (result []*model.SysParamCopy1, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysParamCopy1Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysParamCopy1Do) Debug() ISysParamCopy1Do {
	return s.withDO(s.DO.Debug())
}

func (s sysParamCopy1Do) WithContext(ctx context.Context) ISysParamCopy1Do {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysParamCopy1Do) ReadDB() ISysParamCopy1Do {
	return s.Clauses(dbresolver.Read)
}

func (s sysParamCopy1Do) WriteDB() ISysParamCopy1Do {
	return s.Clauses(dbresolver.Write)
}

func (s sysParamCopy1Do) Session(config *gorm.Session) ISysParamCopy1Do {
	return s.withDO(s.DO.Session(config))
}

func (s sysParamCopy1Do) Clauses(conds ...clause.Expression) ISysParamCopy1Do {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysParamCopy1Do) Returning(value interface{}, columns ...string) ISysParamCopy1Do {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysParamCopy1Do) Not(conds ...gen.Condition) ISysParamCopy1Do {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysParamCopy1Do) Or(conds ...gen.Condition) ISysParamCopy1Do {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysParamCopy1Do) Select(conds ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysParamCopy1Do) Where(conds ...gen.Condition) ISysParamCopy1Do {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysParamCopy1Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysParamCopy1Do {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysParamCopy1Do) Order(conds ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysParamCopy1Do) Distinct(cols ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysParamCopy1Do) Omit(cols ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysParamCopy1Do) Join(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysParamCopy1Do) LeftJoin(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysParamCopy1Do) RightJoin(table schema.Tabler, on ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysParamCopy1Do) Group(cols ...field.Expr) ISysParamCopy1Do {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysParamCopy1Do) Having(conds ...gen.Condition) ISysParamCopy1Do {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysParamCopy1Do) Limit(limit int) ISysParamCopy1Do {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysParamCopy1Do) Offset(offset int) ISysParamCopy1Do {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysParamCopy1Do) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysParamCopy1Do {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysParamCopy1Do) Unscoped() ISysParamCopy1Do {
	return s.withDO(s.DO.Unscoped())
}

func (s sysParamCopy1Do) Create(values ...*model.SysParamCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysParamCopy1Do) CreateInBatches(values []*model.SysParamCopy1, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysParamCopy1Do) Save(values ...*model.SysParamCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysParamCopy1Do) First() (*model.SysParamCopy1, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy1), nil
	}
}

func (s sysParamCopy1Do) Take() (*model.SysParamCopy1, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy1), nil
	}
}

func (s sysParamCopy1Do) Last() (*model.SysParamCopy1, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy1), nil
	}
}

func (s sysParamCopy1Do) Find() ([]*model.SysParamCopy1, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysParamCopy1), err
}

func (s sysParamCopy1Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysParamCopy1, err error) {
	buf := make([]*model.SysParamCopy1, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysParamCopy1Do) FindInBatches(result *[]*model.SysParamCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysParamCopy1Do) Attrs(attrs ...field.AssignExpr) ISysParamCopy1Do {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysParamCopy1Do) Assign(attrs ...field.AssignExpr) ISysParamCopy1Do {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysParamCopy1Do) Joins(fields ...field.RelationField) ISysParamCopy1Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysParamCopy1Do) Preload(fields ...field.RelationField) ISysParamCopy1Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysParamCopy1Do) FirstOrInit() (*model.SysParamCopy1, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy1), nil
	}
}

func (s sysParamCopy1Do) FirstOrCreate() (*model.SysParamCopy1, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy1), nil
	}
}

func (s sysParamCopy1Do) FindByPage(offset int, limit int) (result []*model.SysParamCopy1, count int64, err error) {
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

func (s sysParamCopy1Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysParamCopy1Do) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysParamCopy1Do) Delete(models ...*model.SysParamCopy1) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysParamCopy1Do) withDO(do gen.Dao) *sysParamCopy1Do {
	s.DO = *do.(*gen.DO)
	return s
}
