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

func newPublishReqBefore(db *gorm.DB, opts ...gen.DOOption) publishReqBefore {
	_publishReqBefore := publishReqBefore{}

	_publishReqBefore.publishReqBeforeDo.UseDB(db, opts...)
	_publishReqBefore.publishReqBeforeDo.UseModel(&model.PublishReqBefore{})

	tableName := _publishReqBefore.publishReqBeforeDo.TableName()
	_publishReqBefore.ALL = field.NewAsterisk(tableName)
	_publishReqBefore.PublishReqBeforeID = field.NewString(tableName, "publish_req_before_id")
	_publishReqBefore.SourceReqID = field.NewString(tableName, "source_req_id")
	_publishReqBefore.ReqID = field.NewString(tableName, "req_id")
	_publishReqBefore.CreateTime = field.NewString(tableName, "create_time")

	_publishReqBefore.fillFieldMap()

	return _publishReqBefore
}

type publishReqBefore struct {
	publishReqBeforeDo

	ALL                field.Asterisk
	PublishReqBeforeID field.String
	SourceReqID        field.String
	ReqID              field.String
	CreateTime         field.String

	fieldMap map[string]field.Expr
}

func (p publishReqBefore) Table(newTableName string) *publishReqBefore {
	p.publishReqBeforeDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publishReqBefore) As(alias string) *publishReqBefore {
	p.publishReqBeforeDo.DO = *(p.publishReqBeforeDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publishReqBefore) updateTableName(table string) *publishReqBefore {
	p.ALL = field.NewAsterisk(table)
	p.PublishReqBeforeID = field.NewString(table, "publish_req_before_id")
	p.SourceReqID = field.NewString(table, "source_req_id")
	p.ReqID = field.NewString(table, "req_id")
	p.CreateTime = field.NewString(table, "create_time")

	p.fillFieldMap()

	return p
}

func (p *publishReqBefore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publishReqBefore) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["publish_req_before_id"] = p.PublishReqBeforeID
	p.fieldMap["source_req_id"] = p.SourceReqID
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["create_time"] = p.CreateTime
}

func (p publishReqBefore) clone(db *gorm.DB) publishReqBefore {
	p.publishReqBeforeDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publishReqBefore) replaceDB(db *gorm.DB) publishReqBefore {
	p.publishReqBeforeDo.ReplaceDB(db)
	return p
}

type publishReqBeforeDo struct{ gen.DO }

type IPublishReqBeforeDo interface {
	gen.SubQuery
	Debug() IPublishReqBeforeDo
	WithContext(ctx context.Context) IPublishReqBeforeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublishReqBeforeDo
	WriteDB() IPublishReqBeforeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublishReqBeforeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublishReqBeforeDo
	Not(conds ...gen.Condition) IPublishReqBeforeDo
	Or(conds ...gen.Condition) IPublishReqBeforeDo
	Select(conds ...field.Expr) IPublishReqBeforeDo
	Where(conds ...gen.Condition) IPublishReqBeforeDo
	Order(conds ...field.Expr) IPublishReqBeforeDo
	Distinct(cols ...field.Expr) IPublishReqBeforeDo
	Omit(cols ...field.Expr) IPublishReqBeforeDo
	Join(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo
	Group(cols ...field.Expr) IPublishReqBeforeDo
	Having(conds ...gen.Condition) IPublishReqBeforeDo
	Limit(limit int) IPublishReqBeforeDo
	Offset(offset int) IPublishReqBeforeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqBeforeDo
	Unscoped() IPublishReqBeforeDo
	Create(values ...*model.PublishReqBefore) error
	CreateInBatches(values []*model.PublishReqBefore, batchSize int) error
	Save(values ...*model.PublishReqBefore) error
	First() (*model.PublishReqBefore, error)
	Take() (*model.PublishReqBefore, error)
	Last() (*model.PublishReqBefore, error)
	Find() ([]*model.PublishReqBefore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqBefore, err error)
	FindInBatches(result *[]*model.PublishReqBefore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PublishReqBefore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublishReqBeforeDo
	Assign(attrs ...field.AssignExpr) IPublishReqBeforeDo
	Joins(fields ...field.RelationField) IPublishReqBeforeDo
	Preload(fields ...field.RelationField) IPublishReqBeforeDo
	FirstOrInit() (*model.PublishReqBefore, error)
	FirstOrCreate() (*model.PublishReqBefore, error)
	FindByPage(offset int, limit int) (result []*model.PublishReqBefore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublishReqBeforeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publishReqBeforeDo) Debug() IPublishReqBeforeDo {
	return p.withDO(p.DO.Debug())
}

func (p publishReqBeforeDo) WithContext(ctx context.Context) IPublishReqBeforeDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishReqBeforeDo) ReadDB() IPublishReqBeforeDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishReqBeforeDo) WriteDB() IPublishReqBeforeDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishReqBeforeDo) Session(config *gorm.Session) IPublishReqBeforeDo {
	return p.withDO(p.DO.Session(config))
}

