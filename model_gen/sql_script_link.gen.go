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

func newSqlScriptLink(db *gorm.DB, opts ...gen.DOOption) sqlScriptLink {
	_sqlScriptLink := sqlScriptLink{}

	_sqlScriptLink.sqlScriptLinkDo.UseDB(db, opts...)
	_sqlScriptLink.sqlScriptLinkDo.UseModel(&model.SqlScriptLink{})

	tableName := _sqlScriptLink.sqlScriptLinkDo.TableName()
	_sqlScriptLink.ALL = field.NewAsterisk(tableName)
	_sqlScriptLink.LinkID = field.NewString(tableName, "link_id")
	_sqlScriptLink.Code = field.NewString(tableName, "code")
	_sqlScriptLink.IssueTitle = field.NewString(tableName, "issue_title")
	_sqlScriptLink.URL = field.NewString(tableName, "url")
	_sqlScriptLink.IssueKey = field.NewString(tableName, "issue_key")
	_sqlScriptLink.Updated = field.NewString(tableName, "updated")

	_sqlScriptLink.fillFieldMap()

	return _sqlScriptLink
}

type sqlScriptLink struct {
	sqlScriptLinkDo

	ALL        field.Asterisk
	LinkID     field.String
	Code       field.String
	IssueTitle field.String
	URL        field.String
	IssueKey   field.String
	Updated    field.String

	fieldMap map[string]field.Expr
}

func (s sqlScriptLink) Table(newTableName string) *sqlScriptLink {
	s.sqlScriptLinkDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sqlScriptLink) As(alias string) *sqlScriptLink {
	s.sqlScriptLinkDo.DO = *(s.sqlScriptLinkDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sqlScriptLink) updateTableName(table string) *sqlScriptLink {
	s.ALL = field.NewAsterisk(table)
	s.LinkID = field.NewString(table, "link_id")
	s.Code = field.NewString(table, "code")
	s.IssueTitle = field.NewString(table, "issue_title")
	s.URL = field.NewString(table, "url")
	s.IssueKey = field.NewString(table, "issue_key")
	s.Updated = field.NewString(table, "updated")

	s.fillFieldMap()

	return s
}

func (s *sqlScriptLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sqlScriptLink) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["link_id"] = s.LinkID
	s.fieldMap["code"] = s.Code
	s.fieldMap["issue_title"] = s.IssueTitle
	s.fieldMap["url"] = s.URL
	s.fieldMap["issue_key"] = s.IssueKey
	s.fieldMap["updated"] = s.Updated
}

func (s sqlScriptLink) clone(db *gorm.DB) sqlScriptLink {
	s.sqlScriptLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sqlScriptLink) replaceDB(db *gorm.DB) sqlScriptLink {
	s.sqlScriptLinkDo.ReplaceDB(db)
	return s
}

type sqlScriptLinkDo struct{ gen.DO }

