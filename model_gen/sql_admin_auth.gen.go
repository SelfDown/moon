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

func newSqlAdminAuth(db *gorm.DB, opts ...gen.DOOption) sqlAdminAuth {
	_sqlAdminAuth := sqlAdminAuth{}

	_sqlAdminAuth.sqlAdminAuthDo.UseDB(db, opts...)
	_sqlAdminAuth.sqlAdminAuthDo.UseModel(&model.SqlAdminAuth{})

	tableName := _sqlAdminAuth.sqlAdminAuthDo.TableName()
	_sqlAdminAuth.ALL = field.NewAsterisk(tableName)
	_sqlAdminAuth.SqlAdminAuthID = field.NewString(tableName, "sql_admin_auth_id")
	_sqlAdminAuth.EnvType = field.NewString(tableName, "env_type")
	_sqlAdminAuth.SqlAdminRoleID = field.NewString(tableName, "sql_admin_role_id")
	_sqlAdminAuth.AuthBy = field.NewString(tableName, "auth_by")
	_sqlAdminAuth.CreateTime = field.NewTime(tableName, "create_time")
	_sqlAdminAuth.ModifyTime = field.NewTime(tableName, "modify_time")
	_sqlAdminAuth.Comments = field.NewString(tableName, "comments")
	_sqlAdminAuth.UserID = field.NewString(tableName, "user_id")

	_sqlAdminAuth.fillFieldMap()

	return _sqlAdminAuth
}

type sqlAdminAuth struct {
	sqlAdminAuthDo

	ALL            field.Asterisk
	SqlAdminAuthID field.String
	/*
		环境类型
		生产
		非生产
	*/
	EnvType        field.String
	SqlAdminRoleID field.String // 关联 sql_admin_role.sql_admin_role_id
	AuthBy         field.String // 创建这条记录的操作人 user_id
	CreateTime     field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime     field.Time   // 记录修改时间（数据库自动写入）
	Comments       field.String // 备注说明
	UserID         field.String // 被授权人userid

	fieldMap map[string]field.Expr
}

func (s sqlAdminAuth) Table(newTableName string) *sqlAdminAuth {
	s.sqlAdminAuthDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sqlAdminAuth) As(alias string) *sqlAdminAuth {
	s.sqlAdminAuthDo.DO = *(s.sqlAdminAuthDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sqlAdminAuth) updateTableName(table string) *sqlAdminAuth {
	s.ALL = field.NewAsterisk(table)
	s.SqlAdminAuthID = field.NewString(table, "sql_admin_auth_id")
	s.EnvType = field.NewString(table, "env_type")
	s.SqlAdminRoleID = field.NewString(table, "sql_admin_role_id")
	s.AuthBy = field.NewString(table, "auth_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")
	s.UserID = field.NewString(table, "user_id")

	s.fillFieldMap()

	return s
}

func (s *sqlAdminAuth) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sqlAdminAuth) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["sql_admin_auth_id"] = s.SqlAdminAuthID
	s.fieldMap["env_type"] = s.EnvType
	s.fieldMap["sql_admin_role_id"] = s.SqlAdminRoleID
	s.fieldMap["auth_by"] = s.AuthBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["user_id"] = s.UserID
}

func (s sqlAdminAuth) clone(db *gorm.DB) sqlAdminAuth {
	s.sqlAdminAuthDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sqlAdminAuth) replaceDB(db *gorm.DB) sqlAdminAuth {
	s.sqlAdminAuthDo.ReplaceDB(db)
	return s
}

type sqlAdminAuthDo struct{ gen.DO }

