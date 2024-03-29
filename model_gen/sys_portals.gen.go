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

func newSysPortals(db *gorm.DB, opts ...gen.DOOption) sysPortals {
	_sysPortals := sysPortals{}

	_sysPortals.sysPortalsDo.UseDB(db, opts...)
	_sysPortals.sysPortalsDo.UseModel(&model.SysPortals{})

	tableName := _sysPortals.sysPortalsDo.TableName()
	_sysPortals.ALL = field.NewAsterisk(tableName)
	_sysPortals.SysPortalID = field.NewInt32(tableName, "sys_portal_id")
	_sysPortals.PortalURL = field.NewString(tableName, "portal_url")
	_sysPortals.PortalIcon = field.NewString(tableName, "portal_icon")
	_sysPortals.PortalName = field.NewString(tableName, "portal_name")
	_sysPortals.ProtalDesc = field.NewString(tableName, "protal_desc")
	_sysPortals.SysProjectID = field.NewString(tableName, "sys_project_id")
	_sysPortals.PortalLoginUser = field.NewString(tableName, "portal_login_user")
	_sysPortals.PortalLoginPwd = field.NewString(tableName, "portal_login_pwd")
	_sysPortals.SysPortalTypeID = field.NewString(tableName, "sys_portal_type_id")
	_sysPortals.CreateTime = field.NewTime(tableName, "create_time")
	_sysPortals.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysPortals.Comments = field.NewString(tableName, "comments")

	_sysPortals.fillFieldMap()

	return _sysPortals
}

