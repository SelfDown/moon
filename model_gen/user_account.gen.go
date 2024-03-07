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

func newUserAccount(db *gorm.DB, opts ...gen.DOOption) userAccount {
	_userAccount := userAccount{}

	_userAccount.userAccountDo.UseDB(db, opts...)
	_userAccount.userAccountDo.UseModel(&model.UserAccount{})

	tableName := _userAccount.userAccountDo.TableName()
	_userAccount.ALL = field.NewAsterisk(tableName)
	_userAccount.Userid = field.NewString(tableName, "userid")
	_userAccount.RoleID = field.NewString(tableName, "role_id")
	_userAccount.Username = field.NewString(tableName, "username")
	_userAccount.Userpwd = field.NewString(tableName, "userpwd")
	_userAccount.LastLoginTime = field.NewTime(tableName, "last_login_time")
	_userAccount.LastLoginIP = field.NewString(tableName, "last_login_ip")
	_userAccount.Email = field.NewString(tableName, "email")
	_userAccount.Nick = field.NewString(tableName, "nick")
	_userAccount.Statu = field.NewInt32(tableName, "statu")
	_userAccount.Address = field.NewString(tableName, "address")
	_userAccount.Note = field.NewString(tableName, "note")
	_userAccount.Email2 = field.NewString(tableName, "email2")
	_userAccount.CreateTime = field.NewTime(tableName, "create_time")
	_userAccount.LastLoginFailureTime = field.NewTime(tableName, "last_login_failure_time")
	_userAccount.LoginFailureCount = field.NewInt32(tableName, "login_failure_count")
	_userAccount.Avatar = field.NewString(tableName, "avatar")
	_userAccount.ModifyTime = field.NewTime(tableName, "modify_time")
	_userAccount.Comments = field.NewString(tableName, "comments")
	_userAccount.Tel = field.NewString(tableName, "tel")
	_userAccount.WechatUserid = field.NewString(tableName, "wechat_userid")
	_userAccount.AttendanceWechat = field.NewString(tableName, "attendance_wechat")

	_userAccount.fillFieldMap()

	return _userAccount
}

type userAccount struct {
	userAccountDo

	ALL                  field.Asterisk
	Userid               field.String
	RoleID               field.String
	Username             field.String
	Userpwd              field.String
	LastLoginTime        field.Time   // 最后一次登录时间
	LastLoginIP          field.String // 最后一次登录IP
	Email                field.String // 用户邮箱
	Nick                 field.String // 用户昵称
	Statu                field.Int32
	Address              field.String // 物理地址
	Note                 field.String
	Email2               field.String // 用户备用邮箱
	CreateTime           field.Time   // 记录创建时间（数据库自动写入）
	LastLoginFailureTime field.Time   // 最近一次登录失败时间
	LoginFailureCount    field.Int32  // 登录失败计数
	Avatar               field.String // 用户头像base64
	ModifyTime           field.Time   // 记录修改时间（数据库自动写入）
	Comments             field.String // 备注说明
	Tel                  field.String
	WechatUserid         field.String
	AttendanceWechat     field.String

	fieldMap map[string]field.Expr
}

func (u userAccount) Table(newTableName string) *userAccount {
	u.userAccountDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userAccount) As(alias string) *userAccount {
	u.userAccountDo.DO = *(u.userAccountDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userAccount) updateTableName(table string) *userAccount {
	u.ALL = field.NewAsterisk(table)
	u.Userid = field.NewString(table, "userid")
	u.RoleID = field.NewString(table, "role_id")
	u.Username = field.NewString(table, "username")
	u.Userpwd = field.NewString(table, "userpwd")
	u.LastLoginTime = field.NewTime(table, "last_login_time")
	u.LastLoginIP = field.NewString(table, "last_login_ip")
	u.Email = field.NewString(table, "email")
	u.Nick = field.NewString(table, "nick")
	u.Statu = field.NewInt32(table, "statu")
	u.Address = field.NewString(table, "address")
	u.Note = field.NewString(table, "note")
	u.Email2 = field.NewString(table, "email2")
	u.CreateTime = field.NewTime(table, "create_time")
	u.LastLoginFailureTime = field.NewTime(table, "last_login_failure_time")
	u.LoginFailureCount = field.NewInt32(table, "login_failure_count")
	u.Avatar = field.NewString(table, "avatar")
	u.ModifyTime = field.NewTime(table, "modify_time")
	u.Comments = field.NewString(table, "comments")
	u.Tel = field.NewString(table, "tel")
	u.WechatUserid = field.NewString(table, "wechat_userid")
	u.AttendanceWechat = field.NewString(table, "attendance_wechat")

	u.fillFieldMap()

	return u
}

func (u *userAccount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userAccount) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 21)
	u.fieldMap["userid"] = u.Userid
	u.fieldMap["role_id"] = u.RoleID
	u.fieldMap["username"] = u.Username
	u.fieldMap["userpwd"] = u.Userpwd
	u.fieldMap["last_login_time"] = u.LastLoginTime
	u.fieldMap["last_login_ip"] = u.LastLoginIP
	u.fieldMap["email"] = u.Email
	u.fieldMap["nick"] = u.Nick
	u.fieldMap["statu"] = u.Statu
	u.fieldMap["address"] = u.Address
	u.fieldMap["note"] = u.Note
	u.fieldMap["email2"] = u.Email2
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["last_login_failure_time"] = u.LastLoginFailureTime
	u.fieldMap["login_failure_count"] = u.LoginFailureCount
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["modify_time"] = u.ModifyTime
	u.fieldMap["comments"] = u.Comments
	u.fieldMap["tel"] = u.Tel
	u.fieldMap["wechat_userid"] = u.WechatUserid
	u.fieldMap["attendance_wechat"] = u.AttendanceWechat
}

