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

func newUsersPermissions(db *gorm.DB, opts ...gen.DOOption) usersPermissions {
	_usersPermissions := usersPermissions{}

	_usersPermissions.usersPermissionsDo.UseDB(db, opts...)
	_usersPermissions.usersPermissionsDo.UseModel(&model.UsersPermissions{})

	tableName := _usersPermissions.usersPermissionsDo.TableName()
	_usersPermissions.ALL = field.NewAsterisk(tableName)
	_usersPermissions.ID = field.NewInt32(tableName, "id")
	_usersPermissions.UserID = field.NewInt32(tableName, "user_id")
	_usersPermissions.PermissionID = field.NewInt32(tableName, "permission_id")
	_usersPermissions.CreateTime = field.NewTime(tableName, "create_time")
	_usersPermissions.ModifyTime = field.NewTime(tableName, "modify_time")
	_usersPermissions.Comments = field.NewString(tableName, "comments")

	_usersPermissions.fillFieldMap()

	return _usersPermissions
}

type usersPermissions struct {
	usersPermissionsDo

	ALL          field.Asterisk
	ID           field.Int32
	UserID       field.Int32
	PermissionID field.Int32
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (u usersPermissions) Table(newTableName string) *usersPermissions {
	u.usersPermissionsDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u usersPermissions) As(alias string) *usersPermissions {
	u.usersPermissionsDo.DO = *(u.usersPermissionsDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *usersPermissions) updateTableName(table string) *usersPermissions {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.UserID = field.NewInt32(table, "user_id")
	u.PermissionID = field.NewInt32(table, "permission_id")
	u.CreateTime = field.NewTime(table, "create_time")
	u.ModifyTime = field.NewTime(table, "modify_time")
	u.Comments = field.NewString(table, "comments")

	u.fillFieldMap()

	return u
}

func (u *usersPermissions) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *usersPermissions) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["permission_id"] = u.PermissionID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["modify_time"] = u.ModifyTime
	u.fieldMap["comments"] = u.Comments
}

func (u usersPermissions) clone(db *gorm.DB) usersPermissions {
	u.usersPermissionsDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u usersPermissions) replaceDB(db *gorm.DB) usersPermissions {
	u.usersPermissionsDo.ReplaceDB(db)
	return u
}

type usersPermissionsDo struct{ gen.DO }

