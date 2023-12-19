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

func newIssueFlow(db *gorm.DB, opts ...gen.DOOption) issueFlow {
	_issueFlow := issueFlow{}

	_issueFlow.issueFlowDo.UseDB(db, opts...)
	_issueFlow.issueFlowDo.UseModel(&model.IssueFlow{})

	tableName := _issueFlow.issueFlowDo.TableName()
	_issueFlow.ALL = field.NewAsterisk(tableName)
	_issueFlow.IssueFlowID = field.NewString(tableName, "issue_flow_id")
	_issueFlow.IssueStatu = field.NewString(tableName, "issue_statu")
	_issueFlow.Weight = field.NewInt32(tableName, "weight")

	_issueFlow.fillFieldMap()

	return _issueFlow
}

type issueFlow struct {
	issueFlowDo

	ALL         field.Asterisk
	IssueFlowID field.String
	IssueStatu  field.String
	Weight      field.Int32

	fieldMap map[string]field.Expr
}

func (i issueFlow) Table(newTableName string) *issueFlow {
	i.issueFlowDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i issueFlow) As(alias string) *issueFlow {
	i.issueFlowDo.DO = *(i.issueFlowDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *issueFlow) updateTableName(table string) *issueFlow {
	i.ALL = field.NewAsterisk(table)
	i.IssueFlowID = field.NewString(table, "issue_flow_id")
	i.IssueStatu = field.NewString(table, "issue_statu")
	i.Weight = field.NewInt32(table, "weight")

	i.fillFieldMap()

	return i
}

func (i *issueFlow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *issueFlow) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 3)
	i.fieldMap["issue_flow_id"] = i.IssueFlowID
	i.fieldMap["issue_statu"] = i.IssueStatu
	i.fieldMap["weight"] = i.Weight
}

func (i issueFlow) clone(db *gorm.DB) issueFlow {
	i.issueFlowDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i issueFlow) replaceDB(db *gorm.DB) issueFlow {
	i.issueFlowDo.ReplaceDB(db)
	return i
}

type issueFlowDo struct{ gen.DO }

type IIssueFlowDo interface {
	gen.SubQuery
	Debug() IIssueFlowDo
	WithContext(ctx context.Context) IIssueFlowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIssueFlowDo
	WriteDB() IIssueFlowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIssueFlowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIssueFlowDo
	Not(conds ...gen.Condition) IIssueFlowDo
	Or(conds ...gen.Condition) IIssueFlowDo
	Select(conds ...field.Expr) IIssueFlowDo
	Where(conds ...gen.Condition) IIssueFlowDo
	Order(conds ...field.Expr) IIssueFlowDo
	Distinct(cols ...field.Expr) IIssueFlowDo
	Omit(cols ...field.Expr) IIssueFlowDo
	Join(table schema.Tabler, on ...field.Expr) IIssueFlowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIssueFlowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIssueFlowDo
	Group(cols ...field.Expr) IIssueFlowDo
	Having(conds ...gen.Condition) IIssueFlowDo
	Limit(limit int) IIssueFlowDo
	Offset(offset int) IIssueFlowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueFlowDo
	Unscoped() IIssueFlowDo
	Create(values ...*model.IssueFlow) error
	CreateInBatches(values []*model.IssueFlow, batchSize int) error
	Save(values ...*model.IssueFlow) error
	First() (*model.IssueFlow, error)
	Take() (*model.IssueFlow, error)
	Last() (*model.IssueFlow, error)
	Find() ([]*model.IssueFlow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueFlow, err error)
	FindInBatches(result *[]*model.IssueFlow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IssueFlow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIssueFlowDo
	Assign(attrs ...field.AssignExpr) IIssueFlowDo
	Joins(fields ...field.RelationField) IIssueFlowDo
	Preload(fields ...field.RelationField) IIssueFlowDo
	FirstOrInit() (*model.IssueFlow, error)
	FirstOrCreate() (*model.IssueFlow, error)
	FindByPage(offset int, limit int) (result []*model.IssueFlow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIssueFlowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i issueFlowDo) Debug() IIssueFlowDo {
	return i.withDO(i.DO.Debug())
}

func (i issueFlowDo) WithContext(ctx context.Context) IIssueFlowDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i issueFlowDo) ReadDB() IIssueFlowDo {
	return i.Clauses(dbresolver.Read)
}

func (i issueFlowDo) WriteDB() IIssueFlowDo {
	return i.Clauses(dbresolver.Write)
}

func (i issueFlowDo) Session(config *gorm.Session) IIssueFlowDo {
	return i.withDO(i.DO.Session(config))
}

func (i issueFlowDo) Clauses(conds ...clause.Expression) IIssueFlowDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i issueFlowDo) Returning(value interface{}, columns ...string) IIssueFlowDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i issueFlowDo) Not(conds ...gen.Condition) IIssueFlowDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i issueFlowDo) Or(conds ...gen.Condition) IIssueFlowDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i issueFlowDo) Select(conds ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i issueFlowDo) Where(conds ...gen.Condition) IIssueFlowDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i issueFlowDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IIssueFlowDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i issueFlowDo) Order(conds ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i issueFlowDo) Distinct(cols ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i issueFlowDo) Omit(cols ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i issueFlowDo) Join(table schema.Tabler, on ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i issueFlowDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i issueFlowDo) RightJoin(table schema.Tabler, on ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i issueFlowDo) Group(cols ...field.Expr) IIssueFlowDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i issueFlowDo) Having(conds ...gen.Condition) IIssueFlowDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i issueFlowDo) Limit(limit int) IIssueFlowDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i issueFlowDo) Offset(offset int) IIssueFlowDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i issueFlowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueFlowDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i issueFlowDo) Unscoped() IIssueFlowDo {
	return i.withDO(i.DO.Unscoped())
}

func (i issueFlowDo) Create(values ...*model.IssueFlow) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i issueFlowDo) CreateInBatches(values []*model.IssueFlow, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i issueFlowDo) Save(values ...*model.IssueFlow) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i issueFlowDo) First() (*model.IssueFlow, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFlow), nil
	}
}

func (i issueFlowDo) Take() (*model.IssueFlow, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFlow), nil
	}
}

func (i issueFlowDo) Last() (*model.IssueFlow, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFlow), nil
	}
}

func (i issueFlowDo) Find() ([]*model.IssueFlow, error) {
	result, err := i.DO.Find()
	return result.([]*model.IssueFlow), err
}

func (i issueFlowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueFlow, err error) {
	buf := make([]*model.IssueFlow, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i issueFlowDo) FindInBatches(result *[]*model.IssueFlow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i issueFlowDo) Attrs(attrs ...field.AssignExpr) IIssueFlowDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i issueFlowDo) Assign(attrs ...field.AssignExpr) IIssueFlowDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i issueFlowDo) Joins(fields ...field.RelationField) IIssueFlowDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i issueFlowDo) Preload(fields ...field.RelationField) IIssueFlowDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i issueFlowDo) FirstOrInit() (*model.IssueFlow, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFlow), nil
	}
}

func (i issueFlowDo) FirstOrCreate() (*model.IssueFlow, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueFlow), nil
	}
}

func (i issueFlowDo) FindByPage(offset int, limit int) (result []*model.IssueFlow, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i issueFlowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i issueFlowDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i issueFlowDo) Delete(models ...*model.IssueFlow) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *issueFlowDo) withDO(do gen.Dao) *issueFlowDo {
	i.DO = *do.(*gen.DO)
	return i
}
