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

func newSqlAdminAPIToken(db *gorm.DB, opts ...gen.DOOption) sqlAdminAPIToken {
	_sqlAdminAPIToken := sqlAdminAPIToken{}

	_sqlAdminAPIToken.sqlAdminAPITokenDo.UseDB(db, opts...)
	_sqlAdminAPIToken.sqlAdminAPITokenDo.UseModel(&model.SqlAdminAPIToken{})

	tableName := _sqlAdminAPIToken.sqlAdminAPITokenDo.TableName()
	_sqlAdminAPIToken.ALL = field.NewAsterisk(tableName)
	_sqlAdminAPIToken.SqlAdminAPIToken = field.NewString(tableName, "sql_admin_api_token")
	_sqlAdminAPIToken.UserID = field.NewString(tableName, "user_id")
	_sqlAdminAPIToken.CreateTime = field.NewTime(tableName, "create_time")
	_sqlAdminAPIToken.ModifyTime = field.NewTime(tableName, "modify_time")
	_sqlAdminAPIToken.Comments = field.NewString(tableName, "comments")

	_sqlAdminAPIToken.fillFieldMap()

	return _sqlAdminAPIToken
}

type sqlAdminAPIToken struct {
	sqlAdminAPITokenDo

	ALL              field.Asterisk
	SqlAdminAPIToken field.String // 令牌
	UserID           field.String // 令牌使用人
	CreateTime       field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime       field.Time   // 记录修改时间（数据库自动写入）
	Comments         field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sqlAdminAPIToken) Table(newTableName string) *sqlAdminAPIToken {
	s.sqlAdminAPITokenDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sqlAdminAPIToken) As(alias string) *sqlAdminAPIToken {
	s.sqlAdminAPITokenDo.DO = *(s.sqlAdminAPITokenDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sqlAdminAPIToken) updateTableName(table string) *sqlAdminAPIToken {
	s.ALL = field.NewAsterisk(table)
	s.SqlAdminAPIToken = field.NewString(table, "sql_admin_api_token")
	s.UserID = field.NewString(table, "user_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sqlAdminAPIToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sqlAdminAPIToken) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["sql_admin_api_token"] = s.SqlAdminAPIToken
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sqlAdminAPIToken) clone(db *gorm.DB) sqlAdminAPIToken {
	s.sqlAdminAPITokenDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sqlAdminAPIToken) replaceDB(db *gorm.DB) sqlAdminAPIToken {
	s.sqlAdminAPITokenDo.ReplaceDB(db)
	return s
}

type sqlAdminAPITokenDo struct{ gen.DO }

type ISqlAdminAPITokenDo interface {
	gen.SubQuery
	Debug() ISqlAdminAPITokenDo
	WithContext(ctx context.Context) ISqlAdminAPITokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISqlAdminAPITokenDo
	WriteDB() ISqlAdminAPITokenDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISqlAdminAPITokenDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISqlAdminAPITokenDo
	Not(conds ...gen.Condition) ISqlAdminAPITokenDo
	Or(conds ...gen.Condition) ISqlAdminAPITokenDo
	Select(conds ...field.Expr) ISqlAdminAPITokenDo
	Where(conds ...gen.Condition) ISqlAdminAPITokenDo
	Order(conds ...field.Expr) ISqlAdminAPITokenDo
	Distinct(cols ...field.Expr) ISqlAdminAPITokenDo
	Omit(cols ...field.Expr) ISqlAdminAPITokenDo
	Join(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo
	Group(cols ...field.Expr) ISqlAdminAPITokenDo
	Having(conds ...gen.Condition) ISqlAdminAPITokenDo
	Limit(limit int) ISqlAdminAPITokenDo
	Offset(offset int) ISqlAdminAPITokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlAdminAPITokenDo
	Unscoped() ISqlAdminAPITokenDo
	Create(values ...*model.SqlAdminAPIToken) error
	CreateInBatches(values []*model.SqlAdminAPIToken, batchSize int) error
	Save(values ...*model.SqlAdminAPIToken) error
	First() (*model.SqlAdminAPIToken, error)
	Take() (*model.SqlAdminAPIToken, error)
	Last() (*model.SqlAdminAPIToken, error)
	Find() ([]*model.SqlAdminAPIToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlAdminAPIToken, err error)
	FindInBatches(result *[]*model.SqlAdminAPIToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SqlAdminAPIToken) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISqlAdminAPITokenDo
	Assign(attrs ...field.AssignExpr) ISqlAdminAPITokenDo
	Joins(fields ...field.RelationField) ISqlAdminAPITokenDo
	Preload(fields ...field.RelationField) ISqlAdminAPITokenDo
	FirstOrInit() (*model.SqlAdminAPIToken, error)
	FirstOrCreate() (*model.SqlAdminAPIToken, error)
	FindByPage(offset int, limit int) (result []*model.SqlAdminAPIToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISqlAdminAPITokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sqlAdminAPITokenDo) Debug() ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Debug())
}

func (s sqlAdminAPITokenDo) WithContext(ctx context.Context) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sqlAdminAPITokenDo) ReadDB() ISqlAdminAPITokenDo {
	return s.Clauses(dbresolver.Read)
}

func (s sqlAdminAPITokenDo) WriteDB() ISqlAdminAPITokenDo {
	return s.Clauses(dbresolver.Write)
}

func (s sqlAdminAPITokenDo) Session(config *gorm.Session) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Session(config))
}

