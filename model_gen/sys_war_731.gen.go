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

func newSysWar731(db *gorm.DB, opts ...gen.DOOption) sysWar731 {
	_sysWar731 := sysWar731{}

	_sysWar731.sysWar731Do.UseDB(db, opts...)
	_sysWar731.sysWar731Do.UseModel(&model.SysWar731{})

	tableName := _sysWar731.sysWar731Do.TableName()
	_sysWar731.ALL = field.NewAsterisk(tableName)
	_sysWar731.WarID = field.NewString(tableName, "war_id")
	_sysWar731.WarGroupid = field.NewString(tableName, "war_groupid")
	_sysWar731.WarArtifactid = field.NewString(tableName, "war_artifactid")
	_sysWar731.WarName = field.NewString(tableName, "war_name")
	_sysWar731.HospitalCode = field.NewString(tableName, "hospital_code")
	_sysWar731.DeployType = field.NewString(tableName, "deploy_type")
	_sysWar731.WebServiceType = field.NewString(tableName, "web_service_type")
	_sysWar731.CreateTime = field.NewTime(tableName, "create_time")
	_sysWar731.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysWar731.Comments = field.NewString(tableName, "comments")

	_sysWar731.fillFieldMap()

	return _sysWar731
}

type sysWar731 struct {
	sysWar731Do

	ALL           field.Asterisk
	WarID         field.String // 参数名称
	WarGroupid    field.String // 参数值
	WarArtifactid field.String
	WarName       field.String
	HospitalCode  field.String
	DeployType    field.String // 0 jboss 部署 1 模版部署
	/*
		web服务类型
		0-前台 portal
		1-后台 server
	*/
	WebServiceType field.String
	CreateTime     field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime     field.Time   // 记录修改时间（数据库自动写入）
	Comments       field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysWar731) Table(newTableName string) *sysWar731 {
	s.sysWar731Do.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysWar731) As(alias string) *sysWar731 {
	s.sysWar731Do.DO = *(s.sysWar731Do.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysWar731) updateTableName(table string) *sysWar731 {
	s.ALL = field.NewAsterisk(table)
	s.WarID = field.NewString(table, "war_id")
	s.WarGroupid = field.NewString(table, "war_groupid")
	s.WarArtifactid = field.NewString(table, "war_artifactid")
	s.WarName = field.NewString(table, "war_name")
	s.HospitalCode = field.NewString(table, "hospital_code")
	s.DeployType = field.NewString(table, "deploy_type")
	s.WebServiceType = field.NewString(table, "web_service_type")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysWar731) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysWar731) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["war_id"] = s.WarID
	s.fieldMap["war_groupid"] = s.WarGroupid
	s.fieldMap["war_artifactid"] = s.WarArtifactid
	s.fieldMap["war_name"] = s.WarName
	s.fieldMap["hospital_code"] = s.HospitalCode
	s.fieldMap["deploy_type"] = s.DeployType
	s.fieldMap["web_service_type"] = s.WebServiceType
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysWar731) clone(db *gorm.DB) sysWar731 {
	s.sysWar731Do.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysWar731) replaceDB(db *gorm.DB) sysWar731 {
	s.sysWar731Do.ReplaceDB(db)
	return s
}

type sysWar731Do struct{ gen.DO }