type IUsersPermissionsDo interface {
	gen.SubQuery
	Debug() IUsersPermissionsDo
	WithContext(ctx context.Context) IUsersPermissionsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUsersPermissionsDo
	WriteDB() IUsersPermissionsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUsersPermissionsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUsersPermissionsDo
	Not(conds ...gen.Condition) IUsersPermissionsDo
	Or(conds ...gen.Condition) IUsersPermissionsDo
	Select(conds ...field.Expr) IUsersPermissionsDo
	Where(conds ...gen.Condition) IUsersPermissionsDo
	Order(conds ...field.Expr) IUsersPermissionsDo
	Distinct(cols ...field.Expr) IUsersPermissionsDo
	Omit(cols ...field.Expr) IUsersPermissionsDo
	Join(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo
	Group(cols ...field.Expr) IUsersPermissionsDo
	Having(conds ...gen.Condition) IUsersPermissionsDo
	Limit(limit int) IUsersPermissionsDo
	Offset(offset int) IUsersPermissionsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersPermissionsDo
	Unscoped() IUsersPermissionsDo
	Create(values ...*model.UsersPermissions) error
	CreateInBatches(values []*model.UsersPermissions, batchSize int) error
	Save(values ...*model.UsersPermissions) error
	First() (*model.UsersPermissions, error)
	Take() (*model.UsersPermissions, error)
	Last() (*model.UsersPermissions, error)
	Find() ([]*model.UsersPermissions, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UsersPermissions, err error)
	FindInBatches(result *[]*model.UsersPermissions, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UsersPermissions) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUsersPermissionsDo
	Assign(attrs ...field.AssignExpr) IUsersPermissionsDo
	Joins(fields ...field.RelationField) IUsersPermissionsDo
	Preload(fields ...field.RelationField) IUsersPermissionsDo
	FirstOrInit() (*model.UsersPermissions, error)
	FirstOrCreate() (*model.UsersPermissions, error)
	FindByPage(offset int, limit int) (result []*model.UsersPermissions, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUsersPermissionsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u usersPermissionsDo) Debug() IUsersPermissionsDo {
	return u.withDO(u.DO.Debug())
}

func (u usersPermissionsDo) WithContext(ctx context.Context) IUsersPermissionsDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u usersPermissionsDo) ReadDB() IUsersPermissionsDo {
	return u.Clauses(dbresolver.Read)
}

func (u usersPermissionsDo) WriteDB() IUsersPermissionsDo {
	return u.Clauses(dbresolver.Write)
}

func (u usersPermissionsDo) Session(config *gorm.Session) IUsersPermissionsDo {
	return u.withDO(u.DO.Session(config))
}

func (u usersPermissionsDo) Clauses(conds ...clause.Expression) IUsersPermissionsDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u usersPermissionsDo) Returning(value interface{}, columns ...string) IUsersPermissionsDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u usersPermissionsDo) Not(conds ...gen.Condition) IUsersPermissionsDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u usersPermissionsDo) Or(conds ...gen.Condition) IUsersPermissionsDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u usersPermissionsDo) Select(conds ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u usersPermissionsDo) Where(conds ...gen.Condition) IUsersPermissionsDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u usersPermissionsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUsersPermissionsDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u usersPermissionsDo) Order(conds ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u usersPermissionsDo) Distinct(cols ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u usersPermissionsDo) Omit(cols ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u usersPermissionsDo) Join(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u usersPermissionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u usersPermissionsDo) RightJoin(table schema.Tabler, on ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u usersPermissionsDo) Group(cols ...field.Expr) IUsersPermissionsDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u usersPermissionsDo) Having(conds ...gen.Condition) IUsersPermissionsDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u usersPermissionsDo) Limit(limit int) IUsersPermissionsDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u usersPermissionsDo) Offset(offset int) IUsersPermissionsDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u usersPermissionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersPermissionsDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u usersPermissionsDo) Unscoped() IUsersPermissionsDo {
	return u.withDO(u.DO.Unscoped())
}

func (u usersPermissionsDo) Create(values ...*model.UsersPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u usersPermissionsDo) CreateInBatches(values []*model.UsersPermissions, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u usersPermissionsDo) Save(values ...*model.UsersPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u usersPermissionsDo) First() (*model.UsersPermissions, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsersPermissions), nil
	}
}

func (u usersPermissionsDo) Take() (*model.UsersPermissions, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsersPermissions), nil
	}
}

func (u usersPermissionsDo) Last() (*model.UsersPermissions, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsersPermissions), nil
	}
}

func (u usersPermissionsDo) Find() ([]*model.UsersPermissions, error) {
	result, err := u.DO.Find()
	return result.([]*model.UsersPermissions), err
}

func (u usersPermissionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UsersPermissions, err error) {
	buf := make([]*model.UsersPermissions, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u usersPermissionsDo) FindInBatches(result *[]*model.UsersPermissions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u usersPermissionsDo) Attrs(attrs ...field.AssignExpr) IUsersPermissionsDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u usersPermissionsDo) Assign(attrs ...field.AssignExpr) IUsersPermissionsDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u usersPermissionsDo) Joins(fields ...field.RelationField) IUsersPermissionsDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u usersPermissionsDo) Preload(fields ...field.RelationField) IUsersPermissionsDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u usersPermissionsDo) FirstOrInit() (*model.UsersPermissions, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsersPermissions), nil
	}
}

func (u usersPermissionsDo) FirstOrCreate() (*model.UsersPermissions, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsersPermissions), nil
	}
}

func (u usersPermissionsDo) FindByPage(offset int, limit int) (result []*model.UsersPermissions, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u usersPermissionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u usersPermissionsDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u usersPermissionsDo) Delete(models ...*model.UsersPermissions) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *usersPermissionsDo) withDO(do gen.Dao) *usersPermissionsDo {
	u.DO = *do.(*gen.DO)
	return u
}