type ISqlScriptLinkDo interface {
	gen.SubQuery
	Debug() ISqlScriptLinkDo
	WithContext(ctx context.Context) ISqlScriptLinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISqlScriptLinkDo
	WriteDB() ISqlScriptLinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISqlScriptLinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISqlScriptLinkDo
	Not(conds ...gen.Condition) ISqlScriptLinkDo
	Or(conds ...gen.Condition) ISqlScriptLinkDo
	Select(conds ...field.Expr) ISqlScriptLinkDo
	Where(conds ...gen.Condition) ISqlScriptLinkDo
	Order(conds ...field.Expr) ISqlScriptLinkDo
	Distinct(cols ...field.Expr) ISqlScriptLinkDo
	Omit(cols ...field.Expr) ISqlScriptLinkDo
	Join(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo
	Group(cols ...field.Expr) ISqlScriptLinkDo
	Having(conds ...gen.Condition) ISqlScriptLinkDo
	Limit(limit int) ISqlScriptLinkDo
	Offset(offset int) ISqlScriptLinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlScriptLinkDo
	Unscoped() ISqlScriptLinkDo
	Create(values ...*model.SqlScriptLink) error
	CreateInBatches(values []*model.SqlScriptLink, batchSize int) error
	Save(values ...*model.SqlScriptLink) error
	First() (*model.SqlScriptLink, error)
	Take() (*model.SqlScriptLink, error)
	Last() (*model.SqlScriptLink, error)
	Find() ([]*model.SqlScriptLink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlScriptLink, err error)
	FindInBatches(result *[]*model.SqlScriptLink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SqlScriptLink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISqlScriptLinkDo
	Assign(attrs ...field.AssignExpr) ISqlScriptLinkDo
	Joins(fields ...field.RelationField) ISqlScriptLinkDo
	Preload(fields ...field.RelationField) ISqlScriptLinkDo
	FirstOrInit() (*model.SqlScriptLink, error)
	FirstOrCreate() (*model.SqlScriptLink, error)
	FindByPage(offset int, limit int) (result []*model.SqlScriptLink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISqlScriptLinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sqlScriptLinkDo) Debug() ISqlScriptLinkDo {
	return s.withDO(s.DO.Debug())
}

func (s sqlScriptLinkDo) WithContext(ctx context.Context) ISqlScriptLinkDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sqlScriptLinkDo) ReadDB() ISqlScriptLinkDo {
	return s.Clauses(dbresolver.Read)
}

func (s sqlScriptLinkDo) WriteDB() ISqlScriptLinkDo {
	return s.Clauses(dbresolver.Write)
}

func (s sqlScriptLinkDo) Session(config *gorm.Session) ISqlScriptLinkDo {
	return s.withDO(s.DO.Session(config))
}

func (s sqlScriptLinkDo) Clauses(conds ...clause.Expression) ISqlScriptLinkDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sqlScriptLinkDo) Returning(value interface{}, columns ...string) ISqlScriptLinkDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sqlScriptLinkDo) Not(conds ...gen.Condition) ISqlScriptLinkDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sqlScriptLinkDo) Or(conds ...gen.Condition) ISqlScriptLinkDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sqlScriptLinkDo) Select(conds ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sqlScriptLinkDo) Where(conds ...gen.Condition) ISqlScriptLinkDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sqlScriptLinkDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISqlScriptLinkDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sqlScriptLinkDo) Order(conds ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sqlScriptLinkDo) Distinct(cols ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sqlScriptLinkDo) Omit(cols ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sqlScriptLinkDo) Join(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sqlScriptLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sqlScriptLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sqlScriptLinkDo) Group(cols ...field.Expr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sqlScriptLinkDo) Having(conds ...gen.Condition) ISqlScriptLinkDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sqlScriptLinkDo) Limit(limit int) ISqlScriptLinkDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sqlScriptLinkDo) Offset(offset int) ISqlScriptLinkDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sqlScriptLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISqlScriptLinkDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sqlScriptLinkDo) Unscoped() ISqlScriptLinkDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sqlScriptLinkDo) Create(values ...*model.SqlScriptLink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sqlScriptLinkDo) CreateInBatches(values []*model.SqlScriptLink, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sqlScriptLinkDo) Save(values ...*model.SqlScriptLink) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sqlScriptLinkDo) First() (*model.SqlScriptLink, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlScriptLink), nil
	}
}

func (s sqlScriptLinkDo) Take() (*model.SqlScriptLink, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlScriptLink), nil
	}
}

func (s sqlScriptLinkDo) Last() (*model.SqlScriptLink, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlScriptLink), nil
	}
}

func (s sqlScriptLinkDo) Find() ([]*model.SqlScriptLink, error) {
	result, err := s.DO.Find()
	return result.([]*model.SqlScriptLink), err
}

func (s sqlScriptLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SqlScriptLink, err error) {
	buf := make([]*model.SqlScriptLink, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sqlScriptLinkDo) FindInBatches(result *[]*model.SqlScriptLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sqlScriptLinkDo) Attrs(attrs ...field.AssignExpr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sqlScriptLinkDo) Assign(attrs ...field.AssignExpr) ISqlScriptLinkDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sqlScriptLinkDo) Joins(fields ...field.RelationField) ISqlScriptLinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sqlScriptLinkDo) Preload(fields ...field.RelationField) ISqlScriptLinkDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sqlScriptLinkDo) FirstOrInit() (*model.SqlScriptLink, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlScriptLink), nil
	}
}

func (s sqlScriptLinkDo) FirstOrCreate() (*model.SqlScriptLink, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SqlScriptLink), nil
	}
}

func (s sqlScriptLinkDo) FindByPage(offset int, limit int) (result []*model.SqlScriptLink, count int64, err error) {
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

func (s sqlScriptLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sqlScriptLinkDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sqlScriptLinkDo) Delete(models ...*model.SqlScriptLink) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sqlScriptLinkDo) withDO(do gen.Dao) *sqlScriptLinkDo {
	s.DO = *do.(*gen.DO)
	return s
}
