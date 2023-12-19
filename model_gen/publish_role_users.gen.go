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

func newPublishRoleUsers(db *gorm.DB, opts ...gen.DOOption) publishRoleUsers {
	_publishRoleUsers := publishRoleUsers{}

	_publishRoleUsers.publishRoleUsersDo.UseDB(db, opts...)
	_publishRoleUsers.publishRoleUsersDo.UseModel(&model.PublishRoleUsers{})

	tableName := _publishRoleUsers.publishRoleUsersDo.TableName()
	_publishRoleUsers.ALL = field.NewAsterisk(tableName)
	_publishRoleUsers.PublishRoleUsersID = field.NewString(tableName, "publish_role_users_id")
	_publishRoleUsers.ReqID = field.NewString(tableName, "req_id")
	_publishRoleUsers.UserID = field.NewString(tableName, "user_id")
	_publishRoleUsers.Type = field.NewString(tableName, "type")

	_publishRoleUsers.fillFieldMap()

	return _publishRoleUsers
}

type publishRoleUsers struct {
	publishRoleUsersDo

	ALL                field.Asterisk
	PublishRoleUsersID field.String
	ReqID              field.String
	UserID             field.String // 用户ID
	Type               field.String

	fieldMap map[string]field.Expr
}

func (p publishRoleUsers) Table(newTableName string) *publishRoleUsers {
	p.publishRoleUsersDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publishRoleUsers) As(alias string) *publishRoleUsers {
	p.publishRoleUsersDo.DO = *(p.publishRoleUsersDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publishRoleUsers) updateTableName(table string) *publishRoleUsers {
	p.ALL = field.NewAsterisk(table)
	p.PublishRoleUsersID = field.NewString(table, "publish_role_users_id")
	p.ReqID = field.NewString(table, "req_id")
	p.UserID = field.NewString(table, "user_id")
	p.Type = field.NewString(table, "type")

	p.fillFieldMap()

	return p
}

func (p *publishRoleUsers) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publishRoleUsers) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["publish_role_users_id"] = p.PublishRoleUsersID
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["type"] = p.Type
}

func (p publishRoleUsers) clone(db *gorm.DB) publishRoleUsers {
	p.publishRoleUsersDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publishRoleUsers) replaceDB(db *gorm.DB) publishRoleUsers {
	p.publishRoleUsersDo.ReplaceDB(db)
	return p
}

type publishRoleUsersDo struct{ gen.DO }

type IPublishRoleUsersDo interface {
	gen.SubQuery
	Debug() IPublishRoleUsersDo
	WithContext(ctx context.Context) IPublishRoleUsersDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublishRoleUsersDo
	WriteDB() IPublishRoleUsersDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublishRoleUsersDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublishRoleUsersDo
	Not(conds ...gen.Condition) IPublishRoleUsersDo
	Or(conds ...gen.Condition) IPublishRoleUsersDo
	Select(conds ...field.Expr) IPublishRoleUsersDo
	Where(conds ...gen.Condition) IPublishRoleUsersDo
	Order(conds ...field.Expr) IPublishRoleUsersDo
	Distinct(cols ...field.Expr) IPublishRoleUsersDo
	Omit(cols ...field.Expr) IPublishRoleUsersDo
	Join(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo
	Group(cols ...field.Expr) IPublishRoleUsersDo
	Having(conds ...gen.Condition) IPublishRoleUsersDo
	Limit(limit int) IPublishRoleUsersDo
	Offset(offset int) IPublishRoleUsersDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishRoleUsersDo
	Unscoped() IPublishRoleUsersDo
	Create(values ...*model.PublishRoleUsers) error
	CreateInBatches(values []*model.PublishRoleUsers, batchSize int) error
	Save(values ...*model.PublishRoleUsers) error
	First() (*model.PublishRoleUsers, error)
	Take() (*model.PublishRoleUsers, error)
	Last() (*model.PublishRoleUsers, error)
	Find() ([]*model.PublishRoleUsers, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishRoleUsers, err error)
	FindInBatches(result *[]*model.PublishRoleUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PublishRoleUsers) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublishRoleUsersDo
	Assign(attrs ...field.AssignExpr) IPublishRoleUsersDo
	Joins(fields ...field.RelationField) IPublishRoleUsersDo
	Preload(fields ...field.RelationField) IPublishRoleUsersDo
	FirstOrInit() (*model.PublishRoleUsers, error)
	FirstOrCreate() (*model.PublishRoleUsers, error)
	FindByPage(offset int, limit int) (result []*model.PublishRoleUsers, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublishRoleUsersDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publishRoleUsersDo) Debug() IPublishRoleUsersDo {
	return p.withDO(p.DO.Debug())
}

func (p publishRoleUsersDo) WithContext(ctx context.Context) IPublishRoleUsersDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishRoleUsersDo) ReadDB() IPublishRoleUsersDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishRoleUsersDo) WriteDB() IPublishRoleUsersDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishRoleUsersDo) Session(config *gorm.Session) IPublishRoleUsersDo {
	return p.withDO(p.DO.Session(config))
}

