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

func newCollectDocParams(db *gorm.DB, opts ...gen.DOOption) collectDocParams {
	_collectDocParams := collectDocParams{}

	_collectDocParams.collectDocParamsDo.UseDB(db, opts...)
	_collectDocParams.collectDocParamsDo.UseModel(&model.CollectDocParams{})

	tableName := _collectDocParams.collectDocParamsDo.TableName()
	_collectDocParams.ALL = field.NewAsterisk(tableName)
	_collectDocParams.DocParamsID = field.NewString(tableName, "doc_params_id")
	_collectDocParams.CollectDocID = field.NewString(tableName, "collect_doc_id")
	_collectDocParams.Name = field.NewString(tableName, "name")
	_collectDocParams.Desc = field.NewString(tableName, "desc")
	_collectDocParams.Type = field.NewString(tableName, "type")
	_collectDocParams.Must = field.NewString(tableName, "must")
	_collectDocParams.OrderIndex = field.NewInt32(tableName, "order_index")

	_collectDocParams.fillFieldMap()

	return _collectDocParams
}

type collectDocParams struct {
	collectDocParamsDo

	ALL          field.Asterisk
	DocParamsID  field.String
	CollectDocID field.String
	Name         field.String
	Desc         field.String
	Type         field.String
	Must         field.String
	OrderIndex   field.Int32

	fieldMap map[string]field.Expr
}

func (c collectDocParams) Table(newTableName string) *collectDocParams {
	c.collectDocParamsDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collectDocParams) As(alias string) *collectDocParams {
	c.collectDocParamsDo.DO = *(c.collectDocParamsDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collectDocParams) updateTableName(table string) *collectDocParams {
	c.ALL = field.NewAsterisk(table)
	c.DocParamsID = field.NewString(table, "doc_params_id")
	c.CollectDocID = field.NewString(table, "collect_doc_id")
	c.Name = field.NewString(table, "name")
	c.Desc = field.NewString(table, "desc")
	c.Type = field.NewString(table, "type")
	c.Must = field.NewString(table, "must")
	c.OrderIndex = field.NewInt32(table, "order_index")

	c.fillFieldMap()

	return c
}

func (c *collectDocParams) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collectDocParams) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["doc_params_id"] = c.DocParamsID
	c.fieldMap["collect_doc_id"] = c.CollectDocID
	c.fieldMap["name"] = c.Name
	c.fieldMap["desc"] = c.Desc
	c.fieldMap["type"] = c.Type
	c.fieldMap["must"] = c.Must
	c.fieldMap["order_index"] = c.OrderIndex
}

func (c collectDocParams) clone(db *gorm.DB) collectDocParams {
	c.collectDocParamsDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c collectDocParams) replaceDB(db *gorm.DB) collectDocParams {
	c.collectDocParamsDo.ReplaceDB(db)
	return c
}

type collectDocParamsDo struct{ gen.DO }

type ICollectDocParamsDo interface {
	gen.SubQuery
	Debug() ICollectDocParamsDo
	WithContext(ctx context.Context) ICollectDocParamsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICollectDocParamsDo
	WriteDB() ICollectDocParamsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICollectDocParamsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectDocParamsDo
	Not(conds ...gen.Condition) ICollectDocParamsDo
	Or(conds ...gen.Condition) ICollectDocParamsDo
	Select(conds ...field.Expr) ICollectDocParamsDo
	Where(conds ...gen.Condition) ICollectDocParamsDo
	Order(conds ...field.Expr) ICollectDocParamsDo
	Distinct(cols ...field.Expr) ICollectDocParamsDo
	Omit(cols ...field.Expr) ICollectDocParamsDo
	Join(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo
	Group(cols ...field.Expr) ICollectDocParamsDo
	Having(conds ...gen.Condition) ICollectDocParamsDo
	Limit(limit int) ICollectDocParamsDo
	Offset(offset int) ICollectDocParamsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDocParamsDo
	Unscoped() ICollectDocParamsDo
	Create(values ...*model.CollectDocParams) error
	CreateInBatches(values []*model.CollectDocParams, batchSize int) error
	Save(values ...*model.CollectDocParams) error
	First() (*model.CollectDocParams, error)
	Take() (*model.CollectDocParams, error)
	Last() (*model.CollectDocParams, error)
	Find() ([]*model.CollectDocParams, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectDocParams, err error)
	FindInBatches(result *[]*model.CollectDocParams, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CollectDocParams) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectDocParamsDo
	Assign(attrs ...field.AssignExpr) ICollectDocParamsDo
	Joins(fields ...field.RelationField) ICollectDocParamsDo
	Preload(fields ...field.RelationField) ICollectDocParamsDo
	FirstOrInit() (*model.CollectDocParams, error)
	FirstOrCreate() (*model.CollectDocParams, error)
	FindByPage(offset int, limit int) (result []*model.CollectDocParams, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectDocParamsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectDocParamsDo) Debug() ICollectDocParamsDo {
	return c.withDO(c.DO.Debug())
}

func (c collectDocParamsDo) WithContext(ctx context.Context) ICollectDocParamsDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectDocParamsDo) ReadDB() ICollectDocParamsDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectDocParamsDo) WriteDB() ICollectDocParamsDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectDocParamsDo) Session(config *gorm.Session) ICollectDocParamsDo {
	return c.withDO(c.DO.Session(config))
}

