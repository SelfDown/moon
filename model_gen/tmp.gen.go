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

func newTmp(db *gorm.DB, opts ...gen.DOOption) tmp {
	_tmp := tmp{}

	_tmp.tmpDo.UseDB(db, opts...)
	_tmp.tmpDo.UseModel(&model.Tmp{})

	tableName := _tmp.tmpDo.TableName()
	_tmp.ALL = field.NewAsterisk(tableName)
	_tmp.MonitorItemkeyDescID = field.NewString(tableName, "monitor_itemkey_desc_id")
	_tmp.Itemkey = field.NewString(tableName, "itemkey")
	_tmp.ItemkeyDesc = field.NewString(tableName, "itemkey_desc")

	_tmp.fillFieldMap()

	return _tmp
}

type tmp struct {
	tmpDo

	ALL                  field.Asterisk
	MonitorItemkeyDescID field.String
	Itemkey              field.String
	ItemkeyDesc          field.String

	fieldMap map[string]field.Expr
}

func (t tmp) Table(newTableName string) *tmp {
	t.tmpDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tmp) As(alias string) *tmp {
	t.tmpDo.DO = *(t.tmpDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tmp) updateTableName(table string) *tmp {
	t.ALL = field.NewAsterisk(table)
	t.MonitorItemkeyDescID = field.NewString(table, "monitor_itemkey_desc_id")
	t.Itemkey = field.NewString(table, "itemkey")
	t.ItemkeyDesc = field.NewString(table, "itemkey_desc")

	t.fillFieldMap()

	return t
}

func (t *tmp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tmp) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["monitor_itemkey_desc_id"] = t.MonitorItemkeyDescID
	t.fieldMap["itemkey"] = t.Itemkey
	t.fieldMap["itemkey_desc"] = t.ItemkeyDesc
}

func (t tmp) clone(db *gorm.DB) tmp {
	t.tmpDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tmp) replaceDB(db *gorm.DB) tmp {
	t.tmpDo.ReplaceDB(db)
	return t
}

type tmpDo struct{ gen.DO }

type ITmpDo interface {
	gen.SubQuery
	Debug() ITmpDo
	WithContext(ctx context.Context) ITmpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITmpDo
	WriteDB() ITmpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITmpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITmpDo
	Not(conds ...gen.Condition) ITmpDo
	Or(conds ...gen.Condition) ITmpDo
	Select(conds ...field.Expr) ITmpDo
	Where(conds ...gen.Condition) ITmpDo
	Order(conds ...field.Expr) ITmpDo
	Distinct(cols ...field.Expr) ITmpDo
	Omit(cols ...field.Expr) ITmpDo
	Join(table schema.Tabler, on ...field.Expr) ITmpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITmpDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITmpDo
	Group(cols ...field.Expr) ITmpDo
	Having(conds ...gen.Condition) ITmpDo
	Limit(limit int) ITmpDo
	Offset(offset int) ITmpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITmpDo
	Unscoped() ITmpDo
	Create(values ...*model.Tmp) error
	CreateInBatches(values []*model.Tmp, batchSize int) error
	Save(values ...*model.Tmp) error
	First() (*model.Tmp, error)
	Take() (*model.Tmp, error)
	Last() (*model.Tmp, error)
	Find() ([]*model.Tmp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Tmp, err error)
	FindInBatches(result *[]*model.Tmp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Tmp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITmpDo
	Assign(attrs ...field.AssignExpr) ITmpDo
	Joins(fields ...field.RelationField) ITmpDo
	Preload(fields ...field.RelationField) ITmpDo
	FirstOrInit() (*model.Tmp, error)
	FirstOrCreate() (*model.Tmp, error)
	FindByPage(offset int, limit int) (result []*model.Tmp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITmpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tmpDo) Debug() ITmpDo {
	return t.withDO(t.DO.Debug())
}

func (t tmpDo) WithContext(ctx context.Context) ITmpDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tmpDo) ReadDB() ITmpDo {
	return t.Clauses(dbresolver.Read)
}

func (t tmpDo) WriteDB() ITmpDo {
	return t.Clauses(dbresolver.Write)
}

func (t tmpDo) Session(config *gorm.Session) ITmpDo {
	return t.withDO(t.DO.Session(config))
}

func (t tmpDo) Clauses(conds ...clause.Expression) ITmpDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tmpDo) Returning(value interface{}, columns ...string) ITmpDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tmpDo) Not(conds ...gen.Condition) ITmpDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tmpDo) Or(conds ...gen.Condition) ITmpDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tmpDo) Select(conds ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tmpDo) Where(conds ...gen.Condition) ITmpDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tmpDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITmpDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tmpDo) Order(conds ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tmpDo) Distinct(cols ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tmpDo) Omit(cols ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tmpDo) Join(table schema.Tabler, on ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tmpDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITmpDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tmpDo) RightJoin(table schema.Tabler, on ...field.Expr) ITmpDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tmpDo) Group(cols ...field.Expr) ITmpDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tmpDo) Having(conds ...gen.Condition) ITmpDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tmpDo) Limit(limit int) ITmpDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tmpDo) Offset(offset int) ITmpDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tmpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITmpDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tmpDo) Unscoped() ITmpDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tmpDo) Create(values ...*model.Tmp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tmpDo) CreateInBatches(values []*model.Tmp, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tmpDo) Save(values ...*model.Tmp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tmpDo) First() (*model.Tmp, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tmp), nil
	}
}

func (t tmpDo) Take() (*model.Tmp, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tmp), nil
	}
}

func (t tmpDo) Last() (*model.Tmp, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tmp), nil
	}
}

func (t tmpDo) Find() ([]*model.Tmp, error) {
	result, err := t.DO.Find()
	return result.([]*model.Tmp), err
}

func (t tmpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Tmp, err error) {
	buf := make([]*model.Tmp, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tmpDo) FindInBatches(result *[]*model.Tmp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tmpDo) Attrs(attrs ...field.AssignExpr) ITmpDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tmpDo) Assign(attrs ...field.AssignExpr) ITmpDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tmpDo) Joins(fields ...field.RelationField) ITmpDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tmpDo) Preload(fields ...field.RelationField) ITmpDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tmpDo) FirstOrInit() (*model.Tmp, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tmp), nil
	}
}

func (t tmpDo) FirstOrCreate() (*model.Tmp, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Tmp), nil
	}
}

func (t tmpDo) FindByPage(offset int, limit int) (result []*model.Tmp, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tmpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tmpDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tmpDo) Delete(models ...*model.Tmp) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tmpDo) withDO(do gen.Dao) *tmpDo {
	t.DO = *do.(*gen.DO)
	return t
}