func (s sqlAdminAPITokenDo) Clauses(conds ...clause.Expression) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sqlAdminAPITokenDo) Returning(value interface{}, columns ...string) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sqlAdminAPITokenDo) Not(conds ...gen.Condition) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sqlAdminAPITokenDo) Or(conds ...gen.Condition) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sqlAdminAPITokenDo) Select(conds ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sqlAdminAPITokenDo) Where(conds ...gen.Condition) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sqlAdminAPITokenDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISqlAdminAPITokenDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sqlAdminAPITokenDo) Order(conds ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sqlAdminAPITokenDo) Distinct(cols ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sqlAdminAPITokenDo) Omit(cols ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sqlAdminAPITokenDo) Join(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sqlAdminAPITokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sqlAdminAPITokenDo) RightJoin(table schema.Tabler, on ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sqlAdminAPITokenDo) Group(cols ...field.Expr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sqlAdminAPITokenDo) Having(conds ...gen.Condition) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sqlAdminAPITokenDo) Limit(limit int) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sqlAdminAPITokenDo) Offset(offset int) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sqlAdminAPITokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sqlAdminAPITokenDo) Unscoped() ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sqlAdminAPITokenDo) Create(values ...*model.SqlAdminAPIToken) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sqlAdminAPITokenDo) CreateInBatches(values []*model.SqlAdminAPIToken, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sqlAdminAPITokenDo) Save(values ...*model.SqlAdminAPIToken) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sqlAdminAPITokenDo) First() (*model.SqlAdminAPIToken, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAPIToken), nil
	}
}

func (s sqlAdminAPITokenDo) Take() (*model.SqlAdminAPIToken, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAPIToken), nil
	}
}

func (s sqlAdminAPITokenDo) Last() (*model.SqlAdminAPIToken, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAPIToken), nil
	}
}

func (s sqlAdminAPITokenDo) Find() ([]*model.SqlAdminAPIToken, error) {
	result, err := s.DO.Find()
	return result.([]*model.SqlAdminAPIToken), err
}

func (s sqlAdminAPITokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlAdminAPIToken, err error) {
	buf := make([]*model.SqlAdminAPIToken, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sqlAdminAPITokenDo) FindInBatches(result *[]*model.SqlAdminAPIToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sqlAdminAPITokenDo) Attrs(attrs ...field.AssignExpr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sqlAdminAPITokenDo) Assign(attrs ...field.AssignExpr) ISqlAdminAPITokenDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sqlAdminAPITokenDo) Joins(fields ...field.RelationField) ISqlAdminAPITokenDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sqlAdminAPITokenDo) Preload(fields ...field.RelationField) ISqlAdminAPITokenDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sqlAdminAPITokenDo) FirstOrInit() (*model.SqlAdminAPIToken, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAPIToken), nil
	}
}

func (s sqlAdminAPITokenDo) FirstOrCreate() (*model.SqlAdminAPIToken, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlAdminAPIToken), nil
	}
}

func (s sqlAdminAPITokenDo) FindByPage(offset int, limit int) (result []*model.SqlAdminAPIToken, count int64, err error) {
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

func (s sqlAdminAPITokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sqlAdminAPITokenDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sqlAdminAPITokenDo) Delete(models ...*model.SqlAdminAPIToken) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sqlAdminAPITokenDo) withDO(do gen.Dao) *sqlAdminAPITokenDo {
	s.DO = *do.(*gen.DO)
	return s
}
