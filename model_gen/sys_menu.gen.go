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

func newSysMenu(db *gorm.DB, opts ...gen.DOOption) sysMenu {
	_sysMenu := sysMenu{}

	_sysMenu.sysMenuDo.UseDB(db, opts...)
	_sysMenu.sysMenuDo.UseModel(&model.SysMenu{})

	tableName := _sysMenu.sysMenuDo.TableName()
	_sysMenu.ALL = field.NewAsterisk(tableName)
	_sysMenu.MenuID = field.NewString(tableName, "menu_id")
	_sysMenu.MenuPid = field.NewString(tableName, "menu_pid")
	_sysMenu.MenuText = field.NewString(tableName, "menu_text")
	_sysMenu.MenuURL = field.NewString(tableName, "menu_url")
	_sysMenu.MenuIcon = field.NewString(tableName, "menu_icon")
	_sysMenu.Statu = field.NewInt32(tableName, "statu")
	_sysMenu.Expand = field.NewString(tableName, "expand")
	_sysMenu.NewTab = field.NewString(tableName, "new_tab")
	_sysMenu.SysProjectID = field.NewInt32(tableName, "sys_project_id")
	_sysMenu.RoleAuthLevel = field.NewInt32(tableName, "role_auth_level")
	_sysMenu.MenuIconCSS = field.NewString(tableName, "menu_icon_css")
	_sysMenu.MenuDot = field.NewString(tableName, "menu_dot")
	_sysMenu.CreateTime = field.NewTime(tableName, "create_time")
	_sysMenu.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysMenu.Comments = field.NewString(tableName, "comments")

	_sysMenu.fillFieldMap()

	return _sysMenu
}