func (p publishRoleUsersDo) Clauses(conds ...clause.Expression) IPublishRoleUsersDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishRoleUsersDo) Returning(value interface{}, columns ...string) IPublishRoleUsersDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishRoleUsersDo) Not(conds ...gen.Condition) IPublishRoleUsersDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishRoleUsersDo) Or(conds ...gen.Condition) IPublishRoleUsersDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishRoleUsersDo) Select(conds ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishRoleUsersDo) Where(conds ...gen.Condition) IPublishRoleUsersDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishRoleUsersDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPublishRoleUsersDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishRoleUsersDo) Order(conds ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishRoleUsersDo) Distinct(cols ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishRoleUsersDo) Omit(cols ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishRoleUsersDo) Join(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishRoleUsersDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishRoleUsersDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishRoleUsersDo) Group(cols ...field.Expr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishRoleUsersDo) Having(conds ...gen.Condition) IPublishRoleUsersDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishRoleUsersDo) Limit(limit int) IPublishRoleUsersDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishRoleUsersDo) Offset(offset int) IPublishRoleUsersDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishRoleUsersDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishRoleUsersDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishRoleUsersDo) Unscoped() IPublishRoleUsersDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishRoleUsersDo) Create(values ...*model.PublishRoleUsers) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishRoleUsersDo) CreateInBatches(values []*model.PublishRoleUsers, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishRoleUsersDo) Save(values ...*model.PublishRoleUsers) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishRoleUsersDo) First() (*model.PublishRoleUsers, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRoleUsers), nil
	}
}

func (p publishRoleUsersDo) Take() (*model.PublishRoleUsers, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRoleUsers), nil
	}
}

func (p publishRoleUsersDo) Last() (*model.PublishRoleUsers, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRoleUsers), nil
	}
}

func (p publishRoleUsersDo) Find() ([]*model.PublishRoleUsers, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublishRoleUsers), err
}

func (p publishRoleUsersDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishRoleUsers, err error) {
	buf := make([]*model.PublishRoleUsers, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishRoleUsersDo) FindInBatches(result *[]*model.PublishRoleUsers, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishRoleUsersDo) Attrs(attrs ...field.AssignExpr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishRoleUsersDo) Assign(attrs ...field.AssignExpr) IPublishRoleUsersDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishRoleUsersDo) Joins(fields ...field.RelationField) IPublishRoleUsersDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishRoleUsersDo) Preload(fields ...field.RelationField) IPublishRoleUsersDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishRoleUsersDo) FirstOrInit() (*model.PublishRoleUsers, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRoleUsers), nil
	}
}

func (p publishRoleUsersDo) FirstOrCreate() (*model.PublishRoleUsers, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishRoleUsers), nil
	}
}

func (p publishRoleUsersDo) FindByPage(offset int, limit int) (result []*model.PublishRoleUsers, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p publishRoleUsersDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishRoleUsersDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishRoleUsersDo) Delete(models ...*model.PublishRoleUsers) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishRoleUsersDo) withDO(do gen.Dao) *publishRoleUsersDo {
	p.DO = *do.(*gen.DO)
	return p
}
