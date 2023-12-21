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

func newUserRoleHospital(db *gorm.DB, opts ...gen.DOOption) userRoleHospital {
	_userRoleHospital := userRoleHospital{}

	_userRoleHospital.userRoleHospitalDo.UseDB(db, opts...)
	_userRoleHospital.userRoleHospitalDo.UseModel(&model.UserRoleHospital{})

	tableName := _userRoleHospital.userRoleHospitalDo.TableName()
	_userRoleHospital.ALL = field.NewAsterisk(tableName)
	_userRoleHospital.RoleID = field.NewString(tableName, "role_id")
	_userRoleHospital.HospitalCode = field.NewString(tableName, "hospital_code")
	_userRoleHospital.IsDefault = field.NewInt32(tableName, "is_default")
	_userRoleHospital.CreateTime = field.NewTime(tableName, "create_time")
	_userRoleHospital.ModifyTime = field.NewTime(tableName, "modify_time")
	_userRoleHospital.Comments = field.NewString(tableName, "comments")

	_userRoleHospital.fillFieldMap()

	return _userRoleHospital
}

type userRoleHospital struct {
	userRoleHospitalDo

	ALL          field.Asterisk
	RoleID       field.String
	HospitalCode field.String
	IsDefault    field.Int32  // 1 默认医院  0 非默认
	CreateTime   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime   field.Time   // 记录修改时间（数据库自动写入）
	Comments     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (u userRoleHospital) Table(newTableName string) *userRoleHospital {
	u.userRoleHospitalDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRoleHospital) As(alias string) *userRoleHospital {
	u.userRoleHospitalDo.DO = *(u.userRoleHospitalDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRoleHospital) updateTableName(table string) *userRoleHospital {
	u.ALL = field.NewAsterisk(table)
	u.RoleID = field.NewString(table, "role_id")
	u.HospitalCode = field.NewString(table, "hospital_code")
	u.IsDefault = field.NewInt32(table, "is_default")
	u.CreateTime = field.NewTime(table, "create_time")
	u.ModifyTime = field.NewTime(table, "modify_time")
	u.Comments = field.NewString(table, "comments")

	u.fillFieldMap()

	return u
}

func (u *userRoleHospital) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRoleHospital) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["role_id"] = u.RoleID
	u.fieldMap["hospital_code"] = u.HospitalCode
	u.fieldMap["is_default"] = u.IsDefault
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["modify_time"] = u.ModifyTime
	u.fieldMap["comments"] = u.Comments
}

func (u userRoleHospital) clone(db *gorm.DB) userRoleHospital {
	u.userRoleHospitalDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userRoleHospital) replaceDB(db *gorm.DB) userRoleHospital {
	u.userRoleHospitalDo.ReplaceDB(db)
	return u
}

type userRoleHospitalDo struct{ gen.DO }