func (u userAccount) clone(db *gorm.DB) userAccount {
	u.userAccountDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userAccount) replaceDB(db *gorm.DB) userAccount {
	u.userAccountDo.ReplaceDB(db)
	return u
}

type userAccountDo struct{ gen.DO }

type IUserAccountDo interface {
	gen.SubQuery
	Debug() IUserAccountDo
	WithContext(ctx context.Context) IUserAccountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserAccountDo
	WriteDB() IUserAccountDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserAccountDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserAccountDo
	Not(conds ...gen.Condition) IUserAccountDo
	Or(conds ...gen.Condition) IUserAccountDo
	Select(conds ...field.Expr) IUserAccountDo
	Where(conds ...gen.Condition) IUserAccountDo
	Order(conds ...field.Expr) IUserAccountDo
	Distinct(cols ...field.Expr) IUserAccountDo
	Omit(cols ...field.Expr) IUserAccountDo
	Join(table schema.Tabler, on ...field.Expr) IUserAccountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserAccountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserAccountDo
	Group(cols ...field.Expr) IUserAccountDo
	Having(conds ...gen.Condition) IUserAccountDo
	Limit(limit int) IUserAccountDo
	Offset(offset int) IUserAccountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAccountDo
	Unscoped() IUserAccountDo
	Create(values ...*model.UserAccount) error
	CreateInBatches(values []*model.UserAccount, batchSize int) error
	Save(values ...*model.UserAccount) error
	First() (*model.UserAccount, error)
	Take() (*model.UserAccount, error)
	Last() (*model.UserAccount, error)
	Find() ([]*model.UserAccount, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAccount, err error)
	FindInBatches(result *[]*model.UserAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserAccount) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserAccountDo
	Assign(attrs ...field.AssignExpr) IUserAccountDo
	Joins(fields ...field.RelationField) IUserAccountDo
	Preload(fields ...field.RelationField) IUserAccountDo
	FirstOrInit() (*model.UserAccount, error)
	FirstOrCreate() (*model.UserAccount, error)
	FindByPage(offset int, limit int) (result []*model.UserAccount, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserAccountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userAccountDo) Debug() IUserAccountDo {
	return u.withDO(u.DO.Debug())
}

func (u userAccountDo) WithContext(ctx context.Context) IUserAccountDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userAccountDo) ReadDB() IUserAccountDo {
	return u.Clauses(dbresolver.Read)
}

func (u userAccountDo) WriteDB() IUserAccountDo {
	return u.Clauses(dbresolver.Write)
}

func (u userAccountDo) Session(config *gorm.Session) IUserAccountDo {
	return u.withDO(u.DO.Session(config))
}

func (u userAccountDo) Clauses(conds ...clause.Expression) IUserAccountDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userAccountDo) Returning(value interface{}, columns ...string) IUserAccountDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userAccountDo) Not(conds ...gen.Condition) IUserAccountDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userAccountDo) Or(conds ...gen.Condition) IUserAccountDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userAccountDo) Select(conds ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userAccountDo) Where(conds ...gen.Condition) IUserAccountDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userAccountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserAccountDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userAccountDo) Order(conds ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userAccountDo) Distinct(cols ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userAccountDo) Omit(cols ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userAccountDo) Join(table schema.Tabler, on ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userAccountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userAccountDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userAccountDo) Group(cols ...field.Expr) IUserAccountDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userAccountDo) Having(conds ...gen.Condition) IUserAccountDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userAccountDo) Limit(limit int) IUserAccountDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userAccountDo) Offset(offset int) IUserAccountDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userAccountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAccountDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userAccountDo) Unscoped() IUserAccountDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userAccountDo) Create(values ...*model.UserAccount) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userAccountDo) CreateInBatches(values []*model.UserAccount, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userAccountDo) Save(values ...*model.UserAccount) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userAccountDo) First() (*model.UserAccount, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccount), nil
	}
}

func (u userAccountDo) Take() (*model.UserAccount, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccount), nil
	}
}

func (u userAccountDo) Last() (*model.UserAccount, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccount), nil
	}
}

func (u userAccountDo) Find() ([]*model.UserAccount, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserAccount), err
}

func (u userAccountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAccount, err error) {
	buf := make([]*model.UserAccount, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userAccountDo) FindInBatches(result *[]*model.UserAccount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userAccountDo) Attrs(attrs ...field.AssignExpr) IUserAccountDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userAccountDo) Assign(attrs ...field.AssignExpr) IUserAccountDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userAccountDo) Joins(fields ...field.RelationField) IUserAccountDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userAccountDo) Preload(fields ...field.RelationField) IUserAccountDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userAccountDo) FirstOrInit() (*model.UserAccount, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccount), nil
	}
}

func (u userAccountDo) FirstOrCreate() (*model.UserAccount, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccount), nil
	}
}

func (u userAccountDo) FindByPage(offset int, limit int) (result []*model.UserAccount, count int64, err error) {
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

func (u userAccountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userAccountDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userAccountDo) Delete(models ...*model.UserAccount) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userAccountDo) withDO(do gen.Dao) *userAccountDo {
	u.DO = *do.(*gen.DO)
	return u
}
