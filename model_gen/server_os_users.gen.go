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

func newServerOsUsers(db *gorm.DB, opts ...gen.DOOption) serverOsUsers {
	_serverOsUsers := serverOsUsers{}

	_serverOsUsers.serverOsUsersDo.UseDB(db, opts...)
	_serverOsUsers.serverOsUsersDo.UseModel(&model.ServerOsUsers{})

	tableName := _serverOsUsers.serverOsUsersDo.TableName()
	_serverOsUsers.ALL = field.NewAsterisk(tableName)
	_serverOsUsers.ServerOsUsersID = field.NewString(tableName, "server_os_users_id")
	_serverOsUsers.UserName = field.NewString(tableName, "user_name")
	_serverOsUsers.UserPwd = field.NewString(tableName, "user_pwd")
	_serverOsUsers.UserHome = field.NewString(tableName, "user_home")
	_serverOsUsers.OsPrivileageCode = field.NewString(tableName, "os_privileage_code")
	_serverOsUsers.ServerOsUserID = field.NewInt32(tableName, "server_os_user_id")
	_serverOsUsers.CreateTime = field.NewTime(tableName, "create_time")
	_serverOsUsers.ModifyTime = field.NewTime(tableName, "modify_time")
	_serverOsUsers.Comments = field.NewString(tableName, "comments")
	_serverOsUsers.DefaultUser = field.NewString(tableName, "default_user")
	_serverOsUsers.ServerID = field.NewString(tableName, "server_id")
	_serverOsUsers.CreateShell = field.NewString(tableName, "create_shell")
	_serverOsUsers.OssoftUserGroupID = field.NewString(tableName, "ossoft_user_group_id")

	_serverOsUsers.fillFieldMap()

	return _serverOsUsers
}

type serverOsUsers struct {
	serverOsUsersDo

	ALL               field.Asterisk
	ServerOsUsersID   field.String
	UserName          field.String // 用户名
	UserPwd           field.String // 服务器密码
	UserHome          field.String // 用户主目录
	OsPrivileageCode  field.String // 用户操作系统权限等级，见码表
	ServerOsUserID    field.Int32  // 用户ID
	CreateTime        field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime        field.Time   // 记录修改时间（数据库自动写入）
	Comments          field.String // 备注说明
	DefaultUser       field.String
	ServerID          field.String
	CreateShell       field.String
	OssoftUserGroupID field.String

	fieldMap map[string]field.Expr
}

func (s serverOsUsers) Table(newTableName string) *serverOsUsers {
	s.serverOsUsersDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverOsUsers) As(alias string) *serverOsUsers {
	s.serverOsUsersDo.DO = *(s.serverOsUsersDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverOsUsers) updateTableName(table string) *serverOsUsers {
	s.ALL = field.NewAsterisk(table)
	s.ServerOsUsersID = field.NewString(table, "server_os_users_id")
	s.UserName = field.NewString(table, "user_name")
	s.UserPwd = field.NewString(table, "user_pwd")
	s.UserHome = field.NewString(table, "user_home")
	s.OsPrivileageCode = field.NewString(table, "os_privileage_code")
	s.ServerOsUserID = field.NewInt32(table, "server_os_user_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")
	s.DefaultUser = field.NewString(table, "default_user")
	s.ServerID = field.NewString(table, "server_id")
	s.CreateShell = field.NewString(table, "create_shell")
	s.OssoftUserGroupID = field.NewString(table, "ossoft_user_group_id")

	s.fillFieldMap()

	return s
}

func (s *serverOsUsers) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverOsUsers) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["server_os_users_id"] = s.ServerOsUsersID
	s.fieldMap["user_name"] = s.UserName
	s.fieldMap["user_pwd"] = s.UserPwd
	s.fieldMap["user_home"] = s.UserHome
	s.fieldMap["os_privileage_code"] = s.OsPrivileageCode
	s.fieldMap["server_os_user_id"] = s.ServerOsUserID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["default_user"] = s.DefaultUser
	s.fieldMap["server_id"] = s.ServerID
	s.fieldMap["create_shell"] = s.CreateShell
	s.fieldMap["ossoft_user_group_id"] = s.OssoftUserGroupID
}

func (s serverOsUsers) clone(db *gorm.DB) serverOsUsers {
	s.serverOsUsersDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverOsUsers) replaceDB(db *gorm.DB) serverOsUsers {
	s.serverOsUsersDo.ReplaceDB(db)
	return s
}

type serverOsUsersDo struct{ gen.DO }