type sysMenu struct {
	sysMenuDo

	ALL           field.Asterisk
	MenuID        field.String // 菜单主键
	MenuPid       field.String // 父菜单ID
	MenuText      field.String // 菜单中文名称
	MenuURL       field.String // 菜单链接
	MenuIcon      field.String // 菜单图标
	Statu         field.Int32  // 1 启用，0禁用
	Expand        field.String // 父菜单是否默认展开
	NewTab        field.String // 是否在新的浏览器标签页打开
	SysProjectID  field.Int32  // 所属项目ID
	RoleAuthLevel field.Int32  // 角色权限
	MenuIconCSS   field.String // css图标
	/*
		菜单是否显示小红点，用于提示用户这个是功能新功能或特性
		1 显示小红点
		0 不显示小红点
	*/
	MenuDot    field.String
	CreateTime field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime field.Time   // 记录修改时间（数据库自动写入）
	Comments   field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysMenu) Table(newTableName string) *sysMenu {
	s.sysMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysMenu) As(alias string) *sysMenu {
	s.sysMenuDo.DO = *(s.sysMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysMenu) updateTableName(table string) *sysMenu {
	s.ALL = field.NewAsterisk(table)
	s.MenuID = field.NewString(table, "menu_id")
	s.MenuPid = field.NewString(table, "menu_pid")
	s.MenuText = field.NewString(table, "menu_text")
	s.MenuURL = field.NewString(table, "menu_url")
	s.MenuIcon = field.NewString(table, "menu_icon")
	s.Statu = field.NewInt32(table, "statu")
	s.Expand = field.NewString(table, "expand")
	s.NewTab = field.NewString(table, "new_tab")
	s.SysProjectID = field.NewInt32(table, "sys_project_id")
	s.RoleAuthLevel = field.NewInt32(table, "role_auth_level")
	s.MenuIconCSS = field.NewString(table, "menu_icon_css")
	s.MenuDot = field.NewString(table, "menu_dot")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["menu_pid"] = s.MenuPid
	s.fieldMap["menu_text"] = s.MenuText
	s.fieldMap["menu_url"] = s.MenuURL
	s.fieldMap["menu_icon"] = s.MenuIcon
	s.fieldMap["statu"] = s.Statu
	s.fieldMap["expand"] = s.Expand
	s.fieldMap["new_tab"] = s.NewTab
	s.fieldMap["sys_project_id"] = s.SysProjectID
	s.fieldMap["role_auth_level"] = s.RoleAuthLevel
	s.fieldMap["menu_icon_css"] = s.MenuIconCSS
	s.fieldMap["menu_dot"] = s.MenuDot
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysMenu) clone(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysMenu) replaceDB(db *gorm.DB) sysMenu {
	s.sysMenuDo.ReplaceDB(db)
	return s
}

type sysMenuDo struct{ gen.DO }

type ISysMenuDo interface {
	gen.SubQuery
	Debug() ISysMenuDo
	WithContext(ctx context.Context) ISysMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysMenuDo
	WriteDB() ISysMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysMenuDo
	Not(conds ...gen.Condition) ISysMenuDo
	Or(conds ...gen.Condition) ISysMenuDo
	Select(conds ...field.Expr) ISysMenuDo
	Where(conds ...gen.Condition) ISysMenuDo
	Order(conds ...field.Expr) ISysMenuDo
	Distinct(cols ...field.Expr) ISysMenuDo
	Omit(cols ...field.Expr) ISysMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISysMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo
	Group(cols ...field.Expr) ISysMenuDo
	Having(conds ...gen.Condition) ISysMenuDo
	Limit(limit int) ISysMenuDo
	Offset(offset int) ISysMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo
	Unscoped() ISysMenuDo
	Create(values ...*model.SysMenu) error
	CreateInBatches(values []*model.SysMenu, batchSize int) error
	Save(values ...*model.SysMenu) error
	First() (*model.SysMenu, error)
	Take() (*model.SysMenu, error)
	Last() (*model.SysMenu, error)
	Find() ([]*model.SysMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error)
	FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysMenuDo
	Assign(attrs ...field.AssignExpr) ISysMenuDo
	Joins(fields ...field.RelationField) ISysMenuDo
	Preload(fields ...field.RelationField) ISysMenuDo
	FirstOrInit() (*model.SysMenu, error)
	FirstOrCreate() (*model.SysMenu, error)
	FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysMenuDo) Debug() ISysMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s sysMenuDo) WithContext(ctx context.Context) ISysMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysMenuDo) ReadDB() ISysMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysMenuDo) WriteDB() ISysMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysMenuDo) Session(config *gorm.Session) ISysMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysMenuDo) Clauses(conds ...clause.Expression) ISysMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysMenuDo) Returning(value interface{}, columns ...string) ISysMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysMenuDo) Not(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysMenuDo) Or(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysMenuDo) Select(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysMenuDo) Where(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysMenuDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysMenuDo) Order(conds ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysMenuDo) Distinct(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysMenuDo) Omit(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysMenuDo) Join(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysMenuDo) Group(cols ...field.Expr) ISysMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysMenuDo) Having(conds ...gen.Condition) ISysMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysMenuDo) Limit(limit int) ISysMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysMenuDo) Offset(offset int) ISysMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysMenuDo) Unscoped() ISysMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysMenuDo) Create(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysMenuDo) CreateInBatches(values []*model.SysMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysMenuDo) Save(values ...*model.SysMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysMenuDo) First() (*model.SysMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Take() (*model.SysMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Last() (*model.SysMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) Find() ([]*model.SysMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysMenu), err
}

func (s sysMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysMenu, err error) {
	buf := make([]*model.SysMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysMenuDo) FindInBatches(result *[]*model.SysMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysMenuDo) Attrs(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysMenuDo) Assign(attrs ...field.AssignExpr) ISysMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysMenuDo) Joins(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysMenuDo) Preload(fields ...field.RelationField) ISysMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysMenuDo) FirstOrInit() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FirstOrCreate() (*model.SysMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysMenu), nil
	}
}

func (s sysMenuDo) FindByPage(offset int, limit int) (result []*model.SysMenu, count int64, err error) {
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

func (s sysMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysMenuDo) Delete(models ...*model.SysMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysMenuDo) withDO(do gen.Dao) *sysMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}