type ISqlAdminAuthDo interface {
	gen.SubQuery
	Debug() ISqlAdminAuthDo
	WithContext(ctx context.Context) ISqlAdminAuthDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISqlAdminAuthDo
	WriteDB() ISqlAdminAuthDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISqlAdminAuthDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISqlAdminAuthDo
	Not(conds ...gen.Condition) ISqlAdminAuthDo
	Or(conds ...gen.Condition) ISqlAdminAuthDo
	Select(conds ...field.Expr) ISqlAdminAuthDo
	Where(conds ...gen.Condition) ISqlAdminAuthDo
	Order(conds ...field.Expr) ISqlAdminAuthDo
	Distinct(cols ...field.Expr) ISqlAdminAuthDo
	Omit(cols ...field.Expr) ISqlAdminAuthDo
	Join(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo
	Group(cols ...field.Expr) ISqlAdminAuthDo
	Having(conds ...gen.Condition) ISqlAdminAuthDo
	Limit(limit int) ISqlAdminAuthDo
	Offset(offset int) ISqlAdminAuthDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlAdminAuthDo
	Unscoped() ISqlAdminAuthDo
	Create(values ...*model.SqlAdminAuth) error
	CreateInBatches(values []*model.SqlAdminAuth, batchSize int) error
	Save(values ...*model.SqlAdminAuth) error
	First() (*model.SqlAdminAuth, error)
	Take() (*model.SqlAdminAuth, error)
	Last() (*model.SqlAdminAuth, error)
	Find() ([]*model.SqlAdminAuth, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlAdminAuth, err error)
	FindInBatches(result *[]*model.SqlAdminAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SqlAdminAuth) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISqlAdminAuthDo
	Assign(attrs ...field.AssignExpr) ISqlAdminAuthDo
	Joins(fields ...field.RelationField) ISqlAdminAuthDo
	Preload(fields ...field.RelationField) ISqlAdminAuthDo
	FirstOrInit() (*model.SqlAdminAuth, error)
	FirstOrCreate() (*model.SqlAdminAuth, error)
	FindByPage(offset int, limit int) (result []*model.SqlAdminAuth, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISqlAdminAuthDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sqlAdminAuthDo) Debug() ISqlAdminAuthDo {
	return s.withDO(s.DO.Debug())
}

func (s sqlAdminAuthDo) WithContext(ctx context.Context) ISqlAdminAuthDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sqlAdminAuthDo) ReadDB() ISqlAdminAuthDo {
	return s.Clauses(dbresolver.Read)
}

func (s sqlAdminAuthDo) WriteDB() ISqlAdminAuthDo {
	return s.Clauses(dbresolver.Write)
}

func (s sqlAdminAuthDo) Session(config *gorm.Session) ISqlAdminAuthDo {
	return s.withDO(s.DO.Session(config))
}

func (s sqlAdminAuthDo) Clauses(conds ...clause.Expression) ISqlAdminAuthDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sqlAdminAuthDo) Returning(value interface{}, columns ...string) ISqlAdminAuthDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sqlAdminAuthDo) Not(conds ...gen.Condition) ISqlAdminAuthDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sqlAdminAuthDo) Or(conds ...gen.Condition) ISqlAdminAuthDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sqlAdminAuthDo) Select(conds ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sqlAdminAuthDo) Where(conds ...gen.Condition) ISqlAdminAuthDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sqlAdminAuthDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISqlAdminAuthDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sqlAdminAuthDo) Order(conds ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sqlAdminAuthDo) Distinct(cols ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sqlAdminAuthDo) Omit(cols ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sqlAdminAuthDo) Join(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sqlAdminAuthDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sqlAdminAuthDo) RightJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sqlAdminAuthDo) Group(cols ...field.Expr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sqlAdminAuthDo) Having(conds ...gen.Condition) ISqlAdminAuthDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sqlAdminAuthDo) Limit(limit int) ISqlAdminAuthDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sqlAdminAuthDo) Offset(offset int) ISqlAdminAuthDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sqlAdminAuthDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlAdminAuthDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sqlAdminAuthDo) Unscoped() ISqlAdminAuthDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sqlAdminAuthDo) Create(values ...*model.SqlAdminAuth) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sqlAdminAuthDo) CreateInBatches(values []*model.SqlAdminAuth, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sqlAdminAuthDo) Save(values ...*model.SqlAdminAuth) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sqlAdminAuthDo) First() (*model.SqlAdminAuth, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAuth), nil
	}
}

func (s sqlAdminAuthDo) Take() (*model.SqlAdminAuth, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAuth), nil
	}
}

func (s sqlAdminAuthDo) Last() (*model.SqlAdminAuth, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAuth), nil
	}
}

func (s sqlAdminAuthDo) Find() ([]*model.SqlAdminAuth, error) {
	result, err := s.DO.Find()
	return result.([]*model.SqlAdminAuth), err
}

func (s sqlAdminAuthDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlAdminAuth, err error) {
	buf := make([]*model.SqlAdminAuth, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sqlAdminAuthDo) FindInBatches(result *[]*model.SqlAdminAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sqlAdminAuthDo) Attrs(attrs ...field.AssignExpr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sqlAdminAuthDo) Assign(attrs ...field.AssignExpr) ISqlAdminAuthDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sqlAdminAuthDo) Joins(fields ...field.RelationField) ISqlAdminAuthDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sqlAdminAuthDo) Preload(fields ...field.RelationField) ISqlAdminAuthDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sqlAdminAuthDo) FirstOrInit() (*model.SqlAdminAuth, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAuth), nil
	}
}

func (s sqlAdminAuthDo) FirstOrCreate() (*model.SqlAdminAuth, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAuth), nil
	}
}

func (s sqlAdminAuthDo) FindByPage(offset int, limit int) (result []*model.SqlAdminAuth, count int64, err error) {
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

func (s sqlAdminAuthDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sqlAdminAuthDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sqlAdminAuthDo) Delete(models ...*model.SqlAdminAuth) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sqlAdminAuthDo) withDO(do gen.Dao) *sqlAdminAuthDo {
	s.DO = *do.(*gen.DO)
	return s
}