type IServerOsUsersDo interface {
	gen.SubQuery
	Debug() IServerOsUsersDo
	WithContext(ctx context.Context) IServerOsUsersDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerOsUsersDo
	WriteDB() IServerOsUsersDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerOsUsersDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerOsUsersDo
	Not(conds ...gen.Condition) IServerOsUsersDo
	Or(conds ...gen.Condition) IServerOsUsersDo
	Select(conds ...field.Expr) IServerOsUsersDo
	Where(conds ...gen.Condition) IServerOsUsersDo
	Order(conds ...field.Expr) IServerOsUsersDo
	Distinct(cols ...field.Expr) IServerOsUsersDo
	Omit(cols ...field.Expr) IServerOsUsersDo
	Join(table schema.Tabler, on ...field.Expr) IServerOsUsersDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerOsUsersDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerOsUsersDo
	Group(cols ...field.Expr) IServerOsUsersDo
	Having(conds ...gen.Condition) IServerOsUsersDo
	Limit(limit int) IServerOsUsersDo
	Offset(offset int) IServerOsUsersDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerOsUsersDo
	Unscoped() IServerOsUsersDo
	Create(values ...*model.ServerOsUsers) error
	CreateInBatches(values []*model.ServerOsUsers, batchSize int) error
	Save(values ...*model.ServerOsUsers) error
	First() (*model.ServerOsUsers, error)
	Take() (*model.ServerOsUsers, error)
	Last() (*model.ServerOsUsers, error)
	Find() ([]*model.ServerOsUsers, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerOsUsers, err error)
	FindInBatches(result *[]*model.ServerOsUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerOsUsers) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerOsUsersDo
	Assign(attrs ...field.AssignExpr) IServerOsUsersDo
	Joins(fields ...field.RelationField) IServerOsUsersDo
	Preload(fields ...field.RelationField) IServerOsUsersDo
	FirstOrInit() (*model.ServerOsUsers, error)
	FirstOrCreate() (*model.ServerOsUsers, error)
	FindByPage(offset int, limit int) (result []*model.ServerOsUsers, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerOsUsersDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverOsUsersDo) Debug() IServerOsUsersDo {
	return s.withDO(s.DO.Debug())
}

func (s serverOsUsersDo) WithContext(ctx context.Context) IServerOsUsersDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverOsUsersDo) ReadDB() IServerOsUsersDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverOsUsersDo) WriteDB() IServerOsUsersDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverOsUsersDo) Session(config *gorm.Session) IServerOsUsersDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverOsUsersDo) Clauses(conds ...clause.Expression) IServerOsUsersDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverOsUsersDo) Returning(value interface{}, columns ...string) IServerOsUsersDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverOsUsersDo) Not(conds ...gen.Condition) IServerOsUsersDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverOsUsersDo) Or(conds ...gen.Condition) IServerOsUsersDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverOsUsersDo) Select(conds ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverOsUsersDo) Where(conds ...gen.Condition) IServerOsUsersDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverOsUsersDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerOsUsersDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverOsUsersDo) Order(conds ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverOsUsersDo) Distinct(cols ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverOsUsersDo) Omit(cols ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverOsUsersDo) Join(table schema.Tabler, on ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverOsUsersDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverOsUsersDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverOsUsersDo) Group(cols ...field.Expr) IServerOsUsersDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverOsUsersDo) Having(conds ...gen.Condition) IServerOsUsersDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverOsUsersDo) Limit(limit int) IServerOsUsersDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverOsUsersDo) Offset(offset int) IServerOsUsersDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverOsUsersDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerOsUsersDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverOsUsersDo) Unscoped() IServerOsUsersDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverOsUsersDo) Create(values ...*model.ServerOsUsers) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverOsUsersDo) CreateInBatches(values []*model.ServerOsUsers, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverOsUsersDo) Save(values ...*model.ServerOsUsers) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverOsUsersDo) First() (*model.ServerOsUsers, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerOsUsers), nil
	}
}

func (s serverOsUsersDo) Take() (*model.ServerOsUsers, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerOsUsers), nil
	}
}

func (s serverOsUsersDo) Last() (*model.ServerOsUsers, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerOsUsers), nil
	}
}

func (s serverOsUsersDo) Find() ([]*model.ServerOsUsers, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerOsUsers), err
}

func (s serverOsUsersDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerOsUsers, err error) {
	buf := make([]*model.ServerOsUsers, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverOsUsersDo) FindInBatches(result *[]*model.ServerOsUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverOsUsersDo) Attrs(attrs ...field.AssignExpr) IServerOsUsersDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverOsUsersDo) Assign(attrs ...field.AssignExpr) IServerOsUsersDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverOsUsersDo) Joins(fields ...field.RelationField) IServerOsUsersDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverOsUsersDo) Preload(fields ...field.RelationField) IServerOsUsersDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverOsUsersDo) FirstOrInit() (*model.ServerOsUsers, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerOsUsers), nil
	}
}

func (s serverOsUsersDo) FirstOrCreate() (*model.ServerOsUsers, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerOsUsers), nil
	}
}

func (s serverOsUsersDo) FindByPage(offset int, limit int) (result []*model.ServerOsUsers, count int64, err error) {
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

func (s serverOsUsersDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverOsUsersDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverOsUsersDo) Delete(models ...*model.ServerOsUsers) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverOsUsersDo) withDO(do gen.Dao) *serverOsUsersDo {
	s.DO = *do.(*gen.DO)
	return s
}