type sysPortals struct {
	sysPortalsDo

	ALL             field.Asterisk
	SysPortalID     field.Int32  // UUID
	PortalURL       field.String // 前台地址
	PortalIcon      field.String // 要求用 class 类
	PortalName      field.String // portal名称
	ProtalDesc      field.String // portal描述
	SysProjectID    field.String // 与  sys_project_id 关联，描述portal所属项目
	PortalLoginUser field.String // 站点登陆用户名
	PortalLoginPwd  field.String // 站点登陆密码
	SysPortalTypeID field.String // 链接所属类型
	CreateTime      field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time   // 记录修改时间（数据库自动写入）
	Comments        field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysPortals) Table(newTableName string) *sysPortals {
	s.sysPortalsDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysPortals) As(alias string) *sysPortals {
	s.sysPortalsDo.DO = *(s.sysPortalsDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysPortals) updateTableName(table string) *sysPortals {
	s.ALL = field.NewAsterisk(table)
	s.SysPortalID = field.NewInt32(table, "sys_portal_id")
	s.PortalURL = field.NewString(table, "portal_url")
	s.PortalIcon = field.NewString(table, "portal_icon")
	s.PortalName = field.NewString(table, "portal_name")
	s.ProtalDesc = field.NewString(table, "protal_desc")
	s.SysProjectID = field.NewString(table, "sys_project_id")
	s.PortalLoginUser = field.NewString(table, "portal_login_user")
	s.PortalLoginPwd = field.NewString(table, "portal_login_pwd")
	s.SysPortalTypeID = field.NewString(table, "sys_portal_type_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysPortals) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysPortals) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["sys_portal_id"] = s.SysPortalID
	s.fieldMap["portal_url"] = s.PortalURL
	s.fieldMap["portal_icon"] = s.PortalIcon
	s.fieldMap["portal_name"] = s.PortalName
	s.fieldMap["protal_desc"] = s.ProtalDesc
	s.fieldMap["sys_project_id"] = s.SysProjectID
	s.fieldMap["portal_login_user"] = s.PortalLoginUser
	s.fieldMap["portal_login_pwd"] = s.PortalLoginPwd
	s.fieldMap["sys_portal_type_id"] = s.SysPortalTypeID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysPortals) clone(db *gorm.DB) sysPortals {
	s.sysPortalsDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysPortals) replaceDB(db *gorm.DB) sysPortals {
	s.sysPortalsDo.ReplaceDB(db)
	return s
}

type sysPortalsDo struct{ gen.DO }

type ISysPortalsDo interface {
	gen.SubQuery
	Debug() ISysPortalsDo
	WithContext(ctx context.Context) ISysPortalsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysPortalsDo
	WriteDB() ISysPortalsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysPortalsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysPortalsDo
	Not(conds ...gen.Condition) ISysPortalsDo
	Or(conds ...gen.Condition) ISysPortalsDo
	Select(conds ...field.Expr) ISysPortalsDo
	Where(conds ...gen.Condition) ISysPortalsDo
	Order(conds ...field.Expr) ISysPortalsDo
	Distinct(cols ...field.Expr) ISysPortalsDo
	Omit(cols ...field.Expr) ISysPortalsDo
	Join(table schema.Tabler, on ...field.Expr) ISysPortalsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysPortalsDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysPortalsDo
	Group(cols ...field.Expr) ISysPortalsDo
	Having(conds ...gen.Condition) ISysPortalsDo
	Limit(limit int) ISysPortalsDo
	Offset(offset int) ISysPortalsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysPortalsDo
	Unscoped() ISysPortalsDo
	Create(values ...*model.SysPortals) error
	CreateInBatches(values []*model.SysPortals, batchSize int) error
	Save(values ...*model.SysPortals) error
	First() (*model.SysPortals, error)
	Take() (*model.SysPortals, error)
	Last() (*model.SysPortals, error)
	Find() ([]*model.SysPortals, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysPortals, err error)
	FindInBatches(result *[]*model.SysPortals, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysPortals) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysPortalsDo
	Assign(attrs ...field.AssignExpr) ISysPortalsDo
	Joins(fields ...field.RelationField) ISysPortalsDo
	Preload(fields ...field.RelationField) ISysPortalsDo
	FirstOrInit() (*model.SysPortals, error)
	FirstOrCreate() (*model.SysPortals, error)
	FindByPage(offset int, limit int) (result []*model.SysPortals, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysPortalsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysPortalsDo) Debug() ISysPortalsDo {
	return s.withDO(s.DO.Debug())
}

func (s sysPortalsDo) WithContext(ctx context.Context) ISysPortalsDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysPortalsDo) ReadDB() ISysPortalsDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysPortalsDo) WriteDB() ISysPortalsDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysPortalsDo) Session(config *gorm.Session) ISysPortalsDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysPortalsDo) Clauses(conds ...clause.Expression) ISysPortalsDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysPortalsDo) Returning(value interface{}, columns ...string) ISysPortalsDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysPortalsDo) Not(conds ...gen.Condition) ISysPortalsDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysPortalsDo) Or(conds ...gen.Condition) ISysPortalsDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysPortalsDo) Select(conds ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysPortalsDo) Where(conds ...gen.Condition) ISysPortalsDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysPortalsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysPortalsDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysPortalsDo) Order(conds ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysPortalsDo) Distinct(cols ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysPortalsDo) Omit(cols ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysPortalsDo) Join(table schema.Tabler, on ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysPortalsDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysPortalsDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysPortalsDo) Group(cols ...field.Expr) ISysPortalsDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysPortalsDo) Having(conds ...gen.Condition) ISysPortalsDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysPortalsDo) Limit(limit int) ISysPortalsDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysPortalsDo) Offset(offset int) ISysPortalsDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysPortalsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysPortalsDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysPortalsDo) Unscoped() ISysPortalsDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysPortalsDo) Create(values ...*model.SysPortals) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysPortalsDo) CreateInBatches(values []*model.SysPortals, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysPortalsDo) Save(values ...*model.SysPortals) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysPortalsDo) First() (*model.SysPortals, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPortals), nil
	}
}

func (s sysPortalsDo) Take() (*model.SysPortals, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPortals), nil
	}
}

func (s sysPortalsDo) Last() (*model.SysPortals, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPortals), nil
	}
}

func (s sysPortalsDo) Find() ([]*model.SysPortals, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysPortals), err
}

func (s sysPortalsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysPortals, err error) {
	buf := make([]*model.SysPortals, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysPortalsDo) FindInBatches(result *[]*model.SysPortals, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysPortalsDo) Attrs(attrs ...field.AssignExpr) ISysPortalsDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysPortalsDo) Assign(attrs ...field.AssignExpr) ISysPortalsDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysPortalsDo) Joins(fields ...field.RelationField) ISysPortalsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysPortalsDo) Preload(fields ...field.RelationField) ISysPortalsDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysPortalsDo) FirstOrInit() (*model.SysPortals, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPortals), nil
	}
}

func (s sysPortalsDo) FirstOrCreate() (*model.SysPortals, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPortals), nil
	}
}

func (s sysPortalsDo) FindByPage(offset int, limit int) (result []*model.SysPortals, count int64, err error) {
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

func (s sysPortalsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysPortalsDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysPortalsDo) Delete(models ...*model.SysPortals) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysPortalsDo) withDO(do gen.Dao) *sysPortalsDo {
	s.DO = *do.(*gen.DO)
	return s
}
