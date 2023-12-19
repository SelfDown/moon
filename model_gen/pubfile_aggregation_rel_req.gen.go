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

func newPubfileAggregationRelReq(db *gorm.DB, opts ...gen.DOOption) pubfileAggregationRelReq {
	_pubfileAggregationRelReq := pubfileAggregationRelReq{}

	_pubfileAggregationRelReq.pubfileAggregationRelReqDo.UseDB(db, opts...)
	_pubfileAggregationRelReq.pubfileAggregationRelReqDo.UseModel(&model.PubfileAggregationRelReq{})

	tableName := _pubfileAggregationRelReq.pubfileAggregationRelReqDo.TableName()
	_pubfileAggregationRelReq.ALL = field.NewAsterisk(tableName)
	_pubfileAggregationRelReq.AggRelReqID = field.NewString(tableName, "agg_rel_req_id")
	_pubfileAggregationRelReq.ReqID = field.NewString(tableName, "req_id")
	_pubfileAggregationRelReq.AggID = field.NewString(tableName, "agg_id")

	_pubfileAggregationRelReq.fillFieldMap()

	return _pubfileAggregationRelReq
}

type pubfileAggregationRelReq struct {
	pubfileAggregationRelReqDo

	ALL         field.Asterisk
	AggRelReqID field.String // ID
	ReqID       field.String // 确认单申请ID
	AggID       field.String // ID

	fieldMap map[string]field.Expr
}

func (p pubfileAggregationRelReq) Table(newTableName string) *pubfileAggregationRelReq {
	p.pubfileAggregationRelReqDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pubfileAggregationRelReq) As(alias string) *pubfileAggregationRelReq {
	p.pubfileAggregationRelReqDo.DO = *(p.pubfileAggregationRelReqDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pubfileAggregationRelReq) updateTableName(table string) *pubfileAggregationRelReq {
	p.ALL = field.NewAsterisk(table)
	p.AggRelReqID = field.NewString(table, "agg_rel_req_id")
	p.ReqID = field.NewString(table, "req_id")
	p.AggID = field.NewString(table, "agg_id")

	p.fillFieldMap()

	return p
}

func (p *pubfileAggregationRelReq) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pubfileAggregationRelReq) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["agg_rel_req_id"] = p.AggRelReqID
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["agg_id"] = p.AggID
}

func (p pubfileAggregationRelReq) clone(db *gorm.DB) pubfileAggregationRelReq {
	p.pubfileAggregationRelReqDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pubfileAggregationRelReq) replaceDB(db *gorm.DB) pubfileAggregationRelReq {
	p.pubfileAggregationRelReqDo.ReplaceDB(db)
	return p
}

type pubfileAggregationRelReqDo struct{ gen.DO }

type IPubfileAggregationRelReqDo interface {
	gen.SubQuery
	Debug() IPubfileAggregationRelReqDo
	WithContext(ctx context.Context) IPubfileAggregationRelReqDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPubfileAggregationRelReqDo
	WriteDB() IPubfileAggregationRelReqDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPubfileAggregationRelReqDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPubfileAggregationRelReqDo
	Not(conds ...gen.Condition) IPubfileAggregationRelReqDo
	Or(conds ...gen.Condition) IPubfileAggregationRelReqDo
	Select(conds ...field.Expr) IPubfileAggregationRelReqDo
	Where(conds ...gen.Condition) IPubfileAggregationRelReqDo
	Order(conds ...field.Expr) IPubfileAggregationRelReqDo
	Distinct(cols ...field.Expr) IPubfileAggregationRelReqDo
	Omit(cols ...field.Expr) IPubfileAggregationRelReqDo
	Join(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo
	Group(cols ...field.Expr) IPubfileAggregationRelReqDo
	Having(conds ...gen.Condition) IPubfileAggregationRelReqDo
	Limit(limit int) IPubfileAggregationRelReqDo
	Offset(offset int) IPubfileAggregationRelReqDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPubfileAggregationRelReqDo
	Unscoped() IPubfileAggregationRelReqDo
	Create(values ...*model.PubfileAggregationRelReq) error
	CreateInBatches(values []*model.PubfileAggregationRelReq, batchSize int) error
	Save(values ...*model.PubfileAggregationRelReq) error
	First() (*model.PubfileAggregationRelReq, error)
	Take() (*model.PubfileAggregationRelReq, error)
	Last() (*model.PubfileAggregationRelReq, error)
	Find() ([]*model.PubfileAggregationRelReq, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PubfileAggregationRelReq, err error)
	FindInBatches(result *[]*model.PubfileAggregationRelReq, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PubfileAggregationRelReq) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPubfileAggregationRelReqDo
	Assign(attrs ...field.AssignExpr) IPubfileAggregationRelReqDo
	Joins(fields ...field.RelationField) IPubfileAggregationRelReqDo
	Preload(fields ...field.RelationField) IPubfileAggregationRelReqDo
	FirstOrInit() (*model.PubfileAggregationRelReq, error)
	FirstOrCreate() (*model.PubfileAggregationRelReq, error)
	FindByPage(offset int, limit int) (result []*model.PubfileAggregationRelReq, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPubfileAggregationRelReqDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pubfileAggregationRelReqDo) Debug() IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Debug())
}

