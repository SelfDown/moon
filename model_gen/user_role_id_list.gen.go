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

func newUserRoleIDList(db *gorm.DB, opts ...gen.DOOption) userRoleIDList {
	_userRoleIDList := userRoleIDList{}

	_userRoleIDList.userRoleIDListDo.UseDB(db, opts...)
	_userRoleIDList.userRoleIDListDo.UseModel(&model.UserRoleIDList{})

	tableName := _userRoleIDList.userRoleIDListDo.TableName()
	_userRoleIDList.ALL = field.NewAsterisk(tableName)
	_userRoleIDList.RoleListID = field.NewInt32(tableName, "role_list_id")
	_userRoleIDList.UserID = field.NewString(tableName, "user_id")
	_userRoleIDList.RoleID = field.NewString(tableName, "role_id")
	_userRoleIDList.CreateTime = field.NewString(tableName, "create_time")
	_userRoleIDList.ModifyTime = field.NewString(tableName, "modify_time")
	_userRoleIDList.OrderWeight = field.NewInt32(tableName, "order_weight")
	_userRoleIDList.UserGroupID = field.NewString(tableName, "user_group_id")
	_userRoleIDList.UserRoleType = field.NewString(tableName, "user_role_type")

	_userRoleIDList.fillFieldMap()

	return _userRoleIDList
}

type userRoleIDList struct {
	userRoleIDListDo

	ALL          field.Asterisk
	RoleListID   field.Int32
	UserID       field.String
	RoleID       field.String
	CreateTime   field.String
	ModifyTime   field.String
	OrderWeight  field.Int32
	UserGroupID  field.String
	UserRoleType field.String

	fieldMap map[string]field.Expr
}

func (u userRoleIDList) Table(newTableName string) *userRoleIDList {
	u.userRoleIDListDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRoleIDList) As(alias string) *userRoleIDList {
	u.userRoleIDListDo.DO = *(u.userRoleIDListDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRoleIDList) updateTableName(table string) *userRoleIDList {
	u.ALL = field.NewAsterisk(table)
	u.RoleListID = field.NewInt32(table, "role_list_id")
	u.UserID = field.NewString(table, "user_id")
	u.RoleID = field.NewString(table, "role_id")
	u.CreateTime = field.NewString(table, "create_time")
	u.ModifyTime = field.NewString(table, "modify_time")
	u.OrderWeight = field.NewInt32(table, "order_weight")
	u.UserGroupID = field.NewString(table, "user_group_id")
	u.UserRoleType = field.NewString(table, "user_role_type")

	u.fillFieldMap()

	return u
}

func (u *userRoleIDList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRoleIDList) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["role_list_id"] = u.RoleListID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["role_id"] = u.RoleID
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["modify_time"] = u.ModifyTime
	u.fieldMap["order_weight"] = u.OrderWeight
	u.fieldMap["user_group_id"] = u.UserGroupID
	u.fieldMap["user_role_type"] = u.UserRoleType
}

func (u userRoleIDList) clone(db *gorm.DB) userRoleIDList {
	u.userRoleIDListDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userRoleIDList) replaceDB(db *gorm.DB) userRoleIDList {
	u.userRoleIDListDo.ReplaceDB(db)
	return u
}

type userRoleIDListDo struct{ gen.DO }