type IUserRoleHospitalDo interface {
	gen.SubQuery
	Debug() IUserRoleHospitalDo
	WithContext(ctx context.Context) IUserRoleHospitalDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserRoleHospitalDo
	WriteDB() IUserRoleHospitalDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserRoleHospitalDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserRoleHospitalDo
	Not(conds ...gen.Condition) IUserRoleHospitalDo
	Or(conds ...gen.Condition) IUserRoleHospitalDo
	Select(conds ...field.Expr) IUserRoleHospitalDo
	Where(conds ...gen.Condition) IUserRoleHospitalDo
	Order(conds ...field.Expr) IUserRoleHospitalDo
	Distinct(cols ...field.Expr) IUserRoleHospitalDo
	Omit(cols ...field.Expr) IUserRoleHospitalDo
	Join(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo
	Group(cols ...field.Expr) IUserRoleHospitalDo
	Having(conds ...gen.Condition) IUserRoleHospitalDo
	Limit(limit int) IUserRoleHospitalDo
	Offset(offset int) IUserRoleHospitalDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserRoleHospitalDo
	Unscoped() IUserRoleHospitalDo
	Create(values ...*model.UserRoleHospital) error
	CreateInBatches(values []*model.UserRoleHospital, batchSize int) error
	Save(values ...*model.UserRoleHospital) error
	First() (*model.UserRoleHospital, error)
	Take() (*model.UserRoleHospital, error)
	Last() (*model.UserRoleHospital, error)
	Find() ([]*model.UserRoleHospital, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRoleHospital, err error)
	FindInBatches(result *[]*model.UserRoleHospital, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserRoleHospital) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserRoleHospitalDo
	Assign(attrs ...field.AssignExpr) IUserRoleHospitalDo
	Joins(fields ...field.RelationField) IUserRoleHospitalDo
	Preload(fields ...field.RelationField) IUserRoleHospitalDo
	FirstOrInit() (*model.UserRoleHospital, error)
	FirstOrCreate() (*model.UserRoleHospital, error)
	FindByPage(offset int, limit int) (result []*model.UserRoleHospital, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserRoleHospitalDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userRoleHospitalDo) Debug() IUserRoleHospitalDo {
	return u.withDO(u.DO.Debug())
}

func (u userRoleHospitalDo) WithContext(ctx context.Context) IUserRoleHospitalDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRoleHospitalDo) ReadDB() IUserRoleHospitalDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRoleHospitalDo) WriteDB() IUserRoleHospitalDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRoleHospitalDo) Session(config *gorm.Session) IUserRoleHospitalDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRoleHospitalDo) Clauses(conds ...clause.Expression) IUserRoleHospitalDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRoleHospitalDo) Returning(value interface{}, columns ...string) IUserRoleHospitalDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRoleHospitalDo) Not(conds ...gen.Condition) IUserRoleHospitalDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRoleHospitalDo) Or(conds ...gen.Condition) IUserRoleHospitalDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRoleHospitalDo) Select(conds ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRoleHospitalDo) Where(conds ...gen.Condition) IUserRoleHospitalDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRoleHospitalDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserRoleHospitalDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userRoleHospitalDo) Order(conds ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRoleHospitalDo) Distinct(cols ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRoleHospitalDo) Omit(cols ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRoleHospitalDo) Join(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRoleHospitalDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRoleHospitalDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRoleHospitalDo) Group(cols ...field.Expr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRoleHospitalDo) Having(conds ...gen.Condition) IUserRoleHospitalDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRoleHospitalDo) Limit(limit int) IUserRoleHospitalDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRoleHospitalDo) Offset(offset int) IUserRoleHospitalDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRoleHospitalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserRoleHospitalDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRoleHospitalDo) Unscoped() IUserRoleHospitalDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRoleHospitalDo) Create(values ...*model.UserRoleHospital) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRoleHospitalDo) CreateInBatches(values []*model.UserRoleHospital, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRoleHospitalDo) Save(values ...*model.UserRoleHospital) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRoleHospitalDo) First() (*model.UserRoleHospital, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleHospital), nil
	}
}

func (u userRoleHospitalDo) Take() (*model.UserRoleHospital, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleHospital), nil
	}
}

func (u userRoleHospitalDo) Last() (*model.UserRoleHospital, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleHospital), nil
	}
}

func (u userRoleHospitalDo) Find() ([]*model.UserRoleHospital, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserRoleHospital), err
}

func (u userRoleHospitalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRoleHospital, err error) {
	buf := make([]*model.UserRoleHospital, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRoleHospitalDo) FindInBatches(result *[]*model.UserRoleHospital, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRoleHospitalDo) Attrs(attrs ...field.AssignExpr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRoleHospitalDo) Assign(attrs ...field.AssignExpr) IUserRoleHospitalDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRoleHospitalDo) Joins(fields ...field.RelationField) IUserRoleHospitalDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRoleHospitalDo) Preload(fields ...field.RelationField) IUserRoleHospitalDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRoleHospitalDo) FirstOrInit() (*model.UserRoleHospital, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleHospital), nil
	}
}

func (u userRoleHospitalDo) FirstOrCreate() (*model.UserRoleHospital, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRoleHospital), nil
	}
}

func (u userRoleHospitalDo) FindByPage(offset int, limit int) (result []*model.UserRoleHospital, count int64, err error) {
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

func (u userRoleHospitalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRoleHospitalDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRoleHospitalDo) Delete(models ...*model.UserRoleHospital) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRoleHospitalDo) withDO(do gen.Dao) *userRoleHospitalDo {
	u.DO = *do.(*gen.DO)
	return u
}