func (p publishReqBeforeDo) Clauses(conds ...clause.Expression) IPublishReqBeforeDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishReqBeforeDo) Returning(value interface{}, columns ...string) IPublishReqBeforeDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishReqBeforeDo) Not(conds ...gen.Condition) IPublishReqBeforeDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishReqBeforeDo) Or(conds ...gen.Condition) IPublishReqBeforeDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishReqBeforeDo) Select(conds ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishReqBeforeDo) Where(conds ...gen.Condition) IPublishReqBeforeDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishReqBeforeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPublishReqBeforeDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishReqBeforeDo) Order(conds ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishReqBeforeDo) Distinct(cols ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishReqBeforeDo) Omit(cols ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishReqBeforeDo) Join(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishReqBeforeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishReqBeforeDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishReqBeforeDo) Group(cols ...field.Expr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishReqBeforeDo) Having(conds ...gen.Condition) IPublishReqBeforeDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishReqBeforeDo) Limit(limit int) IPublishReqBeforeDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishReqBeforeDo) Offset(offset int) IPublishReqBeforeDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishReqBeforeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqBeforeDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishReqBeforeDo) Unscoped() IPublishReqBeforeDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishReqBeforeDo) Create(values ...*model.PublishReqBefore) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishReqBeforeDo) CreateInBatches(values []*model.PublishReqBefore, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishReqBeforeDo) Save(values ...*model.PublishReqBefore) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishReqBeforeDo) First() (*model.PublishReqBefore, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqBefore), nil
	}
}

func (p publishReqBeforeDo) Take() (*model.PublishReqBefore, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqBefore), nil
	}
}

func (p publishReqBeforeDo) Last() (*model.PublishReqBefore, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqBefore), nil
	}
}

func (p publishReqBeforeDo) Find() ([]*model.PublishReqBefore, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublishReqBefore), err
}

func (p publishReqBeforeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqBefore, err error) {
	buf := make([]*model.PublishReqBefore, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishReqBeforeDo) FindInBatches(result *[]*model.PublishReqBefore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishReqBeforeDo) Attrs(attrs ...field.AssignExpr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishReqBeforeDo) Assign(attrs ...field.AssignExpr) IPublishReqBeforeDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishReqBeforeDo) Joins(fields ...field.RelationField) IPublishReqBeforeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishReqBeforeDo) Preload(fields ...field.RelationField) IPublishReqBeforeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishReqBeforeDo) FirstOrInit() (*model.PublishReqBefore, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqBefore), nil
	}
}

func (p publishReqBeforeDo) FirstOrCreate() (*model.PublishReqBefore, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqBefore), nil
	}
}

func (p publishReqBeforeDo) FindByPage(offset int, limit int) (result []*model.PublishReqBefore, count int64, err error) {
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

func (p publishReqBeforeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishReqBeforeDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishReqBeforeDo) Delete(models ...*model.PublishReqBefore) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishReqBeforeDo) withDO(do gen.Dao) *publishReqBeforeDo {
	p.DO = *do.(*gen.DO)
	return p
}