func (c collectDocParamsDo) Clauses(conds ...clause.Expression) ICollectDocParamsDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectDocParamsDo) Returning(value interface{}, columns ...string) ICollectDocParamsDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectDocParamsDo) Not(conds ...gen.Condition) ICollectDocParamsDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectDocParamsDo) Or(conds ...gen.Condition) ICollectDocParamsDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectDocParamsDo) Select(conds ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectDocParamsDo) Where(conds ...gen.Condition) ICollectDocParamsDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectDocParamsDo) Order(conds ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectDocParamsDo) Distinct(cols ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectDocParamsDo) Omit(cols ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectDocParamsDo) Join(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectDocParamsDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectDocParamsDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectDocParamsDo) Group(cols ...field.Expr) ICollectDocParamsDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectDocParamsDo) Having(conds ...gen.Condition) ICollectDocParamsDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectDocParamsDo) Limit(limit int) ICollectDocParamsDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectDocParamsDo) Offset(offset int) ICollectDocParamsDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectDocParamsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectDocParamsDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectDocParamsDo) Unscoped() ICollectDocParamsDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectDocParamsDo) Create(values ...*model.CollectDocParams) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectDocParamsDo) CreateInBatches(values []*model.CollectDocParams, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectDocParamsDo) Save(values ...*model.CollectDocParams) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectDocParamsDo) First() (*model.CollectDocParams, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDocParams), nil
	}
}

func (c collectDocParamsDo) Take() (*model.CollectDocParams, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDocParams), nil
	}
}

func (c collectDocParamsDo) Last() (*model.CollectDocParams, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDocParams), nil
	}
}

func (c collectDocParamsDo) Find() ([]*model.CollectDocParams, error) {
	result, err := c.DO.Find()
	return result.([]*model.CollectDocParams), err
}

func (c collectDocParamsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectDocParams, err error) {
	buf := make([]*model.CollectDocParams, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectDocParamsDo) FindInBatches(result *[]*model.CollectDocParams, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectDocParamsDo) Attrs(attrs ...field.AssignExpr) ICollectDocParamsDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectDocParamsDo) Assign(attrs ...field.AssignExpr) ICollectDocParamsDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectDocParamsDo) Joins(fields ...field.RelationField) ICollectDocParamsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectDocParamsDo) Preload(fields ...field.RelationField) ICollectDocParamsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectDocParamsDo) FirstOrInit() (*model.CollectDocParams, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDocParams), nil
	}
}

func (c collectDocParamsDo) FirstOrCreate() (*model.CollectDocParams, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectDocParams), nil
	}
}

func (c collectDocParamsDo) FindByPage(offset int, limit int) (result []*model.CollectDocParams, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectDocParamsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectDocParamsDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c collectDocParamsDo) Delete(models ...*model.CollectDocParams) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *collectDocParamsDo) withDO(do gen.Dao) *collectDocParamsDo {
	c.DO = *do.(*gen.DO)
	return c
}
