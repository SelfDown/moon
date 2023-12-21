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

func newSysTemplate(db *gorm.DB, opts ...gen.DOOption) sysTemplate {
	_sysTemplate := sysTemplate{}

	_sysTemplate.sysTemplateDo.UseDB(db, opts...)
	_sysTemplate.sysTemplateDo.UseModel(&model.SysTemplate{})

	tableName := _sysTemplate.sysTemplateDo.TableName()
	_sysTemplate.ALL = field.NewAsterisk(tableName)
	_sysTemplate.SysTemplateID = field.NewString(tableName, "sys_template_id")
	_sysTemplate.TemplateContent = field.NewString(tableName, "template_content")
	_sysTemplate.TemplateName = field.NewString(tableName, "template_name")
	_sysTemplate.Comments = field.NewString(tableName, "comments")
	_sysTemplate.CreateTime = field.NewTime(tableName, "create_time")
	_sysTemplate.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysTemplate.UserID = field.NewString(tableName, "user_id")

	_sysTemplate.fillFieldMap()

	return _sysTemplate
}

type sysTemplate struct {
	sysTemplateDo

	ALL             field.Asterisk
	SysTemplateID   field.String
	TemplateContent field.String
	TemplateName    field.String
	Comments        field.String
	CreateTime      field.Time // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time
	UserID          field.String

	fieldMap map[string]field.Expr
}

func (s sysTemplate) Table(newTableName string) *sysTemplate {
	s.sysTemplateDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysTemplate) As(alias string) *sysTemplate {
	s.sysTemplateDo.DO = *(s.sysTemplateDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysTemplate) updateTableName(table string) *sysTemplate {
	s.ALL = field.NewAsterisk(table)
	s.SysTemplateID = field.NewString(table, "sys_template_id")
	s.TemplateContent = field.NewString(table, "template_content")
	s.TemplateName = field.NewString(table, "template_name")
	s.Comments = field.NewString(table, "comments")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.UserID = field.NewString(table, "user_id")

	s.fillFieldMap()

	return s
}

func (s *sysTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysTemplate) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["sys_template_id"] = s.SysTemplateID
	s.fieldMap["template_content"] = s.TemplateContent
	s.fieldMap["template_name"] = s.TemplateName
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["user_id"] = s.UserID
}

func (s sysTemplate) clone(db *gorm.DB) sysTemplate {
	s.sysTemplateDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysTemplate) replaceDB(db *gorm.DB) sysTemplate {
	s.sysTemplateDo.ReplaceDB(db)
	return s
}

type sysTemplateDo struct{ gen.DO }

type ISysTemplateDo interface {
	gen.SubQuery
	Debug() ISysTemplateDo
	WithContext(ctx context.Context) ISysTemplateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysTemplateDo
	WriteDB() ISysTemplateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysTemplateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysTemplateDo
	Not(conds ...gen.Condition) ISysTemplateDo
	Or(conds ...gen.Condition) ISysTemplateDo
	Select(conds ...field.Expr) ISysTemplateDo
	Where(conds ...gen.Condition) ISysTemplateDo
	Order(conds ...field.Expr) ISysTemplateDo
	Distinct(cols ...field.Expr) ISysTemplateDo
	Omit(cols ...field.Expr) ISysTemplateDo
	Join(table schema.Tabler, on ...field.Expr) ISysTemplateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysTemplateDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysTemplateDo
	Group(cols ...field.Expr) ISysTemplateDo
	Having(conds ...gen.Condition) ISysTemplateDo
	Limit(limit int) ISysTemplateDo
	Offset(offset int) ISysTemplateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTemplateDo
	Unscoped() ISysTemplateDo
	Create(values ...*model.SysTemplate) error
	CreateInBatches(values []*model.SysTemplate, batchSize int) error
	Save(values ...*model.SysTemplate) error
	First() (*model.SysTemplate, error)
	Take() (*model.SysTemplate, error)
	Last() (*model.SysTemplate, error)
	Find() ([]*model.SysTemplate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTemplate, err error)
	FindInBatches(result *[]*model.SysTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysTemplate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysTemplateDo
	Assign(attrs ...field.AssignExpr) ISysTemplateDo
	Joins(fields ...field.RelationField) ISysTemplateDo
	Preload(fields ...field.RelationField) ISysTemplateDo
	FirstOrInit() (*model.SysTemplate, error)
	FirstOrCreate() (*model.SysTemplate, error)
	FindByPage(offset int, limit int) (result []*model.SysTemplate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysTemplateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysTemplateDo) Debug() ISysTemplateDo {
	return s.withDO(s.DO.Debug())
}

func (s sysTemplateDo) WithContext(ctx context.Context) ISysTemplateDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysTemplateDo) ReadDB() ISysTemplateDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysTemplateDo) WriteDB() ISysTemplateDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysTemplateDo) Session(config *gorm.Session) ISysTemplateDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysTemplateDo) Clauses(conds ...clause.Expression) ISysTemplateDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysTemplateDo) Returning(value interface{}, columns ...string) ISysTemplateDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysTemplateDo) Not(conds ...gen.Condition) ISysTemplateDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysTemplateDo) Or(conds ...gen.Condition) ISysTemplateDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysTemplateDo) Select(conds ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysTemplateDo) Where(conds ...gen.Condition) ISysTemplateDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysTemplateDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysTemplateDo) Order(conds ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysTemplateDo) Distinct(cols ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysTemplateDo) Omit(cols ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysTemplateDo) Join(table schema.Tabler, on ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysTemplateDo) Group(cols ...field.Expr) ISysTemplateDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysTemplateDo) Having(conds ...gen.Condition) ISysTemplateDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysTemplateDo) Limit(limit int) ISysTemplateDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysTemplateDo) Offset(offset int) ISysTemplateDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysTemplateDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysTemplateDo) Unscoped() ISysTemplateDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysTemplateDo) Create(values ...*model.SysTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysTemplateDo) CreateInBatches(values []*model.SysTemplate, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysTemplateDo) Save(values ...*model.SysTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysTemplateDo) First() (*model.SysTemplate, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTemplate), nil
	}
}

func (s sysTemplateDo) Take() (*model.SysTemplate, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTemplate), nil
	}
}

func (s sysTemplateDo) Last() (*model.SysTemplate, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTemplate), nil
	}
}

func (s sysTemplateDo) Find() ([]*model.SysTemplate, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysTemplate), err
}

func (s sysTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysTemplate, err error) {
	buf := make([]*model.SysTemplate, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysTemplateDo) FindInBatches(result *[]*model.SysTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysTemplateDo) Attrs(attrs ...field.AssignExpr) ISysTemplateDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysTemplateDo) Assign(attrs ...field.AssignExpr) ISysTemplateDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysTemplateDo) Joins(fields ...field.RelationField) ISysTemplateDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysTemplateDo) Preload(fields ...field.RelationField) ISysTemplateDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysTemplateDo) FirstOrInit() (*model.SysTemplate, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTemplate), nil
	}
}

func (s sysTemplateDo) FirstOrCreate() (*model.SysTemplate, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysTemplate), nil
	}
}

func (s sysTemplateDo) FindByPage(offset int, limit int) (result []*model.SysTemplate, count int64, err error) {
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

func (s sysTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysTemplateDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysTemplateDo) Delete(models ...*model.SysTemplate) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysTemplateDo) withDO(do gen.Dao) *sysTemplateDo {
	s.DO = *do.(*gen.DO)
	return s
}