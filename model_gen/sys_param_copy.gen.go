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

func newSysParamCopy(db *gorm.DB, opts ...gen.DOOption) sysParamCopy {
	_sysParamCopy := sysParamCopy{}

	_sysParamCopy.sysParamCopyDo.UseDB(db, opts...)
	_sysParamCopy.sysParamCopyDo.UseModel(&model.SysParamCopy{})

	tableName := _sysParamCopy.sysParamCopyDo.TableName()
	_sysParamCopy.ALL = field.NewAsterisk(tableName)
	_sysParamCopy.SysParamID = field.NewString(tableName, "sys_param_id")
	_sysParamCopy.ParamName = field.NewString(tableName, "param_name")
	_sysParamCopy.ParamValue = field.NewString(tableName, "param_value")
	_sysParamCopy.Notes = field.NewString(tableName, "notes")
	_sysParamCopy.HospitalCode = field.NewString(tableName, "hospital_code")
	_sysParamCopy.SysProjectID = field.NewInt32(tableName, "sys_project_id")
	_sysParamCopy.CreateTime = field.NewTime(tableName, "create_time")
	_sysParamCopy.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysParamCopy.Comments = field.NewString(tableName, "comments")

	_sysParamCopy.fillFieldMap()

	return _sysParamCopy
}

type sysParamCopy struct {
	sysParamCopyDo

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

func (s sysParamCopy) Table(newTableName string) *sysParamCopy {
	s.sysParamCopyDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysParamCopy) As(alias string) *sysParamCopy {
	s.sysParamCopyDo.DO = *(s.sysParamCopyDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysParamCopy) updateTableName(table string) *sysParamCopy {
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

func (s *sysParamCopy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysParamCopy) fillFieldMap() {
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

func (s sysParamCopy) clone(db *gorm.DB) sysParamCopy {
	s.sysParamCopyDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysParamCopy) replaceDB(db *gorm.DB) sysParamCopy {
	s.sysParamCopyDo.ReplaceDB(db)
	return s
}

type sysParamCopyDo struct{ gen.DO }

type ISysParamCopyDo interface {
	gen.SubQuery
	Debug() ISysParamCopyDo
	WithContext(ctx context.Context) ISysParamCopyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysParamCopyDo
	WriteDB() ISysParamCopyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysParamCopyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysParamCopyDo
	Not(conds ...gen.Condition) ISysParamCopyDo
	Or(conds ...gen.Condition) ISysParamCopyDo
	Select(conds ...field.Expr) ISysParamCopyDo
	Where(conds ...gen.Condition) ISysParamCopyDo
	Order(conds ...field.Expr) ISysParamCopyDo
	Distinct(cols ...field.Expr) ISysParamCopyDo
	Omit(cols ...field.Expr) ISysParamCopyDo
	Join(table schema.Tabler, on ...field.Expr) ISysParamCopyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysParamCopyDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysParamCopyDo
	Group(cols ...field.Expr) ISysParamCopyDo
	Having(conds ...gen.Condition) ISysParamCopyDo
	Limit(limit int) ISysParamCopyDo
	Offset(offset int) ISysParamCopyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysParamCopyDo
	Unscoped() ISysParamCopyDo
	Create(values ...*model.SysParamCopy) error
	CreateInBatches(values []*model.SysParamCopy, batchSize int) error
	Save(values ...*model.SysParamCopy) error
	First() (*model.SysParamCopy, error)
	Take() (*model.SysParamCopy, error)
	Last() (*model.SysParamCopy, error)
	Find() ([]*model.SysParamCopy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysParamCopy, err error)
	FindInBatches(result *[]*model.SysParamCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysParamCopy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysParamCopyDo
	Assign(attrs ...field.AssignExpr) ISysParamCopyDo
	Joins(fields ...field.RelationField) ISysParamCopyDo
	Preload(fields ...field.RelationField) ISysParamCopyDo
	FirstOrInit() (*model.SysParamCopy, error)
	FirstOrCreate() (*model.SysParamCopy, error)
	FindByPage(offset int, limit int) (result []*model.SysParamCopy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysParamCopyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysParamCopyDo) Debug() ISysParamCopyDo {
	return s.withDO(s.DO.Debug())
}

func (s sysParamCopyDo) WithContext(ctx context.Context) ISysParamCopyDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysParamCopyDo) ReadDB() ISysParamCopyDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysParamCopyDo) WriteDB() ISysParamCopyDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysParamCopyDo) Session(config *gorm.Session) ISysParamCopyDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysParamCopyDo) Clauses(conds ...clause.Expression) ISysParamCopyDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysParamCopyDo) Returning(value interface{}, columns ...string) ISysParamCopyDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysParamCopyDo) Not(conds ...gen.Condition) ISysParamCopyDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysParamCopyDo) Or(conds ...gen.Condition) ISysParamCopyDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysParamCopyDo) Select(conds ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysParamCopyDo) Where(conds ...gen.Condition) ISysParamCopyDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysParamCopyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysParamCopyDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysParamCopyDo) Order(conds ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysParamCopyDo) Distinct(cols ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysParamCopyDo) Omit(cols ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysParamCopyDo) Join(table schema.Tabler, on ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysParamCopyDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysParamCopyDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysParamCopyDo) Group(cols ...field.Expr) ISysParamCopyDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysParamCopyDo) Having(conds ...gen.Condition) ISysParamCopyDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysParamCopyDo) Limit(limit int) ISysParamCopyDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysParamCopyDo) Offset(offset int) ISysParamCopyDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysParamCopyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysParamCopyDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysParamCopyDo) Unscoped() ISysParamCopyDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysParamCopyDo) Create(values ...*model.SysParamCopy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysParamCopyDo) CreateInBatches(values []*model.SysParamCopy, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysParamCopyDo) Save(values ...*model.SysParamCopy) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysParamCopyDo) First() (*model.SysParamCopy, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy), nil
	}
}

func (s sysParamCopyDo) Take() (*model.SysParamCopy, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy), nil
	}
}

func (s sysParamCopyDo) Last() (*model.SysParamCopy, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy), nil
	}
}

func (s sysParamCopyDo) Find() ([]*model.SysParamCopy, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysParamCopy), err
}

func (s sysParamCopyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysParamCopy, err error) {
	buf := make([]*model.SysParamCopy, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysParamCopyDo) FindInBatches(result *[]*model.SysParamCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysParamCopyDo) Attrs(attrs ...field.AssignExpr) ISysParamCopyDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysParamCopyDo) Assign(attrs ...field.AssignExpr) ISysParamCopyDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysParamCopyDo) Joins(fields ...field.RelationField) ISysParamCopyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysParamCopyDo) Preload(fields ...field.RelationField) ISysParamCopyDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysParamCopyDo) FirstOrInit() (*model.SysParamCopy, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy), nil
	}
}

func (s sysParamCopyDo) FirstOrCreate() (*model.SysParamCopy, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysParamCopy), nil
	}
}

func (s sysParamCopyDo) FindByPage(offset int, limit int) (result []*model.SysParamCopy, count int64, err error) {
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

func (s sysParamCopyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysParamCopyDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysParamCopyDo) Delete(models ...*model.SysParamCopy) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysParamCopyDo) withDO(do gen.Dao) *sysParamCopyDo {
	s.DO = *do.(*gen.DO)
	return s
}