type IUserRoleIDListDo interface {
	gen.SubQuery
	Debug() IUserRoleIDListDo
	WithContext(ctx context.Context) IUserRoleIDListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserRoleIDListDo
	WriteDB() IUserRoleIDListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserRoleIDListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserRoleIDListDo
	Not(conds ...gen.Condition) IUserRoleIDListDo
	Or(conds ...gen.Condition) IUserRoleIDListDo
	Select(conds ...field.Expr) IUserRoleIDListDo
	Where(conds ...gen.Condition) IUserRoleIDListDo
	Order(conds ...field.Expr) IUserRoleIDListDo
	Distinct(cols ...field.Expr) IUserRoleIDListDo
	Omit(cols ...field.Expr) IUserRoleIDListDo
	Join(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo
	Group(cols ...field.Expr) IUserRoleIDListDo
	Having(conds ...gen.Condition) IUserRoleIDListDo
	Limit(limit int) IUserRoleIDListDo
	Offset(offset int) IUserRoleIDListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserRoleIDListDo
	Unscoped() IUserRoleIDListDo
	Create(values ...*model.UserRoleIDList) error
	CreateInBatches(values []*model.UserRoleIDList, batchSize int) error
	Save(values ...*model.UserRoleIDList) error
	First() (*model.UserRoleIDList, error)
	Take() (*model.UserRoleIDList, error)
	Last() (*model.UserRoleIDList, error)
	Find() ([]*model.UserRoleIDList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRoleIDList, err error)
	FindInBatches(result *[]*model.UserRoleIDList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserRoleIDList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserRoleIDListDo
	Assign(attrs ...field.AssignExpr) IUserRoleIDListDo
	Joins(fields ...field.RelationField) IUserRoleIDListDo
	Preload(fields ...field.RelationField) IUserRoleIDListDo
	FirstOrInit() (*model.UserRoleIDList, error)
	FirstOrCreate() (*model.UserRoleIDList, error)
	FindByPage(offset int, limit int) (result []*model.UserRoleIDList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserRoleIDListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userRoleIDListDo) Debug() IUserRoleIDListDo {
	return u.withDO(u.DO.Debug())
}

func (u userRoleIDListDo) WithContext(ctx context.Context) IUserRoleIDListDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRoleIDListDo) ReadDB() IUserRoleIDListDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRoleIDListDo) WriteDB() IUserRoleIDListDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRoleIDListDo) Session(config *gorm.Session) IUserRoleIDListDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRoleIDListDo) Clauses(conds ...clause.Expression) IUserRoleIDListDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRoleIDListDo) Returning(value interface{}, columns ...string) IUserRoleIDListDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRoleIDListDo) Not(conds ...gen.Condition) IUserRoleIDListDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRoleIDListDo) Or(conds ...gen.Condition) IUserRoleIDListDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRoleIDListDo) Select(conds ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRoleIDListDo) Where(conds ...gen.Condition) IUserRoleIDListDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRoleIDListDo) Order(conds ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRoleIDListDo) Distinct(cols ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRoleIDListDo) Omit(cols ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRoleIDListDo) Join(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRoleIDListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRoleIDListDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRoleIDListDo) Group(cols ...field.Expr) IUserRoleIDListDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRoleIDListDo) Having(conds ...gen.Condition) IUserRoleIDListDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRoleIDListDo) Limit(limit int) IUserRoleIDListDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRoleIDListDo) Offset(offset int) IUserRoleIDListDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRoleIDListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserRoleIDListDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRoleIDListDo) Unscoped() IUserRoleIDListDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRoleIDListDo) Create(values ...*model.UserRoleIDList) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRoleIDListDo) CreateInBatches(values []*model.UserRoleIDList, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRoleIDListDo) Save(values ...*model.UserRoleIDList) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRoleIDListDo) First() (*model.UserRoleIDList, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleIDList), nil
	}
}

func (u userRoleIDListDo) Take() (*model.UserRoleIDList, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleIDList), nil
	}
}

func (u userRoleIDListDo) Last() (*model.UserRoleIDList, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleIDList), nil
	}
}

func (u userRoleIDListDo) Find() ([]*model.UserRoleIDList, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserRoleIDList), err
}

func (u userRoleIDListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRoleIDList, err error) {
	buf := make([]*model.UserRoleIDList, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRoleIDListDo) FindInBatches(result *[]*model.UserRoleIDList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRoleIDListDo) Attrs(attrs ...field.AssignExpr) IUserRoleIDListDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRoleIDListDo) Assign(attrs ...field.AssignExpr) IUserRoleIDListDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRoleIDListDo) Joins(fields ...field.RelationField) IUserRoleIDListDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRoleIDListDo) Preload(fields ...field.RelationField) IUserRoleIDListDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRoleIDListDo) FirstOrInit() (*model.UserRoleIDList, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleIDList), nil
	}
}

func (u userRoleIDListDo) FirstOrCreate() (*model.UserRoleIDList, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleIDList), nil
	}
}

func (u userRoleIDListDo) FindByPage(offset int, limit int) (result []*model.UserRoleIDList, count int64, err error) {
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

func (u userRoleIDListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRoleIDListDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRoleIDListDo) Delete(models ...*model.UserRoleIDList) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRoleIDListDo) withDO(do gen.Dao) *userRoleIDListDo {
	u.DO = *do.(*gen.DO)
	return u
}