func (p pubfileAggregationRelReqDo) WithContext(ctx context.Context) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pubfileAggregationRelReqDo) ReadDB() IPubfileAggregationRelReqDo {
	return p.Clauses(dbresolver.Read)
}

func (p pubfileAggregationRelReqDo) WriteDB() IPubfileAggregationRelReqDo {
	return p.Clauses(dbresolver.Write)
}

func (p pubfileAggregationRelReqDo) Session(config *gorm.Session) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Session(config))
}

func (p pubfileAggregationRelReqDo) Clauses(conds ...clause.Expression) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pubfileAggregationRelReqDo) Returning(value interface{}, columns ...string) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pubfileAggregationRelReqDo) Not(conds ...gen.Condition) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pubfileAggregationRelReqDo) Or(conds ...gen.Condition) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pubfileAggregationRelReqDo) Select(conds ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pubfileAggregationRelReqDo) Where(conds ...gen.Condition) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pubfileAggregationRelReqDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPubfileAggregationRelReqDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pubfileAggregationRelReqDo) Order(conds ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pubfileAggregationRelReqDo) Distinct(cols ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pubfileAggregationRelReqDo) Omit(cols ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pubfileAggregationRelReqDo) Join(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pubfileAggregationRelReqDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pubfileAggregationRelReqDo) RightJoin(table schema.Tabler, on ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pubfileAggregationRelReqDo) Group(cols ...field.Expr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pubfileAggregationRelReqDo) Having(conds ...gen.Condition) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pubfileAggregationRelReqDo) Limit(limit int) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pubfileAggregationRelReqDo) Offset(offset int) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pubfileAggregationRelReqDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pubfileAggregationRelReqDo) Unscoped() IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pubfileAggregationRelReqDo) Create(values ...*model.PubfileAggregationRelReq) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pubfileAggregationRelReqDo) CreateInBatches(values []*model.PubfileAggregationRelReq, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pubfileAggregationRelReqDo) Save(values ...*model.PubfileAggregationRelReq) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pubfileAggregationRelReqDo) First() (*model.PubfileAggregationRelReq, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggregationRelReq), nil
	}
}

func (p pubfileAggregationRelReqDo) Take() (*model.PubfileAggregationRelReq, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggregationRelReq), nil
	}
}

func (p pubfileAggregationRelReqDo) Last() (*model.PubfileAggregationRelReq, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggregationRelReq), nil
	}
}

func (p pubfileAggregationRelReqDo) Find() ([]*model.PubfileAggregationRelReq, error) {
	result, err := p.DO.Find()
	return result.([]*model.PubfileAggregationRelReq), err
}

func (p pubfileAggregationRelReqDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PubfileAggregationRelReq, err error) {
	buf := make([]*model.PubfileAggregationRelReq, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pubfileAggregationRelReqDo) FindInBatches(result *[]*model.PubfileAggregationRelReq, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pubfileAggregationRelReqDo) Attrs(attrs ...field.AssignExpr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pubfileAggregationRelReqDo) Assign(attrs ...field.AssignExpr) IPubfileAggregationRelReqDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pubfileAggregationRelReqDo) Joins(fields ...field.RelationField) IPubfileAggregationRelReqDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pubfileAggregationRelReqDo) Preload(fields ...field.RelationField) IPubfileAggregationRelReqDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pubfileAggregationRelReqDo) FirstOrInit() (*model.PubfileAggregationRelReq, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggregationRelReq), nil
	}
}

func (p pubfileAggregationRelReqDo) FirstOrCreate() (*model.PubfileAggregationRelReq, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggregationRelReq), nil
	}
}

func (p pubfileAggregationRelReqDo) FindByPage(offset int, limit int) (result []*model.PubfileAggregationRelReq, count int64, err error) {
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

func (p pubfileAggregationRelReqDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pubfileAggregationRelReqDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pubfileAggregationRelReqDo) Delete(models ...*model.PubfileAggregationRelReq) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pubfileAggregationRelReqDo) withDO(do gen.Dao) *pubfileAggregationRelReqDo {
	p.DO = *do.(*gen.DO)
	return p
}