type ISysWar731Do interface {
	gen.SubQuery
	Debug() ISysWar731Do
	WithContext(ctx context.Context) ISysWar731Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysWar731Do
	WriteDB() ISysWar731Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysWar731Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysWar731Do
	Not(conds ...gen.Condition) ISysWar731Do
	Or(conds ...gen.Condition) ISysWar731Do
	Select(conds ...field.Expr) ISysWar731Do
	Where(conds ...gen.Condition) ISysWar731Do
	Order(conds ...field.Expr) ISysWar731Do
	Distinct(cols ...field.Expr) ISysWar731Do
	Omit(cols ...field.Expr) ISysWar731Do
	Join(table schema.Tabler, on ...field.Expr) ISysWar731Do
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysWar731Do
	RightJoin(table schema.Tabler, on ...field.Expr) ISysWar731Do
	Group(cols ...field.Expr) ISysWar731Do
	Having(conds ...gen.Condition) ISysWar731Do
	Limit(limit int) ISysWar731Do
	Offset(offset int) ISysWar731Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysWar731Do
	Unscoped() ISysWar731Do
	Create(values ...*model.SysWar731) error
	CreateInBatches(values []*model.SysWar731, batchSize int) error
	Save(values ...*model.SysWar731) error
	First() (*model.SysWar731, error)
	Take() (*model.SysWar731, error)
	Last() (*model.SysWar731, error)
	Find() ([]*model.SysWar731, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysWar731, err error)
	FindInBatches(result *[]*model.SysWar731, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysWar731) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysWar731Do
	Assign(attrs ...field.AssignExpr) ISysWar731Do
	Joins(fields ...field.RelationField) ISysWar731Do
	Preload(fields ...field.RelationField) ISysWar731Do
	FirstOrInit() (*model.SysWar731, error)
	FirstOrCreate() (*model.SysWar731, error)
	FindByPage(offset int, limit int) (result []*model.SysWar731, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysWar731Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysWar731Do) Debug() ISysWar731Do {
	return s.withDO(s.DO.Debug())
}

func (s sysWar731Do) WithContext(ctx context.Context) ISysWar731Do {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysWar731Do) ReadDB() ISysWar731Do {
	return s.Clauses(dbresolver.Read)
}

func (s sysWar731Do) WriteDB() ISysWar731Do {
	return s.Clauses(dbresolver.Write)
}

func (s sysWar731Do) Session(config *gorm.Session) ISysWar731Do {
	return s.withDO(s.DO.Session(config))
}

func (s sysWar731Do) Clauses(conds ...clause.Expression) ISysWar731Do {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysWar731Do) Returning(value interface{}, columns ...string) ISysWar731Do {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysWar731Do) Not(conds ...gen.Condition) ISysWar731Do {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysWar731Do) Or(conds ...gen.Condition) ISysWar731Do {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysWar731Do) Select(conds ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysWar731Do) Where(conds ...gen.Condition) ISysWar731Do {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysWar731Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysWar731Do {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysWar731Do) Order(conds ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysWar731Do) Distinct(cols ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysWar731Do) Omit(cols ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysWar731Do) Join(table schema.Tabler, on ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysWar731Do) LeftJoin(table schema.Tabler, on ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysWar731Do) RightJoin(table schema.Tabler, on ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysWar731Do) Group(cols ...field.Expr) ISysWar731Do {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysWar731Do) Having(conds ...gen.Condition) ISysWar731Do {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysWar731Do) Limit(limit int) ISysWar731Do {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysWar731Do) Offset(offset int) ISysWar731Do {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysWar731Do) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysWar731Do {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysWar731Do) Unscoped() ISysWar731Do {
	return s.withDO(s.DO.Unscoped())
}

func (s sysWar731Do) Create(values ...*model.SysWar731) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysWar731Do) CreateInBatches(values []*model.SysWar731, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysWar731Do) Save(values ...*model.SysWar731) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysWar731Do) First() (*model.SysWar731, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysWar731), nil
	}
}

func (s sysWar731Do) Take() (*model.SysWar731, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysWar731), nil
	}
}

func (s sysWar731Do) Last() (*model.SysWar731, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysWar731), nil
	}
}

func (s sysWar731Do) Find() ([]*model.SysWar731, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysWar731), err
}

func (s sysWar731Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysWar731, err error) {
	buf := make([]*model.SysWar731, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysWar731Do) FindInBatches(result *[]*model.SysWar731, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysWar731Do) Attrs(attrs ...field.AssignExpr) ISysWar731Do {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysWar731Do) Assign(attrs ...field.AssignExpr) ISysWar731Do {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysWar731Do) Joins(fields ...field.RelationField) ISysWar731Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysWar731Do) Preload(fields ...field.RelationField) ISysWar731Do {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysWar731Do) FirstOrInit() (*model.SysWar731, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysWar731), nil
	}
}

func (s sysWar731Do) FirstOrCreate() (*model.SysWar731, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysWar731), nil
	}
}

func (s sysWar731Do) FindByPage(offset int, limit int) (result []*model.SysWar731, count int64, err error) {
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

func (s sysWar731Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysWar731Do) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysWar731Do) Delete(models ...*model.SysWar731) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysWar731Do) withDO(do gen.Dao) *sysWar731Do {
	s.DO = *do.(*gen.DO)
	return s
}
